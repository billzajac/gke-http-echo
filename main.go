package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
)

func health_handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Configure the content type
	w.Header().Set("Content-Type", "text/plain")

	// Get the Project ID from metadata
	metadataHost := "metadata.google.internal"
	_, err := net.LookupHost(metadataHost)
	if err == nil {
		client := &http.Client{}
		req, err := http.NewRequest("GET", "http://"+metadataHost+"/computeMetadata/v1/project/project-id", nil)
		check(err, "building request")
		req.Header.Add("Metadata-Flavor", "Google")
		//req.Header.Add("X-Google-Metadata-Request", "True")
		resp, err := client.Do(req)
		check(err, "curling metadata request")
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		check(err, "reading metadata body")
		fmt.Fprintf(w, "Request to: %s", body)
	}

	// Get the raw request
	requestDump, err := httputil.DumpRequest(r, true)
	check(err, "reading request")

	// Write out the response
	w.Write(requestDump)
}

func check(err error, note string) {
	if err != nil {
		fmt.Println("%s: %v", note, err)
		//os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", health_handler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
