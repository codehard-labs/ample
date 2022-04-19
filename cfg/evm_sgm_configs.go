package cfg

import (
	"github.com/hashwavelab/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	EVMSGMConfigSCollectionName = "evm_sgm_configs"
)

type EVMSGMConfig struct {
	// Basic Params
	Name          string  `bson:"name" json:"name"`
	Ver           int     `bson:"ver" json:"ver"`
	EIP1559       bool    `bson:"eip1559" json:"eip1559"`
	Mode          string  `bson:"mode" json:"mode"`
	ConstantForGP float64 `bson:"constantForGP" json:"constantForGP"` // this can be seen as the min gas price
	// Constraint Params
	MaxUSDCost   float64 `bson:"maxUSDCost" json:"maxUSDCost"`
	MaxCostRatio float64 `bson:"maxCostRatio" json:"maxCostRatio"`
	// GP := f(Suggested)
	CoefOfSuggestedPower1 float64 `bson:"coefOfSuggestedPower1" json:"coefOfSuggestedPower1"`
	// CostRatio := f(EP)
	ConstantForCostRatio float64 `bson:"constantForCostRatio" json:"constantForCostRatio"`
	CoefOfEPPower1       float64 `bson:"coefOfEPPower1" json:"coefOfEPPower1"`
	CoefOfEPPower2       float64 `bson:"coefOfEPPower2" json:"coefOfEPPower2"`
	CoefOfLnEP           float64 `bson:"coefOfLnEP" json:"coefOfLnEP"`
	// a / (1 + e ^ -(b * EP))
	CoefOfNominatorOfSigmoid float64 `bson:"coefOfNominatorOfSigmoid" json:"coefOfNominatorOfSigmoid"`
	CoefOfEPOfSigmoid        float64 `bson:"coefOfEPOfSigmoid" json:"coefOfEPOfSigmoid"`
}

func GetEVMSGMConfig(chainName string, ver int) (EVMSGMConfig, error) {
	var res EVMSGMConfig
	var err error
	if ver == 0 {
		filter := &bson.M{"name": chainName}
		ops := options.FindOne().SetSort(bson.D{{Key: "ver", Value: -1}})
		err = mongo.FindOne(EVMSGMConfigSCollectionName, filter, &res, ops)
	} else {
		filter := &bson.M{"name": chainName, "ver": ver}
		err = mongo.FindOne(EVMSGMConfigSCollectionName, filter, &res)
	}
	return res, err
}
