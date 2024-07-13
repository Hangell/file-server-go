package main

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"log"
	"net/http"
	"os"
)

/*
* @User admin
* @Pass teste
*/
func Secret(user, realm string) string {
	if user == "admin" {
		return "$1$/q9tTBog$KtdXn0eQTpiUP9g//xLaS1" //https://unix4lyfe.org/crypt/ MD5 Crypt: md5 salt
	}
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretorio> <porta>")
		os.Exit(1)
	}

	httpDir := os.Args[1] //position 0 é o nome do arquivo
	port := os.Args[2]

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))
	fmt.Printf("SErvidor na porta %s...", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

