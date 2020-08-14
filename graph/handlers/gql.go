package handlers

import (
	"gcode/graph"
	"gcode/graph/generated"
	"gcode/orm"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	cfg := generated.Config{
		Resolvers: &graph.Resolver{
			ORM: orm, // pass in the ORM instance in the resolvers to be used
		},
	}

	h := handler.GraphQL(generated.NewExecutableSchema(cfg))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler defines a handler to expose the Playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
