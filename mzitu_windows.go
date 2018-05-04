package main


import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"github.com/satori/go.uuid"    // 生成图片文件名
	"net/http"
)
// 下载图片
func download(img_url string) int {
    uid, _ := uuid.NewV4()
    file_name := uid.String() + ".jpg"
    fmt.Println(file_name)

    resp, _ := http.Get(img_url)
    body, _ := ioutil.ReadAll(resp.Body)
    out, _ := os.Create(file_name)
    io.Copy(out, bytes.NewReader(body))

    return 0
}