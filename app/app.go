package app

import (
	"io"

)

type(
	App interface {
		SetLog( io.Writer ) App                                                 //　设置全局log实时显示终端
		LogGoOn() App                                                           // 继续log打印
		LogRest() App
		Init( mode int, port int, master string, w ...io.Writer ) App           // 使用App前必须进行先　Init初始化,　SetLog()除外
		ReInit( mode int, prot int, master string, w ...io.Writer ) App         // 切换运行模式闭幕式重设log打印目标
		GetAppConf( k ...string ) interface{}                                   // 获取全局参数
		SetAppConf( k string, v interface{} ) App                               // 设置全局参数　(　client模式下不调用该方法　)
		SpiderPrepare( original []string ) App                                  // 须在设置全局运行参数后Run()前调用（client模式下不调用该方法）
		Run()                                                                   // 阻塞式运行直至任务完成（须在所有应当配置项配置完成后调用）
		Stop()                                                                  // Offline 模式下中途终止任务（对外为阻塞式运行直至当前任务终止）
		IsRunning() bool                                                        // 检查任务是否正在运行
		IsPause() bool                                                          // 检查任务是否处于暂停状态
		IsStopped() bool                                                        // 检查任务是否已经终止
		PauseRecover()                                                          // Offline 模式下暂停\恢复任务
		Status() int                                                            // 返回当前状态
		GetSpiderLib()                                                          // 获取全部蜘蛛种类
		GetSpiderByName( string )                                               // 通过名字获取某蜘蛛
		GetSpiderQueue()                                                        // 获取蜘蛛队列接口实例
		GetOutputLib()[]string                                                  // 获取全部输出方式
		GetTaskJar()                                                            // 返回任务库

	}
)