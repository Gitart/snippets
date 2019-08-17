
// Session
func session(w http.ResponseWriter, r *http.Request)(sess data.Session, err error){  
cookie, err := r.Cookie("_cookie")  
 if err == nil {sess = data.Session{Uuid: cookie.Value}    
   if ok, _ := sess.Check(); !ok {err = errors.New("Invalid session")}  
  }    
return
}

// The session function retrieves the cookie from the request: 
// cookie, err := r.Cookie("_cookie")Listing 2.5session utility function in util.go

