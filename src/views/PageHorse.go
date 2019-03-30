package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type PageHorse struct {
	widget *widgets.QWidget
}

func NewPageHorse() *PageHorse {
	ph := new(PageHorse)
	ph.init()
	return ph
}

func (ph *PageHorse) init() {
	ph.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2("ui/page_horse.ui")

	file.Open(core.QIODevice__ReadOnly)
	loader.Load(file, ph.widget)
	file.Close()

	var (
		ui_pushbutton = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("roundButton", core.Qt__FindChildrenRecursively).Pointer())
		ui_title      = widgets.NewQLabelFromPointer(ph.widget.FindChild("title", core.Qt__FindChildrenRecursively).Pointer())
		ui_subtitle   = widgets.NewQLabelFromPointer(ph.widget.FindChild("subtitle", core.Qt__FindChildrenRecursively).Pointer())

		ui_pushButton_2 = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("pushButton_2", core.Qt__FindChildrenRecursively).Pointer())
		ui_pushButton_3 = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("pushButton_3", core.Qt__FindChildrenRecursively).Pointer())
		ui_pushButton_4 = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("pushButton_4", core.Qt__FindChildrenRecursively).Pointer())
		ui_pushButton_5 = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("pushButton_5", core.Qt__FindChildrenRecursively).Pointer())
		ui_pushButton_6 = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("pushButton_6", core.Qt__FindChildrenRecursively).Pointer())

		tbQuanpan    = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbQuanpan", core.Qt__FindChildrenRecursively).Pointer())
		tbAnweizhi   = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbAnweizhi", core.Qt__FindChildrenRecursively).Pointer())
		tbXinrenqu   = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbXinrenqu", core.Qt__FindChildrenRecursively).Pointer())
		tbHuifuqu    = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbHuifuqu", core.Qt__FindChildrenRecursively).Pointer())
		tbShangbanqu = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbShangbanqu", core.Qt__FindChildrenRecursively).Pointer())
		tbXitong     = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbXitong", core.Qt__FindChildrenRecursively).Pointer())
		tbShouji     = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbShouji", core.Qt__FindChildrenRecursively).Pointer())
		tbWinPE      = widgets.NewQToolButtonFromPointer(ph.widget.FindChild("tbWinPE", core.Qt__FindChildrenRecursively).Pointer())
	)

	ui_pushbutton.SetStyleSheet(`border:none;
		background-color: #16da6c;
		border-radius: 30px;
		color: #FFFFFF;
		font-size: 25px;`)
	ui_title.SetStyleSheet(`color: #444444;font-size: 25px;`)
	ui_subtitle.SetStyleSheet(`color: #9b9999;font-size: 16px;`)
	ui_pushButton_2.SetStyleSheet("border-image: url(resources/horse/icon_pb_bg_1.png)")
	ui_pushButton_3.SetStyleSheet("border-image: url(resources/horse/icon_pb_bg_2.png)")
	ui_pushButton_4.SetStyleSheet("border-image: url(resources/horse/icon_pb_bg_3.png)")
	ui_pushButton_5.SetStyleSheet("border-image: url(resources/horse/icon_pb_bg_4.png)")
	ui_pushButton_6.SetStyleSheet("border-image: url(resources/horse/icon_pb_bg_5.png)")

	var tb = NewToolButton3(tbQuanpan, "全盘查杀", "resources/horse/icon_quanpansaomiao.png", 34, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					`)
	tb = NewToolButton3(tbAnweizhi, "按位置查杀", "resources/horse/icon_anweizhi.png", 34, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					`)
	//
	tb = NewToolButton3(tbXinrenqu, "信任区", "resources/horse/icon_xinrenqu.png", 32, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					`)
	tb = NewToolButton3(tbHuifuqu, "恢复区", "resources/horse/icon_huifuqu.png", 32, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #FF0000;
					color:rgb(124, 124, 124);
					`)
	tb = NewToolButton3(tbShangbanqu, "上报区", "resources/horse/icon_shangbaoqu.png", 32, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					`)
	tb = NewToolButton3(tbXitong, "系统急救箱", "resources/horse/icon_icon_1.png", 32, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					`)
	tb = NewToolButton3(tbShouji, "手机急救箱", "resources/horse/icon_icon_2.png", 32, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					`)

	tb = NewToolButton3(tbWinPE, "WinPE版", "resources/horse/icon_icon_3.png", 32, 32)
	tb.Widget().SetStyleSheet(`
					border: 0px;
					border-radius: 0px;
					background-color: #00000000;
					color:rgb(124, 124, 124);
					`)
}

func (ph *PageHorse) Widget() *widgets.QWidget {
	return ph.widget
}
