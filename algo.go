package main


import "math/rand"



func fractal(u float64, typeFractal string) ([]float64 ,[]float64) {


	//var u int = 10e4

	var x []float64 
	var y []float64



	x = append(x, 0)
	y = append(y, 0)
	
	for  i := 0; i<int(u); i++  {
		var randNum int = rand.Intn(2) + 1
		
		
		if typeFractal == "custom1"{	

			if  randNum==1  {
				x[i], y[i] = customF1(x[i], y[i])
			}else{
				x[i], y[i] = customF2(x[i], y[i])
			}

		}else if typeFractal == "custom2"{

			if  randNum==1  {
				x[i], y[i] = custom2F1(x[i], y[i])
			}else{
				x[i], y[i] = custom2F2(x[i], y[i])
			}

		}else if typeFractal == "levyCurve"{

			if  randNum==1  {
				x[i], y[i] = levyCurveF1(x[i], y[i])
			}else{
				x[i], y[i] = levyCurveF2(x[i], y[i])
			}

		}else if typeFractal == "dragonCurve"{

			if  randNum==1  {
				x[i], y[i] = dragonCurveF1(x[i], y[i])
			}else{
				x[i], y[i] = dragonCurveF2(x[i], y[i])
			}

		}
		

		x = append(x, x[i])
		y = append(y, y[i])
	}

	return x,y
}





func customF1(x float64, y float64) (float64, float64) {
	return (x-y)/2 - 0.5, (x+y)/2 
}

func customF2(x, y float64) (float64, float64) {
	return (x+y)/2 + 0.5, (x-y)/2 + 0.5 

}



func custom2F1(x float64, y float64) (float64, float64) {
	return (x-y)/2, -(x+y)/2 
}

func custom2F2(x, y float64) (float64, float64) {
	return (x+y)/2 + 0.5, (x-y)/2 + 0.5 

}



func levyCurveF1(x float64, y float64) (float64, float64) {
	return (x-y)/2, (x+y)/2 
}

func levyCurveF2(x, y float64) (float64, float64) {
	return (x+y)/2 + 0.5, -(x-y)/2 + 0.5 

}



func dragonCurveF1(x float64, y float64) (float64, float64) {
	return (x-y)/2, (x+y)/2 
}

func dragonCurveF2(x, y float64) (float64, float64) {
	return -(x+y)/2 + 0.5, (x-y)/2 + 0.5 

}

