// 命令行界面版

package cmd

import (
	"flag"
	// "fmt"
	"strconv"
	// "strings"

	"github.com/johnlion/btsite/runtime/status"
)

var(
	spiderflag *string                              // 蜘蛛列表
)

func Flag(){
	// 分类说明
	flag.String("c ******************************************** only for cmd ******************************************** -c", "", "")

	// 蜘蛛列表
	spiderflag = flag.String(
		"c_spider",
		"",
		func() string{
			var spiderlist string
			// for k, v := range app.LogicApp
			return "      <蜘蛛列表: 选择多蜘蛛以 \",\" 间隔>\r\n" + spiderlist
		}())

	// 备注说明
	flag.String(
		"c_z",
		"",
		" CMD-EXAPLE: $ syd -_ui=cmd -a_mode=" + strconv.Itoa( status.OFFLINE ) + " -c_spider=3,8 -a_outtype=csV -a_thread=20 -a_dockercap=5000 -a_proxyminute=0 -a_keyins=\"  -a_limit=10 -a_success=true -a_failure=true\n",

	)
}