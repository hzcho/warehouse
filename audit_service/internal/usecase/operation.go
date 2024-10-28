package usecase

import (
	"audit/internal/domain/model"
	"audit/internal/domain/net/request"
	"audit/internal/domain/repository"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation struct {
	operationRepo repository.Operation
	log           *logrus.Logger
}

func NewOperation(operationRepo repository.Operation, log *logrus.Logger) *Operation {
	return &Operation{
		operationRepo: operationRepo,
		log:           log,
	}
}

func (u *Operation) GetById(ctx context.Context, id primitive.ObjectID) (model.OperationLog, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/operation/GetById",
	})
	log.Info()

	operLog, err := u.operationRepo.GetById(ctx, id)
	if err != nil {
		log.Error(err)
		return model.OperationLog{}, err
	}

	return operLog, nil
}

func (u *Operation) GetAll(ctx context.Context, filter request.GetAllFilter) ([]model.OperationLog, error) {
	log := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/operation/GetAll",
	})
	log.Info(fmt.Sprintf("%v %v %v %v %v", filter.UserId, filter.ProductId, filter.OperationType, filter.Limit, filter.Page))

	operLogs, err := u.operationRepo.GetAll(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return operLogs, nil
}

func (u *Operation) Save(ctx context.Context, log model.OperationLog) (primitive.ObjectID, error) {
	lg := u.log.WithFields(logrus.Fields{
		"op": "internal/usecase/operation/Save",
	})
	lg.Info()

	log.OprationId = primitive.NewObjectID()

	id, err := u.operationRepo.Save(ctx, log)
	if err != nil {
		lg.Error(err)
		return primitive.NilObjectID, nil
	}

	return id, nil
}
