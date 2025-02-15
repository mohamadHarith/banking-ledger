package repository

import (
	"context"
	"log"

	"github.com/mohamadHarith/banking-ledger/services/transaction-logger-service/internal/config"
	"github.com/mohamadHarith/banking-ledger/shared/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	mgo *mongo.Client
	db  *mongo.Database
}

func New() *Repository {
	conf := config.GetConf()

	log.Println(conf.MongoDB.User)
	log.Println(conf.MongoDB.Password)

	// uri := fmt.Sprintf("mongodb://%v:%v@localhost:27017/?authSource=banking&authMechanism=SCRAM-SHA-256", conf.MongoDB.User, conf.MongoDB.Password)

	uri := "mongodb://user:password@localhost:27017/?retryWrites=true&loadBalanced=false&serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&authSource=banking&authMechanism=SCRAM-SHA-256"

	mgo, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = mgo.Ping(context.Background(), nil)
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

func (r *Repository) GetTransactionLogs(
	ctx context.Context,
	page uint32,
	userId, accountId string,
) (
	txnLogs []entity.TransactionLog,
	totalNoOfRecords, totalNoOfPages uint32,
	hasNextPage, hasPrevPage bool,
	err error,
) {
	limit := uint32(30)

	skip := (page - 1) * limit

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}}) // Optional: sort by ID

	cursor, err := r.db.Collection("transaction_log").Find(ctx, bson.M{
		"userId":    userId,
		"accountId": accountId,
	})
	if err != nil {
		return
	}

	err = cursor.All(ctx, &txnLogs)
	if err != nil {
		return
	}

	count, err := r.db.Collection("transaction_log").CountDocuments(ctx, bson.M{
		"userId":    userId,
		"accountId": accountId,
	})
	if err != nil {
		return
	}

	totalNoOfRecords = uint32(count)
	totalNoOfPages = totalNoOfRecords / limit
	if totalNoOfRecords%limit != 0 {
		totalNoOfPages++
	}

	if totalNoOfPages > page {
		hasNextPage = true
	}

	if page > 1 && totalNoOfPages >= page {
		hasPrevPage = true
	}

	return
}
