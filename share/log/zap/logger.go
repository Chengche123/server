package log

import "go.uber.org/zap"

var Logger = func() *zap.Logger {
	r, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return r
}()
