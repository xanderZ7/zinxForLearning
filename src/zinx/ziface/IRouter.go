package ziface

// 路由抽象接口
type IRouter interface {
	// 在处理conn业务之前的钩子方法
	PreHandle(request IRequest)
	// 在处理conn业务的主方法 hook
	Handle(request IRequest)
	// 在处理conn业务之后的钩子方法
	PostHandle(request IRequest)
}
