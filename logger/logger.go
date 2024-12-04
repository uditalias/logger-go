package logger

import (
	"github.com/payleet/zlogger"
	"math/rand"
)

type Level int

type OptsFn func(*Options)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Options struct {
	Level      Level
	Statistics int
	Data       []interface{}
}

func defaultOptions() Options {
	return Options{
		Level:      INFO,
		Statistics: 1,
		Data:       nil,
	}
}

func WithStatistics(statistics int) OptsFn {
	return func(opts *Options) {
		opts.Statistics = statistics
	}
}

func WithLevel(level Level) OptsFn {
	return func(opts *Options) {
		opts.Level = level
	}
}

func WithData(data []interface{}) OptsFn {
	return func(opts *Options) {
		opts.Data = data
	}
}

func WithError(opts *Options) {
	opts.Level = ERROR
}

func WithFatal(opts *Options) {
	opts.Level = FATAL
}

func WithWarn(opts *Options) {
	opts.Level = WARN
}

func WithInfo(opts *Options) {
	opts.Level = INFO
}

func WithDebug(opts *Options) {
	opts.Level = DEBUG
}

type Logger interface {
	Log(string, ...OptsFn)
	Error(string, ...OptsFn)
	Fatal(string, ...OptsFn)
	Warn(string, ...OptsFn)
	Info(string, ...OptsFn)
	Debug(string, ...OptsFn)
}

type logger struct {
	options Options
}

func NewLogger() Logger {
	return logger{
		options: defaultOptions(),
	}
}

func NewLoggerWithOptions(options Options) Logger {
	return logger{
		options: options,
	}
}

func (l logger) Log(message string, opts ...OptsFn) {
	options := l.options // copy by value

	for _, fn := range opts {
		fn(&options)
	}

	if !isValidRandom(options.Statistics) {
		return
	}

	switch options.Level {
	case DEBUG:
		zlogger.Debug(message, options.Data)
	case ERROR:
		zlogger.Error(message, options.Data)
	case FATAL:
		zlogger.Fatal(message, options.Data)
	case INFO:
		zlogger.Info(message, options.Data)
	case WARN:
		// warn level is not supported by zlogger so we use error log
		zlogger.Error(message, options.Data)
	}
}

func (l logger) Error(message string, opts ...OptsFn) {
	l.Log(message, append(opts, WithError)...)
}

func (l logger) Fatal(message string, opts ...OptsFn) {
	l.Log(message, append(opts, WithFatal)...)
}

func (l logger) Warn(message string, opts ...OptsFn) {
	l.Log(message, append(opts, WithWarn)...)
}

func (l logger) Info(message string, opts ...OptsFn) {
	l.Log(message, append(opts, WithInfo)...)
}

func (l logger) Debug(message string, opts ...OptsFn) {
	l.Log(message, append(opts, WithDebug)...)
}

func isValidRandom(max int) bool {
	if max <= 0 {
		max = 1
	}

	return rand.Intn(max)+1 == max
}
