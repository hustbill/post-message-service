package models

import (

    "gopkg.in/mgo.v2/bson"
)

type (
     Post struct {
        Id          bson.ObjectId `json:"id" bson:"_id"`
        Type        string      `json:"type"`
        //Active      bool        `json:"active"`
        Content     string      `json:"content"`
        UserId      int64       `json:"user-id"`
        //CreatedAt   time.Time   `json:"created-at"` 
        //UpdatedAt   time.Time   `json:"updated-at"` 
    }
)

type Posts [] Post
