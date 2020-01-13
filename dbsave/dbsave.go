package dbsave

import (
	"github.com/go-ozzo/ozzo-dbx"
	//"github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"../xmlparser"
	"reflect"
	"strings"
	"fmt"
)

var db dbx.DB

func Connection() *dbx.DB {
	db, _ := dbx.Open("mysql", "root:123@/shmotki")
	return db
}

func SaveCategories(categories *xmlparser.Categories) error {
	v := reflect.ValueOf(categories.Category)
	db := Connection()
	for i := 0; i < v.Len(); i++ {
		id := v.Index(i).FieldByName("Id").Int()
		parentId := v.Index(i).FieldByName("ParentId").Int()
		value := v.Index(i).FieldByName("Value").String()

		var categoryTmp = new(xmlparser.Category)

		_ = db.Builder.Select("id", "parent_id", "description").From("categories").Where(dbx.HashExp{"id": id}).One(&categoryTmp)

		if categoryTmp.Id == 0 {
			_, err := db.Builder.Insert("categories", dbx.Params{
				"id":          id,
				"parent_id":   parentId,
				"description": value,
			}).Execute()

			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

func SaveProducts(products *xmlparser.Products) {
	v := reflect.ValueOf(products.Product)
	db := Connection()
	var sb strings.Builder
	rowCnt := 0


	for i := 0; i < v.Len(); i++ {
		batchRow := ""
		rowCnt++


		productId := v.Index(i).FieldByName("ParentId").String()
		article := v.Index(i).FieldByName("Article").String()
		name := v.Index(i).FieldByName("Name").String()
		description := v.Index(i).FieldByName("Description").String()
		available := v.Index(i).FieldByName("Available").String()
		merchantId := v.Index(i).FieldByName("MerchantId").Int()
		gsProductKey := v.Index(i).FieldByName("GsProductKey").String()
		gsCategoryId := v.Index(i).FieldByName("GsCategoryId").Int()
		picture := v.Index(i).FieldByName("Picture").String()
		thumbnail := v.Index(i).FieldByName("Thumbnail").String()
		originalPicture := v.Index(i).FieldByName("OriginalPicture").String()
		vendor := v.Index(i).FieldByName("Vendor").String()
		model := v.Index(i).FieldByName("Model").String()
		oldprice := v.Index(i).FieldByName("Oldprice").Float()
		url := v.Index(i).FieldByName("Url").String()
		destinationUrl := v.Index(i).FieldByName("DestinationUrl").String()
		currencyId	:= v.Index(i).FieldByName("CurrencyId").String()
		price := v.Index(i).FieldByName("Price").Float()
		age := v.Index(i).FieldByName("Age").String()
		composition := v.Index(i).FieldByName("Composition").String()
		otherPictures := v.Index(i).FieldByName("OtherPictures").String()

		batchRow := fmt.Sprintf("(%s, %s, %s, %s, %s, %d, %s, %d, %s, %s, %s, %s, %s, %v, %s, %s, %s, %v, %s, %s, %s)", )

		if(rowCnt > 100) {
			_ db.NewQuery("INSERT INTO ")
		}
	}
}
