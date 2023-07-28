package util

import (
    "github.com/pkg/errors"
    "github.com/spf13/cast"
)

func IsValidPort[T int | int32 | uint | uint32 | uint64 | int64 | string](port T) error {
    p, err := cast.ToIntE(port)
    if err != nil {
        return errors.Wrap(err,"端口转换错误")
    }

    if p >= 0 && p < 65535 {
        return nil
    }
    return errors.Errorf("%d不是一个合格的[0-65535]端口", p)
}
