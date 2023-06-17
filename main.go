package main

import (
	"fmt"
	"os"

	"github.com/NIRVXSH/NIRVXSH-shop-project/config"
	"github.com/NIRVXSH/NIRVXSH-shop-project/pkg/databases"
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

	fmt.Println(db)
}
