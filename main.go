package main

import (
	"code-breaker/config"
	"code-breaker/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

func main() {
	config.Init()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Handle(
		os.Getenv("ASSETS_PATH")+"*",
		http.StripPrefix(os.Getenv("ASSETS_PATH"), http.FileServer(http.Dir("assets"))),
	)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		views.RenderTemplate(w, "index.html", map[string]interface{}{
			"Name": "Andy",
		})
	})

	r.Post("/message", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		message := r.Form.Get("message")
		if message == "" {
			log.Println("Message empty")
			return
		}

		letters := LettersToRandomSymbols()
		words := getWordsAsSlices(strings.Split(message, " "), letters)

		views.RenderTemplate(w, "code.html", map[string]interface{}{
			"message": []rune(message),
			"letters": letters,
			"words":   words,
		})
	})

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		return
	}
}

func getWordsAsSlices(messageWords []string, symbols map[string]string) [][]string {
	var words [][]string

	for _, word := range messageWords {
		var letters []string
		for _, letter := range word {
			letters = append(letters, symbols[string(unicode.ToUpper(letter))])
		}
		words = append(words, letters)
	}

	return words
}
