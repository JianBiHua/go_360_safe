// Package views WaveBallWidget
package views

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"math"
	"math/rand"
	"time"
)

// 定义个气泡
// s: 上升速度
type Bubble struct {
	x int
	y int
	w int
	h int

	s int
}

// BUBBLEMAX 气泡最大个数
const BUBBLEMAX = 8

// 气泡数组
var bubbleList [BUBBLEMAX]*Bubble

// WaveBallWidgetState is WaveBallWidget state
type WaveBallState int

const (
	// WaveBallWidgetState__Red 红色状态
	WaveBallState__Red = iota
	// WaveBallWidgetState__Yellow 黄色状态
	WaveBallState__Yellow
	// WaveBallWidgetState__Green 绿色状态
	WaveBallState__Green
)

// WaveBallWidgetState is WaveBallWidget state
type WaveBallMode int

const (
	// 波浪模式
	WaveBallMode__Wave = iota
	// 图片模式
	WaveBallMode__Icon
)

// WaveBallWidget is wave ball widget
type WaveBallWidget struct {
	*widgets.QWidget

	// 进度，从0到100
	progress int

	// 状态
	state WaveBallState

	// 外圈旋转角度
	angle float64

	// 显示模式
	mode WaveBallMode

	// 画笔
	painter *gui.QPainter

	// 父窗口起始坐标
	// 如果没有这个偏移坐标，则无法显示绘制的图形。
	pX int
	pY int

	// 图片路径
	icon string

	// 波浪前进距离
	waveOffset int

	// 波浪计时器
	waveTicker *time.Ticker
	// 气泡计时器
	BubbleTicker *time.Ticker
}

// NewWaveBallWidget is 构造函数
// w: 传入的Widget
func NewWaveBallWidget(w *widgets.QWidget, x, y int) *WaveBallWidget {
	var b = WaveBallWidget{widgets.NewQWidget(w, 0), 0, 0, 0, 0, nil, x, y, "", 0, nil, nil}
	b.init()
	b.startWave()
	//b.StartBubble()
	return &b
}

// init  WaveBallWidget init variables
// 初始化变量。
func (b *WaveBallWidget) init() {
	// 默认波形模式
	b.mode = WaveBallMode__Wave
	// 默认角度
	b.angle = 135

	// 添加Paint
	b.ConnectPaintEvent(b.onPaint)
}

// 启动波浪
func (b *WaveBallWidget) startWave() {
	b.waveTicker = time.NewTicker(50 * time.Millisecond)

	go func() {
		for {
			select {
			case <-b.waveTicker.C:

				if b.waveOffset > -b.Width() {
					b.waveOffset -= 3
				} else {
					b.waveOffset = (b.waveOffset % 3)
				}

				b.Update()
			}
		}
	}()
}

// SetMode 设置模式
func (b *WaveBallWidget) SetMode(m WaveBallMode) {
	b.mode = m
	// 重新绘制
	b.Update()
}

// SetProgress 设置进度
func (b *WaveBallWidget) SetProgress(p int) {
	b.progress = p
	// 重新绘制
	b.Update()
}

// 设置图标
func (b *WaveBallWidget) SetIcon(icon string) {
	// 波浪不用了, 关闭了
	b.waveTicker.Stop()
	b.StopBubble()

	b.progress = 0
	b.mode = WaveBallMode__Icon
	b.icon = icon
	b.Update()
}

// 开始气泡效果
func (b *WaveBallWidget) StartBubble() {
	b.BubbleTicker = time.NewTicker(50 * time.Millisecond)

	go func() {
		for {
			if b.BubbleTicker == nil {
				time.Sleep(1)
			} else {

				select {
				case <-b.BubbleTicker.C:

					// 外环转动
					b.angle += 5
					if b.angle >= 360 {
						b.angle = float64(int(b.angle) % 5)
					}

					// 检测并移动气泡。
					for i := 0; i < BUBBLEMAX; i++ {
						var bb = bubbleList[i]
						if bb == nil {
							// 创建气泡,
							// 随机位置，随机大小(5 到 (width/BUBBLEMAX-6))，随机速度
							var u = (b.Width() - 50) / (BUBBLEMAX)
							var row = rand.Intn(u)
							var width = 5 + rand.Intn(u-5)
							var speed = 1 + rand.Intn(4)

							// 左右下，各留出25个像素
							bubbleList[i] = &Bubble{row*width + width/2 + 25, b.Height() - 25 - width/2, width, width, speed}
							bb = bubbleList[i]
						}

						// 移动气泡，只是改变了y坐标
						bb.y -= bb.s

						y := float64(b.Height()) * float64((1 - float64(b.progress)/100.0))
						if float64(bb.y) < y {
							bubbleList[i] = nil
						}
					}

					b.Update()
				}
			}
		}
	}()
}

