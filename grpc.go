package main

import (
	"context"
	"log"
	"net"

	"ample/core"
	pb "ample/pb"

	"google.golang.org/grpc"
)

const (
	port = ":8880"
)

type server struct {
	pb.UnimplementedAmpleServer
	db *ConfigDB
}

func (s *server) GetAllSources(ctx context.Context, in *pb.EmptyRequest) (*pb.AllSourcesReply, error) {
	sources := make([]*pb.AllSourcesReply_SourceReply, 0)
	s.db.RLock()
	defer s.db.RUnlock()
	for _, s := range s.db.Sources {
		sources = append(sources, &pb.AllSourcesReply_SourceReply{
			Name:    s.Name,
			Type:    s.Type,
			On:      s.On,
			Trading: s.Trading,
		})
	}
	return &pb.AllSourcesReply{Ok: true, Sources: sources}, nil
}

func (s *server) GetAllTokens(ctx context.Context, in *pb.EmptyRequest) (*pb.AllTokensReply, error) {
	tokens := make([]*pb.AllTokensReply_TokenReply, 0)
	s.db.RLock()
	defer s.db.RUnlock()
	for _, t := range s.db.Tokens {
		tokens = append(tokens, &pb.AllTokensReply_TokenReply{
			AssetName:  t.AssetName,
			GroupName:  t.GroupName,
			Source:     t.Source,
			Identifier: t.Identifier,
			Decimals:   t.Decimals,
			Trading:    t.Trading,
		})
	}
	return &pb.AllTokensReply{Ok: true, Tokens: tokens}, nil
}

func (s *server) GetAllObexTradingPairs(ctx context.Context, in *pb.EmptyRequest) (*pb.AllObexTradingPairsReply, error) {
	pairs := make([]*pb.AllObexTradingPairsReply_ObexTradingPairReply, 0)
	s.db.RLock()
	defer s.db.RUnlock()
	for _, p := range s.db.ObexTradingPairs {
		pairs = append(pairs, &pb.AllObexTradingPairsReply_ObexTradingPairReply{
			PairName:     p.PairName,
			PairType:     p.PairType,
			ExchangeName: p.ExchangeName,
			QuoteAsset:   p.QuoteAsset,
			BaseAsset:    p.BaseAsset,
			StepSize:     p.StepSize,
			TakerFee:     p.TakerFee,
			Trading:      p.Trading,
		})
	}
	return &pb.AllObexTradingPairsReply{Ok: true, Pairs: pairs}, nil
}

func (s *server) GetAllEvmChainClients(ctx context.Context, in *pb.EmptyRequest) (*pb.AllEvmChainClientsReply, error) {
	clients := make([]*pb.AllEvmChainClientsReply_EvmChainClientReply, 0)
	s.db.RLock()
	defer s.db.RUnlock()
	for _, c := range s.db.EvmChainClients {
		clients = append(clients, &pb.AllEvmChainClientsReply_EvmChainClientReply{
			ChainName:              c.ChainName,
			ChainId:                c.ChainID,
			NativeAsset:            c.NativeAsset,
			QueryRpcAddress:        c.QueryRPCAddress,
			ExecRpcAddress:         c.ExecRPCAddress,
			TradingContractAddress: c.TradingContractAddress,
			TradingContractVersion: c.TradingContractVersion,
			RegisteredWorkers:      c.RegisteredWorkers,
			GasMode:                c.GasMode,
			GasType:                c.GasType,
			GasSpecs:               c.GasSpecs,
		})
	}
	return &pb.AllEvmChainClientsReply{Ok: true, Clients: clients}, nil
}

func (s *server) GetAllUniV2Dexs(ctx context.Context, in *pb.EmptyRequest) (*pb.AllUniV2DexsReply, error) {
	dexs := make([]*pb.AllUniV2DexsReply_UniV2DexReply, 0)
	s.db.RLock()
	defer s.db.RUnlock()
	for _, d := range s.db.UniV2Dexs {
		dexs = append(dexs, &pb.AllUniV2DexsReply_UniV2DexReply{
			Name:           d.Name,
			Source:         d.Source,
			RouterAddress:  d.RouterAddress,
			FactoryAddress: d.FactoryAddress,
			SubType:        d.SubType,
			FeeRev:         d.FeeRev,
			Trading:        d.Trading,
		})
	}
	return &pb.AllUniV2DexsReply{Ok: true, Dexs: dexs}, nil
}

func (s *server) GetAllSourceHoldingControls(ctx context.Context, in *pb.EmptyRequest) (*pb.AllSourceHoldingControlsReply, error) {
	src_ctrls := make([]*pb.AllSourceHoldingControlsReply_SourceHoldingControlReply, 0)
	s.db.RLock()
	defer s.db.RUnlock()
	for _, c := range s.db.SourceHoldingControls {
		ctrls := make([]*pb.AllSourceHoldingControlsReply_SourceHoldingControlReply_HoldingControlReply, 0)
		// populate ctrls
		for _, cc := range c.Controls {
			ctrls = append(ctrls, &pb.AllSourceHoldingControlsReply_SourceHoldingControlReply_HoldingControlReply{
				Asset:             cc.Asset,
				WeightControlOn:   cc.WeightControlOn,
				WeightMin:         cc.WeightMin,
				WeightMax:         cc.WeightMax,
				QuantityControlOn: cc.QuantityControlOn,
				QuantityMin:       cc.QuantityMin,
				QuantityMax:       cc.QuantityMax,
			})
		}
		src_ctrls = append(src_ctrls, &pb.AllSourceHoldingControlsReply_SourceHoldingControlReply{
			SourceName: c.SourceName,
			WeightSum:  c.WeightSum,
			Controls:   ctrls,
		})
	}
	return &pb.AllSourceHoldingControlsReply{Ok: true, SourceControls: src_ctrls}, nil
}

func (s *server) GetRawJsonConfig(ctx context.Context, in *pb.KeyRequest) (*pb.RawJsonConfigReply, error) {
	raw_string, err := core.ReadRawJSONString(in.Key)
	if err != nil {
		return &pb.RawJsonConfigReply{Ok: false, RawString: raw_string}, nil
	}
	return &pb.RawJsonConfigReply{Ok: true, RawString: raw_string}, nil
}

func InitGrpcServer(db *ConfigDB) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAmpleServer(s, &server{db: db})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
