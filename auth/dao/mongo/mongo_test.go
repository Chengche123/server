package mongo

import (
	"comic/share/database/mongo/mgotesting"
	util "comic/share/database/mongo/util"
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

var mongoURL string

func TestResolveOpenID(t *testing.T) {
	ctx := context.Background()

	logger, err := zap.NewDevelopment()
	if err != nil {
		t.Fatalf("failed to create zap logger: %v", err)
	}

	mo, err := NewMongoCol(context.Background(),
		logger,
		mongoURL,
		"comic")
	if err != nil {
		t.Fatalf("failed to create mongo connect: %v", err)
	}

	_, err = mo.col.InsertMany(ctx, []interface{}{
		bson.M{
			filedOpenID:  "openid_1",
			util.FiledID: mustObjectID("6055a96e48509813082fb913"),
		},
		bson.M{
			filedOpenID:  "openid_2",
			util.FiledID: mustObjectID("6055a96e48509813082fb914"),
		},
	})
	if err != nil {
		t.Fatalf("failed to insertMany: %v", err)
	}

	mo.objIDCreater = func() primitive.ObjectID {
		id, _ := primitive.ObjectIDFromHex("6055a96e48509813082fb915")
		return id
	}

	_, err = mo.ResolveOpenID(ctx, "openid_3")
	if err != nil {
		t.Fatalf("failed to Resolve openID: %v", err)
	}

	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			name:   "existing_user",
			openID: "openid_1",
			want:   "6055a96e48509813082fb913",
		},
		{
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "6055a96e48509813082fb914",
		},
		{
			name:   "new_user",
			openID: "openid_3",
			want:   "6055a96e48509813082fb915",
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			id, err := mo.ResolveOpenID(context.Background(), cs.openID)
			if err != nil {
				t.Errorf("failed to resolve openID %q: %v", cs.openID, err)
			}

			if id != cs.want {
				t.Errorf("resolve openID: %q; want: %q; got: %q; ", cs.openID, cs.want, id)
			}
		})
	}
}

func mustObjectID(hex string) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		panic(err)
	}
	return id
}

func TestMain(m *testing.M) {
	os.Exit(mgotesting.RunWithMongoInDocker(m, &mongoURL))
}
