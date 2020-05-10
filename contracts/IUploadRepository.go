package contracts

import "nikan.dev/pronto/entities"

type IUploadRepository interface {
	Save(file entities.File) (entities.File,error)
	Get(file []entities.File) ([]entities.File,error)
	Exist(Slug string) (b bool,e error)
}