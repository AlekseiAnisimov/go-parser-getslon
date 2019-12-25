package xmlparser;

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "strings"
)

var xmlfile string

var xmlblob string

type Categories struct {
    Category string
}

func SetXmlFile(path string) {
    xmlfile = path
}

func ReadXmlData() error {
    openfile, err := ioutil.ReadFile(xmlfile)

    if err != nil {
    	return err;
    }

    err = Categories.UnmarshalXML(openfile)

    if err != nil {
        return err;
    }

    return nil
}

func (Categories) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var s string
    if err := d.DecodeElement(&s, &start); err != nil {
        return err
    }

    if strings.ToLower(s) == "categories" {

        fmt.Println(s)
    }
    return nil
}