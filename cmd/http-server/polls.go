package main

import (
	"fmt"
	"net/http"
	"strings"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
)

func (app *App) handleCreatePoll(rw http.ResponseWriter, r *http.Request) {
	var arg sqlc.CreatePollsParams
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	app.PollManager.Create(r.Context(), arg)

	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, err, "")
		return
	}
	sendResponse(rw, http.StatusCreated, arg, "poll created")
}

func (app *App) handleGetPoll(rw http.ResponseWriter, r *http.Request) {

	entry := r.Context().Value(Pools{})

	fmt.Println("entry_bool", entry)

	if entry == true {

		polls := app.PollManager.Get(r.Context())
		sendResponse(rw, http.StatusCreated, polls, "polls_section")
		return
	}

	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")

}

func (app *App) handleUploadImage(rw http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(10 << 20)

	// get handler for filename, size and headers
	file, handler, err := r.FormFile("image")
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid name/ file missing")
		return
	}
	defer file.Close()

	// allow only png
	if !(strings.Contains(handler.Header.Get("Content-Type"), "png")) {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid image format, only .png allowed")
		return
	}

	// putting image in image store
	fileName := id + ".png"
	s3ImageUrl, err := app.objectStore.Put(file, fileName)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

}