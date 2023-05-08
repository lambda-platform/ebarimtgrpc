package server

import (
	"context"
	"github.com/lambda-platform/ebarimt/posapi"
	"github.com/lambda-platform/ebarimtgrpc/ebarimt"
	ebarimtProto "github.com/lambda-platform/ebarimtgrpc/proto"
)

func (server *Server) GetInformation(ctx context.Context, in *ebarimtProto.GetInformationRequest) (*ebarimtProto.InformationOutput, error) {

	api := ebarimt.PosAPI(in.PosID)
	info, err := api.GetInformation()
	if err != nil {
		return nil, err
	}
	ExtraInfo := ebarimtProto.ExtraInfo{
		CountBill:    info.ExtraInfo.CountBill,
		CountLottery: int32(info.ExtraInfo.CountLottery),
		LastSentDate: info.ExtraInfo.LastSentDate,
		PosVersion:   info.ExtraInfo.PosVersion,
	}
	InformationOutput := ebarimtProto.InformationOutput{
		RegisterNo: info.RegisterNo,
		BranchNo:   info.BranchNo,
		PosId:      info.PosID,
		DbDirPath:  info.DBDirPath,
		ExtraInfo:  &ExtraInfo,
	}
	return &InformationOutput, nil
}

func (server *Server) CheckApi(ctx context.Context, in *ebarimtProto.CheckApiRequest) (*ebarimtProto.CheckOutput, error) {
	api := ebarimt.PosAPI(in.PosID)
	res, err := api.CheckApi()
	if err != nil {
		return nil, err
	}

	return &ebarimtProto.CheckOutput{
		Success: res.Success,
		Database: &ebarimtProto.DatabaseStatus{
			Success: res.Database.Success,
			Message: res.Database.Message,
		},
		Config: &ebarimtProto.ConfigStatus{
			Success: res.Config.Success,
			Message: res.Config.Message,
		},
		Network: &ebarimtProto.NetworkStatus{
			Success: res.Network.Success,
			Message: res.Network.Message,
		},
	}, nil
}

func (server *Server) CallFunction(ctx context.Context, in *ebarimtProto.CallFunctionRequest) (*ebarimtProto.CallFunctionResponse, error) {

	api := ebarimt.PosAPI(in.PosID)
	res, err := api.CallFunction(in.FunctionName, in.Params)
	if err != nil {
		return nil, err
	}

	return &ebarimtProto.CallFunctionResponse{
		Result: res,
	}, nil
}

func (server *Server) Put(ctx context.Context, in *ebarimtProto.PutInput) (*ebarimtProto.PutOutput, error) {

	api := ebarimt.PosAPI(in.PosID)
	output, err := api.Put(ConvertFromGRPCPutInput(in))
	if err != nil {
		return nil, err
	}
	return &ebarimtProto.PutOutput{
		Success:           output.Success,
		RegisterNo:        output.RegisterNo,
		BillId:            output.BillID,
		Date:              output.Date,
		MacAddress:        output.MacAddress,
		InternalCode:      output.InternalCode,
		BillType:          output.BillType,
		QrData:            output.QRData,
		Lottery:           output.Lottery,
		LotteryWarningMsg: output.LotteryWarningMsg,
		ErrorCode:         int32(output.ErrorCode),
		Message:           output.Message,
	}, nil
}

func ConvertFromGRPCPutInput(input *ebarimtProto.PutInput) posapi.PutInput {
	stocks := make([]posapi.Stock, len(input.Stocks))
	for i, stock := range input.Stocks {
		stocks[i] = posapi.Stock{
			Code:        stock.Code,
			Name:        stock.Name,
			MeasureUnit: stock.MeasureUnit,
			Qty:         stock.Qty,
			UnitPrice:   stock.UnitPrice,
			TotalAmount: stock.TotalAmount,
			CityTax:     stock.CityTax,
			Vat:         stock.Vat,
			BarCode:     stock.BarCode,
		}
	}

	bankTransactions := make([]posapi.BankTransaction, len(input.BankTransactions))
	for i, transaction := range input.BankTransactions {
		bankTransactions[i] = posapi.BankTransaction{
			Rrn:          transaction.Rrn,
			BankID:       transaction.BankId,
			TerminalID:   transaction.TerminalId,
			ApprovalCode: transaction.ApprovalCode,
			Amount:       transaction.Amount,
		}
	}

	return posapi.PutInput{
		Amount:           input.Amount,
		Vat:              input.Vat,
		CashAmount:       input.CashAmount,
		NonCashAmount:    input.NonCashAmount,
		CityTax:          input.CityTax,
		DistrictCode:     input.DistrictCode,
		PosNo:            input.PosNo,
		CustomerNo:       input.CustomerNo,
		BillType:         input.BillType,
		BillIDSuffix:     input.BillIdSuffix,
		ReturnBillID:     input.ReturnBillId,
		TaxType:          input.TaxType,
		InvoiceID:        input.InvoiceId,
		ReportMonth:      input.ReportMonth,
		BranchNo:         input.BranchNo,
		Stocks:           stocks,
		BankTransactions: bankTransactions,
	}
}

