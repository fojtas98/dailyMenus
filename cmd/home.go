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
	"fmt"

	restaurants "github.com/fojtas98/CLI/restaurants"
	"github.com/spf13/cobra"
)

// homeCmd represents the home command
var homeCmd = &cobra.Command{
	Use:   "home",
	Short: "Get menus of pubs/restaurants near home",
	Run: func(cmd *cobra.Command, args []string) {
		getMenu()
	},
}

func init() {
	rootCmd.AddCommand(homeCmd)
}

func getMenu() {
	var menu []string
	menu = append(menu, restaurants.GetMenuFromLetsMeat()...)
	menu = append(menu, restaurants.GetMenuFrombulvar()...)
	menu = append(menu, restaurants.GetMenuFromvinohradskypivovar()...)
	menu = append(menu, restaurants.GetMenuFromutelleru()...)

	for _, dish := range menu {
		fmt.Println(dish)
	}

}