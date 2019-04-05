package views

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"time"
)

// 尖角的显示方向
type ArrowDirection int

const (
	ArrowDirectionUp = iota
	ArrowDirectionDown
	ArrowDirectionRight
	ArrowDirectionLeft
)

// 尖角的宽高
const ArrowWidth = 10.0

// 阴影宽度
const ShadowWidth = 2.0

type PopupWidget struct {
	*widgets.QWidget

	// 尖角的方向
	arrowDirection ArrowDirection

	// 尖角显示的位置. 距离左边或者上面的距离【理论上应该是0到width;
	// 但实际上要考虑箭头的宽高，
	// 所以这个值的范围应为 ArrowWidth/2 到 width-ArrowWidth/2;
	// 如果画成圆角矩形了，又得计算圆角的宽度】
	luDis float64

	// 当前的宽度， 用来做动画【0到1】
	currentProgress float64

	// 透明度、用来做动画的【0到1】
	currentAlpha float64

	// 当前是否是显示状态。
	isShow bool

	// 动画计数器
	ticker *time.Ticker

	//
	pX float64

	//
	pY float64

	//
	loginWidget *widgets.QWidget
}

// 构造函数
// p: 父类指针，添加到对象里
// d： 方向
func NewPopupWidget(p widgets.QWidget_ITF, d ArrowDirection, dis float64, rx float64, ry float64) *PopupWidget {
	// 默认不可见，为了做动画.
	pw := PopupWidget{widgets.NewQWidget(p, 0), d, dis, 0.0, 1.0, false, nil, rx, ry, nil}
	pw.init()
	return &pw
}

// 初始化函数
func (p *PopupWidget) init() {
	// 窗体背景去除
	//	p.SetAutoFillBackground(false);
	//
	p.SetStyleSheet("background-color:transparent")

	// 默认遮罩，什么也不显示
	p.SetMask2(gui.NewQRegion2(int(p.luDis),
		0,
		0,
		p.Height(),
		gui.QRegion__Rectangle))

	// 连接画图事件
	p.ConnectPaintEvent(p.onPaint)

	//
	p.showSubviews()

}

// 显示子控件，随便加的为了看下效果罢了
func (p *PopupWidget) showSubviews() {

	p.loginWidget = widgets.NewQWidget(p, 0)
	//p.loginWidget.HideDefault()
	p.loginWidget.SetGeometry2(ShadowWidth, ShadowWidth+ArrowWidth, p.Width()-ShadowWidth*2, p.Height()-ShadowWidth*2-ArrowWidth)

	userLabel := widgets.NewQLabel2("用户名:", p.loginWidget, 0)
	userLabel.SetGeometry2(10, 10, 60, 30)

	user := widgets.NewQTextEdit(p.loginWidget)
	user.SetGeometry2(70, 10, 70, 30)
	user.SetStyleSheet("background-color: #A0A0A0")

	btn := widgets.NewQPushButton(p.loginWidget)
	btn.SetGeometry2(0, 50, 100, 50)
	btn.SetText("登陆")
}

func (p *PopupWidget) SetGeometry3(x int, y int, w int, h int) {
	p.SetGeometry2(x, y, w, h)
	p.loginWidget.SetGeometry2(ShadowWidth, ShadowWidth+ArrowWidth, p.Width()-ShadowWidth*2, p.Height()-ShadowWidth*2-ArrowWidth)
}

// 显示
// interval：显示的持续时间, 1000 = 1s
func (p *PopupWidget) Show(interval int) {
	// 直接调用Show2, 代码统一
	p.Show2(interval, nil)
	p.Update()
}

// 显示
// interval：显示的持续时间, 1000 = 1s
// f: 动画结束时相应，可能某些人喜欢在动画结束后执行一些操作，就可以这么写
func (p *PopupWidget) Show2(interval int, f func(widget *PopupWidget)) {
	if !p.isShow {
		// 只有隐藏状态才显示
		p.ticker = time.NewTicker(20 * time.Millisecond)

		// 默认遮罩，什么也不显示
		p.SetMask2(gui.NewQRegion2(int(p.luDis),
			0,
			0,
			p.Height(),
			gui.QRegion__Rectangle))

		p.loginWidget.ShowDefault()

		go func() {
			defer p.ticker.Stop()

			var ticker = 0
			for {
				select {
				case <-p.ticker.C:
					// 1.0 / (interval/20)
					p.currentProgress += 20.0 / float64(interval)
					if p.currentProgress > 1 {
						p.currentProgress = 1
						// 都跑到了，就直接让它结束吧，
						ticker = interval
					}

					// 修改遮罩位置
					// 添加一个遮罩层，我们的遮罩应该以箭头为中心
					// 这样就能算出实际应该遮罩的位置，改变currentProgress， 就出现了从箭头位置向两边延伸的效果
					p.SetMask2(gui.NewQRegion2(int(p.luDis*(1-p.currentProgress)),
						0,
						int(p.currentProgress*float64(p.Width())),
						p.Height(),
						gui.QRegion__Rectangle))

					// 计数
					ticker += 20
					if ticker >= interval {

						//
						p.isShow = true

						// 告诉上层，动画完成了.
						if f != nil {
							f(p)
						}

						return
					}

					// 刷新
					p.Update()

				}
			}
		}()
	}
}

