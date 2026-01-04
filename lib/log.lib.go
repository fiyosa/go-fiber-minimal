package lib

import (
	"encoding/json"
	"fmt"
	"go-fiber-minimal/config"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	LogFile    logFileManager
	LogConsole logConsoleManager
)

type logFileManager struct {
	logger *log.Logger
}

type logConsoleManager struct {
	logger *log.Logger
}

// Init must be called to configure both loggers
func (l logFileManager) Init() {
	dirLogs := "./logs"

	// ================== Log Console ==================

	// Setup LogConsole - writes to console only
	LogConsole.logger = log.New(os.Stdout, "", 0)

	// Setup LogConsole - writes to console AND file
	// multiWriter := io.MultiWriter(os.Stdout, file)
	// LogConsole.logger = log.New(multiWriter, "", log.LstdFlags)

	// =================================================

	if config.Env.APP_ENV == "production" {
		return
	}

	// Create logs directory if it doesn't exist
	_, err := os.Stat(dirLogs)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirLogs, 0755)
		if err != nil {
			fmt.Printf("Error creating directory logs: %v \n\n", err.Error())
			os.Exit(1)
		}
	}

	currentDate := time.Now().Format("2006-01-02")
	// Open log file
	file, err := os.OpenFile(dirLogs+"/fiber_"+currentDate+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v \n\n", err.Error())
		os.Exit(1)
	}
	// Setup LogFile - writes only to file
	LogFile.logger = log.New(file, "", 0)

}

// Helper to formatting log arguments to JSON if complex type
func parseContents(v ...any) []any {
	args := make([]any, len(v))
	for i, arg := range v {
		switch val := arg.(type) {
		case string:
			args[i] = val
		case error:
			args[i] = val.Error()
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
			args[i] = val
		default:
			// Convert maps, structs, slices to JSON string
			b, err := json.MarshalIndent(val, "", "  ")
			if err != nil {
				args[i] = val
			} else {
				args[i] = string(b)
			}
		}
	}
	return args
}

func formatLog(level string, v ...any) []any {
	dt := time.Now().Format("2006-01-02 15:04:05")

	fileInfo := "unknown:0"
	// Mencari caller yang bukan dari file ini (log.lib.go)
	for i := 1; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && filepath.Base(file) != "log.lib.go" {
			fileInfo = fmt.Sprintf("%s:%d", filepath.Base(file), line)
			break
		}
	}

	prefix := fmt.Sprintf("[%s] %s %s", level, dt, fileInfo)
	return append([]any{prefix}, parseContents(v...)...)
}

// LogFile methods - writes only to file
func (l logFileManager) Info(v ...any) {
	l.logger.Println(formatLog("INFO", v...)...)
}

func (l logFileManager) Error(v ...any) {
	l.logger.Println(formatLog("ERROR", v...)...)
}

func (l logFileManager) Warn(v ...any) {
	l.logger.Println(formatLog("WARN", v...)...)
}

func (l logFileManager) Debug(v ...any) {
	l.logger.Println(formatLog("DEBUG", v...)...)
}

func (l logFileManager) Fatal(v ...any) {
	l.logger.Fatalln(formatLog("FATAL", v...)...)
}

func (l logFileManager) Panic(v ...any) {
	l.logger.Panicln(formatLog("PANIC", v...)...)
}

// LogConsole methods - writes to console and file
func (l logConsoleManager) Info(v ...any) {
	l.logger.Println(formatLog("INFO", v...)...)
}

func (l logConsoleManager) Error(v ...any) {
	l.logger.Println(formatLog("ERROR", v...)...)
}

func (l logConsoleManager) Warn(v ...any) {
	l.logger.Println(formatLog("WARN", v...)...)
}

func (l logConsoleManager) Debug(v ...any) {
	l.logger.Println(formatLog("DEBUG", v...)...)
}

func (l logConsoleManager) Fatal(v ...any) {
	l.logger.Fatalln(formatLog("FATAL", v...)...)
}

func (l logConsoleManager) Panic(v ...any) {
	l.logger.Panicln(formatLog("PANIC", v...)...)
}
