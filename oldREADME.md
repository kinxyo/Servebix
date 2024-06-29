# Go API Tutorial

Creating REST APIs in Go 1.22.1 using only `net/http`.

<!-- need to change this -->

> [!NOTE]
> This is essentially a "copy" (for lack of a better word) of [Tiago's video](https://www.youtube.com/watch?v=npzXQSL4oWo), but with more explanation so as to even beginners could understand.

## Begin

First create 2 important files a fresh directory:

- `main.go`
- `go.mod`

Proceed to write in `go.mod`:

```go
module <name> // `name` is supposed to be a unique string here. You may use a github repo link.

go 1.22.1
```

## Understanding API structure

Let's first start with how API is written in Go.

### Basic Components

There are 2 components to writing API:-

1. **Server**:
2. **Router**:

As is the webserver architecture, you send request and you receive a response.

`Server` starts the application at the specified address then listens to the incoming requests. Once a request enters, the `router` then generates a response accordingly, and then the server sends that response back (or serves the response) to the sender.

Now, to be more accurate, the role of the router is to map a function (often referred as _handler_) to a path. So, `router` doesn't exactly handles the request; the _handlers_ are created separately for that, `router` instead maps handler to a path.

So, to create a basic webserver in Go, we will need a `server` (that **listens** to request and **serves** the response) and `router` (that generates that response).

```go
package main

import ("net/http")

func main() {

    // ROUTER (path, handler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]bytes("Hello World!"))
    })

    // SERVER
    http.ListenAndServe(":3000", nil)

}
```

In router, we bind a function to a path. `"/"` being the entry/root path, and send `"Hello World!"` as a response.

Now, don't worry about `w.Write([]bytes("Hello World!"))` if it seems complicated. It's a way of generating a response. There are numerous ways of generating response.

Right now, just know that you can run this file by running `go run main.go` in the terminal.

### Generating Response

A response can be anything from a string to html file, to json or even some form of media.

To write a string based response, We can use the following methods:-

- `W.Write()`: All communications between servers happen in bytes. This is a more efficient way for writing a response as it directly writes in bytes. It is the most preferable way to generate a string based response if you don't want to format that string.
- `fmt.Fprintf()`: This allows you to generate formatted string, meaning you can add variables to the string or even include html tags. Try the below example:

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello World</h1>")
})
```

### Handling Requests

Now as we look at the below code, we can realize if our function was very long, it would make the codebase tidy. Espcially, when we have a lot of paths to deal with in our API.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // ...
})
```

To preserve ther readability of the code, we extract the functions into a different file. And to keep all our "handler" files together, we shift them into a dedicated directory named `handler/`.

`handler/home.go`:

```go
package handler

import (
    "net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

```

Now to use this is our main file, we will specify in the import statement:

```go
import (
    "net/http"
    "<module_name>/<package_name>"
)
```

so in my case it will be:

```go
import (
    "net/http"
    "github.com/kinxyo/knowledge-box/go-api/handler"
)
```

Now, I will reference the function in the `router` this way:

```go
http.HandleFunc('/', handler.Home)
```

### Full code

With this being done, our full code looks something like this:

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/kinxyo/knowledge-box/go-api/handler"
)

func main() {

    http.HandleFunc("/", handler.Home))

    http.ListenAndServe(":8000", nil)

}

```

We have created a basic REST API server in Go.

### QoL improvement

A little quality of life improvement here would be to

Now, as we know there are multiple types of requests (such as GET, POST, DELETE),
so a convenient thing about Go is that we can easily specify our request type in the path specification itself !!!

traditionally, you would first fetch the request then apply an `if-else` condition to check its type, Go eliminates this whole boiler plate.

here's how you'd do it:

```go
package main

import (
	"fmt"
	"net/http"
	"github.com/kinxyo/knowledge-box/go-api/handler", 
)

func main() {

	http.HandleFunc("GET /", handler.Func1)
	http.HandleFunc("POST /users/{id}", handler.Func2)
	http.HandleFunc("PUT /users/{id}", handler.Func3)
	http.HandleFunc("DELETE /users/{id}", handler.Func4)
	http.HandleFunc("OPTIONS /users/{id}", handler.Func5)


	fmt.Println("Server listening on port 8000")

	http.ListenAndServe()

}
```

## More organization

### Separation of concerns

Now, starting a server isn't the only concern of `main.go` file. It can also have various other tasks to perform such as setting up database, \[...]
so it's best practice to separate the concerns of each function in its own file, hence, our logic for establishing a server can be shifted to another file.
you may name that file anything, i'm naming it `api.go` for this tutorial. \[gotta find the name convention for it].


`api.go`

```go
package main

import (
	"fmt"
	"net/http"
)

func serverRun(addr string) {

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>working :rocket: </h1>")
	})

	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Printf("Server listening on port %s\n", addr)

	server.ListenAndServe()

}

```

As you may have seen, here we use a strange function `NewServeMux` and assign its value to `router`. Why and what is happening? \[...]

`main.go`

```go
package main

func main() {

	serverRun(":3000")

}
```

### Interfacing

There can potentially be more things in `api.go` to do, so it's rather encouraged to make a data structure that stores all related information.

so we create a struct called `Server`.

```go
package main

import (
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func ServerIni(addr string) *Server { // we return an instance of the struct with the variable we want to insert; think of it as a `constructor`.
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>working :rocket: </h1>")
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started running at %s", server.Addr)

	return server.ListenAndServe()
}
```

## More functions

### Path parameter

```go
package main

import (
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func ServerIni(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("userID")
		w.Write([]byte("User ID: " + id))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started running at %s", server.Addr)

	return server.ListenAndServe()
}
```

### Wildcard

not sure if "/*" or "/" is. \[will complete it later]

### Middleware

#### logging

we'll start by creating a function:

```go
func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s | path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
```

here, \[explain the function...]

then, we wrap it around router like this \[rephrase sentence]:

```go
	server := http.Server{
		Addr:    s.addr,
		Handler: RequestLoggerMiddleware(router),
	}
```

#### auth

```
curl -H "Authorization: Bearer token" localhost:3000/users/
```

### Subrouting

---

### Directory Structure

<!-- NEEDS TO CHANGE -->

This is very basic example of directory structure needed to create a structured API. Of course, it will differ from real-world project as it's been tonned down for simplicity sake to focus only on API part.

```bash
.
├── api.go
├── go.mod
├── main.go
└── README.md
```

As can be noticed, we seperated the concerns of our API-related stuff in a seperate file from `main.go`.
A lot of configurations are needed to be written for creating an API that's serving a large system, so writing them outside of `main.go` helps us maintain more clean and modular code.

### Coding

#### Set up

`main.go`

We give it a package name `main` for the obvious reason that it is the entry point of our application.

```go
package main

func main() {
    //...
}
```

`api.go`

However, we **also** give this a package name `main` as it is an integral part of the entry point.

```go
package main

// ...
```

#### Starting with a server

`api.go`

First we create an `APIServer` struct, the purpose of which is to hold together the configuration data for the server.

```go
package main

type Server struct {
    addr string //address of sever
    // any other dependencies (such as database)
}



```