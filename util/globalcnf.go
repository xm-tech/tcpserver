package util

type globalCnf struct {
	WorkerPoolSize   uint32 // Worker 数量
	MaxWorkerTaskLen uint32 // 单 Worker 对应任务队列的最大长度, 也即消息积压最多条数
}

var GlobalCnf *globalCnf

func init() {
	GlobalCnf = &globalCnf{
		WorkerPoolSize:   20,
		MaxWorkerTaskLen: 1024,
	}
}
