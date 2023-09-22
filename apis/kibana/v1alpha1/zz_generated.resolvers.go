/*
Copyright Elasticsearch B.V. All rights reserved.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ActionConnector.
func (mg *ActionConnector) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SpaceID),
		Extract:      resource.ExtractParamPath("space_id", true),
		Reference:    mg.Spec.ForProvider.SpaceIDRef,
		Selector:     mg.Spec.ForProvider.SpaceIDSelector,
		To: reference.To{
			List:    &SpaceList{},
			Managed: &Space{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SpaceID")
	}
	mg.Spec.ForProvider.SpaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SpaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AlertingRule.
func (mg *AlertingRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Actions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Actions[i3].ID),
			Extract:      resource.ExtractParamPath("connector_id", true),
			Reference:    mg.Spec.ForProvider.Actions[i3].IDRef,
			Selector:     mg.Spec.ForProvider.Actions[i3].IDSelector,
			To: reference.To{
				List:    &ActionConnectorList{},
				Managed: &ActionConnector{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Actions[i3].ID")
		}
		mg.Spec.ForProvider.Actions[i3].ID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Actions[i3].IDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SpaceID),
		Extract:      resource.ExtractParamPath("space_id", true),
		Reference:    mg.Spec.ForProvider.SpaceIDRef,
		Selector:     mg.Spec.ForProvider.SpaceIDSelector,
		To: reference.To{
			List:    &SpaceList{},
			Managed: &Space{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SpaceID")
	}
	mg.Spec.ForProvider.SpaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SpaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SLO.
func (mg *SLO) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SpaceID),
		Extract:      resource.ExtractParamPath("space_id", true),
		Reference:    mg.Spec.ForProvider.SpaceIDRef,
		Selector:     mg.Spec.ForProvider.SpaceIDSelector,
		To: reference.To{
			List:    &SpaceList{},
			Managed: &Space{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SpaceID")
	}
	mg.Spec.ForProvider.SpaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SpaceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Space.
func (mg *Space) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SpaceID),
		Extract:      resource.ExtractParamPath("space_id", true),
		Reference:    mg.Spec.ForProvider.SpaceIDRef,
		Selector:     mg.Spec.ForProvider.SpaceIDSelector,
		To: reference.To{
			List:    &SpaceList{},
			Managed: &Space{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SpaceID")
	}
	mg.Spec.ForProvider.SpaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SpaceIDRef = rsp.ResolvedReference

	return nil
}
