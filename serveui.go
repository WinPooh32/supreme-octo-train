package main

import (
	"encoding/csv"
	"io"
	"log"
	"mime"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func prepareForecast(data []float64, lacks []YearLacks) (forecast, upperLimit, filtered, restored []float64) {
	smoothHistory := movingavg(data, 2)
	upperLimit = confidenceUpperLimit(smoothHistory, 4)

	filtered = approximateByRegression(upperLimit)
	coefs := calcYearCoefficient(data, filtered)
	multCoefficient(filtered, coefs)

	//Внутри восстанавливает filtered!
	restored = restoreLacks(filtered, lacks)

	forecast = buildForecast(approximateByRegression(filtered[0 : len(filtered)-weeksperyear]))

	return
}

func app(r *gin.Engine) {
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
		if strings.HasPrefix(c.Request.URL.Path, staticPath) {
			c.Status(404)
			return
		}

		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, html, "Прогноз закупок")
	})
}

func api(r *gin.Engine) {
	r.GET("/forecast", func(c *gin.Context) {

		response := make([]gin.H, 0, 10)

		//Считываем товары
		fileIn, _ := os.Open("продажи.csv")
		defer fileIn.Close()
		reader := csv.NewReader(fileIn)

		for i := 0; ; i++ {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			//Парсим провалы
			lacks := make([]YearLacks, 0, 4)

			const yearsField = 1

			years, _ := strconv.Atoi(record[yearsField])
			rawLacks := record[2 : 2+years]

			for _, v := range rawLacks {
				lack := parseLackRange(v)
				lacks = append(lacks, lack)
			}

			// реверс т.к. годы не в том порядке в файле
			// так же для данных продаж
			lacks = reverseLacks(lacks)

			//Считываем название и статистику продаж
			dataBegin := 2 + years
			name := record[0]
			row := reverse(toFloat64(record[dataBegin:]))

			forecast, upperLimit, filtered, restored := prepareForecast(row, lacks)

			response = append(response, gin.H{
				"id":         i,
				"name":       name,
				"data":       row,
				"forecast":   forecast,
				"upperLimit": upperLimit,
				"filtered":   filtered,
				"restored": restored,
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

		c.Next()
	}
}

func serveui(dir string) {
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
