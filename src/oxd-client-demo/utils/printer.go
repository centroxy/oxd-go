package utils

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func DisplayResponse(w http.ResponseWriter,v interface{}){
	value, _ := json.Marshal(v)
	fmt.Fprintln(w,string(value))
}