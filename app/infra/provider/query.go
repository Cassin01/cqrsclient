package provider

import (
	"cqrsclient/infra/connect"

	"github.com/Cassin01/samplepb/pb"
)

// Queryクライアントプロバイダ
type QueryClientProvider struct {
	Category pb.CategoryQueryClient // カテゴリ用クライアント
	Product  pb.ProductQueryClient  // 商品用クライアント
}

// コンストラクタ
func NewQueryClientProvider(connector connect.ServiceConnector) (*QueryClientProvider, error) {
	if connect, err := connector.Connect(); err != nil {
		return nil, err
	} else {
		// カテゴリ用クライアントを生成する
		category := pb.NewCategoryQueryClient(connect)
		// 商品用クライアントを生成する
		product := pb.NewProductQueryClient(connect)
		// プロバイダの生成
		return &QueryClientProvider{Category: category, Product: product}, nil
	}
}
