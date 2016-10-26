package main

import (
	"github.com/johnlion/btsite/exec"

)

func main(){
	// 设置运行时默认操作界面，并开始运行
	// 运行软件前，可设置 -a_ui 参数为"web"、"gui"或"cmd"，指定本次运行的操作界面
	// 其中"gui"仅支持Windows系统
	var uiDefault string = "cmd"
	exec.DefaultRun( uiDefault )
}


