package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func TestGetDBConf(t *testing.T) {

	conf := new(DBConf)
	conf.DB.DBName = "hoge"
	conf.DB.Password = "password"
	conf.DB.User = "User"
	conf.DB.Host = "8080"

	data, err := yaml.Marshal(conf)
	if err != nil {
		t.Fatal(err)
	}

	filename := "conf_file/testfile.yaml"
	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		t.Fatal(err)
	}

	t.Run("yamlで記述された特定のファイルを読み込むことが出来る。", func(t *testing.T) {

		srvConf = serverConf{
			EnvName:    dev.String(),
			DbConfPath: fmt.Sprint(filename),
		}

		dbconf := srvConf.GetDBConf()

		assert.Equal(t, conf.DB.DBName, dbconf.DB.DBName)
		assert.Equal(t, conf.DB.Password, dbconf.DB.Password)
		assert.Equal(t, conf.DB.User, dbconf.DB.User)
		assert.Equal(t, conf.DB.Host, dbconf.DB.Host)
	})

	// test で作成したファイルを削除
	os.Remove(filename)
}
