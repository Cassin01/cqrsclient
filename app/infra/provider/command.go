package provider

import (
	"cqrsclient/infra/connect"

	"github.com/Cassin01/samplepb/pb"
)

// Commandクライアントプロバイダ
type CommandClientProvider struct {
	Category pb.CategoryCommandClient
	Product  pb.ProductCommandClient
}

// コンストラクタ
func NewCommandClientProvider(connector connect.ServiceConnector) (*CommandClientProvider, error) {
	// Command Serviceの接続
	connect, err := connector.Connect()
	if err != nil {
		return nil, err
	}

	// カテゴリ用クライアントを作成する
	category := pb.NewCategoryCommandClient(connect)

	// 商品用クライアントを作成する
	product := pb.NewProductCommandClient(connect)

	// プロバイダの生成
	return &CommandClientProvider{Category: category, Product: product}, nil
}
