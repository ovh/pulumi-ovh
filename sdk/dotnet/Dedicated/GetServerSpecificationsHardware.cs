// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// Use this data source to get the hardward information about a dedicated server associated with your OVHcloud Account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var spec = Ovh.Dedicated.GetServerSpecificationsHardware.Invoke(new()
        ///     {
        ///         ServiceName = "myserver",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetServerSpecificationsHardwareResult> InvokeAsync(GetServerSpecificationsHardwareArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerSpecificationsHardwareResult>("ovh:Dedicated/getServerSpecificationsHardware:getServerSpecificationsHardware", args ?? new GetServerSpecificationsHardwareArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the hardward information about a dedicated server associated with your OVHcloud Account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var spec = Ovh.Dedicated.GetServerSpecificationsHardware.Invoke(new()
        ///     {
        ///         ServiceName = "myserver",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetServerSpecificationsHardwareResult> Invoke(GetServerSpecificationsHardwareInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerSpecificationsHardwareResult>("ovh:Dedicated/getServerSpecificationsHardware:getServerSpecificationsHardware", args ?? new GetServerSpecificationsHardwareInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the hardward information about a dedicated server associated with your OVHcloud Account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var spec = Ovh.Dedicated.GetServerSpecificationsHardware.Invoke(new()
        ///     {
        ///         ServiceName = "myserver",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetServerSpecificationsHardwareResult> Invoke(GetServerSpecificationsHardwareInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerSpecificationsHardwareResult>("ovh:Dedicated/getServerSpecificationsHardware:getServerSpecificationsHardware", args ?? new GetServerSpecificationsHardwareInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerSpecificationsHardwareArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetServerSpecificationsHardwareArgs()
        {
        }
        public static new GetServerSpecificationsHardwareArgs Empty => new GetServerSpecificationsHardwareArgs();
    }

    public sealed class GetServerSpecificationsHardwareInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
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
        /// <summary>
        /// Server boot mode
        /// </summary>
        public readonly string BootMode;
        /// <summary>
        /// Number of cores per processor
        /// </summary>
        public readonly double CoresPerProcessor;
        /// <summary>
        /// Default hardware raid size for this disk group
        /// </summary>
        public readonly Outputs.GetServerSpecificationsHardwareDefaultHardwareRaidSizeResult DefaultHardwareRaidSize;
        /// <summary>
        /// Default hardware raid type for this disk group
        /// </summary>
        public readonly string DefaultHardwareRaidType;
        /// <summary>
        /// Expansion card description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Details about the groups of disks in the server
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerSpecificationsHardwareDiskGroupResult> DiskGroups;
        /// <summary>
        /// Details about the server's expansion cards
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerSpecificationsHardwareExpansionCardResult> ExpansionCards;
        /// <summary>
        /// Server form factor
        /// </summary>
        public readonly string FormFactor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// RAM capacity
        /// </summary>
        public readonly Outputs.GetServerSpecificationsHardwareMemorySizeResult MemorySize;
        /// <summary>
        /// Server motherboard
        /// </summary>
        public readonly string Motherboard;
        /// <summary>
        /// Number of processors in this dedicated server
        /// </summary>
        public readonly double NumberOfProcessors;
        /// <summary>
        /// Processor architecture bit
        /// </summary>
        public readonly string ProcessorArchitecture;
        /// <summary>
        /// Processor name
        /// </summary>
        public readonly string ProcessorName;
        public readonly string ServiceName;
        /// <summary>
        /// Number of threads per processor
        /// </summary>
        public readonly double ThreadsPerProcessor;
        /// <summary>
        /// Capacity of the USB keys installed on your server, if any
        /// </summary>
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
