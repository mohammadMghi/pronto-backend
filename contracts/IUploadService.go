package contracts

import (
	"mime/multipart"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/payloads"
)

type IUploadService interface {
	Get(payloads.UploadPayload)(entities.File,error)
	Upload(file *multipart.FileHeader,payload payloads.UploadPayload) (entities.File ,error)

	CheckExist(filename string)(b bool,err error)
}