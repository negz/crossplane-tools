/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package method

import (
	"fmt"
	"go/types"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/google/go-cmp/cmp"
)

type MockObject struct {
	types.Object

	Named string
}

func (o MockObject) Name() string {
	return o.Named
}

func TestNewSetConditions(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// SetConditions of this Type.
func (t *Type) SetConditions(c ...runtime.Condition) {
	t.Status.SetConditions(c...)
}
`
	f := jen.NewFilePath("pkg")
	NewSetConditions("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewSetConditions(): -want, +got\n%s", diff)
	}
}

func TestNewGetCondition(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// GetCondition of this Type.
func (t *Type) GetCondition(ct runtime.ConditionType) runtime.Condition {
	return t.Status.GetCondition(ct)
}
`
	f := jen.NewFilePath("pkg")
	NewGetCondition("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewGetCondition(): -want, +got\n%s", diff)
	}
}

func TestNewSetProviderConfigReference(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// SetProviderConfigReference of this Type.
func (t *Type) SetProviderConfigReference(r *runtime.Reference) {
	t.Spec.Crossplane.ProviderConfigReference = r
}
`
	f := jen.NewFilePath("pkg")
	NewSetProviderConfigReference("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewSetProviderConfigReference(): -want, +got\n%s", diff)
	}
}

func TestNewGetProviderConfigReference(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// GetProviderConfigReference of this Type.
func (t *Type) GetProviderConfigReference() *runtime.Reference {
	return t.Spec.Crossplane.ProviderConfigReference
}
`
	f := jen.NewFilePath("pkg")
	NewGetProviderConfigReference("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewGetProviderConfigReference(): -want, +got\n%s", diff)
	}
}

func TestNewSetPublishConnectionDetailsTo(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// SetPublishConnectionDetailsTo of this Type.
func (t *Type) SetPublishConnectionDetailsTo(r *runtime.PublishConnectionDetailsTo) {
	t.Spec.Crossplane.PublishConnectionDetailsTo = r
}
`
	f := jen.NewFile("pkg")
	NewSetPublishConnectionDetailsTo("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewSetPublishConnectionDetailsTo(): -want, +got\n%s", diff)
	}
}

func TestNewGetPublishConnectionDetailsTo(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// GetPublishConnectionDetailsTo of this Type.
func (t *Type) GetPublishConnectionDetailsTo() *runtime.PublishConnectionDetailsTo {
	return t.Spec.Crossplane.PublishConnectionDetailsTo
}
`
	f := jen.NewFile("pkg")
	NewGetPublishConnectionDetailsTo("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewGetPublishConnectionDetailsTo(): -want, +got\n%s", diff)
	}
}

func TestNewSetManagementPolicies(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// SetManagementPolicies of this Type.
func (t *Type) SetManagementPolicies(r runtime.ManagementPolicies) {
	t.Spec.Crossplane.ManagementPolicies = r
}
`
	f := jen.NewFilePath("pkg")
	NewSetManagementPolicies("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewSetManagementPolicies(): -want, +got\n%s", diff)
	}
}

func TestNewGetManagementPolicies(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// GetManagementPolicies of this Type.
func (t *Type) GetManagementPolicies() runtime.ManagementPolicies {
	return t.Spec.Crossplane.ManagementPolicies
}
`
	f := jen.NewFilePath("pkg")
	NewGetManagementPolicies("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewGetManagementPolicies(): -want, +got\n%s", diff)
	}
}

func TestNewSetUsers(t *testing.T) {
	want := `package pkg

// SetUsers of this Type.
func (t *Type) SetUsers(i int64) {
	t.Status.Users = i
}
`
	f := jen.NewFilePath("pkg")
	NewSetUsers("t")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewSetUsers(): -want, +got\n%s", diff)
	}
}

func TestNewGetUsers(t *testing.T) {
	want := `package pkg

// GetUsers of this Type.
func (t *Type) GetUsers() int64 {
	return t.Status.Users
}
`
	f := jen.NewFilePath("pkg")
	NewGetUsers("t")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewGetUsers(): -want, +got\n%s", diff)
	}
}

func TestNewManagedGetItems(t *testing.T) {
	want := `package pkg

import resource "example.org/resource"

// GetItems of this Type.
func (t *Type) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(t.Items))
	for i := range t.Items {
		items[i] = &t.Items[i]
	}
	return items
}
`
	f := jen.NewFilePath("pkg")
	NewManagedGetItems("t", "example.org/resource")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewManagedGetItems(): -want, +got\n%s", diff)
	}
}

func TestNewSetRootProviderConfigReference(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// SetProviderConfigReference of this Type.
func (t *Type) SetProviderConfigReference(r runtime.Reference) {
	t.ProviderConfigReference = r
}
`
	f := jen.NewFilePath("pkg")
	NewSetRootProviderConfigReference("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewSetRootProviderConfigReference(): -want, +got\n%s", diff)
	}
}

func TestNewGetRootProviderConfigReference(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// GetProviderConfigReference of this Type.
func (t *Type) GetProviderConfigReference() runtime.Reference {
	return t.ProviderConfigReference
}
`
	f := jen.NewFilePath("pkg")
	NewGetRootProviderConfigReference("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewGetRootProviderConfigReference(): -want, +got\n%s", diff)
	}
}

func TestNewSetRootResourceReference(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// SetResourceReference of this Type.
func (t *Type) SetResourceReference(r runtime.TypedReference) {
	t.ResourceReference = r
}
`
	f := jen.NewFilePath("pkg")
	NewSetRootResourceReference("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewSetRootResourceReference(): -want, +got\n%s", diff)
	}
}

func TestNewGetRootResourceReference(t *testing.T) {
	want := `package pkg

import runtime "example.org/runtime"

// GetResourceReference of this Type.
func (t *Type) GetResourceReference() runtime.TypedReference {
	return t.ResourceReference
}
`
	f := jen.NewFilePath("pkg")
	NewGetRootResourceReference("t", "example.org/runtime")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewGetRootResourceReference(): -want, +got\n%s", diff)
	}
}

func TestNewProviderConfigUsageGetItems(t *testing.T) {
	want := `package pkg

import resource "example.org/resource"

// GetItems of this Type.
func (t *Type) GetItems() []resource.ProviderConfigUsage {
	items := make([]resource.ProviderConfigUsage, len(t.Items))
	for i := range t.Items {
		items[i] = &t.Items[i]
	}
	return items
}
`
	f := jen.NewFilePath("pkg")
	NewProviderConfigUsageGetItems("t", "example.org/resource")(f, MockObject{Named: "Type"})
	if diff := cmp.Diff(want, fmt.Sprintf("%#v", f)); diff != "" {
		t.Errorf("NewProviderConfigUsageGetItems(): -want, +got\n%s", diff)
	}
}
