package db

import "api-postgresql/configs"

func OpenConnection() {
	conf := configs.GetDB()
}
