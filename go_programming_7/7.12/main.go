package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var db database

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var dbTemplate = template.Must(template.New("dblist").Parse(`
<h1>dblist</h1>
<table>
<tr style='text-align: left'>
  <th>name</th>
  <th>price</th>
</tr>
{{range $key,$value:=.}}
<tr>
  <td>{{$key}}</td>
  <td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

var mutex sync.RWMutex

//exercise7.12 并发执行下的增删改查
func main() {
	db = database{"shoes": 50, "socks": 5, "liaokun": 10000}
	http.HandleFunc("/list", list)
	http.HandleFunc("/update", update)
	http.HandleFunc("/search", search)
	http.HandleFunc("/delete", remove)
	http.HandleFunc("/create", create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func create(writer http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price := request.URL.Query().Get("price")
	if _, ok := db[item]; ok {
		update(writer, request)
	} else {
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(writer, "The parameter is illegality ")
		}
		mutex.Lock()
		db[item] = dollars(p)
		if err = dbTemplate.Execute(writer, db); err != nil {
			log.Fatal(err)
		}
		mutex.Unlock()
	}
}

func remove(writer http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		mutex.Lock()
		delete(db, item)
		if err := dbTemplate.Execute(writer, db); err != nil {
			log.Fatal(err)
		}
		mutex.Unlock()
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "no such item:%s", item)
	}
}

func search(writer http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	if _, ok := db[item]; ok {
		if err := dbTemplate.Execute(writer, database{item: db[item]}); err != nil {
			log.Fatal(err)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "no such item:%s", item)
	}
}

func update(writer http.ResponseWriter, request *http.Request) {
	item := request.URL.Query().Get("item")
	price := request.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		newPrice, err := strconv.ParseFloat(price, 32)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Price: %q is invalid.\n", price)
		} else {
			mutex.Lock()
			db[item] = dollars(newPrice)
			if err := dbTemplate.Execute(writer, db); err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
		}

	} else {
		writer.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(writer, "no such item: %q\n", item)
	}
}

func list(writer http.ResponseWriter, request *http.Request) {
	mutex.RLock()
	if err := dbTemplate.Execute(writer, db); err != nil {
		log.Fatal(err)
	}
	mutex.RUnlock()
}
