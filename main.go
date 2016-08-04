package main

import (
	"fmt"
	"html/template"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"time"
)

type editData struct {
	Width  int
	Height int
	Key    string
}

func main() {
	imageStore := getImageDirectory()

	server := newImageServer(imageStore)
	http.HandleFunc("/", server.handleMain)
	http.HandleFunc("/edit", server.handleEdit)
	http.HandleFunc("/image", server.handleImage)

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	log.Print("Listening on port " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type imageServer struct {
	dir          string
	keyValidator *regexp.Regexp
	editTemplate *template.Template
	mainTemplate *template.Template
}

func newImageServer(dir string) *imageServer {
	keyValidator, _ := regexp.Compile("^[0-9]+$")

	return &imageServer{
		dir:          dir,
		keyValidator: keyValidator,
		editTemplate: template.Must(template.New("edit").Parse(string(MustAsset("static/edit.html")))),
		mainTemplate: template.Must(template.New("edit").Parse(string(MustAsset("static/main.html")))),
	}
}

func (s *imageServer) handleMain(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("image")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		key := time.Now().Format("20060102150405")
		destPath := s.getImagePath(key)
		f, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			panic(err)
		}

		http.Redirect(w, r, "edit?key="+key, http.StatusTemporaryRedirect)
		return
	}

	w.Header().Set("content-type", "text/html")
	err := s.mainTemplate.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func (s *imageServer) handleEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	}
	key := r.URL.Query().Get("key")
	if !s.keyValidator.MatchString(key) {
		http.Error(w, "Bad key", http.StatusBadRequest)
		return
	}
	f, err := os.Open(s.getImagePath(key))
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
		} else {
			http.Error(w, "Failed to open file: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		http.Error(w, "Failed to process image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/html")
	err = s.editTemplate.Execute(w, editData{img.Bounds().Dx(), img.Bounds().Dy(), key})
	if err != nil {
		panic(err)
	}
}

func (s *imageServer) handleImage(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if !s.keyValidator.MatchString(key) {
		http.Error(w, "Bad key", http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, s.getImagePath(key))
}

func (s *imageServer) getImagePath(key string) string {
	return path.Join(s.dir, key+".jpg")
}

func getImageDirectory() string {
	tmp := os.Getenv("TMPDIR")
	if tmp == "" {
		tmp = "/tmp"
	}
	imageStore := path.Join(tmp, "imagestore")
	if _, err := os.Stat(imageStore); err != nil {
		if err = os.MkdirAll(imageStore, 0700); err != nil {
			panic(err)
		}
	}
	return imageStore
}
