package loggerfactory

import (
	"github.com/pkg/errors"

	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container/loggerfactory/zap"
)

// receiver for zap factory
type ZapFactory struct{}

// build zap logger
func (mf *ZapFactory) Build(lc *config.LogConfig) error {
	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
