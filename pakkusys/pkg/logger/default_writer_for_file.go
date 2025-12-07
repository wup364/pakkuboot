// SPDX-License-Identifier: MIT
// Copyright (C) 2023 WuPeng <wup364@outlook.com>.

// 默认日志写入文件实现
package logger

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/wup364/pakku/pkg/fileutil"
	"github.com/wup364/pakku/pkg/strutil"
)

const (
	logFileSuffix = ".log"
)

// NewLoggerWriter4File 实例化日志持久化对象
func NewLoggerWriter4File(logdir, logName string, maxLogFiles int) io.Writer {
	return &loggerWriter4File{
		logdir:       logdir,
		logName:      logName,
		maxLogFiles:  maxLogFiles,
		locker:       new(sync.Mutex),
		lastDateTime: time.Now().Format("2006-01-02"),
	}
}

// loggerWriter4File 本地日志文件写入
type loggerWriter4File struct {
	logdir        string
	logName       string
	maxLogFiles   int
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
		logName := lw.logName + "-" + lw.lastDateTime + logFileSuffix
		if currentWriter, err := fileutil.GetWriter(strutil.Parse2UnixPath(lw.logdir + "/" + logName)); err != nil {
			panic(err)
		} else {
			lw.currentWriter = currentWriter
		}
	}

	// 日期变化
	dateTime := time.Now().Format("2006-01-02")
	if lw.lastDateTime != dateTime {
		// 只有关闭成功才会换下一个文件
		if err := lw.currentWriter.Close(); err == nil {
			logName := lw.logName + "-" + dateTime + logFileSuffix
			if currentWriter, err := fileutil.GetWriter(strutil.Parse2UnixPath(lw.logdir + "/" + logName)); err != nil {
				fmt.Println(err.Error())
			} else {
				lw.lastDateTime = dateTime
				lw.currentWriter = currentWriter

				// 清理旧日志文件，如果数量超过限制
				go lw.cleanOldLogFiles()
			}
		}
	}

	return lw.currentWriter
}

// cleanOldLogFiles 清理旧的日志文件，保持最多15个日志文件
func (lw *loggerWriter4File) cleanOldLogFiles() {
	var err error

	// 获取日志文件目录中的所有日志文件
	var dirEntry []fs.DirEntry
	if dirEntry, err = os.ReadDir(lw.logdir); err != nil {
		fmt.Println("Error reading log directory:", err)
		return
	}

	// 筛选出所有日志文件，按修改时间排序
	var logFiles []os.FileInfo
	for _, file := range dirEntry {
		if !file.IsDir() && strings.HasPrefix(file.Name(), lw.logName) && strings.HasSuffix(file.Name(), logFileSuffix) {
			if f, err := file.Info(); nil == err {
				logFiles = append(logFiles, f)
			} else {
				fmt.Println("Error reading log directory:", err)
				return
			}
		}
	}

	// 如果文件数超过 15 个，则删除最旧的文件
	if len(logFiles) > lw.maxLogFiles {
		sort.Slice(logFiles, func(i, j int) bool {
			return logFiles[i].ModTime().Before(logFiles[j].ModTime())
		})

		// 删除最旧的文件
		for _, file := range logFiles[:len(logFiles)-lw.maxLogFiles] {
			if err := os.Remove(filepath.Join(lw.logdir, file.Name())); err != nil {
				fmt.Println("Error deleting old log file:", err)
			}
		}
	}
}
