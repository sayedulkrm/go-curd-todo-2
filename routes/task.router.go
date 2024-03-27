package routes

import (
	"net/http"

	"github.com/sayedulkrm/go-curd-todo-2/controllers"
)

func TaskRoutes() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("GET /alltasks", controllers.GetAllTasks)
	router.HandleFunc("POST /task/create", controllers.CreateNewTask)
	router.HandleFunc("PUT /task", controllers.UpdateTask)
	router.HandleFunc("DELETE /task", controllers.DeleteTask)

	return router

}
