package core

import (
	"os"
	"reflect"
	"spiky/pkg/codec"
	"spiky/pkg/utils"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	logrus.SetLevel(logrus.ErrorLevel)
	code := m.Run()
	os.Exit(code)
}

func TestModelCreation(t *testing.T) {
	csts := utils.NewDefaultConstants()
	output_size := 2
	codec := codec.NewLatencyCodec(255, csts)
	input := NewLayer("Input", 2)
	output := NewLayer("Output", output_size)
	DenseConnection(input, output, csts)
	model := NewModel(codec, []*Layer{input, output}, csts)
	model.Encode([]float64{255, 255})
	model.Run()
	prediction := model.Decode()
	assert.Equal(t, len(prediction), output_size, "Invalid prediction size")
	assert.Equal(t, reflect.TypeOf(prediction[0]).Kind(), reflect.Float64)
}

func TestModelVisit(t *testing.T) {
	csts := utils.NewDefaultConstants()
	input := NewLayer("Input", 2)
	hidden := NewLayer("Hidden", 10)
	output := NewLayer("Output", 10)
	codec := codec.NewLatencyCodec(255, csts)
	DenseConnection(input, hidden, csts)
	DenseConnection(hidden, output, csts)
	model := NewModel(codec, []*Layer{input, hidden, output}, csts)
	visitCount := 0
	model.Visit(func(n *Neuron) {
		visitCount++
	})
	assert.Equal(t, 22, visitCount, "Invalid neuron visit")
}
