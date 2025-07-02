package config

import (
	"flag"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cond "github.com/tmazitov/auth_service.git/pkg/conductor"
	"github.com/tmazitov/service"
)

type ServiceFlags struct {
	ConfigPath          string
	Mode                string
	Core                service.ServiceConfig
	Docs                DocsConfig
	Storage             StorageConfig
	Cache               RedisConfig
	AMQP                cond.AMQPConfig
	GRPCUserServiceAddr string
	JwtSecret           string
}

func Setup() (*Config, error) {
	var (
		flags ServiceFlags = ServiceFlags{}
		conf  *Config
		err   error
	)

	// Constant info

	flags.Core.Name = "auth_service"
	flags.Core.Version = "v0"
	flags.Core.Prefix = "auth"
	flags.Docs.Title = "Auth service"
	flags.Docs.Description = "Service for user authorization and authentication"

	// Main config

	flag.IntVar(&flags.Core.Port, "port", 5000, "Port for the service")
	flag.StringVar(&flags.ConfigPath, "config", "./json", "Path to the service json")
	flag.StringVar(&flags.Mode, "mode", "debug", "Service mode (release or debug)")
	flag.StringVar(&flags.GRPCUserServiceAddr, "grpc_user_service", "localhost:50021", "GRPC listener address for the user service")

	// DB flags

	flag.StringVar(&flags.Storage.Addr, "db_addr", "localhost:5432", "Address of the database")
	flag.StringVar(&flags.Storage.Database, "db_name", "auth_db", "Database name")
	flag.StringVar(&flags.Storage.User, "db_user", "auth_client", "Username to get access to the database")
	flag.StringVar(&flags.Storage.Password, "db_pass", "auth_client", "Password to get access to the database")
	flags.Storage.SSL = false

	// Redis flags

	flag.StringVar(&flags.Cache.Addr, "cache_addr", "localhost:6379", "Address for connection to the cache db ")
	flag.IntVar(&flags.Cache.DB, "cache_db", 0, "Number of the cache db ")

	// AMQP flags
	flag.StringVar(&flags.AMQP.Host, "amqp_host", "localhost", "AMQP host address")
	flag.IntVar(&flags.AMQP.Port, "amqp_port", 5672, "AMQP port")
	flag.StringVar(&flags.AMQP.User, "amqp_user", "guest", "AMQP username")
	flag.StringVar(&flags.AMQP.Pass, "amqp_pass", "guest", "AMQP password")

	// JWT flags
	flag.StringVar(&flags.JwtSecret, "jwt_secret", "supersecret", "JWT secret")

	flag.Parse()

	if conf, err = NewConfig(flags.ConfigPath); err != nil {
		return nil, err
	}

	conf.Jwt.Secret = flags.JwtSecret
	conf.Service = &flags.Core
	conf.Docs = &flags.Docs
	conf.Storage = &flags.Storage
	conf.Redis = &flags.Cache
	conf.Conductor.AMQPConfig = flags.AMQP
	conf.GRPC = &GRPCConfig{
		UserServiceAddress: flags.GRPCUserServiceAddr,
	}

	if flags.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	conf.CORS = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true, // If you want to allow credentials (cookies, etc.)
		MaxAge:           12 * time.Hour,
	}

	return conf, nil
}
