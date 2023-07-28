package checksum

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func GenerateMd5CheckSum(file multipart.File) string {
	md := md5.New()
	io.Copy(md, file)
	return fmt.Sprintf("%x", md.Sum(nil))

}
