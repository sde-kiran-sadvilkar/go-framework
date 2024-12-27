package controllers

import "net/http"





func (c *Controllers) UserLogin(w http.ResponseWriter, r *http.Request) {
	err := c.Core.Render.Page(w, r, "login", nil, nil)
	if err != nil {
		c.Core.ErrorLog.Println(err)
	}
}

func (c *Controllers) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := c.Models.Users.GetByEmail(email)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	matches, err := user.CheckPassword(password)
	if err != nil {
		w.Write([]byte("Error validating password"))
		return
	}

	if !matches {
		w.Write([]byte("Invalid password!"))
		return
	}

	c.Core.Session.Put(r.Context(), "userID", user.ID)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (c *Controllers) Logout(w http.ResponseWriter, r *http.Request) {
	c.Core.Session.RenewToken(r.Context())
	c.Core.Session.Remove(r.Context(), "userID")
	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}