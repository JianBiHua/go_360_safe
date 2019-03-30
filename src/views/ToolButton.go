/**
 * Copyright (C), 2015-2019, 简笔画工作室
 * FileName: ToolButton.go
 * Author: 简笔画
 * Date: 2019.03.28 22:20
 * Description: 360安全卫士的工具按钮
 * History:
 * <author> <time> <version> <desc>
 * 作者姓名 修改时间 版本号 描述
 */
package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type ToolButton struct {
	// 父控件，用于显示ToolButton
	parent *widgets.QWidget

	// 显示文字
	Text string
	// 显示图片
	Icon string
	// 图片的宽高
	Icon_width int
	// 图片的高度
	Icon_height int

	// QToolButton控件
	widget *widgets.QToolButton

	// 一个用户标志位，区分是那个QToolButton被点击
	UserObject int
}

// 默认的款到单元
const WIDTH_UNIT = 50

// 析构函数，创建一个ToolButton
func NewToolButton(parent *widgets.QWidget, text, icon string, width, height int) *ToolButton {
	var tb = ToolButton{parent, text, icon, width, height, nil, -1}
	tb.init()
	return &tb
}

// 析构函数，创建一个ToolButton
func NewToolButton2(parent *widgets.QWidget, text, icon string) *ToolButton {
	var tb = ToolButton{parent, text, icon, WIDTH_UNIT, WIDTH_UNIT, nil, -1}
	tb.init()
	return &tb
}

// 析构函数，创建一个ToolButton
func NewToolButton3(widget *widgets.QToolButton, text, icon string, w, h int) *ToolButton {
	var tb = ToolButton{nil, text, icon, w, w, nil, -1}
	tb.widget = widget
	tb.init()
	return &tb
}

// init
// 初始化变量函数
func (tb *ToolButton) init() {
	tb.showWidget()
}

// showWidget
// 显示ToolButton
func (tb *ToolButton) showWidget() {
	if tb.widget == nil {
		tb.widget = widgets.NewQToolButton(tb.parent)
	}
	action := widgets.NewQAction(nil)
	action.SetIconText(tb.Text)
	action.SetIcon(gui.NewQIcon5(tb.Icon))
	//action.SetIconText()
	//
	tb.widget.SetAutoRaise(false)
	tb.widget.SetDefaultAction(action)
	tb.widget.SetIconSize(core.NewQSize2(tb.Icon_width, tb.Icon_height))
	tb.widget.SetToolButtonStyle(core.Qt__ToolButtonTextUnderIcon)
}

// Widget
// 返回QToolButton对象。
func (tb *ToolButton) Widget() *widgets.QToolButton {
	return tb.widget
}

// ConnectClicked
// 当点击时调用。
func (tb *ToolButton) ConnectClicked(fClicked func(tbutton *ToolButton)) {
	tb.widget.ConnectClicked(func(checked bool) {
		// 被点击时返回
		fClicked(tb)
	})
}

// SetUserObject
// 设置一个用户变量
func (tb *ToolButton) SetUserObject(flag int) {
	tb.UserObject = flag
}

// SetChecked
// 设置是否被选中，选中则为高亮状态。
//
// 解释一下：
// 本来是不需要这样去修改stylesheet的。
// 直接在qss里面加入: （意思是：page_selected属性为真是，使用该样式，否则使用默认样式 ）
// *[page_selected="true"] { /* italicize selected tabs */
//    background-color: qlineargradient(x1:0, y1:0, x2:0, y2:1,
//    stop:0 rgba(0, 0, 0, 0),
//    stop:1 rgba(0, 0, 0, 32));
// }
// 而代码中只需要:
// 	tb.widget.SetProperty("page_selected", core.NewQVariant11(true))
// 或者:
// tb.widget.SetProperty("page_selected", core.NewQVariant11(true))
// 但是似乎无法运行，不知道哪里不对。呵呵，只能先这样了.
func (tb *ToolButton) SetChecked(checked bool) {
	tb.widget.SetProperty("page_selected", core.NewQVariant11(true))
	if !checked {
		// 这个是默认背景色，
		tb.widget.SetStyleSheet(`border: 0px;
			border-radius: 0px;
			background-color: #00000000;
			color:rgb(255, 255, 255);
			padding-top: 5px;
			padding-bottom: 20px;`)
	} else {
		// 这个是一个渐变色的背景，点击时的效果。
		// 具体怎么用，自行百度吧
		// 解释下:
		// x1:0, y1:0, x2:0, y2:1 代表纵向绘图
		// stop:0 rgba(0, 0, 0, 0)  从透明色
		// stop:1 rgba(0, 0, 0, 32) 到0.32半透明色
		tb.widget.SetStyleSheet(`background-color: qlineargradient(x1:0, y1:0, x2:0, y2:1,
			stop:0 rgba(0, 0, 0, 0),
			stop:1 rgba(0, 0, 0, 32));`)
	}
}
