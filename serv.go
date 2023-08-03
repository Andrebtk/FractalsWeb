package main

import "fmt"
import "net/http"
import "text/template"
import "io/ioutil"

import "strconv"
import "time"


type Points struct {
	X []float64
	Y []float64
}

type FractalRenderData struct {
	Fract      Points
	RenderType string
}



func main(){

	


	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){

		w.Header().Add("Content-Type", "text/html")

		var nPoints float64 = 0
		var fractalType string = ""
		var renderType string = "animated"

		var fileContent []byte

		if  req.Method == "POST"  {
			req.ParseForm()

			nPointsStr := req.PostFormValue("numberOfPoints")
			nPoints, _ = strconv.ParseFloat(nPointsStr, 64)

			nPoints = nPoints/10

			fractalType = req.PostFormValue("fractalType")

			renderType = req.PostFormValue("renderType")
			
			fmt.Printf("Number of Points: %s\n", nPoints)
			fmt.Printf("Selected Fractal: %s\n", fractalType)
			fmt.Printf("Render Type : %s\n", renderType)


		}

		
		if  renderType == "animated"{
			fileContent , _ = ioutil.ReadFile("static/ani.html")
		}else if renderType == "instant"  {
			fileContent , _ = ioutil.ReadFile("static/main.html")
		}
		


		doc := string(fileContent)
		
		

		templates, err := template.New("template").Parse(doc)
		if err != nil{ print("ERROR") }

		start := time.Now()
		var x,y []float64 = fractal(nPoints, fractalType)
		elapsed := time.Since(start)

		

		var fract  = Points{
			X: x,
			Y: y,
		}

		fractalRenderData := FractalRenderData{
			Fract:      fract,
			RenderType: fractalType,
		}
		

		start2 := time.Now()
		templates.Execute(w, fractalRenderData)
		elapsed2 := time.Since(start2)

		fmt.Printf("Execution time for fractal func -> %s\n", elapsed)
		fmt.Printf("Execution time for render -> %s\n\n", elapsed2)
	})


	http.HandleFunc("/ani", func(w http.ResponseWriter, req *http.Request){

		w.Header().Add("Content-Type", "text/html")

		
		fileContent , _ := ioutil.ReadFile("static/animation.html")

		html := string(fileContent)

		w.Write([]byte(html))


	})

	
	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	print("Server on : http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
