package cfg

import (
	"github.com/hashwavelab/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	SourcesCollectionName = "sources"
)

type Source struct {
	Name    string `bson:"name" json:"name"`
	Type    string `bson:"type" json:"type"`
	On      bool   `bson:"on" json:"on"`
	Trading bool   `bson:"trading" json:"trading"`
}

func GetAllSources() ([]Source, error) {
	var res []Source
	err := mongo.FindAll(SourcesCollectionName, mongo.EmptyFilter, &res)
	return res, err
}

func GetOneSource(source string) (Source, error) {
	var res Source
	filter := &bson.M{"name": source}
	err := mongo.FindOne(SourcesCollectionName, filter, &res)
	return res, err
}
