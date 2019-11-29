
package neunetlib

import (
    "math"
)

func New() n *Network {
    n.Weights = []
    n.Biases = []
    return
}

type Network struct {
    Weights []float64
    Biases []float64
}

func (n *Network) feedforward(){
    
}

func sigmoid(z float64) float64 {
    return 1.0 / 1.0 + math.Exp(-z)
}