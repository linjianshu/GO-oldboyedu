package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

func main() {
	consoleLogger := NewConsoleLogger(ERROR)
	Logger.Debug(consoleLogger, "hello world")
	Logger.Info(consoleLogger, "hello world")
	Logger.Warning(consoleLogger, "hello world")
	Logger.Error(consoleLogger, "hello world%d", 10)
	Logger.Fatal(consoleLogger, "hello world")

	fileLogger := NewFileLogger(ERROR, ".", "logdemo.txt", "logdemoerr.txt", 128)
	Logger(fileLogger).Debug("hello world %d", 10)
	Logger(fileLogger).Error("fatal error %v", "某树被榨干了...")
}

type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warning(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}

type ConsoleLogger struct {
	LogLevel MODE
}

type MODE = int

const (
	DEBUG MODE = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func NewConsoleLogger(logLevel MODE) ConsoleLogger {
	return ConsoleLogger{
		LogLevel: logLevel,
	}
}

func (m ConsoleLogger) Debug(s string, a ...interface{}) {
	if m.LogLevel >= DEBUG {
		var msg string
		if a == nil {
			msg = s
		} else {
			msg = fmt.Sprintf(s, a)
		}
		fmt.Printf("[%v] [DEBUG] this is a debug log, value :%v\n", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Println(getInfo(2))
	}
}
func (m ConsoleLogger) Info(s string, a ...interface{}) {
	if m.LogLevel >= INFO {
		var msg string
		if a == nil {
			msg = s
		} else {
			msg = fmt.Sprintf(s, a)
		}
		fmt.Printf("[%v] [INFO] this is a info log, value :%v\n", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Println(getInfo(2))
	}
}
func (m ConsoleLogger) Warning(s string, a ...interface{}) {
	if m.LogLevel >= WARNING {
		var msg string
		if a == nil {
			msg = s
		} else {
			msg = fmt.Sprintf(s, a)
		}
		fmt.Printf("[%v] [WARNING] this is a warning log, value :%v\n", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Println(getInfo(2))
	}
}
func (m ConsoleLogger) Error(s string, a ...interface{}) {
	if m.LogLevel >= ERROR {
		var msg string
		if a == nil {
			msg = s
		} else {
			msg = fmt.Sprintf(s, a)
		}
		fmt.Printf("[%v] [ERROR] this is a error log, value :%v \n", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Println(getInfo(2))
	}
}
func (m ConsoleLogger) Fatal(s string, a ...interface{}) {
	if m.LogLevel >= FATAL {
		var msg string
		if a == nil {
			msg = s
		} else {
			msg = fmt.Sprintf(s, a)
		}
		fmt.Printf("[%v] [FATAL] this is a fatal log, value :%v\n", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Println(getInfo(2))
	}
}

func getInfo(n int) string {
	caller, file, line, ok := runtime.Caller(n)
	if !ok {
		return "error"
	}
	name := runtime.FuncForPC(caller).Name()
	return "fileLocation: " + path.Base(file) + ", method: " + name + ", line: " + strconv.Itoa(line)
}

type FileLogger struct {
	LogLevel    MODE
	fileName    string //日志文件的名称
	filePath    string //日志文件的路径
	errFileName string //错误日志单独记录
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

func NewFileLogger(logLevel MODE, filePath, fileName, errFileName string, maxFileSize int64) *FileLogger {
	file, err := os.OpenFile(path.Join(filePath, fileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	errfile, err := os.OpenFile(path.Join(filePath, errFileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	return &FileLogger{
		LogLevel:    logLevel,
		fileName:    fileName,
		filePath:    filePath,
		errFileName: errFileName,
		maxFileSize: maxFileSize,
		fileObj:     file,
		errFileObj:  errfile,
	}
}
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) checkSize(fileObj *os.File) bool {
	stat, err := fileObj.Stat()
	if err != nil {
		return false
	}
	if stat.Size() > f.maxFileSize {
		return true
	} else {
		return false
	}
}
func (f *FileLogger) SplitLogFile() {
	//需要切割文件
	//1.关闭当前文件
	f.fileObj.Close()
	//2.rename 备份一下 xx.log -> xx.log.bak201908031709
	nowStr := time.Now().Format("20060102150405000")
	bakFilePath := path.Join(f.filePath, f.fileName) + ".bak" + nowStr
	err := os.Rename(path.Join(f.filePath, f.fileName), bakFilePath)
	if err != nil {
		return
	}
	//3.打开一个新的日志文件
	newFile, err := os.OpenFile(path.Join(f.filePath, f.fileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	//4.将打开的新日志文件对象赋值给 f.fileObj
	f.fileObj = newFile

}
func (f *FileLogger) Debug(s string, a ...interface{}) {
	if f.LogLevel >= DEBUG {
		if f.checkSize(f.fileObj) {
			f.SplitLogFile()
		}
		msg := fmt.Sprintf(s, a)
		msg = fmt.Sprintf("[%v] [DEBUG] this is a debug log, value :%v", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Fprintln(f.fileObj)
		fmt.Fprintln(f.fileObj, msg)
		fmt.Fprintln(f.fileObj, getInfo(2))
		f.Close()
	}
}

func (f *FileLogger) Info(s string, a ...interface{}) {
	if f.LogLevel >= INFO {

		msg := fmt.Sprintf(s, a)
		msg = fmt.Sprintf("[%v] [INFO] this is a info log, value :%v", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Fprintln(f.fileObj)
		fmt.Fprintln(f.fileObj, msg)
		fmt.Fprintln(f.fileObj, getInfo(2))
		f.Close()
	}
}
func (f *FileLogger) Warning(s string, a ...interface{}) {
	if f.LogLevel >= WARNING {

		msg := fmt.Sprintf(s, a)
		msg = fmt.Sprintf("[%v] [WARNING] this is a warning log, value :%v", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Fprintln(f.fileObj)
		fmt.Fprintln(f.fileObj, msg)
		fmt.Fprintln(f.fileObj, getInfo(2))
		f.Close()
	}
}
func (f *FileLogger) Error(s string, a ...interface{}) {
	if f.LogLevel >= ERROR {

		msg := fmt.Sprintf(s, a)
		msg = fmt.Sprintf("[%v] [ERROR] this is a error log, value :%v", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Fprintln(f.fileObj)
		fmt.Fprintln(f.fileObj, msg)
		fmt.Fprintln(f.fileObj, getInfo(2))

		fmt.Fprintln(f.errFileObj)
		fmt.Fprintln(f.errFileObj, msg)
		fmt.Fprintln(f.errFileObj, getInfo(2))
		f.Close()
	}
}
func (f *FileLogger) Fatal(s string, a ...interface{}) {
	if f.LogLevel >= FATAL {

		msg := fmt.Sprintf(s, a)
		msg = fmt.Sprintf("[%v] [FATAL] this is a Fatal log, value :%v", time.Now().Format("2006-01-02 15-04-05"), msg)
		fmt.Fprintln(f.fileObj)
		fmt.Fprintln(f.fileObj, msg)
		fmt.Fprintln(f.fileObj, getInfo(2))

		fmt.Fprintln(f.errFileObj)
		fmt.Fprintln(f.errFileObj, msg)
		fmt.Fprintln(f.errFileObj, getInfo(2))
		f.Close()
	}
}
