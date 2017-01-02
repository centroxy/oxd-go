//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
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