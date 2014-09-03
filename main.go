package main

import (
	"bufio"
	"net/http"
	"strconv"
	"os"
	"runtime"
	"encoding/json"
	"io/ioutil"
	"html/template"
	"bytes"
	"time"
	"fmt"
	//"github.com/hoisie/mustache"
	//"github.com/Wuvist/mustache"
	"github.com/fromYukki/mustache"
	pongo2 "gopkg.in/flosch/pongo2.v1"
)

func generateGo() {
	f, err := os.Create("./templates/inheritance.go/b0.go.tpl")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("<h1>500 blocks</h1>\n")

	for i := 1; i < 501; i++ {
		is := strconv.Itoa(i)
		ism := strconv.Itoa(i - 1)

		ff, errr := os.Create("./templates/inheritance.go/b" + is + ".go.tpl")
		if errr != nil {
			panic(err)
		}

		ww := bufio.NewWriter(ff)
		ww.WriteString("{{template \"b" + ism + ".go.tpl\"}}\n")
		ww.Flush()

		ff.Close()
	}

	w.Flush()
}

func generateMustache() {
	f, err := os.Create("./templates/inheritance.mustache/b0.mustache")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("<h1>500 blocks</h1>\n")

	for i := 1; i < 501; i++ {
		is := strconv.Itoa(i)
		ism := strconv.Itoa(i - 1)

		ff, errr := os.Create("./templates/inheritance.mustache/b" + is + ".mustache")
		if errr != nil {
			panic(err)
		}

		ww := bufio.NewWriter(ff)
		ww.WriteString("{{> b" + ism + "}}\n")
		ww.Flush()

		ff.Close()
	}

	w.Flush()
}

