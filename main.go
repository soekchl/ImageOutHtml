// ImageOutHtml project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/soekchl/myUtils"
)

const (
	HTML  = `<!DOCTYPE html><HTML><body><div align="center"> %s </div></body></HTML>`
	IMAGE = `<img alt="" src="%s" style="vertical-align:bottom"><br/>`
	HREF2 = `<a href="%s"> 下一章</a>`
	HREF1 = `<a href="%s">上一章 </a>`
)

type TEMP struct {
	image_str string
	file_name string
}

var (
	all_images_list []TEMP
)

func main() {
	show()
}

func show() {
	dir, err := ioutil.ReadDir(".")
	if err != nil {
		Error(err)
		return
	}

	all_images_list = make([]TEMP, len(dir))

	i := 1 // 偏差值
	if len(dir) > 0 {
		i = getStrInt(dir[0].Name())
	}

	for _, v := range dir {
		if v.IsDir() {
			all_images_list[getStrInt(v.Name())-i].image_str = createHtml(v.Name())
			all_images_list[getStrInt(v.Name())-i].file_name = v.Name() + ".html"
		}
	}

	h1 := ""
	h2 := ""

	for i, v := range all_images_list {
		if len(v.image_str) < 1 {
			continue
		}
		if i == 0 {
			h1 = v.file_name
		} else {
			h1 = all_images_list[i-1].file_name
		}
		if i == len(all_images_list)-1 {
			h2 = v.file_name
		} else {
			h2 = all_images_list[i+1].file_name
		}

		v.image_str += fmt.Sprintf(HREF1, h1)
		v.image_str += fmt.Sprintf(HREF2, h2)

		outFile(v.file_name, v.image_str)
	}
}
func createHtml(file_name string) (out_file_str string) {
	dir, err := ioutil.ReadDir(file_name)
	if err != nil {
		Error(err)
		return
	}

	images_str := make([]string, len(dir))

	i := 1 // 偏差值
	if len(dir) > 0 {
		i = getStrInt(dir[0].Name())
	}

	for _, v := range dir {
		images_str[getStrInt(v.Name())-i] = fmt.Sprintf(IMAGE, ".//"+file_name+"//"+v.Name())
	}

	out_file_str = strings.Join(images_str, "")

	return
}

func getStrInt(name string) (dest int) {
	f := true
	for i, _ := range name {
		if name[i] >= 48 && name[i] <= 57 {
			dest *= 10
			dest += charToInt(name[i])
			f = false
		} else if f {
			continue
		} else {
			break
		}
	}
	return dest
}

func charToInt(c byte) int {
	return int(c - 48)
}

func outFile(file_name, out_file_str string) {
	out_file_str = fmt.Sprintf(HTML, out_file_str)

	fi, err := os.Create(file_name)
	if err != nil {
		Error(err)
		return
	}
	fi.WriteString(out_file_str)
	fi.Close()
}
