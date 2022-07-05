//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package service

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	rpc_service "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_service"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &bankService_{}
		},
	})
	bankServiceStructDescriptor := &autowire.StructDescriptor{
		Alias: "github.com/alibaba/ioc-golang/example/transaction_rpc/server/pkg/service/api.BankServiceIOCRPCClient",
		Factory: func() interface{} {
			return &BankService{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*BankService)
			var constructFunc BankServiceConstructFunc = InitBankService
			return constructFunc(impl)
		},
		TransactionMethodsMap: map[string]string{
			"AddMoney":    "AddMoneyRollout",
			"RemoveMoney": "RemoveMoneyRollout",
		},
	}
	rpc_service.RegisterStructDescriptor(bankServiceStructDescriptor)
}

type BankServiceConstructFunc func(impl *BankService) (*BankService, error)
type bankService_ struct {
	GetMoney_           func(id int) int
	AddMoney_           func(id, num int) error
	AddMoneyRollout_    func(id, num int, errMsg string) error
	RemoveMoney_        func(id, num int) error
	RemoveMoneyRollout_ func(id, num int, errMsg string) error
}

func (b *bankService_) GetMoney(id int) int {
	return b.GetMoney_(id)
}

func (b *bankService_) AddMoney(id, num int) error {
	return b.AddMoney_(id, num)
}

func (b *bankService_) AddMoneyRollout(id, num int, errMsg string) error {
	return b.AddMoneyRollout_(id, num, errMsg)
}

func (b *bankService_) RemoveMoney(id, num int) error {
	return b.RemoveMoney_(id, num)
}

func (b *bankService_) RemoveMoneyRollout(id, num int, errMsg string) error {
	return b.RemoveMoneyRollout_(id, num, errMsg)
}

type BankServiceIOCInterface interface {
	GetMoney(id int) int
	AddMoney(id, num int) error
	AddMoneyRollout(id, num int, errMsg string) error
	RemoveMoney(id, num int) error
	RemoveMoneyRollout(id, num int, errMsg string) error
}

func GetBankServiceRpc() (*BankService, error) {
	i, err := rpc_service.GetImpl(util.GetSDIDByStructPtr(new(BankService)))
	if err != nil {
		return nil, err
	}
	impl := i.(*BankService)
	return impl, nil
}

func GetBankServiceIOCInterfaceRpc() (BankServiceIOCInterface, error) {
	i, err := rpc_service.GetImplWithProxy(util.GetSDIDByStructPtr(new(BankService)))
	if err != nil {
		return nil, err
	}
	impl := i.(BankServiceIOCInterface)
	return impl, nil
}
