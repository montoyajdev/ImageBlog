package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

// NewUsers is used o create a new Users controller
// This function wil panic  if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// New is used to render form where user can create new user account
//GET/signup
type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type SignupForm struct {
	Email   string `schema:"email"`
	Passwod string `schema:"password"`
}

//Create  is used to process the signup form when a user
// submits it. This is used to create a newuser account
//POST/signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
