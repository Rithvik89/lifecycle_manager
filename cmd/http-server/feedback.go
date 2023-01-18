package main

import (
	"net/http"
)

func (app *App) handleGetFeedback(rw http.ResponseWriter, r *http.Request) {

	x := (r.Context().Value("user")).(UserDetails)

	if x.Feedback {

		feedback := app.FeedbackManager.GetFeedback(r.Context())
		sendResponse(rw, http.StatusCreated, feedback, "feedback_section")
		return
	}

	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")
}
