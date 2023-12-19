package znet

import "zinx/src/zinx/ziface"

// 实现router时，先嵌入这个BaseRouter基类，然后根据需要对这个
// 基类的方法进行重写就行
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

func (br *BaseRouter) Handle(request ziface.IRequest) {

}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
