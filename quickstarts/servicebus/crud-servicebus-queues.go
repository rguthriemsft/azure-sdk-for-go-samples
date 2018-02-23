// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package servicebus

import (
	"context"
	"fmt"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
)

// CreateServiceBusQueue creates a new Service Bus Queue in the specified Service Bus Namespace
func CreateServiceBusQueue(context context.Context, namespaceName string, queueName string) (*servicebus.SBQueue, error) {
	queuesClient := servicebus.NewQueuesClient(helpers.SubscriptionID())
	queuesClient.Authorizer, _ = getAuthorizer()

	queue, createErr := queuesClient.CreateOrUpdate(
		context,
		helpers.ResourceGroupName(),
		namespaceName,
		queueName,
		servicebus.SBQueue{})

	if createErr != nil {
		return nil, fmt.Errorf("Failed to create queue: %q", createErr.Error())
	}

	return &queue, nil
}

// UpdateServiceBusQueueProperties updates an existing Service Bus Queue with the specified properties
func UpdateServiceBusQueueProperties(context context.Context, namespaceName string, queueName string, queueProperties servicebus.SBQueueProperties) error {
	queuesClient := servicebus.NewQueuesClient(helpers.SubscriptionID())
	queuesClient.Authorizer, _ = getAuthorizer()

	_, updateErr := queuesClient.CreateOrUpdate(
		context,
		helpers.ResourceGroupName(),
		namespaceName,
		queueName,
		servicebus.SBQueue{
			SBQueueProperties: &queueProperties,
		})

	if updateErr != nil {
		return fmt.Errorf("Failed to update queue properties: %q", updateErr.Error())
	}

	return nil
}

// DeleteServiceBusQueue deletes an existing Service Bus Queue from the specified Service Bus Namespace
func DeleteServiceBusQueue(context context.Context, namespaceName string, queueName string) error {
	queuesClient := servicebus.NewQueuesClient(helpers.SubscriptionID())
	queuesClient.Authorizer, _ = getAuthorizer()

	_, deleteErr := queuesClient.Delete(
		context,
		helpers.ResourceGroupName(),
		namespaceName,
		queueName)

	if deleteErr != nil {
		return fmt.Errorf("Failed to delete queue: %q", deleteErr.Error())
	}

	return nil
}
