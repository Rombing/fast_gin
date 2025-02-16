package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
)

type MyLog struct {
}

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogrus() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(MyLog{})
	logrus.AddHook(&MyHook{
		logPath: "logs",
	})
}

type MyHook struct {
	file    *os.File // 当前打开日志文件
	errFile *os.File // 当前打开错误日志文件
	date    string   // 当前日志时间
	logPath string   // 日志路径
	mu      sync.Mutex
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {

	hook.mu.Lock()
	defer hook.mu.Unlock()
	//1.写入到文件
	msg, _ := entry.String()
	date := entry.Time.Format("2006-01-02")

	if hook.date != date {
		hook.rotateFiles(date)
		hook.date = date
	}
	if entry.Level <= logrus.ErrorLevel {
		hook.errFile.Write([]byte(msg))
	}
	hook.file.Write([]byte(msg))
	//2.按时间分片
	//3.错误日志单独存放
	return nil
}

func (hook *MyHook) rotateFiles(timer string) error {
	if hook.file != nil {
		hook.file.Close()
	}
	if hook.file == nil {
		logDir := fmt.Sprintf("%s/%s", hook.logPath, timer)
		os.MkdirAll(logDir, 0666)
		logPath := fmt.Sprintf("%s/info.log", logDir)
		file, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		hook.file = file

		errLogPath := fmt.Sprintf("%s/error.log", logDir)
		errFile, _ := os.OpenFile(errLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		hook.errFile = errFile
	}
	return nil
}

func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
