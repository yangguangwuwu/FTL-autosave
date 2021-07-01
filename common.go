package main

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Autofiles []map[string]interface{}
var AutoFilesName []string
var AutoFilesListName binding.ExternalStringList

func LoadAutoSave(filename string)  {
	log.Printf("载入存档 %s",filename)
	last := strings.LastIndex(filename, ".")
	dirStr := filename[last+1:last+14]
	lasti := strings.LastIndex(dirStr, "-")
	end := dirStr[lasti:]
	endNew := strings.Replace(end,"-","_",1)
	dirStr = strings.Replace(dirStr,end,endNew,1)
	dirStr = "auto/" + dirStr + "/"
	path := dirStr + filename
	//检测FLT    保存 退回主菜单

	// 重命名
	_,err := CopyFile(path, "continue.sav")
	if err != nil {
		log.Println(err.Error())
	}
}

func GetFilelist(path string) {
	Autofiles = Autofiles[0:0]
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		Autofiles = append(Autofiles, map[string]interface{}{
			"time":f.ModTime().Unix(),
			"fileName":f.Name(),
			"path":path,
		})
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

}

//文件是否存在方法定义
func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true

	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func GetDirName() string {
	now := time.Now() //获取当前时间
	//fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	return fmt.Sprintf("%d-%02d-%02d_%02d", year, month, day, hour)
}

func GetTimeDate(format string) string {
	now := time.Now() //获取当前时间
	//fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	return fmt.Sprintf(format, year, month, day, hour, minute, second)
}

func CopyFile(srcFileName string, dstFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)

	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer srcFile.Close()


	os.Remove(dstFileName)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)

}
