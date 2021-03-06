// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/go-iost/rpc/pb (interfaces: ApiServiceServer)

// Package main is a generated GoMock package.
package main

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	pb "github.com/iost-official/go-iost/rpc/pb"
	reflect "reflect"
)

// MockApiServiceServer is a mock of ApiServiceServer interface
type MockApiServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockApiServiceServerMockRecorder
}

// MockApiServiceServerMockRecorder is the mock recorder for MockApiServiceServer
type MockApiServiceServerMockRecorder struct {
	mock *MockApiServiceServer
}

// NewMockApiServiceServer creates a new mock instance
func NewMockApiServiceServer(ctrl *gomock.Controller) *MockApiServiceServer {
	mock := &MockApiServiceServer{ctrl: ctrl}
	mock.recorder = &MockApiServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiServiceServer) EXPECT() *MockApiServiceServerMockRecorder {
	return m.recorder
}

// ExecTransaction mocks base method
func (m *MockApiServiceServer) ExecTransaction(arg0 context.Context, arg1 *pb.TransactionRequest) (*pb.TxReceipt, error) {
	ret := m.ctrl.Call(m, "ExecTransaction", arg0, arg1)
	ret0, _ := ret[0].(*pb.TxReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecTransaction indicates an expected call of ExecTransaction
func (mr *MockApiServiceServerMockRecorder) ExecTransaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTransaction", reflect.TypeOf((*MockApiServiceServer)(nil).ExecTransaction), arg0, arg1)
}

// GetAccount mocks base method
func (m *MockApiServiceServer) GetAccount(arg0 context.Context, arg1 *pb.GetAccountRequest) (*pb.Account, error) {
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(*pb.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount
func (mr *MockApiServiceServerMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockApiServiceServer)(nil).GetAccount), arg0, arg1)
}

