package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/redis/go-redis/v9"

	"github.com/tmazitov/auth_service.git/docs"
	"github.com/tmazitov/auth_service.git/internal/config"
	"github.com/tmazitov/auth_service.git/internal/handlers"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/internal/storage"
	service "github.com/tmazitov/auth_service.git/pkg/service"
)

func setupDocs(config *config.Config) {
	docs.SwaggerInfo.Title = config.Docs.Title
	docs.SwaggerInfo.Description = config.Docs.Description
	docs.SwaggerInfo.Version = config.Service.Version
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s/%s/api", config.Service.Prefix, config.Service.Version)
}

func setupMetrics(service *service.Service) {
	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.SetSlowTime(10)
	metrics.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	metrics.Use(service.GetCore())
}

type ServiceFlags struct {
	ConfigPath string
	Storage    config.StorageConfig
	Redis      config.RedisConfig
}

func setupFlags() ServiceFlags {
	var flags ServiceFlags = ServiceFlags{}

	// Main config

	flag.StringVar(&flags.ConfigPath, "config", "./config.json", "Path to the service config.json")

	// DB flags

	flag.StringVar(&flags.Storage.Addr, "db_addr", "localhost:5432", "Address of the database")
	flag.StringVar(&flags.Storage.Database, "db_name", "auth_db", "Database name")
	flag.StringVar(&flags.Storage.User, "db_user", "auth_client", "Username to get access to the database")
	flag.StringVar(&flags.Storage.Password, "db_pass", "auth_client", "Password to get access to the database")
	flags.Storage.SSL = false

	// Redis flags

	flag.StringVar(&flags.Redis.Addr, "cache_addr", "localhost:6379", "Address for connection to the cache db ")
	flag.IntVar(&flags.Redis.DB, "cache_db", 0, "Number of the cache db ")

	flag.Parse()

	return flags
}

func main() {

	var (
		auth          *service.Service
		conf          *config.Config
		st            *staff.Staff
		flags         ServiceFlags
		storageClient *storage.Storage
		redisClient   *redis.Client
		err           error
	)

	flags = setupFlags()

	if conf, err = config.NewConfig(flags.ConfigPath, &flags.Storage, &flags.Redis); err != nil {
		panic(err)
	}

	fmt.Println(conf.Redis.Addr)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: "",
		DB:       conf.Redis.DB,
	})

	if storageClient, err = storage.NewStorage(conf.DB); err != nil {
		panic(err)
	}

	st = staff.NewStaff(conf)
	st.SetStorage(storageClient)
	st.SetJwt(redisClient, conf.Jwt.Secret)
	if err = st.SetConductor(redisClient, conf.Conductor); err != nil {
		panic(err)
	}

	auth = service.NewService(conf.Service)

	setupDocs(conf)
	setupMetrics(auth)
	auth.SetupMiddleware([]gin.HandlerFunc{
		gin.ErrorLogger(),
	})

	// register the `/metrics` route.
	auth.SetupDocs(handlers.ServiceDocs())
	auth.SetupHandlers(handlers.ServiceEndpoints(st))
	auth.Start()
}
