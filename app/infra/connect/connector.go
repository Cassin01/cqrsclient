package connect

import "google.golang.org/grpc"

// サーバーへの接続インターフェース
type ServiceConnector interface {
	// サーバーへ接続する
	Connect() (*grpc.ClientConn, error)
}
