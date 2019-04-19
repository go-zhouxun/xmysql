package xmysql

type XMySQLConfig struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	DBName   string `json:"dbName"`
	User     string `json:"user"`
	Password string `json:"password"`
	MaxConn  int    `json:"maxConn"`
	MaxIdle  int    `json:"maxIdle"`
}
