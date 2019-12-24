package xmlparser;

import (
    "encoding/xml"
)

var xmlfile string

var xmlblob 

type Category int

type Product int

func main() {
    
}

func (file *xmlfile) SetXmlFile(path string) nil {
    *file = path
    
    return nil;
}

func (file *xmlfile, blob *xmlblob) readXmlData() {
    
}