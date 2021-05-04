package interceptor

import (
	"context"
	"encoding/json"
	"errors"

	zlog "comic/share/log/zap"

	"github.com/micro/go-micro/v2/server"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

type FaceFrontError struct {
	// 如果字段名为code,就会掉进micro框架的坑里面
	StatusCode uint32 `json:"status_code"`

	Message string `json:"message"`
}

type ErrorInterceptor struct {
	logger *zap.Logger
}

func NewErrorInterceptor() (server.HandlerWrapper, error) {

	e := &ErrorInterceptor{
		logger: zlog.Logger,
	}

	return e.warpHandler, nil
}

func (e *ErrorInterceptor) warpHandler(handler server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := handler(ctx, req, rsp)
		if err != nil {

			s, ok := status.FromError(err)
			if ok {
				ferr := &FaceFrontError{
					StatusCode: uint32(s.Code()),
					Message:    s.Message(),
				}
				bs, _ := json.Marshal(ferr)
				return errors.New(string(bs))
			}

			e.logger.Info("get a invalid grpc error, please repaire", zap.String("invalid err", err.Error()))
			return err
		}

		return nil
	}
}