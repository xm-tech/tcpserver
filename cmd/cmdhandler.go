package cmd

import (
	"github.com/xm-tech/tcpserver/iface"
	"github.com/xm-tech/tcpserver/util"
)

type CmdHandler struct {
	WorkerPoolSize uint32
	TaskQuene      []chan iface.IRequest
}

func NewCmdHandler() *CmdHandler {
	return &CmdHandler{
		WorkerPoolSize: util.GlobalConf.WorkerPoolSize,
		TaskQuene:      make([]chan iface.IRequest, util.GlobalConf.WorkerPoolSize),
	}
}

func (self *CmdHandler) OnRequest(request iface.IRequest) {
	connID := request.GetConnection().GetConnID()
	workerID := connID % self.WorkerPoolSize
	self.TaskQuene[workerID] <- request
}

func (self *CmdHandler) StartWorker(workerID int, taskQuene chan iface.IRequest) {
	for {
		select {
		case req := <-taskQuene:
			// 处理业务
			Cmds.GetCommand(int(req.GetMsgID())).Exec(req)
		}
	}
}

func (self *CmdHandler) Run() {
	for i := 0; i < int(self.WorkerPoolSize); i++ {
		// 一个 Worker 被启动，给当前Worker对应的任务队列开辟空间
		self.TaskQuene[i] = make(chan iface.IRequest, util.GlobalConf.MaxWorkerTaskLen)
		go self.StartWorker(i, self.TaskQuene[i])
	}
}
