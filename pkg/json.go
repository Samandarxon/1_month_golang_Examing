package pkg

import (
	"encoding/json"
	"log"
	"os"
)

func Read(fileName string) ([]any, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	var objects []interface{}
	err = json.Unmarshal(data, &objects)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return objects, nil
}

func Write(fileName string, data []any) error {

	body, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return err
}
