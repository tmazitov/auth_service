package main

import (
	"flag"

	"github.com/tmazitov/auth_service.git/internal/config"
	"github.com/tmazitov/auth_service.git/pkg/service"
)

type ServiceFlags struct {
	ConfigPath string
	Core       service.ServiceConfig
	Docs       config.DocsConfig
	Storage    config.StorageConfig
	Cache      config.RedisConfig
}

func setupConfig() (*config.Config, error) {
	var (
		flags ServiceFlags = ServiceFlags{}
		conf  *config.Config
		err   error
	)

	// Constant info

	flags.Core.Name = "auth_service"
	flags.Core.Version = "1.0.0"
	flags.Core.Prefix = "auth"
	flags.Docs.Title = "Auth service"
	flags.Docs.Description = "Service for user authorization and authentication"

	// Main config
	flag.IntVar(&flags.Core.Port, "port", 5000, "Port for the service")
	flag.StringVar(&flags.ConfigPath, "config", "./config.json", "Path to the service config.json")

	// DB flags

	flag.StringVar(&flags.Storage.Addr, "db_addr", "localhost:5432", "Address of the database")
	flag.StringVar(&flags.Storage.Database, "db_name", "auth_db", "Database name")
	flag.StringVar(&flags.Storage.User, "db_user", "auth_client", "Username to get access to the database")
	flag.StringVar(&flags.Storage.Password, "db_pass", "auth_client", "Password to get access to the database")
	flags.Storage.SSL = false

	// Redis flags

	flag.StringVar(&flags.Cache.Addr, "cache_addr", "localhost:6379", "Address for connection to the cache db ")
	flag.IntVar(&flags.Cache.DB, "cache_db", 0, "Number of the cache db ")

	flag.Parse()

	if conf, err = config.NewConfig(flags.ConfigPath); err != nil {
		return nil, err
	}

	conf.Service = &flags.Core
	conf.Docs = &flags.Docs
	conf.Storage = &flags.Storage
	conf.Redis = &flags.Cache

	return conf, nil
}
