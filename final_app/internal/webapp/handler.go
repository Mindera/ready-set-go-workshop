package webapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

const (
	healthOk = "PONG"
)

// Handler http handler
type Handler struct {
	client *redis.Client
}

func (h *Handler) healthHandler(w http.ResponseWriter, req *http.Request) {
	/**
	 ** d) INSERT YOUR CODE BELOW
	 ** Status should return OK if connection with datastore is ok;
	 ** and should return FAIL if not.
	 **/
	rsp, err := h.client.Ping().Result()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"status\": \"FAIL\", \"error\":\"%v\"}", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"status\": \"OK\", \"connected\": \"%v\"}", rsp == healthOk)
	}
}

func (h *Handler) getHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	value, err := h.client.Get(path).Result()
	if err == redis.Nil {
		http.NotFound(w, req)
		return
	} else if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value + "\n"))
}

/**
 ** c.2) INSERT YOUR CODE BELOW
 ** You should implement a function that given a key (request URL path) and a value (request body) add it to the datastore
 ** If everything goes well it should return the body; if not return http.StatusBadRequest (header).
 **/
func (h *Handler) putHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path[1:]
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b := string(body)
	err = h.client.Set(path, b, 0).Err()
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

/**
 ** e.1) INSERT YOUR CODE BELOW
 ** You should implement a logger handler and it should send to stdout:
 ** {request end date} | {request Method} | {time to execute in msec} | {request path} 
 **/
func (h *Handler) loggerHandler(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(rw, req)
		end := time.Now()

		fmt.Printf("%v | %-2s | %12v %2s\n",
			end.Format("2006/01/02 15:04:05"),
			req.Method,
			end.Sub(start),
			req.URL.Path)
	})
}
