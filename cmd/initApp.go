/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/spf13/cobra"
)

var appName string
var modelString string

// initAppCmd represents the initApp command
var initAppCmd = &cobra.Command{
	Use:   "initApp",
	Short: "Initializes Fiber microservice",
	Long: `
This command will initialize a Golang Fiber microservice and set
up the folder structure as well.

Usage: copped-fiber-start initApp --app-name=<app-name> --models='<entity>, <entity-2>'.
`,
	Run: func(cmd *cobra.Command, args []string) {
		newServicePath := "/Users/disjosh/go/src/github.com/COPPED/" + appName
		// newServicePath := "/Users/disjosh/go/src/github.com/COPPED/copped-fiber-starter/" + appName
		createApp(newServicePath)
		createFolders(newServicePath)
		createModels(newServicePath)
		createServer(newServicePath)
		createCmdFile(newServicePath)
		copyShellScript(newServicePath)
		copyReadMe(newServicePath)
	},
}

func createApp(newServicePath string) {
	if err := os.Mkdir(newServicePath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done creating app.\n")
}

func createFolders(newServicePath string) {
	nestedCmd := "cmd/" + appName
	nestedInternal := "internal/" + appName
	folders := []string{
		"build", "cmd", nestedCmd, "docs", "internal", "internal/app", "internal/database",
		nestedInternal, nestedInternal + "/repo", nestedInternal + "/rest", nestedInternal + "/service",
		"pkg", "pkg/models",
	}
	for _, folder := range folders {
		folderPath := newServicePath + "/" + folder
		if err := os.Mkdir(folderPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Done creating directory: %s.\nPath: %s\n", folder, folderPath)
	}
}

func createModels(newServicePath string) {
	fmt.Println("TEST")
	models := strings.Split(modelString, ",")
	for _, model := range models {
		formattedModel := strings.ReplaceAll(model, " ", "")
		fmt.Printf("formattedModel: %s\n", formattedModel)
		fileName := formattedModel + ".go"
		filePath, filePathErr := filepath.Abs(newServicePath)
		if filePathErr != nil {
			log.Fatal(filePathErr)
		}
		outputFile := filePath + "/pkg/models/" + fileName
		processTemplate("struct.tmpl", formattedModel, outputFile)
	}
	fmt.Printf("Done creating models.\n")
}

func createServer(newServicePath string) {
	fileName := "http.go"
	filePath, filePathErr := filepath.Abs(newServicePath)
	if filePathErr != nil {
		log.Fatal(filePathErr)
	}
	outputFile := filePath + "/internal/app/" + fileName
	processTemplate("http.tmpl", "", outputFile)

	fmt.Printf("Done creating server in app/http.go.\n")
}

func createCmdFile(newServicePath string) {
	fileName := "main.go"
	filePath, filePathErr := filepath.Abs(newServicePath)
	if filePathErr != nil {
		log.Fatal(filePathErr)
	}
	outputFile := filePath + "/cmd/" + appName + "/" + fileName
	processTemplate("main.tmpl", appName, outputFile)

	fmt.Printf("Done creating driver app in cmd/main.go.\n")
}

func processTemplate(fileName string, model string, outputFile string) {
	tmpl := template.Must(template.New("").Funcs(sprig.FuncMap()).ParseFiles(fileName))
	var processed bytes.Buffer
	err := tmpl.ExecuteTemplate(&processed, fileName, model)
	if err != nil {
		log.Fatal(err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Couldn't format template: %v'n", err)
	}
	fmt.Println("Writing file: ", outputFile)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(f)
	w.WriteString(string(formatted))
	w.Flush()
}

func copyShellScript(newServicePath string) {
	destFilePath := newServicePath + "/initProject.sh"
	_, err := copy("initProject.sh", destFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done copying shell script to: %s\n", destFilePath)
	// cmd, err := exec.Command("/bin/sh", destFilePath).Output()
	// if err != nil {
	// 	log.Fatalf("Error running .sh script: %+v", err)
	// }
	// output := string(cmd)
	// fmt.Printf("output: %s\n", output)
	// fmt.Println("Done initializing modules.")
}

func copyReadMe(newServicePath string) {
	destFilePath := newServicePath + "/README.md"
	_, err := copy("README.md", destFilePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done copying README.md file to: %s\n", destFilePath)
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func init() {
	rootCmd.AddCommand(initAppCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initAppCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initAppCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	initAppCmd.Flags().StringVarP(&appName, "app-name", "a", "", "The name you want your app defined as.")
	initAppCmd.Flags().StringVarP(&modelString, "models", "m", "", "Names of models you will define.")
}
