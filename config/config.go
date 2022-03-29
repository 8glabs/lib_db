package config

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
