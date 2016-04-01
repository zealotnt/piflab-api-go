package lib

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	router *mux.Router
	DB     *DB
	ENV    string
	PORT   string
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}

func (app *App) Close() {
	app.DB.Close()
}

func (app *App) AddRoutes(routes Routes) {
	for _, route := range routes {
		handler := Logger(route.Handler(app))

		app.router.
			// Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)
	}
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":"+app.PORT, app))
}

func (app *App) Request(method string, route string, body string) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, route, strings.NewReader(body))
	request.RemoteAddr = "127.0.0.1:8080"

	if method == "POST" || method == "PUT" {
		if body != "" && body[0:1] == "{" {
			request.Header.Set("Content-Type", "application/json")
		} else {
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		}
	}

	response := httptest.NewRecorder()
	app.ServeHTTP(response, request)

	return response
}

func NewApp() *App {
	return &App{
		router: newRouter(),
		DB:     newDB(),
		PORT:   getPort(),
		ENV:    getEnv(),
	}
}

func getPort() string {
	env := os.Getenv("PORT")

	if env == "" {
		return "8080"
	}

	return env
}

func getEnv() string {
	env := os.Getenv("ENV")

	if env == "" {
		return "production"
	}

	return env
}

func newRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func parseDB(driver string, env string) string {
	switch driver {
	case "postgres":
		return env
	case "mysql":
		re, _ := regexp.Compile(`(\w*):\/\/(.+)`)
		result := re.FindStringSubmatch(env)
		return result[2] + "?parseTime=true"
	default:
		panic("Driver is not supported")
	}
}

func newDB() *DB {
	re, _ := regexp.Compile(`(\w*):\/\/`)
	result := re.FindStringSubmatch(os.Getenv("DATABASE_URL"))

	if result == nil {
		panic("Can't find driver for DB")
	}

	driver := result[1]

	gorm_arg := parseDB(driver, os.Getenv("DATABASE_URL"))

	db, err := gorm.Open(driver, gorm_arg)
	if err != nil {
		panic(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.DB().Ping()
	if err != nil {
		panic(err.Error())
	}

	return &DB{db.Debug()}
}
