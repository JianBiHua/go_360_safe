// Package views WaveBallWidget
package views

import "github.com/therecipe/qt/widgets"

// WaveBallWidget is wave ball widget
type WaveBallWidget struct {
	widget *widgets.QWidget
}

// NewWaveBallWidget is 构造函数
// w: 宽度
// h: 高度
func NewWaveBallWidget(w, h int) *WaveBallWidget {
	var b = WaveBallWidget{}
	b.init()
	return &b
}

// NewWaveBallWidget2 is 构造函数
// x:
// y:
// w: 宽度
// h: 高度
func NewWaveBallWidget2(x, y, w, h int) *WaveBallWidget {
	var b = WaveBallWidget{}
	b.init()
	return &b
}

// init  WaveBallWidget init variables
// 初始化变量。
func (b *WaveBallWidget) init() {
}

// Widget return WaveBallWidget widget
// 返回波浪球控件.
func (b *WaveBallWidget) Widget() *widgets.QWidget {
	return b.widget
}
