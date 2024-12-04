# Custom Logger with Configurable Options
This project is a proof of concept for implementing a custom logging utility in Go with customizable options. It demonstrates how to build a flexible logger that allows you to configure various logging options using functional options.
## Overview
The purpose of this project is to create a logger that can be easily configured with custom options for different logging levels and additional functionalities.
## Features
- **Multiple Log Levels:** Configure log messages to be displayed at different levels such as DEBUG, INFO, WARN, ERROR, and FATAL.
- **Customizable Options:** Modify options such as log level, statistics, and additional data through functional options.

## Usage
Below is a brief overview of how to utilize the custom logger in your application:
1. **Create an Instance:** You can create a new logger instance with default options or pass customized options.
```go
   log := logger.NewLogger()
   customLog := logger.NewLoggerWithOptions(logger.Options{Level: logger.DEBUG, Statistics: 5})
```
2. **Log Messages:** Use the logger to log messages at different levels. You can pass custom options for each log entry.
```go
   log.Info("This is an info message")
   log.Error("This is an error message", logger.WithData([]interface{}{"additional", "context"}))
```
3. **Setting Options:** You can adjust options dynamically using functional options:
```go
   log.Info("Sample message", logger.WithLevel(logger.WARN), logger.WithStatistics(10))
```
## Functional Options
- **WithLevel:** Set the logging level (e.g., DEBUG, INFO).
- **WithStatistics:** Adjust the probability or conditions under which a log entry is displayed.
- **WithData:** Attach additional data to the log message.

## Example
Here is an example of how to set up a logger and log messages with it:

```go
package main

import "github.com/uditalias/logger-go/logger"

func main() {
	log := logger.NewLogger()

	log.Debug("Debugging message")
	log.Info("Information message", logger.WithStatistics(10))
	log.Error("Error message", logger.WithData([]interface{}{"error", "details"}))
}
```

## License
This project is distributed under the MIT License. See `LICENSE` for more information.

## Contributing
Contributions are welcome. Feel free to submit a pull request or open an issue if you find a bug or have a feature request.