package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	//Define data for the welcome message template
	data := map[string]interface{}{"name": "Darshan"}

	teml, err := template.New("welcome").Parse("Welcome, {{.name}}!, How are you?")
	if err != nil {
		panic(err)
	}

	err = teml.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// this will handle error with panic internally
	teml1 := template.Must(template.New("welcome").Parse("Welcome, {{.name}}!, How are you?"))
	err = teml1.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define named Templates in map for different task
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}!, We are glad you joined.",
		"notification": "{{.name}}, You have new notification: {{.notificationMsg}}",
		"error":        "Opps! An error occured: {{.errorMessage}}",
	}

	//Parse and store templates

	parsedTemplates := make(map[string]*template.Template)

	for name, templ := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(templ))
	}

	for {
		//Show Menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option:")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		data := make(map[string]interface{})
		var tmpl *template.Template

		switch choice {
		case "1":
			data["name"] = name
			tmpl = parsedTemplates["welcome"]
		case "2":
			fmt.Println("Enter your notification message:")
			notiMsg, _ := reader.ReadString('\n')
			notiMsg = strings.TrimSpace(notiMsg)
			data["notificationMsg"] = notiMsg
			data["name"] = name
			tmpl = parsedTemplates["notification"]
		case "3":
			fmt.Println("Enter your error message:")
			errMsg, _ := reader.ReadString('\n')
			errMsg = strings.TrimSpace(errMsg)
			data["errorMessage"] = errMsg

			tmpl = parsedTemplates["error"]
		case "4":
			fmt.Println("Exiting......")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option!")
			continue
		}

		err1 := tmpl.Execute(os.Stdout, data)
		if err1 != nil {
			fmt.Println("Error executing template: ", err1)
		}

	}
}
