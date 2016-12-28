package page

import (
	"net/http"
	"fmt"
)

func LogoutPageSite(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Successfull logout")
}
