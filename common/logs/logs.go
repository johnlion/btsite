package logs

import (
	"io"
	"github.com/johnlion/btsite/common/logs/logs"
	"path"
	"github.com/johnlion/btsite/config"
	"os"
)

type(
	Logs interface {
		// 设置实时log信息显示终端
		SetOutput( show io.Writer ) Logs
		// 暂停输出日志
		Rest()
		// 恢复暂停状态，继续输出日志
		GoOn()
		// 按先后顺序实时截获日志，每次返回1条，normal标记日志是否被关闭
		StealOne() (level int, msg string, normal bool)
		// 正常关闭日志输出
		Close()
		// 返回运行状态，如0,"RUN"
		Status() (int, string)
		DelLogger(adaptername string) error
		SetLogger(adaptername string, config map[string]interface{}) error

		// 以下打印方法除正常log输出外，若为客户端或服务端模式还将进行socket信息发送
		Debug(format string, v ...interface{})
		Informational(format string, v ...interface{})
		App(format string, v ...interface{})
		Notice(format string, v ...interface{})
		Warning(format string, v ...interface{})
		Error(format string, v ...interface{})
		Critical(format string, v ...interface{})
		Alert(format string, v ...interface{})
		Emergency(format string, v ...interface{})
	}

	mylog struct {
		*logs.BeeLogger
	}
)


var Log = func() Logs{
	p, _ := path.Split( config.LOG )
	//　不存在目录时分创建目录
	d , err := os.Stat(p)
	if err != nil || d.IsDir(){
		if err := os.MkdirAll( p, 0777 ); err != nil{
			 Log.Error( "Error:　%v\n", err )
		}
	}
}