package views

import "github.com/therecipe/qt/widgets"

type WidgetInterface interface {
	Widget() *widgets.QWidget
}
