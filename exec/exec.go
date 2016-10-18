package exec

import (
	"runtime"
	"github.com/johnlion/btsite/config"
	"github.com/johnlion/btsite/core"
)

var (
	baseController core.Base
)

func init(){
	// 开启最大核心数运行
	runtime.GOMAXPROCS(runtime.NumCPU())

	baseController.Name = config.FULL_NAME

}

func DefaultRun( uiDefault string ) {

	baseController.Print( uiDefault )

	numbers := make( map[string] interface{} )
	numbers["one"] =1
	numbers["two"] =2.03
	numbers["three"] ="hello"
	numbers["four"] =4
	numbers["five"] =5
	baseController.Print_data_json( numbers )

}



