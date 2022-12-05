package controllers

import "go-morclinic/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	//AppVersion routes
	s.Router.HandleFunc("/api/mobile/app/version", middlewares.SetMiddlewareJSON(s.GetAppversion)).Methods("GET")

	//Fasilitas routes
	s.Router.HandleFunc("/api/mobile/fasilitas/all", middlewares.SetMiddlewareJSON(s.GetFasilitas)).Methods("GET")
	s.Router.HandleFunc("/api/mobile/fasilitas", middlewares.SetMiddlewareJSON(s.GetFasilitasByIdKlinik)).Methods("GET").Queries("id_klinik", "{id_klinik}")

	//News routes
	s.Router.HandleFunc("/api/mobile/news/all", middlewares.SetMiddlewareJSON(s.GetNews)).Methods("GET")
	s.Router.HandleFunc("/api/mobile/news", middlewares.SetMiddlewareJSON(s.GetNewsByIdKlinik)).Methods("GET").Queries("id_klinik", "{id_klinik}")
}
