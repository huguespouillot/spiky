package nodes

import (
	"errors"
	"spiky/pkg/core"
	"spiky/pkg/edges"
)

type baseNode struct {
	id        string
	potential float64
	position  core.Point
	spikes    map[core.Time]bool

	kernel    core.Kernel
	synapses  []core.Edge
	dendrites []core.Edge
}

func (n *baseNode) GetId() string {
	return n.id
}

func (n *baseNode) GetPosition() core.Point {
	return n.position
}

func (n *baseNode) Connect(target core.Node) core.Edge {
	edge := edges.New(n, target)
	n.AddSynapse(edge)
	target.AddDendrite(edge)
	return edge
}

func (node *baseNode) Compute(time core.Time, queue *core.Queue) {
	spiked := node.kernel.Compute(node, time)
	if spiked {
		for _, syn := range node.synapses {
			queue.Add(time+core.Time(syn.GetDelay()), syn.GetTarget())
		}
	}
}

func (n *baseNode) GetSynapses() []core.Edge {
	return n.synapses
}

func (n *baseNode) GetDendrites() []core.Edge {
	return n.dendrites
}

func (n *baseNode) SetSpike(time core.Time, spiked bool) {
	n.spikes[time] = spiked
}

func (n *baseNode) GetSpike(time core.Time) bool {
	return n.spikes[time]
}

func (n *baseNode) GetSpikeRate(startTime core.Time, endTime core.Time) (float64, error) {
	if startTime >= endTime {
		return 0.0, errors.New("invalid time range")
	}
	spikeCount := 0.0
	for _, v := range n.spikes {
		if v {
			spikeCount++
		}
	}
	return spikeCount / float64(endTime-startTime), nil
}

func (n *baseNode) GetChildren() []core.Node {
	var slice = make([]core.Node, len(n.synapses))
	for i, syn := range n.synapses {
		slice[i] = syn.GetTarget()
	}
	return slice
}

func (n *baseNode) GetParents() []core.Node {
	var slice = make([]core.Node, len(n.dendrites))
	for i, syn := range n.dendrites {
		slice[i] = syn.GetSource()
	}
	return slice
}

func (n *baseNode) AddSynapse(edge core.Edge) {
	n.synapses = append(n.synapses, edge)
}

func (n *baseNode) AddDendrite(edge core.Edge) {
	n.dendrites = append(n.synapses, edge)
}
