package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/satori/go.uuid" // 生成图片文件名
)

func FindImageAlt(alt string) (path string) {
	if len(alt) == 0 {
		return ""
	}
	index := UnicodeIndex(alt, '第')
	fmt.Printf("Index = %d\n", index)
	path = SubString(alt, 0, index-1)
	return path

}

//https://blog.csdn.net/wowzai/article/details/8941865
//begin
func UnicodeIndex(str string, substr rune) int {
	// 子串在字符串的字节位置
	result := strings.IndexRune(str, substr)

	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	//当 end < 0的时候，表示无法找到字符串，直接返回  “”
	if end < 0 {
		return ""
	}
	fmt.Printf("begin = %d\nend= %d\n", begin, end)
	// 返回子串
	return string(rs[begin:end])
}

//end

//判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 下载图片
func download(img_url string, alt string) int {
	var mutex *sync.Mutex
	mutex = new(sync.Mutex)
	mutex.Lock()

	path := FindImageAlt(alt)
	PathSep := ""

	//当目录不存在的时候，我们创建这个目录，用于存放套图
	if exists, _ := PathExists(path); exists != true {
		os.Mkdir(path, os.ModePerm)
	}

	if os.IsPathSeparator('\\') {
		PathSep = "\\"
	} else {
		PathSep = "/"
	}

	fmt.Printf("path = %s\n", path)
	uid, _ := uuid.NewV4()
	file_name := path + PathSep + uid.String() + ".jpg"
	fmt.Println(file_name)

	resp, _ := http.Get(img_url)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(file_name)
	io.Copy(out, bytes.NewReader(body))

	mutex.Unlock()

	return 0
}
