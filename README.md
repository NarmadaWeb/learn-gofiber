# Complete Guide to Learning Go Fiber v2 🚀

![Go Fiber Logo](https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/gofiber/fiber/v2)](https://goreportcard.com/report/github.com/gofiber/fiber/v2)
[![GoDoc](https://godoc.org/github.com/gofiber/fiber/v2?status.svg)](https://pkg.go.dev/github.com/gofiber/fiber/v2)
[![Release Version](https://img.shields.io/github/v/release/gofiber/fiber)](https://github.com/gofiber/fiber/releases)
[![License](https://img.shields.io/github/license/gofiber/fiber)](https://github.com/gofiber/fiber/blob/master/LICENSE)
[![Code Test](https://img.shields.io/github/actions/workflow/status/gofiber/fiber/test.yml?branch=master)](https://github.com/gofiber/fiber/actions/workflows/test.yml)
[![Code Coverage](https://coveralls.io/repos/github/gofiber/fiber/badge.svg?branch=master)](https://coveralls.io/github/gofiber/fiber?branch=master)

## Language [English🇬🇧](english.md) & [Indonesian🇮🇩](indonesia.md)
---
Welcome to the complete guide for learning **Go Fiber v2**, a Go web framework inspired by Express.js, built on top of [Fasthttp](https://github.com/valyala/fasthttp), the fastest HTTP engine for Go. Fiber is designed for **ease of development** with **high performance** and **low memory consumption**.

This guide aims to be a comprehensive resource for beginner to intermediate developers looking to build web applications, RESTful APIs, or microservices using Go Fiber v2.

---

## Table of Contents 📖

1.  [Introduction](#1-introduction-)
	*   [What is Go Fiber?](#what-is-go-fiber)
	*   [Why Choose Fiber?](#why-choose-fiber)
	*   [Key Features](#key-features)
	*   [Who is This Guide For?](#who-is-this-guide-for)
2.  [Prerequisites](#2-prerequisites-)
3.  [Installation](#3-installation-)
4.  [Getting Started: The Classic "Hello, World!"](#4-getting-started-the-classic-hello-world-)
	*   [Creating a New Project](#creating-a-new-project)
	*   [Basic Code](#basic-code)
	*   [Running the Application](#running-the-application)
	*   [Testing the Application](#testing-the-application)
5.  [Fiber Core Concepts](#5-fiber-core-concepts-)
	*   [Fiber Application (`fiber.App`)](#fiber-application-fiberapp)
	*   [Routing](#routing)
		*   [Basic HTTP Methods](#basic-http-methods)
		*   [Route Parameters](#route-parameters)
		*   [Optional Parameters & Wildcards](#optional-parameters--wildcards)
		*   [Route Groups](#route-groups)
		*   [Route Naming](#route-naming)
		*   [Listing Routes](#listing-routes)
	*   [Middleware](#middleware)
		*   [What is Middleware?](#what-is-middleware)
		*   [Using Built-in Middleware](#using-built-in-middleware)
		*   [Creating Custom Middleware](#creating-custom-middleware)
		*   [Middleware Execution Order](#middleware-execution-order)
		*   [Skipping Middleware (`Next`)](#skipping-middleware-next)
		*   [Middleware Scope (Global, Group, Route)](#middleware-scope-global-group-route)
		*   [Third-Party Middleware](#third-party-middleware)
	*   [Context (`fiber.Ctx`)](#context-fiberctx)
		*   [Accessing Request Information](#accessing-request-information)
		*   [Sending Responses](#sending-responses)
		*   [Passing Data (Locals)](#passing-data-locals)
		*   [Binding Request Data](#binding-request-data)
	*   [Request Handling](#request-handling-)
		*   [Reading Headers](#reading-headers)
		*   [Reading Query Parameters](#reading-query-parameters)
		*   [Reading Route Parameters](#reading-route-parameters)
		*   [Reading Request Body](#reading-request-body)
		*   [File Upload](#file-upload)
	*   [Response Handling](#response-handling-)
		*   [Setting Status Code](#setting-status-code)
		*   [Setting Headers](#setting-headers)
		*   [Sending Various Data Types](#sending-various-data-types)
	*   [Error Handling](#error-handling-)
		*   [Returning Errors from Handlers](#returning-errors-from-handlers)
		*   [Custom Errors (`fiber.NewError`)](#custom-errors-fibernewerror)
		*   [Custom Error Handler](#custom-error-handler)
		*   [Recover Middleware](#recover-middleware)
	*   [Configuration (`fiber.Config`)](#configuration-fiberconfig)
		*   [Common Configuration Options](#common-configuration-options)
		*   [Prefork Configuration](#prefork-configuration)
	*   [Template Engine](#template-engine-)
		*   [Template Engine Concepts](#template-engine-concepts)
		*   [Using Built-in Template Engine (HTML)](#using-built-in-template-engine-html)
		*   [Using Other Template Engines](#using-other-template-engines)
		*   [Layouts](#layouts)
	*   [Serving Static Files](#serving-static-files-)
	*   [Request Validation](#request-validation-)
		*   [Importance of Validation](#importance-of-validation)
		*   [Using a Validator Library](#using-a-validator-library)
		*   [Implementation Example](#implementation-example)
6.  [Advanced Topics](#6-advanced-topics-)
	*   [WebSocket](#websocket)
	*   [Server-Sent Events (SSE)](#server-sent-events-sse)
	*   [Database Integration](#database-integration)
	*   [Authentication & Authorization (JWT, Sessions)](#authentication--authorization-jwt-sessions)
	*   [Testing](#testing)
	*   [Deployment](#deployment)
	*   [Performance & Optimization](#performance--optimization)
	*   [Project Structure](#project-structure)
	*   [Graceful Shutdown](#graceful-shutdown)
7.  [Example Application (Simple CRUD)](#7-example-application-simple-crud-)
8.  [API Documentation](#8-api-documentation-)
9.  [Best Practices](#9-best-practices-)
10. [Contributing](#10-contributing-)
11. [License](#11-license-)
12. [Acknowledgements](#12-acknowledgements-)

---

## 1. Introduction 🌟

### What is Go Fiber?

Go Fiber is a web framework for the Go programming language strongly inspired by [Express.js](https://expressjs.com/), the highly popular Node.js framework. Fiber is built on top of [Fasthttp](https://github.com/valyala/fasthttp), a high-performance HTTP library for Go. Its goal is to provide a familiar and easy-to-use interface for developers who might be accustomed to Express, while leveraging the speed and efficiency of Go and Fasthttp.

Fiber focuses on:

*   **High Performance:** Utilizes Fasthttp to achieve very high throughput and low latency.
*   **Low Memory Allocation:** Designed to minimize memory allocation during request processing.
*   **Ease of Use:** An expressive and easy-to-learn API, especially if you have an Express.js background.
*   **Flexibility:** A rich middleware ecosystem and extensibility.

### Why Choose Fiber?

There are several compelling reasons to choose Fiber for your next Go web project:

1.  **Incredible Speed:** Thanks to Fasthttp, Fiber is one of the fastest Go frameworks available. This is crucial for high-traffic applications or those requiring minimal latency.
2.  **Rapid Development:** The intuitive, Express-like API allows you to build applications quickly. Good documentation and an active community also help.
3.  **Efficient Memory Usage:** Important for applications running in resource-constrained environments or for reducing hosting costs.
4.  **Middleware Ecosystem:** Fiber provides many built-in middlewares (like logger, recovery, CORS) and makes it easy to integrate or create custom middleware.
5.  **Expressive Routing:** A robust routing system supports parameters, wildcards, groups, and route naming.
6.  **Template Engine Support:** Easily integrates with various Go template engines.
7.  **WebSocket & SSE Support:** Has built-in support or official middleware for real-time communication.
8.  **API Focused:** Well-suited for building modern RESTful APIs.

### Key Features

*   Robust routing
*   Static file serving
*   Middleware & Next() support
*   Express-inspired API
*   Template Engine Support (Go, Pug, Amber, etc.)
*   Rapidly growing middleware ecosystem
*   WebSocket support
*   Server-Sent Events (SSE)
*   Rate Limiter
*   Built on Fasthttp
*   Easy configuration
*   And much more...

### Who is This Guide For?

*   **Beginner Go Developers:** Who want to learn to build web applications with Go using a modern framework.
*   **Experienced Developers (from other languages):** Such as Node.js (Express), Python (Flask/Django), Ruby (Rails) who want to switch to Go for web development.
*   **Intermediate Go Developers:** Who want to delve deeper into Fiber's features and best practices.

It is expected that you have a basic understanding of the Go language (syntax, data types, functions, structs, interfaces) and fundamental HTTP concepts (request, response, methods, status codes).

---

## 2. Prerequisites 🛠️

Before you start learning and using Go Fiber, make sure you have:

1.  **Go Installed:** Go version 1.17 or newer is recommended. You can download and install Go from the [official website](https://go.dev/dl/).
	*   Verify the installation by opening a terminal and running: `go version`
2.  **Basic Understanding of Go:**
	*   Basic Go syntax (variable declaration, data types, control flow like `if`, `for`).
	*   Functions and methods.
	*   Structs and Interfaces.
	*   Goroutines and Channels (a basic understanding helps, but isn't mandatory to start).
	*   Package and module management (`go mod`).
3.  **Basic Understanding of Web & HTTP:**
	*   Client-Server model.
	*   HTTP Request & Response.
	*   HTTP Methods (GET, POST, PUT, DELETE, etc.).
	*   HTTP Status Codes (200 OK, 404 Not Found, 500 Internal Server Error, etc.).
	*   HTTP Headers.
	*   JSON (JavaScript Object Notation) as a common data exchange format.
4.  **Terminal/Command Prompt:** You will frequently use the terminal to run Go commands (build, run, test) and other tools like `curl`.
5.  **Text Editor or IDE:** Choose your favorite code editor (VS Code, GoLand, Vim, Emacs, etc.) that has good Go support.

---

## 3. Installation ⚙️

Installing Fiber is very easy using Go's module system.

1.  **Create a Project Directory:**
	```bash
	mkdir my-fiber-project
	cd my-fiber-project
	```

2.  **Initialize Go Modules:**
	If you're starting a new project, initialize Go modules. Replace `your-module-name` with an appropriate module path (e.g., `github.com/username/my-fiber-project`).
	```bash
	go mod init your-module-name
	```
	This will create a `go.mod` file in your project directory.

3.  **Add Fiber as a Dependency:**
	Run the following command to download and add Fiber v2 to your project:
	```bash
	go get -u github.com/gofiber/fiber/v2
	```
	This command will:
	*   Download the Fiber package and its dependencies.
	*   Add `github.com/gofiber/fiber/v2` as a `require` in your `go.mod` file.
	*   Create or update the `go.sum` file containing dependency checksums.

You are now ready to start using Fiber in your Go code!

---

## 4. Getting Started: The Classic "Hello, World!" 👋

Let's create the simplest Fiber application to ensure everything is working correctly.

### Creating a New Project

If you haven't already, follow the steps in the [Installation](#3-installation-) section.

### Basic Code

Create a new file named `main.go` in your project directory and add the following code:

```go
// main.go
package main

import (
	"log" // Package for logging

	"github.com/gofiber/fiber/v2" // Import Fiber v2 package
)

func main() {
	// 1. Create a new Fiber app instance
	//    We can pass custom configuration here,
	//    but for now, we'll use the defaults.
	app := fiber.New()

	// 2. Define a route for HTTP GET requests to the path "/" (root)
	//    When a GET request comes to "/", the anonymous handler function will be executed.
	app.Get("/", func(c *fiber.Ctx) error {
		// 'c' is a pointer to the Fiber Context, which holds request information
		// and provides methods for sending responses.

		// Send a simple string response "Hello, World! 👋" to the client.
		// The SendString method automatically sets the Content-Type to text/plain.
		return c.SendString("Hello, World! 👋")
	})

	// 3. Start the HTTP server on port 3000
	//    The server will listen for incoming connections on "0.0.0.0:3000".
	//    app.Listen() is a blocking call, meaning the program will stop here
	//    and keep running until the server is stopped (e.g., with Ctrl+C).
	log.Fatal(app.Listen(":3000"))
	// We use log.Fatal to catch any errors that might occur during server startup
	// (e.g., port already in use) and stop the program.
}
```

**Code Explanation:**

1.  **`import`**: We import the standard Go `log` package for logging messages (especially server start errors) and the `fiber` package itself.
2.  **`fiber.New()`**: Creates a new instance of the Fiber application. This is the starting point for every Fiber app. You can pass a `fiber.Config` struct to this function to customize the app's behavior (discussed later).
3.  **`app.Get("/", ...)`**: Defines a *route*. It tells Fiber that when an HTTP request with the `GET` method comes to the `/` path, the provided *handler* function (an anonymous function in this case) should be executed.
4.  **`func(c *fiber.Ctx) error`**: This is the *route handler*. All handlers in Fiber receive a pointer to `fiber.Ctx` (Context) and return an `error`. The Context (`c`) provides access to request data (like headers, parameters, body) and methods to build and send the response. If the handler returns `nil`, it means the request was processed successfully. If it returns an error, Fiber will handle it (usually by sending an error response to the client).
5.  **`c.SendString(...)`**: A method on the Context to send a string response with a `200 OK` status code and `Content-Type: text/plain`.
6.  **`app.Listen(":3000")`**: Starts the HTTP server and makes it listen for connections on port 3000 on all network interfaces (`:3000` is shorthand for `0.0.0.0:3000`). This function blocks execution, keeping the server running to accept requests.
7.  **`log.Fatal(...)`**: If `app.Listen` returns an error during startup (e.g., port 3000 is already in use), `log.Fatal` prints the error message to the console and terminates the program.

### Running the Application

Open your terminal, navigate to your project directory (`my-fiber-project`), and run the command:

```bash
go run main.go
```

You should see output similar to this (the Fiber logo and port information):

```
┌───────────────────────────────────────────────────┐
│                   Fiber v2.xx.x                   │
│               http://127.0.0.1:3000               │
│       (bound on host 0.0.0.0 and port 3000)       │
│                                                   │
│ Handlers ........... 1  Processes ........... 1   │
│ Prefork ....... Disabled  PID ............. xxxxx │
└───────────────────────────────────────────────────┘
```

This indicates that your Fiber server is running and ready to accept requests on port 3000.

### Testing the Application

There are two easy ways to test:

1.  **Using a Web Browser:** Open your browser and navigate to `http://localhost:3000` or `http://127.0.0.1:3000`. You should see the text "Hello, World! 👋" displayed on the page.

2.  **Using `curl` (from another terminal):** Open a new terminal window (leave the server running in the first one) and run:
	```bash
	curl http://localhost:3000
	```
	You will get the output:
	```
	Hello, World! 👋
	```

Congratulations! You have successfully created and run your first Go Fiber application.

---

## 5. Fiber Core Concepts 🧠

This section covers the fundamental concepts you need to understand to work effectively with Fiber.

### Fiber Application (`fiber.App`)

The `*fiber.App` object you create with `fiber.New()` is the core of your application. This object is used to:

*   Register routes.
*   Apply middleware.
*   Configure application settings.
*   Start the HTTP server.

```go
// Create an instance with default configuration
app := fiber.New()

// Create an instance with custom configuration
appWithConfig := fiber.New(fiber.Config{
	AppName:      "My Cool App v1.0",
	Prefork:      true, // Enable prefork mode (discussed later)
	ErrorHandler: myCustomErrorHandler, // Set a custom error handler
})
```

### Routing

Routing is the process of determining how an application responds to a client request to a specific endpoint (URI or path) and HTTP method (GET, POST, etc.). Fiber provides a very flexible and fast routing system.

#### Basic HTTP Methods

You can register routes for all standard HTTP methods using the corresponding methods on the `fiber.App` instance:

```go
app := fiber.New()

// GET: Retrieve data
app.Get("/users", func(c *fiber.Ctx) error {
	return c.SendString("Retrieving user list (GET)")
})

// POST: Create new data
app.Post("/users", func(c *fiber.Ctx) error {
	return c.SendString("Creating a new user (POST)")
})

// PUT: Update data entirely
app.Put("/users/:id", func(c *fiber.Ctx) error {
	id := c.Params("id") // Get the 'id' parameter from the URL
	return c.SendString("Updating user with ID: " + id + " (PUT)")
})

// PATCH: Update partial data
app.Patch("/users/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Partially updating user data ID: " + id + " (PATCH)")
})

// DELETE: Remove data
app.Delete("/users/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.SendString("Deleting user with ID: " + id + " (DELETE)")
})

// OPTIONS: Get allowed communication options
app.Options("/info", func(c *fiber.Ctx) error {
	c.Set("Allow", "GET, POST, OPTIONS")
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
})

// HEAD: Same as GET but without the response body (only headers)
app.Head("/status", func(c *fiber.Ctx) error {
	c.Set("X-App-Status", "OK")
	return c.SendStatus(fiber.StatusOK) // 200 OK (without body)
})

// app.All() to match all HTTP methods
app.All("/universal", func(c *fiber.Ctx) error {
	return c.SendString("This endpoint responds to all HTTP methods.")
})
```

#### Route Parameters

Often, you need to capture dynamic segments from the URL, like a user ID or an article slug. Fiber allows this with *route parameters*, prefixed with a colon (`:`).

```go
// Route: /users/:userId/books/:bookId
app.Get("/users/:userId/books/:bookId", func(c *fiber.Ctx) error {
	// Get parameter values using c.Params("parameterName")
	userId := c.Params("userId")
	bookId := c.Params("bookId")

	return c.SendString("User ID: " + userId + ", Book ID: " + bookId)
	// Example request: GET /users/123/books/abc
	// Output: User ID: 123, Book ID: abc
})
```

Fiber also provides slightly more efficient ways to retrieve parameters if you know their data type:

```go
app.Get("/product/:id", func(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // Try to parse the "id" parameter as an integer
	if err != nil {
		// If the parameter is not a valid integer, send a 400 error
		return c.Status(fiber.StatusBadRequest).SendString("Product ID must be a number")
	}
	// ... process with id (type int) ...
	return c.JSON(fiber.Map{"product_id": id, "status": "found"})
})
```
`ParamsInt`, `ParamsBool`, `ParamsFloat` methods are available.

#### Optional Parameters & Wildcards

*   **Optional Parameter:** Mark a parameter with a question mark (`?`) to make it optional. Your handler needs to check if the parameter exists.
	```go
	// Route: /order/:id?
	app.Get("/order/:id?", func(c *fiber.Ctx) error {
		id := c.Params("id") // Will be empty if not present in URL
		if id == "" {
			return c.SendString("Displaying all orders")
			// Request: GET /order
		}
		return c.SendString("Displaying order details ID: " + id)
		// Request: GET /order/55
	})
	```

*   **Wildcard (`*`)**: Matches anything (including `/`). Useful for capturing long or variable paths. The value is retrieved with `c.Params("*")`.
	```go
	// Route: /files/*
	app.Get("/files/*", func(c *fiber.Ctx) error {
		filePath := c.Params("*") // Gets everything after /files/
		return c.SendString("Accessing file at path: " + filePath)
		// Request: GET /files/images/logo.png -> Output: Accessing file at path: images/logo.png
		// Request: GET /files/docs/report.pdf -> Output: Accessing file at path: docs/report.pdf
	})
	```
	*   **Important:** The wildcard `*` must be at the end of the route path.

*   **Wildcard Parameter (`+`)**: Similar to `*`, but *must* match at least one character.
	```go
	// Route: /user/+
	app.Get("/user/+", func(c *fiber.Ctx) error {
		name := c.Params("+") // Gets everything after /user/
		return c.SendString("User profile: " + name)
		// Request: GET /user/johndoe -> Output: User profile: johndoe
		// Request: GET /user/jane/doe (won't match if StrictRouting is on, matches if off)
		// Request: GET /user/ (won't match because + requires at least 1 character)
	})
	```
	*   **Difference between `*` and `+`**: `*` can match an empty string (if at the end of the URL), `+` cannot. `*` matches everything including `/`, `+` is typically used for a single non-empty path segment.

#### Route Groups

Route groups are very useful for organizing routes that share a common path prefix or middleware.

```go
app := fiber.New()

// Create a group for all routes under /api/v1
apiV1 := app.Group("/api/v1")

// Add middleware specific to this group (e.g., authentication)
apiV1.Use(func(c *fiber.Ctx) error {
	log.Println("API V1 Middleware executed!")
	// Check authentication headers here...
	return c.Next() // Continue to the next handler/middleware
})

// Routes within the group (path relative to the group prefix)
// Handler for GET /api/v1/status
apiV1.Get("/status", func(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "API v1 OK"})
})

// Handler for GET /api/v1/users
apiV1.Get("/users", func(c *fiber.Ctx) error {
	// ... logic to fetch users ...
	return c.JSON([]fiber.Map{{"id": 1, "name": "User One"}, {"id": 2, "name": "User Two"}})
})

// You can create groups within groups (nested groups)
admin := apiV1.Group("/admin")
admin.Use(adminAuthMiddleware) // Admin-specific middleware

// Handler for POST /api/v1/admin/settings
admin.Post("/settings", func(c *fiber.Ctx) error {
	// ... logic to save admin settings ...
	return c.JSON(fiber.Map{"message": "Admin settings saved"})
})
```
Using `app.Group()` makes the code more structured, readable, and reduces redundancy (e.g., not having to type `/api/v1` repeatedly or apply the same middleware to many routes manually).

#### Route Naming

You can give names to your routes. This is useful if you need to generate URLs dynamically elsewhere in your application (e.g., in templates or redirects).

```go
// Give the name "user.profile" to the route
app.Get("/users/:id/profile", func(c *fiber.Ctx) error {
	// ...
	return c.SendString("This is the profile page")
}).Name("user.profile")

// In another handler, generate a URL for the named route
app.Get("/dashboard", func(c *fiber.Ctx) error {
	// Generate URL for user with ID 123
	profileURL, err := c.GetRouteURL("user.profile", fiber.Map{
		"id": "123", // Provide value for the :id parameter
	})
	if err != nil {
		return err // Handle error if route not found or params mismatch
	}
	// profileURL will contain "/users/123/profile"

	// Redirect to the user's profile page
	// return c.Redirect(profileURL)

	return c.SendString("User 123 Profile URL: " + profileURL)
})
```
Route naming improves code maintainability because you don't need to hardcode URLs in multiple places. If the route path changes, you only need to change it in one place (the route definition), and all `GetRouteURL` calls will automatically generate the correct URL.

#### Listing Routes

For debugging or introspection purposes, you can get a list of all routes registered in your Fiber application.

```go
app.Get("/debug/routes", func(c *fiber.Ctx) error {
	// Get a slice of all routes
	routes := app.GetRoutes(true) // true to include Fiber's internal routes

	// Format the output (e.g., as JSON)
	var routeList []fiber.Map
	for _, route := range routes {
		routeList = append(routeList, fiber.Map{
			"method": route.Method,
			"name":   route.Name,
			"path":   route.Path,
			"params": route.Params,
		})
	}
	return c.JSON(routeList)
})

// Run the server and access /debug/routes to see the result
```

### Middleware

Middleware is one of the most powerful concepts in modern web frameworks like Fiber. Middleware is a function that has access to the **Context (`fiber.Ctx`)** object and the **`Next()`** function in the application's request-response cycle.

#### What is Middleware?

Middleware functions act as "gatekeepers" or "intermediate processors" that execute **before** or **after** your main route handlers. Their tasks can vary widely:

*   **Logging:** Recording details of every incoming request.
*   **Authentication/Authorization:** Checking user credentials or access permissions.
*   **Data Validation/Sanitization:** Validating input data or cleaning it up.
*   **Compression:** Compressing the response body (e.g., with Gzip).
*   **CORS (Cross-Origin Resource Sharing):** Setting headers to allow APIs to be accessed from different domains.
*   **Rate Limiting:** Limiting the number of requests from a single IP within a specific time period.
*   **Header Manipulation:** Adding, modifying, or removing request/response headers.
*   **Error Handling:** Catching errors that occur in subsequent handlers.
*   **Caching:** Caching responses for identical requests.

#### Using Built-in Middleware

Fiber comes with many ready-to-use, common middlewares. You can find them in the `github.com/gofiber/fiber/v2/middleware/...` package.

The way to use them is by calling `app.Use()` or the `Use()` method on a route group.

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger" // Import logger middleware
	"github.com/gofiber/fiber/v2/middleware/recover" // Import recover middleware
	"github.com/gofiber/fiber/v2/middleware/cors"    // Import CORS middleware
)

func main() {
	app := fiber.New()

	// 1. Use Middleware Globally (applies to all routes)
	app.Use(recover.New()) // Middleware to catch panics and send 500 Internal Server Error
	app.Use(logger.New(logger.Config{ // Middleware for request logging
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New()) // Middleware to enable CORS with default config

	// Your Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello with Middleware!")
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Oops, a panic occurred!") // The recover middleware will catch this
	})

	log.Fatal(app.Listen(":3000"))
}
```

When you run this application and make a request (e.g., `curl http://localhost:3000`), you will see the request log in your terminal because `logger.New()` is active. If you access `/panic`, you won't see the server crash; instead, you'll get a `500 Internal Server Error` response because `recover.New()` caught the panic.

Other popular built-in middlewares:

*   `basicauth`: HTTP Basic authentication.
*   `compress`: Response compression (Gzip, Deflate, Brotli).
*   `etag`: ETag header generation for caching.
*   `limiter`: Rate limiting.
*   `monitor`: Displays application performance metrics.
*   `pprof`: Go application profiling.
*   `requestid`: Adds a unique ID to each request.
*   `session`: Session management.
*   And many more (see the Fiber documentation).

#### Creating Custom Middleware

You can easily create your own middleware. A custom middleware is just a function that follows the `func(*fiber.Ctx) error` signature.

```go
package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Example Custom Middleware: Adds X-Request-Time header
func TimerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now() // Record start time

		// Call the next handler/middleware in the chain
		// This is crucial! Without c.Next(), the request won't reach the route handler.
		err := c.Next()

		// Code here executes AFTER the route handler finishes
		stop := time.Now()
		duration := stop.Sub(start)

		// Add a custom header to the response
		c.Set("X-Request-Time", duration.String())
		log.Printf("Request to %s took %s", c.Path(), duration)

		// Return any error from the subsequent handler/middleware
		return err
	}
}

// Example Custom Middleware: Simple API Key Header Check
func APIKeyAuthMiddleware(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the value of the 'X-API-Key' header from the request
		key := c.Get("X-API-Key")

		if key == "" {
			log.Println("Auth Middleware: X-API-Key header not found")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "X-API-Key header is required",
			})
		}

		if key != apiKey {
			log.Printf("Auth Middleware: Invalid API Key: %s", key)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid API Key",
			})
		}

		// If API Key is valid, proceed to the next handler
		log.Println("Auth Middleware: API Key valid")
		return c.Next()
	}
}

func main() {
	app := fiber.New()

	// Use built-in logger middleware
	app.Use(logger.New())

	// Use custom timer middleware (globally)
	app.Use(TimerMiddleware())

	// Public route
	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(50 * time.Millisecond) // Simulate work
		return c.SendString("Public Page")
	})

	// Group routes that require an API Key
	api := app.Group("/api")
	// Apply API Key middleware only to the /api group
	api.Use(APIKeyAuthMiddleware("secret-key-123"))

	api.Get("/data", func(c *fiber.Ctx) error {
		time.Sleep(100 * time.Millisecond) // Simulate work
		return c.JSON(fiber.Map{"message": "This is your secret data!"})
	})

	log.Fatal(app.Listen(":3000"))
}
```

**Explanation:**

*   **`TimerMiddleware`**: Records the time before calling `c.Next()` and after it returns, then calculates the duration and adds it as a response header.
*   **`APIKeyAuthMiddleware`**: Receives the valid `apiKey` when created. Inside, it checks the `X-API-Key` header. If it's missing or doesn't match, it sends a `401 Unauthorized` response and *does not* call `c.Next()`, stopping the request. If it matches, it calls `c.Next()` so the request can proceed to the `/api/data` route handler.
*   Note how `TimerMiddleware` is applied globally (`app.Use`), while `APIKeyAuthMiddleware` is only applied to the `/api` group (`api.Use`).

#### Middleware Execution Order

Middleware executes in the order they are added using `app.Use()` or `group.Use()`.

```go
app.Use(Middleware1)
app.Use(Middleware2)

api := app.Group("/api")
api.Use(Middleware3)

api.Get("/test", Handler)

// Execution order for GET /api/test:
// 1. Middleware1
// 2. Middleware2
// 3. Middleware3
// 4. Handler
```

#### Skipping Middleware (`Next`)

The `c.Next()` function is crucial within middleware. Its purpose is to **pass control to the next middleware or handler** in the execution chain.

*   If a middleware **calls `c.Next()`**, execution continues to the next function. Code after `c.Next()` in the middleware will execute *after* the subsequent handlers/middlewares complete.
*   If a middleware **does not call `c.Next()`**, the request-response cycle stops at that middleware. That middleware is then fully responsible for sending a response to the client. This is useful for cases like failed authentication, failed validation, or caching hits.

#### Middleware Scope (Global, Group, Route)

You can apply middleware at different levels:

1.  **Global:** Using `app.Use()`. This middleware runs for **every** incoming request to the application, before route matching occurs (except for middleware like `Static`).
	```go
	app.Use(logger.New()) // Runs for all requests
	```

2.  **Group:** Using `group.Use()`. This middleware runs only for requests matching the group's prefix.
	```go
	admin := app.Group("/admin")
	admin.Use(requireAdminLogin) // Runs only for routes under /admin
	admin.Get("/dashboard", ...)
	```

3.  **Route-Specific:** You can pass one or more middleware functions *before* the main route handler when defining the route.
	```go
	func rateLimitMiddleware(c *fiber.Ctx) error { /* ... */ return c.Next() }
	func specificAuthMiddleware(c *fiber.Ctx) error { /* ... */ return c.Next() }
	func finalHandler(c *fiber.Ctx) error { /* ... */ }

	// rateLimitMiddleware and specificAuthMiddleware will run before finalHandler
	// only for POST requests to /submit
	app.Post("/submit", rateLimitMiddleware, specificAuthMiddleware, finalHandler)
	```

#### Third-Party Middleware

Besides the built-in middleware, many middlewares are developed by the Fiber community or other Go developers. You can find them on GitHub or other sources. Using them is typically the same: import the package and use it with `app.Use()`.

Make sure to read the documentation for third-party middleware to understand its behavior and configuration.

### Context (`fiber.Ctx`)

The `*fiber.Ctx` object is the heart of request handling in Fiber. Every handler and middleware function receives a pointer to this object. `Ctx` provides everything you need to:

*   Access information about the incoming request (method, path, headers, query params, route params, body, client IP, etc.).
*   Send a response back to the client (set status code, headers, send body in various formats).
*   Pass data between middleware and handlers within a single request cycle.
*   Manage the request lifecycle (e.g., calling `Next()`).

Let's look at some of the most important `Ctx` methods:

#### Accessing Request Information

*   **`c.Method()`**: Gets the HTTP method (string, e.g., "GET", "POST").
*   **`c.Path()`**: Gets the request path (string, e.g., "/users/123").
*   **`c.BaseURL()`**: Gets the base URL (e.g., "http://example.com").
*   **`c.OriginalURL()`**: Gets the original URL including the query string.
*   **`c.Hostname()`**: Gets the hostname from the `Host` header.
*   **`c.IP()`**: Gets the client's IP address (considers proxy headers like `X-Forwarded-For`).
*   **`c.IPs()`**: Gets the list of IPs if proxies are involved (from `X-Forwarded-For`).
*   **`c.Protocol()`**: Gets the request protocol (string, e.g., "http", "https").
*   **`c.Secure()`**: Checks if the connection uses HTTPS (boolean).
*   **`c.Get(key string, defaultValue ...string)`**: Gets a request header value. `key` is case-insensitive.
	```go
	ua := c.Get("User-Agent")
	auth := c.Get("Authorization", "default_value_if_not_found")
	```
*   **`c.Params(key string, defaultValue ...string)`**: Gets a route parameter value.
	```go
	userID := c.Params("id")
	```
*   **`c.ParamsInt(key string)`**, **`c.ParamsFloat(key string)`**, **`c.ParamsBool(key string)`**: Gets a route parameter and converts it to the corresponding type. Returns an error if conversion fails.
*   **`c.Query(key string, defaultValue ...string)`**: Gets a query parameter value from the URL.
	```go
	// URL: /search?q=fiber&page=2
	searchTerm := c.Query("q") // "fiber"
	page := c.Query("page", "1") // "2" (defaults to "1" if not present)
	```
*   **`c.QueryParser(out interface{}) error`**: Parses the query string into a Go struct. Useful for complex search/filter parameters.
	```go
	type SearchQuery struct {
		Query string `query:"q"`
		Limit int    `query:"limit"`
		Page  int    `query:"page"`
	}
	var sq SearchQuery
	if err := c.QueryParser(&sq); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid query")
	}
	// sq.Query, sq.Limit, sq.Page will be populated
	```
*   **`c.Body()`**: Gets the request body as `[]byte`. Use this if you need raw access to the body.
	```go
	rawBody := c.Body()
	// Be careful: Reading the body can consume memory for large bodies.
	// Consider BodyParser or body size limits.
	```
*   **`c.BodyParser(out interface{}) error`**: Parses the request body into a Go struct. It automatically detects the `Content-Type` (JSON, XML, form) and performs unmarshaling. This is the most common and recommended way to handle input data.
	```go
	type CreateUserInput struct {
		Name  string `json:"name" xml:"name" form:"name"`
		Email string `json:"email" xml:"email" form:"email"`
	}
	var input CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body: " + err.Error()})
	}
	// input.Name and input.Email will be populated
	```
*   **`c.FormValue(key string, defaultValue ...string)`**: Gets a value from a form (application/x-www-form-urlencoded or multipart/form-data).
*   **`c.FormFile(key string)`**: Gets an uploaded file from a multipart/form-data form. Returns `*multipart.FileHeader`.
*   **`c.SaveFile(fileheader *multipart.FileHeader, path string)`**: Saves an uploaded file to the specified path.
*   **`c.Is(contentType string)`**: Checks if the request `Content-Type` matches (e.g., `c.Is("json")`).
*   **`c.Accepts(offers ...string)`**: Checks the client's `Accept` header and determines the best supported content type (e.g., `c.Accepts("json", "html")`).

#### Sending Responses

*   **`c.SendStatus(statusCode int)`**: Sends a response with only a status code (no body).
	```go
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
	```
*   **`c.Status(statusCode int)`**: Sets the status code for the next response. Useful for chaining with body-sending methods.
	```go
	return c.Status(fiber.StatusCreated).JSON(newUser) // 201 Created
	```
*   **`c.Set(key string, val string)`**: Sets a response header.
	```go
	c.Set("X-Custom-Header", "My Value")
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON) // Another way to set Content-Type
	```
*   **`c.Append(key string, values ...string)`**: Appends a value to an existing header (e.g., `Link` or `Set-Cookie`).
*   **`c.SendString(body string)`**: Sends a string response with `Content-Type: text/plain`.
*   **`c.Send(body []byte)`**: Sends a response body as a byte slice. `Content-Type` will be detected automatically (if possible) or default to `application/octet-stream`.
*   **`c.JSON(data interface{})`**: Sends a JSON response. Converts `data` (struct, map, slice) to JSON and sets `Content-Type: application/json`. Very commonly used for APIs.
	```go
	user := User{ID: 1, Name: "Fiber"}
	return c.JSON(user)
	// or
	return c.JSON(fiber.Map{"status": "success", "data": user})
	```
*   **`c.XML(data interface{})`**: Sends an XML response.
*   **`c.Render(name string, bind interface{}, layouts ...string)`**: Renders an HTML template (requires template engine configuration).
*   **`c.SendFile(filepath string, compress ...bool)`**: Sends a file as the response. `Content-Type` is usually detected from the file extension. The `compress` option (default true) enables Gzip compression if the client supports it.
	```go
	return c.SendFile("./public/images/logo.png")
	```
*   **`c.Download(filepath string, filename ...string)`**: Similar to `SendFile`, but adds a `Content-Disposition: attachment` header, telling the browser to download the file instead of displaying it. You can provide a custom download filename.
	```go
	return c.Download("./private/report.pdf", "Monthly Report.pdf")
	```
*   **`c.Redirect(location string, status ...int)`**: Sends a redirect response (default status 302 Found).
	```go
	return c.Redirect("/login", fiber.StatusTemporaryRedirect) // 307
	```
*   **`c.Cookie(cookie *fiber.Cookie)`**: Sets a response cookie.
	```go
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "random-session-string",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true, // Only send over HTTPS
		SameSite: "Lax",
	})
	```
*   **`c.ClearCookie(key ...string)`**: Clears a cookie from the client's browser.

#### Passing Data (Locals)

Sometimes you need to pass data from one middleware to another or to the main route handler within the *same request cycle*. For example, an authentication middleware might verify a user and then pass the user ID or user object to the handler. `c.Locals()` is the way to do this.

`c.Locals()` acts like a `map[string]interface{}` bound to that specific request context.

```go
// Middleware: Get user data (e.g., from token)
func UserAuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	// ... validate token and get user info ...
	user := User{ID: 123, Role: "admin"} // Example user data

	// Store user data in Locals
	c.Locals("currentUser", user)
	c.Locals("requestID", "xyz-789") // Can store other data types

	log.Println("UserAuthMiddleware: User found and stored in Locals")
	return c.Next() // Continue to the next handler
}

// Route Handler: Use data from Locals
func GetUserProfile(c *fiber.Ctx) error {
	// Get data from Locals
	reqID := c.Locals("requestID").(string) // Requires type assertion
	user, ok := c.Locals("currentUser").(User) // Use type assertion with 'ok' check

	log.Printf("GetUserProfile: Request ID = %s", reqID)

	if !ok {
		log.Println("GetUserProfile: User data not found in Locals!")
		// This shouldn't happen if UserAuthMiddleware always runs before
		return c.Status(fiber.StatusInternalServerError).SendString("Internal error: user data missing")
	}

	// Use user data
	log.Printf("GetUserProfile: Fetching profile for user ID %d (%s)", user.ID, user.Role)
	return c.JSON(fiber.Map{
		"message":   "User Profile",
		"user_id":   user.ID,
		"user_role": user.Role,
		"request_id": reqID,
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Apply auth middleware before the profile handler
	app.Get("/profile", UserAuthMiddleware, GetUserProfile)

	log.Fatal(app.Listen(":3000"))
}
```

**Important:**
*   Data in `c.Locals()` only exists for the lifetime of a single request. The next request will have empty `Locals`.
*   When retrieving data from `Locals()`, you need to perform a *type assertion* (e.g., `.(string)`, `.(User)`) because `Locals()` stores values as `interface{}`. Always check the second boolean `ok` return value from the type assertion to handle cases where the key doesn't exist or the type is wrong.

#### Binding Request Data

As mentioned earlier, `c.BodyParser()`, `c.QueryParser()`, and `c.ParamsParser()` (for route params) are convenient ways to automatically parse incoming request data and populate Go structs. This significantly reduces boilerplate code for manual data extraction and conversion.

```go
type ProductFilter struct {
	Category string `query:"category"` // From query string ?category=...
	MaxPrice int    `query:"max_price"`// From query string ?max_price=...
	SortBy   string `query:"sort"`     // From query string ?sort=...
}

type UpdateProductInput struct {
	Name        string  `json:"name" form:"name"` // From JSON body or Form data
	Description *string `json:"description" form:"description"` // Pointer for optional value
	Price       float64 `json:"price" form:"price" validate:"required,gt=0"` // validation
	IsActive    bool    `json:"is_active" form:"is_active"`
}

type ProductRouteParams struct {
	ProductID int `params:"id"` // From route parameter /products/:id
}

func SearchProducts(c *fiber.Ctx) error {
	var filter ProductFilter
	// Bind query params to filter struct
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid filters"})
	}
	// Use filter.Category, filter.MaxPrice, filter.SortBy
	// ... search logic ...
	return c.JSON(fiber.Map{"message": "Search results", "filters": filter})
}

func UpdateProduct(c *fiber.Ctx) error {
	var params ProductRouteParams
	// Bind route params to params struct
	if err := c.ParamsParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Product ID"})
	}

	var input UpdateProductInput
	// Bind JSON/Form body to input struct
	if err := c.BodyParser(&input); err != nil {
		// Check if error is due to empty body (if allowed)
		if err == fiber.ErrUnprocessableEntity {
			 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body empty or malformed"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data: " + err.Error()})
	}

	// Here you can add additional validation (see Validation section)
	// validate := validator.New()
	// if err := validate.Struct(input); err != nil { ... }

	// Use params.ProductID and input.Name, input.Description, etc.
	// ... product update logic ...
	log.Printf("Updating product ID %d with data: %+v", params.ProductID, input)
	return c.JSON(fiber.Map{"message": "Product updated successfully", "id": params.ProductID})
}

func main() {
	app := fiber.New()
	app.Get("/products/search", SearchProducts) // e.g., /products/search?category=books&max_price=50
	app.Put("/products/:id", UpdateProduct)     // e.g., PUT /products/123 with JSON body
	log.Fatal(app.Listen(":3000"))
}

```
Using struct tags (`query:`, `json:`, `form:`, `params:`, `xml:`, `header:`) tells the parser how to map field names in the request to your Go struct fields.

### Request Handling 📥

This section summarizes specific ways to get various types of data from client requests using `fiber.Ctx`.

#### Reading Headers

Use `c.Get("Header-Name")`. Header names are case-insensitive.

```go
func HandleRequest(c *fiber.Ctx) error {
	userAgent := c.Get(fiber.HeaderUserAgent) // Fiber constant for common header name
	apiKey := c.Get("X-API-Key")
	acceptHeader := c.Get("Accept")

	log.Printf("User-Agent: %s", userAgent)
	log.Printf("API Key: %s", apiKey)
	log.Printf("Accept: %s", acceptHeader)

	// Check if the request accepts JSON
	if c.Accepts("application/json") != "" || c.Accepts("json") != "" {
		return c.JSON(fiber.Map{"message": "You requested JSON"})
	}

	return c.SendString("Headers received")
}
```

#### Reading Query Parameters

Query parameters are the part of the URL after the question mark (`?`), e.g., `/search?q=term&page=1`. Use `c.Query("key")` or `c.QueryParser(&struct)`.

```go
func SearchHandler(c *fiber.Ctx) error {
	// Manual way
	searchTerm := c.Query("q")
	page := c.Query("page", "1") // With default value "1"
	limit, err := c.QueryInt("limit", 10) // Parse to int, default 10
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Parameter 'limit' must be a number")
	}

	log.Printf("Searching for '%s', Page: %s, Limit: %d", searchTerm, page, limit)

	// Using struct binding
	type SearchParams struct {
		Query    string `query:"q"`
		Page     int    `query:"page" default:"1"`
		Limit    int    `query:"limit" default:"10"`
		Sort     string `query:"sort"`
		ShowMeta bool   `query:"show_meta"`
	}
	var params SearchParams
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
	}
	log.Printf("Struct Binding - Searching for '%s', Page: %d, Limit: %d, Sort: '%s', Meta: %t",
		params.Query, params.Page, params.Limit, params.Sort, params.ShowMeta)

	// ... search logic ...
	return c.JSON(fiber.Map{"results": "...", "params_used": params})
}
```
*Note:* The `default:"value"` tag can be used in the struct for `QueryParser` to provide default values if the query parameter is missing from the URL.

#### Reading Route Parameters

Parameters defined in the route path (e.g., `/users/:id`). Use `c.Params("key")` or `c.ParamsParser(&struct)`.

```go
type UserRouteParams struct {
	UserID int `params:"userId"` // Field name must match :userId in route definition
}

// Route: /users/:userId/orders/:orderId
func GetUserOrder(c *fiber.Ctx) error {
	// Manual way
	userIdStr := c.Params("userId")
	orderIdStr := c.Params("orderId")
	log.Printf("Manual - User ID: %s, Order ID: %s", userIdStr, orderIdStr)

	// Using struct binding (only for UserID)
	var params UserRouteParams
	if err := c.ParamsParser(&params); err != nil {
		 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
	}
	log.Printf("Struct Binding - User ID: %d", params.UserID)
	// You still need to get orderId manually if not bound
	orderId, err := c.ParamsInt("orderId")
	 if err != nil {
		 return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Order ID"})
	}
	log.Printf("Struct Binding - Order ID: %d", orderId)


	// ... logic to fetch order data ...
	return c.JSON(fiber.Map{"user_id": params.UserID, "order_id": orderId})
}
```

#### Reading Request Body

For requests like POST, PUT, PATCH that carry data in the body. The most common way is using `c.BodyParser(&struct)`.

```go
type CreatePostInput struct {
	Title   string   `json:"title" form:"title" validate:"required"`
	Content string   `json:"content" form:"content" validate:"required"`
	Tags    []string `json:"tags" form:"tags"` // Can be array/slice
}

func CreatePostHandler(c *fiber.Ctx) error {
	var input CreatePostInput

	// BodyParser handles JSON, Form (urlencoded/multipart), XML
	if err := c.BodyParser(&input); err != nil {
		log.Printf("Error parsing body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to process request body",
			"details": err.Error(),
		})
	}

	// (Optional but recommended) Perform validation
	// validate := validator.New()
	// if err := validate.Struct(input); err != nil { ... return validation errors ...}

	log.Printf("Creating new post: Title='%s', Content='%s', Tags=%v", input.Title, input.Content, input.Tags)
	// ... logic to save post to database ...

	// Return the newly created post data (e.g., with ID)
	newPost := fiber.Map{
		"id": 1, // Example ID from DB
		"title": input.Title,
		"content": input.Content,
		"tags": input.Tags,
	}
	return c.Status(fiber.StatusCreated).JSON(newPost) // 201 Created
}
```

If you need the raw body (e.g., for processing webhook signatures):

```go
func WebhookHandler(c *fiber.Ctx) error {
	rawBody := c.Body() // Get []byte

	// Do something with rawBody, e.g., verify signature
	signature := c.Get("X-Webhook-Signature")
	if !verifySignature(rawBody, signature) {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid signature")
	}

	// If signature is valid, you might still want to parse the body
	var payload map[string]interface{}
	// Use fiber.Unmarshal if body has already been read
	// if err := fiber.Unmarshal(rawBody, &payload); err != nil {
	//     return c.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON payload")
	// }
	// Or if BodyParser fails after c.Body() was called (might happen)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse JSON payload after reading raw body")
	}

	log.Printf("Webhook received: %+v", payload)
	// ... process webhook event ...

	return c.SendStatus(fiber.StatusOK) // Send 200 OK
}
```
**Caution:** Calling `c.Body()` reads the entire body into memory. If you subsequently call `c.BodyParser()`, the parser might not be able to re-read the body (depending on Fiber/Fasthttp internal implementation). If you need both the raw body *and* parsing, read the raw body first, then use an appropriate unmarshal function (e.g., `fiber.Unmarshal` for JSON) on the raw bytes.

#### File Upload

Files are typically uploaded using `multipart/form-data`. Fiber makes handling this easy.

```go
import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)


func UploadFileHandler(c *fiber.Ctx) error {
	// 1. Get the file header from the form field named "file_upload"
	fileHeader, err := c.FormFile("file_upload")
	if err != nil {
		log.Printf("Error getting file: %v", err)
		// Check if error is due to field not found
		if err.Error() == "multipart: no such file" || err == http.ErrMissingFile {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Field 'file_upload' not found or empty"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to process file upload"})
	}

	// 2. (Optional) Get other form fields if present
	description := c.FormValue("description", "No description")

	// 3. (Optional) Validate file (size, MIME type)
	maxSize := int64(5 * 1024 * 1024) // 5 MB
	if fileHeader.Size > maxSize {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{"error": "File size exceeds 5MB limit"})
	}

	allowedMIMETypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"application/pdf": true,
	}
	// Detect MIME type from file header (safer than extension)
	file, err := fileHeader.Open()
	if err != nil {
		 return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer file.Close() // Ensure file is closed
	// Read a small portion for MIME type detection
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		 return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read file for MIME detection"})
	}
	// Reset the read pointer in case Read consumed necessary bytes
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to reset file pointer"})
	}
	// Detect the content type using the buffer
	mimeType := http.DetectContentType(buffer)


	if !allowedMIMETypes[mimeType] {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(fiber.Map{
			"error": "Unsupported file type",
			"detected_mime": mimeType,
		})
	}


	// 4. Determine save path (make unique if necessary)
	//    IMPORTANT: Never use fileHeader.Filename directly as path
	//    as it can contain malicious characters (e.g., "../"). Always sanitize!
	safeFilename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filepath.Base(fileHeader.Filename)) // Example unique name
	savePath := filepath.Join("./uploads", safeFilename) // Save in ./uploads directory

	// Ensure uploads directory exists
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		log.Printf("Error creating uploads directory: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to prepare file storage"})
	}

	// 5. Save the file to disk using c.SaveFile or io.Copy
	// c.SaveFile is easier:
	// err = c.SaveFile(fileHeader, savePath)
	// If using SaveFile, no need to manually Open/Seek/Copy

	// Or using io.Copy (gives more control, uses the already opened file)
	dst, err := os.Create(savePath)
	if err != nil {
		log.Printf("Error creating destination file %s: %v", savePath, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}
	defer dst.Close()

	_, err = io.Copy(dst, file) // Copy from the uploaded file stream to the destination file
	if err != nil {
		log.Printf("Error saving file %s: %v", savePath, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
	}

	log.Printf("File '%s' (description: '%s', size: %d, mime: %s) uploaded successfully to %s",
		fileHeader.Filename, description, fileHeader.Size, mimeType, savePath)

	// Send success response
	return c.JSON(fiber.Map{
		"message":       "File uploaded successfully!",
		"original_name": fileHeader.Filename,
		"saved_path":    savePath,
		"size":          fileHeader.Size,
		"mime_type":     mimeType,
		"description":   description,
	})
}

```
**Key Points for File Upload:**
*   Use `c.FormFile()` to get the `*multipart.FileHeader`.
*   Use `c.FormValue()` to get other text fields in the same form.
*   **Always validate** file size and MIME type. Don't trust the file extension from the client. Use `http.DetectContentType`.
*   **Sanitize the filename** before saving to disk to prevent *path traversal attacks*. `filepath.Base()` helps get only the filename. Generate unique filenames (e.g., with timestamp or UUID) to avoid conflicts.
*   Use `c.SaveFile()` for an easy way to save, or `fileHeader.Open()` and `io.Copy` for more control.
*   Ensure the destination directory exists (`os.MkdirAll`).
*   Set a request body size limit in Fiber config (`BodyLimit`) to prevent DoS attacks with large files.

### Response Handling 📤

This section summarizes how to send various types of responses back to the client.

#### Setting Status Code

Use `c.SendStatus(code)` to send only the status, or `c.Status(code).<SendMethod>(...)` to set the status before sending the body.

```go
app.Post("/items", func(c *fiber.Ctx) error {
	// ... logic to create item ...
	newItem := Item{ID: 5, Name: "New Item"}
	return c.Status(fiber.StatusCreated).JSON(newItem) // 201 Created
})

app.Get("/items/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := findItemByID(id) // Imaginary function
	if err != nil {
		// If item not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"}) // 404 Not Found
	}
	return c.JSON(item) // 200 OK (default if not set)
})

app.Delete("/items/:id", func(c *fiber.Ctx) error {
	// ... logic to delete item ...
	return c.SendStatus(fiber.StatusNoContent) // 204 No Content (common for successful DELETE)
})
```

#### Setting Headers

Use `c.Set(key, value)` or `c.Append(key, value)`.

```go
app.Get("/data", func(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8) // text/plain; charset=utf-8
	c.Set("X-RateLimit-Limit", "100")
	c.Set("X-RateLimit-Remaining", "99")
	// Set Cache-Control header
	c.Set(fiber.HeaderCacheControl, "public, max-age=3600") // Cache for 1 hour
	return c.SendString("This is plain text data with custom headers.")
})

app.Post("/login", func(c *fiber.Ctx) error {
	// ... login validation ...
	c.Cookie(&fiber.Cookie{ // Setting a session cookie
		Name:     "session",
		Value:    "secret-user-123",
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return c.JSON(fiber.Map{"message": "Login successful"})
})
```

#### Sending Various Data Types

*   **String:** `c.SendString("...")` -> `text/plain`
*   **Byte Slice:** `c.Send([]byte{...})` -> `application/octet-stream` (or detected)
*   **JSON:** `c.JSON(data)` -> `application/json`
*   **XML:** `c.XML(data)` -> `application/xml`
*   **HTML (from Template):** `c.Render("template.html", data)` -> `text/html` (requires template engine setup)
*   **File (inline):** `c.SendFile("./path/to/file")` -> Detected MIME type
*   **File (download):** `c.Download("./path/to/file", "download_name.ext")` -> Detected MIME type + `Content-Disposition: attachment`
*   **Redirect:** `c.Redirect("/new/location", 302)`

```go
// Example sending a structured JSON error response
func GetResource(c *fiber.Ctx) error {
	if !userHasPermission(c) { // Imaginary function
		// 403 Forbidden
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"code":    "ACCESS_DENIED",
			"message": "You do not have permission to access this resource.",
		})
	}
	// ...
	return c.JSON(fiber.Map{"data": "..."})
}
```

### Error Handling 💣

Good error handling is crucial for robust applications. Fiber provides several mechanisms for this.

#### Returning Errors from Handlers

The most basic way is to return an `error` from your handler or middleware.

```go
import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm" // Example using GORM error
)

func GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	item, err := database.FindItem(id) // DB function can return error

	if err != nil {
		// If error is 'record not found'
		if errors.Is(err, gorm.ErrRecordNotFound) { // Example with GORM
			log.Printf("Item %s not found", id)
			// Return standard Fiber error for Not Found
			return fiber.ErrNotFound // Will result in 404 Not Found
		}
		// For other database errors
		log.Printf("Database error fetching item %s: %v", id, err)
		// Return a generic error
		return fiber.ErrInternalServerError // Will result in 500 Internal Server Error
		// or a custom error
		// return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve item data")
	}

	return c.JSON(item)
}
```
Fiber has several convenient pre-defined errors (like `fiber.ErrBadRequest`, `fiber.ErrNotFound`, `fiber.ErrUnauthorized`, etc.) that map directly to the appropriate HTTP status codes.

#### Custom Errors (`fiber.NewError`)

If the pre-defined errors aren't sufficient, you can create custom Fiber errors with specific status codes and messages using `fiber.NewError(statusCode, message)`.

```go
import (
	"errors"
	"log"
	"github.com/gofiber/fiber/v2"
	"your-payment-gateway-sdk" // Imaginary SDK
)

func ProcessPayment(c *fiber.Ctx) error {
	// ... input validation logic ...
	if !inputValid {
		// Create a 400 Bad Request error with a specific message
		return fiber.NewError(fiber.StatusBadRequest, "Payment data incomplete or invalid.")
	}

	err := paymentGateway.Charge(...) // Imaginary function
	if err != nil {
		// Handle specific errors from the payment gateway
		if errors.Is(err, paymentGateway.ErrInsufficientFunds) {
			// Create a 402 Payment Required error
			return fiber.NewError(fiber.StatusPaymentRequired, "Insufficient funds.")
		}
		// Other gateway errors
		log.Printf("Payment gateway error: %v", err)
		// Create a 503 Service Unavailable error
		return fiber.NewError(fiber.StatusServiceUnavailable, "Payment service is currently unavailable.")
	}

	return c.JSON(fiber.Map{"status": "Payment successful"})
}
```

#### Custom Error Handler

By default, when a handler returns an `error` (or `fiber.Error`), Fiber catches it and sends an appropriate HTTP response (using the status code from `fiber.Error` or defaulting to 500, and the message from the error).

You can **fully customize** how these errors are turned into HTTP responses by providing a custom `ErrorHandler` in the Fiber configuration. This is useful for:

*   Formatting all error responses consistently (e.g., always in a specific JSON format).
*   Hiding internal error details in production environments.
*   Centralized error logging.
*   Sending errors to monitoring systems (like Sentry, Datadog).

```go
package main

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Custom Error Handler Function
func MyCustomErrorHandler(c *fiber.Ctx, err error) error {
	// Default status code is 500
	code := fiber.StatusInternalServerError
	message := "An internal server error occurred."

	// Check if error is a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}

	// Log internal errors in detail (only on the server)
	log.Printf("[ErrorHandler] Status: %d, Error: %v, Path: %s", code, err, c.Path())

	// Send error to Sentry/Datadog etc. here if needed

	// Don't send internal error details to the client in production
	isProduction := os.Getenv("APP_ENV") == "production"
	if isProduction && code == fiber.StatusInternalServerError {
		message = "Sorry, an unexpected error occurred."
	}

	// Set Content-Type if not already set
	// (Important if error occurs before Content-Type is set by handler)
	if c.Get(fiber.HeaderContentType) == "" {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	}

	// Send consistent JSON error response
	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"code":    code, // Or your internal application error code
		"message": message,
		// "details": err.Error(), // AVOID this in production
	})
}

func main() {
	app := fiber.New(fiber.Config{
		// Register our custom error handler
		ErrorHandler: MyCustomErrorHandler,
	})

	// Important: Recover middleware should still be used to catch panics
	// before the custom ErrorHandler is called.
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// This handler succeeds
		return c.SendString("OK")
	})

	app.Get("/notfound", func(c *fiber.Ctx) error {
		// Return a standard Fiber error
		return fiber.ErrNotFound // Will be handled by MyCustomErrorHandler -> 404 JSON
	})

	app.Get("/badrequest", func(c *fiber.Ctx) error {
		// Return a custom error
		return fiber.NewError(fiber.StatusBadRequest, "Parameter 'q' is required.") // -> 400 JSON
	})

	app.Get("/dberror", func(c *fiber.Ctx) error {
		// Simulate a non-Fiber error
		simulatedError := errors.New("database connection failed")
		return simulatedError // Will be handled by MyCustomErrorHandler -> 500 JSON
	})

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Something went terribly wrong!") // Will be caught by Recover, then to ErrorHandler -> 500 JSON
	})


	log.Fatal(app.Listen(":3000"))
}
```

With a custom `ErrorHandler`, all errors returned from handlers (including panics caught by `recover`) will pass through this function, giving you full control over the final error response.

#### Recover Middleware

The `recover` middleware (from `github.com/gofiber/fiber/v2/middleware/recover`) is crucial. Its function is to catch *panics* that might occur in your handlers or middleware.

Without `recover`, a panic would cause the server to crash. With `recover`, the panic is caught, converted into an error (usually `500 Internal Server Error`), and then processed by the `ErrorHandler` (default or custom).

**Always use `recover.New()` as one of your first global middlewares.**

```go
app := fiber.New(fiber.Config{ErrorHandler: MyCustomErrorHandler})

// Recover should be registered BEFORE other middleware/handlers that might panic
app.Use(recover.New(recover.Config{
	EnableStackTrace: true, // Log stack trace (useful during development)
}))

// Other Middleware and Routes...
app.Get("/panic-now", func(c *fiber.Ctx) error {
	myMap := map[string]string{}
	// This will cause a panic (nil map dereference)
	myMap["key"] = "value" // <-- PANIC!
	return c.SendString("Won't reach here")
})
```

### Configuration (`fiber.Config`)

When creating a Fiber app instance with `fiber.New()`, you can pass a `fiber.Config` struct to customize various aspects of its behavior.

```go
import (
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
)

// MyCustomErrorHandler function defined elsewhere
// func MyCustomErrorHandler(c *fiber.Ctx, err error) error { ... }

config := fiber.Config{
	// Prefork spawns multiple Go processes listening on the same port.
	// Useful for leveraging all CPU cores without manual Goroutine logic.
	// Not compatible with some stateful middleware (e.g., default session).
	// Default: false
	Prefork: os.Getenv("APP_ENV") == "production", // Enable only in production

	// App name, appears in 'Server' header if ServerHeader isn't set.
	AppName: "My Awesome App v1.1",

	// Override the default 'Server' header ('Fiber'). Empty string hides it.
	ServerHeader: "MyWebServer",
	// ServerHeader: "", // Hide the Server header

	// Strict routing. If true, '/foo' and '/foo/' are treated differently.
	// Default: false
	StrictRouting: false,

	// Case-sensitive routing. If true, '/Foo' and '/foo' are different.
	// Default: false
	CaseSensitive: false,

	// Maximum request body size in bytes. Prevents DoS.
	// Default: 4 * 1024 * 1024 (4MB)
	BodyLimit: 10 * 1024 * 1024, // 10 MB

	// Configuration for the template engine (see Template Engine section).
	// Views: ...,

	// Custom error handler (see Error Handling section).
	ErrorHandler: MyCustomErrorHandler,

	// Maximum number of allowed request headers.
	// Default: 1024
	ReadBufferSize: 8192, // Increase if handling large headers

	// Timeout for reading the entire request (including body).
	// Default: No limit (relies on OS timeout)
	ReadTimeout: 5 * time.Second,

	// Timeout for writing the response.
	// Default: No limit
	WriteTimeout: 10 * time.Second,

	// Timeout for idle (keep-alive) connections.
	// Default: No limit (relies on OS timeout)
	IdleTimeout: 60 * time.Second,

	// Other configurations... (see fiber.Config documentation)
	// DisableKeepalive: false,
	// ReduceMemoryUsage: false, // Can reduce memory but might be slightly slower
	// GETOnly: false, // Only allow GET methods
	// EnablePrintRoutes: true, // Print routes on startup
	// Network: "tcp", // Network protocol (tcp, tcp4, tcp6)
}

app := fiber.New(config)
```

Choose the configuration options that suit your application's needs, especially `Prefork`, `BodyLimit`, `ErrorHandler`, and `Timeouts` for production applications.

#### Prefork Configuration

`Prefork` mode is a unique Fiber feature (leveraging the SO_REUSEPORT feature on Linux/BSD). When `Prefork: true`, Fiber will:

1.  Create as many *child processes* as there are available CPU cores (or based on `runtime.GOMAXPROCS(0)`).
2.  Each child process runs an identical instance of the Fiber application.
3.  All child processes listen on the *same port*.
4.  The operating system kernel distributes incoming connections to one of the available child processes (kernel-level load balancing).

**Advantages:**

*   **Automatic Multi-core Utilization:** An easy way to make your application use all CPU cores without explicitly managing Goroutines for handling requests in parallel.
*   **Potentially Higher Throughput:** Can increase the number of requests per second that can be handled.
*   **Isolation:** If one child process crashes due to a panic (though `recover` should catch it), other child processes keep running and serving requests.

**Disadvantages/Considerations:**

*   **Stateful Middleware:** Middleware that stores state in the process memory (like Fiber's default memory-based `session` middleware) **will not work correctly**. Each child process has its own memory, so a session created in one process won't be recognized by another. You need to use external state storage (like Redis, Memcached, database) for sessions, rate limiting, etc., if using prefork.
*   **Debugging:** Slightly more complex to debug as you have multiple processes.
*   **Overhead:** Starting multiple processes consumes more memory than a single process with many Goroutines.
*   **Linux/BSD Only:** The SO_REUSEPORT feature isn't widely available or works the same way on Windows or macOS.

**When to Use Prefork?**

*   *Stateless* applications (not relying on in-memory state between requests).
*   CPU-bound applications that need to utilize all cores.
*   Running in a Linux/BSD environment in production.

**When to Avoid Prefork?**

*   Using memory-based stateful middleware.
*   Needing global state shared among all handlers (requires inter-process synchronization if using prefork, easier with Goroutines in a single process).
*   Running on Windows/macOS.
*   I/O-bound applications (waiting on database, network) might not benefit much compared to regular Goroutines.

When in doubt, start without `Prefork` (`false`) and enable it later if benchmarking shows significant benefits and you are prepared to handle the state implications.

### Template Engine 📄

Fiber allows you to render dynamic HTML using various Go template engines. Fiber itself doesn't include an engine but provides an *interface* (`fiber.Views`) that engine adapters can implement.

#### Template Engine Concepts

A template engine separates presentation logic (HTML) from business logic (Go). You create template files (e.g., `.html`, `.tmpl`, `.pug`) containing HTML markup mixed with the engine's specific syntax to display dynamic data, loop through items, perform conditionals, etc.

#### Using Built-in Template Engine (HTML)

Go has a built-in `html/template` package that is HTML-safe (performs automatic escaping to prevent XSS). Fiber provides an adapter for this.

1.  **Install Adapter:**
	```bash
	go get -u github.com/gofiber/template/html/v2
	```

2.  **Configure in Fiber:**
	Create a directory to store your template files (e.g., `views`). Create a simple template file, e.g., `views/index.html`:
	```html
	<!-- views/index.html -->
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{.Title}}</title> <!-- Display 'Title' data -->
	</head>
	<body>
		<h1>{{.Header}}</h1>
		<p>Welcome to the example page!</p>

		<h2>Item List:</h2>
		{{if .Items}} <!-- Check if Items exists and is not empty -->
			<ul>
				{{range .Items}} <!-- Loop through the Items slice -->
					<li>{{.}}</li> <!-- Display each item -->
				{{end}}
			</ul>
		{{else}}
			<p>No items to display.</p>
		{{end}}
	</body>
	</html>
	```

	Configure the engine in `main.go`:
	```go
	package main

	import (
		"log"

		"github.com/gofiber/fiber/v2"
		"github.com/gofiber/template/html/v2" // Import adapter
	)

	func main() {
		// 1. Create the engine instance, point it to the views directory
		//    Using Reload: true is recommended during development
		//    so template changes are reflected without restarting the server.
		engine := html.New("./views", ".html") // Look for .html files in ./views
		engine.Reload(true) // Enable reload during development
		engine.Debug(true) // Enable debug logging during development

		// 2. Create Fiber app with the configured engine
		app := fiber.New(fiber.Config{
			Views: engine, // Tell Fiber to use this engine
		})

		// 3. Define a route that uses c.Render()
		app.Get("/", func(c *fiber.Ctx) error {
			// Data to pass to the template
			data := fiber.Map{
				"Title":  "Main Page",
				"Header": "Hello from Fiber!",
				"Items":  []string{"Apple", "Banana", "Cherry"},
			}
			// Render the "index" template (without the .html extension)
			// and pass the data to it.
			return c.Render("index", data)
		})

		app.Get("/empty", func(c *fiber.Ctx) error {
			// Example with no items
			data := fiber.Map{
				"Title":  "Empty Page",
				"Header": "No items here",
				"Items":  nil, // Or []string{}
			}
			return c.Render("index", data)
		})

		log.Fatal(app.Listen(":3000"))
	}
	```

3.  **Run and Test:** Run `go run main.go` and open `http://localhost:3000` in your browser. You'll see the rendered HTML with data from the handler. Access `http://localhost:3000/empty` to see the `else` condition.

#### Using Other Template Engines

Fiber supports many other popular engines via separate adapters (usually in `github.com/gofiber/template/...`). Examples: Pug, Amber, Handlebars, Jet.

For instance, using **Handlebars**:

1.  **Install Adapter:**
	```bash
	go get -u github.com/gofiber/template/handlebars/v2
	```
2.  **Create Handlebars Template:** (e.g., `views/profile.hbs`)
	```handlebars
	<!-- views/profile.hbs -->
	<!DOCTYPE html>
	<html>
	<head>
		<title>Profile {{user.Name}}</title>
	</head>
	<body>
		<h1>User Profile</h1>
		<p>ID: {{user.ID}}</p>
		<p>Name: {{user.Name}}</p>
		<p>Email: {{user.Email}}</p>

		{{#if isAdmin}}
			<p><strong>Status: Administrator</strong></p>
		{{else}}
			<p>Status: Regular User</p>
		{{/if}}
	</body>
	</html>
	```
3.  **Configure in Fiber:**
	```go
	// ... other imports ...
	import (
		"github.com/gofiber/fiber/v2"
		"github.com/gofiber/template/handlebars/v2" // Import handlebars adapter
		"log"
	)

	func main() {
		// Create Handlebars engine
		engine := handlebars.New("./views", ".hbs") // .hbs extension
		engine.Reload(true)

		app := fiber.New(fiber.Config{
			Views: engine,
		})

		app.Get("/profile/:id", func(c *fiber.Ctx) error {
			// Example data
			userData := fiber.Map{
				"ID": c.Params("id"),
				"Name": "John Doe",
				"Email": "john.doe@example.com",
			}
			isAdmin := c.QueryBool("admin", false) // Check for ?admin=true query

			// Pass data to the profile.hbs template
			return c.Render("profile", fiber.Map{
				"user": userData,
				"isAdmin": isAdmin,
			})
		})

		log.Fatal(app.Listen(":3000"))
	}
	```

The process is similar for other engines: install the adapter, create templates with that engine's syntax, and configure the engine in `fiber.Config`.

#### Layouts

Many template engines (including `html/template` with `define`/`template` functions and the Handlebars adapter) support the concept of *layouts*. A layout is a base template (HTML skeleton) that defines the common page structure (header, footer, sidebar), and page-specific content is inserted into it.

Using layouts reduces HTML code duplication across many templates.

Example with `html/template`:

1.  **Create Layout (`views/layouts/main.html`):**
	```html
	<!-- views/layouts/main.html -->
	<!DOCTYPE html>
	<html>
	<head>
		<title>{{template "title" .}} - My Website</title> <!-- Call the title block -->
		<link rel="stylesheet" href="/static/css/style.css">
	</head>
	<body>
		<header>Common Header</header>
		<main>
			{{template "content" .}} <!-- Call the content block -->
		</main>
		<footer>Common Footer</footer>
	</body>
	</html>
	```

2.  **Create Content Template (`views/about.html`):**
	```html
	<!-- views/about.html -->
	{{define "title"}}About Us{{end}} <!-- Define the title block -->

	{{define "content"}} <!-- Define the content block -->
		<h2>About Our Company</h2>
		<p>This is the about page.</p>
	{{end}}
	```

3.  **Configure Engine and Render with Layout:**
	```go
	// ... imports ...
	import (
		"log"
		"github.com/gofiber/fiber/v2"
		"github.com/gofiber/template/html/v2"
	)


	func main() {
		// Load all .html templates from the views directory
		engine := html.New("./views", ".html")
		engine.Reload(true)

		app := fiber.New(fiber.Config{ Views: engine })

		app.Get("/about", func(c *fiber.Ctx) error {
			// When rendering 'about', also pass the layout file name 'main'
			// Data will be available in both templates (layout and content)
			return c.Render("about", fiber.Map{}, "layouts/main")
		})

		// Other routes can use the same layout
		app.Get("/contact", func(c *fiber.Ctx) error {
			// Create views/contact.html similar to about.html
			return c.Render("contact", fiber.Map{"Email": "info@example.com"}, "layouts/main")
		})

		log.Fatal(app.Listen(":3000"))
	}
	```

When `c.Render("about", data, "layouts/main")` is called:
*   Fiber (via the engine) loads `views/about.html` and `views/layouts/main.html`.
*   The engine executes `layouts/main.html` as the base template.
*   When `{{template "content" .}}` is encountered in the layout, the engine looks for the `{{define "content"}}` block in `about.html` and renders it there. The same applies to `{{template "title" .}}`.
*   The data (`fiber.Map{}`) is passed to both templates.

Other template engine adapters might handle layouts differently (e.g., Handlebars uses Partials and Helpers). Consult the specific adapter's documentation.

### Serving Static Files 📁

Web applications almost always need to serve static files like CSS, JavaScript, images, fonts, etc. Fiber provides the efficient `Static` middleware for this.

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Create a 'public' directory in your project root
	// Put CSS, JS, image files inside it
	// Example: ./public/css/style.css
	//         ./public/js/script.js
	//         ./public/images/logo.png

	// Register the Static middleware
	// First argument is the URL prefix
	// Second argument is the directory path on the filesystem
	app.Static("/static", "./public")
	// Now:
	// Request to /static/css/style.css will serve the file ./public/css/style.css
	// Request to /static/images/logo.png will serve the file ./public/images/logo.png

	// You can register multiple static directories
	app.Static("/assets", "./assets") // Serve files from ./assets under the /assets URL

	// Serving files from the root URL (e.g., favicon.ico or index.html)
	// Use the "/" prefix
	app.Static("/", "./root_files")
	// Request to /favicon.ico will serve ./root_files/favicon.ico

	// Your other application routes
	app.Get("/", func(c *fiber.Ctx) error {
		// Example HTML referencing static files
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Fiber App</title>
			<link rel="stylesheet" href="/static/css/style.css">
			<link rel="icon" href="/favicon.ico">
		</head>
		<body>
			<h1>Welcome!</h1>
			<img src="/static/images/logo.png" alt="Logo">
			<script src="/static/js/script.js"></script>
		</body>
		</html>
		`
		// Send as HTML
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
		return c.SendString(html)
	})

	log.Fatal(app.Listen(":3000"))
}
```

**`Static` Configuration:**

The `Static` middleware can be further configured:

```go
app.Static("/downloadables", "./files_to_download", fiber.Static{
	// Compress static files (gzip, brotli) if client supports it.
	// Default: false (in v2. Was true in older versions)
	Compress: true,

	// Allow byte range requests (important for video/audio streaming).
	// Default: false (in v2)
	ByteRange: true,

	// Allow directory browsing (shows file list if index.html is missing).
	// Default: false (Don't enable in production unless really needed!)
	Browse: false,

	// Default index file name looked for when accessing a directory.
	// Default: "index.html"
	Index: "default.html",

	// Cache-Control max-age duration (seconds). 0 means no cache.
	// Default: 0
	MaxAge: 3600, // Cache for 1 hour
})
```

**Important:**
*   Place the `Static` middleware **before** your route definitions if there's a potential path conflict (e.g., if you have a route `/static/users` and also a static directory `/static`). Fiber will try to match static files first.
*   Ensure the directory path you provide to `app.Static` is correct relative to the location where you run your compiled Go application binary.

### Request Validation ✅

Validating incoming data from the client (body, query params, route params) is a crucial step for application security and data integrity. Never trust client input!

#### Importance of Validation

*   **Security:** Prevents attacks like SQL Injection, Cross-Site Scripting (XSS), and other exploits that leverage invalid data.
*   **Data Integrity:** Ensures data stored in your database or processed by your application has the correct format and values.
*   **User Experience:** Provides clear feedback to users if their input is incorrect.
*   **Application Stability:** Prevents errors or panics caused by unexpected data.

#### Using a Validator Library

Fiber doesn't include a built-in validation library, but it integrates easily with popular libraries like [`go-playground/validator`](https://github.com/go-playground/validator). This is a very powerful and widely used library in the Go ecosystem.

1.  **Install Validator:**
	```bash
	go get github.com/go-playground/validator/v10
	```

2.  **Add Validation Tags to Structs:**
	Use the `validate` tag on struct fields you want to validate. The library supports many built-in validation rules (required, email, url, min, max, len, uuid, etc.) and allows custom validations.

	```go
	type RegisterUserInput struct {
		Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Age      int    `json:"age" validate:"omitempty,gte=18,lte=120"` // gte=greater than or equal, lte=less than or equal
		Website  string `json:"website" validate:"omitempty,url"` // omitempty: validate only if field is not empty
		UserType string `json:"user_type" validate:"required,oneof=admin user guest"` // Must be one of these values
	}

	type ProductFilter struct {
		Category string `query:"category" validate:"omitempty,alpha"` // Only letters
		MaxPrice *int   `query:"max_price" validate:"omitempty,gt=0"` // Pointer allows nil/empty, validate if present
		Page     int    `query:"page" validate:"omitempty,min=1"`
	}
	```
	See the [`go-playground/validator` documentation](https://pkg.go.dev/github.com/go-playground/validator/v10) for a full list of available tags.

#### Implementation Example

You need to create a validator instance and call its `Struct()` method after binding the request data.

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10" // Import validator
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Global validator instance (or create per request if different configs needed)
var validate = validator.New()

// Struct for a more informative validation error response
type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

// Helper function to format validation errors
func formatValidationErrors(err error) []ValidationErrorResponse {
	var errors []ValidationErrorResponse

	// Check if the error is of type ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			// Create more user-friendly error messages (simple example)
			var message string
			switch fieldErr.Tag() {
			case "required":
				message = "This field is required."
			case "email":
				message = "Invalid email format."
			case "min":
				 message = "Minimum value/length is " + fieldErr.Param() + "."
			case "max":
				 message = "Maximum value/length is " + fieldErr.Param() + "."
			 case "alphanum":
				 message = "Must contain only letters and numbers."
			case "oneof":
				message = "Must be one of: " + strings.Replace(fieldErr.Param(), " ", ", ", -1)
			default:
				message = "Invalid field (" + fieldErr.Tag() + ")"
			}

			errors = append(errors, ValidationErrorResponse{
				Field:   fieldErr.Field(), // Struct field name
				Tag:     fieldErr.Tag(),   // Validation tag that failed
				Message: message,          // Custom message
			})
		}
	} else {
		// If error is not ValidationErrors (rare if input is err from validate.Struct)
		log.Printf("Warning: Unexpected validation error: %v", err)
		errors = append(errors, ValidationErrorResponse{Message: "Unknown validation error"})
	}
	return errors
}

// Handler for user registration
func RegisterUserHandler(c *fiber.Ctx) error {
	input := new(RegisterUserInput) // Use new() to get a pointer

	// 1. Bind request body to struct
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Failed to process request body.",
			"details": err.Error(),
		})
	}

	// 2. Perform validation on the bound struct
	err := validate.Struct(input)
	if err != nil {
		// If validation fails, format errors and send 400 Bad Request
		validationErrors := formatValidationErrors(err)
		log.Printf("Validation failed for registration: %v", validationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Provided data is invalid.",
			"errors":  validationErrors, // Send details of validation errors
		})
	}

	// 3. If validation succeeds, proceed
	log.Printf("Valid registration received: Username=%s, Email=%s, Age=%d", input.Username, input.Email, input.Age)
	// ... logic to save user to database ...

	// Send success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User registered successfully.",
		"user": fiber.Map{ // Don't send password back!
			"username": input.Username,
			"email":    input.Email,
			"age": input.Age,
			"website": input.Website,
			"user_type": input.UserType,
		},
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/register", RegisterUserHandler)

	log.Fatal(app.Listen(":3000"))
}
```

**How to Test:**
*   **Valid Request:**
	```bash
	curl -X POST http://localhost:3000/register -H "Content-Type: application/json" -d '{"username": "johndoe", "email": "john.doe@example.com", "password": "password123", "age": 25, "user_type": "user"}'
	```
	*(Should return 201 Created)*
*   **Invalid Request (multiple errors):**
	```bash
	curl -X POST http://localhost:3000/register -H "Content-Type: application/json" -d '{"username": "jo", "email": "john.doe@", "password": "pass", "age": 15, "user_type": "superadmin"}'
	```
	*(Should return 400 Bad Request with a list of validation errors in the JSON body)*

**Validation Tips:**

*   **Validate Early:** Perform validation as soon as you receive the data.
*   **Don't Rely Solely on Frontend Validation:** Browser validation (JavaScript) is good for UX but *must* be re-validated on the backend because clients can bypass frontend validation.
*   **Provide Clear Error Messages:** Help users correct their input. Format validation error responses consistently.
*   **Use Pointers for Optional Fields:** If a field (like `Age` or `Website`) is optional *and* has validation rules (like `gte=18` or `url`), use a pointer (`*int`, `*string`). This way, if the client doesn't send the field, its value will be `nil`, and the `omitempty` tag will skip validation. If using non-pointer types (`int`, `string`), Go's zero value (0 or "") would be validated, which might not be desired.
*   **Custom Validation:** `go-playground/validator` allows you to register custom validation functions for more complex logic (e.g., checking if a username already exists in the database).

---

## 6. Advanced Topics 🧭

After mastering the basics, let's explore some more advanced features and concepts in Fiber and Go web development.

### WebSocket

WebSockets provide a full-duplex communication channel over a single TCP connection. This is very useful for real-time applications like chats, live notifications, online games, etc.

Fiber provides an easy-to-use WebSocket middleware package: `github.com/gofiber/contrib/websocket`.

1.  **Install:**
	```bash
	go get github.com/gofiber/contrib/websocket
	```

2.  **Example Usage:**
	```go
	package main

	import (
		"log"

		"github.com/gofiber/contrib/websocket" // Import websocket
		"github.com/gofiber/fiber/v2"
		"github.com/gofiber/fiber/v2/middleware/logger"
	)

	func main() {
		app := fiber.New()
		app.Use(logger.New())

		// Middleware to ensure the request is a WebSocket upgrade
		app.Use("/ws", func(c *fiber.Ctx) error {
			// Check if headers indicate a WebSocket upgrade request
			if websocket.IsWebSocketUpgrade(c) {
				c.Locals("allowed", true)
				return c.Next()
			}
			// If not a WebSocket request, send 426 Upgrade Required
			return fiber.ErrUpgradeRequired
		})

		// Handler for WebSocket connections on the path /ws/:id
		app.Get("/ws/:id", websocket.New(func(conn *websocket.Conn) {
			// conn is a *websocket.Conn wrapping the connection
			// Get parameters from the original URL (before upgrade)
			id := conn.Params("id")
			log.Printf("WebSocket connected for ID: %s from %s", id, conn.RemoteAddr())

			// Variables for message type, message, and error
			var (
				mt  int
				msg []byte
				err error
			)

			// Infinite loop to read messages from the client
			for {
				// Read message from client
				// conn.ReadMessage() is a blocking call
				if mt, msg, err = conn.ReadMessage(); err != nil {
					// If there's an error (connection closed, etc.), log and break the loop
					log.Println("Error reading message:", err)
					break // Exit loop, closes connection on server side
				}

				log.Printf("Message received from ID %s: %s (Type: %d)", id, msg, mt)

				// Send message back to client (echo)
				// You can send text (websocket.TextMessage) or binary (websocket.BinaryMessage)
				if err = conn.WriteMessage(mt, msg); err != nil {
					log.Println("Error writing message:", err)
					break // Exit if writing fails
				}
			}
			// Code after the loop executes when the connection is closed (by client or server)
			log.Printf("WebSocket disconnected for ID: %s", id)
			// Perform cleanup if necessary (e.g., remove user from online list)
		}))

		log.Fatal(app.Listen(":3000"))
	}
	```

**Explanation:**
*   `websocket.IsWebSocketUpgrade(c)`: Checks for `Connection: Upgrade` and `Upgrade: websocket` headers.
*   `websocket.New(handler)`: The main middleware that handles the upgrade handshake and calls your `handler` function once the connection is established. The `handler` receives a `*websocket.Conn`.
*   `*websocket.Conn`: The object for interacting with the WebSocket connection. Key methods include:
	*   `ReadMessage()`: Reads an incoming message (blocking).
	*   `WriteMessage(messageType int, data []byte)`: Sends a message.
	*   `WriteJSON(v interface{})`, `ReadJSON(v interface{})`: Sends/receives data in JSON format.
	*   `Close()`: Closes the connection.
	*   `LocalAddr()`, `RemoteAddr()`: Gets local/remote addresses.
	*   `Params(key string)`, `Query(key string)`, `Cookies(key string)`, `Locals(key string)`: Accesses data from the HTTP context *before* the upgrade occurred.
*   `for {}` loop: A common pattern to continuously read messages from the client until the connection is closed or an error occurs.

**Managing Multiple Connections:**
For chat or notification apps, you need a way to track all active connections and send messages to specific connections or all connections (broadcast). This usually involves using maps (with mutexes for concurrency) or Go channels.

```go
import (
	"log"
	"sync" // Needed for mutex if not using channels for sync

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)


// Example (very simple) connection management for broadcasting
var clients = make(map[*websocket.Conn]bool) // Map of active connections
var register = make(chan *websocket.Conn)     // Channel to register new connections
var broadcast = make(chan []byte)             // Channel for broadcast messages
var unregister = make(chan *websocket.Conn)   // Channel to remove connections
var mutex = &sync.Mutex{} // Mutex for protecting the clients map if not using channels for all access


func runHub() {
	for {
		select {
		case conn := <-register:
			mutex.Lock()
			clients[conn] = true
			mutex.Unlock()
			log.Println("Connection registered:", conn.RemoteAddr())
		case message := <-broadcast:
			log.Println("Broadcasting message:", string(message))
			// Send message to all registered clients
			mutex.Lock() // Lock before iterating over clients
			for conn := range clients {
				// Use a separate goroutine for writing to avoid blocking the hub
				go func(c *websocket.Conn, msg []byte) {
					if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Println("Broadcast error to", c.RemoteAddr(), err)
						// Enqueue for unregister if write fails
						unregister <- c
					}
				}(conn, message)
			}
			mutex.Unlock()
		case conn := <-unregister:
			mutex.Lock()
			// Remove connection from map if still present
			if _, ok := clients[conn]; ok {
				log.Println("Connection unregistered:", conn.RemoteAddr())
				delete(clients, conn)
				conn.Close() // Ensure it's closed
			}
			mutex.Unlock()
		}
	}
}


// In main.go
func main() {
	// ... setup fiber app ...

	// Run the hub in a separate goroutine
	go runHub()

	app.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		// 1. Register the new connection
		register <- conn
		log.Printf("New WebSocket connection from: %s", conn.RemoteAddr())

		// 2. Ensure the connection is unregistered when the handler function exits (connection closes)
		defer func() {
			unregister <- conn
		}()

		// 3. Loop reading messages from this client
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				// Check for specific close errors vs. general errors
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					 log.Printf("Websocket read error: %v", err)
				} else {
					log.Printf("Websocket connection closed normally: %s", conn.RemoteAddr())
				}
				break // Exit loop on error/close
			}
		   if mt == websocket.TextMessage {
				// Send received message to the broadcast channel
				log.Printf("Message from %s: %s", conn.RemoteAddr(), string(msg))
				broadcast <- msg
			} else {
				 log.Printf("Received non-text message type: %d", mt)
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
}

```
This is a basic example. Production applications require better error handling, more efficient data structures (especially for many connections), and possibly separating the hub logic into its own package. Using channels for all map access (`register`, `broadcast`, `unregister`) can eliminate the need for the explicit mutex.

### Server-Sent Events (SSE)

SSE is a web standard allowing a server to push updates (events) to a client over a long-lived HTTP connection. Unlike WebSockets, SSE is *unidirectional* (server-to-client). It's simpler than WebSockets and suitable for cases like live news feeds, status updates, notifications, etc., where bidirectional communication isn't needed.

Fiber doesn't have a specific built-in SSE middleware like WebSocket, but SSE is easily implemented using a regular handler because it's fundamentally an HTTP response with `Content-Type: text/event-stream` and a specific message format.

```go
package main

import (
	"bufio"
	"context" // Needed for context checking
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp" // Import fasthttp for Stream
)

func sseHandler(c *fiber.Ctx) error {
	// 1. Set necessary headers for SSE
	c.Set(fiber.HeaderContentType, "text/event-stream")
	c.Set(fiber.HeaderCacheControl, "no-cache")
	c.Set(fiber.HeaderConnection, "keep-alive")
	c.Set(fiber.HeaderTransferEncoding, "chunked") // Or ensure response isn't buffered

	// 2. Use c.Context().SetBodyStreamWriter for streaming responses
	// This allows us to write to the response incrementally.
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		log.Println("SSE Client connected:", c.IP())
		eventID := 0
		ticker := time.NewTicker(2 * time.Second) // Use a ticker for regular events
		defer ticker.Stop() // Ensure ticker is stopped when done

		// Loop to send events periodically
		for {
			select {
			// Check if the client connection is still active
			case <-c.Context().Done(): // Fiber/Fasthttp's context
				log.Println("SSE Client disconnected (context done):", c.IP())
				return // Stop the loop if client disconnects

			// Wait for the next tick
			case t := <-ticker.C:
				eventID++
				// SSE message format:
				// id: <unique_id>\n
				// event: <event_name>\n (optional)
				// data: <your_data>\n\n (data can be multiline if prefixed with 'data: ')

				// Send 'server-time' event with current time data
				fmt.Fprintf(w, "id: %d\n", eventID)
				fmt.Fprintf(w, "event: server-time\n")
				fmt.Fprintf(w, "data: {\"time\": \"%s\"}\n\n", t.Format(time.RFC3339))

				// Send another example event (ping)
				if eventID % 5 == 0 {
					fmt.Fprintf(w, "id: %d-ping\n", eventID)
					fmt.Fprintf(w, "event: ping\n") // Event without data
					fmt.Fprintf(w, "data: \n\n")
				}

				// Flush the buffer to ensure data is sent to the client
				if err := w.Flush(); err != nil {
					// Flushing errors usually mean the client disconnected
					log.Printf("SSE Error flushing / client disconnected: %v", err)
					// Stop the goroutine if flushing fails
					// (c.Context().Done() might not have triggered yet)
					return
				}
				log.Printf("SSE event %d sent to %s", eventID, c.IP())
			}
		}
	})

	// Important: Return nil here because the response is written by the stream writer.
	// Fiber/Fasthttp handles finalizing the response after the stream writer finishes.
	return nil
}

func main() {
	app := fiber.New()

	app.Get("/events", sseHandler)

	// Simple HTML page to test SSE from the browser
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString(`
			<!DOCTYPE html>
			<html>
			<head><title>SSE Test</title></head>
			<body>
				<h1>SSE Test</h1>
				<ul id="events"></ul>
				<script>
					const eventsList = document.getElementById('events');
					console.log("Connecting to SSE stream...");
					const evtSource = new EventSource("/events"); // Connect to the SSE endpoint

					// Handler for named 'server-time' events
					evtSource.addEventListener("server-time", function(event) {
						console.log("Received server-time event:", event.data);
						try {
							const data = JSON.parse(event.data);
							const newItem = document.createElement("li");
							newItem.textContent = "Server Time: " + data.time + " (ID: " + event.lastEventId + ")";
							eventsList.appendChild(newItem);
						} catch (e) {
							console.error("Failed to parse server-time data:", e);
						}
					});

					// Handler for named 'ping' events
					evtSource.addEventListener("ping", function(event) {
						console.log("Received ping event (ID: " + event.lastEventId + ")");
						const newItem = document.createElement("li");
						newItem.textContent = "PING! (ID: " + event.lastEventId + ")";
						eventsList.appendChild(newItem);
					});

					// Handler for default 'message' events (if 'event:' line is omitted)
					evtSource.onmessage = function(event) {
						console.log("Received generic message:", event.data);
						const newItem = document.createElement("li");
						newItem.textContent = "Generic Message: " + event.data;
						eventsList.appendChild(newItem);
					};

					// Handler for connection errors
					evtSource.onerror = function(err) {
						console.error("EventSource failed:", err);
						const newItem = document.createElement("li");
						newItem.textContent = "Connection error or stream closed!";
						newItem.style.color = "red";
						eventsList.appendChild(newItem);
						// Browser might try to reconnect automatically.
						// Close manually if needed:
						// evtSource.close();
					};

					 evtSource.onopen = function() {
						console.log("SSE connection opened.");
						const newItem = document.createElement("li");
						newItem.textContent = "Connection Opened";
						newItem.style.color = "green";
						eventsList.appendChild(newItem);
					};

				</script>
			</body>
			</html>
		`)
	})

	log.Fatal(app.Listen(":3000"))
}
```

**Explanation:**
*   **Headers:** `Content-Type: text/event-stream`, `Cache-Control: no-cache`, `Connection: keep-alive` are essential.
*   **`c.Context().SetBodyStreamWriter(func(w *bufio.Writer))`**: This is the key to streaming in Fiber/Fasthttp. You get a `bufio.Writer` connected directly to the response stream.
*   **Message Format:** Each message ends with two newlines (`\n\n`). Messages can have `id`, `event`, and `data` fields.
*   **`w.Flush()`**: Crucial to call after writing each event to ensure data is sent to the client and not held in the buffer.
*   **Handling Disconnects:** Checking `c.Context().Done()` or errors during `w.Flush()` are ways to detect if the client has closed the connection.
*   **Client-Side (JavaScript):** Uses the browser's built-in `EventSource` object to connect and listen for events. `addEventListener` is used for named events, `onmessage` for default events, `onerror` for errors.

SSE is a great alternative to WebSockets if you only need server-to-client data push.

### Database Integration

Fiber is a web framework, not an ORM (Object-Relational Mapper) or database library. Database integration is done using standard Go libraries or third-party ones, just like in any other Go application.

General steps:

1.  **Choose a Database Library:**
	*   **`database/sql` (Go Built-in):** Standard interface for SQL databases. You need a specific driver for your database (e.g., `github.com/lib/pq` for PostgreSQL, `github.com/go-sql-driver/mysql` for MySQL, `github.com/mattn/go-sqlite3` for SQLite). Provides full control but requires writing manual SQL queries.
	*   **`sqlx` (`github.com/jmoiron/sqlx`):** An extension to `database/sql` that simplifies scanning query results into structs and working with slices/maps. Still requires writing SQL.
	*   **GORM (`gorm.io/gorm`):** The most popular ORM for Go. Provides a high-level API for queries, migrations, relationships, hooks, etc. Reduces manual SQL writing but has its own learning curve.
	*   **NoSQL Drivers:** If using MongoDB, Cassandra, etc., use the official or community drivers for that database (e.g., `go.mongodb.org/mongo-driver`).

2.  **Initialize Connection (Connection Pool):**
	*   Don't open a new connection for every request. Use a *connection pool* managed by the database library.
	*   Initialize the pool when the application starts (e.g., in the `main` function or a separate init function).
	*   Store the pool object (e.g., `*sql.DB`, `*sqlx.DB`, `*gorm.DB`) so it can be accessed by handlers.

3.  **Access Database in Handlers:**
	*   The simplest way is to make the database pool object a global variable (not ideal for testing and large scale).
	*   Better ways:
		*   **Dependency Injection:** Inject the pool object into your handler structs.
		*   **Middleware + Locals:** Create middleware that adds the pool object (or a transaction) to `c.Locals()` for handlers to use.
		*   **Receiver Method:** Define handlers as methods on a struct that has access to the database pool.

**Example with `database/sql` and PostgreSQL (Driver: `lib/pq`):**

```go
package main

import (
	"context" // Recommended for DB operations
	"database/sql"
	"errors" // For checking specific errors like sql.ErrNoRows
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq" // Import driver (blank import)
)

// Struct for product data
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// --- Database Setup ---

// Global variable (simple approach, consider DI for larger projects)
var db *sql.DB

// Function to initialize DB connection
func initDatabase() (*sql.DB, error) {
	// Get DSN (Data Source Name) from environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Fallback DSN for development if env var not set
		dsn = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
		log.Println("DATABASE_URL not set, using default DSN:", dsn)
	}

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25) // Max number of open connections
	db.SetMaxIdleConns(25) // Max number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Max time a connection can be reused

	// Try pinging the database to ensure connection is successful
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		db.Close() // Close if ping fails
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection successful!")
	return db, nil
}

// --- Handlers ---

// Handler to get all products
func getProductsHandler(c *fiber.Ctx) error {
	// Use request context for potential cancellation/timeout
	rows, err := db.QueryContext(c.Context(), "SELECT id, name, price, created_at FROM products ORDER BY created_at DESC")
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return fiber.ErrInternalServerError // Send 500
	}
	defer rows.Close() // Important: Always close rows

	products := []Product{} // Slice to hold results
	for rows.Next() {
		var p Product
		// Scan query result into Product struct fields
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
			log.Printf("Error scanning product row: %v", err)
			return fiber.ErrInternalServerError
		}
		products = append(products, p)
	}

	// Check for errors during row iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating product rows: %v", err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(products) // Send results as JSON
}

// Handler to create a new product
func createProductHandler(c *fiber.Ctx) error {
	// 1. Bind & Validate input
	input := new(struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required,gt=0"`
	})

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	// (Add validation using a validator here if needed)
	// err := validate.Struct(input)
	// if err != nil { ... }

	// 2. Execute INSERT query
	var newID int
	var createdAt time.Time
	// Use request context
	err := db.QueryRowContext(c.Context(),
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, created_at",
		input.Name, input.Price,
	).Scan(&newID, &createdAt) // Scan the returned ID and created_at

	if err != nil {
		log.Printf("Error inserting product: %v", err)
		// Handle specific errors (e.g., constraint violation) if necessary
		return fiber.ErrInternalServerError
	}

	// 3. Create response
	newProduct := Product{
		ID:        newID,
		Name:      input.Name,
		Price:     input.Price,
		CreatedAt: createdAt,
	}

	return c.Status(fiber.StatusCreated).JSON(newProduct)
}

func main() {
	// Initialize database on startup
	var err error
	db, err = initDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Ensure connection is closed when app exits (better handled in graceful shutdown)
	// defer db.Close()

	app := fiber.New()
	app.Use(logger.New())

	// Product routes
	app.Get("/products", getProductsHandler)
	app.Post("/products", createProductHandler)
	// (Add handlers for GET /products/:id, PUT /products/:id, DELETE /products/:id)

	// Implement Graceful Shutdown (see Graceful Shutdown section)
	// ...

	log.Fatal(app.Listen(":3000"))
}
```

**Dependency Injection Pattern (Better):**

```go
// handlers/product_handler.go
package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"myproject/models" // Assuming models are defined here
)

// Handler Struct with DB dependency
type ProductHandler struct {
	DB *sql.DB
}

// Constructor for the handler
func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{DB: db}
}

// Make handlers methods on the ProductHandler struct
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	rows, err := h.DB.QueryContext(c.Context(), "SELECT id, name, price, created_at FROM products ORDER BY created_at DESC")
	if err != nil {
		log.Printf("Error querying products: %v", err)
		return fiber.ErrInternalServerError
	}
	defer rows.Close()

	products := []models.Product{} // Use models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
			log.Printf("Error scanning product row: %v", err)
			return fiber.ErrInternalServerError
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error iterating product rows: %v", err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(products)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	input := new(struct {
		Name  string  `json:"name" validate:"required"`
		Price float64 `json:"price" validate:"required,gt=0"`
	})
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}
	// ... validation ...

	var newID int
	var createdAt time.Time
	err := h.DB.QueryRowContext(c.Context(),
		"INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, created_at",
		input.Name, input.Price,
	).Scan(&newID, &createdAt)
	if err != nil {
		log.Printf("Error inserting product: %v", err)
		return fiber.ErrInternalServerError
	}

	newProduct := models.Product{
		ID:        newID,
		Name:      input.Name,
		Price:     input.Price,
		CreatedAt: createdAt,
	}
	return c.Status(fiber.StatusCreated).JSON(newProduct)
}


// main.go
package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"context"
	"errors"
	"net/http"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq" // Driver

	"myproject/handlers" // Adjust import path
	"myproject/models"   // Adjust import path
)

// initDatabase function same as before
func initDatabase() (*sql.DB, error) { /* ... */ return nil, nil }

func main() {
	db, err := initDatabase()
	if err != nil { log.Fatalf("DB init failed: %v", err) }

	app := fiber.New()
	app.Use(logger.New())

	// Create handler instance, injecting the DB
	productHandler := handlers.NewProductHandler(db)

	// Register routes using methods from the handler instance
	api := app.Group("/api") // Example grouping
	api.Get("/products", productHandler.GetProducts)
	api.Post("/products", productHandler.CreateProduct)
	// ... other product routes ...

	// Graceful shutdown logic (simplified)
	go func() {
		if err := app.Listen(":3000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Close DB connection during shutdown
	defer func() {
		log.Println("Closing database connection...")
		if err := db.Close(); err != nil {
		   log.Printf("Error closing DB: %v", err)
		}
	}()


	if err := app.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
```
This DI pattern makes handlers more testable (you can mock `*sql.DB`) and better structured.

### Authentication & Authorization (JWT, Sessions)

Securing API endpoints or web pages is a common requirement. Two popular approaches are:

1.  **Sessions (Session-based Authentication):**
	*   **How it works:**
		1.  User logs in with username/password.
		2.  Server verifies credentials.
		3.  Server creates a unique session ID, stores it (in memory, Redis, DB) along with user data (e.g., user ID).
		4.  Server sends the session ID to the client, usually stored in a *cookie*.
		5.  For subsequent requests, the client sends the session cookie.
		6.  Server retrieves the session ID from the cookie, looks it up in the session store, and gets the associated user data. If the session is valid, access is granted.
		7.  On logout, the session is deleted on the server.
	*   **Pros:** Relatively simple concept, stateful (easy to store session data), server-side logout is straightforward (delete session).
	*   **Cons:** Stateful (requires server-side storage, can be a bottleneck), scalability issues (if using memory & prefork/multi-server), less suitable for purely stateless APIs or mobile apps (cookies aren't always ideal).
	*   **Implementation in Fiber:** Use the `session` middleware (`github.com/gofiber/fiber/v2/middleware/session`). This middleware supports various *storage* providers (memory, Redis, PostgreSQL, MySQL, etc.).

2.  **Tokens (JWT - JSON Web Token):**
	*   **How it works:**
		1.  User logs in.
		2.  Server verifies credentials.
		3.  Server creates a *token* (JWT) containing a *payload* (user data like ID, role, etc.) and a *signature* (using a secret key known only to the server).
		4.  Server sends the token to the client. Client stores it (localStorage, sessionStorage, Authorization header).
		5.  For subsequent requests, the client sends the token (usually in the `Authorization: Bearer <token>` header).
		6.  Server receives the token, *verifies the signature* using the secret key. If the signature is valid, the server *trusts* the payload within the token (no need to check DB/session store).
		7.  Server uses data from the payload (e.g., user ID) for authorization.
	*   **Pros:** Stateless (server doesn't need to store session state), suitable for APIs & microservices, scalable, good for mobile apps, many libraries available.
	*   **Cons:** Tokens can be larger than session IDs, logout is more complex (tokens can't be invalidated on the server by default, requires blacklisting/short expiry), payload data is visible (though tamper-proof without the secret key), requires refresh token handling for long sessions.
	*   **Implementation in Fiber:** Use the official JWT middleware (`github.com/gofiber/contrib/jwt`) or a popular Go JWT library (`github.com/golang-jwt/jwt/v5`) and create custom middleware.

**Example Custom JWT Middleware (using `golang-jwt/jwt/v5`):**

```go
// middleware/auth.go
package middleware

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing tokens (store in env var!)
var jwtSecret = []byte(getJwtSecret()) // Use helper func

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("Warning: JWT_SECRET environment variable not set. Using default insecure secret.")
		return "a-very-insecure-secret-key-replace-me" // !! CHANGE THIS !!
	}
	return secret
}


// Struct for custom JWT claims (payload)
type MyCustomClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims // Embed standard claims (Issuer, Subject, Audience, ExpiresAt, NotBefore, IssuedAt, JWTID)
}

// Function to generate a JWT token
func GenerateJWT(userID int, role string) (string, error) {
	// Set custom claims
	claims := MyCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			// Set expiration time (e.g., 1 hour)
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app", // Your application name
			Subject:   "user-auth",
		},
	}

	// Create a new token with claims and HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	return signedToken, nil
}

// Middleware to protect routes
func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": "error", "message": "Missing Authorization Header",
			})
		}

		// Check format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": "error", "message": "Invalid Authorization format (must be 'Bearer <token>')",
			})
		}
		tokenString := parts[1]

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Ensure signing method is HMAC (HS256) as we used
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Return our secret key
			return jwtSecret, nil
		})

		if err != nil {
			log.Printf("Error parsing/validating token: %v", err)
			// Check specific error types (e.g., token expired)
			if errors.Is(err, jwt.ErrTokenExpired) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status":"error", "message": "Token has expired"})
			}
			// Other errors (invalid signature, malformed token, etc.)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status":"error", "message": "Invalid token"})
		}

		// If token is valid, extract claims
		if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
			// Store user information from claims into Locals for handler use
			c.Locals("userID", claims.UserID)
			c.Locals("userRole", claims.Role)
			c.Locals("jwtClaims", claims) // Store all claims if needed
			log.Printf("JWT Middleware: Access granted for User ID %d (Role: %s)", claims.UserID, claims.Role)
			return c.Next() // Proceed to the handler
		}

		// If claims couldn't be cast or token is invalid (should have been caught by err check)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status":"error", "message": "Invalid token or claims"})
	}
}

// (Optional) Middleware for role-based authorization
func AuthorizeRole(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get role from Locals (assuming Protected middleware ran first)
		role, ok := c.Locals("userRole").(string)
		if !ok {
			log.Println("Authorize Middleware: User role not found in Locals")
			 return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Access denied (role unknown)"})
		}

		// Check if user's role is in the list of allowed roles
		isAllowed := false
		for _, allowed := range allowedRoles {
			if role == allowed {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			log.Printf("Authorize Middleware: Access denied for role '%s'. Allowed roles: %v", role, allowedRoles)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Insufficient permissions"}) // 403 Forbidden
		}

		log.Printf("Authorize Middleware: Access allowed for role '%s'", role)
		return c.Next() // Proceed if allowed
	}
}


// --- Example Usage in main.go or auth_routes.go ---
/*
package main

import (
	"fmt"
	"log"
	"myproject/middleware" // Import your middleware package
	"github.com/gofiber/fiber/v2"
	// ... other imports
)

func main() {
	app := fiber.New()
	// ... other middleware (logger, recover) ...

	// --- Authentication Routes ---
	auth := app.Group("/auth")
	auth.Post("/login", func(c *fiber.Ctx) error {
		// 1. Get username/password from body
		// 2. Verify against database
		// 3. If valid, generate JWT
		userID := 123 // Example ID from DB
		userRole := "admin" // Example role from DB
		token, err := middleware.GenerateJWT(userID, userRole)
		if err != nil {
			 log.Printf("Failed to generate token: %v", err)
			 return fiber.ErrInternalServerError
		}
		// 4. Send token to client
		return c.JSON(fiber.Map{"token": token})
	})

	// --- Protected Routes ---
	api := app.Group("/api")
	api.Use(middleware.Protected()) // Apply JWT middleware here

	api.Get("/me", func(c *fiber.Ctx) error {
		// Get data from Locals set by middleware
		userID := c.Locals("userID").(int)
		userRole := c.Locals("userRole").(string)
		return c.JSON(fiber.Map{"user_id": userID, "role": userRole})
	})

	// --- Protected Routes with Role Authorization ---
	adminApi := api.Group("/admin")
	// Only users with 'admin' role can access endpoints in this group
	adminApi.Use(middleware.AuthorizeRole("admin"))

	adminApi.Get("/users", func(c *fiber.Ctx) error {
		// Only admins can reach here
		return c.JSON(fiber.Map{"message": "List of all users (admin only)"})
	})

	// Another endpoint accessible by 'admin' or 'editor' roles
	api.Post("/articles", middleware.AuthorizeRole("admin", "editor"), func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(int)
		return c.JSON(fiber.Map{"message": fmt.Sprintf("Creating new article (by user %d)", userID)})
	})


	log.Fatal(app.Listen(":3000"))
}
*/
```

**Important for JWT:**
*   **Secret Key:** Keep your secret key confidential! Store it in environment variables or a secrets management system, don't hardcode it.
*   **Expiration:** Always set an expiration time (`ExpiresAt`) on tokens to limit their validity period.
*   **HTTPS:** Always use HTTPS to transmit tokens to prevent theft.
*   **Payload:** Don't store sensitive data in the JWT payload as it's easily readable (though tamper-proof). Store user IDs, roles, or other non-sensitive information.
*   **Refresh Tokens:** For longer sessions, implement a *refresh token* mechanism. The main token (access token) has a short lifespan (minutes/hours), while the refresh token has a longer lifespan (days/weeks) and is stored more securely. The client uses the refresh token to obtain a new access token without logging in again.
*   **Blacklisting (Optional):** If you need a way to invalidate tokens before they expire (e.g., on password change or forced logout), you need a server-side blacklist implementation (e.g., in Redis) to store invalid token IDs. The JWT middleware needs to check this blacklist.

Choose the authentication strategy (session or token) that best suits your application's needs.

### Testing

Writing automated tests is crucial to ensure your application works as expected and to prevent regressions when you make changes. Go has strong built-in testing support, and it works well for testing Fiber applications.

The common approach is to use the `net/http/httptest` package to create mock HTTP requests and pass them directly to your Fiber handler without needing to run a real HTTP server.

**Test Structure:**

*   Test files must end with `_test.go` (e.g., `main_test.go`, `handlers/product_handler_test.go`).
*   Test functions must start with `Test` and accept `t *testing.T` (e.g., `TestGetProductsHandler`).
*   Use Go's `testing` package and an assertion library (optional but helpful) like `testify/assert` or `testify/require`.

**Example Test for a Simple Handler:**

Suppose we have a `main.go` like this:
```go
// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/echo", func(c *fiber.Ctx) error {
		body := c.Body() // Get raw body
		// Send back the same body with the original content type
		return c.Status(fiber.StatusOK).Send(body)
	})

	app.Get("/users/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.JSON(fiber.Map{"user": name})
	})
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}
```

Now create a `main_test.go` file:

```go
// main_test.go
package main // Must be in the same package as the code being tested

import (
	"encoding/json" // Need this for JSON handling
	"io"
	"net/http"          // Import net/http for status code constants etc.
	"net/http/httptest" // Import httptest
	"strings"
	"testing" // Import testing

	"github.com/gofiber/fiber/v2" // Import fiber
	"github.com/stretchr/testify/assert" // Import testify/assert (optional)
	"github.com/stretchr/testify/require" // Import testify/require (optional)
)

// Helper function to create a request and get a response
func performRequest(app *fiber.App, method, path string, body io.Reader, headers map[string]string) *http.Response {
	// Create a mock HTTP request
	req := httptest.NewRequest(method, path, body)

	// Add headers if provided
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Execute the request through the Fiber handler
	// app.Test() is Fiber's specific method for testing
	resp, err := app.Test(req, -1) // -1 for no timeout
	if err != nil {
		// Fail test if app.Test() fails (rare)
		panic(err) // or t.Fatalf("app.Test failed: %v", err) if t is available
	}

	return resp // Return the standard http.Response
}


// Test for the GET / route
func TestGetRoot(t *testing.T) {
	// 1. Setup Fiber Application (just for this test)
	app := fiber.New()
	setupRoutes(app) // Register routes

	// 2. Create Request
	// req := httptest.NewRequest("GET", "/", nil)
	// resp, _ := app.Test(req)
	resp := performRequest(app, "GET", "/", nil, nil)
	defer resp.Body.Close() // Important: Always close the response body

	// 3. Assertions (using testify/assert)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200 OK")

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Should be able to read response body") // require: fail test if error occurs
	assert.Equal(t, "Hello, World!", string(bodyBytes), "Response body should be 'Hello, World!'")
}

// Test for the POST /echo route
func TestPostEcho(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	// Data to be sent in the body
	requestBody := `{"message": "ping"}`
	headers := map[string]string{"Content-Type": "application/json"} // Set header

	resp := performRequest(app, "POST", "/echo", strings.NewReader(requestBody), headers)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200 OK")

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	// Compare response body with request body
	assert.JSONEq(t, requestBody, string(bodyBytes), "Response body should be the same as the request body (JSON)")
	// Check response Content-Type (although the handler doesn't set it explicitly, Fiber/Fasthttp might detect it)
	// assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
}

// Test for the GET /users/:name route
func TestGetUserByName(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	userName := "alice"
	path := "/users/" + userName

	resp := performRequest(app, "GET", path, nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200 OK")
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json", "Content-Type should be application/json")

	// Read and unmarshal the JSON response body
	var result map[string]string
	err := json.NewDecoder(resp.Body).Decode(&result) // Import "encoding/json"
	require.NoError(t, err, "Should be able to decode JSON response")

	// Check JSON content
	assert.Equal(t, userName, result["user"], "Value of 'user' in JSON should match the path parameter")
}

// Test for a non-existent route (404 Not Found)
func TestNotFound(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	resp := performRequest(app, "GET", "/path/does/not/exist", nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Status code should be 404 Not Found")
}

// Don't forget to import "encoding/json"
```

**Running Tests:**

Open a terminal in your project directory and run:

```bash
go test ./... -v
```

*   `go test`: The basic command to run tests.
*   `./...`: Run tests in the current directory and all subdirectories.
*   `-v`: Verbose mode, shows the names of tests being run and their results (PASS/FAIL).

**Testing with Databases:**

Testing handlers that interact with a database requires additional strategies:

1.  **Separate Test Database:** The best approach is to use a separate test database (e.g., a temporary Docker database or a dedicated test database). Your tests connect to this database, perform necessary data setup, run the handler, and then clean up the data (or roll back transactions). This ensures tests run in isolation and do not interfere with the development/production database.
2.  **Database Mocking:** Using mocking libraries (like `sqlmock` for `database/sql` or GORM's mocking features) to simulate database behavior without actually connecting to a database. This is faster but might not capture all real database behaviors. Suitable for pure unit tests of handler logic.

**Example (Concept) Testing Handler with DB:**

```go
// product_handler_test.go
package handlers_test // Use _test package for black-box testing

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time" // Needed for checking CreatedAt

	"myproject/handlers" // Import the actual handler package
	"myproject/models"   // Import model (e.g., struct Product)
	_ "github.com/lib/pq" // Import driver for side effects (registration)
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to set up the test DB (simple example, ideally use Docker/migrations)
func setupTestDB(t *testing.T) *sql.DB {
	dsn := "postgres://user:password@localhost:5433/testdb?sslmode=disable" // Test DB on a different port
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err, "Must be able to connect to the test DB")
	// Clean the table before the test (or use transactions)
	_, err = db.Exec("DELETE FROM products")
	require.NoError(t, err)
	return db
}

func TestCreateProductHandler(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Setup Fiber App and Handler with the test DB
	app := fiber.New()
	productHandler := handlers.NewProductHandler(db) // Use DI
	app.Post("/products", productHandler.CreateProduct)

	// Input data
	inputData := map[string]interface{}{
		"name": "Test Product",
		"price": 99.99,
	}
	body, _ := json.Marshal(inputData)

	// Create request
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := app.Test(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Should be 201 Created")

	// Check response body
	var createdProduct models.Product
	err = json.NewDecoder(resp.Body).Decode(&createdProduct)
	require.NoError(t, err)
	assert.Equal(t, inputData["name"], createdProduct.Name)
	assert.Equal(t, inputData["price"], createdProduct.Price)
	assert.NotZero(t, createdProduct.ID, "Product ID should be generated")
	assert.NotZero(t, createdProduct.CreatedAt, "CreatedAt should be populated")

	// (Optional) Verify data in the test DB
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM products WHERE id = $1", createdProduct.ID).Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count, "Product should be saved in the DB")
}

// ... other tests (GetProducts, GetProductByID, etc.) ...
```

**Testing Tips:**
*   **Focus on Input and Output:** Handler tests should focus on whether different request inputs produce the expected response outputs (status code, headers, body).
*   **Isolation:** Aim to test each handler separately. If a handler calls another service, consider mocking that service in the handler's unit test.
*   **Use Test Suites:** For more structured tests, use Go's test suite features or `testify/suite` for better setup/teardown.
*   **Code Coverage:** Use `go test ./... -cover` to see the percentage of your code covered by tests. Aim for high coverage, but focus on testing critical paths and complex logic.

### Deployment

Once your application is ready, the next step is to deploy it to a production server. There are many ways to deploy Go/Fiber applications:

1.  **Compile & Copy Binary:**
	*   **How it Works:** Compile your Go application into a *single static binary* on your development machine (or CI/CD server), then copy that binary to the production server and run it.
	*   **Compile:**
		```bash
		# For Linux (most common for servers)
		GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o myapp main.go

		# -ldflags="-s -w": Strips debug symbols and DWARF info, resulting in a smaller binary.
		# -o myapp: Output binary name.
		```
	*   **Run on Server:** Copy `myapp` to the server, then run: `./myapp`
	*   **Pros:** Very simple, minimal dependencies on the server (just the binary), fast startup.
	*   **Cons:** Requires manual server setup (OS, firewall, reverse proxy), needs a way to run the binary as a service (to auto-restart on crash or server reboot).

2.  **Using a Supervisor (Systemd, SupervisorD):**
	*   **How it Works:** Use a process manager like `systemd` (common in modern Linux) or `supervisord` to manage your application binary as a *service*. The manager handles start, stop, auto-restart, logging, etc.
	*   **Example Systemd Unit File (`/etc/systemd/system/myapp.service`):**
		```ini
		[Unit]
		Description=My Fiber Application
		After=network.target # Start after the network is ready

		[Service]
		User=appuser       # Run as a non-root user
		Group=appuser
		WorkingDirectory=/path/to/app/directory # Application's working directory
		EnvironmentFile=/path/to/app/directory/.env # Load env vars from file (optional)
		ExecStart=/path/to/binary/myapp           # Path to your binary
		Restart=on-failure                       # Restart if it fails
		RestartSec=5s                            # Wait 5 seconds before restarting
		StandardOutput=journal                   # Redirect stdout to journald
		StandardError=journal                    # Redirect stderr to journald
		LimitNOFILE=65536                        # Increase file descriptor limit (important for many connections)

		[Install]
		WantedBy=multi-user.target # Enable at multi-user boot
		```
	*   **Systemd Commands:**
		*   `sudo systemctl daemon-reload`: Reload config after creating/modifying the service file.
		*   `sudo systemctl enable myapp`: Enable the service to start on boot.
		*   `sudo systemctl start myapp`: Start the service now.
		*   `sudo systemctl status myapp`: Check the service status.
		*   `sudo systemctl stop myapp`: Stop the service.
		*   `sudo journalctl -u myapp -f`: View service logs (follow new logs).
	*   **Pros:** Reliable process management, auto-restart, centralized logging.
	*   **Cons:** Still requires manual server setup.

3.  **Docker:**
	*   **How it Works:** Package your application and its dependencies (including the Go runtime if needed, though static binaries are more common) into a *Docker image*. Run this image as a *container* on any server with Docker installed.
	*   **Example `Dockerfile` (Multi-stage build for small binary):**
		```dockerfile
		# --- Stage 1: Build ---
		FROM golang:1.21-alpine AS builder

		WORKDIR /app

		# Copy go mod and sum, then download dependencies (leverages Docker cache)
		COPY go.mod go.sum ./
		RUN go mod download

		# Copy the rest of the source code
		COPY . .

		# Build static binary
		# CGO_ENABLED=0 is important for static builds on Alpine
		# -ldflags="-s -w" for smaller size
		RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /myapp main.go

		# --- Stage 2: Run ---
		# Use a very small base image (scratch or alpine)
		FROM alpine:latest
		# FROM scratch # Smallest image, but no shell/tools

		WORKDIR /app

		# (Optional) Install CA certificates if the app needs outgoing HTTPS connections
		RUN apk --no-cache add ca-certificates

		# Copy the built binary from the builder stage
		COPY --from=builder /myapp /myapp

		# (Optional) Copy static files or templates if not embedded
		# COPY --from=builder /app/public ./public
		# COPY --from=builder /app/views ./views

		# Expose the port the Fiber app listens on
		EXPOSE 3000

		# Set default environment variables (can be overridden at runtime)
		ENV APP_ENV=production
		ENV LISTEN_ADDR=:3000
		# ENV DATABASE_URL=...

		# Command to run the application when the container starts
		# Entrypoint ensures the binary is the main command
		ENTRYPOINT ["/myapp"]
		# Cmd can be added for default arguments to the entrypoint
		# CMD ["--some-flag"]
		```
	*   **Build & Run:**
		*   `docker build -t myapp-image .`
		*   `docker run -d -p 8080:3000 --name myapp-container -e DATABASE_URL="xxx" myapp-image`
			*   `-d`: Run detached (background).
			*   `-p 8080:3000`: Map host port 8080 to container port 3000.
			*   `--name`: Name the container.
			*   `-e`: Set environment variables.
	*   **Pros:** Consistent environment (dev, staging, prod), isolated dependencies, easily scalable, portable, rich tooling (Docker Compose, Kubernetes).
	*   **Cons:** Need to learn Docker, images can be large if not optimized, container overhead (small).

4.  **Kubernetes (K8s):**
	*   **How it Works:** Container orchestration platform. You define your application's desired state (deployments, services, ingresses, configmaps, secrets) in YAML files, and Kubernetes handles deployment, scaling, load balancing, self-healing, etc., across a cluster of servers.
	*   **Pros:** Highly scalable, high availability, manages complexity for large applications/microservices.
	*   **Cons:** Very complex, steep learning curve, requires cluster infrastructure.

5.  **Platform as a Service (PaaS):**
	*   **How it Works:** Platforms like Heroku, Google App Engine, AWS Elastic Beanstalk, Render, Fly.io. You typically just push your code (or Docker image), and the platform handles the infrastructure, deployment, scaling, logging.
	*   **Pros:** Very easy to use, focus on code not infrastructure, auto-scaling (depending on platform).
	*   **Cons:** Less flexible than self-hosted/IaaS solutions, can be more expensive at scale, potential vendor lock-in.

6.  **Serverless (Functions as a Service - FaaS):**
	*   **How it Works:** Platforms like AWS Lambda, Google Cloud Functions, Azure Functions. You deploy individual handler function code. The platform automatically runs the function when requests come in and scales automatically (even to zero when idle).
	*   **Pros:** Pay-per-use, extreme auto-scaling, no server management.
	*   **Cons:** Cold starts (latency on the first request after idle), runtime/duration limits, state management is difficult, less suitable for stateful apps or long-running connections (WebSocket/SSE need other solutions). Good for stateless APIs or small tasks.

**Recommendations:**

*   **Small/Personal Projects:** Compile & Copy + Systemd/SupervisorD, or PaaS (Render, Fly.io).
*   **Medium Apps/Startups:** Docker + Docker Compose (for multi-container apps), or PaaS.
*   **Large Apps/Microservices:** Docker + Kubernetes, or a scalable PaaS.
*   **Very Specific APIs/Event-Driven:** Serverless (FaaS).

**Reverse Proxy (Nginx, Caddy, Traefik):**

It is almost always **recommended** to run your Fiber application behind a *reverse proxy* in production.

*   **Reverse Proxy Tasks:**
	*   **SSL/TLS Termination:** Handles HTTPS encryption, allowing your app to run plain HTTP behind the proxy.
	*   **Load Balancing:** Distributes traffic if you run multiple instances of your app.
	*   **Caching:** Caches static or dynamic responses.
	*   **Compression:** Compresses responses (though Fiber can do this too).
	*   **Rate Limiting/Security:** Adds an extra layer of security.
	*   **Serving Static Files:** Can be more efficient than Fiber for high-volume static files.
	*   **Host/Path-Based Routing:** Directs requests to different applications on the same server.

**Simple Nginx Configuration Example:**

```nginx
# /etc/nginx/sites-available/myapp.conf

server {
	listen 80; # Listen on port 80 (HTTP)
	server_name yourdomain.com www.yourdomain.com; # Your domain(s)

	# Redirect HTTP to HTTPS (if using SSL)
	# location / {
	#     return 301 https://$host$request_uri;
	# }

	# If using SSL (recommended), configure in a separate server block
	# listen 443 ssl http2;
	# server_name yourdomain.com www.yourdomain.com;
	# ssl_certificate /path/to/fullchain.pem;
	# ssl_certificate_key /path/to/privkey.pem;
	# include /etc/letsencrypt/options-ssl-nginx.conf; # Additional SSL settings
	# ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

	location / {
		proxy_pass http://localhost:3000; # Forward requests to the Fiber app on port 3000
		proxy_set_header Host $host;             # Pass the original Host header
		proxy_set_header X-Real-IP $remote_addr; # Pass the real client IP
		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header X-Forwarded-Proto $scheme; # Pass the protocol (http/https)

		# Additional settings for WebSocket if needed
		# proxy_http_version 1.1;
		# proxy_set_header Upgrade $http_upgrade;
		# proxy_set_header Connection "upgrade";
		# proxy_read_timeout 86400; # Long timeout for WebSocket connections
	}

	# (Optional) Let Nginx serve static files directly
	# location /static/ {
	#     alias /path/to/app/directory/public/;
	#     expires 7d;      # Cache in browser for 7 days
	#     access_log off; # Turn off access logging for static files
	# }
}

# Don't forget to create a symlink:
# sudo ln -s /etc/nginx/sites-available/myapp.conf /etc/nginx/sites-enabled/
# sudo systemctl restart nginx
```

### Performance & Optimization

Fiber is already very fast by default, but here are some things you can do to optimize performance further:

*   **Prefork:** As discussed earlier, `Prefork: true` can help utilize all CPU cores for stateless applications on Linux/BSD. Benchmark for your specific case.
*   **Avoid Excessive Allocations:** Go is fast, but the garbage collector (GC) can be a bottleneck. Avoid unnecessary memory allocations in hot paths (frequently called handlers). Use `sync.Pool` for frequently reused objects. Fasthttp and Fiber already use pools internally extensively.
*   **Use `fiber.Map` vs `map[string]interface{}`:** For simple JSON responses, `fiber.Map` might be slightly more efficient as it's an alias but might have internal pool-related optimizations.
*   **Fast Binding:** `BodyParser` is quite fast, but if you know the input format for sure (e.g., always JSON), using `c.App().JSONDecoder(body, &out)` (after `c.Body()`) might be slightly faster as it skips type detection. Measure the difference.
*   **Selective Middleware:** Only use the middleware you actually need. Apply middleware at the narrowest scope (route/group) if it's not globally required.
*   **Efficient Database Queries:** The most common bottleneck is often the database.
	*   Use proper indexes on frequently queried columns.
	*   Avoid N+1 queries (fetch related data in one or two queries, not many small ones).
	*   Use `SELECT` only for the columns you need, not `SELECT *`.
	*   Optimize slow queries (use `EXPLAIN ANALYZE` in your DB).
	*   Use a well-configured connection pool.
*   **Caching:** Implement caching for data that changes infrequently or is expensive to generate.
	*   **HTTP Caching:** Use `Cache-Control`, `ETag`, `Last-Modified` headers so browsers or CDNs can cache responses. Fiber's `etag` middleware can help.
	*   **Server-Side Caching:** Store results of DB queries or expensive computations in memory (e.g., using `ristretto`, `go-cache`) or an external store (Redis, Memcached). Cache middleware exists for Fiber (built-in or third-party) that can automatically cache responses based on URL/headers.
*   **Compression:** Enable Gzip/Brotli compression (e.g., with the `compress` middleware or in the reverse proxy) to reduce response sizes and speed up transfers.
*   **Application Profiling:** Use Go's profiling tools (`pprof`) to accurately identify CPU and memory bottlenecks. Fiber makes integrating `pprof` easy via the `pprof` middleware. Analyze the results with `go tool pprof`.
*   **Build Flags:** Use `-ldflags="-s -w"` when building for production (smaller binary).
*   **Latest Go Version:** Always use the latest stable Go version, as it often includes performance improvements.

**Important:** Avoid premature optimization. Write clean, correct code first, then use profiling to find *real* bottlenecks before optimizing.

### Project Structure

Organizing your project code well becomes important as the application grows. There's no single "correct" way, but here are some common patterns:

**1. Flat Structure (Simple):**

Suitable for very small projects. All files in the root.

```
my-fiber-project/
├── go.mod
├── go.sum
├── main.go         # Fiber setup, DB, global Middleware, Routes
├── handlers.go     # All handler functions
├── models.go       # Data struct definitions (User, Product)
└── main_test.go
```

**2. Feature/Domain-Based:**

Group code by feature or business domain.

```
my-fiber-project/
├── go.mod
├── go.sum
├── config/             # Config loading functions/structs
│   └── config.go
├── internal/           # Internal packages (cannot be imported from outside the project)
│   ├── database/       # DB connection & initialization code
│   │   └── database.go
│   ├── middleware/     # Custom middleware
│   │   └── auth.go
│   ├── models/         # Data/domain structs
│   │   ├── user.go
│   │   └── product.go
│   ├── auth/           # Authentication feature
│   │   ├── handler.go  # Auth-related handlers (login, register)
│   │   ├── service.go  # Auth business logic (optional)
│   │   └── routes.go   # Setup /auth/* routes
│   ├── product/        # Product feature
│   │   ├── handler.go  # Product CRUD handlers
│   │   ├── repository.go # Product DB access logic (interface + implementation)
│   │   ├── service.go  # Product business logic (optional)
│   │   └── routes.go   # Setup /products/* routes
│   └── ...             # Other features/domains
├── cmd/                # Application entry points (main package)
│   └── api/
│       └── main.go     # Main Fiber setup, DB init, call route setups
├── web/                # Frontend files (if any)
│   ├── templates/
│   └── static/
└── Dockerfile
```
*   **`internal/`**: Packages here can only be imported by other code within `my-fiber-project`. This is good for hiding implementation details.
*   **Separation of Concerns:**
	*   `handlers` (or `xxx/handler.go`): Receive Fiber requests, parse input, call services/repositories, format responses. Should *not* contain business logic or direct DB queries.
	*   `service` (optional): Contains core business logic, decoupled from HTTP or DB. Can be reused.
	*   `repository` (or `store`, `dao`): Contains data access logic (DB queries). Usually an interface and concrete implementations (e.g., `PostgresProductRepository`). This aids testing and storage swapping.
	*   `routes` (or `router`): Defines Fiber routes and maps them to the appropriate handlers.
*   **Dependency Injection:** Fits well with this structure. `main.go` initializes the DB, creates repository instances, then services, then handlers (injecting dependencies downwards), and finally registers the routes.

**3. Layered Structure:**

Similar to the domain structure, but the primary grouping is by architectural layer.

```
my-fiber-project/
├── go.mod
├── go.sum
├── config/
├── pkg/                # Reusable libraries/utilities (can be imported from outside)
│   └── validator/
├── internal/
│   ├── delivery/       # Presentation layer (e.g., HTTP)
│   │   ├── http/
│   │   │   ├── middleware/
│   │   │   ├── handler/    # All HTTP handlers
│   │   │   │   ├── product_handler.go
│   │   │   │   └── user_handler.go
│   │   │   └── router.go   # Setup all HTTP routes
│   │   └── grpc/       # (If using gRPC delivery)
│   ├── usecase/        # Business logic layer (services)
│   │   ├── product_usecase.go
│   │   └── user_usecase.go
│   ├── repository/     # Data access layer
│   │   ├── postgres/   # Repo implementation for Postgres
│   │   │   ├── product_repo.go
│   │   │   └── user_repo.go
│   │   └── redis/      # Repo implementation for Redis (cache, etc.)
│   ├── domain/         # Core entity/model definitions & repo/usecase interfaces
│   │   ├── product.go
│   │   └── user.go
├── cmd/api/main.go     # Application entry point
└── ...
```

Choose the structure that makes the most sense for your project's size and complexity. Start simple and refactor as it grows. Consistency is key.

### Graceful Shutdown

When you need to stop or restart your application server (e.g., during a new version deployment or maintenance), it's important to do it *gracefully*. This means:

1.  The server stops accepting *new* connections.
2.  The server waits for *currently active* connections to finish processing (within a certain time limit).
3.  After all connections are finished or the timeout is reached, the server truly stops.

This prevents abrupt disconnections that can cause errors for clients or data loss.

Fiber (via Fasthttp) provides the `app.Shutdown()` or `app.ShutdownWithTimeout()` methods for this. You need to listen for OS signals (like `SIGINT` from Ctrl+C or `SIGTERM` from `kill`/systemd) to trigger the shutdown.

```go
package main

import (
	// "context" // Required for ShutdownWithContext (optional but good)
	"errors"    // For checking specific errors
	"log"
	"net/http" // For http.ErrServerClosed
	"os"
	"os/signal" // To catch OS signals
	"syscall"   // For signal constants (SIGINT, SIGTERM)
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// ... import DB, etc. ...
)

// func initDatabase() (*sql.DB, error) { ... }
// var db *sql.DB

func main() {
	// --- Application Setup ---
	app := fiber.New(fiber.Config{
		// ReadTimeout/WriteTimeout/IdleTimeout help with graceful shutdown
		IdleTimeout: 5 * time.Second,
	})
	app.Use(logger.New())

	// Database setup (example)
	// db, err := initDatabase()
	// if err != nil { log.Fatalf("DB init failed: %v", err) }
	// defer func() {
	//     log.Println("Closing database connection...")
	//     if err := db.Close(); err != nil {
	//         log.Printf("Error closing DB: %v", err)
	//     }
	// }() // Defer DB closure

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("Processing request to / ...")
		time.Sleep(3 * time.Second) // Simulate long work
		log.Println("Finished processing request to /")
		return c.SendString("Request processed!")
	})

	// --- Start Server in Goroutine ---
	// Run app.Listen in a goroutine so it doesn't block the main goroutine,
	// allowing us to wait for shutdown signals.
	go func() {
		listenAddr := ":3000" // Get from config/env var
		log.Printf("Server starting to listen on %s", listenAddr)
		if err := app.Listen(listenAddr); err != nil {
			// Errors other than http.ErrServerClosed are fatal
			if !errors.Is(err, http.ErrServerClosed) {
				 log.Fatalf("Failed to run server: %v", err)
			}
			 log.Println("Server stopped listening.")
		}
	}()

	// --- Wait for Shutdown Signal ---
	quit := make(chan os.Signal, 1) // Channel to receive signals
	// Catch SIGINT (Ctrl+C) and SIGTERM (kill default, systemd stop)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	receivedSignal := <-quit
	log.Printf("Received shutdown signal: %s. Starting graceful shutdown...", receivedSignal)

	// --- Graceful Shutdown Process ---
	// Give a deadline for shutdown (e.g., 10 seconds)
	// shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// Call app.Shutdown() or app.ShutdownWithTimeout()
	// err := app.ShutdownWithContext(shutdownCtx) // Fiber v2.45.0+

	// Or without context timeout (Fiber will wait for idle connections based on IdleTimeout)
	err := app.Shutdown()

	if err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Graceful shutdown complete. Server exiting.")
	// Here, the deferred DB connection closure will also run.
}
```

**Explanation:**
1.  **Start Server in Goroutine:** `app.Listen()` is blocking. We run it in a goroutine so `main` can continue to wait for signals.
2.  **Setup Signal Channel:** Create the `quit` channel and use `signal.Notify` to direct `SIGINT` and `SIGTERM` signals to it.
3.  **Wait for Signal:** `<-quit` blocks execution until one of the specified signals is received.
4.  **Call `app.Shutdown()`:** Once a signal is received, call `app.Shutdown()`. Fiber will:
	*   Stop accepting new connections.
	*   Wait for active connections to finish (limited by `IdleTimeout` in `fiber.Config`).
	*   Close the listener.
5.  **Handle Shutdown Error:** Check for errors from `Shutdown()`.
6.  **Other Cleanup:** If other resources need closing (like DB connections), do it after `Shutdown()` completes or use `defer` in `main`.

With this implementation, when you stop the application (Ctrl+C or `systemctl stop`), it will attempt to finish ongoing requests before exiting completely.

---

## 7. Example Application (Simple CRUD) 📝

Let's create a simple RESTful API to manage "Book" data using Fiber and storing data in memory (for simplicity; use a database in a real application).

```go
package main

import (
	"errors" // Needed for errors.Is
	"fmt"
	"log"
	"net/http" // Needed for http.ErrServerClosed
	"os"
	"os/signal"
	"sync" // For Mutex to prevent race conditions
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// --- Model ---
type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// --- Input Validation Struct ---
type CreateBookInput struct {
	Title  string `json:"title" validate:"required,min=3"`
	Author string `json:"author" validate:"required,min=3"`
}

type UpdateBookInput struct {
	Title  *string `json:"title" validate:"omitempty,min=3"` // Pointer to make optional
	Author *string `json:"author" validate:"omitempty,min=3"` // Pointer to make optional
}

// --- In-Memory Storage ---
var (
	books      = make(map[int]Book) // Map to store books (ID -> Book)
	nextBookID = 1                   // Book ID counter
	bookMutex  = &sync.RWMutex{}     // Mutex to protect access to the books map
	validate   = validator.New()     // Validator instance
)

// --- Error Response Helper ---
func validationErrorResponse(err error) fiber.Map {
	var errorsSlice []fiber.Map // Renamed from 'errors' to avoid conflict
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			errorsSlice = append(errorsSlice, fiber.Map{
				"field":   fieldErr.Field(),
				"tag":     fieldErr.Tag(),
				"message": fmt.Sprintf("Validation failed on '%s' with tag '%s'", fieldErr.Field(), fieldErr.Tag()),
			})
		}
	}
	return fiber.Map{"status": "fail", "errors": errorsSlice}
}

// --- Handlers ---

// GET /books - Get all books
func getBooks(c *fiber.Ctx) error {
	bookMutex.RLock() // Lock for reading
	defer bookMutex.RUnlock()

	// Create a slice from map values
	bookList := make([]Book, 0, len(books))
	for _, book := range books {
		bookList = append(bookList, book)
	}

	return c.JSON(fiber.Map{"status": "success", "data": bookList})
}

// GET /books/:id - Get a book by ID
func getBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid book ID"})
	}

	bookMutex.RLock()
	defer bookMutex.RUnlock()

	book, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Book with ID %d not found", id)})
	}

	return c.JSON(fiber.Map{"status": "success", "data": book})
}

// POST /books - Create a new book
func createBook(c *fiber.Ctx) error {
	input := new(CreateBookInput)

	// Parse & Validate Body
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrorResponse(err))
	}

	bookMutex.Lock() // Lock for writing
	defer bookMutex.Unlock()

	// Create a new book
	newBook := Book{
		ID:        nextBookID,
		Title:     input.Title,
		Author:    input.Author,
		CreatedAt: time.Now(),
	}

	// Store in the map
	books[nextBookID] = newBook
	nextBookID++ // Increment ID for the next book

	log.Printf("New book created: %+v", newBook)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": newBook})
}

// PUT /books/:id - Update a book
func updateBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid book ID"})
	}

	input := new(UpdateBookInput)
	// Parse & Validate Body
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
	}
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrorResponse(err))
	}

	bookMutex.Lock()
	defer bookMutex.Unlock()

	// Check if the book exists
	book, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Book with ID %d not found", id)})
	}

	// Update fields if provided in the input (using pointers in UpdateBookInput)
	if input.Title != nil {
		book.Title = *input.Title
	}
	if input.Author != nil {
		book.Author = *input.Author
	}

	// Save the updated book back to the map
	books[id] = book

	log.Printf("Book updated: %+v", book)
	return c.JSON(fiber.Map{"status": "success", "data": book})
}

// DELETE /books/:id - Delete a book
func deleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid book ID"})
	}

	bookMutex.Lock()
	defer bookMutex.Unlock()

	// Check if the book exists before deleting
	_, exists := books[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("Book with ID %d not found", id)})
	}

	// Delete the book from the map
	delete(books, id)

	log.Printf("Book with ID %d deleted", id)
	// Send 204 No Content for successful DELETE
	return c.SendStatus(fiber.StatusNoContent)
}

// --- Setup & Main ---
func setupRoutes(app *fiber.App) {
	// Group for book API
	bookApi := app.Group("/books")

	bookApi.Get("/", getBooks)       // GET /books
	bookApi.Post("/", createBook)    // POST /books
	bookApi.Get("/:id", getBook)     // GET /books/:id
	bookApi.Put("/:id", updateBook)  // PUT /books/:id
	bookApi.Delete("/:id", deleteBook) // DELETE /books/:id
}

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second, // Example for graceful shutdown
	})

	// Middleware
	app.Use(recover.New()) // Catch panics
	app.Use(logger.New())  // Log requests

	// Setup routes
	setupRoutes(app)

	// Handle 404 for unmatched routes
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": fmt.Sprintf("Endpoint '%s' with method '%s' not found.", c.Path(), c.Method()),
		})
	})

	// --- Graceful Shutdown Logic ---
	go func() {
		listenAddr := ":3000"
		log.Printf("Server starting to listen on %s", listenAddr)
		if err := app.Listen(listenAddr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to run server: %v", err)
		}
		log.Println("Server stopped listening.")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	receivedSignal := <-quit
	log.Printf("Received shutdown signal: %s. Starting graceful shutdown...", receivedSignal)

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}

	log.Println("Graceful shutdown complete. Server exiting.")
}
```

**How to Run and Test:**

1.  Save the code above as `main.go`.
2.  Run: `go run main.go`.
3.  Use `curl` or Postman/Insomnia to test the endpoints:
	*   **Create:** `curl -X POST http://localhost:3000/books -H "Content-Type: application/json" -d '{"title": "Go Programming Blueprints", "author": "Mat Ryer"}'`
	*   **Create (Invalid):** `curl -X POST http://localhost:3000/books -H "Content-Type: application/json" -d '{"title": "Go"}'` (Author required)
	*   **Get All:** `curl http://localhost:3000/books`
	*   **Get One:** `curl http://localhost:3000/books/1` (Replace 1 with an existing ID)
	*   **Update:** `curl -X PUT http://localhost:3000/books/1 -H "Content-Type: application/json" -d '{"author": "M. Ryer"}'` (Updates only the author)
	*   **Delete:** `curl -X DELETE http://localhost:3000/books/1`
	*   **Get Deleted:** `curl http://localhost:3000/books/1` (Should return 404 Not Found)
	*   **Not Found Route:** `curl http://localhost:3000/nonexistent` (Should return 404 JSON from the middleware)

This example demonstrates CRUD basics, routing, body parsing, input validation, using a map with a mutex for simple storage, and basic error handling.

---

## 8. API Documentation 📚

Good API documentation is crucial so that consumers (frontend developers, other teams, public users) can understand how to use your API.

*   **Go Code Documentation (GoDoc):**
	*   Always add comments to your public functions, structs, interfaces, and packages following the GoDoc format.
	*   Use `go doc ./...` to view documentation locally.
	*   Publish to [pkg.go.dev](https://pkg.go.dev/) for public access.
	*   This is important for other developers using your library/code.

*   **RESTful API Documentation (Specification):**
	*   Document each API endpoint:
		*   Path (e.g., `/users/:id`)
		*   HTTP Method (GET, POST, PUT, DELETE)
		*   Brief description of the endpoint's purpose.
		*   Parameters (path, query, header) - name, type, required/optional, description.
		*   Request Body (if any) - format (JSON, XML), example, field descriptions, validation rules.
		*   Success Response - status code, body format, example, field descriptions.
		*   Error Responses - status code, body format, possible error codes, descriptions.
		*   Required Authentication/Authorization.
	*   **Tools:**
		*   **OpenAPI Specification (Swagger):** The industry standard for describing RESTful APIs in JSON or YAML format.
			*   Many tools can generate interactive UIs from OpenAPI specs (Swagger UI, Redoc).
			*   Can be written manually or generated from code/comments.
			*   Fiber has integrations (unofficial but popular) like `github.com/arsmn/fiber-swagger/v3` that can generate specs from comments or structs and serve Swagger UI.
			*   Example comments for `fiber-swagger`:
				```go
				// @Summary      Get a book by ID
				// @Description  Get details for a specific book
				// @Tags         Books
				// @Accept       json
				// @Produce      json
				// @Param        id   path      int  true  "Book ID"
				// @Success      200  {object}  BookResponseDoc // Define separate doc struct if needed
				// @Failure      400  {object}  ErrorResponseDoc
				// @Failure      404  {object}  ErrorResponseDoc
				// @Router       /books/{id} [get]
				func getBook(c *fiber.Ctx) error { ... }
				```
		*   **API Blueprint:** Another Markdown-based format for API documentation.
		*   **Postman Collections:** Can be exported as basic documentation.

Choose a format and tool that works for you, but **ensure your API documentation is always kept up-to-date** as the code changes.

---

## 9. Best Practices ✨

Some tips and best practices when developing applications with Go Fiber:

1.  **Clear Project Structure:** Use a logical directory structure (feature-based or layer-based) as the project grows.
2.  **Centralized Configuration:** Load configuration (port, DB DSN, secret keys) from environment variables or config files, don't hardcode. Use libraries like `viper` or structs with `envconfig`.
3.  **Consistent Error Handling:** Use a custom `ErrorHandler` to format all error responses uniformly and hide internal details in production. Log errors verbosely on the server.
4.  **Thorough Input Validation:** Always validate *all* client input (body, query, params, headers) using libraries like `go-playground/validator`.
5.  **Use Middleware Wisely:** Leverage middleware for cross-cutting concerns (logging, auth, recovery, CORS). Apply them at the appropriate scope (global, group, route).
6.  **Dependency Injection (DI):** Use DI to manage dependencies (like DB connections) for more testable and structured code. Avoid global variables where possible.
7.  **Write Tests:** Write unit tests for business logic/services/repositories and integration tests for your HTTP handlers. Aim for good code coverage.
8.  **Security:**
	*   Use HTTPS in production (SSL termination at the reverse proxy).
	*   Protect against common attacks (SQLi, XSS, CSRF - Fiber doesn't automatically protect against CSRF for session/HTML form apps; CSRF middleware is needed).
	*   Don't expose sensitive information in logs or error responses.
	*   Keep secret keys (JWT, API keys) confidential.
	*   Set `BodyLimit` to prevent DoS.
	*   Use strong Authentication & Authorization.
9.  **Effective Logging:** Use structured loggers (e.g., `zerolog`, `zap`) in addition to Fiber's logger. Log relevant context (request ID, user ID if available) for easier debugging. Set different log levels for development and production.
10. **Graceful Shutdown:** Implement graceful shutdown to handle server termination properly.
11. **Performance:** Profile your application to find bottlenecks before optimizing. Utilize caching where appropriate.
12. **Documentation:** Write GoDoc for your code and document your API clearly.
13. **Dependency Management:** Use Go Modules (`go.mod`, `go.sum`) to manage dependencies. Update dependencies periodically (watch for breaking changes).
14. **Use Context (`context.Context`):** Pass `c.Context()` (which is a `context.Context`) down to long-running or I/O-bound function calls (like DB queries, outgoing HTTP requests) to support cancellation and timeouts.

---

## 10. Contributing 🤝

Fiber is an open-source project. Contributions are always welcome! Check the official contribution guidelines in the Fiber repository: [CONTRIBUTING.md](https://github.com/gofiber/fiber/blob/master/.github/CONTRIBUTING.md).

Ways to contribute:

*   Reporting bugs.
*   Submitting feature requests.
*   Writing or improving documentation.
*   Submitting Pull Requests with bug fixes or new features (ensure you follow code style and include tests).
*   Helping answer questions in GitHub Discussions or other community channels.

---

## 11. License 📜

Go Fiber is released under the **MIT License**. See the [LICENSE](https://github.com/gofiber/fiber/blob/master/LICENSE) file for full details.

---

## 12. Acknowledgements 🙏

*   The Go Fiber core team and all its contributors.
*   The creator of Fasthttp, Valyala.
*   The Express.js community for the inspiration.
*   The incredible Go community.

---

Hopefully, this comprehensive guide is helpful in your journey learning and using Go Fiber v2! Happy coding! 🎉

