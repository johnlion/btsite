package cache

import (
	"time"
	"sync/atomic"
	"runtime"
)

//**************************************任务运行时公共配置****************************************\\

// 任务运行时公共配置
type AppConf struct {
	Mode                                    int                     // 节点角色
	Port                                    int                     // 主节点端口
	Master                                  string                  // 服务器(主节点)地址,　不含端口
	ThreadNum                               int                     // 全局最大并发量
	PauseTime                               int64                   // 暂停时长参考/ms( 随机:　PauseTime/2 ~ PauseTime*2 )
	OutType                                 string                  // 输出方式
	DockerCap                               int                     // 分段转储容器容量
	DockerQueueCap                          int                     // 分段输出池容量,　不小于2
	Limit                                   int64                   // 采集上限,0为不限,　若在规则中设置初始值为LIMIT 则为自定义限制,否则默认限制请求数
	ProxyMinute                             int64                   // 代理IP更换的间隔分钟数
	SuccessInherit                          bool                    // 继承历史成功记录
	FailureInherit                          bool                    // 继承历史失败记录
	//选填项
	Keyins                                  string                  // 自定义输入,后期切分为多个任务的 Keyin自定义配置

}

// 设置初始值即默认值
var Task = new ( AppConf )

// 根据Task.DockerCap智能调整分段输出池容量Task.DockerQueueCap

func AutoDockerCap(){
	switch {
	case    Task.DockerCap <= 10:
		Task.DockerQueueCap = 500
	case    Task.DockerCap <=500:
		Task.DockerQueueCap = 200
	case    Task.DockerCap <=1000:
		Task.DockerQueueCap = 100
	case    Task.DockerCap <=10000:
		Task.DockerQueueCap = 50
	case    Task.DockerCap <=100000:
		Task.DockerQueueCap = 10
	default:
		Task.DockerQueueCap = 4
	}
}

//****************************************任务报告*******************************************\\

type Report struct {
	SpiderName      string
	Keyin           string
	DataNum         uint64
	FileNum         uint64
	Time            time.Duration
}


var (
	// 点击开始按钮的时间点
	StartTime time.Time
	// 文本数据小结报告
	ReportChan chan *Report
	// 请求页面总数[]uint{总数，失败数}
	pageSum [2]uint64

)

// 重置页面计数
func ResetPageCount(){
	pageSum = [2]uint64{}
}

// 0 返回总下载页数，负数 返回失败数，正数 返回成功数
func GetPageCount( i int ) uint64{
	switch
	{
	case i > 0 :
		// 返回成功数
		return pageSum[0]
	case i < 0 :
		// 返回失败数
		return pageSum[1]
	case i == 0:
		// 返回总数
		return pageSum[0] + pageSum[1]
	}
	// 返回总数
	return pageSum[0] + pageSum[1]
}

func PageSuccCount(){
	atomic.AddUint64( &pageSum[0], 1 )
}

func PageFailCount(){
	atomic.AddUint64( &pageSum[1], 1 )
}

//****************************************init函数执行顺序控制*******************************************\\

var initOrder = make( map[int]bool )

// 标记当前init()已执行完毕
func ExecInit( order int ){
	initOrder[order] = true
}

// 等待指定init()执行完毕
// 需并发协程中调用
func WaitInit( order int ){
	for !initOrder[order]{
		runtime.Gosched()
	}
}

//****************************************初始化*******************************************\\
func init(){
	// 任务报告
	ReportChan = make( chan *Report )

	// 根据Task.DockerCap智能调整分段输出池容量Task.DockerQueueCap
	AutoDockerCap()
}