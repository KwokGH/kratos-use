package data

import (
	"context"
	"kratos-use/ent"
	"kratos-use/ent/schema/dbo/mysql"
	"kratos-use/internal/biz"
	"kratos-use/internal/conf"
	"kratos-use/pkg/cache"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewDB,
	mysql.NewBaseDBData,
	NewTransaction,
	NewData,
	NewUserRepo,
	NewDiaryRepo,
)

// Data .
type Data struct {
	*mysql.BaseDBData
	cache     cache.ICache
	bootstrap *conf.Bootstrap
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(
	bootstrap *conf.Bootstrap,
	baseDB *mysql.BaseDBData,
) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}

	return &Data{
		BaseDBData: baseDB,
		bootstrap:  bootstrap,
	}, cleanup, nil
}

// NewDB 数据库
func NewDB(bootstrap *conf.Bootstrap) (*ent.Client, error) {
	c := bootstrap.Data

	client, err := mysql.Open(c.Database.Source)
	if err != nil {
		log.Infow("msg", "open mysql error", "err", err)
		return nil, err
	}

	// 自动迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Infow("msg", "failed creating schema resources", "err", err)
		return nil, err
	}

	return client, nil
}
