// Package col implements creation of columns.
package col

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type col struct {
	size       int
	isMax      bool
	components []core.Component
	config     *entity.Config
	style      *props.Cell
}

// New is responsible to create an instance of core.Col.
func New(size ...int) core.Col {
	if len(size) == 0 {
		return &col{isMax: true}
	}

	return &col{size: size[0]}
}

// Add is responsible to add a component to a core.Col.
func (c *col) Add(components ...core.Component) core.Col {
	c.components = append(c.components, components...)
	return c
}

// GetSize returns the size of a core.Col.
func (c *col) GetSize() int {
	if c.isMax {
		return c.config.MaxGridSize
	}

	return c.size
}

// GetStructure returns the Structure of a core.Col.
func (c *col) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "col",
		Value:   c.size,
		Details: c.style.ToMap(),
	}

	if c.isMax {
		if len(str.Details) == 0 {
			str.Details = make(map[string]interface{})
		}
		str.Details["is_max"] = true
	}

	node := node.New(str)

	for _, c := range c.components {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

// Render renders a core.Col into a PDF context.
func (c *col) Render(provider core.Provider, cell entity.Cell, createCell bool) {
	if createCell {
		provider.CreateCol(cell.Width, cell.Height, c.config, c.style)
	}

	for _, component := range c.components {
		component.Render(provider, &cell)
	}
}

// SetConfig set the config for the component.
func (c *col) SetConfig(config *entity.Config) {
	c.config = config
	for _, component := range c.components {
		component.SetConfig(config)
	}
}

// WithStyle sets the style for the column.
func (c *col) WithStyle(style *props.Cell) core.Col {
	c.style = style
	return c
}
