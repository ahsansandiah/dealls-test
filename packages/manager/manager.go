package manager

import (
	"context"
	"fmt"
	"os"

	jwtAuth "github.com/ahsansandiah/dealls-test/packages/auth/jwt"
	middlewareAuth "github.com/ahsansandiah/dealls-test/packages/auth/middleware"
	httpClient "github.com/ahsansandiah/dealls-test/packages/client/http"
	"github.com/ahsansandiah/dealls-test/packages/config"
	"github.com/ahsansandiah/dealls-test/packages/json"
	logger "github.com/ahsansandiah/dealls-test/packages/log"
	"github.com/ahsansandiah/dealls-test/packages/server"
	gormDatabase "github.com/ahsansandiah/dealls-test/packages/storage/gorm"
	cache "github.com/ahsansandiah/dealls-test/packages/storage/redis"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Manager interface {
	GetConfig() *config.Config
	GetServer() *server.Server
	GetGorm() *gorm.DB
	GetCache() *redis.Client
	GetJwt() jwtAuth.Jwt
	GetMiddleware() middlewareAuth.Middleware
	GetLog() logger.Log
	GetJson() json.Json
	GetHttp() httpClient.Http
}

type manager struct {
	config         *config.Config
	server         *server.Server
	db             *gorm.DB
	cache          *redis.Client
	jwtAuth        jwtAuth.Jwt
	middlewareAuth middlewareAuth.Middleware
	logger         logger.Log
	json           json.Json
	httpClient     httpClient.Http
}

func NewInit() (Manager, error) {
	lg := logger.NewLog()
	ctx := context.Background()

	fmt.Println(os.Getenv("APP_ENV"))
	cfg, err := config.NewConfig()
	if err != nil {
		lg.ErrorLog(ctx, err)
		return nil, err
	}

	srv := server.NewServer(cfg)

	jwt := jwtAuth.NewJwt(cfg)

	database, err := gormDatabase.NewGorm(cfg).Connect()
	if err != nil {
		lg.ErrorLog(ctx, err)
		return nil, err
	}

	clHttp := httpClient.NewHttp(cfg, lg)
	clHttp.Connect()

	json := json.NewJson(lg)

	cache, err := cache.NewRedis(cfg).Connect()
	if err != nil {
		lg.ErrorLog(ctx, err)
		return nil, err
	}

	middleware := middlewareAuth.NewMiddleware(cfg, cache, lg, json)

	return &manager{
		config:         cfg,
		server:         srv,
		db:             database,
		cache:          cache,
		jwtAuth:        jwt,
		middlewareAuth: middleware,
		logger:         lg,
		httpClient:     clHttp,
		json:           json,
	}, nil
}

func (sm *manager) GetConfig() *config.Config {
	return sm.config
}

func (sm *manager) GetServer() *server.Server {
	return sm.server
}

func (sm *manager) GetGorm() *gorm.DB {
	return sm.db
}

func (sm *manager) GetCache() *redis.Client {
	return sm.cache
}

func (sm *manager) GetJwt() jwtAuth.Jwt {
	return sm.jwtAuth
}

func (sm *manager) GetMiddleware() middlewareAuth.Middleware {
	return sm.middlewareAuth
}

func (sm *manager) GetLog() logger.Log {
	return sm.logger
}

func (sm *manager) GetJson() json.Json {
	return sm.json
}

func (sm *manager) GetHttp() httpClient.Http {
	return sm.httpClient
}
