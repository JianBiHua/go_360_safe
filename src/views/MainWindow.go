/**
 * Copyright (C), 2015-2019, 简笔画工作室
 * FileName: MainWindow.go
 * Author: 简笔画
 * Date: 2019.03.26 21:06
 * Description: 360安全卫士，主窗口文件。
 * History:
 * <author> <time> <version> <desc>
 * 作者姓名 修改时间 版本号 描述
 */
package views

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"os"
)

// MainWindow is a main window
// 显示主界面。
type MainWindow struct {
	//
	app *widgets.QApplication
	// 窗体
	window *widgets.QMainWindow

	// 起始点击坐标
	startPos  *core.QPoint
	windowPos *core.QPoint
}

// 模拟构造函数
// 一个全局函数
func NewMainWindow() *MainWindow {
	wm := &MainWindow{}
	wm.init()
	return wm
}

// Show is show main windows
// 让应用程序进入循环。
// is public func
func (mw *MainWindow) Show() {
	mw.showMainWindow()
	mw.showSubviews()
}

// Exec is application into loop
//
// 让应用程序进入循环。
func (mw *MainWindow) Exec() {
	mw.app.Exec()
}

// Init is init variables
// 初始化一些变量等
func (mw *MainWindow) init() {

}

// Show is to show main Window
// 显示主窗口
func (mw *MainWindow) showMainWindow() {

	mw.app = widgets.NewQApplication(len(os.Args), os.Args)
	// 设置qss，相当于html中的css文件
	data, err := ioutil.ReadFile("qss/stylesheet.qss")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	mw.app.SetStyleSheet(string(data))

	window := widgets.NewQMainWindow(nil, 0)
	// 设置标题
	window.SetWindowTitle("360手机助手")
	// 设置宽高
	window.SetMinimumSize2(920, 580)
	// 无法缩放
	window.SetMaximumHeight(580)
	window.SetMaximumWidth(920)
	// 设置背景色
	window.SetStyleSheet("background-color: #F5F5F5")
	// 去掉标题栏
	window.SetWindowFlags(core.Qt__FramelessWindowHint)
	// 启动鼠标移动
	window.SetMouseTracking(true)
	// 鼠标移动
	window.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
		// 鼠标左键按下时处理
		if event.Button() == core.Qt__LeftButton {
			window.Move2(mw.windowPos.X()+event.GlobalX()-mw.startPos.X(),
				mw.windowPos.Y()+event.GlobalY()-mw.startPos.Y())
		}
	})
	// 鼠标按下
	window.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		// 左键按下时处理
		if event.Button() == core.Qt__LeftButton {
			// 保存坐标
			mw.startPos = event.GlobalPos()
			mw.windowPos = window.Pos()
		}
	})

	window.Show()
}

// ShowSubViews show all sub views
//
// 显示子界面
func (mw *MainWindow) showSubviews() {

}
