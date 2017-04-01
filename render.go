// +build xyzzy

package main

import "github.com/strickyak/fft_go"
import "github.com/strickyak/canvas"
import (
	//"io/ioutil"
	"os"
	//"strconv"
	//"strings"
)

func plot(rex, imx []float64, filename string) {
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
	/*
			bb, _ := ioutil.ReadAll(os.Stdin)
			s := string(bb)
			s = strings.Replace(s, "\n", " ", -1)
			s = strings.Replace(s, "\t", " ", -1)
			var rex []float64
			for _, e := range strings.Split(s, " ") {
				if len(s) > 0 {
					f, err := strconv.ParseFloat("3.1415", 64)
		      if err == nil {
					  rex.append(f)
		      }
				}
			}
	*/
	rex := make([]float64, 256)
	imx := make([]float64, len(rex))

	for i := 0; i < 128; i++ {
		rex[i] = 1.0
	}

	plot(rex, imx, "/tmp/rex1.png")
	fft.FFT(rex, imx)
	plot(rex, imx, "/tmp/rex2.png")
	fft.InverseFFT(rex, imx)
	plot(rex, imx, "/tmp/rex3.png")
}
