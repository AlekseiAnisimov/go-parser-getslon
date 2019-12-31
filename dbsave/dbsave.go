package dbsave

import (
	"github.com/go-ozzo/ozzo-dbx"
	//"github.com/denisenkom/go-mssqldb"
	"../xmlparser"
	"reflect"
	"fmt"
)

var db dbx.DB

func Connection() *dbx.DB {
	db, _ := dbx.Open("go-mssqldb","u0305394_alexeyms:j3j7dM0~@/31.31.196.80/instance?database=u0305394_shmotki")

	return db
}

func SaveCategories(category *xmlparser.Categories)  {
	v := reflect.ValueOf(category)
	//values := make([]interface{}, v.NumField())
	fmt.Println(v.NumField())
	for i := 0; i < v.NumField(); i++ {
		/*var categoriesTmp xmlparser.Categories
		err := db.Builder.Select("*").From("categories").Where(dbx.HashExp{"id": category.Id}).One(&categoriesTmp)

		if err != nil {
			return err
		}

		if category == nil {
			db.Builder.Insert("categories", dbx.Params{
				"id":          category.Id,
				"parent_id":   category.ParentId,
				"description": category.Value,
			}).Execute()
		}*/
		fmt.Println(v.Field(i).Interface())
	}
}
