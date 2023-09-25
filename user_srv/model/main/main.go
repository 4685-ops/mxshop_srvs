package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"io"
	"strings"
)

func genMd5(code string) string {

	hash := md5.New()
	_, _ = io.WriteString(hash, code)
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	//fmt.Println(genMd5("123456"))
	//dsn := "root:123456@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//
	//newLogger := logger.New(
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // 禁用彩色打印
	//	},
	//)
	//
	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true,
	//	},
	//	Logger: newLogger,
	//})
	//
	//// 迁移 schema
	//err := db.AutoMigrate(&model.User{}) //此处应该有建表语句 sql
	//if err != nil {
	//	return
	//}

	// 方式一：使用默认选项
	//salt, encodedPwd := password.Encode("generic password", nil)
	//fmt.Println(salt)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	//// 方式二：使用自定义选项
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	fmt.Println(check) // true

}
