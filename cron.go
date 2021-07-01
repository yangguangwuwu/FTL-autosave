package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func initFilesDate()  {
	GetFilelist("auto")
	AutoFilesName = AutoFilesName[0:0]
	for _,item := range Autofiles{
		AutoFilesName = append(AutoFilesName, item["fileName"].(string))
	}
	log.Printf("文件数据同步 ！")
}

func initCron()  {
	initFilesDate()
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		initFilesDate()
		AutoFilesListName.Reload()
	})
	c.AddFunc("*/1 * * * *", func() { //每分钟
		dir := "auto"
		dirName := dir + "/" + GetDirName()
		if IsPathExists(dirName) == false {
			if IsPathExists(dir) == false {
				os.Mkdir(dir, os.ModePerm)
			}
			os.Mkdir(dirName, os.ModePerm)
		}
		timeDate := GetTimeDate("%d-%02d-%02d-%02d-%02d-%02d")
		_, err := CopyFile("./continue.sav", "./"+dirName+"/continue.sav."+timeDate)
		if err == nil {
			timeDate = GetTimeDate("%d-%02d-%02d %02d:%02d:%02d")
			log.Printf("FTL自动存档成功！%s ！", timeDate)
		}
	})

	c.Start()
}
