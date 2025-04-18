// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

// NOTE: this file is generated via 'go:generate' - manual changes will be overwritten

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = SpringCloudAppAssociationId{}

func TestSpringCloudAppAssociationIDFormatter(t *testing.T) {
	actual := NewSpringCloudAppAssociationID("12345678-1234-9876-4563-123456789012", "resGroup1", "spring1", "app1", "bind1").ID()
	expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/bindings/bind1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestSpringCloudAppAssociationID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SpringCloudAppAssociationId
	}{
		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
			Error: true,
		},

		{
			// missing SpringName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/",
			Error: true,
		},

		{
			// missing value for SpringName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/",
			Error: true,
		},

		{
			// missing AppName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/",
			Error: true,
		},

		{
			// missing value for AppName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/",
			Error: true,
		},

		{
			// missing BindingName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/",
			Error: true,
		},

		{
			// missing value for BindingName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/bindings/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/bindings/bind1",
			Expected: &SpringCloudAppAssociationId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				ResourceGroup:  "resGroup1",
				SpringName:     "spring1",
				AppName:        "app1",
				BindingName:    "bind1",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/RESGROUP1/PROVIDERS/MICROSOFT.APPPLATFORM/SPRING/SPRING1/APPS/APP1/BINDINGS/BIND1",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := SpringCloudAppAssociationID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.SpringName != v.Expected.SpringName {
			t.Fatalf("Expected %q but got %q for SpringName", v.Expected.SpringName, actual.SpringName)
		}
		if actual.AppName != v.Expected.AppName {
			t.Fatalf("Expected %q but got %q for AppName", v.Expected.AppName, actual.AppName)
		}
		if actual.BindingName != v.Expected.BindingName {
			t.Fatalf("Expected %q but got %q for BindingName", v.Expected.BindingName, actual.BindingName)
		}
	}
}

func TestSpringCloudAppAssociationIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *SpringCloudAppAssociationId
	}{
		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
			Error: true,
		},

		{
			// missing SpringName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/",
			Error: true,
		},

		{
			// missing value for SpringName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/",
			Error: true,
		},

		{
			// missing AppName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/",
			Error: true,
		},

		{
			// missing value for AppName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/",
			Error: true,
		},

		{
			// missing BindingName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/",
			Error: true,
		},

		{
			// missing value for BindingName
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/bindings/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/bindings/bind1",
			Expected: &SpringCloudAppAssociationId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				ResourceGroup:  "resGroup1",
				SpringName:     "spring1",
				AppName:        "app1",
				BindingName:    "bind1",
			},
		},

		{
			// lower-cased segment names
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/bindings/bind1",
			Expected: &SpringCloudAppAssociationId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				ResourceGroup:  "resGroup1",
				SpringName:     "spring1",
				AppName:        "app1",
				BindingName:    "bind1",
			},
		},

		{
			// upper-cased segment names
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/SPRING/spring1/APPS/app1/BINDINGS/bind1",
			Expected: &SpringCloudAppAssociationId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				ResourceGroup:  "resGroup1",
				SpringName:     "spring1",
				AppName:        "app1",
				BindingName:    "bind1",
			},
		},

		{
			// mixed-cased segment names
			Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/SpRiNg/spring1/ApPs/app1/BiNdInGs/bind1",
			Expected: &SpringCloudAppAssociationId{
				SubscriptionId: "12345678-1234-9876-4563-123456789012",
				ResourceGroup:  "resGroup1",
				SpringName:     "spring1",
				AppName:        "app1",
				BindingName:    "bind1",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := SpringCloudAppAssociationIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.SpringName != v.Expected.SpringName {
			t.Fatalf("Expected %q but got %q for SpringName", v.Expected.SpringName, actual.SpringName)
		}
		if actual.AppName != v.Expected.AppName {
			t.Fatalf("Expected %q but got %q for AppName", v.Expected.AppName, actual.AppName)
		}
		if actual.BindingName != v.Expected.BindingName {
			t.Fatalf("Expected %q but got %q for BindingName", v.Expected.BindingName, actual.BindingName)
		}
	}
}
