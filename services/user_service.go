package services

import (
	"back/db"
	"back/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func (u *UserService) GetUserByID(ctx context.Context) (any, error) {
	response, err := u.queries.GetUserByID(ctx, 0)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewUserService() (*UserService, error) {
	env := utils.NewEnv()
	log.Printf("Database configuration: Host=%s, Port=%s, DB=%s, User=%s",
		env.DBHost, env.DBPort, env.DBName, env.DBUser)

	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?connect_timeout=5",
		env.DBUser,
		env.DBPassword,
		"localhost",
		env.DBPort,
		env.DBName,
	)

	// Create pool configuration
	config, err := pgxpool.ParseConfig(connString)
	log.Print(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %v", err)
	}

	// Set some reasonable pool settings
	config.MaxConns = 10
	config.MinConns = 1
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	// Create connection with timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Print("Attempting to create connection pool...")
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	log.Print("Connection pool created, attempting to ping...")

	// Ping with timeout
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping failed: %v", err)
	}

	log.Print("Database connection successful!")

	return &UserService{
		queries: db.New(pool),
		pool:    pool,
	}, nil
}
