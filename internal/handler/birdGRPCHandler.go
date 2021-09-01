package handler

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"main/internal/models"
	"main/internal/proto"
	"main/internal/service"
)

const birdHandler = "bird"

type BirdGRPCHandler struct {
	birbService service.BirdService

	proto.UnimplementedBirdServiceServer
}

func (b *BirdGRPCHandler) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateReply, error) {

	bird := models.Bird{}
	bird.Name = request.GetName()
	bird.Type = request.GetType()

	err := b.birbService.Create(ctx, bird, request.GetToken())
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "create",
		}).Errorf("unable to create bird %v", err)

	}

	mes := "bird created"
	return &proto.CreateReply{Message: &mes}, err
}

func (b *BirdGRPCHandler) Get(ctx context.Context, request *proto.GetRequest) (*proto.GetReply, error) {

	token := request.GetToken()

	resultBird, err := b.birbService.Get(ctx, token)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "get",
		}).Errorf("unable to get bird %v", err)

		return &proto.GetReply{Bird: []byte{}}, err
	}
	marshaledBird, err := json.Marshal(resultBird)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "get",
		}).Errorf("unable to marshal bird %v", err)

		return &proto.GetReply{Bird: marshaledBird}, err
	}

	return &proto.GetReply{Bird: marshaledBird}, err
}

func (b *BirdGRPCHandler) Change(ctx context.Context, request *proto.ChangeRequest) (*proto.ChangeReply, error) {

	bird := models.Bird{}
	bird.Name = request.GetName()
	bird.Type = request.GetType()

	token := request.GetToken()

	err := b.birbService.Change(ctx, bird, token)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "change",
		}).Errorf("unable to change bird %v", err)

	}

	mes := "bird changed"

	return &proto.ChangeReply{Message: &mes}, err
}

func (b *BirdGRPCHandler) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteReply, error) {

	token := request.GetToken()

	err := b.birbService.Delete(ctx, token)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": birdHandler,
			"action":  "change",
		}).Errorf("unable to change bird %v", err)

		mes := err.Error()

		return &proto.DeleteReply{Message: &mes}, err
	}

	mes := "bird deleted"

	return &proto.DeleteReply{Message: &mes}, err
}

func NewBirdGRPCHandler(birdService service.BirdService) proto.BirdServiceServer {
	return &BirdGRPCHandler{
		birbService: birdService,
	}
}
