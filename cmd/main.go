package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/marcelluseasley/saturday-night-something-to-do/handlers"
	"github.com/marcelluseasley/saturday-night-something-to-do/migrations"
	"github.com/marcelluseasley/saturday-night-something-to-do/service"
	"github.com/marcelluseasley/saturday-night-something-to-do/storage"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	handleDBMigrations(dbPool)

	pgStorage := storage.NewPostgresStorage(dbPool)
	service := service.NewUserService(pgStorage)
	userHandler := handlers.NewUserHandler(service)
	app := setupRoutes(userHandler)

	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Fiber was successful shutdown.")
}

func handleDBMigrations(pool *pgxpool.Pool) {
	c, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer c.Release()
	migrator, err := migrations.NewMigrator(c.Conn())
	if err != nil {
		log.Fatal(err)
	}

	now, exp, info, err := migrator.Info()
	if err != nil {
		panic(err)
	}
	if now < exp {

		println("migration needed, current state:")
		println(info)

		err = migrator.Migrate()
		if err != nil {
			panic(err)
		}
		println("migration successful!")
	} else {
		println("no database migration needed")
	}
}

func setupRoutes(uh *handlers.UserHandler) *fiber.App {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(healthcheck.New())

	app.Post("/users", uh.CreateUser)
	app.Get("/users/:id", uh.GetUser)
	app.Patch("/users", uh.UpdateUser)
	app.Delete("/users/:id", uh.DeleteUser)
	app.Get("/users", uh.ListUsers)

	return app
}
