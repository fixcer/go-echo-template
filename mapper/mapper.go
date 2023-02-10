package mapper

import (
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/mitchellh/mapstructure"
	"go-backend-template/pkg/logging"
	"reflect"
	"time"
)

// ToDateHookFunc converts time.Time to *time.Time
func ToDateHookFunc() mapstructure.DecodeHookFunc {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		logging.Log.Info("From: ", from.Type(), " To: ", to.Type(), " Value: ", from.Interface())
		if from.Type() == reflect.TypeOf(&time.Time{}) && to.Type() == reflect.TypeOf(&types.Date{
			Time: time.Time{},
		}) {
			logging.Log.Debug("Converting time.Time to *time.Time")
			return &types.Date{
				Time: *from.Interface().(*time.Time),
			}, nil
		}

		return from.Interface(), nil
	}
}
