package main

import (
	//"fmt"
	"log"
	"net/http"
	//"time"

	"example/story_maker/backend/database"
	"example/story_maker/backend/model"
	"example/story_maker/backend/requestStructures"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)



func handleGetRoot(c *gin.Context) {
    c.IndentedJSON(
		http.StatusOK, 
		map[string]string{"value" : "Hello world"},
	)
}

func handleGetStories(c *gin.Context, db model.Model) {
	c.IndentedJSON(
		http.StatusOK, 
		db.GetStories(),
	)
}

func handleDeleteStories(c *gin.Context, db model.Model) {
	var data requestStructures.RequestBodyDeleteStories
	if err := c.BindJSON(&data); err != nil {
		log.Printf("Deleting stories failed with error: %v\n", err)
        c.IndentedJSON(
			http.StatusNotModified, 
			map[string]bool{"success" : false},
		)
		return
    }
	
	db.DeleteStories(data.Ids)
	c.IndentedJSON(
		http.StatusOK, 
		map[string]any{"success" : true},
	)
}

func handleAddStory(c *gin.Context, db model.Model) {
	var newStory model.Story

    // Call BindJSON to bind the received JSON to newStory.
	// TODO: make all fields in json required
    if err := c.BindJSON(&newStory); err != nil {
		log.Printf("Adding story failed with error: %v\n", err)
        c.IndentedJSON(
			http.StatusNotModified, 
			map[string]bool{"success" : false},
		)
		return
    }
	addesStory := db.AddStory(newStory)
	c.IndentedJSON(
		http.StatusOK, 
		map[string]any{"success" : true, "story" : addesStory},
	)
}

func main() {
	// model initialization
	db := model.Model(database.MakeDatabase())
	log.Println("Db: ", db)

	// middleware
	router := gin.Default()
	router.Use(cors.Default()) // CORS - allow all origins

	// routes
    router.GET("/", handleGetRoot)
	router.GET(
		"/stories", 
		func(c *gin.Context){handleGetStories(c, db)},
	)
	router.POST(
		"/add-story", 
		func(c *gin.Context){handleAddStory(c, db)},
	)
	router.DELETE(
		"/delete-stories",
		func(c *gin.Context){handleDeleteStories(c, db)},
	)

    router.Run("localhost:8000")
}
