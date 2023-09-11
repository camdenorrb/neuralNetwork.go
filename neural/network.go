package neural

type Network struct {
	InputNodes   uint64
	HiddenNodes  uint64
	OutputNodes  uint64
	LearningRate mat.Dense
}
