// Package views WidgetInterface
package views

import "github.com/therecipe/qt/widgets"

// WidgetInterface is widget interface
type WidgetInterface interface {
	// Widget is return widget object
	Widget() *widgets.QWidget
}
