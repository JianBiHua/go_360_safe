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

const WIDTH = 920
const HEIGHT = 580

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
	mw.window.Show()
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

	mw.window = widgets.NewQMainWindow(nil, 0)
	// 设置标题
	mw.window.SetWindowTitle("360手机助手")
	// 设置宽高
	mw.window.SetMinimumSize2(WIDTH, HEIGHT)
	// 无法缩放
	mw.window.SetMaximumHeight(HEIGHT)
	mw.window.SetMaximumWidth(WIDTH)
	// 设置背景色
	mw.window.SetStyleSheet("background-color: #22CB64")
	// 去掉标题栏
	mw.window.SetWindowFlags(core.Qt__FramelessWindowHint)
	// 启动鼠标移动
	mw.window.SetMouseTracking(true)
	// 鼠标移动
	mw.window.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
		// 鼠标左键按下时处理
		if event.Button() == core.Qt__LeftButton {
			mw.window.Move2(mw.windowPos.X()+event.GlobalX()-mw.startPos.X(),
				mw.windowPos.Y()+event.GlobalY()-mw.startPos.Y())
		}
	})
	// 鼠标按下
	mw.window.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		// 左键按下时处理
		if event.Button() == core.Qt__LeftButton {
			// 保存坐标
			mw.startPos = event.GlobalPos()
			mw.windowPos = mw.window.Pos()
		}
	})

}

// ShowSubViews show all sub views
//
// 显示子界面
func (mw *MainWindow) showSubviews() {

	mw.showTopLayout()

}

// showTopLayout show top layout
// 显示顶部的布局
func (mw *MainWindow) showTopLayout() {

	widget := widgets.NewQWidget(mw.window, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	widget.SetGeometry2(0, 0, WIDTH, 30)

	// ==============右侧四个按钮， 从右往左=============
	// 关闭按钮
	close := widgets.NewQPushButton(widget)
	close.SetGeometry2(WIDTH-30, 5, 20, 20)
	// 设置这个有什么用呢? 为了让qss找到这个QPushButton
	close.SetObjectName("close")
	close.ConnectClicked(func(checked bool) {
		// 应用退出
		mw.app.Exit(0)
	})

	// 最小化按钮
	min := widgets.NewQPushButton(widget)
	min.SetGeometry2(WIDTH-60, 5, 20, 20)
	min.SetObjectName("min")
	min.ConnectClicked(func(checked bool) {
		// 应用最小化， 他喵的居然无效。
		mw.window.ShowMinimized()
	})

	// 菜单按钮
	menu := widgets.NewQPushButton(widget)
	menu.SetGeometry2(WIDTH-90, 5, 20, 20)
	menu.SetObjectName("menu")

	// 主题按钮
	theme := widgets.NewQPushButton(widget)
	theme.SetGeometry2(WIDTH-120, 5, 20, 20)
	theme.SetObjectName("theme")

	// 显示图标
	icon := widgets.NewQWidget(widget, core.Qt__Widget)
	icon.SetGeometry2(5, 5, 20, 20)
	icon.SetStyleSheet("border-image: url(resources/icon.png)")

	// 显示名字
	title := widgets.NewQLabel(widget, core.Qt__Widget)
	title.SetGeometry2(35, 5, 100, 20)
	title.SetText("360安全卫士11")
	title.SetStyleSheet("color: #FFFFFF")
}
