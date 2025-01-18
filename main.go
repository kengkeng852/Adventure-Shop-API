package main

import (
	"github.com/kengkeng852/adventure-shop-api/config"
	"github.com/kengkeng852/adventure-shop-api/databases"
	"github.com/kengkeng852/adventure-shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
