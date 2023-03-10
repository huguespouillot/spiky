package monitoring

import (
	"spiky/pkg/core"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockedLayer struct {
	mock.Mock
}

func (ml *MockedLayer) Visit(fn func(int, *core.Neuron)) {
	ml.Called(fn)
}

func (ml *MockedLayer) Size() int {
	return 10
}

func TestSpikeMonitor(t *testing.T) {
	layer := new(MockedLayer)
	layer.On("Visit", mock.Anything).Return()
	monitor := NewSpikeMonitor(layer, 10)
	ticker := time.NewTicker(30 * time.Millisecond)
	monitor.Open(ticker.C)
	duration := 65 * time.Millisecond
	time.Sleep(duration)
	layer.AssertNumberOfCalls(t, "Visit", 2)
	ticker.Stop()
}
