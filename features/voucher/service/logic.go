package service

import "my-tourist-ticket/features/voucher"

type voucherService struct {
	voucherData voucher.VoucherDataInterface
}

func New(repo voucher.VoucherDataInterface) voucher.VoucherServiceInterface {
	return &voucherService{
		voucherData: repo,
	}
}

// Create implements voucher.VoucherServiceInterface.
func (service *voucherService) Create(input voucher.Core) error {
	err := service.voucherData.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

// Update implements voucher.VoucherServiceInterface.
func (service *voucherService) Update(voucherId int, input voucher.Core) error {
	err := service.voucherData.Update(voucherId, input)
	if err != nil {
		return err
	}

	return nil
}
