package controllers

import "go-morclinic/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// example
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
	//end example

	// Login V2 Route
	s.Router.HandleFunc("/api/mobile/login/uname", middlewares.SetMiddlewareJSON(s.LoginWithUsername)).Methods("POST")
	s.Router.HandleFunc("/api/mobile/login/norm", middlewares.SetMiddlewareJSON(s.LoginWithNorm)).Methods("POST")

	//Pengguna
	s.Router.HandleFunc("/api/mobile/pengguna/resetpassword", middlewares.SetMiddlewareJSON(s.LoginWithNorm)).Methods("POST")

	//AppVersion routes
	s.Router.HandleFunc("/api/mobile/app/version", middlewares.SetMiddlewareJSON(s.GetAppversion)).Methods("GET")

	//Fasilitas routes
	s.Router.HandleFunc("/api/mobile/fasilitas/all", middlewares.SetMiddlewareJSON(s.GetFasilitas)).Methods("GET")
	s.Router.HandleFunc("/api/mobile/fasilitas", middlewares.SetMiddlewareJSON(s.GetFasilitasByIdKlinik)).Methods("GET").Queries("id_klinik", "{id_klinik}")

	//News routes
	s.Router.HandleFunc("/api/mobile/news/all", middlewares.SetMiddlewareJSON(s.GetNews)).Methods("GET")
	s.Router.HandleFunc("/api/mobile/news", middlewares.SetMiddlewareJSON(s.GetNewsByIdKlinik)).Methods("GET").Queries("id_klinik", "{id_klinik}")
	s.Router.HandleFunc("/api/mobile/news/list_three", middlewares.SetMiddlewareJSON(s.GetListThreeByIdKlinik)).Methods("GET").Queries("id_klinik", "{id_klinik}")

	//Jadwal dokter routes
	s.Router.HandleFunc("/api/mobile/jadwaldokter", middlewares.SetMiddlewareJSON(s.JadwalDokter)).Methods("GET").Queries("id_klinik", "{id_klinik}").Queries("nama_hari", "{nama_hari}")

	//Klinik routes
	s.Router.HandleFunc("/api/mobile/klinik", middlewares.SetMiddlewareJSON(s.GetKlinik)).Methods("GET")
	s.Router.HandleFunc("/api/mobile/klinik/detail", middlewares.SetMiddlewareJSON(s.GetKlinikDetail)).Methods("GET").Queries("id_klinik", "{id_klinik}")
}
