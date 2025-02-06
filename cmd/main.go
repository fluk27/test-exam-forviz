package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"test-exam-forviz/config"
	"test-exam-forviz/internal/models"
	"test-exam-forviz/internal/repositories/db"
	"test-exam-forviz/internal/routers"
	"test-exam-forviz/internal/services"
	"test-exam-forviz/loggers"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	cfg := config.InitConfig()
	loggers.InitLogger(cfg.App)

	DB := initSqlite(cfg.Sqlite)
	migrateDB(DB, models.BookRepository{})
	// repository
	bookRepo := db.NewBookRepository(DB)

	// service
	bookSvc := services.NewBookService(bookRepo)

	e := routers.InitRouter(bookSvc)
	go run(e, cfg.App)
	quit := make(chan os.Signal, 1)
	<-quit
	loggers.Info("receive signal: shutting down...\n")
	if err := e.Shutdown(context.Background()); err != nil {
		loggers.Fatal(err.Error())
	}

}

func run(e *echo.Echo, cfg config.App) {
	serverPort := fmt.Sprintf(":%v", cfg.Port)
	err := e.Start(serverPort)
	if err != nil {
		loggers.Fatal("router shutdown....")
	}

}

func initSqlite(configSqlite config.Sqlite) *gorm.DB {
	sqlitePath := configSqlite.Name
	if strings.TrimSpace(configSqlite.Path) != "" {
		sqlitePath = fmt.Sprintf("%v/%v", configSqlite.Path, configSqlite.Name)

	}
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		loggers.Fatal(fmt.Sprintf("cannot connect sqlite error=%v", err.Error()), zap.Error(err))
	}
	sql, _ := db.DB()
	sql.SetMaxIdleConns(configSqlite.MaxIdleConns)
	sql.SetMaxOpenConns(configSqlite.MaxOpenConns)
	sql.SetConnMaxLifetime(configSqlite.MaxLifeTimeMinutes)
	loggers.Info("connect DB successfully.")
	return db
}

func migrateDB(db *gorm.DB, bookTable models.BookRepository) {
	err := db.AutoMigrate(bookTable)
	if err != nil {
		loggers.Fatal(fmt.Sprintf("AutoMigrate error:%v", err.Error()), zap.Error(err))
	}
	loggers.Info("migrate DB successfully.")
}
