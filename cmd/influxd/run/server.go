package run

import (
	"fmt"
	"time"

	"github.com/influxdb/influxdb/cluster"
	"github.com/influxdb/influxdb/meta"
	"github.com/influxdb/influxdb/services/admin"
	"github.com/influxdb/influxdb/services/collectd"
	"github.com/influxdb/influxdb/services/graphite"
	"github.com/influxdb/influxdb/services/httpd"
	"github.com/influxdb/influxdb/services/opentsdb"
	"github.com/influxdb/influxdb/services/udp"
	"github.com/influxdb/influxdb/tsdb"
)

// Server represents a container for the metadata and storage data and services.
// It is built using a Config and it manages the startup and shutdown of all
// services in the proper order.
type Server struct {
	MetaStore     *meta.Store
	TSDBStore     *tsdb.Store
	QueryExecutor *tsdb.QueryExecutor
	PointsWriter  *cluster.PointsWriter
	ShardWriter   *cluster.ShardWriter

	Services []Service
}

// NewServer returns a new instance of Server built from a config.
func NewServer(c *Config) *Server {
	// Construct base meta store and data store.
	s := &Server{
		MetaStore: meta.NewStore(c.Meta),
		TSDBStore: tsdb.NewStore(c.Data.Dir),
	}

	// Initialize query executor.
	s.QueryExecutor = tsdb.NewQueryExecutor(s.TSDBStore)
	s.QueryExecutor.MetaStore = s.MetaStore
	s.QueryExecutor.MetaStatementExecutor = &meta.StatementExecutor{Store: s.MetaStore}

	// Set the shard writer
	s.ShardWriter = cluster.NewShardWriter(time.Duration(c.Cluster.ShardWriterTimeout))

	// Initialize points writer.
	s.PointsWriter = cluster.NewPointsWriter()
	s.PointsWriter.MetaStore = s.MetaStore
	s.PointsWriter.TSDBStore = s.TSDBStore
	s.PointsWriter.ShardWriter = s.ShardWriter

	// Append services.
	s.appendClusterService(c.Cluster)
	s.appendAdminService(c.Admin)
	s.appendHTTPDService(c.HTTPD)
	s.appendCollectdService(c.Collectd)
	s.appendOpenTSDBService(c.OpenTSDB)
	s.appendUDPService(c.UDP)
	for _, g := range c.Graphites {
		s.appendGraphiteService(g)
	}

	return s
}

func (s *Server) appendClusterService(c cluster.Config) {
	srv := cluster.NewService(c)
	srv.TSDBStore = s.TSDBStore
	s.Services = append(s.Services, srv)
}

func (s *Server) appendAdminService(c admin.Config) {
	if !c.Enabled {
		return
	}
	srv := admin.NewService(c)
	s.Services = append(s.Services, srv)
}

func (s *Server) appendHTTPDService(c httpd.Config) {
	if !c.Enabled {
		return
	}
	srv := httpd.NewService(c)
	srv.Handler.MetaStore = s.MetaStore
	srv.Handler.QueryExecutor = s.QueryExecutor
	srv.Handler.PointsWriter = s.PointsWriter
	s.Services = append(s.Services, srv)
}

func (s *Server) appendCollectdService(c collectd.Config) {
	if !c.Enabled {
		return
	}
	srv := collectd.NewService(c)
	s.Services = append(s.Services, srv)
}

func (s *Server) appendOpenTSDBService(c opentsdb.Config) {
	if !c.Enabled {
		return
	}
	srv := opentsdb.NewService(c)
	s.Services = append(s.Services, srv)
}

func (s *Server) appendGraphiteService(c graphite.Config) {
	if !c.Enabled {
		return
	}
	srv := graphite.NewService(c)
	s.Services = append(s.Services, srv)
}

func (s *Server) appendUDPService(c udp.Config) {
	if !c.Enabled {
		return
	}
	srv := udp.NewService(c)
	srv.Server.PointsWriter = s.PointsWriter
	s.Services = append(s.Services, srv)
}

// Open opens the meta and data store and all services.
func (s *Server) Open() error {
	if err := func() error {
		// Open meta store.
		if err := s.MetaStore.Open(); err != nil {
			return fmt.Errorf("open meta store: %s", err)
		}

		// Wait for the store to initialize.
		<-s.MetaStore.Ready()

		// Open TSDB store.
		if err := s.TSDBStore.Open(); err != nil {
			return fmt.Errorf("open tsdb store: %s", err)
		}

		for _, service := range s.Services {
			if err := service.Open(); err != nil {
				return fmt.Errorf("open service: %s", err)
			}
		}

		return nil

	}(); err != nil {
		s.Close()
		return err
	}

	return nil
}

// Close shuts down the meta and data stores and all services.
func (s *Server) Close() error {
	if s.MetaStore != nil {
		s.MetaStore.Close()
	}
	if s.TSDBStore != nil {
		s.TSDBStore.Close()
	}
	for _, service := range s.Services {
		service.Close()
	}
	return nil
}

// Service represents a service attached to the server.
type Service interface {
	Open() error
	Close() error
}
