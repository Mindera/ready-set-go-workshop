package webapp

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"status\": \"\"}")
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

/**
 ** e.1) INSERT YOUR CODE BELOW
 ** You should implement a logger handler and it should send to stdout:
 ** {request end date} | {request Method} | {time to execute in msec} | {request path}
 **/
