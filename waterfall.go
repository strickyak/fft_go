// +build main

package main

import fft "github.com/strickyak/fft_go"
import (
	"flag"
	"fmt"
	"io"
	"math"
)

var N = flag.Int("n", 256, "size of FFT")

// printPPM prints the FFT waterfall as a PPM textfile ("plain PPM").
// See http://netpbm.sourceforge.net/doc/ppm.html#plainppm
// It would be nice to add better scaling & colors.
func printPPM(rr, ii [][]float64) {
	fmt.Printf("P3\n")
	fmt.Printf("  %d %d\n", *N, len(rr))
	fmt.Printf("    255\n")

	for x := 0; x < len(rr); x++ {
		for y := 0; y < *N; y++ {
			dist := math.Sqrt(rr[x][y]*rr[x][y] + ii[x][y]*ii[x][y])
			//z := int(rr[x][y])
			//if z < 0 {
			//z = -z
			//}
			z := (int(dist) * 255 / 30000) & 255
			fmt.Printf("%d %d %d   ", z, z, z)
		}
		fmt.Printf("\n")
	}
}

// main reads integer amplitude sequence from stdin,
// does FFTs of chunks of --n samples,
// and crudely prints the FFT waterfall as a Plain PPM file.
func main() {
	flag.Parse()
	var rr, ii [][]float64
	var done bool

	for !done {
		reals, imags := make([]float64, *N), make([]float64, *N)
		var i int

		for i = 0; i < *N; i++ {
			var sample int
			n, err := fmt.Scanf("%d", &sample)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				done = true
				break
			}

			reals[i], imags[i] = float64(sample), 0.0
		}

		for ; i < *N; i++ {
			reals[i], imags[i] = 0.0, 0.0
		}

		fft.FFT(reals, imags)
		rr, ii = append(rr, reals), append(ii, imags)
		for i = 0; i < *N; i++ {
			//println(i, reals[i], imags[i])
		}
	}

	printPPM(rr, ii)
}
