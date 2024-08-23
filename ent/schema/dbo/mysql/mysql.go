package mysql

import (
	"context"
	ent_o "entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"kratos-use/ent"
	"kratos-use/ent/intercept"
	"time"
)

type contextTxKey struct{}

type BaseDBData struct {
	db *ent.Client
}

func NewBaseDBData(db *ent.Client) *BaseDBData {
	return &BaseDBData{
		db: db,
	}
}

func Open(source string) (*ent.Client, error) {
	drv, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	log.Info("mysql连接成功")
	// 获取数据库驱动中的sql.DB对象。
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv), ent.Debug())
	client.Intercept(
		intercept.Func(func(ctx context.Context, q intercept.Query) error {
			if ent_o.QueryFromContext(ctx).Limit == nil {
				q.Limit(500)
			}
			return nil
		}),
	)
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()

			defer func() {
				log.Context(ctx).Infof("数据库操作日志: Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()

			switch m.Op() {
			case ent.OpCreate:
				err := m.SetField("created_at", time.Now().Unix())
				if err != nil {
					log.Context(ctx).Errorw("msg", "设置创建时间字段失败", "err", err)
					return nil, err
				}

			case ent.OpUpdateOne, ent.OpUpdate:
				err := m.SetField("updated_at", time.Now().Unix())
				if err != nil {
					log.Context(ctx).Errorw("msg", "设置更新时间字段失败", "err", err)
					return nil, err
				}
			case ent.OpDeleteOne, ent.OpDelete:
				err := m.SetField("deleted_at", time.Now().Unix())
				if err != nil {
					log.Context(ctx).Errorw("msg", "设置删除时间字段失败", "err", err)
					return nil, err
				}
			}

			return next.Mutate(ctx, m)
		})
	})

	return client, nil
}

func (b *BaseDBData) DB(ctx context.Context) *ent.Client {
	tx, ok := ctx.Value(contextTxKey{}).(*ent.Tx)
	if ok {
		return tx.Client()
	}

	return b.db
}

func (b *BaseDBData) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	log.Info("begin transaction")

	var (
		err error
		tx  *ent.Tx
	)
	tx, err = b.db.Tx(ctx)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, contextTxKey{}, tx)

	defer func() {
		if v := recover(); v != nil {
			log.Warnw("msg", "transaction panic", "recover error", err)

			err1 := tx.Rollback().Error
			if err1 != nil {
				log.Warnw("msg", "rollback transaction failed", "err", err, "err1", err1)
			} else {
				log.Info("rollback transaction successfully")
			}
		}
	}()

	if err := fn(ctx); err != nil {
		log.Errorw("msg", "func execute failed", "err", err)

		if rerr := tx.Rollback(); rerr != nil {
			log.Errorw("msg", "rolling back transaction error", "err", err)
		} else {
			log.Info("rolling back transaction successfully")
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	log.Info("commit transaction successfully")

	return nil
}
