package random

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/quotes", getQuotes)
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := Asset("random.html")
	fmt.Fprint(w, string(data))
}

func getQuotes(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)
	response, err := client.Get("http://quotesondesign.com/wp-json/posts?filter[orderby]=rand&filter[posts_per_page]=1&callback=")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)

}

//go-bindata -prefix "./static/" -pkg random ./static/random.html
