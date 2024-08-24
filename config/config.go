package config

import (
	"context"

	"github.com/Sirpyerre/payment-platform/pkg/logger"

	"github.com/sethvargo/go-envconfig"
)

type Configuration struct {
	Environment  string `env:"ENVIRONMENT, default=local"`
	Server       Server
	DBConfig     DBConfig
	SecretKey    string `env:"SECRET_KEY,required"`
	BankProvider PaymentBank
}

type Server struct {
	Environment string `env:"GO_ENV"`
	LogLevel    string `env:"LOG_LEVEL,default=DEBUG"`
	Port        int    `env:"PORT,default=8080"`
	Version     string `env:"VERSION_API,default=0.0.1"`
}

// DBConfig :
type DBConfig struct {
	Server          string `env:"POSTGRES_SERVER,required"`
	Database        string `env:"POSTGRES_DATABASE,required"`
	Port            int    `env:"POSTGRES_PORT,default=5432"`
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	ConnectTimeOut  int    `env:"POSTGRES_CONNECT_TIMEOUT,default=15"`
	MaxOpenConns    int    `env:"POSTGRES_MAX_OPEN_CONNS,default=30"`
	MaxIdleConns    int    `env:"POSTGRES_MAX_IDLE_CONNS,default=25"`
	ConnMaxLifetime int    `env:"POSTGRES_CONN_MAX_LIFETIME,default=30"`
	QueryTimeout    int    `env:"POSTGRES_QUERY_TIMEOUT,default=60"`
}

type PaymentBank struct {
	URL     string `env:"PAYMENT_BANK_URL"`
	Timeout int    `env:"PAYMENT_BANK_TIMEOUT,default=30"`
}

func NewConfiguration() *Configuration {
	configuration := new(Configuration)
	readConfigEnv(configuration)
	return configuration
}

func readConfigEnv(configuration *Configuration) {
	ctx := context.Background()
	if err := envconfig.Process(ctx, configuration); err != nil {
		logger.GetLogger().FatalIfError("config", "readConfigEnv", err)
	}
}
