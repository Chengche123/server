package controler

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type ParamSet struct {
	Logger *zap.Logger

	CodeResolver   CodeResolver
	OpenIDResolver OpenIDResolver

	TokenExpireIn  time.Duration
	TokenGenerator TokenGenerator

	UserResolver UserResolver
}

func (p *ParamSet) check() error {
	if p.Logger == nil {
		p.Logger, _ = zap.NewDevelopment()
	}

	if p.CodeResolver == nil {
		return fmt.Errorf("invalid CodeResolver: nil")
	}

	if p.OpenIDResolver == nil {
		return fmt.Errorf("invalid OpenIDResolver: nil")
	}

	if p.TokenExpireIn == 0 {
		return fmt.Errorf("invalid TokenExpireIn: 0")
	}

	if p.TokenGenerator == nil {
		return fmt.Errorf("invalid TokenGenerator: nil")
	}

	if p.UserResolver == nil {
		return fmt.Errorf("invalid UserResolver: nil")
	}

	return nil
}

func (p *ParamSet) Build() (*service, error) {
	err := p.check()
	if err != nil {
		return nil, fmt.Errorf("cannot open auth service: %v", err)
	}

	return &service{
		logger: p.Logger,

		codeResolver:   p.CodeResolver,
		openIDResolver: p.OpenIDResolver,

		tokenExpireIn:  p.TokenExpireIn,
		tokenGenerator: p.TokenGenerator,

		userResolver: p.UserResolver,
	}, nil
}
