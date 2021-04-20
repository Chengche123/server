package mongo

import (
	"context"
	"fmt"

	util "comic/share/mongo/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const filedOpenID = "open_id"
const colNameAccount = "account"

type mongoCol struct {
	logger       *zap.Logger
	col          *mongo.Collection
	objIDCreater func() primitive.ObjectID
}

func (m *mongoCol) Close() error {
	return nil
}

func (m *mongoCol) ResolveOpenID(ctx context.Context, openID string) (string, error) {
	objID := m.objIDCreater()

	res := m.col.FindOneAndUpdate(ctx, bson.M{
		filedOpenID: openID,
	}, util.SetOnInsert(bson.M{
		util.FiledID: objID,
		filedOpenID:  openID,
	}), options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))

	err := res.Err()
	if err != nil {
		return "", fmt.Errorf("failed to FindOneAndUpdate: %v", err)
	}

	var row util.ObjID

	err = res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("failed to decode result: %v", err)
	}

	return row.ID.Hex(), nil
}

// NewMongoDB create a dao
func NewMongoCol(ctx context.Context, logger *zap.Logger, cstr string, dbName string) (*mongoCol, error) {

	mc, err := mongo.Connect(ctx, options.Client().ApplyURI(cstr))
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongodb: %v", err)
	}
	col := mc.Database(dbName).Collection(colNameAccount)

	return &mongoCol{
		logger:       logger,
		col:          col,
		objIDCreater: primitive.NewObjectID,
	}, nil
}
