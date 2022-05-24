package service

import (
	"log"
	"net/http"

	"github.com/estiam/ticketing/auth"
	"github.com/estiam/ticketing/model"
	"github.com/gin-gonic/gin"
)

// GetAllUser donnes en JSON la liste des users
func (s *Service) GetAllUser(ctx *gin.Context) {
	us, err := s.db.GetAllUser()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, us)
}

// CreateUser est un handler qui permet de créer un nouvel user
func (s *Service) CreateUser(ctx *gin.Context) {
	var payload model.User
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.CreateUser(&payload)
	if err != nil {
		log.Printf("error create user %T %#v", err, err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

// DeleteUser permet de supprimer un user avec son id passer en paramètre
func (s *Service) DeleteUser(ctx *gin.Context) {
	err := s.db.DeleteUser(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"user": "deleted"})
}

// UpdateUser met à jour un user avec son id passer en paramètre
func (s *Service) UpdateUser(ctx *gin.Context) {
	payload := make(map[string]interface{})
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.UpdateUser(ctx.Param("id"), payload)
	if err != nil {
		log.Println("error create user", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// GetUserById récupère un user à partir du l'id donné en paramètre
func (s *Service) GetUserById(ctx *gin.Context) {
	u, err := s.db.GetUserById(ctx.Param("id"))
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

// Login fournit un jwt après que le user ce soit connecter avec son  email, password
func (s *Service) Login(ctx *gin.Context) {
	var payload model.LoginPayload
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println("service user: login user parse payload", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.GetUserByEmail(payload.Email)
	if err != nil {
		log.Println("servce user: login user db GetUserByEmail", err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if u.Password != payload.Password {
		log.Println("service user: the password doesn't match")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtValue, err := auth.NewJWT(u)
	if u.Password != payload.Password {
		log.Println("service user: create JWT", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
}
