package main

import (
	    "net/http"
	    "time"
)


// **********************************************************
// Write cookies
// Cookies_write(w, "Key",  "Values",  12) 12-Days
// **********************************************************
func Cookies_write(w http.ResponseWriter, Name, Value string, Days int){
	   Expire := time.Now().AddDate(0, 0, Days)
	   http.SetCookie(w, &http.Cookie{Name:Name, Value:Value, Expires:Expire, HttpOnly: true, Path:"/"})
}

// **********************************************************
// Read cookies
// Cookies_write(w, "Key",  "Values",  12) 12-Days
// **********************************************************
func Cookies_read(r *http.Request, Name string) *http.Cookie {
	   ck, _ := r.Cookie(Name)
	   return ck
}
