package configs

import "os"

func GetConnectServer() *PostgresSql {
	return &PostgresSql{
		Host:     os.Getenv("DB_Host"),
		Port:     os.Getenv("DB_Port"),
		DBname:   os.Getenv("DB_Name"),
		Password: os.Getenv("DB_Password"),
		Username: os.Getenv("DB_User"),
		SSlmode:  os.Getenv("SSlmode"),
	}
}
