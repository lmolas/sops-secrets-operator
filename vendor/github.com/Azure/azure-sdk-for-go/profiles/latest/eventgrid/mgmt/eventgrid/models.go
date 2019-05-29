// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package eventgrid

import original "github.com/Azure/azure-sdk-for-go/services/eventgrid/mgmt/2019-01-01/eventgrid"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type EndpointType = original.EndpointType

const (
	EndpointTypeEventHub                     EndpointType = original.EndpointTypeEventHub
	EndpointTypeEventSubscriptionDestination EndpointType = original.EndpointTypeEventSubscriptionDestination
	EndpointTypeHybridConnection             EndpointType = original.EndpointTypeHybridConnection
	EndpointTypeStorageQueue                 EndpointType = original.EndpointTypeStorageQueue
	EndpointTypeWebHook                      EndpointType = original.EndpointTypeWebHook
)

type EndpointTypeBasicDeadLetterDestination = original.EndpointTypeBasicDeadLetterDestination

const (
	EndpointTypeDeadLetterDestination EndpointTypeBasicDeadLetterDestination = original.EndpointTypeDeadLetterDestination
	EndpointTypeStorageBlob           EndpointTypeBasicDeadLetterDestination = original.EndpointTypeStorageBlob
)

type EventSubscriptionProvisioningState = original.EventSubscriptionProvisioningState

const (
	AwaitingManualAction EventSubscriptionProvisioningState = original.AwaitingManualAction
	Canceled             EventSubscriptionProvisioningState = original.Canceled
	Creating             EventSubscriptionProvisioningState = original.Creating
	Deleting             EventSubscriptionProvisioningState = original.Deleting
	Failed               EventSubscriptionProvisioningState = original.Failed
	Succeeded            EventSubscriptionProvisioningState = original.Succeeded
	Updating             EventSubscriptionProvisioningState = original.Updating
)

type ResourceRegionType = original.ResourceRegionType

const (
	GlobalResource   ResourceRegionType = original.GlobalResource
	RegionalResource ResourceRegionType = original.RegionalResource
)

type TopicProvisioningState = original.TopicProvisioningState

const (
	TopicProvisioningStateCanceled  TopicProvisioningState = original.TopicProvisioningStateCanceled
	TopicProvisioningStateCreating  TopicProvisioningState = original.TopicProvisioningStateCreating
	TopicProvisioningStateDeleting  TopicProvisioningState = original.TopicProvisioningStateDeleting
	TopicProvisioningStateFailed    TopicProvisioningState = original.TopicProvisioningStateFailed
	TopicProvisioningStateSucceeded TopicProvisioningState = original.TopicProvisioningStateSucceeded
	TopicProvisioningStateUpdating  TopicProvisioningState = original.TopicProvisioningStateUpdating
)

type TopicTypeProvisioningState = original.TopicTypeProvisioningState

const (
	TopicTypeProvisioningStateCanceled  TopicTypeProvisioningState = original.TopicTypeProvisioningStateCanceled
	TopicTypeProvisioningStateCreating  TopicTypeProvisioningState = original.TopicTypeProvisioningStateCreating
	TopicTypeProvisioningStateDeleting  TopicTypeProvisioningState = original.TopicTypeProvisioningStateDeleting
	TopicTypeProvisioningStateFailed    TopicTypeProvisioningState = original.TopicTypeProvisioningStateFailed
	TopicTypeProvisioningStateSucceeded TopicTypeProvisioningState = original.TopicTypeProvisioningStateSucceeded
	TopicTypeProvisioningStateUpdating  TopicTypeProvisioningState = original.TopicTypeProvisioningStateUpdating
)

