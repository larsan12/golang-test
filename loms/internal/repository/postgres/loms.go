package repository

import (
	"context"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/postgres/transactor"
	"route256/loms/internal/repository/schema"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
)

type LomsRepo struct {
	transactor.QueryEngineProvider
}

func NewLomsRepo(provider transactor.QueryEngineProvider) *LomsRepo {
	return &LomsRepo{
		QueryEngineProvider: provider,
	}
}

var (
	ordersColumns            = []string{"order_id", "user_id", "status", "updated_at", "created_at"}
	stocksReservationColumns = []string{"order_id", "sku", "warehouse_id", "count", "status"}
	stocksColumns            = []string{"warehouse_id", "sku", "count"}
	psql                     = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

const (
	ordersTable            = "orders"
	stocksReservationTable = "stocks_reservation"
	stocksTable            = "stocks"
)

const (
	OrderStatusNew      = "new"
	OrderStatusPaid     = "paid"
	OrderStatusCanceled = "cancelled"
	ReserveStatusPaid   = "paid"
)

func (r *LomsRepo) GetStockReservations(ctx context.Context, orderId int64) ([]domain.StockReservation, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var stockReservations []schema.StockReservation
	rawQuery, args, err := psql.Select(stocksReservationColumns...).
		From(stocksReservationTable).
		Where(sq.Eq{"order_id": orderId}).
		ToSql()

	if err != nil {
		return nil, err
	}
	errorr := pgxscan.Select(ctx, db, &stockReservations, rawQuery, args...)
	if errorr != nil {
		return nil, errorr
	}

	result := make([]domain.StockReservation, len(stockReservations))
	if err != nil {
		return nil, err
	}
	for i, stock := range stockReservations {
		result[i].OrderId = stock.OrderId
		result[i].Sku = stock.Sku
		result[i].WarehouseId = stock.WarehouseId
		result[i].Count = stock.Count
		result[i].Status = stock.Status
	}
	return result, nil
}

func (r *LomsRepo) UpdateStock(ctx context.Context, sku uint32, warehouseId int64, diff int64) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var operator sq.Sqlizer
	if diff > 0 {
		operator = sq.Expr("count+?", diff)
	} else {
		operator = sq.Expr("count-?", -diff)
	}
	rawQuery, args, err := psql.Update(stocksTable).
		Where(sq.Eq{"warehouse_id": warehouseId, "sku": sku}).
		Set("count", operator).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}

func (r *LomsRepo) CreateStockReservation(ctx context.Context, stockReserv domain.StockReservation) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Insert(stocksReservationTable).
		Columns("order_id", "sku", "warehouse_id", "count").
		Values(stockReserv.OrderId, stockReserv.Sku, stockReserv.WarehouseId, stockReserv.Count).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}

func (r *LomsRepo) DeleteStockReservation(ctx context.Context, stockReserv domain.StockReservation) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Delete(stocksReservationTable).
		Where(sq.Eq{"order_id": stockReserv.OrderId, "warehouse_id": stockReserv.WarehouseId, "sku": stockReserv.Sku}).
		ToSql()
	if err != nil {
		return err
	}
	_, err = db.Exec(ctx, rawQuery, args...)
	return err
}

func (r *LomsRepo) CreateOrder(ctx context.Context, user int64) (int64, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Insert(ordersTable).
		Columns("user_id", "status").
		Values(user, OrderStatusNew).
		Suffix("RETURNING order_id").
		ToSql()
	if err != nil {
		return 0, err
	}
	var order schema.Order
	err = pgxscan.Get(ctx, db, &order, rawQuery, args...)
	if err != nil {
		return 0, err
	}
	return order.OrderId, nil
}

func (r *LomsRepo) GetOrder(ctx context.Context, orderId int64) (domain.Order, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var order domain.Order
	var schemaOrder schema.Order
	rawQuery, args, err := psql.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"order_id": orderId}).
		ToSql()

	if err != nil {
		return order, err
	}
	errorr := pgxscan.Get(ctx, db, &schemaOrder, rawQuery, args...)
	if errorr != nil {
		return order, errorr
	}
	order.User = int64(schemaOrder.UserId)
	order.Status = schemaOrder.Status
	return order, errorr
}

func (r *LomsRepo) GetOrderItems(ctx context.Context, orderId int64) ([]domain.OrderItem, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var items []domain.OrderItem
	var schemaItems []schema.OrderItem
	rawQuery, args, err := psql.Select("order_id", "sku", "SUM(count) as count").
		From(stocksReservationTable).
		Where(sq.Eq{"order_id": orderId}).
		GroupBy("order_id", "sku").
		ToSql()

	if err != nil {
		return items, err
	}
	errorr := pgxscan.Select(ctx, db, &schemaItems, rawQuery, args...)
	items = make([]domain.OrderItem, len(schemaItems))
	for i, item := range schemaItems {
		items[i].Count = uint16(item.Count)
		items[i].Sku = item.Sku
	}
	return items, errorr
}

func (r *LomsRepo) GetStocks(ctx context.Context, sku uint32) ([]domain.StockItem, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var stocks []schema.Stock
	rawQuery, args, err := psql.Select(stocksColumns...).
		From(stocksTable).
		Where(sq.Eq{"sku": sku}).
		Where(sq.Gt{"count": 0}).
		ToSql()

	if err != nil {
		return nil, err
	}
	errorr := pgxscan.Select(ctx, db, &stocks, rawQuery, args...)
	if errorr != nil {
		return nil, errorr
	}

	result := make([]domain.StockItem, len(stocks))
	if err != nil {
		return nil, err
	}
	for i, stock := range stocks {
		result[i].Count = stock.Count
		result[i].WarehouseID = stock.WarehouseId
	}
	return result, nil
}

func (r *LomsRepo) OrderSetStatus(ctx context.Context, orderId int64, status string) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Update(ordersTable).
		Where(sq.Eq{"order_id": orderId}).
		Set("status", status).
		Set("updated_at", time.Now()).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}

func (r *LomsRepo) ReservSetStatuses(ctx context.Context, orderId int64, status string) error {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Update(stocksReservationTable).
		Where(sq.Eq{"order_id": orderId}).
		Set("status", status).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}

func (r *LomsRepo) GetOldOrders(ctx context.Context, dateFrom time.Time) ([]domain.Order, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var orders []schema.Order
	rawQuery, args, err := psql.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.LtOrEq{"updated_at": dateFrom}).
		ToSql()

	if err != nil {
		return nil, err
	}
	errorr := pgxscan.Select(ctx, db, &orders, rawQuery, args...)
	if errorr != nil {
		return nil, errorr
	}

	result := make([]domain.Order, len(orders))
	if err != nil {
		return nil, err
	}
	for i, order := range orders {
		result[i].OrderId = order.OrderId
		result[i].Status = order.Status
		result[i].User = int64(order.UserId)
	}
	return result, nil
}
