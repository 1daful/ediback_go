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
	"bytes"
	"encoding/json"
	"io"
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
	w.WriteHeader(http.StatusCreated)
	preflight(w, r)
	//url := params(w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//respByte, _ := run("GET", url.String(), nil)
	respByte, _ := run(r)
	w.Write(respByte)
	w.WriteHeader(http.StatusOK)
}

func toGet(w http.ResponseWriter, r *http.Request) ([]byte, interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	preflight(w, r)
	//url := params(w, r)
	respByte, respInterface := run(r)
	w.Write(respByte)
	w.WriteHeader(http.StatusOK)
	return respByte, respInterface
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
	preflight(w, r)
	//url := params(w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//var data map[string]interface{}
	//err := json.NewDecoder(r.Body).Decode(&data)
	/*if err != nil {
	    http.Error(w, err.Error(), http.StatusBadRequest)
	    return
	}*/
	//respByte, _ := run("POST", url.String(), r.Body)
	respByte, _ := run(r)
	w.Write(respByte)
	// Send a response back to the client
	w.WriteHeader(http.StatusOK)
}

func toPost(w http.ResponseWriter, r *http.Request) ([]byte, interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	preflight(w, r)
	//url := params(w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil, nil
	}
	req, err := makeRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, nil
	}
	respByte, respInterface := run(req)
	w.Write(respByte)
	// Send a response back to the client
	w.WriteHeader(http.StatusOK)
	return respByte, respInterface
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

/*func params(w http.ResponseWriter, r *http.Request) url.URL {
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

/*query := r.URL.Query()
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
}*/

func run(r *http.Request) ([]byte, interface{}) {
	//type Response struct{}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	/*req, err := http.NewRequest(method, url, data)

	if err != nil {
		fmt.Print(err.Error())
		log.Println(err)
	}*/

	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	//reqBytes, err := ioutil.ReadAll(req.Body)
	//if (err != nil) {
	//	log.Fatal(err)
	//}

	resp, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	var responseObject interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseObject)
	if err != nil {
		log.Println(err)
	}
	//json.Unmarshal(bodyBytes, &responseObject)
	/*mai := make(map[string]interface{})
	fmt.Println(["e"])*/
	//for k, v := range responseObject {
	/*keys := make([]string, 0, len(responseObject))
	for k2 := range mai {
		keys = append(keys, k2)

	}*/
	b, err := json.Marshal(responseObject)
	if err != nil {
		log.Println(err)
	}
	return b, responseObject
}

/*func run(method string, url string, data io.Reader) ([]byte, interface{}) {
	//type Response struct{}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		fmt.Print(err.Error())
		log.Println(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	//reqBytes, err := ioutil.ReadAll(req.Body)
	//if (err != nil) {
	//	log.Fatal(err)
	//}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	var responseObject interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseObject)
	if err != nil {
		log.Println(err)
	}
	//json.Unmarshal(bodyBytes, &responseObject)

	b, err := json.Marshal(responseObject)
	if err != nil{
		log.Println(err)
	}
	return b, responseObject
}*/

func toRun(r *http.Request) {
	//type Response struct{}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Type", "application/json")
	/*reqBytes, err := ioutil.ReadAll(req.Body)
	if (err != nil) {
		log.Fatal(err)
	}*/

	resp, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	var responseObject interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseObject)
	if err != nil {
		log.Println(err)
	}
	_, err = json.Marshal(responseObject)
	if err != nil {
		log.Println(err)
	}
}

/*func setParam(w http.ResponseWriter, r *http.Request) {
    params := r.Header
    g := params.Get("params")
}*/

func makeRequest(r *http.Request) (*http.Request, error) {
	type Request struct {
		url    string
		config map[string]string
		data   map[string]interface{}
	}
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(req.url)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for key, value := range req.config {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	// Create the HTTP request with optional body
	var body io.Reader
	reqMethod := http.MethodGet
	dataBytes, err := json.Marshal(req.data)
	if err != nil {
		return nil, err
	}
	if req.data != nil {
		body = bytes.NewBuffer(dataBytes)
		reqMethod = http.MethodPost
	}

	request, err := http.NewRequest(reqMethod, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set appropriate headers (e.g., Content-Type)
	request.Header.Set("Content-Type", "application/json") // Adjust as needed

	return request, nil
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("/schedule", schedule).Methods(http.MethodGet)

	search := r.PathPrefix("/search").Subrouter()
	search.HandleFunc("", getSearch).Methods(http.MethodGet)

	handler := cors.AllowAll().Handler(r)
	//api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":2000", handler))
}
