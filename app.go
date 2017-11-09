package main

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/graphql", middleware(CreateSchema()))
	log.Println("Running..")
	http.ListenAndServe(":8080", nil)
}

func middleware(schema graphql.Schema) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(data),
		})

		json.NewEncoder(writer).Encode(result)
	}
}