// 结束起气泡效果
func (b *WaveBallWidget) StopBubble() {
	if b.BubbleTicker != nil {
		b.BubbleTicker.Stop()
		b.BubbleTicker = nil
	}
	// 清除所有气泡
	for i := 0; i < BUBBLEMAX; i++ {
		bubbleList[i] = nil
	}
	// 外环复位
	b.angle = 135
	b.Update()
}

// Widget return WaveBallWidget widget
// 返回波浪球控件.
func (b *WaveBallWidget) Widget() *widgets.QWidget {
	return b.Window()
}

// 画图函数
func (b *WaveBallWidget) onPaint(event *gui.QPaintEvent) {
	//// 创建画笔
	var device = b.BackingStore().PaintDevice() //window.Painters[0]
	b.painter = gui.NewQPainter2(device)
	// 反走样
	b.painter.SetRenderHint(gui.QPainter__Antialiasing, true)

	// 画外圈
	b.drawRound()
	// 画完整圆，或者波浪半圆
	if b.mode == WaveBallMode__Icon || b.progress >= 90 {
		b.drawFullWave()
	} else {
		// 画两个，一前以后，后面半透明，前面不透明
		b.drawWave(true)
		b.drawWave(false)
	}
	// 画图片或者文字
	if b.mode == WaveBallMode__Wave {
		b.drawText()
	} else {
		b.drawIcon()
	}

	b.drawBubbles()

	// 结束绘制
	b.painter.End()
	b.painter.DestroyQPainter()
}

// 获取画笔颜色
func (b *WaveBallWidget) getColor() *gui.QColor {
	if b.progress < 30 {
		return gui.NewQColor3(0xff, 0x84, 0x38, 0xFF)
	} else if b.progress < 60 {
		return gui.NewQColor3(255, 127, 127, 255)
	} else {
		return gui.NewQColor3(0x31, 0xef, 0x1a, 0xFF)
	}
}

