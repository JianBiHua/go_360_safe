/**
 * Copyright (C), 2015-2019, 简笔画工作室
 * FileName: main.go
 * Author: 简笔画
 * Date: 2019.03.26 21:06
 * Description: 360安全卫士，入口文件。
 * History:
 * <author> <time> <version> <desc>
 * 作者姓名 修改时间 版本号 描述
 */
package main

import (
	"GoWorkspace/qo_360_safe/src/views"
	"fmt"
	"strconv"
)

// 入口
func main() {

	fmt.Println(strconv.FormatFloat(100.0, 'f', -1, 64))

	// 打印一个开始消息
	fmt.Println("start Go360Safe")
	// 创建mainwindow
	mainWindow := views.NewMainWindow()
	// 显示主窗口
	mainWindow.Show()
	// 进入循环
	mainWindow.Exec()
}
