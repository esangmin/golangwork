package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	"work/trace"

	"github.com/stretchr/objx"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

var avatars Avatar = TryAvatars{
	UseFileSystemAvatar,
	UseAuthAvatar,
	UseGravatar}

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
}

var host = flag.String("host", ":8080", "The addr of the application.")

func main() {

	flag.Parse() // 플래그 파싱

	// gomniauth 설정
	gomniauth.SetSecurityKey("032427rehtuadowtsierusaeratym")
	gomniauth.WithProviders(
		google.New("821325730969-5hb4kqmgtelq1dhidpaubrmb5ls3m6ud.apps.googleusercontent.com", "EaobphGJG0E_tJjl-NQOvvyo", "http://localhost:8080/auth/callback/google"),
		github.New("1da2eaf7d2c43cbed253", "2fcf8c0f720b62a66770ca4474aeed8400f2df2d", "http://localhost:8080/auth/callback/github"),
		facebook.New("2723504014379643", "b1c78df92a1a4f2db101b9f968aa248c", "http://localhost:8080/auth/callback/facebook"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.Handle("/upload", &templateHandler{filename: "upload.html"})
	http.HandleFunc("/uploader", uploaderHandler)
	http.Handle("/avatars/", http.StripPrefix("/avatars/", http.FileServer(http.Dir("./avatars"))))

	// 방을 가져옴
	go r.run()

	// 웹 서버 시작
	log.Println("Starting web server on", *host)
	if err := http.ListenAndServe(*host, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
