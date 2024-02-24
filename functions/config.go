package functions

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateInitialConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}
	infoFolder := filepath.Join(homeDir, "Add-translations-config")
	fileName := "user_config.json"

	// Create the folder if it doesn't exist
	if _, err := os.Stat(infoFolder); os.IsNotExist(err) {
		err := os.Mkdir(infoFolder, 0755)
		if err != nil {
			return err
		}
	}
	// Create the file if it doesn't exist
	if _, err := os.Stat(infoFolder + "/" + fileName); os.IsNotExist(err) {
		_, err := os.Create(infoFolder + "/" + fileName)
		if err != nil {
			return err
		}
	} else {
		return nil
	}
	// Write the initial content to the file
	err = os.WriteFile(infoFolder+"/"+fileName, []byte(`{
		"languages": [
			"en_US",
			"en_GB",
			"en_AU",
			"en_NZ",
			"en_CA",
			"de_DE",
			"es_MX",
			"es_ES",
			"fr_CA",
			"fr_FR",
			"nl_NL",
			"pt_BR" 
		],
		"file_tipe": "json"
		}`), 0644)

	if err != nil {
		return err
	}
	return nil
}
