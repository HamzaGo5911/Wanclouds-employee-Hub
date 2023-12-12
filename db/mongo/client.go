package mongo

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/HamzaGo5911/wanclouds-employee-hub/config"
	"github.com/HamzaGo5911/wanclouds-employee-hub/db"
	domainerr "github.com/HamzaGo5911/wanclouds-employee-hub/errors"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

const (
	employeeCollection = "employee"
	homeCollection     = "home"
	officeCollection   = "office"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

var (
	once sync.Once
	cli  *mongo.Client
)

// NewClient initializes a mongo database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	once.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
		log().Infof("initializing mongodb: %s", uri)
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			log().Errorf("failed to connect to MongoDB: %v", err)
			return
		}
		cli = client
	})

	if cli == nil {
		return nil, errors.New("failed to initialize MongoDB client")
	}

	return &client{conn: cli}, nil
}

// AddEmployee adds a new employee to the database.
func (c *client) AddEmployee(employee *models.Employee) (string, error) {
	if employee.ID == "" {
		employee.ID = uuid.NewV4().String()
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(employeeCollection)
	filter := bson.M{"_id": employee.ID}
	update := bson.M{"$set": employee}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return "", errors.Wrap(err, "failed to update employee")
	}

	return employee.ID, nil
}

// GetEmployeeByID retrieves a employee from the database based on its unique ID.
func (c *client) GetEmployeeByID(id string) (*models.Employee, error) {
	var employee *models.Employee
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(employeeCollection)

	filter := bson.M{"_id": id}
	if err := collection.FindOne(context.TODO(), filter).Decode(&employee); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainerr.NewAPIError(domainerr.NotFound, fmt.Sprintf("employee with ID '%s' not found", id))
		}
		return nil, err
	}

	return employee, nil
}

// DeleteEmployee deletes a employee from the database based on its unique ID.
func (c client) DeleteEmployee(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(employeeCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete employee")
	}

	return nil
}

// ListEmployee retrieves a list of all employee from the database.
func (c client) ListEmployee() ([]*models.Employee, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(employeeCollection)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list employee")
	}
	defer cursor.Close(context.TODO())

	var employee []*models.Employee
	for cursor.Next(context.TODO()) {
		var e models.Employee
		if err := cursor.Decode(&e); err != nil {
			return nil, errors.Wrap(err, "failed to decode employee")
		}
		employee = append(employee, &e)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}

	return employee, nil
}

// AddOffice adds a new office to the database.
func (c *client) AddOffice(office *models.Office) (string, error) {
	if office.ID == "" {
		office.ID = uuid.NewV4().String()
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(officeCollection)
	filter := bson.M{"_id": office.ID}
	update := bson.M{"$set": office}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return "", errors.Wrap(err, "failed to update office")
	}

	return office.ID, nil
}

// GetOfficeByID retrieves a office from the database based on its unique ID.
func (c *client) GetOfficeByID(id string) (*models.Office, error) {
	var office *models.Office
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(officeCollection)

	filter := bson.M{"_id": id}
	if err := collection.FindOne(context.TODO(), filter).Decode(&office); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainerr.NewAPIError(domainerr.NotFound, fmt.Sprintf("office with ID '%s' not found", id))
		}
		return nil, err
	}

	return office, nil
}

// DeleteOffice deletes a office from the database based on its unique ID.
func (c client) DeleteOffice(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(officeCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete office")
	}

	return nil
}

// ListOffice retrieves a list of all Office from the database.
func (c client) ListOffice() ([]*models.Office, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(officeCollection)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list office")
	}

	defer cursor.Close(context.TODO())

	var office []*models.Office
	for cursor.Next(context.TODO()) {
		var o models.Office
		if err := cursor.Decode(&o); err != nil {
			return nil, errors.Wrap(err, "failed to decode office")
		}
		office = append(office, &o)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}

	return office, nil
}

// AddHome adds a new Home to the database.
func (c *client) AddHome(home *models.Home) (string, error) {
	if home.ID == "" {
		home.ID = uuid.NewV4().String()
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(homeCollection)
	filter := bson.M{"_id": home.ID}
	update := bson.M{"$set": home}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(context.TODO(), filter, update, opts); err != nil {
		return "", errors.Wrap(err, "failed to update home")
	}

	return home.ID, nil
}

// GetHomeByID retrieves a home from the database based on its unique ID.
func (c *client) GetHomeByID(id string) (*models.Home, error) {
	var home *models.Home
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(homeCollection)

	filter := bson.M{"_id": id}
	if err := collection.FindOne(context.TODO(), filter).Decode(&home); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainerr.NewAPIError(domainerr.NotFound, fmt.Sprintf("home with ID '%s' not found", id))
		}
		return nil, err
	}

	return home, nil
}

// DeleteHome deletes a home from the database based on its unique ID.
func (c client) DeleteHome(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(homeCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete home")
	}

	return nil
}

// ListHome retrieves a list of all Home from the database.
func (c client) ListHome() ([]*models.Home, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(homeCollection)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to list home")
	}

	defer cursor.Close(context.TODO())

	var home []*models.Home
	for cursor.Next(context.TODO()) {
		var h models.Home
		if err := cursor.Decode(&h); err != nil {
			return nil, errors.Wrap(err, "failed to decode home")
		}
		home = append(home, &h)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}

	return home, nil
}

// Disconnect - closes the db connections
func (c *client) Disconnect(ctx context.Context) error {
	if err := c.conn.Disconnect(ctx); err != nil {
		return errors.Wrap(err, "failed to disconnect mongo client")
	}

	return nil
}
