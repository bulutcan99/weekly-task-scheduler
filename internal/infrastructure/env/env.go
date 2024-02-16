package env

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type ENV struct {
	Host                      string `env:"HOST,required"`
	ServerPort                int    `env:"SERVER_PORT,required"`
	ServerReadTimeout         int    `env:"SERVER_READ_TIMEOUT,required"`
	DbPort                    int    `env:"DB_PORT,required"`
	DbName                    string `env:"DB_NAME,required"`
	ProviderCollection        string `env:"PROVIDER_COLLECTION,required"`
	DeveloperCollection       string `env:"DEVELOPER_COLLECTION,required"`
	TaskCollection            string `env:"TASK_COLLECTION,required"`
	MaxConnsPerHost           int    `env:"MAX_CONNS_PER_HOST,required"`
	MaxConnWaitTimeout        int    `env:"MAX_CONN_WAIT_TIMEOUT,required"`
	ReadTimeout               int    `env:"READ_TIMEOUT,required"`
	MaxIdempotentCallAttempts int    `env:"MAX_IDEMPOTENT_CALL_ATTEMPTS,required"`
	DevNumber                 int    `env:"DEV_NUMBER,required"`
}

var doOnce sync.Once
var Env ENV

func ParseEnv() *ENV {
	doOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("error loading .env file: %v", err)
			os.Exit(0)
		}
		if err := env.Parse(&Env); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(0)
		}
	})
	return &Env
}
