package goredis

import (
	"errors"
	"strings"
	"time"

	rlog "github.com/HarryBird/mo-kit/log/kratos/goredis"
	klog "github.com/go-kratos/kratos/v2/log"
	redis "github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

func NewRedis(opts IOption, l klog.Logger) (*redis.Client, error) {
	if opts.GetAddr() == "" {
		return nil, errors.New("redis: invalid connection address")
	}

	o := &redis.Options{
		Addr:         opts.GetAddr(),
		Password:     "",
		DB:           0,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		PoolFIFO:     true,
		PoolSize:     10,
		MinIdleConns: 2,
		PoolTimeout:  2 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	if opts.GetDb() > 0 {
		o.DB = cast.ToInt(opts.GetDb())
	}

	if opts.GetPassword() != "" {
		o.Password = opts.GetPassword()
	}

	if opts.GetDialTimeout() > 0 {
		o.DialTimeout = int32ToMilliSecond(opts.GetDialTimeout())
	}

	if opts.GetReadTimeout() > 0 {
		o.ReadTimeout = int32ToMilliSecond(opts.GetReadTimeout())
	}

	if opts.GetWriteTimeout() > 0 {
		o.WriteTimeout = int32ToMilliSecond(opts.GetWriteTimeout())
	}

	if opts.GetPoolTimeout() > 0 {
		o.PoolTimeout = int32ToMilliSecond(opts.GetPoolTimeout())
	}

	if opts.GetIdleTimeout() > 0 {
		o.IdleTimeout = int32ToMilliSecond(opts.GetIdleTimeout())
	}

	if opts.GetPoolType() != "" && !strings.EqualFold(opts.GetPoolType(), "FIFO") {
		o.PoolFIFO = false
	}

	if opts.GetPoolSize() > 0 {
		o.PoolSize = cast.ToInt(opts.GetPoolSize())
	}

	if opts.GetMinIdleConns() > 0 {
		o.MinIdleConns = cast.ToInt(opts.GetMinIdleConns())
	}

	// fmt.Printf("goredis: connect options: %+v", o)

	redis.SetLogger(rlog.New(l))

	return redis.NewClient(o), nil
}

func int32ToMilliSecond(t int32) time.Duration {
	return time.Duration(cast.ToInt64(t)) * time.Millisecond
}
