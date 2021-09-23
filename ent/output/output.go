package output

import (
	"github.com/gdower/bhlinker/ent/name"
	"github.com/gdower/bhlinker/ent/reference"
	"github.com/gdower/bhlinker/ent/score"
	bhln "github.com/gnames/bhlnames/domain/entity"
)

type Output struct {
	InputID      string              `json:"inputId"`
	InputName    name.Name           `json:"inputName"`
	InputRef     reference.Reference `json:"inputRef,omitempty"`
	OutputName   string              `json:"outputName,omitempty"`
	EditDistance int                 `json:"editDistance,omitempty"`
	Error        error               `json:"error,omitempty"`
	BHLref       *bhln.Reference     `json:"referenceBHL"`
	score.Score  `json:"score"`
}
