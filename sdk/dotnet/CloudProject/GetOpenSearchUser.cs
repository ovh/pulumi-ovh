// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetOpenSearchUser
    {
        public static Task<GetOpenSearchUserResult> InvokeAsync(GetOpenSearchUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOpenSearchUserResult>("ovh:CloudProject/getOpenSearchUser:getOpenSearchUser", args ?? new GetOpenSearchUserArgs(), options.WithDefaults());

        public static Output<GetOpenSearchUserResult> Invoke(GetOpenSearchUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOpenSearchUserResult>("ovh:CloudProject/getOpenSearchUser:getOpenSearchUser", args ?? new GetOpenSearchUserInvokeArgs(), options.WithDefaults());

        public static Output<GetOpenSearchUserResult> Invoke(GetOpenSearchUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOpenSearchUserResult>("ovh:CloudProject/getOpenSearchUser:getOpenSearchUser", args ?? new GetOpenSearchUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOpenSearchUserArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetOpenSearchUserArgs()
        {
        }
        public static new GetOpenSearchUserArgs Empty => new GetOpenSearchUserArgs();
    }

    public sealed class GetOpenSearchUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetOpenSearchUserInvokeArgs()
        {
        }
        public static new GetOpenSearchUserInvokeArgs Empty => new GetOpenSearchUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetOpenSearchUserResult
    {
        public readonly ImmutableArray<Outputs.GetOpenSearchUserAclResult> Acls;
        public readonly string ClusterId;
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string ServiceName;
        public readonly string Status;

        [OutputConstructor]
        private GetOpenSearchUserResult(
            ImmutableArray<Outputs.GetOpenSearchUserAclResult> acls,

            string clusterId,

            string createdAt,

            string id,

            string name,

            string serviceName,

            string status)
        {
            Acls = acls;
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            ServiceName = serviceName;
            Status = status;
        }
    }
}
