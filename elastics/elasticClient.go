package elastics

import (
    "encoding/json"
    "fmt"
      //   "log"
    // "time"
    "../models"
    
    "github.com/olivere/elastic"
)



func checkClient() {
    // Create a client and connect to http://192.1.199.81:9200
	// client, err := elastic.NewClient(elastic.SetURL("http://192.1.199.81:9200"))
    client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

	if err != nil {
		// Handle error
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping().Do()
	if err != nil {
		// Handle error
		panic(err)
	}
    
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
    
}

func elasticDel() {
    client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
    
    // Delete the index again
    _, err = client.DeleteIndex("products").Do()
    if err != nil {
        // Handle error
        panic(err)
    }
}

func elasticQuery() {
    client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

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
                var p models.Product
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
    
}


func elasticInsert() {
    // Create a client and connect to http://127.0.0.1:9200
    client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))

	if err != nil {
		// Handle error
		panic(err)
	}

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
    
    var product =   "name"
    
    // Index a product (using JSON serialization)
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
}