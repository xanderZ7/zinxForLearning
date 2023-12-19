package ziface

// 实际上是把客户端请求的连接信息和请求的数据包装在一个Request中

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte
}
