package data

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/ChungKiet/cake-interview/ent"
	"github.com/ChungKiet/cake-interview/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	// init postgres driver
	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewCampaignRepo, NewVoucherRepo, NewVoucherCampaignRepo, NewUserVoucherRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(conf *conf.Data) (*Data, func(), error) {
	log := log.NewHelper(log.DefaultLogger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to db: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

// NewData .

func (d *Data) GetClient(ctx context.Context) *ent.Client {
	if tx, ok := d.TxFromContext(ctx); ok {
		return tx.Client()
	}

	return d.db
}

func (d *Data) NewTxContext(ctx context.Context, tx *ent.Tx) context.Context {
	return context.WithValue(ctx, "entTransaction", tx)
}

func (d *Data) TxFromContext(ctx context.Context) (tx *ent.Tx, ok bool) {
	tx, ok = ctx.Value("entTransaction").(*ent.Tx)
	return
}

func (d *Data) TxCompleted(ctx context.Context) context.Context {
	return context.WithValue(ctx, "entTransaction", nil)
}

func (r *Data) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, ok := r.TxFromContext(ctx)
	if !ok || tx == nil {
		ntx, err := r.GetClient(ctx).Tx(ctx)
		ctx = r.NewTxContext(ctx, tx)
		if err != nil {
			return err
		}
		tx = ntx
	}
	// tx, err := r.db.Tx(ctx)
	// if err != nil {
	// 	return err
	// }
	newCtx := r.NewTxContext(ctx, tx)
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(newCtx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
