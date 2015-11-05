package models

import (
    "time"
    "gopkg.in/mgo.v2/bson"
)

type Product struct {
    Id                          bson.ObjectId `json:"id" bson:"_id"`
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


type Products [] Product
