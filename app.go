package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"strings"

	"github.com/luisbeqja/autoAddTranslation/functions"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	functions.CreateInitialConfig()
}

// FOR FUTURE USE: possibility to add directly from an excel file giving just the line of the translation
/* func readXlsxFile(path string) {
	file, err := xlsx.OpenFile("existing.xlsx")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	sheetIndex := 0 // Index of the sheet (e.g., 0 for the first sheet)
	rowIndex := 0   // Index of the row (e.g., 0 for the first row)

	// Specify the range of columns to read (C to N)
	startColIndex := 2 // Index of column C (0-based)
	endColIndex := 13  // Index of column N (0-based)

	row := file.Sheets[sheetIndex].Rows[rowIndex]

	for colIndex := startColIndex; colIndex <= endColIndex; colIndex++ {
		// Check if the column index is within bounds
		if colIndex < len(row.Cells) {
			cell := row.Cells[colIndex]
			fmt.Printf("Column %c: %s\n", 'A'+colIndex, cell.String())
		} else {
			fmt.Printf("Column %c: (empty)\n", 'A'+colIndex)
		}
	}
} */

func (a *App) AddConfiguration(file string) string {
	homeDir, _ := os.UserHomeDir()
	// Your input string
	inputString := file
	path := homeDir + "/Add-translations-config/user_config.json"
	// Write the string to a file
	err := os.WriteFile(path, []byte(inputString), 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return ""
	}
	return "Your configuration has been added to the file."
}

func (a *App) DowloadConfig() string {
	err := functions.DownloadHandler()
	if err != nil {
		return "Error downloading file:" + err.Error()
	}
	return "File user_config.json in Downloads folder"
}

func (a *App) DivideTabs(text string, path string, translationKey string) string {

	homeDir, _ := os.UserHomeDir()
	configFle := homeDir + "/Add-translations-config/user_config.json"
	jsonFile, err := os.Open(configFle)
	if err != nil {
		fmt.Println(err)
		return "Error opening file"
	}
	defer jsonFile.Close()

	// Read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	langsArray := result["languages"].([]interface{})

	fmt.Println(langsArray[0])

	parts := strings.Split(text, "\t")

	for key, part := range parts {
		if len(langsArray) > key {
			langPath := path + "/" + fmt.Sprintf("%v", langsArray[key]) + ".js"
			err := functions.AddTranslation(part, langPath, translationKey)
			if err != nil {
				return "Error adding translation to " + fmt.Sprintf("%v", langsArray[key])
			}
		}

	}
	return "Your translations have been added to all the files."
}
