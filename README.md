A normal function was converted into a HandlerFunc type and used as an HTTP handler by registering it with ServeMux.Handle. Because ordinary functions are frequently used as HTTP handlers in this way, the http package provides a shortcut method: ServeMux.HandleFunc. The HandleFunc registers the handler function for the given pattern. 
It internally (inside the http package) converts into a HandlerFunc type and registers the handler into ServeMux.

The http package provides a couple of shortcut methods for working with DefaultServeMux: http.Handle and http.HandleFunc. 
The http.Handle function registers the handler for the given
pattern in DefaultServeMux, and http.HandleFunc registers the handler function for the given pattern in DefaultServeMux. 
So these functions are just shortcuts to use ServeMux.Handle and ServeMux.HandleFunc in DefaultServeMux. 
The ListenAndServe function uses DefaultServeMux if the second parameter is set as nil instead of providing an http.Handler objec