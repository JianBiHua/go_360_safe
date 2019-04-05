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
	"math"
	"os"
)

// WIDTH is window max and min width
const WIDTH = 920

// HEIGHT is window max and min height
const HEIGHT = 580

// PAGECOUNT is max page count
const PAGECOUNT = 8

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

	// 登陆弹出框
	loginPopupWidget *PopupWidget

	//
	loginWidget *WidgetLog
}

// NewMainWindow 模拟构造函数
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
	mw.loginPopupWidget = nil
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
		fmt.Println("ConnectMouseMoveEvent")
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

	mw.window.ConnectEvent(func(event *core.QEvent) bool {

		if event.Type() == core.QEvent__HoverMove {
			// 获取当前鼠标在屏幕上的绝对坐标
			var pos = core.NewQPoint2(gui.QCursor_Pos().X(), gui.QCursor_Pos().Y())
			// 将全局左边转换成window上的坐标
			var pos2 = mw.window.MapFromGlobal(pos)
			var x = pos2.X()
			var y = pos2.Y()

			// 如果是鼠标滑动
			// 处理鼠标的滑动
			// 实现的逻辑是，
			// 当鼠标移动到loginPopupWidget上或者移动到登陆区域，才显示登陆框，否则，如果显示了则隐藏。
			if mw.loginPopupWidget == nil {
				// 如果loginPopupWidget不存在，移动到登陆区域，才显示登陆框
				if mw.loginWidget.Widget().Geometry().Contains3(x, y) {
					mw.loginPopupWidget = NewPopupWidget(mw.window, ArrowDirectionUp, 90, 700, 90)
					mw.loginPopupWidget.SetGeometry3(700, 90, 180, 100)
					mw.loginPopupWidget.Show(1 * 1000)
					mw.loginPopupWidget.SetWindowFlags(core.Qt__WindowStaysOnTopHint)
					mw.loginPopupWidget.ShowNormal()
				}
			} else {
				// 如果显示了。鼠标移动到除登录区以及loginPopupWidget区域外时，隐藏loginPopupWidget
				if !mw.loginWidget.Widget().Geometry().Contains3(x, y) &&
					!mw.loginPopupWidget.Geometry().Contains3(x, y) {

					mw.loginPopupWidget.Hide2(1*1000, func(widget *PopupWidget) {
						if mw.loginPopupWidget != nil {
							mw.loginPopupWidget.HideDefault()
							//当隐藏后销毁
							mw.loginPopupWidget.DestroyQWidget()
							mw.loginPopupWidget = nil
						}
					})
				}
			}
		}

		// 如果没有这句，你会发现波浪不动了， 因为你把所有的事件都处理了，下面收不到消息了
		return mw.window.EventDefault(event)
	})
}

