package status

// 运行模式
const (
						// val
	UNSET int = iota - 1                    // -1
	OFFLINE                                 // 0
	SERVER                                  // 1
	CLIENT                                  // 2
)

// 数据头部信息
const (
	// 任务请求Header
	REQTASK = iota + 1                      // 1
	// 任务响应流头Header
	TASK                                    // 2
	// 打印Header
	LOG                                     // 3
)

// 运行状态
const (
	STOPPED = iota - 1                      // -1
	STOP                                    // 0
	RUN                                     // 1
	PAUSE                                   // 2
)
