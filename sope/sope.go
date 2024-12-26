package sope

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sde-kiran-sadvilkar/sope/render"
	"github.com/sde-kiran-sadvilkar/sope/session"
)

const version = "1.0.0"

type Sope struct {
	AppName string
	Debug   bool
	Version string
	ErrorLog *log.Logger
	InfoLog *log.Logger
	RootPath string
	Routes *chi.Mux
	Render *render.Render
	JetViews *jet.Set
	Session *scs.SessionManager
	config config
	
}

type config struct{
	port string
	renderer string
	cookie cookieConfig
	sessionType string
}

func (s *Sope) New(rootPath string) error {

	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"controllers", "migrations", "views", "data", "public", "tmp", "logs", "middleware",
		},
	}

	err := s.Init(pathConfig)

	if err != nil {
		return err
	}

	// check .env
	err = s.checkDotEnv(rootPath)

	if err!=nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath+"/.env")

	if err!=nil {
		return err
	}

	infogLog, errorLog := s.startLogger()
	s.InfoLog = infogLog
	s.ErrorLog = errorLog

	s.Debug,_ = strconv.ParseBool(os.Getenv("DEBUG"))
	s.Version = version
	s.RootPath = rootPath

	
	s.config = config{
		port: os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
		cookie: cookieConfig{
			name: os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist: os.Getenv("COOKIE_PERSISTS"),
			secure: os.Getenv("COOKIE_SECURE"),
			domain: os.Getenv("COOKIE_DOMAIN"),

		},
		sessionType: os.Getenv("SESSION_TYPE"),
	}

	sess := session.Session{
		CookieLifeTime: s.config.cookie.lifetime,
		CookiePersist: s.config.cookie.persist,
		CookieName: s.config.cookie.name,
		SessionType: s.config.sessionType,
		CookieDomain: s.config.cookie.domain,
	}

	s.Session = sess.InitSession();

	var views = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views",rootPath)),
		jet.InDevelopmentMode(),
	)

	s.JetViews = views

	s.createRenderer()

	s.Routes = s.routes().(*chi.Mux)


	return nil

}

func (s *Sope) CreateServer(){
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", s.config.port),
		ErrorLog: s.ErrorLog,
		Handler: s.Routes,
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	s.InfoLog.Printf("Listening on port %s", s.config.port)
	err := srv.ListenAndServe()
	s.ErrorLog.Fatal(err)
}

func (s *Sope) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := s.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sope) checkDotEnv(path string) error {
	err := s.CreateFileIfNotExist(fmt.Sprintf("%s/.env",path))

	if err != nil {
		return err
	}

	return nil
}

func (s *Sope) startLogger() (*log.Logger, *log.Logger){

	infoColor:="\033[33m"
	errorColor:="\033[31m"
	restColor:="\033[0m"

	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, infoColor + " INFO \t" + restColor, log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stdout,errorColor + "ERROR\t" + restColor,log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog

}

func (s *Sope) createRenderer() {

	render := render.Render{
		Renderer: s.config.renderer,
		RootPath: s.RootPath,
		Port: s.config.port,
		JetViews: s.JetViews,
	}

	s.Render = &render

}