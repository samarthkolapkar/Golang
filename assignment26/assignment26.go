/*Visit https://merce.co and save the homepage as an HTML file from a browser as
“merce-homepage.html”. Write a program to read the saved HTML file, compress it and store
the compressed file as “merce-homepage.html.zip”. Use ready classes for compression. At
the end, the program shall print HTML file size, Compressed file size*/
package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

//this function is used to print the size of file
func sizeoffile(filename string) int64 {
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error!!")
	}
	return info.Size()
}

//function to create the zip file
func convertZip(filename string) {
	zipfile, err := os.Create(filename + ".zip")
	if err != nil {
		fmt.Println("Error while creating the zip file!!")
	}
	defer zipfile.Close()
	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error while reading the file")
	}
	filewriter, err := zipWriter.Create(filepath.Base(filename))
	filewriter.Write(filedata)
	fmt.Println(filename + ".zip")
}

//This function is to write the content of webpage to the file
func writeToFile(filename string, data []byte) error {
	// Write data to file with read-write permissions for the owner and read-only permissions for others
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	Url := "https://pkg.go.dev/std"
	response, err := http.Get(Url) //used to load the url
	if err != nil {
		fmt.Println("Error:", err)

	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	filename := "test.html"
	writeToFile(filename, body)
	fmt.Println(filename)
	convertZip(filename)
	result := sizeoffile(filename)
	fmt.Println("Size of file is", result, "bytes")

}
