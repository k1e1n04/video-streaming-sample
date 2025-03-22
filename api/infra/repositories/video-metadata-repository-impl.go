package repositories

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	entities2 "github.com/k1e1n04/video-streaming-sample/api/domain/entities"
	"github.com/k1e1n04/video-streaming-sample/api/domain/repositories"
	"github.com/k1e1n04/video-streaming-sample/api/infra/records"
	"github.com/k1e1n04/video-streaming-sample/api/utils"
)

// VideoMetadataRepositoryImpl is an implementation of VideoMetadataRepository
type VideoMetadataRepositoryImpl struct {
	dynamodbClient *dynamodb.Client
}

// NewVideoMetadataRepositoryImpl is a constructor
func NewVideoMetadataRepositoryImpl(dynamodbClient *dynamodb.Client) repositories.VideoMetadataRepository {
	return &VideoMetadataRepositoryImpl{
		dynamodbClient: dynamodbClient,
	}
}

// toRecord is a method to convert VideoMetadataEntity to record
func (r *VideoMetadataRepositoryImpl) toRecord(video entities2.VideoMetadataEntity) records.VideoMetadata {
	return records.VideoMetadata{
		ID:        video.ID().Value(),
		Title:     video.Title().Value(),
		CreatedAt: utils.ToDateTimeString(video.CreatedAt()),
	}
}

// toEntity is a method to convert record to VideoMetadataEntity
func (r *VideoMetadataRepositoryImpl) toEntity(video records.VideoMetadata) (*entities2.VideoMetadataEntity, error) {
	entity, err := entities2.RestoreVideoMetadataEntity(
		video.ID,
		video.Title,
		video.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Register is a method to register video metadata
func (r *VideoMetadataRepositoryImpl) Register(ctx context.Context, video entities2.VideoMetadataEntity) error {
	record := r.toRecord(video)

	_, err := r.dynamodbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(records.VideoMetadataTableName),
		Item: map[string]types.AttributeValue{
			"id":         &types.AttributeValueMemberS{Value: record.ID},
			"title":      &types.AttributeValueMemberS{Value: record.Title},
			"created_at": &types.AttributeValueMemberS{Value: record.CreatedAt},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// FindByID is a method to find video metadata by id
func (r *VideoMetadataRepositoryImpl) FindByID(ctx context.Context, videoID entities2.VideoID) (*entities2.VideoMetadataEntity, error) {
	var record records.VideoMetadata
	result, err := r.dynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(records.VideoMetadataTableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: videoID.Value()},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}
	err = attributevalue.UnmarshalMap(result.Item, &record)
	if err != nil {
		return nil, err
	}

	video, err := r.toEntity(record)
	if err != nil {
		return nil, err
	}
	return video, nil
}

// FindPage is a method to find all video metadata
func (r *VideoMetadataRepositoryImpl) FindPage(ctx context.Context, limit int32, lastEvaluatedKey *string) (*utils.Pageable[entities2.VideoMetadataEntity], error) {
	var videos []entities2.VideoMetadataEntity
	var lastEvaluatedKeyMap map[string]types.AttributeValue

	if lastEvaluatedKey != nil {
		lastEvaluatedKeyMap = map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: *lastEvaluatedKey},
		}
	}

	result, err := r.dynamodbClient.Scan(ctx, &dynamodb.ScanInput{
		TableName:            aws.String(records.VideoMetadataTableName),
		Limit:                aws.Int32(int32(limit)),
		ExclusiveStartKey:    lastEvaluatedKeyMap,
		ProjectionExpression: aws.String("id, title, created_at"),
	})
	if err != nil {
		return nil, err
	}

	for _, item := range result.Items {
		var record records.VideoMetadata
		err = attributevalue.UnmarshalMap(item, &record)
		if err != nil {
			return nil, err
		}

		video, err := r.toEntity(record)
		if err != nil {
			return nil, err
		}
		videos = append(videos, *video)
	}

	// get lastEvaluatedKey
	var lastEvaluatedKeyStr *string
	if result.LastEvaluatedKey != nil {
		if val, ok := result.LastEvaluatedKey["id"].(*types.AttributeValueMemberS); ok {
			lastEvaluatedKeyStr = &val.Value
		}
	} else {
		lastEvaluatedKeyStr = nil
	}

	return utils.NewPageable(videos, lastEvaluatedKeyStr), nil
}
