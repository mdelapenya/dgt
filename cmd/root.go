package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mdelapenya/dgt/scrap"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dgt",
	Short: "Inits a web server with a REST API to fetch the sticker for a given plate",
	Long:  `Inits a web server with a REST API to fetch the sticker for a given plate`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

// Execute execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed!")
	}
}

func fetchPlate(c *gin.Context) {
	plate := c.Param("plate")

	sticker := scrap.ProcessPlate(plate, true)

	status := http.StatusOK

	if sticker == "Not found" || sticker == scrap.NotFound {
		status = http.StatusNotFound
	}

	c.JSON(status, gin.H{
		"result": sticker,
	})
}

func runServer() {
	router := gin.Default()

	router.GET("/plates/:plate", fetchPlate)

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.Run()
}