type BaseClient = original.BaseClient
type BasicDeadLetterDestination = original.BasicDeadLetterDestination
type BasicEventSubscriptionDestination = original.BasicEventSubscriptionDestination
type DeadLetterDestination = original.DeadLetterDestination
type EventHubEventSubscriptionDestination = original.EventHubEventSubscriptionDestination
type EventHubEventSubscriptionDestinationProperties = original.EventHubEventSubscriptionDestinationProperties
type EventSubscription = original.EventSubscription
type EventSubscriptionDestination = original.EventSubscriptionDestination
type EventSubscriptionFilter = original.EventSubscriptionFilter
type EventSubscriptionFullURL = original.EventSubscriptionFullURL
type EventSubscriptionProperties = original.EventSubscriptionProperties
type EventSubscriptionUpdateParameters = original.EventSubscriptionUpdateParameters
type EventSubscriptionsClient = original.EventSubscriptionsClient
type EventSubscriptionsCreateOrUpdateFuture = original.EventSubscriptionsCreateOrUpdateFuture
type EventSubscriptionsDeleteFuture = original.EventSubscriptionsDeleteFuture
type EventSubscriptionsListResult = original.EventSubscriptionsListResult
type EventSubscriptionsUpdateFuture = original.EventSubscriptionsUpdateFuture
type EventType = original.EventType
type EventTypeProperties = original.EventTypeProperties
type EventTypesListResult = original.EventTypesListResult
type HybridConnectionEventSubscriptionDestination = original.HybridConnectionEventSubscriptionDestination
type HybridConnectionEventSubscriptionDestinationProperties = original.HybridConnectionEventSubscriptionDestinationProperties
type Operation = original.Operation
type OperationInfo = original.OperationInfo
type OperationsClient = original.OperationsClient
type OperationsListResult = original.OperationsListResult
type Resource = original.Resource
type RetryPolicy = original.RetryPolicy
type StorageBlobDeadLetterDestination = original.StorageBlobDeadLetterDestination
type StorageBlobDeadLetterDestinationProperties = original.StorageBlobDeadLetterDestinationProperties
type StorageQueueEventSubscriptionDestination = original.StorageQueueEventSubscriptionDestination
type StorageQueueEventSubscriptionDestinationProperties = original.StorageQueueEventSubscriptionDestinationProperties
type Topic = original.Topic
type TopicProperties = original.TopicProperties
type TopicRegenerateKeyRequest = original.TopicRegenerateKeyRequest
type TopicSharedAccessKeys = original.TopicSharedAccessKeys
type TopicTypeInfo = original.TopicTypeInfo
type TopicTypeProperties = original.TopicTypeProperties
type TopicTypesClient = original.TopicTypesClient
type TopicTypesListResult = original.TopicTypesListResult
type TopicUpdateParameters = original.TopicUpdateParameters
type TopicsClient = original.TopicsClient
type TopicsCreateOrUpdateFuture = original.TopicsCreateOrUpdateFuture
type TopicsDeleteFuture = original.TopicsDeleteFuture
type TopicsListResult = original.TopicsListResult
type TopicsUpdateFuture = original.TopicsUpdateFuture
type TrackedResource = original.TrackedResource
type WebHookEventSubscriptionDestination = original.WebHookEventSubscriptionDestination
type WebHookEventSubscriptionDestinationProperties = original.WebHookEventSubscriptionDestinationProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewEventSubscriptionsClient(subscriptionID string) EventSubscriptionsClient {
	return original.NewEventSubscriptionsClient(subscriptionID)
}
func NewEventSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) EventSubscriptionsClient {
	return original.NewEventSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopicTypesClient(subscriptionID string) TopicTypesClient {
	return original.NewTopicTypesClient(subscriptionID)
}
func NewTopicTypesClientWithBaseURI(baseURI string, subscriptionID string) TopicTypesClient {
	return original.NewTopicTypesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTopicsClient(subscriptionID string) TopicsClient {
	return original.NewTopicsClient(subscriptionID)
}
func NewTopicsClientWithBaseURI(baseURI string, subscriptionID string) TopicsClient {
	return original.NewTopicsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleEndpointTypeBasicDeadLetterDestinationValues() []EndpointTypeBasicDeadLetterDestination {
	return original.PossibleEndpointTypeBasicDeadLetterDestinationValues()
}
func PossibleEndpointTypeValues() []EndpointType {
	return original.PossibleEndpointTypeValues()
}
func PossibleEventSubscriptionProvisioningStateValues() []EventSubscriptionProvisioningState {
	return original.PossibleEventSubscriptionProvisioningStateValues()
}
func PossibleResourceRegionTypeValues() []ResourceRegionType {
	return original.PossibleResourceRegionTypeValues()
}
func PossibleTopicProvisioningStateValues() []TopicProvisioningState {
	return original.PossibleTopicProvisioningStateValues()
}
func PossibleTopicTypeProvisioningStateValues() []TopicTypeProvisioningState {
	return original.PossibleTopicTypeProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}