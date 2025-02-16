package service

import (
	"context"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"github.com/hltl/GoMall/app/payment/biz/dal/mysql"
	"github.com/hltl/GoMall/app/payment/biz/model"
	payment "github.com/hltl/GoMall/rpc_gen/kitex_gen/payment"
	"gorm.io/gorm"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	if err := card.Validate(true); err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4004001, err.Error())
	}

	tId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(4005001, err.Error())
	}
	err = mysql.DB.WithContext(s.ctx).Transaction(func(tx *gorm.DB) error {
		pl := &model.PaymentLog{
			UserID:        uint(req.UserId),
			OrderId:       req.OrderId,
			TransactionID: tId.String(),
			Amount:        req.Amount,
			PayAt:         time.Now(),
		}
		if err := model.Pay(s.ctx, tx, pl); err != nil {
			return kerrors.NewGRPCBizStatusError(4005002, err.Error())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &payment.ChargeResp{TransactionId: tId.String()}, nil
}
