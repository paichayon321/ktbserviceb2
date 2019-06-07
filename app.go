//
// developer="Phuwanai Thummavet <www.serial-coder.com>"
//

package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func serveReq(w http.ResponseWriter, r *http.Request) {

	urlPath := r.URL.Path
	extURL := os.Getenv("HOST") + ":" + os.Getenv("PORT") + urlPath

	req, err := http.NewRequest("GET", extURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Timeout = 60000 ms
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*60000)
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		return
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write([]byte(bodyString))
}

func main() {

	http.HandleFunc("/", serveReq)
	if err := http.ListenAndServe(":"+os.Getenv("SERVICE_B2_PORT"), nil); err != nil {
		panic(err)
	}
}
