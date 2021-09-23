package input

import (
	"github.com/gdower/bhlio/ent/name"
	"github.com/gdower/bhlio/ent/reference"
)

type Input struct {
	ID                  string `json:"id"`
	name.Name           `json:"name"`
	reference.Reference `json:"reference"`
}
