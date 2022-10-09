package log

import (
	"github.com/hadenlabs/terraform-aws-lambda-function/config"
	"github.com/hadenlabs/terraform-aws-lambda-function/internal/common/log/provider"
	"github.com/hadenlabs/terraform-aws-lambda-function/internal/errors"
)

// New initialize a new Log.
func NewLog(conf config.Config) TracingLogger {
	return Factory(conf)
}

// Factory Log.
func Factory(conf config.Config) (prov TracingLogger) {
	switch conf.Log.Provider {
	case "zap":
		prov = provider.NewZap(conf)
	default:
		panic(errors.Errorf(errors.ErrorParseConfig, "unsupported email provider: %s", conf.Log.Provider))
	}
	return prov
}
