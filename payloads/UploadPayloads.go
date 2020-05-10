package payloads

import internalContracts "nikan.dev/pronto/internals/contracts"

type UploadPayload struct {
	Description string
}


func (i UploadPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {
	return []internalContracts.IValidatable {
		validator.Validatable().Field(i.Description).Name("Description").Require().String(),
	};
}




//func (i UploadPayload) Validation(validator internalContracts.IValidator) []internalContracts.IValidatable {

