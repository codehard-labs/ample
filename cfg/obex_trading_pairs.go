package cfg

import (
	"github.com/hashwavelab/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	ObexTradingPairsCollectionName = "obex_trading_pairs"
)

type ObexTradingPair struct {
	PairName     string `bson:"pairName" json:"pairName"`
	PairType     string `bson:"pairType" json:"pairType"`
	ExchangeName string `bson:"exchangeName" json:"exchangeName"`
	QuoteAsset   string `bson:"quoteAsset" json:"quoteAsset"`
	BaseAsset    string `bson:"baseAsset" json:"baseAsset"`
	StepSize     string `bson:"stepSize" json:"stepSize"`
	TakerFee     string `bson:"takerFee" json:"takerFee"`
	Trading      bool   `bson:"trading" json:"trading"`
}

func GetAllObexTradingPairs() ([]ObexTradingPair, error) {
	var res []ObexTradingPair
	err := mongo.FindAll(ObexTradingPairsCollectionName, mongo.EmptyFilter, &res)
	return res, err
}

func GetObexTradingPairsOfOneExchange(exchange string) ([]ObexTradingPair, error) {
	var res []ObexTradingPair
	filter := &bson.M{"exchangeName": exchange}
	err := mongo.FindAll(ObexTradingPairsCollectionName, filter, &res)
	return res, err
}
