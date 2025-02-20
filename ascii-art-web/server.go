package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	U "ascii-art-web/utils"
)

var _map = make(map[int][8]string)

func InitMap() {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 32; i < 127; i++ {
		_map[i] = U.InsertValue(scanner)
	}
}

func GenerateASCIIArt(text string) string {
	var result strings.Builder
	lines := [8]string{}

	for _, char := range text {
		for i := 0; i < 8; i++ {
			lines[i] += _map[int(char)][i]
		}
	}
	for i := 0; i < 8; i++ {
		result.WriteString(lines[i] + "\n")
	}
	return result.String()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	text := r.FormValue("text")

	if text == "" {
		http.Error(w, "Text input cannot be empty", http.StatusBadRequest)
		return
	}

	asciiArt := GenerateASCIIArt(text)
	fmt.Println(asciiArt)
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, asciiArt)
}

func main() {
	InitMap()

	// Serve static files (CSS) from the "templates" folder
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8081", nil)
}
