package router

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	u := r.PathPrefix("/user").Subrouter()
	userRoutes(u)
	c := r.PathPrefix("/category").Subrouter()
	categoryRoutes(c)
	pr := r.PathPrefix("/product").Subrouter()
	productRoutes(pr)
	m := r.PathPrefix("/movement").Subrouter()
	movementRoutes(m)
	w := r.PathPrefix("/warehouse").Subrouter()
	warehouseRoutes(w)
	cl := r.PathPrefix("/client").Subrouter()
	clientRoutes(cl)
}
