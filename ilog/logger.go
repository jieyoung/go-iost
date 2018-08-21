package ilog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type message struct {
	content string
	level   Level
}

type Logger struct {
	callDepth   int
	lowestLevel Level
	isRunning   int32
	wg          *sync.WaitGroup

	writers   []LogWriter
	msg       chan *message
	flush     chan *sync.WaitGroup
	syncWrite bool

	bufPool *BufPool

	quitCh chan struct{}
}

func New() *Logger {
	return &Logger{
		callDepth:   1,
		lowestLevel: LevelFatal,
		wg:          new(sync.WaitGroup),
		msg:         make(chan *message, 4096),
		flush:       make(chan *sync.WaitGroup, 1),
		bufPool:     NewBufPool(),
		quitCh:      make(chan struct{}, 1),
		syncWrite:   true,
	}
}

func NewConsoleLogger() *Logger {
	logger := New()
	consoleWriter := NewConsoleWriter()
	logger.AddWriter(consoleWriter)
	return logger
}

func (logger *Logger) Start() {
	if !atomic.CompareAndSwapInt32(&logger.isRunning, 0, 1) {
		return
	}
	if len(logger.writers) == 0 {
		fmt.Println("logger's writers is empty.")
		return
	}

	logger.wg.Add(1)
	go func() {
		defer func() {
			atomic.StoreInt32(&logger.isRunning, 0)
			logger.cleanMsg()
			for _, writer := range logger.writers {
				writer.Flush()
				writer.Close()
			}
			logger.wg.Done()
		}()

		for {
			select {
			case <-logger.quitCh:
				return
			case msg := <-logger.msg:
				logger.write(msg)
			case wg := <-logger.flush:
				logger.cleanMsg()
				logger.flushWriters()
				wg.Done()
			}
		}
	}()
}

func (logger *Logger) Stop() {
	if !atomic.CompareAndSwapInt32(&logger.isRunning, 1, 0) {
		return
	}
	logger.quitCh <- struct{}{}
	logger.wg.Wait()
}

func (logger *Logger) SetCallDepth(d int) {
	logger.callDepth = d
}

func (logger *Logger) AddWriter(writer LogWriter) error {
	if err := writer.Init(); err != nil {
		return err
	}
	logger.writers = append(logger.writers, writer)
	if logger.lowestLevel > writer.GetLevel() {
		logger.lowestLevel = writer.GetLevel()
	}
	return nil
}

func (logger *Logger) Flush() {
	if atomic.LoadInt32(&logger.isRunning) == 0 {
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	select {
	case logger.flush <- wg:
		wg.Wait()
	default:
	}
}

func (logger *Logger) AsyncWrite() {
	logger.syncWrite = false
}

func (logger *Logger) write(msg *message) {
	wg := &sync.WaitGroup{}
	for _, writer := range logger.writers {
		if msg.level < writer.GetLevel() {
			continue
		}
		wg.Add(1)
		go func(lw LogWriter) {
			lw.Write(msg.content, msg.level)
			wg.Done()
		}(writer)
	}
	wg.Wait()

}

func (logger *Logger) flushWriters() {
	wg := &sync.WaitGroup{}
	for _, writer := range logger.writers {
		wg.Add(1)
		go func(lw LogWriter) {
			lw.Flush()
			wg.Done()
		}(writer)
	}
	wg.Wait()
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	logger.genMsg(LevelDebug, fmt.Sprintf(format, v...))
}

func (logger *Logger) Info(format string, v ...interface{}) {
	logger.genMsg(LevelInfo, fmt.Sprintf(format, v...))
}

func (logger *Logger) Warn(format string, v ...interface{}) {
	logger.genMsg(LevelWarn, fmt.Sprintf(format, v...))
}

func (logger *Logger) Error(format string, v ...interface{}) {
	logger.genMsg(LevelError, fmt.Sprintf(format, v...))
}

func (logger *Logger) Fatal(format string, v ...interface{}) {
	logger.genMsg(LevelFatal, fmt.Sprintf(format, v...))
	logger.Stop()
	os.Exit(1)
}

func (logger *Logger) genMsg(level Level, log string) {
	if level < logger.lowestLevel {
		return
	}
	if atomic.LoadInt32(&logger.isRunning) == 0 {
		return
	}
	buf := logger.bufPool.Get()
	defer logger.bufPool.Release(buf)

	buf.Write(levelBytes[level])
	buf.WriteString(" ")
	buf.WriteString(time.Now().Format("2006-01-02 15:04:05.000"))
	buf.WriteString(" ")
	buf.WriteString(location(logger.callDepth + 3))
	buf.WriteString(" ")
	buf.WriteString(log)
	buf.WriteString("\n")

	select {
	case logger.msg <- &message{buf.String(), level}:
		// default:
	}

	if logger.syncWrite {
		logger.Flush()
	}
}

func (logger *Logger) cleanMsg() {
	for {
		select {
		case msg := <-logger.msg:
			logger.write(msg)
		default:
			return
		}
	}
}

func location(deep int) string {
	_, file, line, ok := runtime.Caller(deep)
	if !ok {
		file = "???"
		line = 0
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}