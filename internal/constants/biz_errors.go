// SPDX-License-Identifier: MIT
// Copyright (C) 2023 WuPeng <wup364@outlook.com>.

package constants

import (
	"errors"

	"github.com/wup364/pakku/pkg/utypes"
)

// ErrUnknownError 未知错误
var ErrUnknownError = utypes.NewCustomError(errors.New("unknown error"), "SERVER_ERROR")
