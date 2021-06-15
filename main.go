package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	c := cron.New()

	c.AddFunc("*/1 * * * *", func() { //每分钟
		dirName := GetDirName()
		if IsPathExists(dirName) == false {
			os.Mkdir(dirName, os.ModePerm)
		}
		timeDate := GetTimeDate("%d-%02d-%02d-%02d-%02d-%02d")
		_, err := CopyFile("./continue.sav", "./"+dirName+"/continue.sav."+timeDate)
		if err == nil {
			timeDate = GetTimeDate("%d-%02d-%02d %02d:%02d:%02d")
			log.Printf("FTL自动存档成功！%s ！", timeDate)
		}
		//files := GetFilelist(".")
		//fmt.Println(files)
	})

	c.Start()

	for {
		time.Sleep(time.Second)
	}
}

func GetFilelist(path string) []string {
	var files []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		//println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return files

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

	//打开dstFileName

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)

}
