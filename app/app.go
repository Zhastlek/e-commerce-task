package app

import (
	"github.com/gin-gonic/gin"

	"e-comerce/internal/adapters/handlers"
	"e-comerce/internal/adapters/repositories"
	"e-comerce/internal/service"
	"e-comerce/pkg"
)

func Initialize() *gin.Engine {
	router := gin.Default()
	config := pkg.Configuration()
	database := repositories.InitDB(config)
	storageEditor := repositories.NewStorageEditor(database)
	storageGetter := repositories.NewStorageGetter(database)
	serviceEditor := service.NewServiceEditor(storageEditor)
	serviceGetter := service.NewServiceGetter(storageGetter)
	handler := handlers.NewHandler(serviceEditor, serviceGetter)
	handler.Register(router)
	return router
}
