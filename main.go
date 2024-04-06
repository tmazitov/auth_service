package main

import (
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

func main() {

	var (
		auth          *service.Service
		conf          *config.Config
		st            *staff.Staff
		storageClient *storage.Storage
		redisClient   *redis.Client
		err           error
	)

	if conf, err = setupConfig(); err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: "",
		DB:       conf.Redis.DB,
	})

	if storageClient, err = storage.NewStorage(conf.Storage); err != nil {
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
