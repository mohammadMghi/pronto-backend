package mysql

import (
	"github.com/jinzhu/gorm"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/exceptions"
	"nikan.dev/pronto/internals/dependencies"
)

type uploadRepository struct {
	pool *gorm.DB
	deps dependencies.CommonDependencies
}

func NewUploadRepository(deps dependencies.CommonDependencies , pool interface{}) uploadRepository{
	return uploadRepository{pool: pool.(*gorm.DB) , deps: deps}
}

func (repository uploadRepository) Save(file entities.File) (entities.File,error){
	if err := repository.pool.Create(&file).Error;err != nil{
		return file,exceptions.UploadError
	}
	return file,nil
}

//TODO ::need to changes
func (repository uploadRepository)Get(file []entities.File) ([]entities.File,error){
	if err := repository.pool.First(&file).Error;err != nil{
		return file,exceptions.UploadError
	}
	return file,nil
}

//TODO :: here we return error with nil !
func (repository uploadRepository) Exist(filename string) (b bool,e error) {
	fileCount := 0
	repository.pool.Model(&entities.File{}).Where("filename = ?", filename).Count(&fileCount)
	if fileCount > 0 {
		return true,nil
	}
	return false, nil
}
