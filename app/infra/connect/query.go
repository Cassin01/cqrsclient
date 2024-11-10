package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type queryConnector struct{}

func NewQueryConnector() ServiceConnector {
	return &queryConnector{}
}

// QueryServiceの接続を確立する
func (q *queryConnector) Connect() (*grpc.ClientConn, error) {
	server := "queryservice:8083" // Queryサービス名とポート番号
	// gRPCサーバに接続する
	conn, err := grpc.Dial(
		server, // 接続文字列
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
