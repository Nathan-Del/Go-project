package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/estiam/ticketing/auth"
	"github.com/estiam/ticketing/model"
	"github.com/gin-gonic/gin"
)

// GetAllTicket donnes en JSON la liste des tickets
func (s *Service) GetAllTicket(ctx *gin.Context) {
	us, err := s.db.GetAllTicket()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		log.Println("get")
		return
	}
	ctx.JSON(http.StatusOK, us)
}

// CreateTicket est un handler qui permet de créer un nouveau ticket
func (s *Service) CreateTicket(ctx *gin.Context) {
	var payload model.Ticket
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create ticket", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.CreateTicket(&payload)
	if err != nil {
		log.Printf("error create ticket %T %#v", err, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// DeleteTicket permet de supprimer un ticket avec son id passer en paramètre
func (s *Service) DeleteTicket(ctx *gin.Context) {
	err := s.db.DeleteUser(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"ticket": "deleted"})
}

// UpdateTicket met à jour un ticket avec son id passer en paramètre
func (s *Service) UpdateTicket(ctx *gin.Context) {

	payload := make(map[string]interface{})
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create ticket", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Une erreur sera renvoyé si le status est différent de Ouvert, En cours, Terminer, afin de garantir une cohérence des données
	switch payload["status"] {
	case payload["status"] != "Ouvert":
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	case payload["status"] != "En cours":
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	case payload["status"] != "Terminer":
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	u, err := s.db.UpdateTicket(ctx.Param("id"), payload)
	if err != nil {
		log.Println("error create ticket", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// GetTicketById récupère un ticket à partir du l'id donné en paramètre
func (s *Service) GetTicketById(ctx *gin.Context) {
	u, err := s.db.GetTicketById(ctx.Param("id"))
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	sessionValue, ok := ctx.MustGet("session").(*auth.CustomClaims)
	if !ok {
		log.Println("error assert session")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Println("from session get id user", sessionValue.IDUser)

	ctx.JSON(http.StatusOK, u)
}

// GetTicketByUserId permet à un utilisateur de retrouver les tickets qui lui son affecter
func (s *Service) GetTicketsByUserId(ctx *gin.Context) {
	u, err := s.db.GetTicketsByUserId(ctx.Param("id"))
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	sessionValue, ok := ctx.MustGet("session").(*auth.CustomClaims)
	if !ok {
		log.Println("error assert session")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Println("from session get id user", sessionValue.IDUser)

	ctx.JSON(http.StatusOK, u)
}

// UploadFile permet de charger un fichier dans le dossier /upload
func (s *Service) UploadFile(ctx *gin.Context) {

	// Single file
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Println("handler: upload image error", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Println(file.Filename)

	// vérifie que le fichier est bien un png
	if file.Filename[len(file.Filename)-3:] != "png" {
		log.Println("handler: upload image error, is not a png", err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.JSON(http.StatusMethodNotAllowed, fmt.Sprintf("'%s' is not a png!", file.Filename))
		return
	}

	// Upload the file to ./upload/ + file.Filename
	ctx.SaveUploadedFile(file, "./upload/"+file.Filename)

	ctx.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
