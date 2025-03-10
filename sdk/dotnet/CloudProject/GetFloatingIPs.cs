// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetFloatingIPs
    {
        public static Task<GetFloatingIPsResult> InvokeAsync(GetFloatingIPsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFloatingIPsResult>("ovh:CloudProject/getFloatingIPs:getFloatingIPs", args ?? new GetFloatingIPsArgs(), options.WithDefaults());

        public static Output<GetFloatingIPsResult> Invoke(GetFloatingIPsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFloatingIPsResult>("ovh:CloudProject/getFloatingIPs:getFloatingIPs", args ?? new GetFloatingIPsInvokeArgs(), options.WithDefaults());

        public static Output<GetFloatingIPsResult> Invoke(GetFloatingIPsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFloatingIPsResult>("ovh:CloudProject/getFloatingIPs:getFloatingIPs", args ?? new GetFloatingIPsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFloatingIPsArgs : global::Pulumi.InvokeArgs
    {
        [Input("regionName", required: true)]
        public string RegionName { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetFloatingIPsArgs()
        {
        }
        public static new GetFloatingIPsArgs Empty => new GetFloatingIPsArgs();
    }

    public sealed class GetFloatingIPsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetFloatingIPsInvokeArgs()
        {
        }
        public static new GetFloatingIPsInvokeArgs Empty => new GetFloatingIPsInvokeArgs();
    }


    [OutputType]
    public sealed class GetFloatingIPsResult
    {
        public readonly ImmutableArray<Outputs.GetFloatingIPsCloudProjectFloatingipResult> CloudProjectFloatingips;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RegionName;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetFloatingIPsResult(
            ImmutableArray<Outputs.GetFloatingIPsCloudProjectFloatingipResult> cloudProjectFloatingips,

            string id,

            string regionName,

            string serviceName)
        {
            CloudProjectFloatingips = cloudProjectFloatingips;
            Id = id;
            RegionName = regionName;
            ServiceName = serviceName;
        }
    }
}