func (server *Server) PutBatch(ctx context.Context, in *ebarimtProto.PutInputBatch) (*ebarimtProto.PutOutput, error) {
	api := ebarimt.PosAPI(in.PosID)
	output, err := api.PutBatch(ConvertFromGRPCPutInputBatch(in))
	if err != nil {
		return nil, err
	}
	return &ebarimtProto.PutOutput{
		Success:           output.Success,
		RegisterNo:        output.RegisterNo,
		BillId:            output.BillID,
		Date:              output.Date,
		MacAddress:        output.MacAddress,
		InternalCode:      output.InternalCode,
		BillType:          output.BillType,
		QrData:            output.QRData,
		Lottery:           output.Lottery,
		LotteryWarningMsg: output.LotteryWarningMsg,
		ErrorCode:         int32(output.ErrorCode),
		Message:           output.Message,
	}, nil
}

func ConvertFromGRPCPutInputBatch(grpcModel *ebarimtProto.PutInputBatch) posapi.PutInputBatch {
	bills := make([]posapi.Bill, len(grpcModel.GetBills()))

	for i, grpcBill := range grpcModel.Bills {

		stocks := make([]posapi.Stock, len(grpcBill.Stocks))
		for i, stock := range grpcBill.Stocks {
			stocks[i] = posapi.Stock{
				Code:        stock.Code,
				Name:        stock.Name,
				MeasureUnit: stock.MeasureUnit,
				Qty:         stock.Qty,
				UnitPrice:   stock.UnitPrice,
				TotalAmount: stock.TotalAmount,
				CityTax:     stock.CityTax,
				Vat:         stock.Vat,
				BarCode:     stock.BarCode,
			}
		}

		bankTransactions := make([]posapi.BankTransaction, len(grpcBill.BankTransactions))
		for i, transaction := range grpcBill.BankTransactions {
			bankTransactions[i] = posapi.BankTransaction{
				Rrn:          transaction.Rrn,
				BankID:       transaction.BankId,
				TerminalID:   transaction.TerminalId,
				ApprovalCode: transaction.ApprovalCode,
				Amount:       transaction.Amount,
			}
		}
		bills[i] = posapi.Bill{
			Amount:           grpcBill.Amount,
			Vat:              grpcBill.Vat,
			CashAmount:       grpcBill.CashAmount,
			NonCashAmount:    grpcBill.NonCashAmount,
			CityTax:          grpcBill.CityTax,
			DistrictCode:     grpcBill.DistrictCode,
			PosNo:            grpcBill.PosNo,
			CustomerNo:       grpcBill.CustomerNo,
			BillType:         grpcBill.BillType,
			BillIDSuffix:     grpcBill.BillIdSuffix,
			ReturnBillID:     grpcBill.ReturnBillId,
			TaxType:          grpcBill.TaxType,
			InvoiceID:        grpcBill.InvoiceId,
			ReportMonth:      grpcBill.ReportMonth,
			BranchNo:         grpcBill.BranchNo,
			Stocks:           stocks,
			BankTransactions: bankTransactions,
			InternalId:       grpcBill.GetInternalId(),
			RegisterNo:       grpcBill.GetRegisterNo(),
		}
	}

	return posapi.PutInputBatch{
		Group:        grpcModel.GetGroup(),
		Vat:          grpcModel.GetVat(),
		Amount:       grpcModel.GetAmount(),
		BillType:     grpcModel.GetBillType(),
		BillIDSuffix: grpcModel.GetBillIdSuffix(),
		PosNo:        grpcModel.GetPosNo(),
		Bills:        bills,
	}
}
func (server *Server) ReturnBill(ctx context.Context, in *ebarimtProto.BillInput) (*ebarimtProto.BillOutput, error) {

	api := ebarimt.PosAPI(in.PosID)
	res, err := api.ReturnBill(ConvertFromGRPCBillInput(in))
	if err != nil {
		return nil, err
	}

	return &ebarimtProto.BillOutput{
		Success:   res.Success,
		ErrorCode: int32(res.ErrorCode),
		Message:   res.Message,
	}, nil
}
func ConvertFromGRPCBillInput(grpcModel *ebarimtProto.BillInput) posapi.BillInput {
	return posapi.BillInput{
		ReturnBillID: grpcModel.GetReturnBillId(),
		Date:         grpcModel.GetDate(),
	}
}

func (server *Server) SendData(ctx context.Context, in *ebarimtProto.SendDataRequest) (*ebarimtProto.DataOutput, error) {

	api := ebarimt.PosAPI(in.PosID)
	res, err := api.SendData()
	if err != nil {
		return nil, err
	}

	return &ebarimtProto.DataOutput{
		Success:   res.Success,
		ErrorCode: int32(res.ErrorCode),
		Message:   res.Message,
	}, nil
}
