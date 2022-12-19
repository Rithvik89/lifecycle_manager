package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/internal/feedback_manager"
	"github.com/yards22/lcmanager/internal/poll_manager"
	"github.com/yards22/lcmanager/internal/r_manager"
	"github.com/yards22/lcmanager/internal/r_posts_manager"
	"github.com/yards22/lcmanager/internal/r_users_manager"
	"github.com/yards22/lcmanager/internal/t_posts_manager"
	"github.com/yards22/lcmanager/internal/t_users_manager"
	"github.com/yards22/lcmanager/internal/token_manager"
	"github.com/yards22/lcmanager/pkg/env"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open(env.ViperGetEnvVar("DB_DRIVER_NAME"), env.ViperGetEnvVar("DB_DATA_SOURCE_NAME"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initRunnerManagers(app *App) {
	// Initialize managers and add to app
	tokenManager := token_manager.New(sqlc.New(app.db), time.Hour)
	app.managers["tokenManager"] = tokenManager
	trendingPostsManager := t_posts_manager.New(sqlc.New(app.db), 24*(time.Hour))
	app.managers["trendingPostsManager"] = trendingPostsManager
	trendingUserManager := t_users_manager.New(sqlc.New(app.db), 24*(time.Hour))
	app.managers["trendingUserManager"] = trendingUserManager
	recommendedUsersManager := r_users_manager.New(sqlc.New(app.db), 12*(time.Hour))
	app.managers["recommendedUsersManager"] = recommendedUsersManager
	recommendedPostsManager := r_posts_manager.New(sqlc.New(app.db), 6*(time.Hour))
	app.managers["recommendedUsersManager"] = recommendedPostsManager
	ratingManager := r_manager.New(sqlc.New(app.db), time.Minute)
	app.managers["recommendedUsersManager"] = ratingManager
}

func initManagers(app *App) {

	// Initialize API Managers
	app.PollManager = poll_manager.New(sqlc.New(app.db))
	app.FeedbackManager = feedback_manager.New(sqlc.New(app.db))

}

func initServer(app *App) {
	r := chi.NewRouter()
	handler(r)
	srv := http.Server{
		Addr:    env.ViperGetEnvVar("SERVER_ADDR"),
		Handler: r,
	}
	app.srv = &srv
}
