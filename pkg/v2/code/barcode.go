package code

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
)

type barcode struct {
	code   string
	prop   props.Barcode
	config *config.Maroto
}

func NewBar(code string, ps ...props.Barcode) domain.Component {
	prop := props.Barcode{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid()

	return &barcode{
		code: code,
		prop: prop,
	}
}

func NewBarCol(size int, code string, ps ...props.Barcode) domain.Col {
	bar := NewBar(code, ps...)
	return col.New(size).Add(bar)
}

func NewBarRow(height float64, code string, ps ...props.Barcode) domain.Row {
	bar := NewBar(code, ps...)
	c := col.New(0).Add(bar)
	return row.New(height).Add(c)
}

func (b *barcode) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddBarCode(b.code, cell, b.prop)
}

func (b *barcode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "barcode",
		Value: b.code,
	}

	return tree.NewNode(str)
}

func (b *barcode) GetValue() string {
	return b.code
}

func (b *barcode) SetConfig(config *config.Maroto) {
	b.config = config
}
