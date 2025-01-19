Installer GO
$ sudo snap install go

Vérifier que la bonne version de Go est installer
$ go version
Version de Go utilié pour le projet: 1.23.4

$ sudo pacman -S go

Installer les dependances
$go mod tidy


Pour lancer le server
$go run main.go

Pour build un exécutable
$go build main.go

Pour run l'exécutable
$./main






Docker


sudo docker run -p 3000:3000 --rm -v $(pwd):/app -v /app/tmp --name my-go-server-air my-go-server

