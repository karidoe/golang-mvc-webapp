package models

import (
	validator "github.com/asaskevich/govalidator"
	"golang-mvc-webapp/config"
	"golang-mvc-webapp/db"
	"gopkg.in/mgo.v2/bson"
)

type ProductModel struct {}

type ProductItem struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Sku string `json:"sku" valid:"type(string)|required"`
	Name string `json:"name" valid:"type(string)|required"`
	Price float64 `json:"price" valid:"required"`
}

var (
	DB *db.Mongodb
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

func (item *ProductItem) IsValid() (bool, error) {
	valid, err := validator.ValidateStruct(item)
	return valid, err
}