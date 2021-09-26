package config

import (
	"github.com/caarlos0/env"
	"go.uber.org/zap"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var Cfg Configuration

func init() {
	err := Cfg.LoadEnv()

	if Cfg.AppEnv == "development" {
		Cfg.Logger, _ = zap.NewDevelopment()
	} else {
		Cfg.Logger, _ = zap.NewProduction()
	}

	if Cfg.AppEnv == "test" {
		Cfg.ReadOnlyDatabaseURL = Cfg.TestReadOnlyDatabaseURL
		Cfg.DatabaseURL = Cfg.TestDatabaseURL
	}

	log := Cfg.Logger.Sugar().Named("config")

	if err != nil {
		log.Fatalw("Error loading env config",
			"error:", err)
	}
}

type Configuration struct {
	AppEnv string `env:"APP_ENV" envDefault:"development"`
	//AuthSecret string `env:"JWT_AUTH_SECRET" envDefault:"debug"`

	Host               string        `env:"HOST" envDefault:"localhost"`
	Port               string        `env:"PORT" envDefault:"8081"`
	ServerReadTimeout  time.Duration `env:"SERVER_READ_TIMEOUT" envDefault:"30s"`
	ServerWriteTimeout time.Duration `env:"SERVER_WRITE_TIMEOUT" envDefault:"30s"`
	ServiceName        string        `env:"SERVICE_NAME" envDefault:"Token Service"`

	// Database Configuration
	MaxIdleConn int `env:"MAX_IDLE_CONN" envDefault:"10"`
	MaxOpenConn int `env:"MAX_OPEN_CONN" envDefault:"20"`

	// TODO: Change these urls to the correct database
	DatabaseURL         string `env:"DATABASE_URL" envDefault:"postgres://localhost:5432/nft_service?sslmode=disable"` // TODO postgres://localhost:5432/account_service?sslmode=disable
	ReadOnlyDatabaseURL string `env:"READ_ONLY_DATABASE_URL" envDefault:""`                                            // TODO postgres://localhost:5432/account_service?sslmode=disable
	//RedisURL            string `env:"REDIS_URL" envDefault:"redis://localhost:6379"`                                                 // TODO
	OpenSeaURL string `env:"OPENSEA_URL" envDefault:"https://api.opensea.io"`

	TestDatabaseURL         string `env:"TEST_DATABASE_URL" envDefault:""`           // TODO postgres://localhost:5432/account_service_test?sslmode=disable
	TestReadOnlyDatabaseURL string `env:"TEST_READ_ONLY_DATABASE_URL" envDefault:""` // TODO postgres://localhost:5432/account_service_test?sslmode=disable
	// Cache Control Configuration
	ShortCacheSeconds uint `env:"SHORT_CACHE_SECONDS" envDefault:"60"`
	CacheSeconds      uint `env:"CACHE_SECONDS" envDefault:"300"`

	// Rate Limiting Configuration
	RateLimitWindow time.Duration `env:"RATE_LIMIT_WINDOW" envDefault:"60s"` // May use milliseconds
	RateLimitAllows int           `env:"RATE_LIMIT_ALLOWS" envDefault:"60"`

	// Cors Related Configuration
	CORSenabled    bool     `env:"CORS_ENABLED" envDefault:"true"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" `

	APIVersions []string `env:"API_VERSIONS" envDefault:"v1"` // separate by no space commas

	OpenSeaGetTimeout time.Duration `env:"OPENSEA_GET_TIMEOUT" envDefault:"15s"`

	// TODO: Setup an email
	//DomainEmail         string `env:"DOMAIN_EMAIL"`
	//DomainEmailPassword string `env:"DOMAIN_EMAIL_PASSWORD"`

	Logger *zap.Logger
}

func (c *Configuration) LoadEnv() error {
	return env.Parse(c)
}
