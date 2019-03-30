package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// 这个结构体，暂时只需要一个QWidget，之后肯定会被扩充
type PageHome struct {
	widget *widgets.QWidget
}

// 这些都跟之前一样的
func NewPageHome() *PageHome {
	ph := new(PageHome)
	ph.init()
	return ph
}

// 从UI中加载widget
// 核心代码
func (ph *PageHome) init() {
	ph.widget = widgets.NewQWidget(nil, 0)

	// 核心代码1: 加载page_home.ui，加载不同的页面，需要修改*.ui
	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2("ui/page_home.ui")

	// 核心代码2：让page_home.ui与widget关联
	file.Open(core.QIODevice__ReadOnly)
	loader.Load(file, ph.widget)
	file.Close()

	var (
		// 核心代码3： 从ui中获取objectname=roundButton等的控件，
		// 注意NewQSpinBoxFromPointer，NewQLabelFromPointer对应不一样的控件，
		ui_pushbutton = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("roundButton", core.Qt__FindChildrenRecursively).Pointer())
		ui_title      = widgets.NewQLabelFromPointer(ph.widget.FindChild("title", core.Qt__FindChildrenRecursively).Pointer())
		ui_subtitle   = widgets.NewQLabelFromPointer(ph.widget.FindChild("subtitle", core.Qt__FindChildrenRecursively).Pointer())

		fangHuZhongXin = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("fangHuZhongXin", core.Qt__FindChildrenRecursively).Pointer())
		wangGouXianPei = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("wangGouXianPei", core.Qt__FindChildrenRecursively).Pointer())
		fanLeSuo       = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("fanLeSuo", core.Qt__FindChildrenRecursively).Pointer())

		tbApp     = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbApp", core.Qt__FindChildrenRecursively).Pointer())
		tbService = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbService", core.Qt__FindChildrenRecursively).Pointer())
		tbHandler = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbHandler", core.Qt__FindChildrenRecursively).Pointer())
		tbFix     = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbFix", core.Qt__FindChildrenRecursively).Pointer())
		tbMore    = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbMore", core.Qt__FindChildrenRecursively).Pointer())
	)

	// 修改控件的样式
	ui_pushbutton.SetStyleSheet(`border:none;
		background-color: #16da6c;
		border-radius: 30px;
		color: #FFFFFF;
		font-size: 25px;`)
	ui_title.SetStyleSheet(`color: #444444;font-size: 25px;`)
	ui_subtitle.SetStyleSheet(`color: #9b9999;font-size: 16px;`)

	// 这个是ToolButton上节出现的，这里又加了一个构造函数,
	var tb = NewToolButton3(fangHuZhongXin, "防护中心", "resources/other/icon_fanghuzhongxin.png", 40, 40)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					`)

	tb = NewToolButton3(wangGouXianPei, "网购先陪", "resources/other/icon_wanggou.png", 40, 40)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					`)
	tb = NewToolButton3(fanLeSuo, "反勒索服务", "resources/other/icon_fanlesuo.png", 40, 40)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					`)

	tb = NewToolButton3(tbApp, "软件管理", "resources/other/icon_app.png", 42, 42)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					padding-bottom: 10px;
					`)
	tb = NewToolButton3(tbService, "主页修复", "resources/other/icon_fix.png", 42, 42)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					padding-bottom: 10px;
					`)
	tb = NewToolButton3(tbHandler, "手机助手", "resources/other/icon_handler.png", 42, 42)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					padding-bottom: 10px;
					`)
	tb = NewToolButton3(tbFix, "更多功能", "resources/other/icon_more.png", 42, 42)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					padding-bottom: 10px;
					`)
	tb = NewToolButton3(tbMore, "人工服务", "resources/other/icon_help.png", 42, 42)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					font-size: 10px;
					padding-bottom: 10px;
					`)
}

// 这玩意，跟之前都是一样的，为了在其它页面调用到QT的控件指针，方便修改属性等
func (ph *PageHome) Widget() *widgets.QWidget {
	return ph.widget
}
