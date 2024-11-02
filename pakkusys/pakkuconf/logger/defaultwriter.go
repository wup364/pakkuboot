// Copyright (C) 2023 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package logger

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/strutil"
)

// NewLoggerWriter4File 实例化日志持久化对象
func NewLoggerWriter4File(logdir, logName string) io.Writer {
	return &loggerWriter4File{
		logdir:       logdir,
		logName:      logName,
		locker:       new(sync.Mutex),
		lastDateTime: time.Now().Format("2006-01-02"),
	}
}

// loggerWriter4File 本地日志文件写入
type loggerWriter4File struct {
	logdir        string
	logName       string
	lastDateTime  string
	locker        sync.Locker
	currentWriter io.WriteCloser
}

// Write implements io.Writer.
func (lw *loggerWriter4File) Write(p []byte) (n int, err error) {
	return lw.getWriter().Write(p)
}

func (lw *loggerWriter4File) getWriter() io.Writer {
	lw.locker.Lock()
	defer lw.locker.Unlock()

	// 第一次初始化
	if nil == lw.currentWriter {
		logName := lw.logName + "_" + lw.lastDateTime + ".log"
		if currentWriter, err := fileutil.GetWriter(strutil.Parse2UnixPath(lw.logdir + "/" + logName)); nil != err {
			panic(err)
		} else {
			lw.currentWriter = currentWriter
		}
	}

	// 日期变化
	dateTime := time.Now().Format("2006-01-02")
	if lw.lastDateTime != dateTime {
		retry := 100
		var err error
		for err = lw.currentWriter.Close(); nil != err && retry > 0; retry-- {
			fmt.Println(err.Error())
			time.Sleep(time.Millisecond * 10)
		}

		// 只有关闭成功才会换下一个文件
		if err == nil {
			logName := lw.logName + "_" + dateTime + ".log"
			if currentWriter, err := fileutil.GetWriter(strutil.Parse2UnixPath(lw.logdir + "/" + logName)); nil != err {
				fmt.Println(err.Error())
			} else {
				lw.lastDateTime = dateTime
				lw.currentWriter = currentWriter
			}
		}
	}

	return lw.currentWriter
}
