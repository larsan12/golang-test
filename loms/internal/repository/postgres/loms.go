package repository

import (
	"context"
	"errors"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/postgres/transactor"
	"route256/loms/internal/repository/schema"

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
	ordersColumns = []string{"order_id", "user_id", "status"}
	stocksReservationColumns = []string{"order_id", "sku", "warehouse_id", "count", "status"}
	stocksColumns = []string{"warehouse_id", "sku", "count"}
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

const (
	ordersTable = "orders"
	stocksReservationTable = "stocks_reservation"
	stocksTable = "stocks"
)

const (
	OrderStatusNew = "new"
	OrderStatusPaid = "paid"
	OrderStatusCanceled = "cancelled"
	ReserveStatusPaid = "paid"
)

func (r *LomsRepo) GetStocks(ctx context.Context, sku uint32) ([]schema.Stock, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var stocks []schema.Stock
	rawQuery, args, err := psql.Select(stocksColumns...).
		From(stocksTable).
		Where(sq.Eq{"sku": sku}).
		Where(sq.Gt{"count": 0}).
		ToSql()
	
	if err != nil {
		return stocks, err
	}
	errorr := pgxscan.Select(ctx, db, &stocks, rawQuery, args...)
	return stocks, errorr
}

func (r *LomsRepo) GetStockReservations(ctx context.Context, orderId int64) ([]schema.StockReservation, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var stockReservations []schema.StockReservation
	rawQuery, args, err := psql.Select(stocksReservationColumns...).
		From(stocksReservationTable).
		Where(sq.Eq{"order_id": orderId}).
		ToSql()
	
	if err != nil {
		return stockReservations, err
	}
	errorr := pgxscan.Select(ctx, db, &stockReservations, rawQuery, args...)
	return stockReservations, errorr
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

func (r *LomsRepo) CreateStockReservation(ctx context.Context, stockReserv schema.StockReservation) error {
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

func (r *LomsRepo) DeleteStockReservation(ctx context.Context, stockReserv schema.StockReservation) error {
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


func (r *LomsRepo) GetOrder(ctx context.Context, orderId int64) (schema.Order, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var order schema.Order
	rawQuery, args, err := psql.Select(ordersColumns...).
		From(ordersTable).
		Where(sq.Eq{"order_id": orderId}).
		ToSql()
	
	if err != nil {
		return order, err
	}
	errorr := pgxscan.Get(ctx, db, &order, rawQuery, args...)
	return order, errorr
}


func (r *LomsRepo) GetOrderItems(ctx context.Context, orderId int64) ([]schema.OrderItem, error) {
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	var items []schema.OrderItem
	rawQuery, args, err := psql.Select("order_id", "sku", "SUM(count) as count").
		From(stocksReservationTable).
		Where(sq.Eq{"order_id": orderId}).
		GroupBy("order_id", "sku").
		ToSql()
	
	if err != nil {
		return items, err
	}
	errorr := pgxscan.Select(ctx, db, &items, rawQuery, args...)
	return items, errorr
}


func (r *LomsRepo) MakeOrder(ctx context.Context, user int64, items []domain.OrderItem) (int64, error) {
	orderId, err := r.CreateOrder(ctx, user)
	if err != nil {
		return orderId, err
	}

	for _, item := range items {
		stocks, err := r.GetStocks(ctx, item.Sku)
		if err != nil {
			return orderId, err;
		}
		restCount := item.Count
		for _, stock := range stocks {
			var reservation schema.StockReservation = schema.StockReservation{
				OrderId: orderId,
				WarehouseId: stock.WarehouseId,
				Sku: stock.Sku,
			}
			reservation.WarehouseId = stock.WarehouseId
			if stock.Count >= uint64(restCount) {
				reservation.Count = uint32(restCount)
				restCount = 0
			} else {
				restCount = restCount - uint16(stock.Count)
				reservation.Count = uint32(stock.Count)
			}
			err = r.CreateStockReservation(ctx, reservation)
			if err != nil {
				return 0, err
			}
			stock.Count -= uint64(reservation.Count)
			err = r.UpdateStock(ctx, stock.Sku, stock.WarehouseId, -int64(reservation.Count))
			if err != nil {
				return 0, err
			}
			if restCount == 0 {
				break;
			}
		}
		if restCount != 0 {
			return 0, errors.New("not enough stocks")
		}
	}
	return orderId, nil
}

func (r *LomsRepo) Stock(ctx context.Context, sku uint32) ([]domain.StockItem, error) {
	stocks, err := r.GetStocks(ctx, sku)
	if err != nil {
		return nil, err;
	}
	var result []domain.StockItem = make([]domain.StockItem, len(stocks))
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
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}

func (r *LomsRepo) PaidOrder(ctx context.Context, orderId int64) error {
	err := r.OrderSetStatus(ctx, orderId, OrderStatusPaid)
	if err != nil {
		return err
	}
	db := r.QueryEngineProvider.GetQueryEngine(ctx)
	rawQuery, args, err := psql.Update(stocksReservationTable).
		Where(sq.Eq{"order_id": orderId}).
		Set("status", ReserveStatusPaid).
		ToSql()
	if err != nil {
		return err
	}
	if _, err := db.Exec(ctx, rawQuery, args...); err != nil {
		return err
	}
	return nil
}

func (r *LomsRepo) ListOrder(ctx context.Context, orderId int64) (domain.Order, error) {
	var order domain.Order
	repoOrder, err := r.GetOrder(ctx, orderId)
	if err != nil {
		return order, err
	}
	order.Status = repoOrder.Status
	order.User = int64(repoOrder.UserId)
	orderItems, err := r.GetOrderItems(ctx, orderId)
	if err != nil {
		return order, err
	}
	items := make([]domain.OrderItem, len(orderItems))
	for i, item := range orderItems {
		items[i].Count = uint16(item.Count)
		items[i].Sku = item.Sku
	}
	order.Items = items
	return order, nil
}

func (r *LomsRepo) CancelOrder(ctx context.Context, orderId int64) error {
	err := r.OrderSetStatus(ctx, orderId, OrderStatusCanceled)
	if err != nil {
		return err
	}
	reservations, err := r.GetStockReservations(ctx, orderId)
	if err != nil {
		return err
	}

	for _, reserv := range reservations {
		err = r.UpdateStock(ctx, reserv.Sku, reserv.WarehouseId, int64(reserv.Count))
		if err != nil {
			return err
		}
		err = r.DeleteStockReservation(ctx, reserv)
		if err != nil {
			return err
		}
	}
	return nil
}