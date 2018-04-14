package utils

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"strings"
)


func TestGetData(t *testing.T) {

	expectedValue := "Hello, client"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, expectedValue)
	}))
	defer ts.Close()

	result := GetData(ts.URL)
	expectedValueWithReturn := fmt.Sprintf("%s\n",expectedValue)

	if strings.Compare(string(result), expectedValueWithReturn) == 0 {
		fmt.Printf("Works as expected.")
	} else {
		t.Errorf("Expecting: %v Got: %v" +
			" Maybe look for return?",string(result),expectedValue)
	}
	fmt.Printf("%s", result)

}