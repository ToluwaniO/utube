package helper

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetContext() (context.Context, context.CancelFunc) {
	timout := 2 * time.Minute
	return context.WithTimeout(context.Background(), timout)
}

func GetBody(r *http.Request) string {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func IsBlank(s string) bool {
	trimmed := strings.TrimSpace(s)
	return len(trimmed) == 0
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}
