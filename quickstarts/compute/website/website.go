// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package web

import (
	"context"
	"fmt"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/iam"
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2016-09-01/web"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

func getAppServiceClient() web.AppServicePlansClient {
	token, _ := iam.GetResourceManagementToken(iam.AuthGrantType())
	appServiceClient := web.NewAppServicePlansClient(helpers.SubscriptionID())
	appServiceClient.Authorizer = autorest.NewBearerAuthorizer(token)
	appServiceClient.AddToUserAgent(helpers.UserAgent())
	return appServiceClient
}

func getAppsClient() web.AppsClient {
	token, _ := iam.GetResourceManagementToken(iam.AuthGrantType())
	webAppsClient := web.NewAppsClient(helpers.SubscriptionID())
	webAppsClient.Authorizer = autorest.NewBearerAuthorizer(token)
	webAppsClient.AddToUserAgent(helpers.UserAgent())
	return webAppsClient
}

// CreateAppServicePlan creates a new App Service Plan
func CreateAppServicePlan(ctx context.Context, webName string) (appserviceplan web.AppServicePlan, err error) {
	appServiceClient := getAppServiceClient()
	future, err := appServiceClient.CreateOrUpdate(
		ctx,
		helpers.ResourceGroupName(),
		webName,
		web.AppServicePlan{
			Sku: &web.SkuDescription{
				Name:     to.StringPtr("B1"),
				Tier:     to.StringPtr("Basic"),
				Capacity: to.Int32Ptr(1),
			},
		},
	)
	if err != nil {
		return appserviceplan, fmt.Errorf("cannot create Web: %v", err)
	}

	err = future.WaitForCompletion(ctx, appServiceClient.Client)
	if err != nil {
		return appserviceplan, fmt.Errorf("cannot get the web create or update future response: %v", err)
	}

	return future.Result(appServiceClient)

}

// CreateWebSite creates a new web site app
func CreateWebSite(ctx context.Context, webName string) (website web.Site, err error) {
	var tags map[string]*string
	webAppsClient := getAppsClient()
	future, err := webAppsClient.CreateOrUpdate(
		ctx,
		"createorupdate",
		webName,
		web.Site{
			Name:           to.StringPtr(webName),
			Location:       to.StringPtr("southcentralus"),
			Kind:           to.StringPtr("app"),
			Type:           to.StringPtr("Microsoft.Web/sites"),
			Tags:           tags,
			SiteProperties: &web.SiteProperties{},
		},
		to.BoolPtr(false),
		to.BoolPtr(false),
		to.BoolPtr(true),
		"600",
	)
	if err != nil {
		return website, fmt.Errorf("cannot create Web: %v", err)
	}

	err = future.WaitForCompletion(ctx, webAppsClient.Client)
	if err != nil {
		return website, fmt.Errorf("cannot get the web create or update future response: %v", err)
	}

	return future.Result(webAppsClient)
}

// GetAppServicePlan gets the specified App Service plan info
func GetAppServicePlan(ctx context.Context, appServicePlanName string) (web.AppServicePlan, error) {
	appServiceClient := getAppServiceClient()
	return appServiceClient.Get(ctx, helpers.ResourceGroupName(), appServicePlanName)
}

// GetWebSite gets the specified WebSite info
func GetWebSite(ctx context.Context, webAppsName string) (web.Site, error) {
	webAppsClient := getAppsClient()
	return webAppsClient.Get(ctx, helpers.ResourceGroupName(), webAppsName)
}
