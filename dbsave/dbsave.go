package dbsave

import (
	"github.com/go-ozzo/ozzo-dbx"
	//"github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"parser/xmlparser"
	"reflect"
)

var db dbx.DB

func Connection() *dbx.DB {
	db, _ := dbx.Open("mysql", "root2:123@/shmotki")
	return db
}

func SaveCategories(categories *xmlparser.Categories) error {
	v := reflect.ValueOf(categories.Category)
	for i := 0; i < v.Len(); i++ {
		id := v.Index(i).FieldByName("Id").Int()
		parentId := v.Index(i).FieldByName("ParentId").Int()
		value := v.Index(i).FieldByName("Value").String()

		var categoryTmp = new(xmlparser.Category)
		db := Connection()
		_ = db.Builder.Select("id", "parent_id", "description").From("categories").Where(dbx.HashExp{"id": id}).One(&categoryTmp)

		if categoryTmp.Id == 0 {
			db.Builder.Insert("categories", dbx.Params{
				"id":          id,
				"parent_id":   parentId,
				"description": value,
			}).Execute()
		}
	}

	return nil
}
