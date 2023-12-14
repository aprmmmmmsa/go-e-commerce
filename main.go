package main

import (
	"github/mtotheesa/go-e-commerce/config"
	"github/mtotheesa/go-e-commerce/modules/servers"
	"github/mtotheesa/go-e-commerce/pkg/databases"
	"os"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
