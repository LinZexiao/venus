// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/filecoin-project/venus/venus-shared/api/market/client (interfaces: IMarketClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	address "github.com/filecoin-project/go-address"
	datatransfer "github.com/filecoin-project/go-data-transfer"
	retrievalmarket "github.com/filecoin-project/go-fil-markets/retrievalmarket"
	big "github.com/filecoin-project/go-state-types/big"
	internal "github.com/filecoin-project/venus/venus-shared/internal"
	types "github.com/filecoin-project/venus/venus-shared/types"
	market "github.com/filecoin-project/venus/venus-shared/types/market"
	client "github.com/filecoin-project/venus/venus-shared/types/market/client"
	gomock "github.com/golang/mock/gomock"
	cid "github.com/ipfs/go-cid"
	peer "github.com/libp2p/go-libp2p-core/peer"
)

// MockIMarketClient is a mock of IMarketClient interface.
type MockIMarketClient struct {
	ctrl     *gomock.Controller
	recorder *MockIMarketClientMockRecorder
}

// MockIMarketClientMockRecorder is the mock recorder for MockIMarketClient.
type MockIMarketClientMockRecorder struct {
	mock *MockIMarketClient
}

// NewMockIMarketClient creates a new mock instance.
func NewMockIMarketClient(ctrl *gomock.Controller) *MockIMarketClient {
	mock := &MockIMarketClient{ctrl: ctrl}
	mock.recorder = &MockIMarketClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMarketClient) EXPECT() *MockIMarketClientMockRecorder {
	return m.recorder
}

