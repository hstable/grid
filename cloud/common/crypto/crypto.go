package crypto

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
)

//用一个花里胡哨的加密方式来加密密码
func CryptoPwd(password string) string {
	shaed := sha512.Sum512_256([]byte(password))
	pwd := md5.Sum(shaed[:])
	return fmt.Sprintf("%x", pwd)
}

func HashWithSalt(raw []byte, salt string) string {
	a := sha512.New()
	a.Write(raw)
	return fmt.Sprintf("%x", a.Sum([]byte(salt)))
}
