package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/fojtas98/CLI/templates"
	"github.com/spf13/cobra"
)

var newRestaurant struct {
	Name      string
	Website   string
	OpenTag   string
	CloseTag  string
	DishCount int
}

var addToMap struct {
	Funcs string
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

		_, b, _, _ := runtime.Caller(0)
		path := path.Dir(filepath.Dir(b))

		newRestaurant.Website = getInfoFromUser("Website")
		newRestaurant.Name = getInfoFromUser("Name")
		newRestaurant.OpenTag = getInfoFromUser("OpenTag")
		newRestaurant.CloseTag = getInfoFromUser("CloseTag")
		newRestaurant.DishCount, _ = strconv.Atoi(getInfoFromUser("DishCount"))

		resTemp := templates.CreateFromTamplate("res")
		f, _ := os.Create(path + "/restaurants/" + newRestaurant.Name + ".go")
		err := resTemp.Execute(f, newRestaurant)
		if err != nil {
			panic(err)
		}
		files, _ := ioutil.ReadDir(path + "/restaurants")
		for _, f := range files {
			name := f.Name()
			name = name[:len(name)-3]
			if name != "map" && name != "index" {
				addToMap.Funcs += `f["` + name + `"] = ` + "GetMenuFrom" + name + "\n" + "\t"
			}
		}
		mapTemp := templates.CreateFromTamplate("map")
		file, _ := os.Create(path + "/restaurants/" + "map.go")
		mapTemp.Execute(file, addToMap)
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
