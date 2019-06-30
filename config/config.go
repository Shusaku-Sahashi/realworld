package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

/*
環境変数の読み込みはyamlを使用した読み込みとした。
https://github.com/go-yaml/yaml
*/

// TODO: フレキシブルにするにはenvで指定されたconf.{環境名}.ymlファイルを読み込んだ方が良い気がする。
//go:generate stringer -type=envType
type envType int

const (
	dev envType = iota
	prod
)

type ServerConf struct {
	EnvName    string
	DbConfPath string
}

var SrvConf ServerConf
var DBConfig DBConf

type DBConf struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
	} `yaml:"db`
}

func InitializeConf() {
	var _env string

	flag.StringVar(&_env, "env", "dev", "profile type (dev, stg, prg). default is 'dev'")
	flag.Parse()

	// env flag が想定外の場合はfaitalを起こす。
	if _env != dev.String() && _env != prod.String() {
		log.Fatalf("env: %v is not supported.", _env)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("work dir is not obrained. ErrorMessage: %v", err)
	}

	// srvConfを初期化
	SrvConf = ServerConf{
		EnvName:    dev.String(),
		DbConfPath: filepath.Join(path, "config", "conf_file", "conf."+dev.String()+".yml"),
	}

	// dbの設定の初期化を行う。
	DBConfig = SrvConf.GetDBConf()

}

// TODO: 各ファイルの読み込み方は同じになる想定なのでこの読み出し方法を抽象化して、init時に全ての設定ファイルが読み込まれる様にする。
func (cf ServerConf) GetDBConf() DBConf {
	data, err := ioutil.ReadFile(cf.DbConfPath)
	if err != nil {
		log.Fatalf("db confing file can not be read. err: %v", err)
	}

	conf := new(DBConf)
	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		log.Fatalf("db confing file can not be read. err: %v", err)
	}

	log.Println("Read DB Config File.")

	return *conf
}
