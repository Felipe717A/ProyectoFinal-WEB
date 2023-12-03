package main

import (
	"log"
	"net/http"

	"PROYECTO/controllers"
	"PROYECTO/handlers"
	"PROYECTO/models"
	repository "PROYECTO/repository"

	GorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func ConnectDB(url, driver string) (*sqlx.DB, error) {
	pgURL, err := pq.ParseURL(url)
	if err != nil {
		log.Printf("Error al analizar la URL de conexión: %s", err.Error())
		return nil, err
	}

	db, err := sqlx.Connect(driver, pgURL)
	if err != nil {
		log.Printf("Error al conectar a la base de datos: %s", err.Error())
		return nil, err
	}

	log.Printf("Conexión exitosa a la base de datos: %#v", db)
	return db, nil
}

func main() {
	db, err := ConnectDB("postgres://lhvextcg:XgumSZeMyraHVY8WXBeXft9NTSnEqVfX@bubble.db.elephantsql.com/lhvextcg", "postgres")
	if err != nil {
		log.Fatalln("Error conectando a la base de datos", err.Error())
		return
	}

	repo, err := repository.NewRepository[models.F1Teams](db)
	if err != nil {
		log.Fatalln("Error al crear una instancia de repositorio", err.Error())
		return
	}

	controller, err := controllers.NewController(repo)
	if err != nil {
		log.Fatalln("Error al crear una instancia de controlador", err.Error())
		return
	}

	handler, err := handlers.NewHandler(controller)
	if err != nil {
		log.Fatalln("Error al crear una instancia de manejador", err.Error())
		return
	}

	router := mux.NewRouter()

	router.Handle("/f1teams", http.HandlerFunc(handler.ListF1Teams)).Methods(http.MethodGet)
	router.Handle("/f1teams", http.HandlerFunc(handler.CreateF1Team)).Methods(http.MethodPost)
	router.Handle("/f1teams/{id}", http.HandlerFunc(handler.ReadF1Team)).Methods(http.MethodGet)
	router.Handle("/f1teams/{id}", http.HandlerFunc(handler.UpdateF1Team)).Methods(http.MethodPatch)
	router.Handle("/f1teams/{id}", http.HandlerFunc(handler.DeleteF1Team)).Methods(http.MethodDelete)
	headers := GorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := GorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"})
	origins := GorillaHandlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":8080", GorillaHandlers.CORS(headers, methods, origins)(router))
}
