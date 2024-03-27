package routes

import (
	"net/http"

	"github.com/sayedulkrm/go-curd-todo-2/middlewares"
)

func SetupRoutes() http.Handler {
	root := http.NewServeMux()

	// Set up user routes
	tasksRoutes := TaskRoutes()
	root.Handle("/api/v1/", middlewares.LogMiddleware(http.StripPrefix("/api/v1", tasksRoutes).(http.HandlerFunc)))

	// Custom NotFoundHandler
	// root.Handle("/*", http.HandlerFunc(utils.NotFoundResponse))

	root.HandleFunc("/", middlewares.LogMiddleware(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")

		html := `<h1>Server is working. To See Frontend <a href="http://localhost:3000"> Click Here </a></h1>`
		w.Write([]byte(html))
	}))

	// Catching Pages Not Found

	return root
}
