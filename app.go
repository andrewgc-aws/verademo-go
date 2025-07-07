package main

import (
	"embed"
	_ "github.com/boombuler/barcode"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gorilla/securecookie"
	_ "github.com/ncruces/julianday"
	_ "github.com/prometheus/client_golang/prometheus" // use a real subpackage
	_ "github.com/stretchr/testify"
	_ "github.com/tetratelabs/wazero"
	_ "golang.org/x/crypto/bcrypt" // use a real subpackage
	_ "gopkg.in/yaml.v2"
	"log"
	"net/http"
	router "verademo-go/src-app/routes"
	sqlite "verademo-go/src-app/shared/db"
	session "verademo-go/src-app/shared/session"
	"verademo-go/src-app/shared/view"
)

//go:embed resources
var resources embed.FS

//go:embed templates
var templates embed.FS

func main() {
	session.Configure(session.Session{Name: "verademo", SecretKey: "key"})
	sqlite.OpenDB()
	log.Print("\nStarting VerademoGO....")
	log.Print("\nVerademoGO is running.")
	view.ParseTemplates(templates)
	router.SetResources(resources)
	log.Fatal(http.ListenAndServe(":8000", router.Routes()))
}
