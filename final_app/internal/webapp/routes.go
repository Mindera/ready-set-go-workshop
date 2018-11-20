package webapp

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// RouteRegistry data struct for http routes
type RouteRegistry struct {
	Mux *mux.Router
	Client *redis.Client
}

// RegisterRoutes register http routes
func (r RouteRegistry) RegisterRoutes() {
	h := Handler{r.Client}

	/**
	 ** e.2) INSERT YOUR CODE BELOW
	 ** Each handler should make use of the logger handler implemented on f.1)
	 **/
	r.Mux.HandleFunc("/health", h.loggerHandler(h.healthHandler))

	r.Mux.Methods("GET").PathPrefix("/").Handler(h.loggerHandler(h.getHandler))
	/**
	 ** INSERT YOUR CODE BELOW
	 ** c.1) Add one http route with PUT http method for any given path
	 ** c.3) Add the implemented function putHandler (c.2) as a handler of PUT
	 **/
	r.Mux.Methods("PUT").PathPrefix("/").Handler(h.loggerHandler(h.putHandler))
}
