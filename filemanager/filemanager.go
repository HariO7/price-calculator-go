package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	inputPath  string
	outputPath string
}

type IOManger interface {
	ReadDataFromFile() ([]string, error)
	WriteJSON(data any) error
}

func New(inputPath string, outputPath string) FileManager {
	return FileManager{
		inputPath:  inputPath,
		outputPath: outputPath,
	}
}

func (fm FileManager) ReadDataFromFile() ([]string, error) {

	file, err := os.Open(fm.inputPath)

	if err != nil {
		fmt.Println("File cannot be openned")
		fmt.Println(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("error occured while scanning")
		fmt.Println(err)
		file.Close()
		return nil, err
	}

	file.Close()

	return data, nil

}

func (fm FileManager) WriteJSON(data interface{}) error {

	file, err := os.Create(fm.outputPath)

	if err != nil {
		file.Close()
		fmt.Println(err)
		return errors.New("failed creating a new file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		fmt.Println(err)
		return errors.New("failed writing a new file")
	}

	file.Close()
	return nil

}
