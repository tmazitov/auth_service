package main

import (
	"github.com/gin-gonic/gin"
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

	if conf, err = config.NewConfig("config.json"); err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if storageClient, err = storage.NewStorage(conf.DB); err != nil {
		panic(err)
	}

	st = staff.NewStaff(conf)
	st.SetStorage(storageClient)
	st.SetJwt(redisClient, conf.JwtSecret)
	if err = st.SetConductor(redisClient, conf.Conductor); err != nil {
		panic(err)
	}

	auth = service.NewService("auth-service", "5001", "auth")

	docs.SwaggerInfo.Title = "Auth Service"
	docs.SwaggerInfo.Description = "This is a simple auth service"
	docs.SwaggerInfo.Version = "0.2"
	docs.SwaggerInfo.Host = "localhost:5001"
	docs.SwaggerInfo.BasePath = "/auth/v0/api"

	auth.SetupMiddleware([]gin.HandlerFunc{
		gin.ErrorLogger(),
	})

	auth.SetupDocs(handlers.ServiceDocs())
	auth.SetupHandlers(handlers.ServiceEndpoints(st))
	auth.Start()
}
