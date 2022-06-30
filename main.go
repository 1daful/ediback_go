/*Status Codes
1xx Information

100 Continue
101 Switching Protocols
102 Processing

2xx Success

200 OK
201 Created
202 Accepted
203 Non-authoritative Information
204 No Content
205 Reset Content
206 Partial Content
207 Multi-Status
208 Already Reported
226 IM Used

3xx Redirects

300 Multiple Choices
301 Moved Permanently
302 Found
303 See Other
304 Not Modified
305 Use Proxy
307 Temporary Redirect
308 Permanent

4xx Client Error

400 Bad Request
401 Unauthorized
402 Payment Required
403 Forbidden
404 Not Found
405 Method Not Allowed
406 Not Acceptable
407 Proxy Authentication Required
408 Request Timeout
409 Conflict
410 Gone
411 Length Required
412 Precondition Failed
413 Payload Too Large
414 Request-URI Too Long
415 Unsupported Media Type
416 Requested Range Not Satisfiable
417 Expectation Failed
418 I'm a teapot
421 Misdirected Request
422 Unprocessable Entity
423 Locked
424 Failed Dependency
426 Upgrade Required
428 Precondition Required
429 Too Many Requests
431 Request Header Fields Too Large
444 Connection Closed Without Response
451 Unavailable For Legal Reasons
499 Client Closed Request

5xx Server Error

500 Internal Server Error
501 Not Implemented
502 Bad Gateway
503 Service Unavailable
504 Gateway Timeout
505 HTTP Version Not Supported
506 Variant Also Negotiates
507 Insufficient Storage
508 Loop Detected
510 Not Extended
511 Network Authentication Required
599 Network Connect Timeout Error */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	//"strconv"
	//"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	preflight(w, r)
	url := params(w, r)
	resp, _ := run("GET", url.String())
	w.Write(resp)
}

func getSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	preflight(w, r)
	resp := searchoy()
	w.Write(resp)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func preflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Orign, Content-Type, Authorization, cache-control")
	if (*r).Method == "OPTION" {
		return
	}
}

func params(w http.ResponseWriter, r *http.Request) url.URL {
	//pathParams := mux.Vars(r)

	/*userID := -1
	  var err error
	  /*for k, v := range pathParams {
	      mp := make(map[string]interface{})
	      mp[k] = v
	  }*/

	/*if val, ok := pathParams["baseUrl"]; ok {
	    userID, err = strconv.Atoi(val)
	    if err != nil {
	        w.WriteHeader(http.StatusInternalServerError)
	        w.Write([]byte(`{"message": "need a base url"}`))
	        return
	    }
	}*/

	/*commentID := -1
	  if val, ok := pathParams["commentID"]; ok {
	      commentID, err = strconv.Atoi(val)
	      if err != nil {
	          w.WriteHeader(http.StatusInternalServerError)
	          w.Write([]byte(`{"message": "need a number"}`))
	          return
	      }
	  }*/

	query := r.URL.Query()
	//location := query.Get("location")
	baseUrl := query.Get("baseUrl")
	query.Del("baseUrl")
	//newUrl := strings.Replace(r.URL.String(), baseUrl, "", -1)
	u, _ := url.Parse(baseUrl)
	//u.Scheme = "https"
	//u.Host = baseUrl
	u.RawQuery = query.Encode()
	return *u
	//w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func run(method string, url string) ([]byte, interface{}) {
	type Response struct{}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Print(err.Error())
		log.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	/*reqBytes, err := ioutil.ReadAll(req.Body)
	if (err != nil) {
		log.Fatal(err)
	}*/

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	type Res struct {
	}
	var responseObject interface{}
	json.Unmarshal(bodyBytes, &responseObject)
	/*mai := make(map[string]interface{})
	fmt.Println(["e"])*/
	//for k, v := range responseObject {
	/*keys := make([]string, 0, len(responseObject))
	for k2 := range mai {
		keys = append(keys, k2)

	}*/
	b, _ := json.Marshal(responseObject)
	return b, responseObject
}

/*func setParam(w http.ResponseWriter, r *http.Request) {
    params := r.Header
    g := params.Get("params")
}*/

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)

	search := r.PathPrefix("/search").Subrouter()
	search.HandleFunc("", getSearch).Methods(http.MethodGet)

	handler := cors.AllowAll().Handler(r)
	//api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", handler))
}
