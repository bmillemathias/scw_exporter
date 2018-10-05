package config

import (
	"time"
)

// Server defines the general server configuration.
type Server struct {
	Addr string
	Path string
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string
	Pretty bool
}

// Target defines the target specific configuration.
type Target struct {
	Token   string
	Org     string
	Region  string
	Timeout time.Duration
}

// Collector defines the collector specific configuration.
type Collector struct {
	Dashboard      bool
	SecurityGroups bool
	Servers        bool
	Snapshots      bool
	Volumes        bool
}

// Config is a combination of all available configurations.
type Config struct {
	Server    Server
	Logs      Logs
	Target    Target
	Collector Collector
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
