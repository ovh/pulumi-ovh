// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Iam
{
    public static class GetReferenceResourceType
    {
        /// <summary>
        /// Use this data source to list all the IAM resource types.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// data "ovh_iam_reference_resource_type" "types" {
        /// }
        /// ```
        /// </summary>
        public static Task<GetReferenceResourceTypeResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReferenceResourceTypeResult>("ovh:Iam/getReferenceResourceType:getReferenceResourceType", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetReferenceResourceTypeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of resource types
        /// </summary>
        public readonly ImmutableArray<string> Types;

        [OutputConstructor]
        private GetReferenceResourceTypeResult(
            string id,

            ImmutableArray<string> types)
        {
            Id = id;
            Types = types;
        }
    }
}