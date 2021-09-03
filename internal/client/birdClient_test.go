package client

import (
	"Dogs/internal/models"
	"context"
	proto "github.com/INEFFABLE-games/Birds/protocol"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestBirdClient_Create(t *testing.T) {

	grpcConn, err := grpc.Dial("localhost:1323", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}

	client := proto.NewBirdServiceClient(grpcConn)
	birdClient := NewBirdClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	res, err := birdClient.Create(ctx, models.Bird{
		Owner: "me",
		Name:  "Silvi",
		Type:  "Phoenix",
	})

	log.Println(res)

	require.Nil(t, err)
}

func TestBirdClient_Get(t *testing.T) {

	grpcConn, err := grpc.Dial("localhost:1323", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}

	client := proto.NewBirdServiceClient(grpcConn)
	birdClient := NewBirdClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	res, err := birdClient.Get(ctx, "me")

	log.Println(res)

	require.Nil(t, err)
}

func TestBirdClient_Change(t *testing.T) {

	grpcConn, err := grpc.Dial("localhost:1323", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}

	client := proto.NewBirdServiceClient(grpcConn)
	birdClient := NewBirdClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	res, err := birdClient.Change(ctx, models.Bird{
		Owner: "me",
		Name:  "Goran",
		Type:  "Dragon",
	})

	log.Println(res)

	require.Nil(t, err)
}

func TestBirdClient_Delete(t *testing.T) {

	grpcConn, err := grpc.Dial("localhost:1323", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	}

	client := proto.NewBirdServiceClient(grpcConn)
	birdClient := NewBirdClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	res, err := birdClient.Delete(ctx, "me")

	log.Println(res)

	require.Nil(t, err)
}
