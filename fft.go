// FFT does not work; it has some kind of error.
// Some data looks correct, but it does not round-trip.
//
// Based on BASIC code from
// The Scientist and Engineer's Guide to Digital Signal Processing
// By Steven W. Smith, Ph.D.
package fft

import "math"

func Log2(a int) int {
	var i int
	for a > 1 {
		a = a >> 1
		i++
	}
	return i
}

func FFT(rex, imx []float64) {
	n := len(rex)
	m := Log2(n)
	j := n / 2
	//println(n, m, j)

	for i := 1; i <= n/2; i++ {
		if i < j {
			rex[i], rex[j] = rex[j], rex[i]
			imx[i], imx[j] = imx[j], imx[i]
		}
		k := n / 2
		for k <= j {
			j = j - k
			k = k / 2
		}
		j = j + k
	}

	for l := 1; l <= m; l++ {
		le := 1 << uint(l)
		le2 := le / 2
		ur, ui := 1.0, 0.0
		sr := math.Cos(math.Pi / float64(le2))
		si := -math.Sin(math.Pi / float64(le2))

		for j := 1; j <= le2; j++ {
			jm1 := j - 1
			for i := jm1; i <= n-1; i += le {
				ip := i + le2
				tr := rex[ip]*ur - imx[ip]*ui
				ti := rex[ip]*ui + imx[ip]*ur
				rex[ip] = rex[i] - tr
				imx[ip] = imx[i] - ti
				rex[i] += tr
				imx[i] += ti
			}
			ur, ui = ur*sr-ui*si, ur*si+ui*sr
		}
	}
}

func InverseFFT(rex, imx []float64) {
	n := len(rex)
	nf := float64(n)

	for k := 0; k < n; k++ {
		imx[k] = -imx[k]
	}
	FFT(rex, imx)
	for i := 0; i < n; i++ {
		rex[i], imx[i] = rex[i]/nf, -imx[i]/nf
	}
}
