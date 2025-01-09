package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Dolg0ff/todoapp-audit-log/internal/config"
	"github.com/Dolg0ff/todoapp-audit-log/internal/repository"
	"github.com/Dolg0ff/todoapp-audit-log/internal/server"
	"github.com/Dolg0ff/todoapp-audit-log/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.URI, cfg.DB.Username)
	opts := options.Client().ApplyURI(uri)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("Mongo connect: %s", err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)
	db.Collection(os.Getenv("DB_COLLECTION"))

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := server.NewAuditServer(auditService)
	srv := server.New(auditSrv)

	fmt.Println("SERVER STARTED", time.Now())

	if err := srv.ListenAndServe(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
