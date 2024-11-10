package connect

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type commandConnector struct{}

func NewCommandConnector() ServiceConnector {
	return &commandConnector{}
}

// CommandServiceとの接続を確立する
func (c *commandConnector) Connect() (*grpc.ClientConn, error) {
	server := "commandservice:8082" // Commandサービス名とポート番号
	// gRPCサーバーに接続する
	conn, err := grpc.Dial( // gRPCサーバとの接続を確立する
		server, // 接続文字列
		// TLSを使用せず、平文で通信する
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 接続が確率するまで待つか、失敗した場合にエラーを即座に取得する
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
