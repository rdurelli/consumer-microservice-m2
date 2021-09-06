/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"consumer-rabbitmq/controller"

	"github.com/gin-gonic/gin"

	"github.com/spf13/cobra"
)

var (
	router          *gin.Engine
	webServicePort  string
	controllerEmail controller.EmailController
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A web service to discover all information regarding the emails",
	Run: func(cmd *cobra.Command, args []string) {
		router.Run(":" + webServicePort)
	},
}

func init() {
	cobra.OnInitialize(initServer, initEmailController, initRoutes)
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&webServicePort, "port", "p", "8080", "port to create the webservice")
}

func initServer() {
	router = gin.Default()
}
func initEmailController() {
	//repo has been already declared in the package cmd
	controllerEmail = controller.NewEmailController(repo)
}

func initRoutes() {
	api := router.Group("/api")
	{
		api.GET("/emails", controllerEmail.Find)
		api.GET("/emails/showAll", controllerEmail.ShowAll)
	}
}
