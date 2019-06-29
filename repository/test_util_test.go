package repositories

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)



/*
gormのcolum情報からdbのテーブル名称を全て取得する。
tagの取得に関しては下記を参照。
https://hiyosi.tumblr.com/post/100922038678/go-%E3%81%A7-struct-%E3%81%AE%E3%82%BF%E3%82%B0%E6%83%85%E5%A0%B1%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B
*/
func GetDBFields(m interface{}) []string {
	tp := reflect.TypeOf(m)
	var fields []string

	for i := 0; i < tp.NumField(); i++ {
		tag := tp.Field(i).Tag.Get("gorm")
		if strings.Contains(tag, "column") {

			fls := strings.Split(tag, ";")
			for _, fl := range fls {
				if strings.HasPrefix(fl, "column") {
					fields = append(fields, strings.Split(fl, ":")[1])
				}
			}
		}
	}

	return fields
}

// testUtilのテスト
func TestGetDBFields (t *testing.T) {
	
	type testUser struct {
		ID       uint   `gorm:"column:id;primary_key"`
		Username string `gorm:"column:user_name"`
		Email    string `gorm:"column:email"`
		Password string `gorm:"column:password"`
		Bio      string `gorm:"column:bio"`
	}

	fields := GetDBFields(testUser{})		

	assert.Equal(t, "id", fields[0])
	assert.Equal(t, "user_name", fields[1])
	assert.Equal(t, "email", fields[2])
	assert.Equal(t, "password", fields[3])
	assert.Equal(t, "bio", fields[4])
}