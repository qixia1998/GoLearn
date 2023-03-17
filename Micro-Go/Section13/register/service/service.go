package service

import (
	"GoLearn/Micro-Go/Section13/register/discovery"
	"context"
	"errors"
	"log"
)

type Service interface {
	HealthCheck() string

	DiscoveryService(ctx context.Context, serviceName string) ([]*discovery.InstanceInfo, error)
}

var ErrNotServiceInstances = errors.New("instances are not existed")

type RegisterServiceImpl struct {
	discoveryClient *discovery.DiscoveryClient
}

func NewRegisterServiceImpl(discoveryClient *discovery.DiscoveryClient) Service {
	return &RegisterServiceImpl{
		discoveryClient: discoveryClient,
	}
}

func (service *RegisterServiceImpl) DiscoveryService(ctx context.Context, serviceName string) ([]*discovery.InstanceInfo, error) {

	instances, err := service.discoveryClient.DiscoverServices(ctx, serviceName)

	if err != nil {
		log.Printf("get service info err: %s", err)
	}
	if instances == nil || len(instances) == 0 {
		return nil, ErrNotServiceInstances
	}
	return instances, nil
}

func (*RegisterServiceImpl) HealthCheck() string {
	return "OK"
}
