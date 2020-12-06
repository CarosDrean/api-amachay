package router

import "github.com/gorilla/mux"

func Routes(r *mux.Router) {
	u := r.PathPrefix("/user").Subrouter()
	userRoutes(u)
	c := r.PathPrefix("/category").Subrouter()
	categoryRoutes(c)
	pr := r.PathPrefix("/product").Subrouter()
	productRoutes(pr)
	i := r.PathPrefix("/input").Subrouter()
	inputRoutes(i)
	o := r.PathPrefix("/output").Subrouter()
	outputRoutes(o)
}
