package snsiface

import (
	"github.com/awslabs/aws-sdk-go/service/sns"
)

type SNSAPI interface {
	AddPermission(*sns.AddPermissionInput) (*sns.AddPermissionOutput, error)

	ConfirmSubscription(*sns.ConfirmSubscriptionInput) (*sns.ConfirmSubscriptionOutput, error)

	CreatePlatformApplication(*sns.CreatePlatformApplicationInput) (*sns.CreatePlatformApplicationOutput, error)

	CreatePlatformEndpoint(*sns.CreatePlatformEndpointInput) (*sns.CreatePlatformEndpointOutput, error)

	CreateTopic(*sns.CreateTopicInput) (*sns.CreateTopicOutput, error)

	DeleteEndpoint(*sns.DeleteEndpointInput) (*sns.DeleteEndpointOutput, error)

	DeletePlatformApplication(*sns.DeletePlatformApplicationInput) (*sns.DeletePlatformApplicationOutput, error)

	DeleteTopic(*sns.DeleteTopicInput) (*sns.DeleteTopicOutput, error)

	GetEndpointAttributes(*sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error)

	GetPlatformApplicationAttributes(*sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error)

	GetSubscriptionAttributes(*sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error)

	GetTopicAttributes(*sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error)

	ListEndpointsByPlatformApplication(*sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error)

	ListPlatformApplications(*sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error)

	ListSubscriptions(*sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error)

	ListSubscriptionsByTopic(*sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error)

	ListTopics(*sns.ListTopicsInput) (*sns.ListTopicsOutput, error)

	Publish(*sns.PublishInput) (*sns.PublishOutput, error)

	RemovePermission(*sns.RemovePermissionInput) (*sns.RemovePermissionOutput, error)

	SetEndpointAttributes(*sns.SetEndpointAttributesInput) (*sns.SetEndpointAttributesOutput, error)

	SetPlatformApplicationAttributes(*sns.SetPlatformApplicationAttributesInput) (*sns.SetPlatformApplicationAttributesOutput, error)

	SetSubscriptionAttributes(*sns.SetSubscriptionAttributesInput) (*sns.SetSubscriptionAttributesOutput, error)

	SetTopicAttributes(*sns.SetTopicAttributesInput) (*sns.SetTopicAttributesOutput, error)

	Subscribe(*sns.SubscribeInput) (*sns.SubscribeOutput, error)

	Unsubscribe(*sns.UnsubscribeInput) (*sns.UnsubscribeOutput, error)
}