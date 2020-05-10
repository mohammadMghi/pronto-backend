package services

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/entities"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
	"os"
	"path"
	"strconv"
	"strings"
)

type uploadService struct {
	repository contracts.IUploadRepository
	deps dependencies.CommonDependencies
}

func (services uploadService)Get(payload payloads.UploadPayload)(entities.File,error){
	//input,error :=
	return entities.File{},nil
}

func(services uploadService)Upload(file *multipart.FileHeader,payload payloads.UploadPayload) (entities.File ,error){
	//service
	src, err := file.Open()
	if err != nil {
		return entities.File{},err
	}
	defer src.Close()


	//valodation
	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = src.Read(buff)
	if err != nil {
		return entities.File{},err
	}
	filetype := http.DetectContentType(buff)
	//TODO:: file validation check file type with 512 byte of the file FOR sure more security we also need check file extension
	errType :=  FileValidation(filetype)
	if errType != nil{
		return entities.File{},err
	}


	src.Seek(0, 0)



	//init
	i := 0
	myfilename := file.Filename
	nameCon, err :=  services.CheckExist(file.Filename)
	if err!= nil{
		return entities.File{},err
	}

	for nameCon !=false{
		//check file name in database
		//if exist return true , we can continue to looping
		// if return false we can exit to loop ...
		i = i + 1
		file.Filename =  FilenameWithoutExtension(myfilename) + strconv.FormatInt(int64(i), 10) +FilenameExtension(file.Filename)
		if exist, err := services.CheckExist(file.Filename); exist!=false{
			if err!= nil{
				return entities.File{},err
			}
			file.Filename =  FilenameWithoutExtension(file.Filename) + strconv.FormatInt(int64(i), 10) +FilenameExtension(file.Filename)
			nameCon = true
		}else{
			file.Filename =  FilenameWithoutExtension(myfilename) + strconv.FormatInt(int64(i), 10) +FilenameExtension(file.Filename)
			nameCon = false
		}
	}


	dst, err := os.OpenFile("./img/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return entities.File{},err
	}
	defer dst.Close()
	// Destination
	if err != nil {
		return entities.File{},err
	}
	defer dst.Close()



	// Copy
	if _, err = io.Copy(dst,src); err != nil {
		return entities.File{},err
	}

	if err := services.deps.Validator.Validate(payload); err != nil {
		return entities.File{}, err
	}
	return services.repository.Save(entities.File{
		Description : payload.Description,
		Filename:     file.Filename,
	})
}

func (services uploadService)CheckExist(filename string)(b bool,err error){
	if exist,err := services.repository.Exist(filename);exist != false{
		return true,err
	}
     	return false,err
}


func NewUploadService(deps dependencies.CommonDependencies,repository contracts.IUploadRepository) uploadService {
	return uploadService{
		repository,
		deps,
	}
}


//validation

//validate : upload file
func  FileValidation (filetype string ) error{

	switch filetype{
	case "image/jpeg", "image/jpg":
	default:
		return errors.New("Error in format file")
	}
	return nil
}
func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func FilenameExtension(fn string) string {
	return path.Ext(fn)
}