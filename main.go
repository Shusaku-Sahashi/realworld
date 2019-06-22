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
	// 時刻と時刻のマイクロ秒、ディレクトリパスを含めたファイル名を出力
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.SetPrefix("[error]:")
}

func main() {
	r := router.InitRouter()
	if err := r.Run(); err != nil {
		log.Panic(err)
	}
}
