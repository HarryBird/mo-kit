package goredis

type IOption interface {
	GetAddr() string
	GetDb() int32
	GetPassword() string
	GetDialTimeout() int32
	GetIdleTimeout() int32
	GetMaxConnAge() int32
	GetMinIdleConns() int32
	GetPoolSize() int32
	GetPoolTimeout() int32
	GetPoolType() string
	GetReadTimeout() int32
	GetWriteTimeout() int32
}
