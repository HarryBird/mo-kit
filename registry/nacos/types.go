package nacos

type IServerConfig interface {
	GetAddress() string
	GetPort() uint64
	GetScheme() string
	GetContextPath() string
}
