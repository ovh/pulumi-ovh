// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to list all the IAM resource types.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Iam"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Iam.GetReferenceResourceType(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetReferenceResourceType(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetReferenceResourceTypeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetReferenceResourceTypeResult
	err := ctx.Invoke("ovh:Iam/getReferenceResourceType:getReferenceResourceType", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getReferenceResourceType.
type GetReferenceResourceTypeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of resource types
	Types []string `pulumi:"types"`
}
