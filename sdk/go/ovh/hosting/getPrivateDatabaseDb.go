// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an hosting privatedatabase.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Hosting"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Hosting.GetPrivateDatabaseDb(ctx, &hosting.GetPrivateDatabaseDbArgs{
//				DatabaseName: "XXXXXX",
//				ServiceName:  "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPrivateDatabaseDb(ctx *pulumi.Context, args *LookupPrivateDatabaseDbArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDatabaseDbResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateDatabaseDbResult
	err := ctx.Invoke("ovh:Hosting/getPrivateDatabaseDb:getPrivateDatabaseDb", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateDatabaseDb.
type LookupPrivateDatabaseDbArgs struct {
	// Database name
	DatabaseName string `pulumi:"databaseName"`
	// The internal name of your private database
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getPrivateDatabaseDb.
type LookupPrivateDatabaseDbResult struct {
	// Time of the next backup (every day)
	BackupTime string `pulumi:"backupTime"`
	// Creation date of the database
	CreationDate string `pulumi:"creationDate"`
	DatabaseName string `pulumi:"databaseName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Space used by the database (in MB)
	QuotaUsed   int    `pulumi:"quotaUsed"`
	ServiceName string `pulumi:"serviceName"`
	// Users granted to this database
	Users []GetPrivateDatabaseDbUser `pulumi:"users"`
}

func LookupPrivateDatabaseDbOutput(ctx *pulumi.Context, args LookupPrivateDatabaseDbOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateDatabaseDbResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPrivateDatabaseDbResultOutput, error) {
			args := v.(LookupPrivateDatabaseDbArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Hosting/getPrivateDatabaseDb:getPrivateDatabaseDb", args, LookupPrivateDatabaseDbResultOutput{}, options).(LookupPrivateDatabaseDbResultOutput), nil
		}).(LookupPrivateDatabaseDbResultOutput)
}

// A collection of arguments for invoking getPrivateDatabaseDb.
type LookupPrivateDatabaseDbOutputArgs struct {
	// Database name
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The internal name of your private database
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPrivateDatabaseDbOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseDbArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateDatabaseDb.
type LookupPrivateDatabaseDbResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateDatabaseDbResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseDbResult)(nil)).Elem()
}

func (o LookupPrivateDatabaseDbResultOutput) ToLookupPrivateDatabaseDbResultOutput() LookupPrivateDatabaseDbResultOutput {
	return o
}

func (o LookupPrivateDatabaseDbResultOutput) ToLookupPrivateDatabaseDbResultOutputWithContext(ctx context.Context) LookupPrivateDatabaseDbResultOutput {
	return o
}

// Time of the next backup (every day)
func (o LookupPrivateDatabaseDbResultOutput) BackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseDbResult) string { return v.BackupTime }).(pulumi.StringOutput)
}

// Creation date of the database
func (o LookupPrivateDatabaseDbResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseDbResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseDbResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseDbResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateDatabaseDbResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseDbResult) string { return v.Id }).(pulumi.StringOutput)
}

// Space used by the database (in MB)
func (o LookupPrivateDatabaseDbResultOutput) QuotaUsed() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseDbResult) int { return v.QuotaUsed }).(pulumi.IntOutput)
}

func (o LookupPrivateDatabaseDbResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseDbResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Users granted to this database
func (o LookupPrivateDatabaseDbResultOutput) Users() GetPrivateDatabaseDbUserArrayOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseDbResult) []GetPrivateDatabaseDbUser { return v.Users }).(GetPrivateDatabaseDbUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateDatabaseDbResultOutput{})
}
