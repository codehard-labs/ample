package main

import (
	"ample/core"
	"log"
	"sync"
)

type ConfigDB struct {
	sync.RWMutex
	Sources                []core.Source
	Tokens                 []core.Token
	SourcePositionControls []core.SourcePositionControl
	ObexTradingPairs       []core.ObexTradingPair
	EvmChainClients        []core.EvmChainClient
	UniV2Dexs              []core.UniV2Dex
}

func main() {
	sources, err := core.LoadSourcesFromDB()
	if err != nil {
		log.Fatal(err)
	}
	tokens, err := core.LoadTokensFromDB()
	if err != nil {
		log.Fatal(err)
	}
	sourcePositionControls, err := core.LoadSourcePositionControlsFromDB()
	if err != nil {
		log.Fatal(err)
	}
	obexTradingPairs, err := core.LoadObexTradingPairsFromDB()
	if err != nil {
		log.Fatal(err)
	}
	evmChainClients, err := core.LoadEvmChainClientsFromDB()
	if err != nil {
		log.Fatal(err)
	}
	uniV2Dexs, err := core.LoadUniV2DexsFromDB()
	if err != nil {
		log.Fatal(err)
	}
	db := &ConfigDB{
		Sources:                sources,
		Tokens:                 tokens,
		SourcePositionControls: sourcePositionControls,
		ObexTradingPairs:       obexTradingPairs,
		EvmChainClients:        evmChainClients,
		UniV2Dexs:              uniV2Dexs,
	}
	InitGrpcServer(db)
}
