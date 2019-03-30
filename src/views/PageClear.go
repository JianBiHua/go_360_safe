package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// PageClear is clear page
// for clear system rubbish
type PageClear struct {
	widget *widgets.QWidget
}

// NewPageClear is 构造函数
func NewPageClear() *PageClear {
	ph := new(PageClear)
	ph.init()
	return ph
}

// init
// 从UI中加载widget
func (ph *PageClear) init() {
	ph.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2("ui/page_clear.ui")

	file.Open(core.QIODevice__ReadOnly)
	var formWidget = loader.Load(file, ph.widget)
	file.Close()

	var (
		uiPushbutton = widgets.NewQPushButtonFromPointer(ph.widget.FindChild("roundButton", core.Qt__FindChildrenRecursively).Pointer())
		uiTitle      = widgets.NewQLabelFromPointer(ph.widget.FindChild("title", core.Qt__FindChildrenRecursively).Pointer())
		uiSubtitle   = widgets.NewQLabelFromPointer(ph.widget.FindChild("subtitle", core.Qt__FindChildrenRecursively).Pointer())
	)

	uiPushbutton.SetStyleSheet(`border:none;
		background-color: #16da6c;
		border-radius: 30px;
		color: #FFFFFF;
		font-size: 25px;`)
	uiTitle.SetStyleSheet(`color: #444444;font-size: 25px;`)
	uiSubtitle.SetStyleSheet(`color: #9b9999;font-size: 16px;`)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(formWidget, 0, 0)
	ph.widget.SetLayout(layout)
}

// Widget is back a widgets.QWidget pointer
func (ph *PageClear) Widget() *widgets.QWidget {
	return ph.widget
}
