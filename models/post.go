package models

import (
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type Post struct {
        Id          bson.ObjectId `json:"id" bson:"_id"`
        UserId      int64       `json:"user-id"`
        Type        string      `json:"type"`
        Active      bool        `json:"active"`
        Content     json      `json:"content"`
        
        CreatedAt   time.Time   `json:"created-at"` 
        UpdatedAt   time.Time   `json:"updated-at"`
        /*
        file content:   { "title": "xx",    "link": "xx",   "name": "xx",   "comment": "xx"}
        text content:   { "text-message": "xx"}
        video content:  { "title": "xx",    "link": "xx",   "name": "xx",   "comment": "xx"}
        image content:  { "title": "xx",    "link": "xx",   "name": "xx",   "comment": "xx"}
        */
    }


type TextContent struct {
	TextMessage string `json:"text-message"`
}
type MediaContent struct {
	Title string `json:"title"`
	Link string `json:"link"`
	Name string `json:"name"`
	Comment string `json:"comment"`
}

type Posts [] Post
