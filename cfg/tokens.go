package cfg

import (
	"github.com/hashwavelab/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	TokensCollectionName = "tokens"
)

// Asset on a specific Source to be considered
type Token struct {
	AssetName  string `bson:"assetName" json:"assetName"`
	GroupName  string `bson:"groupName" json:"groupName"`
	Source     string `bson:"source" json:"source"`
	Identifier string `bson:"identifier" json:"identifier"` // source specific name
	Decimals   uint32 `bson:"decimals" json:"decimals"`     // only applicable if source is some blockchain
	Trading    bool   `bson:"trading" json:"trading"`
}

func GetAllTokens() ([]Token, error) {
	var res []Token
	err := mongo.FindAll(TokensCollectionName, mongo.EmptyFilter, &res)
	return res, err
}

func GetAllTokensOfOneSource(source string) ([]Token, error) {
	var res []Token
	filter := &bson.M{"source": source}
	err := mongo.FindAll(TokensCollectionName, filter, &res)
	return res, err
}
