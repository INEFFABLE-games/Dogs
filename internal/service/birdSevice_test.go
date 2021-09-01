package service

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main/internal/models"
	"main/internal/repository"
	"testing"
	"time"
)

func initializeDataBases() (*mongo.Client, *sql.DB) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://dbuser1:dbuser1@petscluster.82emx.mongodb.net/Birds?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//connect to mongo
	mongoConn, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "create",
		}).Errorf("unable connect to MonogDB %v", err)
	}

	//connect to postgres
	postgresconn, err := sql.Open("postgres", "port=5432 host=localhost user=postgres password=12345 dbname=dogs sslmode=disable")
	if err != nil {
		log.WithFields(log.Fields{
			"hadler": "Test_UserRepository",
			"action": "Login",
		}).Errorf("unable connect to database %v", err)
	}
	return mongoConn, postgresconn
}

func TestBirdService_Create(t *testing.T) {
	mongoConn, postgresConn := initializeDataBases()

	serv := NewBirdService(*repository.NewBirdsMongoRepository(mongoConn, postgresConn))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3600)
	defer cancel()

	err := serv.Create(ctx, models.Bird{
		Owner: "",
		Name:  "Phoenix",
		Type:  "Firebird",
	}, "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJMb2dpbiI6InRlc3R1c2VyMSIsImV4cCI6MTYzMDQxMTMyM30")

	require.Nil(t, err)
}

func TestBirdService_Get(t *testing.T) {
	mongoConn, postgresConn := initializeDataBases()

	serv := NewBirdService(*repository.NewBirdsMongoRepository(mongoConn, postgresConn))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3600)
	defer cancel()

	resultBird, err := serv.Get(ctx, "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJMb2dpbiI6InRlc3R1c2VyMSIsImV4cCI6MTYzMDQxMTMyM30")
	if err != nil {
		log.Errorf("unable to get bird %v", err)
	}

	require.Equal(t, models.Bird{
		Owner: "testuser1",
		Name:  "Phoenix",
		Type:  "Firebird",
	}, resultBird)
}

func TestBirdService_Change(t *testing.T) {
	mongoConn, postgresConn := initializeDataBases()

	serv := NewBirdService(*repository.NewBirdsMongoRepository(mongoConn, postgresConn))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3600)
	defer cancel()

	err := serv.Change(ctx, models.Bird{
		Owner: "",
		Name:  "Phoenix",
		Type:  "Firebird",
	},
		"eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJMb2dpbiI6InRlc3R1c2VyMSIsImV4cCI6MTYzMDQxMTMyM30")

	require.Nil(t, err)
}

func TestBirdService_Delete(t *testing.T) {
	mongoConn, postgresConn := initializeDataBases()

	serv := NewBirdService(*repository.NewBirdsMongoRepository(mongoConn, postgresConn))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3600)
	defer cancel()

	err := serv.Delete(ctx, "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJMb2dpbiI6InRlc3R1c2VyMSIsImV4cCI6MTYzMDQxMTMyM30")

	require.Nil(t, err)
}
