package api

import (
	"fmt"
	"net/http"
	"time"
)

func Msg(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprint(w, currentTime)
}
