package core

import (
	"fmt"
	"strconv"
	"time"
)

type Base struct {
	Name string
	Data map[string] interface{}
	time1 float64                                        //启动时的时间
	time2 float64                                         //执行结束的时间
	totalTime float64                                    //总消耗的时间
}

func  ( this *Base ) GetName(){
	fmt.Printf( this.Name )
}

func ( this *Base ) SetTime1(){
	this.time1 =float64( time.Now().UnixNano() )
}

func ( this *Base ) SetTime2(){
	this.time2 =  float64( time.Now().UnixNano() )
}

func ( this *Base ) SetTotalTime(){
	this.totalTime = ( this.time2 -  this.time1 ) /1000000
}

func ( this *Base ) GetTotalTime()  string{
	this.SetTotalTime()
	result :="Time:" + string( fmt.Sprintf( "%.4f", this.totalTime )) + "s"
	return result
}


func ( this *Base ) Print( i interface{} ) {
	fmt.Printf("\x1b[33;1m%s%v\n\n\n", " ", i)
}

/*********************************************************
 * 输出色彩分隔行
 *********************************************************/
func ( this *Base ) Print_line_color_str( frontcolor int, backgroundcolor int ) {
	// frontcolor , backgroundcolor
	// 前景 背景 颜色
	// ---------------------------------------
	// 30  40  黑色
	// 31  41  红色
	// 32  42  绿色
	// 33  43  黄色
	// 34  44  蓝色
	// 35  45  紫红色
	// 36  46  青蓝色
	// 37  47  白色
	//
	// 代码 意义
	// -------------------------
	//  0  终端默认设置
	//  1  高亮显示
	//  4  使用下划线
	//  5  闪烁
	//  7  反白显示
	//  8  不可见

	fmt.Printf("\n %c[1;"+ strconv.Itoa( frontcolor ) +";" + strconv.Itoa( backgroundcolor ) + "m%s%c[0m\n\n", 0x1B, "##################################################################################################################", 0x1B)
}

/*********************************************************
 * 输出json data
 *********************************************************/
func (this *Base ) Print_data_json ( m map[string] interface{} ){
	//Example:
	//numbers := make( map[string] interface{} )
	//numbers["one"] =1
	//numbers["two"] =2.03
	//numbers["three"] ="hello"
	//numbers["four"] =4
	//numbers["five"] =5
	//baseController.Print_data_json( numbers )

	this.Print_line_color_str(31, 40)
	for k, v := range m {
		switch vv := v.(type) {
			case string:
				fmt.Println( " ",k, "is string", vv )
			case int:
				fmt.Println( " ",k, "is int", vv )
			case float32:
				fmt.Println( " ",k, "is float32", int32( vv ) )
			case float64:
				fmt.Println( " ", k, "is float64", int64( vv ))
			case nil:
				fmt.Println( " ",k ,"is nil", "null" )
			case map[string]interface{}:
				fmt.Println( " ",k, "is an map:" )
				this.Print_data_json( vv )
			default:
				fmt.Println( " ",k, "is of type I don't know how to handle", fmt.Sprintf( "%T", v ) )
		}
	}
	this.Print_line_color_str(31, 40)
}

