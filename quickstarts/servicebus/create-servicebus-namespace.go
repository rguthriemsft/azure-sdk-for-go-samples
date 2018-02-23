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
	"github.com/Azure/go-autorest/autorest/to"
)

// CreateServiceBusNamespace creates a new Service Bus namespace
func CreateServiceBusNamespace(context context.Context, namespaceName string) (*servicebus.SBNamespace, error) {
	namespacesClient := servicebus.NewNamespacesClient(helpers.SubscriptionID())
	namespacesClient.Authorizer, _ = getAuthorizer()

	createFuture, createErr := namespacesClient.CreateOrUpdate(
		context,
		helpers.ResourceGroupName(),
		namespaceName,
		servicebus.SBNamespace{
			Location: to.StringPtr(helpers.Location()),
			Tags:     *to.StringMapPtr(map[string]string{"sdk-sample": "golang"}),
			Sku: &servicebus.SBSku{
				Name: "Standard",
				Tier: "Standard",
			},
		})

	if createErr != nil {
		return nil, fmt.Errorf("Failed to create Service Bus Namespace: %q", createErr.Error())
	}

	createWaitErr := createFuture.WaitForCompletion(context, namespacesClient.Client)

	if createWaitErr != nil {
		return nil, fmt.Errorf("Failed to create Service Bus Namespace: %q", createWaitErr.Error())
	}

	sbNamespace, createResultErr := createFuture.Result(namespacesClient)

	if createResultErr != nil {
		return nil, fmt.Errorf("Failed to create Service Bus Namespace: %q", createResultErr.Error())
	}

	return &sbNamespace, nil
}
