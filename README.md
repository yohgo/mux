yohgo/mux
===

---

* [Install](#install)
* [Examples](#examples)

---

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/yohgo/mux
```

## Examples

Let's start by creating a list of routes:

```go
func main() {
    router := mux.NewRouter(mux.Routes{
        {Name: "Index", Method: "GET", Path: "/", HandlerFunc: IndexHandler},
        {Name: "Products", Method: "GET", Path: "/products", HandlerFunc: ProductsHandler},
        {Name: "Orders", Method: "GET", Path: "/orders", HandlerFunc: OrdersHandler},
    })
}
```
Here we register three routes mapping URL paths along with the HTTP Method to handlers. This is equivalent to how `http.HandleFunc()` works: if an incoming request URL matches one of the paths, the corresponding handler is called passing (`http.ResponseWriter`, `*http.Request`) as parameters.

After configuring the routes you can now attempt to start the server:

```go
func main() {
    router := mux.NewRouter(mux.Routes{
        {Name: "Index", Method: "GET", Path: "/", HandlerFunc: IndexHandler},
        {Name: "Products", Method: "GET", Path: "/products", HandlerFunc: ProductsHandler},
        {Name: "Orders", Method: "GET", Path: "/orders", HandlerFunc: OrdersHandler},
    })

    // Attempt to start server
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal("Failed to start server")
    }
}
```
