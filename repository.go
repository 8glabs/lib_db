package lib_db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	repoDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// TODO(weiduan): Check connection
	repo := Repository{
		DB: repoDb,
	}
	return &repo, nil
}