// ShowSubViews show all sub views
// 显示子界面
func (mw *MainWindow) showSubviews() {

	mw.showTopLayout()
	mw.showToolBarsLayout()
	mw.showStackedWidget()

	//mw.loginPopupWidget = NewPopupWidget (mw.window, ArrowDirectionUp, 100, 700, 90)
	//mw.loginPopupWidget.SetGeometry2(700, 90, 180, 100)
	//mw.loginPopupWidget.Show(300)
	//mw.loginPopupWidget.SetWindowFlags(core.Qt__WindowStaysOnTopHint);
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
var toolButtonsIconNames = [...]string{
	"resources/tab/tab_home.png",
	"resources/tab/tab_horse.png",
	"resources/tab/tab_clear.png",
	"resources/tab/tab_fix.png",
	"resources/tab/tab_optimize.png",
	"resources/tab/tab_actions.png",
	"resources/tab/tab_market.png",
	"resources/tab/tab_app.png"}

// 名称
var toolButtonsNames = [...]string{
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
	ToolButtonTypeHome = iota
	// 木马扫描
	ToolButtonTypeHorse
	// 电脑清除
	ToolButtonTypeClear
	// 系统修复
	ToolButtonTypeFix
	// 功能大全
	ToolButtonTypeActions
	// 优化加速
	ToolButtonTypeYouhua
	// 360商城
	ToolButtonTypeMarket
	// 软件管家
	ToolButtonTypeApp
)

// showToolBarsLayout show toolbars layout
// 显示工具按钮
func (mw *MainWindow) showToolBarsLayout() {

	widget := widgets.NewQWidget(mw.window, 0)
	// 设置一个QVboxLayout布局
	widget.SetLayout(widgets.NewQVBoxLayout())
	widget.SetGeometry2(0, 30, WIDTH, 90)

	// 添加8个ToolButton
	for i := 0; i < PAGECOUNT; i++ {
		// 创建ToolButton
		toolButtons[i] = NewToolButton2(widget, toolButtonsNames[i],
			toolButtonsIconNames[i])
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
	mw.loginWidget = NewWidgetLog(widget)
	mw.loginWidget.Widget().SetGeometry2(740, 15, 300, 90)
}

// onToolButtonDidClicked
// ToolButton中的QToolButton控件被点击时调用。
func (mw *MainWindow) onToolButtonDidClicked(button *ToolButton) {
	// 看这个UserObject，用来区别到底是那个控件被点击了
	fmt.Println("UserObject", button.UserObject)
	// 只有i==UserObject，说明就是那个控件被点击了
	// 为啥减2呢, 因为360安全卫士就是这么设计的
	// 后面两个得特殊处理的。
	if button.UserObject < PAGECOUNT-2 {
		for i := 0; i < PAGECOUNT-2; i++ {
			toolButtons[i].SetChecked(i == button.UserObject)
		}

		// 改变当前Page
		mw.stackWidget.SetCurrentIndex(button.UserObject)
	}
}

// 获取画笔颜色
func getColor(progress int) *gui.QColor {
	if progress < 30 {
		return gui.NewQColor3(0xff, 0x84, 0x38, 0xFF)
	} else if progress < 60 {
		return gui.NewQColor3(255, 127, 127, 255)
	} else {
		return gui.NewQColor3(0x31, 0xef, 0x1a, 0xFF)
	}
}

// 获取渐变色
func getLGColor(isHalf bool, widget *widgets.QWidget, progress int) *gui.QLinearGradient {
	var linear = gui.NewQLinearGradient3(float64(widget.Width()/2), float64(widget.Height()), float64(widget.Height()/2), 0)
	var alpha = 0xff
	if isHalf {
		alpha = 0x7f
	}
	if progress < 30 {
		linear.SetColorAt(0, gui.NewQColor3(0xFF, 0x9e, 0x7c, alpha))
		linear.SetColorAt(0.3, gui.NewQColor3(0xFF, 0x73, 0x56, alpha))
		linear.SetColorAt(1, gui.NewQColor3(0xFF, 0x68, 0x4e, alpha))
	} else if progress < 60 {
		linear.SetColorAt(0, gui.NewQColor3(0xFF, 0xc1, 0x73, alpha))
		linear.SetColorAt(0.2, gui.NewQColor3(0xFF, 0xae, 0x5c, alpha))
		linear.SetColorAt(0.8, gui.NewQColor3(0xFF, 0x8a, 0x40, alpha))
		linear.SetColorAt(1, gui.NewQColor3(0xF8, 0x93, 0x4A, alpha))
	} else {
		linear.SetColorAt(0, gui.NewQColor3(0x7c, 0xfb, 0x7c, alpha))
		linear.SetColorAt(0.2, gui.NewQColor3(0x53, 0xe9, 0x6e, alpha))
		linear.SetColorAt(0.8, gui.NewQColor3(0x35, 0xe3, 0x85, alpha))
		linear.SetColorAt(1, gui.NewQColor3(0x72, 0xec, 0xa7, alpha))
	}

	return linear
}

func getPoint(w float64, widget *widgets.QWidget, angle float64) *core.QPointF {
	var x = float64(widget.Width()/2) + w*math.Cos((360.0-angle)*3.14/180.0)
	var y = float64(widget.Height()/2) + w*math.Sin((360.0-angle)*3.14/180.0)
	return core.NewQPointF3(x, y)
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
