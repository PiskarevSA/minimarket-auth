package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PiskarevSA/minimarket-points/internal/storage/postgresql"
	"github.com/PiskarevSA/minimarket-points/internal/usecases/signin"
	"github.com/PiskarevSA/minimarket-points/internal/usecases/signup"
	"github.com/PiskarevSA/minimarket-points/pkg/pgx/transactor"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const connUrl = "postgres://user:password@127.0.0.1:5432?sslmode=disable"

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, connUrl)
	if err != nil {
		log.Fatalln(err)
	}

	storage := postgresql.New(pool)

	signup := signup.New(
		"auth",
		storage,
		transactor.New(pool),
		[]byte("jwt"),
		jwt.SigningMethodHS256,
		time.Hour,
		time.Hour,
	)

	result, err := signup.Do(
		ctx, "mtchuikov", "jtuM6mwNvhD3PIHjIwjNfhp1",
	)
	fmt.Println(result, err)

	signin := signin.New(
		storage,
		[]byte("jwt"),
		jwt.SigningMethodHS256,
		time.Hour,
		time.Hour,
	)

	result1, err := signin.Do(
		ctx, "mtchuikov", "jtuM6mwNvD3PIHjIwjNfhp1",
	)
	fmt.Println(result1, err)

}
