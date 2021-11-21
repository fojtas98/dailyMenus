package cmd

import (
	"bufio"
	"fmt"
	"os"

	u "github.com/fojtas98/CLI/template"
	"github.com/spf13/cobra"
)

var newRestaurant struct {
	Name     string
	Website  string
	OpenTag  string
	CloseTag string
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		newRestaurant.Website = getInfoFromUser("Website")
		newRestaurant.Name = getInfoFromUser("Name")
		newRestaurant.OpenTag = getInfoFromUser("OpenTag")
		newRestaurant.CloseTag = getInfoFromUser("CloseTag")

		template := u.CreateFromTamplate()
		currentWorkingDirectory, _ := os.Getwd()
		f, _ := os.Create(currentWorkingDirectory + "/restaurants/" + newRestaurant.Name + ".go")
		err := template.Execute(f, newRestaurant)
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func getInfoFromUser(name string) string {
	fmt.Print("Enter " + name + ": ")
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line = scanner.Text()
	if line == "" {
		fmt.Println(name + "needs to be specified")

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error encountered:", err)
	}
	return line
}