// 获取渐变色
func (b *WaveBallWidget) getLGColor(isHalf bool) *gui.QLinearGradient {
	var linear = gui.NewQLinearGradient3(float64(b.X()/2+b.pX), float64(b.Y()+b.pY),
		float64(b.X()/2+b.pX), float64(b.Y()+b.Height()+b.pY))
	//gui.NewQLinearGradient3(float64(b.Width()/2), float64(b.Height()), float64(b.Height()/2), 0)
	var alpha = 0xff
	if isHalf {
		alpha = 0x7f
	}
	if b.progress < 30 {
		linear.SetColorAt(0, gui.NewQColor3(0xFF, 0x9e, 0x7c, alpha))
		linear.SetColorAt(0.3, gui.NewQColor3(0xFF, 0x73, 0x56, alpha))
		linear.SetColorAt(1, gui.NewQColor3(0xFF, 0x68, 0x4e, alpha))
		//linear.SetColorAt(0, gui.NewQColor2(core.Qt__white))
		//linear.SetColorAt(1, gui.NewQColor3(0xa6, 0xce, 0x39, 255))
	} else if b.progress < 60 {
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

func (b *WaveBallWidget) getPoint(w float64) *core.QPointF {
	var x = float64(b.Width()/2) + w*math.Cos((360.0-b.angle)*3.14/180.0)
	var y = float64(b.Height()/2) + w*math.Sin((360.0-b.angle)*3.14/180.0)
	return core.NewQPointF3(x, y)
}

// 画波浪
func (b *WaveBallWidget) drawWave(isHalf bool) {
	b.painter.Save()

	// 设置颜色。
	b.painter.SetPen3(core.Qt__NoPen)
	b.painter.SetBrush(gui.NewQBrush10(b.getLGColor(isHalf)))

	sx := float64(b.X() + b.pX)
	sy := float64(b.Y() + b.pY)

	// 单元宽度
	u := float64(b.Height() / 2)
	// 根据进度，设置高度
	y := float64(2 * u * float64((1 - float64(b.progress)/100.0)))
	// 波形高度
	waveH := 15.0
	offset := float64(b.waveOffset)

	//
	path := gui.NewQPainterPath()
	path.MoveTo2(offset+sx, sy+y)
	if isHalf {
		path.QuadTo2(offset+u/2+sx, sy+waveH+y, offset+u+sx, sy+y)
		path.QuadTo2(offset+u/2+u+sx, sy-waveH+y, offset+2*u+sx, sy+y)
		path.QuadTo2(offset+u/2+u*2+sx, sy+waveH+y, offset+u*3+sx, sy+y)
		path.QuadTo2(offset+u/2+u*3+sx, sy-waveH+y, offset+4*u+sx, sy+y)
		path.QuadTo2(offset+u/2+u*4+sx, sy+waveH+y, offset+5*u+sx, sy+y)
	} else {
		path.QuadTo2(offset+u/2+sx, sy-waveH+y, offset+u+sx, sy+y)
		path.QuadTo2(offset+u/2+u+sx, sy+waveH+y, offset+2*u+sx, sy+y)
		path.QuadTo2(offset+u/2+u*2+sx, sy-waveH+y, offset+u*3+sx, sy+y)
		path.QuadTo2(offset+u/2+u*3+sx, sy+waveH+y, offset+4*u+sx, sy+y)
		path.QuadTo2(offset+u/2+u*4+sx, sy-waveH+y, offset+5*u+sx, sy+y)
	}
	path.LineTo2(offset+5*u+sx, 2*u+sy+y)
	path.LineTo2(sx, 2*u+sy+y)
	path.LineTo2(sx, sy+y)

	path2 := gui.NewQPainterPath()
	path2.AddEllipse2(6+sx, 6+sy, float64(b.Width()-12), float64(b.Height()-12))
	b.painter.SetClipPath(path, core.Qt__ReplaceClip)
	b.painter.DrawPath(path2)

	b.painter.Restore()
}

// 画满进度的波浪球，就是一个圆球
func (b *WaveBallWidget) drawFullWave() {

	b.painter.Save()

	// 设置颜色。
	b.painter.SetPen3(core.Qt__NoPen)
	b.painter.SetBrush(gui.NewQBrush10(b.getLGColor(false)))

	b.painter.DrawEllipse3(6+b.X()+b.pX, 6+b.Y()+b.pY, b.Width()-12, b.Height()-12)

	b.painter.Restore()
}

// 画外环
func (b *WaveBallWidget) drawRound() {
	b.painter.Save()

	// 设置颜色。
	b.painter.SetPen2(b.getColor())
	b.painter.SetBrush(gui.NewQBrush10(b.getLGColor(false)))

	// 画弧线
	b.painter.DrawArc(core.NewQRectF4(float64(2+b.X()+b.pX), float64(2+b.Y()+b.pY), float64(b.Width()-4), float64(b.Height()-4)), int(b.angle*16), 315*16)
	var point = b.getPoint(float64(b.Width()-4) / 2)
	b.painter.DrawEllipse3(int(point.X())-2+b.X()+b.pX, int(point.Y())+b.Y()-2+b.pY, 4, 4)

	b.painter.Restore()
}

// 画文字
func (b *WaveBallWidget) drawText() {
	b.painter.Save()

	// 设置
	b.painter.SetPen2(gui.NewQColor3(0xFF, 0xFF, 0xFF, 0xFF))
	// 设置字体
	textFont := gui.NewQFont()
	textFont.SetPixelSize(20)
	b.painter.SetFont(textFont)

	//

	var fh = float64(b.progress) / 100.0
	var max = int(float64(b.Height()) * fh)
	var top = float64(b.Height())*float64(1-fh) + float64((max-20)/2)

	// 这个20瞎写的，呵呵，偷偷懒
	b.painter.DrawText3(b.X()+b.Width()/2+b.pX-40, b.Y()+b.pY+int(top), fmt.Sprintf("%d分 点我", b.progress))

	textFont.DestroyQFont()

	b.painter.Restore()
}

// 画图片
func (b *WaveBallWidget) drawIcon() {
	b.painter.Save()

	// 将图片画在中心位置
	b.painter.DrawImage8(core.NewQPoint2(b.X()+b.Width()/2+b.pX-40, b.Y()+b.Height()/2+b.pY-40),
		gui.NewQImage9(b.icon, "png"))

	b.painter.Restore()
}

// 画气泡
func (b *WaveBallWidget) drawBubbles() {
	b.painter.Save()

	// 设置颜色。
	b.painter.SetPen3(core.Qt__NoPen)
	// 半透明
	b.painter.SetBrush(gui.NewQBrush3(gui.NewQColor3(0xFF, 0xFF, 0xFF, 0x5F), core.Qt__SolidPattern))

	// 画气泡
	for i := 0; i < BUBBLEMAX; i++ {
		var bb = bubbleList[i]
		if bb != nil {
			b.painter.DrawEllipse3(bb.x+b.pX+b.X(), bb.y+b.pY+b.Y(), bb.w, bb.h)
		}
	}

	b.painter.Restore()
}
