/*
 * Copyright (C), 2015-2019, 简笔画工作室
 * FileName: MainWindow.go
 * Author: 简笔画
 * Date: 2019.03.26 21:06
 * Description: 360安全卫士，主窗口文件。
 * History:
 * <author> <time> <version> <desc>
 * 作者姓名 修改时间 版本号 描述
 */

// Package views is all user define view struct, Contains one or more QWidgets pointer
package views

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"io/ioutil"
	"os"
)

// WIDTH is window max and min width
const WIDTH = 920

// HEIGHT is window max and min height
const HEIGHT = 580

// PAGE_COUNT is max page count
const PAGE_COUNT = 8

// MainWindow is a main window
// 显示主界面。
type MainWindow struct {
	//
	app *widgets.QApplication
	// 窗体
	window *widgets.QMainWindow
	//
	stackWidget *widgets.QStackedWidget

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
	mw.showToolBarsLayout()
	mw.showStackedWidget()

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

// 用一用GO的数组喽
var toolButtons [8]*ToolButton

// 图片名称
var toolButtons_icon_name = [...]string{
	"resources/tab/tab_home.png",
	"resources/tab/tab_horse.png",
	"resources/tab/tab_clear.png",
	"resources/tab/tab_fix.png",
	"resources/tab/tab_optimize.png",
	"resources/tab/tab_actions.png",
	"resources/tab/tab_market.png",
	"resources/tab/tab_app.png"}

// 名称
var toolButtons_name = [...]string{
	"电脑体验",
	"木马查杀",
	"电脑清理",
	"系统修复",
	"优化加速",
	"功能大全",
	"电360商城",
	"软件管家"}

// 在用用GO的枚举吧，哈哈
const (
	// 电脑体验
	ToolButtonType_HOME = iota
	// 木马扫描
	ToolButtonType_HORSE
	// 电脑清除
	ToolButtonType_CLEAR
	// 系统修复
	ToolButtonType_FIX
	// 功能大全
	ToolButtonType_ACTIONS
	// 优化加速
	ToolButtonType_YOUHUA
	// 360商城
	ToolButtonType_MARKET
	// 软件管家
	ToolButtonType_APP
)

// showToolBarsLayout show toolbars layout
// 显示工具按钮
func (mw *MainWindow) showToolBarsLayout() {

	widget := widgets.NewQWidget(mw.window, 0)
	// 设置一个QVboxLayout布局
	widget.SetLayout(widgets.NewQVBoxLayout())
	widget.SetGeometry2(0, 30, WIDTH, 90)

	// 添加8个ToolButton
	for i := 0; i < PAGE_COUNT; i++ {
		// 创建ToolButton
		toolButtons[i] = NewToolButton2(widget, toolButtons_name[i],
			toolButtons_icon_name[i])
		// 设置显示位置
		toolButtons[i].Widget().SetGeometry2(20+90*i, 0, 90, 90)
		// 设置点击状态
		toolButtons[i].SetChecked(i == 0)
		// 按下时执行
		toolButtons[i].ConnectClicked(mw.onToolButtonDidClicked)
		// 设置一个用户标识
		toolButtons[i].SetUserObject(i)
	}

	// 显示我的章子,
	// 名人作画都会盖个章，则也学学名人，盖个章呗 呵呵。
	wl := NewWidgetLog(widget)
	wl.Widget().SetGeometry2(740, 15, 300, 90)
}

// onToolButtonDidClicked
// ToolButton中的QToolButton控件被点击时调用。
func (mw *MainWindow) onToolButtonDidClicked(button *ToolButton) {
	// 看这个UserObject，用来区别到底是那个控件被点击了
	fmt.Println("UserObject", button.UserObject)
	// 只有i==UserObject，说明就是那个控件被点击了
	// 为啥减2呢, 因为360安全卫士就是这么设计的
	// 后面两个得特殊处理的。
	if button.UserObject < PAGE_COUNT-2 {
		for i := 0; i < PAGE_COUNT-2; i++ {
			toolButtons[i].SetChecked(i == button.UserObject)
		}

		// 改变当前Page
		mw.stackWidget.SetCurrentIndex(button.UserObject)
	}
}

// showStackedWidget
// 显示stackedWidget
func (mw *MainWindow) showStackedWidget() {
	mw.stackWidget = widgets.NewQStackedWidget(mw.window)
	mw.stackWidget.SetGeometry2(0, 120, WIDTH, HEIGHT-120)
	mw.stackWidget.SetCurrentIndex(0)
	mw.stackWidget.SetStyleSheet("background: #FFFFFF")

	// 接口的初次应用
	var widgetInterface WidgetInterface

	// 将PageHome，断言转换为接口
	// 因为go没有继承接口的概念，所以这里实现使用转换来处理的
	widgetInterface = NewPageHome()
	//创建Page,并显示。
	mw.stackWidget.AddWidget(widgetInterface.Widget())

	// 将HorseHome，断言转换为接口
	// 因为go没有继承接口的概念，所以这里实现使用转换来处理的
	widgetInterface = NewPageHorse()
	//创建Page,并显示。
	mw.stackWidget.AddWidget(widgetInterface.Widget())

	widgetInterface = NewPageClear()
	mw.stackWidget.AddWidget(widgetInterface.Widget())

	widgetInterface = NewPageFix()
	//创建Page,并显示。
	mw.stackWidget.AddWidget(widgetInterface.Widget())

	widgetInterface = NewPageOptimize()
	//创建Page,并显示。
	mw.stackWidget.AddWidget(widgetInterface.Widget())

	widgetInterface = NewPageActions()
	//创建Page,并显示。
	mw.stackWidget.AddWidget(widgetInterface.Widget())
}
