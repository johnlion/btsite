package web

import "flag"

var (
	ip              *string
	port            *int
	addr            string
	spiderMenu      []map[string]string
)

// 获取外部参数
func Flag(){
	flag.String("b ******************************************** only for web ******************************************** -b", "", "")
	// web服务器IP与端口号
	ip = flag.String( "b_ip", "0.0.0.0", "  <Web Server IP>" )
	port = flag.Int( "b_port", 9090, "      <Web Server POrt>" )
}