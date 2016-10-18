package core

import (
	"fmt"
	"strconv"
)

type Base struct {
	Name string
	data map[string] interface{}
}

func  ( this *Base ) GetName(){
	fmt.Printf( this.Name )
}


func ( this *Base ) Print( i interface{} ) {
	this.print_line_color_str(31, 40)
	fmt.Printf("\x1b[37;1m%s%v\n\n\n", " ", i)
	this.print_line_color_str(31, 40)
}

/*********************************************************
 * 输出色彩分隔行
 *********************************************************/
func ( this *Base ) print_line_color_str( frontcolor int, backgroundcolor int ) {
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
	this.print_line_color_str(31, 40)
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
	this.print_line_color_str(31, 40)
}


