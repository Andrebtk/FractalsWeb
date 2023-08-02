package main


import "math/rand"


var f1, f2 func(float64, float64) (float64, float64)

func fractal(u float64, typeFractal string) ([]float64 ,[]float64) {


	//var u int = 10e4

	var x []float64 
	var y []float64

	if typeFractal == "custom1" {
		f1 = func(x, y float64) (float64, float64) { return (x - y)/2 - 0.5, (x + y)/2 }
		f2 = func(x, y float64) (float64, float64) { return (x + y)/2 + 0.5, (x - y)/2 + 0.5 }
	} else if typeFractal == "custom2" {
		f1 = func(x, y float64) (float64, float64) { return (x - y)/2, -(x + y)/2 }
		f2 = func(x, y float64) (float64, float64) { return (x + y)/2 + 0.5, (x - y)/2 + 0.5 }
	} else if typeFractal == "levyCurve" {
		f1 = func(x, y float64) (float64, float64) { return (x - y)/2, (x + y)/2 }
		f2 = func(x, y float64) (float64, float64) { return (x + y)/2 + 0.5, -(x - y)/2 + 0.5 }
	} else if typeFractal == "dragonCurve" {
		f1 = func(x, y float64) (float64, float64) { return (x - y)/2, (x + y)/2 }
		f2 = func(x, y float64) (float64, float64) { return -(x + y)/2 + 0.5, (x - y)/2 + 0.5 }
	}

	x = append(x, 0)
	y = append(y, 0)
	
	for  i := 0; i<int(u); i++  {
		var randNum int = rand.Intn(2) + 1
		
		
		if  randNum==1  {
			x[i], y[i] = f1(x[i], y[i])
		}else{
			x[i], y[i] = f2(x[i], y[i])
		}
		

		x = append(x, x[i])
		y = append(y, y[i])
	}

	return x,y
}
