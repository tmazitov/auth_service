package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/tmazitov/auth_service.git/docs"
	"github.com/tmazitov/auth_service.git/internal/config"
	"github.com/tmazitov/auth_service.git/internal/handlers"
	"github.com/tmazitov/auth_service.git/internal/staff"
	"github.com/tmazitov/auth_service.git/internal/storage"
	"github.com/tmazitov/service"
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

	if conf, err = config.Setup(); err != nil {
		log.Fatalf("Failed to setup config: %v", err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: "",
		DB:       conf.Redis.DB,
	})
	defer redisClient.Close()

	if storageClient, err = storage.NewStorage(conf.Storage); err != nil {
		log.Fatalf("Failed to setup storage: %v", err)
	}
	defer storageClient.Close()

	userServiceConn, err := grpc.Dial(conf.GRPC.UserServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer userServiceConn.Close()

	st = staff.NewStaff(userServiceConn, conf)
	st.SetStorage(storageClient)
	st.SetJwt(redisClient, conf.Jwt.Secret)
	if err = st.SetConductor(redisClient, conf.Conductor); err != nil {
		log.Fatalf("Failed to setup conductor: %v", err)
	}

	auth = service.NewService(conf.Service)
	auth.GetCore().Use(cors.New(conf.CORS))
	setupDocs(conf)
	setupMetrics(auth)
	auth.SetupMiddleware([]gin.HandlerFunc{
		gin.ErrorLogger(),
	})

	// register the `/metrics` route.
	docs := handlers.ServiceDocs()[0]
	auth.GetCore().Handle(docs.Method, docs.Path, docs.Handler.AfterMiddleware()...)
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
