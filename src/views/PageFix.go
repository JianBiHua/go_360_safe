package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

type PageFix struct {
	widget *widgets.QWidget
}

func NewPageFix() *PageFix {
	ph := new(PageFix)
	ph.init()
	return ph
}

// 从UI中加载widget
func (ph *PageFix) init() {
	ph.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2("ui/page_fix.ui")

	file.Open(core.QIODevice__ReadOnly)
	loader.Load(file, ph.widget)
	file.Close()

	var (
		ui_pushbutton = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("roundButton", core.Qt__FindChildrenRecursively).Pointer())
		ui_title      = widgets.NewQLabelFromPointer(ph.widget.FindChild("title", core.Qt__FindChildrenRecursively).Pointer())
		ui_subtitle   = widgets.NewQLabelFromPointer(ph.widget.FindChild("subtitle", core.Qt__FindChildrenRecursively).Pointer())
	)

	ui_pushbutton.SetStyleSheet(`border:none;
		background-color: #16da6c;
		border-radius: 30px;
		color: #FFFFFF;
		font-size: 25px;`)
	ui_title.SetStyleSheet(`color: #444444;font-size: 25px;`)
	ui_subtitle.SetStyleSheet(`color: #9b9999;font-size: 16px;`)
}

func (ph *PageFix) Widget() *widgets.QWidget {
	return ph.widget
}
