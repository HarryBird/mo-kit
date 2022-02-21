package goredis

import (
	"context"

	klog "github.com/go-kratos/kratos/v2/log"
)

const (
	infoStr = "%s\n[info] "
)

type Logger struct {
	logger *klog.Helper
}

func New(l klog.Logger) Logger {
	return Logger{
		logger: klog.NewHelper(klog.With(l, "mod", "repo.redis")),
	}
}

func (l Logger) Printf(ctx context.Context, format string, data ...interface{}) {
	l.logger.WithContext(ctx).Infof("[redis] "+format, data...)
}
