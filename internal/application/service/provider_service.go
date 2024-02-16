package service

import (
	"context"
	"errors"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/model/entity"
	"github.com/bulutcan99/weekly-task-scheduler/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProviderService struct {
	providerRepo repository.IProviderRepository
}

func NewProviderService(providerRepo repository.IProviderRepository) *ProviderService {
	return &ProviderService{
		providerRepo,
	}
}

func (ps *ProviderService) GetProviders(ctx context.Context) ([]entity.Provider, error) {
	providers, err := ps.providerRepo.FindProviders(ctx)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}

		return nil, err
	}

	return providers, nil
}

func (ps *ProviderService) AddProvider(ctx context.Context, provider *entity.Provider) error {
	err := ps.providerRepo.InsertProvider(ctx, provider)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProviderService) UpdateProvider(ctx context.Context, provider *entity.Provider) error {
	err := ps.providerRepo.UpdateProvider(ctx, provider)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProviderService) DeleteProviders(ctx context.Context, name string) error {
	err := ps.providerRepo.DeleteProviders(ctx, name)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProviderService) GetProviderWithTasks(ctx context.Context, id string) (*entity.Provider, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	provider, err := ps.providerRepo.GetProviderWithTasks(ctx, objectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}

		return nil, err
	}

	return provider, nil
}

func (ps *ProviderService) GetProviderByUrl(ctx context.Context, url string) (*entity.Provider, error) {
	filter := bson.M{
		"url": url,
	}
	provider, err := ps.providerRepo.GetProvider(ctx, filter)
	if err != nil {
		return nil, err
	}
	return provider, nil
}
