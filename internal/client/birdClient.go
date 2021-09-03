package client

import (
	"Dogs/internal/models"
	"context"
	proto "github.com/INEFFABLE-games/Birds/protocol"
	log "github.com/sirupsen/logrus"
)

// BirdClient struct for bird client
type BirdClient struct {
	client proto.BirdServiceClient
}

// Create send create request to Bird server
func (b BirdClient) Create(ctx context.Context, bird models.Bird) (string, error) {

	res, err := b.client.Create(ctx, &proto.CreateRequest{
		Owner: &bird.Owner,
		Name:  &bird.Name,
		Type:  &bird.Type,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "birds",
			"action":  "create",
		}).Errorf("unable to create bird %v", err)

		return "", err
	}

	return res.GetMessage(), err
}

// Get send create request to Bird server
func (b BirdClient) Get(ctx context.Context, owner string) (models.Bird, error) {

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

// Change send create request to Bird server
func (b BirdClient) Change(ctx context.Context, bird models.Bird) (string, error) {

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

// Delete send create request to Bird server
func (b BirdClient) Delete(ctx context.Context, owner string) (string, error) {

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

// NewBirdClient create new bird client object
func NewBirdClient(client proto.BirdServiceClient) *BirdClient {
	return &BirdClient{client: client}
}
