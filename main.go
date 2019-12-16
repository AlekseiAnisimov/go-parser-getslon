// parser project main.go
package main

import (
	"archive/zip"
	"parser/mail"
	"path"

	//"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("Start! in: ", time.Now())
	pathFile, err := DownloadFile()
	if err != nil {
		fmt.Println(err)
	}

	err = DeleteXmlFiles()
	if err != nil {
		fmt.Println(err)
	}

	Unzip(pathFile, xmlFilePath)
}

func DownloadFile() (filePath string, err error) {
	var filepath string
	var url string

	filepath = xmlFilePath
	url = remoteURL

	resp, err := http.Get(url)
	
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		mail.MailSend()
		fmt.Println("Download error")
	}

	fullFilePath := path.Join(filepath, "file.zip")
	out, err :=  os.Create(fullFilePath)
	
	defer out.Close()
	
	_, err = io.Copy(out, resp.Body)

	return fullFilePath, nil
}

func DeleteXmlFiles() error {
	var fpath string
	fpath = xmlFilePath

	files, err := ioutil.ReadDir(fpath)

	if err != nil {
		return err
	}

	for _, file := range files {
		fileExt := filepath.Ext(file.Name())

		if fileExt == ".xml" {
			err := os.Remove(file.Name())
			
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Unzip(filepathZip string, dst string) error{
	r, err := zip.OpenReader(filepathZip)
	 
	if err != nil {
		return err
	}

	for _, f := range r.File {
		rc, err := f.Open()
		
		if err != nil {
			return err
		}
		
		defer func() {
			if rc.Close(); err != nil {
				panic(err)
			}
		}()
		
		fpath := filepath.Join(dst, f.Name)
		
		if f.FileInfo().IsDir() {
			os.Mkdir(fpath, f.Mode())
		} else {
			var fdir string
			lastIndex := strings.LastIndex(fpath, string(os.PathSeparator))

			if lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, f.Mode())
			if err != nil {
				return err
			}

			f, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
