package Uitls

import (
	"crypto/md5"
	"github.com/astaxie/beego"
	"encoding/hex"
)

//MD5加密 +盐
func EnCode(str string) (result string) {
	md := md5.New()
	md.Write([]byte(str + beego.AppConfig.String("MD5Salt")))
	result = hex.EncodeToString(md.Sum(nil))
	return
}
