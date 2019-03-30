package views

//func TestMainWindow_Show(t *testing.T)  {
//	mw := NewMainWindow()
//	mw.Show()
//}

// 显示界面的代码
func ExampleMainWindow_Show() {
	mw := NewMainWindow()
	mw.Show()
	mw.Exec()
}