func generatePongo() {
	f, err := os.Create("./templates/inheritance.pongo/b0.pongo.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("<h1>500 blocks</h1>\n")

	for i := 1; i < 501; i++ {
		is := strconv.Itoa(i)
		ism := strconv.Itoa(i - 1)

		w.WriteString("{% block b" + is +" %}{% endblock %}\n")
		ff, errr := os.Create("./templates/inheritance.pongo/b" + is + ".pongo.html")
		if errr != nil {
			panic(err)
		}

		ww := bufio.NewWriter(ff)
		ww.WriteString("{% extends \"b" + ism + ".pongo.html\" %}\n{% block b" + is +" %}data" + is +"{% endblock %}\n")
		ww.Flush()

		ff.Close()
	}

	w.Flush()
}

type TestData struct {
	Title		string			`json:"title"`
	Keywords	string			`json:"keywords"`
	Var1		string			`json:"var1"`
	Var2		string			`json:"var2"`
	Var3		string			`json:"var3"`
	Var4		string			`json:"var4"`
	Var5		string			`json:"var5"`
	Var6		string			`json:"var6"`
	Var7		string			`json:"var7"`
	Var8		string			`json:"var8"`
	Var9		string			`json:"var9"`
	Var0		string			`json:"var0"`
	Items		[]TestDataItem	`json:"array"`
}

type TestDataItem struct {
	Id		int		`json:"id"`
	Title	string	`json:"title"`
	Var1	string	`json:"var1"`
	Var2	string	`json:"var2"`
	Var3	string	`json:"var3"`
	Var4	string	`json:"var4"`
	Var5	string	`json:"var5"`
	Var6	string	`json:"var6"`
}

var (
	globData		TestData
	goTpl1			*template.Template
	goTpl2			*template.Template
	goTpl3			*template.Template
	mustacheTpl1	*mustache.Template
	mustacheTpl2	*mustache.Template
	mustacheTpl3	*mustache.Template
	pongoTpl1		*pongo2.Template
	pongoTpl2		*pongo2.Template
	pongoTpl3		*pongo2.Template
	err 			error
)

func printDuration(startTime time.Time) {
	duration := time.Since(startTime).Nanoseconds()
	var durationUnits string
	switch {
	case duration > 2000000:
		durationUnits = "ms"
		duration /= 1000000
	case duration > 1000:
		durationUnits = "μs"
		duration /= 1000
	default:
		durationUnits = "ns"
	}
	fmt.Printf("[%d %s]\n", duration, durationUnits)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())



	// Generate some templates
	generateGo()
	generateMustache()
	generatePongo()

	// Read basic data
	var data []byte
	if data, err = ioutil.ReadFile("data.json"); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(data, &globData); err != nil {
		panic(err)
	}

	// Scan templates
	goTpl1, err = template.ParseFiles("./templates/echo.go.tpl")
	if err != nil {
		panic(err)
	}
	goTpl2, err = template.ParseFiles("./templates/for.go.tpl")
	if err != nil {
		panic(err)
	}
	goTpl3, err = template.ParseGlob("./templates/inheritance.go/*.go.tpl")
	if err != nil {
		panic(err)
	}

	mustacheTpl1, err = mustache.ParseFile("./templates/echo.mustache")
	if err != nil {
		panic(err)
	}
	mustacheTpl2, err = mustache.ParseFile("./templates/for.mustache")
	if err != nil {
		panic(err)
	}
	mustacheTpl3, err = mustache.ParseFile("./templates/inheritance.mustache/b500.mustache")
	if err != nil {
		panic(err)
	}
	pongoTpl1, err = pongo2.FromFile("./templates/echo.pongo.html")
	if err != nil {
		panic(err)
	}
	pongoTpl2, err = pongo2.FromFile("./templates/for.pongo.html")
	if err != nil {
		panic(err)
	}
	pongoTpl3, err = pongo2.FromFile("./templates/inheritance.pongo/b100.pongo.html")
	if err != nil {
		panic(err)
	}


	// Handle request
	http.HandleFunc("/go1", func(w http.ResponseWriter, r *http.Request) {
			/**
				Echo 10000 variables

				Transactions:		       45753 hits
				Availability:		      100.00 %
				Elapsed time:		      299.11 secs
				Data transferred:	     2226.40 MB
				Response time:		        0.65 secs
				Transaction rate:	      152.96 trans/sec
				Throughput:		        7.44 MB/sec
				Concurrency:		       99.86
				Successful transactions:       45753
				Failed transactions:	           0
				Longest transaction:	        1.70
				Shortest transaction:	        0.04

				Go: [186 ms] - [844 ms] / 2 = [515 ms]
			 */
			startTime := time.Now()

			buf := bytes.NewBuffer([]byte{})
			goTpl1.Execute(buf, globData)

			w.WriteHeader(200)
			w.Write(buf.Bytes())

			printDuration(startTime)
		})
	http.HandleFunc("/go2", func(w http.ResponseWriter, r *http.Request) {
			/**
				Echo 10 items from 1000 elements in loop

				Transactions:		       38654 hits
				Availability:		      100.00 %
				Elapsed time:		      257.71 secs
				Data transferred:	     1992.46 MB
				Response time:		        0.67 secs
				Transaction rate:	      149.99 trans/sec
				Throughput:		        7.73 MB/sec
				Concurrency:		       99.82
				Successful transactions:       38654
				Failed transactions:	           0
				Longest transaction:	        1.86
				Shortest transaction:	        0.05

				Go: [167 ms] - [991 ms] / 2 = [579 ms]
			 */
			startTime := time.Now()

			buf := bytes.NewBuffer([]byte{})
			goTpl2.Execute(buf, globData)

			w.WriteHeader(200)
			w.Write(buf.Bytes())

			printDuration(startTime)
		})
	http.HandleFunc("/go3", func(w http.ResponseWriter, r *http.Request) {
			/**
				500 blocks inheritance

				Transactions:		      180590 hits
				Availability:		       99.39 %
				Elapsed time:		      395.80 secs
				Data transferred:	       89.56 MB
				Response time:		        0.05 secs
				Transaction rate:	      456.27 trans/sec
				Throughput:		        0.23 MB/sec
				Concurrency:		       22.42
				Successful transactions:      180590
				Failed transactions:	        1100
				Longest transaction:	        3.59
				Shortest transaction:	        0.00

				Go: [0.9 ms] - [8 ms] / 2 = [4.45 ms]
			 */
			startTime := time.Now()

			buf := bytes.NewBuffer([]byte{})
			goTpl3.ExecuteTemplate(buf, "b500.go.tpl", globData)

			w.WriteHeader(200)
			w.Write(buf.Bytes())

			printDuration(startTime)
		})
	http.HandleFunc("/mustache1", func(w http.ResponseWriter, r *http.Request) {
			/**
				Echo 10000 variables

				Transactions:		      103366 hits
				Availability:		      100.00 %
				Elapsed time:		      257.42 secs
				Data transferred:	     5030.41 MB
				Response time:		        0.24 secs
				Transaction rate:	      401.55 trans/sec
				Throughput:		       19.54 MB/sec
				Concurrency:		       98.12
				Successful transactions:      103366
				Failed transactions:	           0
				Longest transaction:	       14.95
				Shortest transaction:	        0.00

				Go: [8 ms] - [305 ms] / 2 = [156 ms]
			 */
			startTime := time.Now()

			w.WriteHeader(200)
			w.Write([]byte(mustacheTpl1.Render(globData)))

			printDuration(startTime)
		})
	http.HandleFunc("/mustache2", func(w http.ResponseWriter, r *http.Request) {
			/**
				Echo 10 items from 1000 elements in loop

				Transactions:		      256765 hits
				Availability:		      100.00 %
				Elapsed time:		      585.09 secs
				Data transferred:	    12745.49 MB
				Response time:		        0.22 secs
				Transaction rate:	      438.85 trans/sec
				Throughput:		       21.78 MB/sec
				Concurrency:		       98.04
				Successful transactions:      256765
				Failed transactions:	           0
				Longest transaction:	        3.02
				Shortest transaction:	        0.01

				Go: [38 ms] - [180 ms] / 2 = [109 ms]
			 */
			startTime := time.Now()

			w.WriteHeader(200)
			w.Write([]byte(mustacheTpl2.Render(globData)))

			printDuration(startTime)
		})
	http.HandleFunc("/mustache3", func(w http.ResponseWriter, r *http.Request) {
			/**
				500 blocks inheritance

				Transactions:		      147924 hits
				Availability:		       99.57 %
				Elapsed time:		      297.84 secs
				Data transferred:	       73.36 MB
				Response time:		        0.08 secs
				Transaction rate:	      496.66 trans/sec
				Throughput:		        0.25 MB/sec
				Concurrency:		       38.60
				Successful transactions:      147924
				Failed transactions:	         641
				Longest transaction:	       21.35
				Shortest transaction:	        0.00

				Go: [0.4 ms] - [3 ms] / 2 = [1.7 ms]
			 */
			startTime := time.Now()

			w.WriteHeader(200)
			w.Write([]byte(mustacheTpl3.Render(globData)))

			printDuration(startTime)
		})
	http.HandleFunc("/pongo1", func(w http.ResponseWriter, r *http.Request) {
			/**
				Echo 10000 variables

				Transactions:		       72414 hits
				Availability:		      100.00 %
				Elapsed time:		      170.71 secs
				Data transferred:	     3524.10 MB
				Response time:		        0.23 secs
				Transaction rate:	      424.19 trans/sec
				Throughput:		       20.64 MB/sec
				Concurrency:		       99.23
				Successful transactions:       72414
				Failed transactions:	           0
				Longest transaction:	        8.13
				Shortest transaction:	        0.00

			 */
			startTime := time.Now()

			ctx := pongo2.Context{
			"Var0": globData.Var0,
			"Var1": globData.Var1,
			"Var2": globData.Var2,
			"Var3": globData.Var3,
			"Var4": globData.Var4,
			"Var5": globData.Var5,
			"Var6": globData.Var6,
			"Var7": globData.Var7,
			"Var8": globData.Var8,
			"Var9": globData.Var9,
			}
			res, _ := pongoTpl1.ExecuteBytes(ctx)

			w.WriteHeader(200)
			w.Write(res)

			printDuration(startTime)
		})
	http.HandleFunc("/pongo2", func(w http.ResponseWriter, r *http.Request) {
			/**
				Echo 10 items from 1000 elements in loop

				Transactions:		      101368 hits
				Availability:		      100.00 %
				Elapsed time:		      314.93 secs
				Data transferred:	     5611.81 MB
				Response time:		        0.31 secs
				Transaction rate:	      321.87 trans/sec
				Throughput:		       17.82 MB/sec
				Concurrency:		       98.84
				Successful transactions:      101368
				Failed transactions:	           0
				Longest transaction:	        1.24
				Shortest transaction:	        0.02
			 */
			startTime := time.Now()

			ctx := pongo2.Context{
			"Items": globData.Items,
			}
			res, _ := pongoTpl2.ExecuteBytes(ctx)

			w.WriteHeader(200)
			w.Write(res)

			printDuration(startTime)
		})
	http.HandleFunc("/pongo3", func(w http.ResponseWriter, r *http.Request) {
			/**
				500 blocks inheritance

				Transactions:		       49333 hits
				Availability:		       99.60 %
				Elapsed time:		       89.38 secs
				Data transferred:	       52.32 MB
				Response time:		        0.05 secs
				Transaction rate:	      551.95 trans/sec
				Throughput:		        0.59 MB/sec
				Concurrency:		       26.74
				Successful transactions:       49333
				Failed transactions:	         200
				Longest transaction:	        2.90
				Shortest transaction:	        0.00
			 */
			startTime := time.Now()

			res, _ := pongoTpl3.ExecuteBytes(pongo2.Context{})

			w.WriteHeader(200)
			w.Write(res)

			printDuration(startTime)
		})
	http.ListenAndServe(":8080", nil)
}
