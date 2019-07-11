// showHtml project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	. "github.com/soekchl/myUtils"
)

type TEMP struct {
	imageStr string
	fileName string
}

var (
	allImagesList []TEMP
)

func main() {
	dir, err := ioutil.ReadDir("./")
	if err != nil {
		Error(err)
		return
	}

	allImagesList = make([]TEMP, len(dir))

	i := 1 // 偏差值
	if len(dir) > 0 {
		i = getStrInt(dir[0].Name())
	}
	for _, v := range dir {
		if v.IsDir() {
			allImagesList[getStrInt(v.Name())-i].imageStr = createHtml(v.Name())
			allImagesList[getStrInt(v.Name())-i].fileName = v.Name() + ".html"
		}
	}

	h1 := ""
	h2 := ""

	for i, v := range allImagesList {
		if len(v.imageStr) < 1 {
			continue
		}
		if i == 0 {
			h1 = v.fileName
		} else {
			h1 = allImagesList[i-1].fileName
		}
		if i == len(allImagesList)-1 {
			h2 = v.fileName
		} else {
			h2 = allImagesList[i+1].fileName
		}

		outFile(v.fileName, v.imageStr, h1, h2)
	}
}

func createHtml(fileName string) (outFileStr string) {
	defer func() {
		if err := recover(); err != nil {
			Error(fileName, " Err=", err)
			time.Sleep(time.Second * 3)
		}
	}()

	dir, err := ioutil.ReadDir(fileName)
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
		images_str[getStrInt(v.Name())-i] = fmt.Sprintf(imgModel, "./"+fileName+"/"+v.Name())
	}

	outFileStr = strings.Join(images_str, "")

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

func outFile(fileName, outFileStr, pre, next string) {
	outFileStr = fmt.Sprintf(htmlModel, htmlHead, pre, next, outFileStr)
	fi, err := os.Create(fileName)
	if err != nil {
		Error(err)
		return
	}
	fi.WriteString(outFileStr)
	fi.Close()
}
