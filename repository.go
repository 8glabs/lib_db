package lib_db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	DB *gorm.DB
}

type DBConfig struct {
	Type string `mapstructure:"type"`
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
	// Only some databases need this database name
	DBName        string `mapstructure:"dbName"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	EnableTLS     bool   `mapstructure:"enable_tls"`
	ServerName    string `mapstructure:"server_name"`
	ServerCert    string `mapstructure:"server_cert"`
	ClientCert    string `mapstructure:"client_cert"`
	ClientKey     string `mapstructure:"client_key"`
	RecreateTable bool   `mapstructure:"recreate_table"`
}

func NewRepo(conf *DBConfig) (*Repository, error) {
	host := conf.Host
	port := conf.Port
	user := conf.User
	password := conf.Password
	dbname := conf.DBName
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, dbname, port)
	repoDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	// TODO(weiduan): Check connection
	repo := Repository{
		DB: repoDb,
	}
	return &repo, nil
}

func NewDb(conf *DBConfig, db string) (*gorm.DB, error) {
	host := conf.Host
	port := conf.Port
	user := conf.User
	password := conf.Password
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, user, password, db, port)
	repoDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	return repoDb, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// Close
	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}
