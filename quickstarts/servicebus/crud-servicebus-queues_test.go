// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package servicebus

import (
	"context"
	"strings"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/resources"
	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
	"github.com/Azure/go-autorest/autorest/to"
)

func ExampleCreateServiceBusQueue() {
	context := context.Background()

	helpers.SetResourceGroupName(resourceGroupNameSuffix)

	defer resources.Cleanup(context)

	_, err := resources.CreateGroup(context, helpers.ResourceGroupName())

	if err != nil {
		helpers.PrintAndLog("Failed to create resource group.")
		helpers.PrintAndLog(err.Error())

		return
	}

	_, err = CreateServiceBusNamespace(context, serviceBusNamespaceName)

	if err != nil {
		helpers.PrintAndLog("Failed to create Service Bus Namespace.")
		helpers.PrintAndLog(err.Error())

		return
	}

	queueName := "azure-sdk-for-go-sample-queue-" + strings.ToLower(helpers.GetRandomLetterSequence(4))

	_, err = CreateServiceBusQueue(context, serviceBusNamespaceName, queueName)

	if err != nil {
		helpers.PrintAndLog("Failed to create Service Bus Queue.")
		helpers.PrintAndLog(err.Error())

		return
	}

	helpers.PrintAndLog("Created Service Bus Queue.")

	// Output:
	// Created Service Bus Queue.
}

func ExampleDeleteServiceBusQueue() {
	context := context.Background()

	helpers.SetResourceGroupName(resourceGroupNameSuffix)

	defer resources.Cleanup(context)

	_, err := resources.CreateGroup(context, helpers.ResourceGroupName())

	if err != nil {
		helpers.PrintAndLog("Failed to create resource group.")
		helpers.PrintAndLog(err.Error())

		return
	}

	_, err = CreateServiceBusNamespace(context, serviceBusNamespaceName)

	if err != nil {
		helpers.PrintAndLog("Failed to create Service Bus Namespace.")
		helpers.PrintAndLog(err.Error())

		return
	}

	queueName := "azure-sdk-for-go-sample-queue-" + strings.ToLower(helpers.GetRandomLetterSequence(4))

	_, err = CreateServiceBusQueue(context, serviceBusNamespaceName, queueName)

	if err != nil {
		helpers.PrintAndLog("Failed to create Service Bus Queue.")
		helpers.PrintAndLog(err.Error())

		return
	}

	err = DeleteServiceBusQueue(context, serviceBusNamespaceName, queueName)

	if err != nil {
		helpers.PrintAndLog("Failed to delete Service Bus Queue.")
		helpers.PrintAndLog(err.Error())

		return
	}

	helpers.PrintAndLog("Deleted Service Bus Queue.")

	// Output:
	// Deleted Service Bus Queue.
}

func ExampleUpdateServiceBusQueueProperties() {
	context := context.Background()

	helpers.SetResourceGroupName(resourceGroupNameSuffix)

	defer resources.Cleanup(context)

	_, err := resources.CreateGroup(context, helpers.ResourceGroupName())

	if err != nil {
		helpers.PrintAndLog("Failed to create resource group.")
		helpers.PrintAndLog(err.Error())

		return
	}

	_, err = CreateServiceBusNamespace(context, serviceBusNamespaceName)

	if err != nil {
		helpers.PrintAndLog("Failed to create Service Bus Namespace.")
		helpers.PrintAndLog(err.Error())

		return
	}

	queueName := "azure-sdk-for-go-sample-queue-" + strings.ToLower(helpers.GetRandomLetterSequence(4))

	_, err = CreateServiceBusQueue(context, serviceBusNamespaceName, queueName)

	if err != nil {
		helpers.PrintAndLog("Failed to create Service Bus Queue.")
		helpers.PrintAndLog(err.Error())

		return
	}

	err = UpdateServiceBusQueueProperties(
		context,
		serviceBusNamespaceName,
		queueName,
		servicebus.SBQueueProperties{
			MaxDeliveryCount:                 to.Int32Ptr(5),
			DeadLetteringOnMessageExpiration: to.BoolPtr(true),
			DefaultMessageTimeToLive:         to.StringPtr("P1D"),
		})

	if err != nil {
		helpers.PrintAndLog("Failed to update Service Bus Queue properties.")
		helpers.PrintAndLog(err.Error())

		return
	}

	helpers.PrintAndLog("Updated Service Bus Queue properties.")

	// Output:
	// Updated Service Bus Queue properties.
}
