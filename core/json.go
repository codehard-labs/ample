package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const JSON_CFG_FILE_PATH = "cfg/"

func ReadJSONFile(name string, dataPtr interface{}) error {
	jsonFile, err := os.Open(JSON_CFG_FILE_PATH + name + ".json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	err1 := json.Unmarshal(jsonData, dataPtr)
	return err1
}

func SaveJSONFile(name string, dataPtr interface{}) error {
	json, err := json.MarshalIndent(dataPtr, "", "    ")
	if err != nil {
		log.Println("saveJSONFile ERROR", err)
	}
	return ioutil.WriteFile(JSON_CFG_FILE_PATH+name+".json", json, 0644)
}

func ReadRawJSONString(name string) (string, error) {
	jsonData, err := ioutil.ReadFile(JSON_CFG_FILE_PATH + name + ".json")
	if err != nil {
		return "", err
	}
	return strings.Join(strings.Fields(string(jsonData)), " "), nil
}
