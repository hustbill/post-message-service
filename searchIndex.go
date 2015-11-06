package main

import (
	"flag"
	"fmt"
	elastigo "github.com/mattbaird/elastigo/lib"
	"log"
	"os"
    
    //"./models"
)

var (
	host *string = flag.String("host", "localhost", "Elasticsearch Host")
)

func main() {
	c := elastigo.NewConn()
	log.SetFlags(log.LstdFlags)
	flag.Parse()

	// Trace all requests
	c.RequestTracer = func(method, url, body string) {
		log.Printf("Requesting %s %s", method, url)
		log.Printf("Request body: %s", body)
	}

	fmt.Println("host = ", *host)
	// Set the Elasticsearch Host to Connect to
	c.Domain = *host

    // Index a document
    _, err := c.Index("oldindex", "product", "docid_1", nil, `{"name":"bob"}`)
    exitIfErr(err)

    // Index a doc using a map of values
    _, err = c.Index("oldindex", "product", "docid_2", nil, map[string]string{"name": "venkatesh"})
    exitIfErr(err)
    
	// Index a doc using Structs
	_, err = c.Index("testindex", "user", "docid_3", nil,map[string]string{"name": "wena"})
	exitIfErr(err)
        
    // "name":"peanuts","description":"Honey Roasted peanuts","permalink":"",
    //             "tax_category_id":0,"shipping_category_id":0,"deleted_at":"0001-01-01T00:00:00Z","meta_description":"",
    //             "meta_keywords":"","position":0,"is_featured":false,"can_discount":false,"distributor_only_membership":false
    
    // product := models.Product{ 59001, "peanuts", "Honey Roasted peanuts", "Good taste food"}
    //
    //
    // // Index a doc using Structs
    // _, err = c.Index("pdtindex", "product", "producid_3", nil, product)
    //     fmt.Println(product)

    //exitIfErr(err)

	// Search Using Raw json String
	searchJson := `{
	    "query" : {
	        "term" : { "Name" : "venkatesh" }
	    }
	}`
	out, err := c.Search("oldindex", "product", nil, searchJson)
    
	if len(out.Hits.Hits) == 1 {
		fmt.Println("%v", out.Hits.Hits[0].Source)
	}
	exitIfErr(err)

}
func exitIfErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}

