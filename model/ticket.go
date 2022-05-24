package model

type Ticket struct {
	// ID est l'identifiant unique du ticket dans la BDD
	ID string `json:"id"`
	// Title est le titre du ticket
	Title string `json:"title"`
	// Description est la description du ticket qui doit aider le user à résoudre le pb
	Description string `json:"description"`
	// Status est le status du ticket qui peut être à : Ouvert, En cours, Terminer
	Status string `json:"status"`
	// IdUser est l'id du user qui est en charge du ticket
	IdUser string `json:"id_user"`
}
