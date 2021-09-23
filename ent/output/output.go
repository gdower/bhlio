package output

import (
	"github.com/gdower/bhlio/ent/name"
	"github.com/gdower/bhlio/ent/refbhl"
	"github.com/gdower/bhlio/ent/reference"
	"github.com/gdower/bhlio/ent/score"
)

type Output struct {
	InputID      string               `json:"inputId"`
	InputName    name.Name            `json:"inputName"`
	InputRef     reference.Reference  `json:"inputRef,omitempty"`
	OutputName   string               `json:"outputName,omitempty"`
	EditDistance int                  `json:"editDistance,omitempty"`
	Error        error                `json:"error,omitempty"`
	BHLref       *refbhl.ReferenceBHL `json:"referenceBHL"`
	score.Score  `json:"score"`
}
