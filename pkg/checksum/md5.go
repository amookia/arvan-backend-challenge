package checksum

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func GenerateFileMd5CheckSum(file multipart.File) string {
	md := md5.New()
	io.Copy(md, file)
	return fmt.Sprintf("%x", md.Sum(nil))

}

func GenerateByteMd5CheckSum(body []byte) string {
	md := md5.New()
	md.Write(body)
	return fmt.Sprintf("%x", md.Sum(nil))

}