package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
    //"../elastics"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/julienschmidt/httprouter"
    "github.com/olivere/elastic"
    
)


type (
	// ProductController represents the controller for operating on the Product resource
	ProductController struct {
		session *mgo.Session
	}
)

// NewProductController provides a reference to a ProductController with provided mongo session
func NewProductController(s *mgo.Session) *ProductController {
	return &ProductController{s}
}

// Create a new Index in Elasticsearch 
func CreateIndex() {
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)

	// Obtain a client. You can provide your own HTTP client here.
	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog),  elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
    
    // Use the IndexExists service to check if a specified index exists.
    exists, err := client.IndexExists("productindex").Do()
    if err != nil {
        painc(err)
    }
    
    if !exists {
        // Create a new index.
        createIndex, err := client.CreateIndex("productIndex").Do()
        if err != nil {
            panic(err)
        }
        if !createIndex.Acknowledged {
            fmt.Println("Not ackowledged")
        }
    }
    
    // Index a product (using JSON serialization) 
    product1 := models.Product{Name : "peanuts", Description: "good food for afternoon"}
    put1, err := client.Index().
        Index("productindex").
        Type("text").
        Id("1").
        BodyJson(product1).
        Do()
    
        if err != nil {
            panic(err)
        }
        fmt.Printf("Indexed product %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func SearchIndexWithId() {
	errorlog := log.New(os.Stdout, "APP ", log.LstdFlags)

	// Obtain a client. You can provide your own HTTP client here.
	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog),  elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
    
    get1, err := client.Get().
        Index("productindex").
        Type("text").
        Id("1").
        Do()
    
    if err != nil {
        panic(err)
    }
    if get1.Found {
        fmt.Printf("Got document %s in version %d from index %s , type %s\n", get1.Id, get1.Version, get1.Index, 
            get1.Type)
    }
}

// Search with a term query 
func SearchIndexWithTermQuery() {
    termQuery := elasticNewTermQuery("Name", "peanuts")
    searchResult, err := clientSearch().
        Index("productindex").   // search in index "productindex"
        Query(&termQuery).      // specify the query
        Sort("Name", true).     // sort by "name" field, ascending
        From(0).Size(10).       // take documents 0-9
        Pretty(true).           // pretty print request and response JSON
        Do()                    // execute

    if err != nil {
        panic(err)
    }
    
    // searchResult is of type SearchResult and returns hits, suggestions,
    // and all kinds of other information from Elasticsearch.
    fmt.Println("Query took %d milliseconds\n", searchResult.TookInMillis)
    

    if searchResult.Hits != nil {
        fmt.Printf("Found a total of %d products\n", searchResult.Hits.TotalHits)
        
        // Iterat through results 
        for _, hit := range searchResult.Hits.Hits {
            // hit.Index contains the name of the index
            // Deserialize hit.Source into a Product 
            // (could also be just a map[string]interface{})
            
            var p models.Product
            err := json.Unmarshal(*hit.Source, &p)
            if err != nil {
                fmt.Println("Deserializaiton failed")
                panic(err)
            }
            
            // work with product
            fmt.Printf("Product %s , %s\n", p.Name, p.Description)
        }
    } else {
        // Not hits
        fmt.Print("Found no products\n")
    }
}

// GetProduct retrieves an individual post resource
func (uc ProductController) GetProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub product
	u := models.Product{}

	// Fetch product from MongoDB
    /*
	if err := uc.session.DB("post_message_service").C("products").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	} */
    
    if err := searchInElastic(); err != nil {
		w.WriteHeader(404)
		return
    }

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)


	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateProduct creates a new post resource
func (uc ProductController) CreateProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an post to be populated from the body
	u := models.Product{}

	// Populate the post data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.Id = bson.NewObjectId()

	// Write the post to mongo
	uc.session.DB("post_message_service").C("products").Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveProduct removes an existing post resource
func (uc ProductController) RemoveProduct(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove post
	if err := uc.session.DB("post_message_service").C("products").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}