package models

import (
	"golang-mvc-webapp/config"
	"golang-mvc-webapp/db"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct{}

type ProductItem struct {
	Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Sku   string        `json:"sku" validate:"required,numeric"`
	Name  string        `json:"name" validate:"required"`
	Price float64       `json:"price" validate:"required,numeric"`
}

var (
	DB     *db.Mongodb
	dbName = config.Getenv("APP_MONGO_DATABASE")
)

func GetProductModel() *ProductModel {
	return &ProductModel{}
}

func (c *ProductModel) Create(p ProductItem) error {
	session := db.GetMongodb().GetSession()
	defer session.Close()

	err := session.DB(dbName).C("products").Insert(p)
	return err
}

func (c *ProductModel) All() ([]ProductItem, error) {

	session := db.GetMongodb().GetSession()
	defer session.Close()

	var results []ProductItem
	err := session.DB(dbName).C("products").Find(bson.M{}).All(&results)
	return results, err
}
