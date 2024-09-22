package db

import (
	"context"
	"time"

	"github.com/rickalon/GoWebScraper/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	GetUrl(context.Context, string) (error, *data.URL)
	InsertOne(context.Context, ...any) error
	CreateDB()
	IsON() bool
}
type MongoDB struct {
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
	IsOn            bool
}

func NewMongoDB(ctx context.Context, prefix string) (*MongoDB, error) {
	client, err := runConnection(ctx, prefix)
	return &MongoDB{MongoClient: client}, err
}

func runConnection(ctxOpt context.Context, prefix string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctxOpt, 20*time.Second)
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(prefix))
}

func (m *MongoDB) GetUrl(ctx context.Context, url string) (error, *data.URL) {
	filter := bson.M{"urlname": url}
	var result *data.URL
	err := m.MongoCollection.FindOne(ctx, filter).Decode(&result)

	return err, result
}

func (m *MongoDB) CreateDB() {
	m.MongoCollection = m.MongoClient.Database("URI").Collection("URLS")
}

func (m *MongoDB) InsertOne(ctx context.Context, v ...any) error {
	_, err := m.MongoCollection.InsertMany(ctx, v)
	return err
}

func (m *MongoDB) IsON() bool {
	return m.IsOn
}
