package main

import (
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/endpoints/http"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/repositories/mysql"
	"nikan.dev/pronto/services"
)


func main() {

	gateway := drivers.NewEchoGateway()
	config := drivers.NewViperConfiguration()
	validator := drivers.NewOzzoValidator()
	pool := drivers.NewGormDriver().Boot(config,entities.User{}, entities.Category{}, entities.Post{}, entities.Setting{},entities.File{})

	deps := dependencies.CommonDependencies{config,validator}


	userRepository := mysql.NewUserRepository(deps, pool)
	userService := services.NewUserService(deps,userRepository)
	userEndpoint := http.NewUserEndpoint(deps,userService)



	 uploadRepository := mysql.NewUploadRepository(deps,pool)
	 uploadService := services.NewUploadService(deps,uploadRepository)
	 uploadEndpoint := http.NewUploadFileEndpoint(deps,uploadService)


	categoryRepository := mysql.NewCategoryRepository(deps, pool)
	categoryService := services.NewCategoryService(deps,categoryRepository)
	categoryEndpoint := http.NewCategoryEndpoint(deps,categoryService)



	applicationRepository := mysql.NewApplicationRepository(deps, pool)
	applicationService := services.NewApplicationService(deps, applicationRepository)
	applicationEndpoint := http.NewApplicationEndpoint(deps,applicationService)



	postRepository := mysql.NewPostRepository(deps, pool)
	postService := services.NewPostService(deps, postRepository)
	postEndpoint := http.NewPostEndpoint(deps,postService)

	gateway.Boot(config, applicationEndpoint, userEndpoint, categoryEndpoint, postEndpoint,uploadEndpoint)

}
