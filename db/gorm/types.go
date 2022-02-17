package gorm

type Pooler interface {
	GetMaxIdle() int32
	GetMaxOpen() int32
	GetMaxLifeTime() int32
}
