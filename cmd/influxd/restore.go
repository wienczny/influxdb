package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/influxdb/influxdb"
)

// RestoreCommand represents the program execution for "influxd restore".
type RestoreCommand struct {
	// The logger passed to the ticker during execution.
	Logger *log.Logger

	// Standard input/output, overridden for testing.
	Stderr io.Writer
}

// NewRestoreCommand returns a new instance of RestoreCommand with default settings.
func NewRestoreCommand() *RestoreCommand {
	cmd := RestoreCommand{
		Stderr: os.Stderr,
	}

	// Set up logger.
	cmd.Logger = log.New(cmd.Stderr, "", log.LstdFlags)
	return &cmd
}

// Run excutes the program.
func (cmd *RestoreCommand) Run(args ...string) error {

	cmd.Logger.Printf("influxdb restore, version %s, commit %s", version, commit)
	// Parse command line arguments.
	config, path, err := cmd.parseFlags(args)
	if err != nil {
		return err
	}

	return cmd.Restore(config, path)
}

func (cmd *RestoreCommand) Restore(config *Config, path string) error {
	// Remove broker & data directories.
	if err := os.RemoveAll(config.BrokerDir()); err != nil {
		return fmt.Errorf("remove broker dir: %s", err)
	} else if err := os.RemoveAll(config.DataDir()); err != nil {
		return fmt.Errorf("remove data dir: %s", err)
	}

	// Open snapshot file and all incremental backups.
	ssr, files, err := influxdb.OpenFileSnapshotsReader(path)
	if err != nil {
		return fmt.Errorf("open: %s", err)
	}
	defer closeAll(files)

	// Extract manifest.
	//ss, err := ssr.Snapshot()
	//if err != nil {
	//return fmt.Errorf("snapshot: %s", err)
	//}

	// Unpack snapshot files into data directory.
	if err := cmd.unpack(config.DataDir(), ssr); err != nil {
		return fmt.Errorf("unpack: %s", err)
	}

	// Notify user of completion.
	cmd.Logger.Printf("restore complete using %s", path)
	return nil
}

// parseFlags parses and validates the command line arguments.
func (cmd *RestoreCommand) parseFlags(args []string) (*Config, string, error) {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	configPath := fs.String("config", "", "")
	fs.SetOutput(cmd.Stderr)
	fs.Usage = cmd.printUsage
	if err := fs.Parse(args); err != nil {
		return nil, "", err
	}

	// Parse configuration file from disk.
	var config *Config
	var err error
	if *configPath == "" {
		config, err = NewTestConfig()
		log.Println("No config provided, using default settings")
	} else {
		config, err = ParseConfig(*configPath)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Require output path.
	path := fs.Arg(0)
	if path == "" {
		return nil, "", fmt.Errorf("snapshot path required")
	}

	return config, path, nil
}

func closeAll(a []io.Closer) {
	for _, c := range a {
		_ = c.Close()
	}
}

// unpack expands the files in the snapshot archive into a directory.
func (cmd *RestoreCommand) unpack(path string, ssr *influxdb.SnapshotsReader) error {
	// Create root directory.
	if err := os.MkdirAll(path, 0777); err != nil {
		return fmt.Errorf("mkdir: err=%s", err)
	}

	// Loop over files and extract.
	for {
		// Read entry header.
		sf, err := ssr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("next: entry=%s, err=%s", sf.Name, err)
		}

		// Log progress.
		cmd.Logger.Printf("unpacking: %s / idx=%d (%d bytes)", sf.Name, sf.Index, sf.Size)

		// Create parent directory for output file.
		if err := os.MkdirAll(filepath.Dir(filepath.Join(path, sf.Name)), 0777); err != nil {
			return fmt.Errorf("mkdir: entry=%s, err=%s", sf.Name, err)
		}

		if err := func() error {
			// Create output file.
			f, err := os.Create(filepath.Join(path, sf.Name))
			if err != nil {
				return fmt.Errorf("create: entry=%s, err=%s", sf.Name, err)
			}
			defer f.Close()

			// Copy contents from reader.
			if _, err := io.CopyN(f, ssr, sf.Size); err != nil {
				return fmt.Errorf("copy: entry=%s, err=%s", sf.Name, err)
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

// printUsage prints the usage message to STDERR.
func (cmd *RestoreCommand) printUsage() {
	fmt.Fprintf(cmd.Stderr, `usage: influxd restore [flags] PATH

restore uses a snapshot of a data node to rebuild a cluster.

        -config <path>
                          Set the path to the configuration file.
`)
}

// u64tob converts a uint64 into an 8-byte slice.
func u64tob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// btou64 converts an 8-byte slice into an uint64.
func btou64(b []byte) uint64 { return binary.BigEndian.Uint64(b) }
