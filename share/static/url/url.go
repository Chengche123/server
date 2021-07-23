package static

import (
	"net/url"
	zlog "share/log/zap"

	"go.uber.org/zap"
)

func ConverURL(u string) string {
	u1, err := url.Parse(u)
	if err != nil {
		zlog.Logger.Info("invalid url", zap.String("url", u))
		return u
	}

	return "http://212.129.236.77/" + u1.Host + u1.Path
}
