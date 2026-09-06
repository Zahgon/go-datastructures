package optimization

type pbs []*vertexProbabilityBundle

type vertexProbabilityBundle struct {
	probability float64
	vertex      *nmVertex
}

func calculateVVP(guess, vertex *nmVertex, sigma float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calculateSigma(dimensions, guesses int) float64 { _ = "STUB: not implemented"; return 0 }

func (pbs pbs) calculateProbabilities(bestGuess *nmVertex, sigma float64) {
	_ = "STUB: not implemented"
	return
}

func (pbs pbs) sort() { _ = "STUB: not implemented"; return }

func (pbs pbs) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pbs pbs) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pbs pbs) Len() int { _ = "STUB: not implemented"; return 0 }

type results struct {
	vertices vertices

	config NelderMeadConfiguration

	pbs pbs
}

func (results *results) search(result *nmVertex) int { _ = "STUB: not implemented"; return 0 }

func (results *results) exists(result *nmVertex, hint int) bool {
	_ = "STUB: not implemented"
	return false
}

func (results *results) insert(vertex *nmVertex) { _ = "STUB: not implemented"; return }

func (results *results) grab(num int) vertices { _ = "STUB: not implemented"; return *new(vertices) }

func (results *results) reSort(vertex *nmVertex) { _ = "STUB: not implemented"; return }

func newResults(guess *nmVertex, config NelderMeadConfiguration, num int) *results {
	_ = "STUB: not implemented"
	return nil
}
