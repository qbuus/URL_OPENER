package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	pathOfFile    *os.File
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
)

func init() {
	year, month, day := time.Now().Date()

	err := os.MkdirAll("logger/logs", 0755)
	if err != nil {
		fmt.Println("[ERROR] creating logs directory:", err)
		return
	}

	pathOfLogFile, err := filepath.Abs("./logger/logs/" + fmt.Sprintf("%d-%02d-%02d.log", year, month, day))
	if err != nil {
		fmt.Println("[ERROR] getting absolute path:", err)
		return
	}

	pathOfFile, err := os.OpenFile(pathOfLogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("[ERROR] opening log file:", err)
		return
	}

	InfoLogger = log.New(pathOfFile, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(pathOfFile, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(pathOfFile, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(info ...any) {
	InfoLogger.Println(info...)
	fmt.Println(info...)
}

func InfoF(format string, info ...any) {
	InfoLogger.Printf(format, info...)
	fmt.Printf(format, info...)
}

func Warning(info ...any) {
	WarningLogger.Println(info...)
	fmt.Println(info...)
}

func WarningF(format string, info ...any) {
	WarningLogger.Printf(format, info...)
	fmt.Printf(format, info...)
}

func Error(info ...any) {
	ErrorLogger.Println(info...)
	fmt.Println(info...)
}

func ErrorF(format string, info ...any) {
	ErrorLogger.Printf(format, info...)
	fmt.Printf(format, info...)
}
