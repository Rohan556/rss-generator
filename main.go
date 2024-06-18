package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rohan556/rss-generator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Port not found")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB URL not found")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", errorHandler)
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feed", apiCfg.handlerGetAllFeeds)
	v1Router.Post("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follow", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follow/{feed_follow_id}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting at port %v", portString)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port number: ", portString)
}
