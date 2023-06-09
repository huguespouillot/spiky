package core

type Neuron struct {
	id        string
	potential float64
	spikeTime *float64
	synapses  []*Edge
	dendrites []*Edge
}

func (neuron *Neuron) GetSpikeTime() *float64 {
	return neuron.spikeTime
}

func (neuron *Neuron) Fire(world *World) {
	// var currentTime *float64 =
	neuron.spikeTime = new(float64)
	*neuron.spikeTime = world.GetTime()
	for _, syn := range neuron.synapses {
		syn.Forward(world)
	}
	neuron.potential = 0
	world.markDirty(neuron)
}

func (neuron *Neuron) Receive(world *World, signal float64) {
	if neuron.spikeTime != nil {
		return
	}
	neuron.potential = neuron.potential + signal
	if neuron.potential >= world.Const.Threshold {
		neuron.Fire(world)
	}
	world.markDirty(neuron)
}

func (neuron *Neuron) Adjust(world *World, err float64) {
	for _, dend := range neuron.dendrites {
		dend.Adjust(world, err)
	}
}

func (n *Neuron) Reset() {
	n.potential = 0
	n.spikeTime = nil
}

func NewNeuron(id string) *Neuron {
	return &Neuron{
		id:        id,
		potential: 0.0,
		spikeTime: nil,
		synapses:  []*Edge{},
		dendrites: []*Edge{},
	}
}
