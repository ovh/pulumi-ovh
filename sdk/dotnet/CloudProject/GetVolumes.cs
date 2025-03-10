// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetVolumes
    {
        public static Task<GetVolumesResult> InvokeAsync(GetVolumesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVolumesResult>("ovh:CloudProject/getVolumes:getVolumes", args ?? new GetVolumesArgs(), options.WithDefaults());

        public static Output<GetVolumesResult> Invoke(GetVolumesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumesResult>("ovh:CloudProject/getVolumes:getVolumes", args ?? new GetVolumesInvokeArgs(), options.WithDefaults());

        public static Output<GetVolumesResult> Invoke(GetVolumesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVolumesResult>("ovh:CloudProject/getVolumes:getVolumes", args ?? new GetVolumesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVolumesArgs : global::Pulumi.InvokeArgs
    {
        [Input("regionName", required: true)]
        public string RegionName { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetVolumesArgs()
        {
        }
        public static new GetVolumesArgs Empty => new GetVolumesArgs();
    }

    public sealed class GetVolumesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetVolumesInvokeArgs()
        {
        }
        public static new GetVolumesInvokeArgs Empty => new GetVolumesInvokeArgs();
    }


    [OutputType]
    public sealed class GetVolumesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RegionName;
        public readonly string ServiceName;
        public readonly ImmutableArray<Outputs.GetVolumesVolumeResult> Volumes;

        [OutputConstructor]
        private GetVolumesResult(
            string id,

            string regionName,

            string serviceName,

            ImmutableArray<Outputs.GetVolumesVolumeResult> volumes)
        {
            Id = id;
            RegionName = regionName;
            ServiceName = serviceName;
            Volumes = volumes;
        }
    }
}
