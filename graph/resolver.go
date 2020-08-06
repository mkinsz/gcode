package graph

import (
	"gcode/graph/model"
	"gcode/orm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct
type Resolver struct {
	ORM      *orm.ORM
	articles []model.Article
}
