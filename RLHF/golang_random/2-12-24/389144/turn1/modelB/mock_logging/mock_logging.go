package mock_logging

import (
	"fmt"

	"go-programs/RLHF/golang_random/2-12-24/389144/turn1/modelB/logging"

	"github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	rec      *MockLoggerMockRecorder
	logLevel logging.Level
	msgs     []string
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl, rec: &MockLoggerMockRecorder{mock: mock}}
	mock.logLevel = logging.DebugLevel // Default log level to Debug
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.rec
}

// SetLogLevel sets the log level of the mock logger
func (m *MockLogger) SetLogLevel(level logging.Level) {
	m.logLevel = level
}

// AddLogMessage records a log message in the mock
func (m *MockLogger) AddLogMessage(msg string) {
	m.msgs = append(m.msgs, msg)
}

// Debugf logs a debug message (captured in mock)
func (m *MockLogger) Debugf(format string, args ...interface{}) {
	if m.logLevel <= logging.DebugLevel {
		msg := fmt.Sprintf(format, args...)
		m.AddLogMessage(msg)
	}
}

// Other log level implementations omitted for brevity
