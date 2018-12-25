package app

import "net/http"

// App struct is top level application.
// It contains Router,View,Config and private fields.
type App struct {
	router *Router `register:"Router"`
	// routerC map[string]*routerCache
	config *Config `register:"Config"`
}

// create app instance【router and config】
func New() *App {
	a := new(App)
	a.router = NewRouter()
	// a.config, _ = NewConfig("config.json")
	return a
}

// HTTP interface implementation
func (app *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	app.handle(res, req)
}

// 处理程序
func (app *App) handle(res http.ResponseWriter, req *http.Request) {
	// 应用处理的地方
}

func (app *App) method(method string, fn ...Handler) {

}

// 运行 监听端口
func (app *App) Run() {
	// addr := app.config.StringOr("app.server", "localhost:9001")
	// println("http server run at " + addr)
	// e := http.ListenAndServe(addr, app)
	// panic(e)
}
