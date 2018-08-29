package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", myHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func myHandler(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get("http://www.fuxiben.com")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	defer res.Body.Close()
	io.Copy(w, res.Body)

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Fprint(w, err)
	// 	return
	// }
	// fmt.Fprint(w, data)
}
