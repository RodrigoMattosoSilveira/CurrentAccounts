# READY-TO-RUN MIGRATION TEMPLATE (Gin + Fiber + Reverse Proxy)

This template runs:
- Gin server on :8080 (your legacy app)
- Fiber server on :3000 (your new app)
- Unified entrypoint reverse proxy on :80
  → decides whether to forward requests to Gin or Fiber

You update the proxy routing rules as you migrate routes.

📂 Folder Structure
migration/
 ├── proxy/
 │    └── proxy.go
 ├── ginapp/
 │    └── gin.go
 ├── fiberapp/
 │    └── fiber.go
 └── main.go     (runs all three)

## proxy/proxy.go — Reverse Proxy Router
```go
package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func newProxy(target string) *httputil.ReverseProxy {
	u, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(u)
}

func StartReverseProxy() {
	ginProxy := newProxy("http://localhost:8080")
	fiberProxy := newProxy("http://localhost:3000")

	mux := http.NewServeMux()

	// ------------------------------
	// CONFIGURE MIGRATION HERE 👇
	// ------------------------------

	// NEW Fiber routes
	mux.Handle("/new/", fiberProxy)
	mux.Handle("/admin/", fiberProxy)

	// Everything else → Gin
	mux.Handle("/", ginProxy)

	// ------------------------------

	println("[PROXY] Listening on :80")
	http.ListenAndServe(":80", mux)
}
```

You will modify only this file as you migrate routes.

## ginapp/gin.go — Legacy Gin App
```go
package ginapp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartGin() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "GIN: Home")
	})

	r.GET("/old", func(c *gin.Context) {
		c.String(http.StatusOK, "GIN: Old route")
	})

	println("[GIN] Listening on :8080")
	r.Run(":8080")
}
```


Keep your existing Gin app here.

## fiberapp/fiber.go — New Fiber App
```go
package fiberapp

import (
	"github.com/gofiber/fiber/v2"
)

func StartFiber() {
	app := fiber.New()

	app.Get("/new", func(c *fiber.Ctx) error {
		return c.SendString("FIBER: New route")
	})

	app.Get("/admin/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("FIBER: Dashboard")
	})

	println("[FIBER] Listening on :3000")
	app.Listen(":3000")
}
```

Add new routes here during migration.

## main.go — Start All Servers
```go
package main

import (
	"migration/fiberapp"
	"migration/ginapp"
	"migration/proxy"
)

func main() {

	// Start Gin
	go ginapp.StartGin()

	// Start Fiber
	go fiberapp.StartFiber()

	// Start reverse proxy (entrypoint)
	proxy.StartReverseProxy()
}
```

This gives a single binary that runs both frameworks.

You can also split them into separate containers for production.

⭐ You Now Have:

✔ Gin running
✔ Fiber running
✔ A proxy routing traffic to either
✔ Full compatibility with gradual migration
✔ Zero conflict between engines

# STEP-BY-STEP ROUTE MIGRATION CHECKLIST

Use the following numbered checklist to migrate routes cleanly and systematically.

✅ Step 1 — Identify a Route to Migrate

Pick one route group to migrate at a time:

Examples:
- /users/*
- /admin/*
- /profile
- /orders/*

✅ Step 2 — Copy the Handler Logic Into Fiber

If your Gin handler:

```go
func UserList(c *gin.Context) {
    users := queryUsers()
    c.HTML(200, "users.html", gin.H{"users": users})
}
```

Make new Fiber version:

```go
func UserList(c *fiber.Ctx) error {
    users := queryUsers()
    return Render(c, "users.html", fiber.Map{"users": users})
}
```

(Using your shared template loader.)

✅ Step 3 — Add the New Route to Fiber
```go
	app.Get("/users", UserList) 
```

✅ Step 4 — Update the Reverse Proxy to Point Traffic to Fiber

Modify proxy.go:

```go
	mux.Handle("/users/", fiberProxy)
```

Now all /users/* requests bypass Gin and go directly to Fiber.

✅ Step 5 — Test in Local Environment

Check:
- HTML templates render correctly
- Middleware behaves as expected (sessions, cookies, CSRF)
- URL params behave the same
- Responses match previous output (golden-file tests recommended)

✅ Step 6 — Remove the Gin Route

Once Fiber route works perfectly:

Delete from Gin:

```go
r.GET("/users", UserList)
```

(Or keep it temporarily behind a feature flag.)

✅ Step 7 — Repeat for the Next Route

Migrate one route group at a time:

```bash
/users/*     → Fiber
/products/*  → Fiber
/admin/*     → Fiber
/auth/*      → Fiber
```

# 🧪 Optional: Golden File Regression Tests

Fiber version:

```go
resp, _ := app.Test(req)
body, _ := io.ReadAll(resp.Body)
gold := readFile("testdata/users.golden.html")

if string(body) != gold {
    t.Fatalf("output mismatch")
}
```

Ensures HTML output matches your old Gin implementation.

# 🧩 Optional: Shared Template Engine

You can load templates once and reuse in both Gin and Fiber:
```go
var Tmpl = template.Must(template.ParseGlob("templates/**/*.html"))
```

Fiber renderer:

```go
func Render(c *fiber.Ctx, name string, data any) error {
    c.Type("html", "utf-8")
    return Tmpl.ExecuteTemplate(c.Response().BodyWriter(), name, data)
}
```

Gin renderer:

```go
r.SetHTMLTemplate(Tmpl)
```

🏁 Final Notes

This architecture allows:

Feature	Supported?
Gin + Fiber together	✔ Yes
Zero-downtime migration	✔ Yes
Route-by-route migration	✔ Yes
Shared HTML templates	✔ Yes
Shared business logic	✔ Yes
No changes to infrastructure	✔ Yes

You can deploy this approach to real production environments easily.

# References
- [chatgpt](https://chatgpt.com/c/69332840-ee50-832d-a10c-cb8a8b6c0aae)