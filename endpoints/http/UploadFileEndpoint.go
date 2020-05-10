package http

import (
	"github.com/labstack/echo/v4"
	"nikan.dev/pronto/contracts"
	"nikan.dev/pronto/drivers"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/payloads"
)

type IUploadEndpoint struct {
	deps dependencies.CommonDependencies
	service contracts.IUploadService
}

func (endpoint IUploadEndpoint)Boot(transport interface{}){
	t := transport.(*echo.Group)
	group := t.Group("/file")
	group.POST("/upload",endpoint.upload)

}

func (endpoint IUploadEndpoint)upload(ctx echo.Context) error{
	//make payload
	input, errPayload := drivers.RequestToPayload(ctx, new(payloads.UploadPayload))
	if errPayload != nil {
		return errPayload
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}



	payload, err := endpoint.service.Upload(file,*input.(*payloads.UploadPayload))
	return drivers.TypeToResponse(ctx, payload, err)

	//return ctx.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename))
}


func NewUploadFileEndpoint(deps dependencies.CommonDependencies, service contracts.IUploadService) IUploadEndpoint {
	return IUploadEndpoint{deps,service}
}


