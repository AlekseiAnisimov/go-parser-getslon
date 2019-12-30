package xmlparser

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
)

var xmlfile string

var xmlblob string

type Xml struct {

}

type Categories struct {
    XMLName xml.Name    `xml:"categories"`
    Category []string   `xml:"category"`
}

func SetXmlFile(path string) {
    xmlfile = path
}

func ReadXmlData() error {
    openfile, err := ioutil.ReadFile(xmlfile)
    if err != nil {
    	return err;
    }

    var v Categories
    err = xml.Unmarshal([]byte(openfile), &v)

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(v.Category)
    return err
}
