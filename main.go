package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"github.com/PatrikOlin/neo-burton/api"
	"github.com/PatrikOlin/neo-burton/db"
)

var (
	// version string
	addr string
)

func init() {
	pflag.StringVarP(&addr, "address", "a", ":4040", "the address for the api to listen on. Host and port separated by ':'")
	pflag.Parse()

	_, err := db.Open()
	if err != nil {
		log.Fatalln("failed to load db .env")
	}
	fmt.Println("hallå?")
}

func main() {
	// configure logger
	log, _ := zap.NewProduction(zap.WithCaller(false))
	defer func() {
		_ = log.Sync()
	}()

	r := api.GetRouter(log)

	http.ListenAndServe(addr, r)
}
