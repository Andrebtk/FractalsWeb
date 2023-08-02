package main

import "fmt"
import "net/http"
import "text/template"
import "io/ioutil"
import "log"
import "strconv"




type Points struct {
	X []float64
	Y []float64
}



func main(){

	fileContent , err := ioutil.ReadFile("static/main.html")
	if err != nil { log.Fatal(err) }


	doc := string(fileContent)


	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){

		w.Header().Add("Content-Type", "text/html")

		var nPoints float64 = 0
		var fractalType string = ""
		if  req.Method == "POST"  {
			req.ParseForm()

			nPointsStr := req.PostFormValue("numberOfPoints")
			nPoints, _ = strconv.ParseFloat(nPointsStr, 64)

			fractalType = req.PostFormValue("fractalType")

			fmt.Printf("Number of Points: %s\n", nPoints)
			fmt.Printf("Selected Fractal: %s\n\n", fractalType)

		}


		

		templates, err := template.New("template").Parse(doc)
		if err != nil{ print("ERROR") }

		var x,y []float64 = fractal(nPoints, fractalType)


		var fract  = Points{
			X: x,
			Y: y,
		}



		templates.Execute(w, fract)
	})

	
	
	
	fs := http.FileServer(http.Dir("static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	print("Server on : http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
