package token_manager

import (
	"context"
	"fmt"
	"time"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/pkg/runner"
)

type TokenManager struct {
	querier sqlc.Querier
	runner  *runner.Runner
}

func New(querier sqlc.Querier, interval time.Duration) *TokenManager {
	return &TokenManager{querier, runner.New(interval)}
}

// Function to clean expired tokens from DB
func (tm *TokenManager) CleanExpiredTokens(ctx context.Context) {
	err := tm.querier.DeleteExpiredTokens(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (tm *TokenManager) Run() {
	tm.runner.Run(func() {
		tm.CleanExpiredTokens(context.Background())
	})
}

func (tm *TokenManager) Close() {
	tm.runner.Close()
}