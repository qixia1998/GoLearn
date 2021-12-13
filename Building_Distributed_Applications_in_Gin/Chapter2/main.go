package main

import(
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"github.com/rs/xid"
)

var recipes []Recipe
func init() {
	recipes = make([]Recipe, 0)
}
type Recipe struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.Run()
}
