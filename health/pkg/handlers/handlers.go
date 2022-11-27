package handlers

import (
	"net/http"

	"github.com/fontexd/go/health/pkg/checks"
	"github.com/fontexd/go/health/pkg/config"
	"github.com/fontexd/go/health/pkg/jsoncrawler"
	"github.com/fontexd/go/health/pkg/models"
	"github.com/fontexd/go/health/pkg/render"
)

// Repo repsitory used by the handlers
var Repo *Repository

// Repsotity type defining
type Repository struct {
	App *config.Appconfig
}

// Newrepo creates the repo
func NewRepo(a *config.Appconfig) *Repository {

	return &Repository{

		App: a,
	}
}

// Newhandler set the respository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	Urls := []string{"https://oauth.plan2learn.dk/health", "https://media.plan2learn.dk/health"}
	for x := range Urls {

		remoteip := jsoncrawler.JsonData(Urls[x]).Value
		m.App.Session.Put(r.Context(), "remote_ip", remoteip)
	}
	render.RenderTemplate(w, "home.page.tmpl", &models.Templatedata{})

}

func (m *Repository) Environments(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "environments.page.tmpl", &models.Templatedata{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	env := "192.168.10.200"
	stringMap := make(map[string]string)
	stringBool := make(map[string]bool)
	stringMap["test"] = "hello, again"
	stringBool["redis"] = checks.RedisHealth(env)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.Templatedata{
		StringMap:  stringMap,
		StringBool: stringBool,
	})

}
