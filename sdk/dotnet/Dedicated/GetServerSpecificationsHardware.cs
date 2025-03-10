// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    public static class GetServerSpecificationsHardware
    {
        public static Task<GetServerSpecificationsHardwareResult> InvokeAsync(GetServerSpecificationsHardwareArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerSpecificationsHardwareResult>("ovh:Dedicated/getServerSpecificationsHardware:getServerSpecificationsHardware", args ?? new GetServerSpecificationsHardwareArgs(), options.WithDefaults());

        public static Output<GetServerSpecificationsHardwareResult> Invoke(GetServerSpecificationsHardwareInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerSpecificationsHardwareResult>("ovh:Dedicated/getServerSpecificationsHardware:getServerSpecificationsHardware", args ?? new GetServerSpecificationsHardwareInvokeArgs(), options.WithDefaults());

        public static Output<GetServerSpecificationsHardwareResult> Invoke(GetServerSpecificationsHardwareInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerSpecificationsHardwareResult>("ovh:Dedicated/getServerSpecificationsHardware:getServerSpecificationsHardware", args ?? new GetServerSpecificationsHardwareInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerSpecificationsHardwareArgs : global::Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetServerSpecificationsHardwareArgs()
        {
        }
        public static new GetServerSpecificationsHardwareArgs Empty => new GetServerSpecificationsHardwareArgs();
    }

    public sealed class GetServerSpecificationsHardwareInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetServerSpecificationsHardwareInvokeArgs()
        {
        }
        public static new GetServerSpecificationsHardwareInvokeArgs Empty => new GetServerSpecificationsHardwareInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerSpecificationsHardwareResult
    {
        public readonly string BootMode;
        public readonly double CoresPerProcessor;
        public readonly Outputs.GetServerSpecificationsHardwareDefaultHardwareRaidSizeResult DefaultHardwareRaidSize;
        public readonly string DefaultHardwareRaidType;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetServerSpecificationsHardwareDiskGroupResult> DiskGroups;
        public readonly ImmutableArray<Outputs.GetServerSpecificationsHardwareExpansionCardResult> ExpansionCards;
        public readonly string FormFactor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly Outputs.GetServerSpecificationsHardwareMemorySizeResult MemorySize;
        public readonly string Motherboard;
        public readonly double NumberOfProcessors;
        public readonly string ProcessorArchitecture;
        public readonly string ProcessorName;
        public readonly string ServiceName;
        public readonly double ThreadsPerProcessor;
        public readonly ImmutableArray<Outputs.GetServerSpecificationsHardwareUsbKeyResult> UsbKeys;

        [OutputConstructor]
        private GetServerSpecificationsHardwareResult(
            string bootMode,

            double coresPerProcessor,

            Outputs.GetServerSpecificationsHardwareDefaultHardwareRaidSizeResult defaultHardwareRaidSize,

            string defaultHardwareRaidType,

            string description,

            ImmutableArray<Outputs.GetServerSpecificationsHardwareDiskGroupResult> diskGroups,

            ImmutableArray<Outputs.GetServerSpecificationsHardwareExpansionCardResult> expansionCards,

            string formFactor,

            string id,

            Outputs.GetServerSpecificationsHardwareMemorySizeResult memorySize,

            string motherboard,

            double numberOfProcessors,

            string processorArchitecture,

            string processorName,

            string serviceName,

            double threadsPerProcessor,

            ImmutableArray<Outputs.GetServerSpecificationsHardwareUsbKeyResult> usbKeys)
        {
            BootMode = bootMode;
            CoresPerProcessor = coresPerProcessor;
            DefaultHardwareRaidSize = defaultHardwareRaidSize;
            DefaultHardwareRaidType = defaultHardwareRaidType;
            Description = description;
            DiskGroups = diskGroups;
            ExpansionCards = expansionCards;
            FormFactor = formFactor;
            Id = id;
            MemorySize = memorySize;
            Motherboard = motherboard;
            NumberOfProcessors = numberOfProcessors;
            ProcessorArchitecture = processorArchitecture;
            ProcessorName = processorName;
            ServiceName = serviceName;
            ThreadsPerProcessor = threadsPerProcessor;
            UsbKeys = usbKeys;
        }
    }
}
