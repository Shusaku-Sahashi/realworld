package main

import (
	"log"
	"os"

	"github.com/app/realworld/router"
)

func init() {
	// TODO: logパッケージとして切り出す。
	errorLog, err := os.OpenFile("Error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	log.SetOutput(errorLog)
	log.SetFlags(log.)
	log.SetPrefix("[error]:")
}

func main() {
	r := router.InitRouter()
	log.Panic("test")
	if err := r.Run(); err != nil {
		log.Panic(err)
	}
}