// ClientCalcCommP mocks base method.
func (m *MockIMarketClient) ClientCalcCommP(arg0 context.Context, arg1 string) (*client.CommPRet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientCalcCommP", arg0, arg1)
	ret0, _ := ret[0].(*client.CommPRet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientCalcCommP indicates an expected call of ClientCalcCommP.
func (mr *MockIMarketClientMockRecorder) ClientCalcCommP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientCalcCommP", reflect.TypeOf((*MockIMarketClient)(nil).ClientCalcCommP), arg0, arg1)
}

// ClientCancelDataTransfer mocks base method.
func (m *MockIMarketClient) ClientCancelDataTransfer(arg0 context.Context, arg1 datatransfer.TransferID, arg2 peer.ID, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientCancelDataTransfer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientCancelDataTransfer indicates an expected call of ClientCancelDataTransfer.
func (mr *MockIMarketClientMockRecorder) ClientCancelDataTransfer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientCancelDataTransfer", reflect.TypeOf((*MockIMarketClient)(nil).ClientCancelDataTransfer), arg0, arg1, arg2, arg3)
}

// ClientCancelRetrievalDeal mocks base method.
func (m *MockIMarketClient) ClientCancelRetrievalDeal(arg0 context.Context, arg1 retrievalmarket.DealID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientCancelRetrievalDeal", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientCancelRetrievalDeal indicates an expected call of ClientCancelRetrievalDeal.
func (mr *MockIMarketClientMockRecorder) ClientCancelRetrievalDeal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientCancelRetrievalDeal", reflect.TypeOf((*MockIMarketClient)(nil).ClientCancelRetrievalDeal), arg0, arg1)
}

// ClientDataTransferUpdates mocks base method.
func (m *MockIMarketClient) ClientDataTransferUpdates(arg0 context.Context) (<-chan market.DataTransferChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientDataTransferUpdates", arg0)
	ret0, _ := ret[0].(<-chan market.DataTransferChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientDataTransferUpdates indicates an expected call of ClientDataTransferUpdates.
func (mr *MockIMarketClientMockRecorder) ClientDataTransferUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientDataTransferUpdates", reflect.TypeOf((*MockIMarketClient)(nil).ClientDataTransferUpdates), arg0)
}

// ClientDealPieceCID mocks base method.
func (m *MockIMarketClient) ClientDealPieceCID(arg0 context.Context, arg1 cid.Cid) (client.DataCIDSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientDealPieceCID", arg0, arg1)
	ret0, _ := ret[0].(client.DataCIDSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientDealPieceCID indicates an expected call of ClientDealPieceCID.
func (mr *MockIMarketClientMockRecorder) ClientDealPieceCID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientDealPieceCID", reflect.TypeOf((*MockIMarketClient)(nil).ClientDealPieceCID), arg0, arg1)
}

// ClientDealSize mocks base method.
func (m *MockIMarketClient) ClientDealSize(arg0 context.Context, arg1 cid.Cid) (client.DataSize, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientDealSize", arg0, arg1)
	ret0, _ := ret[0].(client.DataSize)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientDealSize indicates an expected call of ClientDealSize.
func (mr *MockIMarketClientMockRecorder) ClientDealSize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientDealSize", reflect.TypeOf((*MockIMarketClient)(nil).ClientDealSize), arg0, arg1)
}

// ClientExport mocks base method.
func (m *MockIMarketClient) ClientExport(arg0 context.Context, arg1 client.ExportRef, arg2 client.FileRef) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientExport", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientExport indicates an expected call of ClientExport.
func (mr *MockIMarketClientMockRecorder) ClientExport(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientExport", reflect.TypeOf((*MockIMarketClient)(nil).ClientExport), arg0, arg1, arg2)
}

// ClientFindData mocks base method.
func (m *MockIMarketClient) ClientFindData(arg0 context.Context, arg1 cid.Cid, arg2 *cid.Cid) ([]client.QueryOffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientFindData", arg0, arg1, arg2)
	ret0, _ := ret[0].([]client.QueryOffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientFindData indicates an expected call of ClientFindData.
func (mr *MockIMarketClientMockRecorder) ClientFindData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientFindData", reflect.TypeOf((*MockIMarketClient)(nil).ClientFindData), arg0, arg1, arg2)
}

// ClientGenCar mocks base method.
func (m *MockIMarketClient) ClientGenCar(arg0 context.Context, arg1 client.FileRef, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGenCar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientGenCar indicates an expected call of ClientGenCar.
func (mr *MockIMarketClientMockRecorder) ClientGenCar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGenCar", reflect.TypeOf((*MockIMarketClient)(nil).ClientGenCar), arg0, arg1, arg2)
}

// ClientGetDealInfo mocks base method.
func (m *MockIMarketClient) ClientGetDealInfo(arg0 context.Context, arg1 cid.Cid) (*client.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetDealInfo", arg0, arg1)
	ret0, _ := ret[0].(*client.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientGetDealInfo indicates an expected call of ClientGetDealInfo.
func (mr *MockIMarketClientMockRecorder) ClientGetDealInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetDealInfo", reflect.TypeOf((*MockIMarketClient)(nil).ClientGetDealInfo), arg0, arg1)
}

// ClientGetDealStatus mocks base method.
func (m *MockIMarketClient) ClientGetDealStatus(arg0 context.Context, arg1 uint64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetDealStatus", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientGetDealStatus indicates an expected call of ClientGetDealStatus.
func (mr *MockIMarketClientMockRecorder) ClientGetDealStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetDealStatus", reflect.TypeOf((*MockIMarketClient)(nil).ClientGetDealStatus), arg0, arg1)
}

// ClientGetDealUpdates mocks base method.
func (m *MockIMarketClient) ClientGetDealUpdates(arg0 context.Context) (<-chan client.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetDealUpdates", arg0)
	ret0, _ := ret[0].(<-chan client.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientGetDealUpdates indicates an expected call of ClientGetDealUpdates.
func (mr *MockIMarketClientMockRecorder) ClientGetDealUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetDealUpdates", reflect.TypeOf((*MockIMarketClient)(nil).ClientGetDealUpdates), arg0)
}

// ClientGetRetrievalUpdates mocks base method.
func (m *MockIMarketClient) ClientGetRetrievalUpdates(arg0 context.Context) (<-chan client.RetrievalInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientGetRetrievalUpdates", arg0)
	ret0, _ := ret[0].(<-chan client.RetrievalInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientGetRetrievalUpdates indicates an expected call of ClientGetRetrievalUpdates.
func (mr *MockIMarketClientMockRecorder) ClientGetRetrievalUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientGetRetrievalUpdates", reflect.TypeOf((*MockIMarketClient)(nil).ClientGetRetrievalUpdates), arg0)
}

// ClientHasLocal mocks base method.
func (m *MockIMarketClient) ClientHasLocal(arg0 context.Context, arg1 cid.Cid) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientHasLocal", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientHasLocal indicates an expected call of ClientHasLocal.
func (mr *MockIMarketClientMockRecorder) ClientHasLocal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientHasLocal", reflect.TypeOf((*MockIMarketClient)(nil).ClientHasLocal), arg0, arg1)
}

// ClientImport mocks base method.
func (m *MockIMarketClient) ClientImport(arg0 context.Context, arg1 client.FileRef) (*client.ImportRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientImport", arg0, arg1)
	ret0, _ := ret[0].(*client.ImportRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientImport indicates an expected call of ClientImport.
func (mr *MockIMarketClientMockRecorder) ClientImport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientImport", reflect.TypeOf((*MockIMarketClient)(nil).ClientImport), arg0, arg1)
}

// ClientListDataTransfers mocks base method.
func (m *MockIMarketClient) ClientListDataTransfers(arg0 context.Context) ([]market.DataTransferChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientListDataTransfers", arg0)
	ret0, _ := ret[0].([]market.DataTransferChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientListDataTransfers indicates an expected call of ClientListDataTransfers.
func (mr *MockIMarketClientMockRecorder) ClientListDataTransfers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientListDataTransfers", reflect.TypeOf((*MockIMarketClient)(nil).ClientListDataTransfers), arg0)
}

// ClientListDeals mocks base method.
func (m *MockIMarketClient) ClientListDeals(arg0 context.Context) ([]client.DealInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientListDeals", arg0)
	ret0, _ := ret[0].([]client.DealInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientListDeals indicates an expected call of ClientListDeals.
func (mr *MockIMarketClientMockRecorder) ClientListDeals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientListDeals", reflect.TypeOf((*MockIMarketClient)(nil).ClientListDeals), arg0)
}

// ClientListImports mocks base method.
func (m *MockIMarketClient) ClientListImports(arg0 context.Context) ([]client.Import, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientListImports", arg0)
	ret0, _ := ret[0].([]client.Import)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientListImports indicates an expected call of ClientListImports.
func (mr *MockIMarketClientMockRecorder) ClientListImports(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientListImports", reflect.TypeOf((*MockIMarketClient)(nil).ClientListImports), arg0)
}

// ClientListRetrievals mocks base method.
func (m *MockIMarketClient) ClientListRetrievals(arg0 context.Context) ([]client.RetrievalInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientListRetrievals", arg0)
	ret0, _ := ret[0].([]client.RetrievalInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientListRetrievals indicates an expected call of ClientListRetrievals.
func (mr *MockIMarketClientMockRecorder) ClientListRetrievals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientListRetrievals", reflect.TypeOf((*MockIMarketClient)(nil).ClientListRetrievals), arg0)
}

// ClientMinerQueryOffer mocks base method.
func (m *MockIMarketClient) ClientMinerQueryOffer(arg0 context.Context, arg1 address.Address, arg2 cid.Cid, arg3 *cid.Cid) (client.QueryOffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientMinerQueryOffer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(client.QueryOffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientMinerQueryOffer indicates an expected call of ClientMinerQueryOffer.
func (mr *MockIMarketClientMockRecorder) ClientMinerQueryOffer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientMinerQueryOffer", reflect.TypeOf((*MockIMarketClient)(nil).ClientMinerQueryOffer), arg0, arg1, arg2, arg3)
}

// ClientQueryAsk mocks base method.
func (m *MockIMarketClient) ClientQueryAsk(arg0 context.Context, arg1 peer.ID, arg2 address.Address) (*client.StorageAsk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientQueryAsk", arg0, arg1, arg2)
	ret0, _ := ret[0].(*client.StorageAsk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientQueryAsk indicates an expected call of ClientQueryAsk.
func (mr *MockIMarketClientMockRecorder) ClientQueryAsk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientQueryAsk", reflect.TypeOf((*MockIMarketClient)(nil).ClientQueryAsk), arg0, arg1, arg2)
}

// ClientRemoveImport mocks base method.
func (m *MockIMarketClient) ClientRemoveImport(arg0 context.Context, arg1 client.ImportID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRemoveImport", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRemoveImport indicates an expected call of ClientRemoveImport.
func (mr *MockIMarketClientMockRecorder) ClientRemoveImport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRemoveImport", reflect.TypeOf((*MockIMarketClient)(nil).ClientRemoveImport), arg0, arg1)
}

// ClientRestartDataTransfer mocks base method.
func (m *MockIMarketClient) ClientRestartDataTransfer(arg0 context.Context, arg1 datatransfer.TransferID, arg2 peer.ID, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRestartDataTransfer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRestartDataTransfer indicates an expected call of ClientRestartDataTransfer.
func (mr *MockIMarketClientMockRecorder) ClientRestartDataTransfer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRestartDataTransfer", reflect.TypeOf((*MockIMarketClient)(nil).ClientRestartDataTransfer), arg0, arg1, arg2, arg3)
}

// ClientRetrieve mocks base method.
func (m *MockIMarketClient) ClientRetrieve(arg0 context.Context, arg1 client.RetrievalOrder) (*client.RestrievalRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRetrieve", arg0, arg1)
	ret0, _ := ret[0].(*client.RestrievalRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientRetrieve indicates an expected call of ClientRetrieve.
func (mr *MockIMarketClientMockRecorder) ClientRetrieve(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRetrieve", reflect.TypeOf((*MockIMarketClient)(nil).ClientRetrieve), arg0, arg1)
}

// ClientRetrieveTryRestartInsufficientFunds mocks base method.
func (m *MockIMarketClient) ClientRetrieveTryRestartInsufficientFunds(arg0 context.Context, arg1 address.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRetrieveTryRestartInsufficientFunds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRetrieveTryRestartInsufficientFunds indicates an expected call of ClientRetrieveTryRestartInsufficientFunds.
func (mr *MockIMarketClientMockRecorder) ClientRetrieveTryRestartInsufficientFunds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRetrieveTryRestartInsufficientFunds", reflect.TypeOf((*MockIMarketClient)(nil).ClientRetrieveTryRestartInsufficientFunds), arg0, arg1)
}

// ClientRetrieveWait mocks base method.
func (m *MockIMarketClient) ClientRetrieveWait(arg0 context.Context, arg1 retrievalmarket.DealID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientRetrieveWait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClientRetrieveWait indicates an expected call of ClientRetrieveWait.
func (mr *MockIMarketClientMockRecorder) ClientRetrieveWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientRetrieveWait", reflect.TypeOf((*MockIMarketClient)(nil).ClientRetrieveWait), arg0, arg1)
}

// ClientStartDeal mocks base method.
func (m *MockIMarketClient) ClientStartDeal(arg0 context.Context, arg1 *client.StartDealParams) (*cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientStartDeal", arg0, arg1)
	ret0, _ := ret[0].(*cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientStartDeal indicates an expected call of ClientStartDeal.
func (mr *MockIMarketClientMockRecorder) ClientStartDeal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientStartDeal", reflect.TypeOf((*MockIMarketClient)(nil).ClientStartDeal), arg0, arg1)
}

// ClientStatelessDeal mocks base method.
func (m *MockIMarketClient) ClientStatelessDeal(arg0 context.Context, arg1 *client.StartDealParams) (*cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientStatelessDeal", arg0, arg1)
	ret0, _ := ret[0].(*cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClientStatelessDeal indicates an expected call of ClientStatelessDeal.
func (mr *MockIMarketClientMockRecorder) ClientStatelessDeal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientStatelessDeal", reflect.TypeOf((*MockIMarketClient)(nil).ClientStatelessDeal), arg0, arg1)
}

// DefaultAddress mocks base method.
func (m *MockIMarketClient) DefaultAddress(arg0 context.Context) (address.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultAddress", arg0)
	ret0, _ := ret[0].(address.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DefaultAddress indicates an expected call of DefaultAddress.
func (mr *MockIMarketClientMockRecorder) DefaultAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultAddress", reflect.TypeOf((*MockIMarketClient)(nil).DefaultAddress), arg0)
}

// MarketAddBalance mocks base method.
func (m *MockIMarketClient) MarketAddBalance(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketAddBalance", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketAddBalance indicates an expected call of MarketAddBalance.
func (mr *MockIMarketClientMockRecorder) MarketAddBalance(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketAddBalance", reflect.TypeOf((*MockIMarketClient)(nil).MarketAddBalance), arg0, arg1, arg2, arg3)
}

// MarketGetReserved mocks base method.
func (m *MockIMarketClient) MarketGetReserved(arg0 context.Context, arg1 address.Address) (big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketGetReserved", arg0, arg1)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketGetReserved indicates an expected call of MarketGetReserved.
func (mr *MockIMarketClientMockRecorder) MarketGetReserved(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketGetReserved", reflect.TypeOf((*MockIMarketClient)(nil).MarketGetReserved), arg0, arg1)
}

// MarketReleaseFunds mocks base method.
func (m *MockIMarketClient) MarketReleaseFunds(arg0 context.Context, arg1 address.Address, arg2 big.Int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketReleaseFunds", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarketReleaseFunds indicates an expected call of MarketReleaseFunds.
func (mr *MockIMarketClientMockRecorder) MarketReleaseFunds(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketReleaseFunds", reflect.TypeOf((*MockIMarketClient)(nil).MarketReleaseFunds), arg0, arg1, arg2)
}

// MarketReserveFunds mocks base method.
func (m *MockIMarketClient) MarketReserveFunds(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketReserveFunds", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketReserveFunds indicates an expected call of MarketReserveFunds.
func (mr *MockIMarketClientMockRecorder) MarketReserveFunds(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketReserveFunds", reflect.TypeOf((*MockIMarketClient)(nil).MarketReserveFunds), arg0, arg1, arg2, arg3)
}

// MarketWithdraw mocks base method.
func (m *MockIMarketClient) MarketWithdraw(arg0 context.Context, arg1, arg2 address.Address, arg3 big.Int) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarketWithdraw", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarketWithdraw indicates an expected call of MarketWithdraw.
func (mr *MockIMarketClientMockRecorder) MarketWithdraw(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarketWithdraw", reflect.TypeOf((*MockIMarketClient)(nil).MarketWithdraw), arg0, arg1, arg2, arg3)
}

// MessagerGetMessage mocks base method.
func (m *MockIMarketClient) MessagerGetMessage(arg0 context.Context, arg1 cid.Cid) (*internal.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessagerGetMessage", arg0, arg1)
	ret0, _ := ret[0].(*internal.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessagerGetMessage indicates an expected call of MessagerGetMessage.
func (mr *MockIMarketClientMockRecorder) MessagerGetMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagerGetMessage", reflect.TypeOf((*MockIMarketClient)(nil).MessagerGetMessage), arg0, arg1)
}

// MessagerPushMessage mocks base method.
func (m *MockIMarketClient) MessagerPushMessage(arg0 context.Context, arg1 *internal.Message, arg2 *types.MessageSendSpec) (cid.Cid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessagerPushMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(cid.Cid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessagerPushMessage indicates an expected call of MessagerPushMessage.
func (mr *MockIMarketClientMockRecorder) MessagerPushMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagerPushMessage", reflect.TypeOf((*MockIMarketClient)(nil).MessagerPushMessage), arg0, arg1, arg2)
}

// MessagerWaitMessage mocks base method.
func (m *MockIMarketClient) MessagerWaitMessage(arg0 context.Context, arg1 cid.Cid) (*types.MsgLookup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessagerWaitMessage", arg0, arg1)
	ret0, _ := ret[0].(*types.MsgLookup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessagerWaitMessage indicates an expected call of MessagerWaitMessage.
func (mr *MockIMarketClientMockRecorder) MessagerWaitMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessagerWaitMessage", reflect.TypeOf((*MockIMarketClient)(nil).MessagerWaitMessage), arg0, arg1)
}
