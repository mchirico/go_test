package utils

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)


// GetData --  basic http GET returning raw data
func GetData(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("GetData err: %v", err)
		return []byte{}
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return data

}
