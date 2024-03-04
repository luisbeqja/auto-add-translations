package main

import (
	"context"

	"github.com/luisbeqja/autoAddTranslation/functions"
	"github.com/luisbeqja/autoAddTranslation/ui"
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

func (a *App) AddConfigurationHandler(file string) string {
	return ui.AddConfigurationHandler(file)
}

func (a *App) DowloadConfigHandler() string {
	return ui.DowloadConfigHandler()
}

func (a *App) AddTranslationHandler(text string, path string, translationKey string) string {
	return ui.AddTranslationHandler(text, path, translationKey)
}
