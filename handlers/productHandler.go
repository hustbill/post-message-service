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

// Check Elasticsearch is available or not
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
    return
    
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

	// Stub post
	u := models.Product{}

	// Fetch post
	if err := uc.session.DB("post_message_service").C("products").FindId(oid).One(&u); err != nil {
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