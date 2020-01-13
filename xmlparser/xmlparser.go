package xmlparser

import (
	"encoding/xml"
	//"io/ioutil"
	"os"
)

var xmlfile string

var xmlblob string

type Xml struct {
}

type Categories struct {
	XMLName  xml.Name   `xml:"categories"`
	Category []Category `xml:"category"`
}

type Category struct {
	Id       int    `xml:"id,attr"`
	ParentId int    `xml:"parent_id,attr"`
	Value    string `xml:",chardata"`
}

type Products struct {
	XMLName xml.Name `xml:"offers"`
	Product []Product `xml:"offer"`
}

type Product struct {
	ProductId string `xml:"id,attr"`
	Article string `xml:"article,attr"`
	Name string `xml:"name"`
	Description string `xml:"description"`
	Available bool `xml:"available,attr"`
	MerchantId int `xml:"merchant_id,attr"`
	GsProductKey string `xml:"gs_product_key,attr"`
	GsCategoryId int `xml:"gs_category_id,attr"`
	Picture string `xml:"picture"`
	Thumbnail string `xml:"thumbnail"`
	OriginalPicture string `xml:"original_picture"`
	Vendor string `xml:"vendor"`
	Model string `xml:"model"`
	Oldprice float32 `xml:"oldprice"`
	Url string `xml:"url"`
	DestinationUrl string `xml:"destination-url-do-not-send-traffic"`
	CurrencyId string `xml:"currencyId"`
	Price float32 `xml:"price"`
	Age string `xml:"возраст,attr"`
	Composition string `xml:"состав,attr"`
	//other_pictures JSON
}

func SetXmlFile(path string) {
	xmlfile = path
}

func ReadXmlData(c *Categories, products *Products) error {
	openfile, err := os.Open(xmlfile)
	if err != nil {
		return err
	}

	decoder := xml.NewDecoder(openfile)
	var inElement string

	for {
		t, _ := decoder.Token()

		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			inElement = se.Name.Local

			if inElement == "categories" {
				decoder.DecodeElement(&c, &se)
			} else if inElement == "offers" {
				decoder.DecodeElement(&products, &se)
			}
		default:
		}
	}

	return nil
}
