package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kashifkhan0771/HostService/config"
	"github.com/kashifkhan0771/HostService/db"
	domainErr "github.com/kashifkhan0771/HostService/errors"
	"github.com/kashifkhan0771/HostService/models"
)

const (
	hostCollection = "hosts"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mysql database connection.
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

func (c *client) AddHost(host *models.Host) (string, error) {
	if host.ID != "" {
		return "", errors.New("id is not empty")
	}

	host.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hostCollection)
	if _, err := collection.InsertOne(context.TODO(), host); err != nil {
		return "", errors.Wrap(err, "failed to add host")
	}

	return host.ID, nil
}

func (c *client) GetHost(id string) (*models.Host, error) {
	var stu *models.Host
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hostCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&stu); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("host: %s not found", id))
		}
		return nil, err
	}
	return stu, nil
}

func (c *client) DeleteHost(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hostCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete host")
	}

	return nil
}

func (c *client) UpdateHost(host *models.Host) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(hostCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": host.ID}, bson.M{"$set": host}); err != nil {
		return errors.Wrap(err, "failed to update host")
	}

	return nil
}
