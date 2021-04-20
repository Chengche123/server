package controler

import (
	"context"
	"encoding/json"
	"net/http"

	authpb "micro/auth/api/gen/v1"

	api "github.com/micro/go-micro/v2/api/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

type Service struct {
	Grpc   authpb.AuthService
	Looger *zap.Logger
}

func (s *Service) Login(ctx context.Context, req *api.Request, res *api.Response) error {
	res.StatusCode = 200
	grpcRes, _ := s.Grpc.Login(ctx, &authpb.LoginRequest{
		Code: req.Get["code"].Values[0],
	})
	resp, _ := json.Marshal(&grpcRes)
	res.Body = string(resp)
	return nil
}

func (s *Service) UserLogin(ctx context.Context, req *api.Request, res *api.Response) error {
	// 仅支持POST方法
	if req.Method != "POST" {
		res.StatusCode = http.StatusMethodNotAllowed
		return nil
	}

	// 从body中读取参数
	var rpcReq authpb.UserLoginRequest
	err := json.Unmarshal([]byte(req.Body), &rpcReq)
	if err != nil {
		res.StatusCode = http.StatusBadRequest
		return nil
	}

	// 用grpc调用微服务
	rpcRes, err := s.Grpc.UserLogin(ctx, &rpcReq)
	if err != nil {
		res.StatusCode = http.StatusBadRequest

		st, _ := status.FromError(err)
		res.Body = st.Message()

		return nil
	}

	// 将rpc响应转换成http响应
	raw, err := json.Marshal(rpcRes)
	if err != nil {
		s.Looger.Error("failed to marshal rpc response", zap.Error(err))

		res.StatusCode = http.StatusInternalServerError
		return nil
	}
	res.StatusCode = http.StatusOK
	res.Body = string(raw)

	return nil
}
