package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	go initCron()
	initKey()
	initWindow()
}

func initKey() {
}



func initWindow()  {
	App := app.New()
	App.Settings().SetTheme(CnTheme{})

	Window := App.NewWindow("FLT存档管理")
	Window.Resize(fyne.NewSize(750, 600))
	ctrlTab := desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: desktop.ControlModifier}
	Window.Canvas().AddShortcut(&ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})


	toolbar := widget.NewToolbar(
		//widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
		//	log.Println("save ")
		//}),
		//widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
		//	log.Println("load ")
		//}),
		//widget.NewToolbarSpacer(),
		//widget.NewToolbarAction(theme.HelpIcon(), func() {
		//	log.Println("Display help")
		//	w := App.NewWindow("帮助")
		//	w.Resize(fyne.NewSize(700, 550))
		//	w.SetContent(widget.NewCard("帮助","使用说明",nil))
		//	w.Show()
		//}),
	)
	help := widget.NewLabel("首先，退回游戏主菜单。然后，在列表选择载入的存档，点击载入，最后在游戏主菜单继续游戏！")
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("功能",theme.HomeIcon(), container.NewGridWithColumns(1,AutoSaveLog()),),
		container.NewTabItemWithIcon("帮助",theme.HelpIcon(), widget.NewCard("帮助","帮助说明",help)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	grid := container.NewBorder(
		toolbar,
		nil,
		nil,
		nil,
		tabs,
	)


	Window.SetContent(grid)
	Window.ShowAndRun()
}

func AutoSaveLog() fyne.CanvasObject {

	data := binding.BindStringList(&AutoFilesName)
	AutoFilesListName = data
	list := widget.NewListWithData(data,func() fyne.CanvasObject {
		return container.NewBorder(nil,nil,nil,widget.NewButtonWithIcon("载入",theme.ContentUndoIcon(), func() {
		}),widget.NewLabel("test"))
	},func(i binding.DataItem, o fyne.CanvasObject) {
		o.(*fyne.Container).Objects[0].(*widget.Label).Bind(i.(binding.String))
		o.(*fyne.Container).Objects[1].(*widget.Button).OnTapped = func() {
			filename,_ := i.(binding.String).Get()
			LoadAutoSave(filename)
		}
	})
	return widget.NewCard("自动存档","",list)
}

func SaveLog() fyne.CanvasObject  {
	return widget.NewCard("手动存档","",nil)
}