// GetBatchContractStorage mocks base method
func (m *MockApiServiceServer) GetBatchContractStorage(arg0 context.Context, arg1 *pb.GetBatchContractStorageRequest) (*pb.GetBatchContractStorageResponse, error) {
	ret := m.ctrl.Call(m, "GetBatchContractStorage", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetBatchContractStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBatchContractStorage indicates an expected call of GetBatchContractStorage
func (mr *MockApiServiceServerMockRecorder) GetBatchContractStorage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBatchContractStorage", reflect.TypeOf((*MockApiServiceServer)(nil).GetBatchContractStorage), arg0, arg1)
}

// GetBlockByHash mocks base method
func (m *MockApiServiceServer) GetBlockByHash(arg0 context.Context, arg1 *pb.GetBlockByHashRequest) (*pb.BlockResponse, error) {
	ret := m.ctrl.Call(m, "GetBlockByHash", arg0, arg1)
	ret0, _ := ret[0].(*pb.BlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByHash indicates an expected call of GetBlockByHash
func (mr *MockApiServiceServerMockRecorder) GetBlockByHash(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByHash", reflect.TypeOf((*MockApiServiceServer)(nil).GetBlockByHash), arg0, arg1)
}

// GetBlockByNumber mocks base method
func (m *MockApiServiceServer) GetBlockByNumber(arg0 context.Context, arg1 *pb.GetBlockByNumberRequest) (*pb.BlockResponse, error) {
	ret := m.ctrl.Call(m, "GetBlockByNumber", arg0, arg1)
	ret0, _ := ret[0].(*pb.BlockResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockByNumber indicates an expected call of GetBlockByNumber
func (mr *MockApiServiceServerMockRecorder) GetBlockByNumber(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockByNumber", reflect.TypeOf((*MockApiServiceServer)(nil).GetBlockByNumber), arg0, arg1)
}

// GetCandidateBonus mocks base method
func (m *MockApiServiceServer) GetCandidateBonus(arg0 context.Context, arg1 *pb.GetAccountRequest) (*pb.CandidateBonus, error) {
	ret := m.ctrl.Call(m, "GetCandidateBonus", arg0, arg1)
	ret0, _ := ret[0].(*pb.CandidateBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandidateBonus indicates an expected call of GetCandidateBonus
func (mr *MockApiServiceServerMockRecorder) GetCandidateBonus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandidateBonus", reflect.TypeOf((*MockApiServiceServer)(nil).GetCandidateBonus), arg0, arg1)
}

// GetChainInfo mocks base method
func (m *MockApiServiceServer) GetChainInfo(arg0 context.Context, arg1 *pb.EmptyRequest) (*pb.ChainInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetChainInfo", arg0, arg1)
	ret0, _ := ret[0].(*pb.ChainInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainInfo indicates an expected call of GetChainInfo
func (mr *MockApiServiceServerMockRecorder) GetChainInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainInfo", reflect.TypeOf((*MockApiServiceServer)(nil).GetChainInfo), arg0, arg1)
}

// GetContract mocks base method
func (m *MockApiServiceServer) GetContract(arg0 context.Context, arg1 *pb.GetContractRequest) (*pb.Contract, error) {
	ret := m.ctrl.Call(m, "GetContract", arg0, arg1)
	ret0, _ := ret[0].(*pb.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract
func (mr *MockApiServiceServerMockRecorder) GetContract(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockApiServiceServer)(nil).GetContract), arg0, arg1)
}

// GetContractStorage mocks base method
func (m *MockApiServiceServer) GetContractStorage(arg0 context.Context, arg1 *pb.GetContractStorageRequest) (*pb.GetContractStorageResponse, error) {
	ret := m.ctrl.Call(m, "GetContractStorage", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetContractStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractStorage indicates an expected call of GetContractStorage
func (mr *MockApiServiceServerMockRecorder) GetContractStorage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractStorage", reflect.TypeOf((*MockApiServiceServer)(nil).GetContractStorage), arg0, arg1)
}

// GetContractStorageFields mocks base method
func (m *MockApiServiceServer) GetContractStorageFields(arg0 context.Context, arg1 *pb.GetContractStorageFieldsRequest) (*pb.GetContractStorageFieldsResponse, error) {
	ret := m.ctrl.Call(m, "GetContractStorageFields", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetContractStorageFieldsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractStorageFields indicates an expected call of GetContractStorageFields
func (mr *MockApiServiceServerMockRecorder) GetContractStorageFields(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractStorageFields", reflect.TypeOf((*MockApiServiceServer)(nil).GetContractStorageFields), arg0, arg1)
}

// GetContractVote mocks base method
func (m *MockApiServiceServer) GetContractVote(arg0 context.Context, arg1 *pb.GetContractRequest) (*pb.ContractVote, error) {
	ret := m.ctrl.Call(m, "GetContractVote", arg0, arg1)
	ret0, _ := ret[0].(*pb.ContractVote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractVote indicates an expected call of GetContractVote
func (mr *MockApiServiceServerMockRecorder) GetContractVote(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractVote", reflect.TypeOf((*MockApiServiceServer)(nil).GetContractVote), arg0, arg1)
}

// GetGasRatio mocks base method
func (m *MockApiServiceServer) GetGasRatio(arg0 context.Context, arg1 *pb.EmptyRequest) (*pb.GasRatioResponse, error) {
	ret := m.ctrl.Call(m, "GetGasRatio", arg0, arg1)
	ret0, _ := ret[0].(*pb.GasRatioResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGasRatio indicates an expected call of GetGasRatio
func (mr *MockApiServiceServerMockRecorder) GetGasRatio(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGasRatio", reflect.TypeOf((*MockApiServiceServer)(nil).GetGasRatio), arg0, arg1)
}

// GetNodeInfo mocks base method
func (m *MockApiServiceServer) GetNodeInfo(arg0 context.Context, arg1 *pb.EmptyRequest) (*pb.NodeInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetNodeInfo", arg0, arg1)
	ret0, _ := ret[0].(*pb.NodeInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeInfo indicates an expected call of GetNodeInfo
func (mr *MockApiServiceServerMockRecorder) GetNodeInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeInfo", reflect.TypeOf((*MockApiServiceServer)(nil).GetNodeInfo), arg0, arg1)
}

// GetProducerVoteInfo mocks base method
func (m *MockApiServiceServer) GetProducerVoteInfo(arg0 context.Context, arg1 *pb.GetProducerVoteInfoRequest) (*pb.GetProducerVoteInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetProducerVoteInfo", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetProducerVoteInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducerVoteInfo indicates an expected call of GetProducerVoteInfo
func (mr *MockApiServiceServerMockRecorder) GetProducerVoteInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducerVoteInfo", reflect.TypeOf((*MockApiServiceServer)(nil).GetProducerVoteInfo), arg0, arg1)
}

// GetRAMInfo mocks base method
func (m *MockApiServiceServer) GetRAMInfo(arg0 context.Context, arg1 *pb.EmptyRequest) (*pb.RAMInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetRAMInfo", arg0, arg1)
	ret0, _ := ret[0].(*pb.RAMInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRAMInfo indicates an expected call of GetRAMInfo
func (mr *MockApiServiceServerMockRecorder) GetRAMInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRAMInfo", reflect.TypeOf((*MockApiServiceServer)(nil).GetRAMInfo), arg0, arg1)
}

// GetToken721Balance mocks base method
func (m *MockApiServiceServer) GetToken721Balance(arg0 context.Context, arg1 *pb.GetTokenBalanceRequest) (*pb.GetToken721BalanceResponse, error) {
	ret := m.ctrl.Call(m, "GetToken721Balance", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetToken721BalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken721Balance indicates an expected call of GetToken721Balance
func (mr *MockApiServiceServerMockRecorder) GetToken721Balance(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken721Balance", reflect.TypeOf((*MockApiServiceServer)(nil).GetToken721Balance), arg0, arg1)
}

// GetToken721Metadata mocks base method
func (m *MockApiServiceServer) GetToken721Metadata(arg0 context.Context, arg1 *pb.GetToken721InfoRequest) (*pb.GetToken721MetadataResponse, error) {
	ret := m.ctrl.Call(m, "GetToken721Metadata", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetToken721MetadataResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken721Metadata indicates an expected call of GetToken721Metadata
func (mr *MockApiServiceServerMockRecorder) GetToken721Metadata(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken721Metadata", reflect.TypeOf((*MockApiServiceServer)(nil).GetToken721Metadata), arg0, arg1)
}

// GetToken721Owner mocks base method
func (m *MockApiServiceServer) GetToken721Owner(arg0 context.Context, arg1 *pb.GetToken721InfoRequest) (*pb.GetToken721OwnerResponse, error) {
	ret := m.ctrl.Call(m, "GetToken721Owner", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetToken721OwnerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken721Owner indicates an expected call of GetToken721Owner
func (mr *MockApiServiceServerMockRecorder) GetToken721Owner(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken721Owner", reflect.TypeOf((*MockApiServiceServer)(nil).GetToken721Owner), arg0, arg1)
}

// GetTokenBalance mocks base method
func (m *MockApiServiceServer) GetTokenBalance(arg0 context.Context, arg1 *pb.GetTokenBalanceRequest) (*pb.GetTokenBalanceResponse, error) {
	ret := m.ctrl.Call(m, "GetTokenBalance", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetTokenBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenBalance indicates an expected call of GetTokenBalance
func (mr *MockApiServiceServerMockRecorder) GetTokenBalance(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenBalance", reflect.TypeOf((*MockApiServiceServer)(nil).GetTokenBalance), arg0, arg1)
}

// GetTokenInfo mocks base method
func (m *MockApiServiceServer) GetTokenInfo(arg0 context.Context, arg1 *pb.GetTokenInfoRequest) (*pb.TokenInfo, error) {
	ret := m.ctrl.Call(m, "GetTokenInfo", arg0, arg1)
	ret0, _ := ret[0].(*pb.TokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokenInfo indicates an expected call of GetTokenInfo
func (mr *MockApiServiceServerMockRecorder) GetTokenInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenInfo", reflect.TypeOf((*MockApiServiceServer)(nil).GetTokenInfo), arg0, arg1)
}

// GetTxByHash mocks base method
func (m *MockApiServiceServer) GetTxByHash(arg0 context.Context, arg1 *pb.TxHashRequest) (*pb.TransactionResponse, error) {
	ret := m.ctrl.Call(m, "GetTxByHash", arg0, arg1)
	ret0, _ := ret[0].(*pb.TransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxByHash indicates an expected call of GetTxByHash
func (mr *MockApiServiceServerMockRecorder) GetTxByHash(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxByHash", reflect.TypeOf((*MockApiServiceServer)(nil).GetTxByHash), arg0, arg1)
}

// GetTxReceiptByTxHash mocks base method
func (m *MockApiServiceServer) GetTxReceiptByTxHash(arg0 context.Context, arg1 *pb.TxHashRequest) (*pb.TxReceipt, error) {
	ret := m.ctrl.Call(m, "GetTxReceiptByTxHash", arg0, arg1)
	ret0, _ := ret[0].(*pb.TxReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxReceiptByTxHash indicates an expected call of GetTxReceiptByTxHash
func (mr *MockApiServiceServerMockRecorder) GetTxReceiptByTxHash(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxReceiptByTxHash", reflect.TypeOf((*MockApiServiceServer)(nil).GetTxReceiptByTxHash), arg0, arg1)
}

// GetVoterBonus mocks base method
func (m *MockApiServiceServer) GetVoterBonus(arg0 context.Context, arg1 *pb.GetAccountRequest) (*pb.VoterBonus, error) {
	ret := m.ctrl.Call(m, "GetVoterBonus", arg0, arg1)
	ret0, _ := ret[0].(*pb.VoterBonus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVoterBonus indicates an expected call of GetVoterBonus
func (mr *MockApiServiceServerMockRecorder) GetVoterBonus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVoterBonus", reflect.TypeOf((*MockApiServiceServer)(nil).GetVoterBonus), arg0, arg1)
}

// ListContractStorage mocks base method
func (m *MockApiServiceServer) ListContractStorage(arg0 context.Context, arg1 *pb.ListContractStorageRequest) (*pb.ListContractStorageResponse, error) {
	ret := m.ctrl.Call(m, "ListContractStorage", arg0, arg1)
	ret0, _ := ret[0].(*pb.ListContractStorageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContractStorage indicates an expected call of ListContractStorage
func (mr *MockApiServiceServerMockRecorder) ListContractStorage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContractStorage", reflect.TypeOf((*MockApiServiceServer)(nil).ListContractStorage), arg0, arg1)
}

// SendTransaction mocks base method
func (m *MockApiServiceServer) SendTransaction(arg0 context.Context, arg1 *pb.TransactionRequest) (*pb.SendTransactionResponse, error) {
	ret := m.ctrl.Call(m, "SendTransaction", arg0, arg1)
	ret0, _ := ret[0].(*pb.SendTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction
func (mr *MockApiServiceServerMockRecorder) SendTransaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockApiServiceServer)(nil).SendTransaction), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockApiServiceServer) Subscribe(arg0 *pb.SubscribeRequest, arg1 pb.ApiService_SubscribeServer) error {
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockApiServiceServerMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockApiServiceServer)(nil).Subscribe), arg0, arg1)
}
