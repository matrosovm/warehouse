package examples

import (
	"context"
	"fmt"
	"net/rpc"

	wh "github.com/matrosovm/warehouse/pkg/api/warehouse"
)

type WarehouseClient interface {
	Reservation(context.Context, *wh.ReservationRequest) (*wh.ReservationResponse, error)
	ReleaseOfReserve(context.Context, *wh.ReleaseOfReservedRequest) (*wh.ReleaseOfReservedResponse, error)
	RemainingProducts(context.Context, *wh.RemainingProductsRequest) (*wh.RemainingProductsResponse, error)
}

type warehouseClient struct {
	client *rpc.Client
}

func NewWarehouseClient(client *rpc.Client) WarehouseClient {
	return &warehouseClient{client: client}
}

func (wC *warehouseClient) Reservation(
	ctx context.Context,
	req *wh.ReservationRequest,
) (*wh.ReservationResponse, error) {
	resp := &wh.ReservationResponse{}

	done := make(chan error)
	go func() {
		done <- wC.client.Call("Service.Reservation", req, resp)
	}()

	select {
	case err := <-done:
		return resp, err
	case <-ctx.Done():
		return nil, fmt.Errorf("RPC call timed out")
	}
}

func (wC *warehouseClient) ReleaseOfReserve(
	ctx context.Context,
	req *wh.ReleaseOfReservedRequest,
) (*wh.ReleaseOfReservedResponse, error) {
	resp := &wh.ReleaseOfReservedResponse{}

	done := make(chan error)
	go func() {
		done <- wC.client.Call("Service.ReleaseOfReserve", req, resp)
	}()

	select {
	case err := <-done:
		return resp, err
	case <-ctx.Done():
		return nil, fmt.Errorf("RPC call timed out")
	}
}

func (wC *warehouseClient) RemainingProducts(
	ctx context.Context,
	req *wh.RemainingProductsRequest,
) (*wh.RemainingProductsResponse, error) {
	resp := &wh.RemainingProductsResponse{}

	done := make(chan error)
	go func() {
		done <- wC.client.Call("Service.RemainingProducts", req, resp)
	}()

	select {
	case err := <-done:
		return resp, err
	case <-ctx.Done():
		return nil, fmt.Errorf("RPC call timed out")
	}
}

func helper(ctx context.Context, some any) error
