package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

/**
 * Copyright (C), 2015-2019, 简笔画工作室
 * FileName: WidgetLog.go
 * Author: 简笔画
 * Date: 2019.03.28 22:20
 * Description: 盖章的Widget呀
 * History:
 * <author> <time> <version> <desc>
 * 作者姓名 修改时间 版本号 描述
 */
type WidgetLog struct {
	// 父类控件
	parent *widgets.QWidget
	// 显示的控件
	widget *widgets.QWidget
}

// NewWidgetLog
// 一样的自定义析构函数
func NewWidgetLog(p *widgets.QWidget) *WidgetLog {
	wl := WidgetLog{p, nil}
	wl.init()
	return &wl
}

// init
// 一样的初始化函数。
func (wl *WidgetLog) init() {
	wl.showWidgets()
}

// showWidgets
// 显示控件
// 为啥我要把这块定义成一个Widget，而不是一个一个创建并显示呢，
// 因为我要给他添加鼠标事件，用于弹出登陆对话框
func (wl *WidgetLog) showWidgets() {
	//
	wl.widget = widgets.NewQWidget(wl.parent, core.Qt__Widget)
	// 添加偶的章子
	nameLabel := widgets.NewQLabel2("简笔画", wl.widget, core.Qt__Widget)
	// 第三个参数是加粗用的，第四个斜体.
	nameLabel.SetFont(gui.NewQFont2("宋体", 30, 100, false))
	nameLabel.SetGeometry2(10, 0, 100, 30)
	nameLabel.SetObjectName("WidgetLog_Label")
	nameLabel.SetAlignment(core.Qt__AlignHCenter)
	// 添加登陆360账号
	loginLabel := widgets.NewQLabel2("登陆360账号", wl.widget, core.Qt__Widget)
	// 第三个参数是加粗用的，第四个斜体.
	loginLabel.SetFont(gui.NewQFont2("微软雅黑", 12, 100, false))
	loginLabel.SetGeometry2(10, 30, 100, 30)
	loginLabel.SetObjectName("WidgetLog_Label")
	// 水平居中
	loginLabel.SetAlignment(core.Qt__AlignHCenter)
	// 添加360图标
	iconWidget := widgets.NewQWidget(wl.widget, core.Qt__Widget)
	iconWidget.SetGeometry2(110, 0, 60, 60)
	iconWidget.SetStyleSheet("border-image: url(resources/icon_w.png)")
}

//
func (wl *WidgetLog) Widget() *widgets.QWidget {
	return wl.widget
}
