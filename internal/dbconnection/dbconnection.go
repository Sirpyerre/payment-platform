package dbconnection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/Sirpyerre/payment-platform/config"
	"github.com/Sirpyerre/payment-platform/pkg/logger"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Connector struct {
	DB *sqlx.DB
}

// NewDBConnection :
func NewDBConnection(cfg *config.Configuration) *Connector {
	connection := &Connector{}
	db, _ := connection.setDBConnection(cfg)

	// ping
	err := db.Ping()
	defer logger.GetLogger().Debugf("dbconnect:%s", "ping database")
	if err != nil {
		logger.GetLogger().FatalIfError("dbconnection", "NewDBConnection", err)
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Minute * time.Duration(180))

	connection.DB = db

	return connection
}

func (c *Connector) setDBConnection(cfg *config.Configuration) (*sqlx.DB, error) {
	connectionString := createDNS(cfg)
	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		logger.GetLogger().FatalIfError("dbconnection", "setDBConnection", err)
		return nil, err
	}
	return db, nil

}

func createDNS(cfg *config.Configuration) string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=%d",
		cfg.DBConfig.Server,
		cfg.DBConfig.Port,
		cfg.DBConfig.Database,
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.ConnectTimeOut,
	)
}
