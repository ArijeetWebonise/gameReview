package app

import (
	"net/http"

	"github.com/Arijeet-webonise/gameReview/pkg/framework"
)

// Handle will be serving only those requests that dont need to be authed
func (app *App) handle(handler func(*framework.Response, *framework.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		handler(&res, &req)
		res.Write()
	})
}

//RenderView renders a view
func (app *App) renderView(viewHandler func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewHandler(w, r)
	})
}

func (app *App) renderSecureView(viewHandler func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := app.CheckIfLogined(r)
		if err != nil {
			app.Log.Error(err)
			app.Handler404Error(w, r)
			return
		}
		viewHandler(w, r)
	})
}
