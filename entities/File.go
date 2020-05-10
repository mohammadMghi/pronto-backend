package entities

import "nikan.dev/pronto/internals/entity"

type File struct {
	entity.BaseEntity
	Description string
	Filename    string
}
