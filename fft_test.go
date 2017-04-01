package fft

import "fmt"
import "testing"

const N = 256

func TestLog2(t *testing.T) {
	var i int
	for i = 0; i < 10; i++ {
		println(i, Log2(i))
	}
}

func TestImpulse(t *testing.T) {
	xr := make([]float64, N)
	xi := make([]float64, N)
	for i := 0; i < N/2; i++ {
		xr[i] = 1.0
	}
	FFT(xr, xi)

	fmt.Printf("-----------------------------------\n")
	for i := 0; i < N; i++ {
		fmt.Printf("[%4d]  %10.5f  %10.5f\n", i, xr[i], xi[i])
	}
	fmt.Printf("-----------------------------------\n")
	InverseFFT(xr, xi)
	for i := 0; i < N; i++ {
		fmt.Printf("[%4d]  %10.5f  %10.5f\n", i, xr[i], xi[i])
	}
	fmt.Printf("-----------------------------------\n")
}
