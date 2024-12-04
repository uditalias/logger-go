package main

import "github.com/uditalias/logger-go/logger"

func main() {

	appLogger := logger.NewLogger()

	appLogger.Log("With default options")

	appLogger.Log(
		"Hello World with statistics",
		logger.WithStatistics(100),
	)

	appLogger.Log(
		"Hello World with log level info and statistics",
		logger.WithLevel(logger.INFO),
		logger.WithStatistics(10),
	)

	appLogger.Error("This is an error log", logger.WithStatistics(100))

	appLogger.Fatal("This is a fatal log with data", logger.WithData(
		[]interface{}{
			"key1", "value1",
			"key2", "value2",
		}))

	appLogger2 := logger.NewLoggerWithOptions(logger.Options{
		Level:      logger.DEBUG,
		Statistics: 100,
	})

	appLogger2.Log("Hello World with custom options")
}
