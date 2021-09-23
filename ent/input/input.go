package input

import (
	"github.com/gdower/bhlinker/ent/name"
	"github.com/gdower/bhlinker/ent/reference"
)

type Input struct {
	ID                  string `json:"id"`
	name.Name           `json:"name"`
	reference.Reference `json:"reference"`
}
