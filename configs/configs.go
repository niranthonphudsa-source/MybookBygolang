package configs

type Configs struct {
	PostgresSQL PostgresSql
	App         Fiber
}

type Fiber struct {
	FiberHost string
	FiberPort string
}

type PostgresSql struct {
	Host     string
	Port     string
	Username string
	DBname   string
	Password string
	SSlmode  string
}
