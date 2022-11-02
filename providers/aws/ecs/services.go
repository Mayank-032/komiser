package instances

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	. "github.com/mlabouardy/komiser/models"
)

func Services(ctx context.Context, cfg aws.Config, account string) ([]Resource, error) {
	resources := make([]Resource, 0)
	var config ecs.ListServicesInput
	ecsClient := ecs.NewFromConfig(cfg)
	output, err := ecsClient.ListServices(context.Background(), &config)
	if err != nil {
		return resources, err
	}

	for _, service := range output.ServiceArns {
		resources = append(resources, Resource{
			Provider:  "AWS",
			Account:   account,
			Service:   "ECS Service",
			Region:    cfg.Region,
			Name:      service,
			Cost:      0,
			FetchedAt: time.Now(),
		})

		if aws.ToString(output.NextToken) == "" {
			break
		}

		config.NextToken = output.NextToken
	}
	log.Printf("[%s] Fetched %d AWS ECS services from %s\n", account, len(resources), cfg.Region)
	return resources, nil
}
