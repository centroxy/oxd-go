//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package page

import (
	"net/http"
	"fmt"
)

func LogoutPageSite(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Successfull logout")
}
