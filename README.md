A normal function was converted into a HandlerFunc type and used as an HTTP handler by registering it with ServeMux.Handle. Because ordinary functions are frequently used as HTTP handlers in this way, the http package provides a shortcut method: ServeMux.HandleFunc. The HandleFunc registers the handler function for the given pattern. 
It internally (inside the http package) converts into a HandlerFunc type and registers the handler into ServeMux.