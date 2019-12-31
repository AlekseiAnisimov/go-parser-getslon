package xmlparser

import (
    "encoding/xml"
    //"io/ioutil"
    "os"
    //"fmt"
)

var xmlfile string

var xmlblob string

type Xml struct {

}

type Categories struct {
    XMLName xml.Name    `xml:"categories"`
    Category []Category   `xml:"category"`
}

type Category struct {
    Id int `xml:"id,attr"`
    ParentId int `xml:"parent_id,attr"`
    Value string `xml:",chardata"`
}

func SetXmlFile(path string) {
    xmlfile = path
}

func ReadXmlData(c *Categories) error {
    openfile, err := os.Open(xmlfile)
    if err != nil {
        return err;
    }

    decoder := xml.NewDecoder(openfile)
    var inElement string
    //var c Categories

    for {
        t, _ := decoder.Token()

        if t == nil {
            break;
        }

        switch se := t.(type) {
        case xml.StartElement:
            inElement = se.Name.Local

            if inElement == "categories" {
                decoder.DecodeElement(&c, &se)
            }
        default:
        }
    }

    //fmt.Println(c.Category)
    return nil
}