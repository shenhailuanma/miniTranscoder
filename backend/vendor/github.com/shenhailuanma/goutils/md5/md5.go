package md5


import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Md5String(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Md5Bytes(data []byte) string {
	ctx := md5.New()
	ctx.Write(data)
	return hex.EncodeToString(ctx.Sum(nil))
}