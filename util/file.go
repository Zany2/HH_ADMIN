package util

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"math/rand"
	"os"
	"strings"
	"time"
)

var r = []byte("0123456789abcdefghijklmnopgrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomStringFileName 随机生成文件名
func RandomStringFileName(n int) string {
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(200)))
	for i := 0; i < n; i++ {
		result = append(result, r[rand.Intn(len(r))])
	}
	return string(result)
}

// GetRandomStringNameNew 随机生成文件名
func GetRandomStringNameNew(filePreFix string, mdEncrypt bool) string {
	nameString := filePreFix + grand.S(26, false) + gconv.String(gtime.Now().Unix()) + grand.Digits(26)
	if mdEncrypt {
		encrypt, _ := gmd5.Encrypt(nameString)
		return encrypt
	} else {
		return nameString
	}
}

// PathExistOrCreat 文件夹是否存在/不存在则创建
func PathExistOrCreat(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}

// JudgeType 判断文件限制类型
func JudgeFileType(srcFileName string, typeList []interface{}) (result bool) {
	result = true
	if srcFileName == "" || len(typeList) <= 0 {
		return result
	}
	for _, typeName := range typeList {
		if strings.Split(srcFileName, ".")[len(strings.Split(srcFileName, "."))-1] == gconv.String(typeName) {
			return false
		}
	}
	return
}
