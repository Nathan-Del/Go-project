// main est responsable de la récupération des envois de configuration
// et de créer tous les objets nécessaires (db type of conn...).
package main

import (
	"github.com/gin-gonic/gin"

	"github.com/estiam/ticketing/auth"
	"github.com/estiam/ticketing/db/postgre"
	"github.com/estiam/ticketing/service"
)

func main() {
	dbConnection := postgre.New("mydb.db")
	s := service.New(dbConnection)

	router := gin.Default()

	verif := auth.VerifyJWT()

	router.GET("/users", s.GetAllUser)
	router.POST("/users", s.CreateUser)
	router.GET("/users/:id", verif, s.GetUserById)
	router.DELETE("/users/:id", s.DeleteUser)
	router.PATCH("/users/:id", s.UpdateUser)
	router.PUT("/users/:id", s.UpdateUser)
	router.POST("/login", s.Login)

	router.GET("/tickets", s.GetAllTicket)
	router.POST("/tickets", s.CreateTicket)
	router.GET("/tickets/:id", verif, s.GetTicketById)
	router.GET("/tickets_users/:id", verif, s.GetTicketsByUserId)
	router.DELETE("/tickets/:id", s.DeleteTicket)
	router.PATCH("/tickets/:id", s.UpdateTicket)
	router.PUT("/tickets/:id", s.UpdateTicket)

	router.POST("upload", s.UploadFile)

	router.Run()
}
