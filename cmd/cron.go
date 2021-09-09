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
	"consumer-rabbitmq/model"
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
)

var (
	minutes int
)

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Use this command to specify minute to monitor if the emails have been sent.",
	Long:  `This command is used to create a job to run every specific minutes.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := gocron.NewScheduler()
		s.Every(uint64(minutes)).Minutes().Do(func() {
			fmt.Println("Vai chamando...")
			emails, _ := repo.ShowAllSentFlag(false)
			for _, email := range *emails {
				user := model.User{
					Name:     email.Name,
					LastName: email.LastName,
					Email:    email.Email,
				}
				err := emailService.SendEmail(user)
				if err != nil {
					fmt.Println("Could not send the email to: ", user.Email, err)
				}
				email.Sent = true
				repo.Update(email.Id, email)

				fmt.Println(email.Email)
			}
		})
		<-s.Start()
	},
}

func init() {
	rootCmd.AddCommand(cronCmd)
	cronCmd.Flags().IntVarP(&minutes, "minutes", "m", 1, "Minutes to monitor if all emails have been sent. Then try to send the email again.")
}
