
package main

import (
    "encoding/json"
    "fmt"
    // "log"
    // "os"
    // "reflect"
	"time"
     "gopkg.in/olivere/elastic.v2"
    

)

type Product struct {
    Id                          int64       `json:"id"`
    Name                        string      `json:"name"`
    Description                 string      `json:"description"`
    Permalink                   string      `json:"permalink"`
    TaxCategoryId               int64       `json:"tax_category_id"`
    ShippingCategoryId          int64       `json:"shipping_category_id"`
    DeletedAt                   time.Time   `json:"deleted_at"` 
    MetaDescription             string      `json:"meta_description"`
    MetaKeywords                string      `json:"meta_keywords"`
    Position                    int64       `json:"position"`
    IsFeatured                  bool        `json:"is_featured"`
    CanDiscount                 bool        `json:"can_discount"`
    DistributorOnlyMembership   bool        `json:"distributor_only_membership"`
}


func main() {
    
    //
    // errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)
    //
    //     // Obtain a client. You can provide your own HTTP client here.
    //     // client, err := elastic.NewClient(elastic.SetErrorLog(errorlog))
    //
    // // Create a client and connect to http://192.168.2.10:9201
    // elastic.SetErrorLog(errorlog)
    
    // client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
  //   if err != nil {
  //       // Handle error
  //       panic(err)
  //   }
  //
  
u := elastic.SetURL("http://192.168.1.59:9200")

client, err := elastic.NewClient(u, elastic.SetSniff(false))
  	if err != nil {
  		// Handle error
         panic(err)
  	}

    	// Trace request and response details like this
    	//client.SetTracer(log.New(os.Stdout, "", 0))

    	// Ping the Elasticsearch server to get e.g. the version number
        // info, code, err := client.Ping().Do()
      //         if err != nil {
      //             // Handle error
      //             panic(err)
      //         }
    	//fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)

    	// Getting the ES version number is quite common, so there's a shortcut
    	esversion, err := client.ElasticsearchVersion("http://192.168.1.59:9200")
    	if err != nil {
    		// Handle error
    		panic(err)
    	}
    	fmt.Printf("Elasticsearch version %s", esversion)


	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("products").Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("products").Do()
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
    
    /*curl -XPOST -H 'Content-Type: application/json' -d '{"Name": "peanuts", "Description": "Honey Roasted peanuts", "MetaDescription": "Good taste food"}' http://127.0.0.1:3000/v1/products 
   Result : {"id":"563ab9e7d126195066000001","name":"peanuts","description":"Honey Roasted peanuts","permalink":"",
               "tax_category_id":0,"shipping_category_id":0,"deleted_at":"0001-01-01T00:00:00Z","meta_description":"",
               "meta_keywords":"","position":0,"is_featured":false,"can_discount":false,"distributor_only_membership":false} */
    
    // Index a product (using JSON serialization)
    product := Product{Name: "peanuts", Description: "Honey Roasted peanuts", MetaDescription: "Good taste food"}
    put1, err := client.Index().
        Index("products").
        Type("product").
        Id("1").
        BodyJson(product).
        Do()
    if err != nil {
        // Handle error
        panic(err)
    }
    fmt.Printf("Indexed product %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)


    //Get product with specified ID
    get1, err := client.Get().
        Index("products").
        Type("product").
        Id("1").
        Do()
    if err != nil {
        // Handle error
        panic(err)
    }
    if get1.Found {
        fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)

    }

    // Flush to make sure the documents got written.
    _, err = client.Flush().Index("products").Do()
    if err != nil {
        panic(err)
    }

        // Search with a term query
        termQuery := elastic.NewTermQuery("name", "annual_fee")
        searchResult, err := client.Search().
            Index("products").   // search in index "products"
            Query(&termQuery).  // specify the query
            Sort("name", true). // sort by "name" field, ascending
            From(0).Size(10).   // take documents 0-9
            Pretty(true).       // pretty print request and response JSON
            Do()                // execute
        if err != nil {
            // Handle error
            panic(err)
        }

        // searchResult is of type SearchResult and returns hits, suggestions,
        // and all kinds of other information from Elasticsearch.
        fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)


        // // TotalHits is another convenience function that works even when something goes wrong.
        fmt.Printf("Found a total of %d products\n", searchResult.TotalHits())


        // Here's how you iterate through results with full control over each step.
        if searchResult.Hits != nil {
            fmt.Printf("Found a total of %d products\n", searchResult.Hits.TotalHits)

            // Iterate through results
            for _, hit := range searchResult.Hits.Hits {
                // hit.Index contains the name of the index

                // Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
                var p Product
                err := json.Unmarshal(*hit.Source, &p)
                if err != nil {
                    // Deserialization failed
                }

                // Work with product
                fmt.Printf("The new product is %s: %s\n", p.Name, p.Description)
            }
        } else {
            // No hits
            fmt.Print("Found no products\n")
        }

    // Delete the index again
    _, err = client.DeleteIndex("products").Do()
    if err != nil {
        // Handle error
        panic(err)
    }



}
