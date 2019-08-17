
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

func index(w http.ResponseWriter, r *http.Request) { 
     threads, err := data.Threads(); if err == nil {
           _, err := session(w, r)    
            public_tmpl_files := []string{"templates/layout.html", 
                                          "templates/public.navbar.html",
                                          "templates/index.html"}    
      
            private_tmpl_files := []string{"templates/layout.html", 
                                           "templates/private.navbar.html",
                                           "templates/index.html"}    
            var templates *template.Template    
            
            if err != nil {      
               templates = template.Must(template.Parse-Files(private_tmpl_files...)) 
            } else {      
               templates = template.Must(template.ParseFiles(public_tmpl_files...))    
            }     
      templates.ExecuteTemplate(w, "layout", threads)       
}}
