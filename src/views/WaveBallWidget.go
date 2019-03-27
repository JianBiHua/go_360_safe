/**
 * Copyright (C), 2015-2019, 简笔画工作室
 * FileName: MainWindow.go
 * Author: 简笔画
 * Date: 2019.03.27 12:32
 * Description: 波浪球。
 * History:
 * <author> <time> <version> <desc>
 * 作者姓名 修改时间 版本号 描述
 */
package views

import "github.com/therecipe/qt/widgets"

type WaveBallWidget struct {
	widget *widgets.QWidget
}

//
func NewWaveBallWidget() *WaveBallWidget {
	var b = WaveBallWidget{}
	b.init()
	return &b
}

//
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

func (b *WaveBallWidget) Show() {
}
