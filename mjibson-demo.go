// +build main

package main

import (
	"fmt"
	"os"

	"github.com/mjibson/go-dsp/fft"
)
import "github.com/strickyak/canvas"

func plotReal(x []float64, filename string) {
	n := len(x)
	c := canvas.NewCanvasWithScale(512, 512, 0, float64(n), -5.0, 5.0)
	for i, e := range x {
		c.SSet(float64(i), e, canvas.Green)
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	c.WritePng(f)
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func plotComplex(vec []complex128, filename string) {
	n := len(vec)
	c := canvas.NewCanvasWithScale(512, 512, 0, float64(n), -5.0, 5.0)
	//var px, py int
	for i, e := range vec {
		//x, y := c.UnScale(real(e), imag(e))
		//if i > 0 {
		//c.PaintTriangle((i-1), py, (i-1), py, (i), y, canvas.Blue)
		//c.PaintTriangle((i-1), px, (i-1), px, (i), x, canvas.Green)
		//}
		c.SSet(float64(i), imag(e), canvas.Blue)
		c.SSet(float64(i), real(e), canvas.Green)
		//px, py = x, y
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	c.WritePng(f)
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func plotDual(rex, imx []float64, filename string) {
	n := len(rex)
	c := canvas.NewCanvasWithScale(512, 512, 0, float64(n), -5.0, 5.0)
	for i, e := range imx {
		c.SSet(float64(i), e, canvas.Blue)
	}
	for i, e := range rex {
		c.SSet(float64(i), e, canvas.Green)
	}
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	c.WritePng(f)
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	x := make([]float64, 64)
	for i := 0; i < 32; i++ {
		x[i] = 1.0
	}
	plotReal(x, "/tmp/fft1.png")
	y := fft.FFTReal(x)
	fmt.Println(y)
	plotComplex(y, "/tmp/fft2.png")
	z := fft.IFFT(y)
	plotComplex(z, "/tmp/fft3.png")
}
