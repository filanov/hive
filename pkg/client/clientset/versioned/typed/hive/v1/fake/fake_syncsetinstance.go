// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSyncSetInstances implements SyncSetInstanceInterface
type FakeSyncSetInstances struct {
	Fake *FakeHiveV1
	ns   string
}

var syncsetinstancesResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "syncsetinstances"}

var syncsetinstancesKind = schema.GroupVersionKind{Group: "hive.openshift.io", Version: "v1", Kind: "SyncSetInstance"}

// Get takes name of the syncSetInstance, and returns the corresponding syncSetInstance object, and an error if there is any.
func (c *FakeSyncSetInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *hivev1.SyncSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(syncsetinstancesResource, c.ns, name), &hivev1.SyncSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncSetInstance), err
}

// List takes label and field selectors, and returns the list of SyncSetInstances that match those selectors.
func (c *FakeSyncSetInstances) List(ctx context.Context, opts v1.ListOptions) (result *hivev1.SyncSetInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(syncsetinstancesResource, syncsetinstancesKind, c.ns, opts), &hivev1.SyncSetInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hivev1.SyncSetInstanceList{ListMeta: obj.(*hivev1.SyncSetInstanceList).ListMeta}
	for _, item := range obj.(*hivev1.SyncSetInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested syncSetInstances.
func (c *FakeSyncSetInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(syncsetinstancesResource, c.ns, opts))

}

// Create takes the representation of a syncSetInstance and creates it.  Returns the server's representation of the syncSetInstance, and an error, if there is any.
func (c *FakeSyncSetInstances) Create(ctx context.Context, syncSetInstance *hivev1.SyncSetInstance, opts v1.CreateOptions) (result *hivev1.SyncSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(syncsetinstancesResource, c.ns, syncSetInstance), &hivev1.SyncSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncSetInstance), err
}

// Update takes the representation of a syncSetInstance and updates it. Returns the server's representation of the syncSetInstance, and an error, if there is any.
func (c *FakeSyncSetInstances) Update(ctx context.Context, syncSetInstance *hivev1.SyncSetInstance, opts v1.UpdateOptions) (result *hivev1.SyncSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(syncsetinstancesResource, c.ns, syncSetInstance), &hivev1.SyncSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncSetInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSyncSetInstances) UpdateStatus(ctx context.Context, syncSetInstance *hivev1.SyncSetInstance, opts v1.UpdateOptions) (*hivev1.SyncSetInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(syncsetinstancesResource, "status", c.ns, syncSetInstance), &hivev1.SyncSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncSetInstance), err
}

// Delete takes name of the syncSetInstance and deletes it. Returns an error if one occurs.
func (c *FakeSyncSetInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(syncsetinstancesResource, c.ns, name), &hivev1.SyncSetInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSyncSetInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(syncsetinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &hivev1.SyncSetInstanceList{})
	return err
}

// Patch applies the patch and returns the patched syncSetInstance.
func (c *FakeSyncSetInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hivev1.SyncSetInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(syncsetinstancesResource, c.ns, name, pt, data, subresources...), &hivev1.SyncSetInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SyncSetInstance), err
}
