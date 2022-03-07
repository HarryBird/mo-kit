package app

import (
	"context"

	"github.com/HarryBird/mo-kit/msgr"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	klog "github.com/go-kratos/kratos/v2/log"
)

func LogRequest(ctx context.Context, log *klog.Helper, fname string, req interface{}) {
	log.WithContext(ctx).Infof("%s Request Begin...", msgr.W(fname))
	log.WithContext(ctx).Debugf("%s Request Param: %+v", msgr.W(fname), req)
}

func LogResponse(ctx context.Context, log *klog.Helper, fname string, resp interface{}) {
	log.WithContext(ctx).Infof("%s Request End...", msgr.W(fname))
	log.WithContext(ctx).Debugf("%s Response: %+v", msgr.W(fname), resp)
}

func LogErrorStack(ctx context.Context, log *klog.Helper, fname string, err error) {
	log.WithContext(ctx).Errorf("%s error stack=%+v", msgr.W(fname), err)
}

func LogErrorRPC(ctx context.Context, log *klog.Helper, fname string, err error) {
	e := kerrors.FromError(err)
	log.WithContext(ctx).Errorf("%s call rpc fail: code=%d, msg=%s, reason=%s, metadata=%v",
		msgr.W(fname), e.GetCode(), e.GetMessage(), e.GetReason(), e.GetMetadata())
}
