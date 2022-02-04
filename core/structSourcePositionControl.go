package core

// Weight can possibily be negative even on a source
// if we are allowed to borrow and sell the asset.
// To ensure consistency, weightSum is also int instead of uint.

type SourcePositionControl struct {
	SourceName string            `json:"sourceName"`
	WeightSum  int32             `json:"weightSum"`
	Controls   []PositionControl `json:"controls"`
}

type PositionControl struct {
	Asset             string `json:"asset"`
	WeightControlOn   bool   `json:"weightControlOn"`
	WeightMin         int32  `json:"weightMin"`
	WeightMax         int32  `json:"weightMax"`
	QuantityControlOn bool   `json:"quantityControlOn"`
	QuantityMin       string `json:"quantityMin"` // this can be some decimals that needs to be accurate
	QuantityMax       string `json:"quantityMax"` // this can be some decimals that needs to be accurate
}

func LoadSourcePositionControlsFromDB() ([]SourcePositionControl, error) {
	name := "sourcePositionControls"
	data := []SourcePositionControl{}
	//
	err := loadFromDB(name, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveSourcePositionControlsToDB(data []SourcePositionControl) error {
	name := "sourcePositionControls"
	//
	return saveToDB(name, &data)
}
