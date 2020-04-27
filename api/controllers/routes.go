package controllers

import "github.com/enocgc/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/loginAdmin", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/registerAdmin", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// LoginCLient Route
	s.Router.HandleFunc("/loginClient", middlewares.SetMiddlewareJSON(s.LoginClient)).Methods("POST")
	//Clients routes

	s.Router.HandleFunc("/registerClient", middlewares.SetMiddlewareJSON(s.CreateUserClient)).Methods("POST")
	s.Router.HandleFunc("/clients", middlewares.SetMiddlewareJSON(s.GetUsersClient)).Methods("GET")     //exitoso
	s.Router.HandleFunc("/clients/{id}", middlewares.SetMiddlewareJSON(s.GetUserClient)).Methods("GET") //exitoso
	// s.Router.HandleFunc("/clients/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/clients/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	// parkin detail
	s.Router.HandleFunc("/addParkinDetail", middlewares.SetMiddlewareJSON(s.CreateParkinDetail)).Methods("POST")

	//Posts routes
	s.Router.HandleFunc("/postsCreate", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
