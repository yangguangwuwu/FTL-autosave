package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/robfig/cron/v3"
	"image/color"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

//go:embed fonts/msyhbd.ttc
var FontBold []byte

//go:embed fonts/msyhl.ttc
var FontLight []byte

//go:embed fonts/msyh.ttc
var FontRegular []byte

type CnTheme struct {}

var _ fyne.Theme = (*CnTheme)(nil)

func (r CnTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (r CnTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (r CnTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Monospace {
		return fyne.NewStaticResource("regular.ttc",FontRegular)
	}
	if style.Bold {
		return fyne.NewStaticResource("bold.ttc",FontBold)
	}
	return fyne.NewStaticResource("regular.ttc",FontRegular)
}

func (r CnTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	go initCron()
	initKey()
	initWindow()
}

func initKey() {
}

func initCron()  {
	c := cron.New()

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

func initWindow()  {
	App := app.New()
	App.Settings().SetTheme(CnTheme{})

	Window := App.NewWindow("FLT存档管理")
	Window.Resize(fyne.NewSize(400, 300))

	ctrlTab := desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: desktop.ControlModifier}
	Window.Canvas().AddShortcut(&ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})


	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			log.Println("save ")
		}),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			log.Println("load ")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
	grid := container.NewBorder(
		toolbar,
		nil,
		nil,
		nil,
		container.NewGridWithColumns(2,AutoSaveLog(),SaveLog()),
	)
	Window.SetContent(grid)
	Window.ShowAndRun()
}

func AutoSaveLog() fyne.CanvasObject {
	files := GetFilelist("auto")
	log.Println(files)


	list := widget.NewListWithData(func() int {
		return len(files)
	}, func() fyne.CanvasObject {
		return container.NewBorder(nil,nil,nil,widget.NewButton("载入", func() {
		}),widget.NewLabel("test"))
	}, func(i binding.DataItem , o fyne.CanvasObject) {

	})

	return widget.NewCard("自动存档","",list)
}

func SaveLog() fyne.CanvasObject  {
	return widget.NewCard("手动存档","",nil)
}

func GetFilelist(path string) []map[string]interface{} {
	var files []map[string]interface{}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, map[string]interface{}{
			"time":f.ModTime().Unix(),
			"fileName":f.Name(),
			"path":path,
		})
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