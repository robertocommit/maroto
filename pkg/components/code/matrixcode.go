// Package code implements creation of barcode, matrixcode and qrcode.
// nolint:dupl
package code

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type matrixCode struct {
	code   string
	prop   props.Rect
	config *entity.Config
}

// NewMatrix is responsible to create an instance of a MatrixCode.
func NewMatrix(code string, barcodeProps ...props.Rect) core.Component {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &matrixCode{
		code: code,
		prop: prop,
	}
}

// NewMatrixCol is responsible to create an instance of a MatrixCode wrapped in a Col.
func NewMatrixCol(size int, code string, ps ...props.Rect) core.Col {
	matrixCode := NewMatrix(code, ps...)
	return col.New(size).Add(matrixCode)
}

// NewMatrixRow is responsible to create an instance of a MatrixCode wrapped in a Row.
func NewMatrixRow(height float64, code string, ps ...props.Rect) core.Row {
	matrixCode := NewMatrix(code, ps...)
	c := col.New().Add(matrixCode)
	return row.New(height).Add(c)
}

// Render renders a MatrixCode into a PDF context.
func (m *matrixCode) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddMatrixCode(m.code, cell, &m.prop)
}

// GetStructure returns the Structure of a MatrixCode.
func (m *matrixCode) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "matrixcode",
		Value:   m.code,
		Details: m.prop.ToMap(),
	}

	return node.New(str)
}

// SetConfig sets the configuration of a MatrixCode.
func (m *matrixCode) SetConfig(config *entity.Config) {
	m.config = config
}
