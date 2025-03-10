// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    public static class GetNasHA
    {
        public static Task<GetNasHAResult> InvokeAsync(GetNasHAArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNasHAResult>("ovh:Dedicated/getNasHA:getNasHA", args ?? new GetNasHAArgs(), options.WithDefaults());

        public static Output<GetNasHAResult> Invoke(GetNasHAInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNasHAResult>("ovh:Dedicated/getNasHA:getNasHA", args ?? new GetNasHAInvokeArgs(), options.WithDefaults());

        public static Output<GetNasHAResult> Invoke(GetNasHAInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNasHAResult>("ovh:Dedicated/getNasHA:getNasHA", args ?? new GetNasHAInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNasHAArgs : global::Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetNasHAArgs()
        {
        }
        public static new GetNasHAArgs Empty => new GetNasHAArgs();
    }

    public sealed class GetNasHAInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetNasHAInvokeArgs()
        {
        }
        public static new GetNasHAInvokeArgs Empty => new GetNasHAInvokeArgs();
    }


    [OutputType]
    public sealed class GetNasHAResult
    {
        public readonly string NasHAURN;
        public readonly bool CanCreatePartition;
        public readonly string CustomName;
        public readonly string Datacenter;
        public readonly string DiskType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Ip;
        public readonly bool Monitored;
        public readonly ImmutableArray<string> PartitionsLists;
        public readonly string ServiceName;
        public readonly double ZpoolCapacity;
        public readonly double ZpoolSize;

        [OutputConstructor]
        private GetNasHAResult(
            string NasHAURN,

            bool canCreatePartition,

            string customName,

            string datacenter,

            string diskType,

            string id,

            string ip,

            bool monitored,

            ImmutableArray<string> partitionsLists,

            string serviceName,

            double zpoolCapacity,

            double zpoolSize)
        {
            this.NasHAURN = NasHAURN;
            CanCreatePartition = canCreatePartition;
            CustomName = customName;
            Datacenter = datacenter;
            DiskType = diskType;
            Id = id;
            Ip = ip;
            Monitored = monitored;
            PartitionsLists = partitionsLists;
            ServiceName = serviceName;
            ZpoolCapacity = zpoolCapacity;
            ZpoolSize = zpoolSize;
        }
    }
}
