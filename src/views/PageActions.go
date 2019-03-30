package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/uitools"
	"github.com/therecipe/qt/widgets"
)

// PageActions is more action's sheet
type PageActions struct {
	widget *widgets.QWidget
}

// NewPageActions is 析构函数
func NewPageActions() *PageActions {
	ph := new(PageActions)
	ph.init()
	return ph
}

// init 初始化变量
// 从UI中加载widget
func (ph *PageActions) init() {
	ph.widget = widgets.NewQWidget(nil, 0)

	var loader = uitools.NewQUiLoader(nil)
	var file = core.NewQFile2("ui/page_actions.ui")

	file.Open(core.QIODevice__ReadOnly)
	var formWidget = loader.Load(file, ph.widget)
	file.Close()

	//var (
	//	ui_inputSpinBox1 = widgets.NewQSpinBoxFromPointer(widget.FindChild("inputSpinBox1", core.Qt__FindChildrenRecursively).Pointer())
	//	ui_inputSpinBox2 = widgets.NewQSpinBoxFromPointer(widget.FindChild("inputSpinBox2", core.Qt__FindChildrenRecursively).Pointer())
	//	ui_outputWidget  = widgets.NewQLabelFromPointer(widget.FindChild("outputWidget", core.Qt__FindChildrenRecursively).Pointer())
	//)
	//
	//ui_inputSpinBox1.ConnectValueChanged(func(value int) {
	//	ui_outputWidget.SetText(fmt.Sprint(value + ui_inputSpinBox2.Value()))
	//})
	//
	//ui_inputSpinBox2.ConnectValueChanged(func(value int) {
	//	ui_outputWidget.SetText(fmt.Sprint(value + ui_inputSpinBox1.Value()))
	//})

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(formWidget, 0, 0)
	ph.widget.SetLayout(layout)
}

// Widget is back a widgets.QWidget pointer
func (ph *PageActions) Widget() *widgets.QWidget {
	return ph.widget
}