// 隐藏
// interval：隐藏的持续时间, 1000 = 1s
func (p *PopupWidget) Hide(interval int) {
	// 直接调用Hide2, 代码统一
	p.Hide2(interval, nil)
}

// 隐藏
// interval：隐藏的持续时间, 1000 = 1s
// f: 动画结束时相应，可能某些人喜欢在动画结束后执行一些操作，就可以这么写
func (p *PopupWidget) Hide2(interval int, f func(widget *PopupWidget)) {
	if p.isShow {
		// 只有显示状态才隐藏
		p.ticker = time.NewTicker(50 * time.Millisecond)

		// 默认遮罩，什么也不显示
		p.SetMask2(gui.NewQRegion2(0,
			0,
			p.Width(),
			p.Height(),
			gui.QRegion__Rectangle))

		go func() {
			defer p.ticker.Stop()

			var ticker = 0
			for {
				select {
				case <-p.ticker.C:
					// 1.0 / (interval/20)
					p.currentProgress -= 20.0 / float64(interval)

					fmt.Println(p.currentProgress, interval, 20.0/float64(interval))
					if p.currentProgress < 0 {
						p.currentProgress = 0
						// 都跑到了，就直接让它结束吧，
						ticker = interval
					}

					// 修改遮罩位置
					// 添加一个遮罩层，我们的遮罩应该以箭头为中心
					// 这样就能算出实际应该遮罩的位置，改变currentProgress， 就出现了从箭头位置向两边延伸的效果
					p.SetMask2(gui.NewQRegion2(int(p.luDis*(1-p.currentProgress)),
						0,
						int(p.currentProgress*float64(p.Width())),
						p.Height(),
						gui.QRegion__Rectangle))

					// 计数
					ticker += 20
					if ticker >= interval {
						//
						p.isShow = false

						// 告诉上层，动画完成了.
						if f != nil {
							f(p)
						}

						// 刷新
						p.Update()

						return
					}

					// 刷新
					p.Update()
				}
			}
		}()
	}
}

// 获取渐变色,阴影颜色
func (p *PopupWidget) getShadowColor() *gui.QLinearGradient {
	var linear = gui.NewQLinearGradient3(float64(0+p.pX), float64(p.Y()+p.Height()/2),
		float64(p.Width())+p.pX, float64(p.Y()+p.Height()/2))

	linear.SetColorAt(0, gui.NewQColor3(0, 0, 0, 0x6F))
	linear.SetColorAt(0.9, gui.NewQColor3(0x0, 0x0, 0x0, 0x3F))
	linear.SetColorAt(1, gui.NewQColor3(0, 0, 0, 0x0F))

	return linear
}

