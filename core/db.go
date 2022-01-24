package core

import "errors"

func loadFromDB(name string, ptr interface{}) error {
	switch DB_MODE {
	case DB_MODE_JSON:
		return ReadJSONFile(name, ptr)
	default:
		return errors.New("no mode set")
	}
}

func saveToDB(name string, ptr interface{}) error {
	switch DB_MODE {
	case DB_MODE_JSON:
		return SaveJSONFile(name, ptr)
	default:
		return errors.New("no mode set")
	}
}
