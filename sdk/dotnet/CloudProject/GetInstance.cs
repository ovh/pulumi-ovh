// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetInstance
    {
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("ovh:CloudProject/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("ovh:CloudProject/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());

        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("ovh:CloudProject/getInstance:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetInstanceArgs()
        {
        }
        public static new GetInstanceArgs Empty => new GetInstanceArgs();
    }

    public sealed class GetInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetInstanceInvokeArgs()
        {
        }
        public static new GetInstanceInvokeArgs Empty => new GetInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly ImmutableArray<Outputs.GetInstanceAddressResult> Addresses;
        public readonly ImmutableArray<Outputs.GetInstanceAttachedVolumeResult> AttachedVolumes;
        public readonly string AvailabilityZone;
        public readonly string FlavorId;
        public readonly string FlavorName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImageId;
        public readonly string InstanceId;
        public readonly string Name;
        public readonly string Region;
        public readonly string ServiceName;
        public readonly string SshKey;
        public readonly string TaskState;

        [OutputConstructor]
        private GetInstanceResult(
            ImmutableArray<Outputs.GetInstanceAddressResult> addresses,

            ImmutableArray<Outputs.GetInstanceAttachedVolumeResult> attachedVolumes,

            string availabilityZone,

            string flavorId,

            string flavorName,

            string id,

            string imageId,

            string instanceId,

            string name,

            string region,

            string serviceName,

            string sshKey,

            string taskState)
        {
            Addresses = addresses;
            AttachedVolumes = attachedVolumes;
            AvailabilityZone = availabilityZone;
            FlavorId = flavorId;
            FlavorName = flavorName;
            Id = id;
            ImageId = imageId;
            InstanceId = instanceId;
            Name = name;
            Region = region;
            ServiceName = serviceName;
            SshKey = sshKey;
            TaskState = taskState;
        }
    }
}
