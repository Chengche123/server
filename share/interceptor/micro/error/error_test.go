package interceptor

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFoo(t *testing.T) {
	err := status.Error(codes.Unauthenticated, "密码错误")
	s, _ := status.FromError(err)
	ferr := &struct {
		Code    uint32 `json:"status_code"`
		Message string `json:"message"`
	}{
		Code:    uint32(s.Code()),
		Message: s.Message(),
	}

	bs, _ := json.Marshal(ferr)
	err = errors.New(string(bs))
	fmt.Printf("%s\n", err.Error())
}
