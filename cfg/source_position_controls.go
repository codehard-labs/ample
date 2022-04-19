package cfg

import (
	"github.com/hashwavelab/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// Weight can possibily be negative even on a source
// if we are allowed to borrow and sell the asset.
// To ensure consistency, weightSum is also int instead of uint.

var (
	SourcePositionControlsCollectionName = "source_position_controls"
)

type SourcePositionControl struct {
	SourceName string            `bson:"sourceName" json:"sourceName"`
	WeightSum  int32             `bson:"weightSum" json:"weightSum"`
	Controls   []PositionControl `bson:"controls" json:"controls"`
}

type PositionControl struct {
	Asset             string `bson:"asset" json:"asset"`
	WeightControlOn   bool   `bson:"weightControlOn" json:"weightControlOn"`
	WeightMin         int32  `bson:"weightMin" json:"weightMin"`
	WeightMax         int32  `bson:"weightMax" json:"weightMax"`
	QuantityControlOn bool   `bson:"quantityControlOn" json:"quantityControlOn"`
	QuantityMin       string `bson:"quantityMin" json:"quantityMin"` // this can be some decimals that needs to be accurate
	QuantityMax       string `bson:"quantityMax" json:"quantityMax"` // this can be some decimals that needs to be accurate
}

func GetSourcePositionControl(source string) (SourcePositionControl, error) {
	var res SourcePositionControl
	filter := &bson.M{"sourceName": source}
	err := mongo.FindOne(SourcePositionControlsCollectionName, filter, &res)
	return res, err
}
