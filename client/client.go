package client

import (
	"context"
	"errors"
	"time"

	"github.com/hashwavelab/ample/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var RPCTimeout = time.Second * 1

const (
	GRPCAddress = "localhost:8880"
)

func connect() (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	return conn, err
}

func GetAllSources() ([]*pb.AllSourcesReply_SourceReply, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewAmpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	r, err := c.GetAllSources(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	if !r.Ok {
		return nil, errors.New("not ok")
	}
	return r.Sources, nil
}

func GetAllEvmChainClients() ([]*pb.AllEvmChainClientsReply_EvmChainClientReply, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewAmpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	r, err := c.GetAllEvmChainClients(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	if !r.Ok {
		return nil, errors.New("not ok")
	}
	return r.Clients, nil
}

func GetAllUniV2Dexs() ([]*pb.AllUniV2DexsReply_UniV2DexReply, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewAmpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	r, err := c.GetAllUniV2Dexs(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	if !r.Ok {
		return nil, errors.New("not ok")
	}
	return r.Dexs, nil
}

func GetAllTokens() ([]*pb.AllTokensReply_TokenReply, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewAmpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	r, err := c.GetAllTokens(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	if !r.Ok {
		return nil, errors.New("not ok")
	}
	return r.Tokens, nil
}

func GetAllSourcePositionControls() ([]*pb.AllSourcePositionControlsReply_SourcePositionControlReply, error) {
	conn, err := connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewAmpleClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), RPCTimeout)
	defer cancel()
	r, err := c.GetAllSourcePositionControls(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	if !r.Ok {
		return nil, errors.New("not ok")
	}
	return r.SourceControls, nil
}
