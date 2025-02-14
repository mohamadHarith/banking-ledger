package repository

import (
	"context"
	"fmt"

	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/config"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	mgo *mongo.Client
	db  *mongo.Database
}

func New() *Repository {
	conf := config.GetConf()

	uri := fmt.Sprintf("mongodb://%v:%v@localhost:27017", conf.MongoDB.User, conf.MongoDB.Password)

	mgo, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	db := mgo.Database("banking")

	return &Repository{
		mgo: mgo,
		db:  db,
	}
}

func (r *Repository) InsertTransactionLog(ctx context.Context, l entity.TransactionLog) error {
	_, err := r.db.Collection("transaction_log").InsertOne(ctx, l)

	return err
}
