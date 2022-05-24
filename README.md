# Ticketing app

Ce projet consiste en la création d'une API en GO. 

L'API est structuré pour recevoir des users et des tickets.


User :

	• ID : string => ID est l'identifiant unique du user dans la BDD

	• FirstName : string => FirstName est le prénom du user

	• LastName : string => LastName est le nom de famille du user

	• Email : string => Email est utilisé pour la connexion et la communication avec le user

	• Password : Password => Password est haché en sha256 du user
	

Tickets :

	• ID string `json:"id"` => ID est l'identifiant unique du ticket dans la BDD

	• Title string `json:"title"` => Title est le titre du ticket

	• Description string `json:"description"` => Description est la description du ticket qui doit aider le user à résoudre le pb

	• Status string `json:"status"` => Status est le status du ticket qui peut être à : Ouvert, En cours, Terminer

	• IdUser string `json:"id_user"` => IdUser est l'id du user qui est en charge du ticket


En effet, le but de ce projet est de créer une API qui permettra la gestion de ticket, comme le ferait une application de ticketing. Pour ce faire, les routes suivantes ont été mises en place :

Users :

	• GET => http://localhost:8080/users

	• POST => http://localhost:8080/users

	• GET => http://localhost:8080/users/:id

	• DELETE => http://localhost:8080/users/:id

	• PATCH => http://localhost:8080/users/:id

	• PUT => http://localhost:8080/users/:id

	• POST => http://localhost:8080/login

Tickets :

	• GET => http://localhost:8080/tickets

	• POST => http://localhost:8080/tickets

	• GET => http://localhost:8080/tickets/:id

	• GET => http://localhost:8080/tickets_users/:id

	• DELETE => http://localhost:8080/tickets/:id

	• PATCH => http://localhost:8080/tickets/:id

	• PUT => http://localhost:8080/tickets/:id

    • POST => http://localhost:8080/upload


## Installation GO 

Pour exécuter ce projet, il est nécessaire d'avoir GO d'installer.
Si ce n'est pas encore le cas voici un lien qui vous explique comment faire : 

https://go.dev/doc/install


## Excécution du projet

Le projet fonctionne avec des conteneurs créés via docker, pour les lancer, vous pouvez exécuter le Makerfile.
Mais si vous êtes sur Windows, il faut lancer la commande dans le Makerfile manuellement dans le terminal :

```
docker build -t ticketing-api .
```

Ensuite, il faut lancer la commande suivante pour lancer les conteneurs de notre image :

```
docker-compose up -d
```

Il est possible de voir la liste des conteneurs avec :

```
docker ps -a
```

## Arrêt du projet

Pour stopper le projet, vous pouvez lancer la commande suivante (CONTAINER ID => l'id des conteneurs lancer visible avec la commande ```docker ps -a```) :

```
docker stop <CONTAINER ID> <CONTAINER ID>
```

Puis on peut supprimer les conteneurs une fois arrêter avec :

```
docker-compose down
```