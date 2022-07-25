package iface

type IServer interface {
	// 启动服务器处理逻辑
	Start()
	// 停滞服务器
	Stop()
	// 运行服务器
	Run()
}
