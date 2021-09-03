package service

import (
	"Dogs/internal/models"
	"context"
	proto "github.com/INEFFABLE-games/Birds/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// BirdService struct for bird client
type BirdService struct {
	client proto.BirdServiceClient
}

// Create send create request to Bird server
func (b BirdService) Create(ctx context.Context, bird models.Bird) (string, error) {

	res, err := b.client.Create(ctx, &proto.CreateRequest{
		Owner: &bird.Owner,
		Name:  &bird.Name,
		Type:  &bird.Type,
	})

	return res.GetMessage(), err
}

// Get send get request to Bird server
func (b BirdService) Get(ctx context.Context, owner string) (models.Bird, error) {

	res, err := b.client.Get(ctx, &proto.GetRequest{Owner: &owner})
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "get",
		}).Errorf("unable to get bird %v", err)

		return models.Bird{}, err
	}

	result := models.Bird{
		Owner: res.GetOwner(),
		Name:  res.GetName(),
		Type:  res.GetType()}

	return result, err
}

// Change send change request to Bird server
func (b BirdService) Change(ctx context.Context, bird models.Bird) (string, error) {

	res, err := b.client.Change(ctx, &proto.ChangeRequest{
		Owner: &bird.Owner,
		Name:  &bird.Name,
		Type:  &bird.Type,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "change",
		}).Errorf("unable to cahnge bird %v", err)

		return res.GetMessage(), err
	}

	return res.GetMessage(), err
}

// Delete send delete request to Bird server
func (b BirdService) Delete(ctx context.Context, owner string) (string, error) {

	res, err := b.client.Delete(ctx, &proto.DeleteRequest{Owner: &owner})
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "delete",
		}).Errorf("unable to delete bird %v", err)

		return res.GetMessage(), err
	}

	return res.GetMessage(), err
}

// NewBirdService create new bird client object
func NewBirdService(grpconn *grpc.ClientConn) *BirdService {
	client := proto.NewBirdServiceClient(grpconn)
	return &BirdService{client: client}
}
