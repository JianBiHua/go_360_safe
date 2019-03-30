package views

// 显示界面的代码
func ExampleMainWindow_Show() {
	mw := NewMainWindow()
	mw.Show()
	mw.Exec()
}