// 画图
// 绘制较简单，就不将painter作为全局变量了
func (p *PopupWidget) onPaint(event *gui.QPaintEvent) {
	//// 创建画笔
	var device = p.BackingStore().PaintDevice() //window.Painters[0]
	var painter = gui.NewQPainter2(device)
	// 反走样
	painter.SetRenderHint(gui.QPainter__Antialiasing, true)

	var shadowPath = gui.NewQPainterPath()
	var path = gui.NewQPainterPath()

	// 我这里只画一个向上的，其它基本一样，就是坐标不一样罢了
	// 我最讨厌go里面的不同类型的运算，转来转去，麻烦....
	switch p.arrowDirection {
	case ArrowDirectionUp:
		// 办法1：这个是通过遮罩实现的，但是看不到左右的边，如果有阴影，效果就不是很好了。
		//// 这些代码是显示一个完整的带尖角的代码。
		//path.MoveTo2(0, ArrowWidth)
		//path.LineTo2(p.luDis-ArrowWidth/2.0, ArrowWidth)
		//// 下面两句是画箭头的
		//path.LineTo2(p.luDis, 0)
		//path.LineTo2(p.luDis+ArrowWidth/2.0, ArrowWidth)
		//path.LineTo2(float64(p.Width()), ArrowWidth)
		//path.LineTo2(float64(p.Width()), float64(p.Height()))
		//path.LineTo2(0, float64(p.Height()))
		//path.LineTo2(0, ArrowWidth)
		//
		//// 添加一个遮罩层，我们的遮罩应该以箭头为中心
		//// 这样就能算出实际应该遮罩的位置，改变currentProgress， 就出现了从箭头位置向两边延伸的效果
		//path2.AddRect2(p.luDis*(1-p.currentProgress), 0, p.currentProgress*float64(p.Width()), float64(p.Height()))
		//painter.DrawPath(path2)

		// 办法2，直接根据currentProgress，动态画这个带尖角的图形，可以稍微往里留一两个像素，那么就能看到阴影了.
		// 注意，这个图形最小宽度应为 ArrowWidth+(ShadowWidth*2)， 左右留出ShadowWidth大小的像素画阴影。

		// 这里要考虑两种情况，
		// 第一： currentProgress*p.width() < ArrowWidth+(ShadowWidth*2) 时
		// 第二： currentProgress*p.width() >= ArrowWidth+(ShadowWidth*2) 时
		// 设置颜色。

		//=======================================================================
		// 阴影的原理，就是将图形, 向左或者向右，向上或者向下移动一定像素，图形透明度降低
		// 画阴影
		painter.SetPen3(core.Qt__NoPen)
		painter.SetBrush(gui.NewQBrush10(p.getShadowColor()))

		var rw = p.currentProgress * float64(p.Width()-ShadowWidth*2)
		var rx = 0.0
		if rw < ArrowWidth+ShadowWidth*2 {
			rx = p.luDis - ArrowWidth/2 + 2
		} else {
			rx = p.luDis*(1-p.currentProgress) + 2
		}

		// 这样就画出了不同的图形，最小宽度ArrowWidth+(ShadowWidth*2)
		shadowPath.MoveTo2(rx+p.pX, ArrowWidth+2+p.pY)
		shadowPath.LineTo2(p.luDis-ArrowWidth/2+p.pX, ArrowWidth+2+p.pY)
		// 画角
		shadowPath.LineTo2(p.luDis+p.pX, 0+2+p.pY)
		shadowPath.LineTo2(p.luDis+ArrowWidth/2+p.pX, ArrowWidth+2+p.pY)
		shadowPath.LineTo2(rx+rw+p.pX, ArrowWidth+2+p.pY)
		// 留出底部的阴影区域
		shadowPath.LineTo2(rx+rw+p.pX, float64(p.Height())-ShadowWidth+2+p.pY)
		shadowPath.LineTo2(rx+p.pX, float64(p.Height())-ShadowWidth+2+p.pY)
		shadowPath.LineTo2(rx+p.pX, ArrowWidth+2+p.pY)
		// 显示图形。
		painter.DrawPath(shadowPath)

		//=======================================================================
		// 画正常的图形
		// 设置颜色。
		painter.SetPen2(gui.NewQColor3(0x7F, 0x7F, 0x7F, 0x7F))
		painter.SetBrush(gui.NewQBrush3(gui.NewQColor3(0xF5, 0xF5, 0xF5, 0xFF), 1))

		rw = p.currentProgress * float64(p.Width()-ShadowWidth*2)
		rx = 0.0
		if rw < ArrowWidth+ShadowWidth*2 {
			rx = p.luDis - ArrowWidth/2
		} else {
			rx = p.luDis * (1 - p.currentProgress)
		}

		// 这样就画出了不同的图形，最小宽度ArrowWidth+(ShadowWidth*2)
		path.MoveTo2(rx+p.pX, ArrowWidth+p.pY)
		path.LineTo2(p.luDis-ArrowWidth/2+p.pX, ArrowWidth+p.pY)
		// 画角
		path.LineTo2(p.luDis+p.pX, 0+p.pY)
		path.LineTo2(p.luDis+ArrowWidth/2+p.pX, ArrowWidth+p.pY)
		path.LineTo2(rx+rw+p.pX, ArrowWidth+p.pY)
		// 留出底部的阴影区域
		path.LineTo2(rx+rw+p.pX, float64(p.Height())-ShadowWidth+p.pY)
		path.LineTo2(rx+p.pX, float64(p.Height())-ShadowWidth+p.pY)
		path.LineTo2(rx+p.pX, ArrowWidth+p.pY)

		// 显示图形。
		painter.DrawPath(path)
	case ArrowDirectionDown:
	}

	// 结束绘制
	painter.End()
	painter.DestroyQPainter()
}
