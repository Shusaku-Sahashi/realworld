package main

import (
	"log"
	"os"

	"github.com/app/realworld/handler"

	"github.com/app/realworld/router"

	"github.com/app/realworld/db"

	"github.com/app/realworld/config"
)

func init() {
	// loggerの初期化
	setLogger()
}

func main() {
	// 各種設定情報の初期化
	config.InitializeConf()
	db := db.DBConn()
	defer db.Close()

	r := router.New()

	// usecaseを登録
	h := handler.Handler{}
	h.Register(r)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func setLogger() {
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
