// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: external/data_feed/data_feed.proto

package data_feed

import policy "github.com/chef/automate/api/external/iam/v2/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.datafeed.DatafeedService/AddDestination", "datafeed:destination", "datafeed:destination:add", "POST", "/api/v0/datafeed/destination", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*AddDestinationRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "url":
					return m.Url
				case "secret":
					return m.Secret
				case "services":
					return m.Services
				case "integration_types":
					return m.IntegrationTypes
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.datafeed.DatafeedService/GetDestination", "datafeed:destination:{id}", "datafeed:destination:get", "GET", "/api/v0/datafeed/destination/{id}", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.datafeed.DatafeedService/DeleteDestination", "destination:destination:{id}", "destination:destination:delete", "DELETE", "/api/v0/datafeed/destination/{id}", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.datafeed.DatafeedService/UpdateDestination", "datafeed:destination:{id}", "datafeed:destination:update", "PATCH", "/api/v0/datafeed/destination/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*UpdateDestinationRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "url":
					return m.Url
				case "secret":
					return m.Secret
				case "services":
					return m.Services
				case "integration_types":
					return m.IntegrationTypes
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.datafeed.DatafeedService/ListDestinations", "datafeed:destinations", "datafeed:destinations:list", "POST", "/api/v0/datafeed/destinations", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.datafeed.DatafeedService/TestDestination", "datafeed:destinations:test", "datafeed:destinations:test", "POST", "/api/v0/datafeed/destinations/test", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*URLValidationRequest); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "url":
					return m.Url
				case "services":
					return m.Services
				case "integration_types":
					return m.IntegrationTypes
				default:
					return ""
				}
			})
		}
		return ""
	})
}
