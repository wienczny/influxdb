package influxdb

import (
	"errors"
	"time"

	"github.com/influxdb/influxdb/meta"
)

const defaultReadTimeout = 5 * time.Second

var ErrTimeout = errors.New("timeout")

// Coordinator handle queries and writes across multiple local and remote
// data nodes.
type Coordinator struct {
	MetaStore meta.Store
}

// Write is coordinates multiple writes across local and remote data nodes
// according the request consistency level
func (c *Coordinator) Write(p *WritePointsRequest) error {

	databaseInfo, err := c.MetaStore.Database(p.Database)
	if err != nil {
		return err
	}
	_ = databaseInfo // deleteme

	retentionPolicyInfo, err := c.MetaStore.RetentionPolicy(p.Database, p.RetentionPolicy)
	if err != nil {
		return err
	}
	_ = retentionPolicyInfo // deleteme

	// TODO: normalize and validate our WritePointsRequest
	// eg: p.Normalize(databaseInfo.DefaultRetentionPolicy)

	// FIXME: use the consistency level specified by the WritePointsRequest
	pol := newConsistencyPolicyN(1) // p.ConsistencyLevel -- how do we get a ConsistencyPolicy from a ConsistencyLevel?

	// FIXME: build set of local and remote point writers
	var ws []PointsWriter

	///-----------------------

	type result struct {
		writerID int
		err      error
	}

	ch := make(chan result, len(ws))
	for i, w := range ws {
		go func(id int, w PointsWriter) {
			err := w.Write(p)
			ch <- result{id, err}
		}(i, w)
	}

	timeout := time.After(defaultReadTimeout)
	for range ws {
		select {
		case <-timeout:
			return ErrTimeout // return timeout error to caller
		case res := <-ch:
			if !pol.IsDone(res.writerID, res.err) {
				continue
			}
			if res.err != nil {
				return res.err
			}
			return nil
		}
	}

	panic("unreachable or bad policy impl")
}

func (c *Coordinator) Execute(q *QueryRequest) chan *Result {
	return nil
}

// remoteWriter is a PointWriter for a remote data node
type remoteWriter struct {
	//ShardInfo []ShardInfo
	//DataNodes DataNodes
}

func (w *remoteWriter) Write(p *WritePointsRequest) error {
	return nil
}
