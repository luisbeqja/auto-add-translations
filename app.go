package main

import (
	"context"
	"fmt"

	"strings"

	"github.com/luisbeqja/autoAddTranslation/functions"

	"github.com/tealeg/xlsx"
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
}

func readXlsxFile(path string) {
	file, err := xlsx.OpenFile("existing.xlsx")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Specify the sheet and row you want to read (assuming 0-based indices)
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
}

func (a *App) DivideTabs(text string, path string, translationKey string) string {
	parts := strings.Split(text, "\t")

	langs := [12]string{"en_US", "en_GB", "en_AU", "en_NZ", "en_CA", "de_DE", "es_MX", "es_ES", "fr_CA", "fr_FR", "nl_NL", "pt_BR"}

	for key, part := range parts {
		langPath := path + "/" + langs[key] + ".js"
		err := functions.AddTranslation(part, langPath, translationKey)
		if err != nil {
			return "Error adding translation to " + langs[key]
		}
	}
	return "Your translations have been added to all the files."
}
