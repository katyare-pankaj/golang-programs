package modelA

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type testOut struct {
	bytes.Buffer
}

func (t *testOut) Write(b []byte) (int, error) {
	return t.Buffer.Write(b)
}

func TestInitLogger(t *testing.T) {
	InitLogger()
	assert.NotNil(t, logger, "Logger should not be nil")
}

func TestInfo(t *testing.T) {
	testWriter := new(testOut)
	zerolog.SetGlobalOutput(testWriter)
	defer zerolog.SetGlobalOutput(os.Stdout)

	fields := map[string]interface{}{
		"userId":  123,
		"request": "GET /api/users",
	}

	Info("User requested data", fields)

	expectedOutput := `{"level":"info","time":"2023-10-04T15:12:34Z","caller":"logger_test.go:44","msg":"User requested data","userId":123,"request":"GET /api/users"}`
	assert.Equal(t, expectedOutput, testWriter.String())
}

func TestDebug(t *testing.T) {
	testWriter := new(testOut)
	zerolog.SetGlobalOutput(testWriter)
	defer zerolog.SetGlobalOutput(os.Stdout)

	fields := map[string]interface{}{
		"data": "Some debug information",
	}

	Debug("Debug message", fields)

	expectedOutput := `{"level":"debug","time":"2023-10-04T15:12:34Z","caller":"logger_test.go:54","msg":"Debug message","data":"Some debug information"}`
	assert.Equal(t, expectedOutput, testWriter.String())
}

func TestError(t *testing.T) {
	testWriter := new(testOut)
	zerolog.SetGlobalOutput(testWriter)
	defer zerolog.SetGlobalOutput(os.Stdout)

	err := fmt.Errorf("Something went wrong")
	fields := map[string]interface{}{
		"status": 500,
	}

	Error(err, "Internal server error", fields)

	expectedOutput := `{"level":"error","time":"2023-10-04T15:12:34Z","caller":"logger_test.go:64","msg":"Internal server error","error":"Something went wrong","status":500}`
	assert.Equal(t, expectedOutput, testWriter.String())
}
