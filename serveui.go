package main

import (
	"path"
	"mime"
	"os"
	"encoding/csv"
	"io"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

func prepareForecast(data []float64) (forecast, upperLimit, filtered []float64){
	smoothHistory := movingavg(data, 2)
	upperLimit = confidenceUpperLimit(smoothHistory, 4)
	
	filtered = approximateByRegression(upperLimit)
	coefs := calcYearCoefficient(data, filtered)
	multCoefficient(filtered, coefs)

	forecast = buildForecast(approximateByRegression(upperLimit[0:len(filtered)-weeksperyear]))

	return
}

func app(r *gin.Engine){
	html := `<!DOCTYPE html>
<html>

<head>
	<meta charset="UTF-8" />	
	<title>%s</title>
</head>

<body>
	<div id="root"></div>

	<!-- Dependencies -->
	<script src="/static/react.development.js"></script>
	<script src="/static/react-dom.development.js"></script>

	<!-- Main -->
	<script src="/static/bundle.js"></script>
</body>

</html>`

	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, staticPath){
			c.Status(404)
			return
		}

		log.Println("NOROUTE")

		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, html, "Прогноз закупок")
	})
}

func api(r *gin.Engine){
	r.GET("/forecast", func(c *gin.Context) {

		response := make([]gin.H, 0, 10)

		//Считываем товары
		fileIn, _ := os.Open("продажи.csv")
		defer fileIn.Close()
		reader := csv.NewReader(fileIn)
	
		for i := 0; ;i++ {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			name := record[0]
			row := reverse(toFloat64(record[1:]))

			forecast, upperLimit, filtered := prepareForecast(row)
			
			response = append(response, gin.H{
				"id": i,
				"name": name,
				"data": row,
				"forecast": forecast,
				"upperLimit": upperLimit,
				"filtered": filtered,
			})
		}

		c.JSON(200, response)
	})
}

func detectMIME() gin.HandlerFunc {
	return func(c *gin.Context) {

		var contentType string 
		ext := path.Ext(c.Request.URL.EscapedPath())
		
		switch ext { 
		case ".js":
			contentType = "application/javascript"
		case ".woff":
			contentType = "application/font-woff"
		case ".woff2":
			contentType = "application/font-woff2"
		default:
			contentType = mime.TypeByExtension(ext)
		}
		
		if len(contentType) > 0 {
			c.Header("Content-Type", contentType)
		}

		log.Println("contentType", contentType, c.Writer.Header().Get("Content-Type"))

		c.Next()
	}
}

func serveui(dir string){
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	r.Use(cors.New(config))
	r.Use(detectMIME())

	app(r)
	api(r)	

	r.StaticFS(staticPath, gin.Dir(dir, true))
	// r.Static(staticPath, dir)
	r.Run() // listen and serve on 0.0.0.0:8080
}