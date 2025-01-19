# Installer GO
$ sudo snap install go
$ sudo pacman -S install
$ sudo apt install golang


Vérifier que la bonne version de Go est installer
$ go version
Version de Go utilié pour le projet: 1.23.4

# HOST-LOCAL

## Installer les dependances
$ go mod tidy

## Pour lancer le server
$ go run main.go

## Pour build un exécutable
$ go build main.go

## Pour run l'exécutable
$./main



# DOCKER
Run in docker with hot reload

## INSTALL
$ sudo snap install docker

## Build Docker image
$ sudo docker buildx build -t go-server .

## Run server in Docker
$ sudo docker run -p 3000:3000 --rm -v $(pwd):/app -v /app/tmp --name go-server-air go-server

