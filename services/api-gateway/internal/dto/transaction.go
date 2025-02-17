package dto

import (
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	pb "github.com/mohamadHarith/banking-ledger/shared/proto/transaction_logger_proto"
)

func FromProtoToEntityTransactionLogs(in []*pb.TransactionLog) (o []entity.TransactionLog) {
	for i := 0; i < len(in); i++ {
		log := in[i]
		o = append(o, entity.TransactionLog{
			Id:          *log.Id,
			AccountId:   *log.AccountId,
			UserId:      *log.UserId,
			Amount:      *log.Amount,
			Balance:     *log.Balance,
			Description: *log.Description,
			CreatedAt:   log.CreatedAt.AsTime(),
		})
	}

	return
}
