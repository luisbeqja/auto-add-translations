package functions

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func AddTranslation(translation string, filePath string, translationKey string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	newLine := translationKey + ":" + "'" + translation + "'" + ","

	lines = append(lines[:1], append([]string{newLine}, lines[1:]...)...)

	output, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}

func DownloadHandler() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return err
	}
	// Specify the download folder path
	downloadFolderPath := filepath.Join(homeDir, "Downloads")

	srcFile, err := os.Open("conf/user_config.json")
	if err != nil {
		return err
	}
	defer srcFile.Close()
	fileName := "user_config.json"

	// Create the destination file in the specified folder
	destFileName := filepath.Join(downloadFolderPath, fileName)
	destFile, err := os.Create(destFileName)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the contents from source to destination
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
