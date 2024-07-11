package types

import "net/http"

//type UserRoutes interface {
//	handleLogIn(w http.ResponseWriter, r *http.Request)
//	handleSignUp(w http.ResponseWriter, r *http.Request)
//}

type DoctorRoutes interface {
	Login(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
}
