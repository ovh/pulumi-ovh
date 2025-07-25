// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    public static class GetNasHAPartition
    {
        /// <summary>
        /// Use this data source to retrieve information about a dedicated HA-NAS partition.
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
        ///     var myNasHaPartition = Ovh.Dedicated.GetNasHAPartition.Invoke(new()
        ///     {
        ///         ServiceName = "zpool-12345",
        ///         Name = "my-zpool-partition",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNasHAPartitionResult> InvokeAsync(GetNasHAPartitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNasHAPartitionResult>("ovh:Dedicated/getNasHAPartition:getNasHAPartition", args ?? new GetNasHAPartitionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a dedicated HA-NAS partition.
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
        ///     var myNasHaPartition = Ovh.Dedicated.GetNasHAPartition.Invoke(new()
        ///     {
        ///         ServiceName = "zpool-12345",
        ///         Name = "my-zpool-partition",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNasHAPartitionResult> Invoke(GetNasHAPartitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNasHAPartitionResult>("ovh:Dedicated/getNasHAPartition:getNasHAPartition", args ?? new GetNasHAPartitionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a dedicated HA-NAS partition.
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
        ///     var myNasHaPartition = Ovh.Dedicated.GetNasHAPartition.Invoke(new()
        ///     {
        ///         ServiceName = "zpool-12345",
        ///         Name = "my-zpool-partition",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNasHAPartitionResult> Invoke(GetNasHAPartitionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNasHAPartitionResult>("ovh:Dedicated/getNasHAPartition:getNasHAPartition", args ?? new GetNasHAPartitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNasHAPartitionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of your dedicated HA-NAS partition.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The service_name of your dedicated HA-NAS.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetNasHAPartitionArgs()
        {
        }
        public static new GetNasHAPartitionArgs Empty => new GetNasHAPartitionArgs();
    }

    public sealed class GetNasHAPartitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of your dedicated HA-NAS partition.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The service_name of your dedicated HA-NAS.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetNasHAPartitionInvokeArgs()
        {
        }
        public static new GetNasHAPartitionInvokeArgs Empty => new GetNasHAPartitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetNasHAPartitionResult
    {
        /// <summary>
        /// Percentage of partition space used in %
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// A brief description of the partition
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// one of "NFS", "CIFS" or "NFS_CIFS"
        /// </summary>
        public readonly string Protocol;
        public readonly string ServiceName;
        /// <summary>
        /// size of the partition in GB
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// Percentage of partition space used by snapshots in %
        /// </summary>
        public readonly int UsedBySnapshots;

        [OutputConstructor]
        private GetNasHAPartitionResult(
            int capacity,

            string description,

            string id,

            string name,

            string protocol,

            string serviceName,

            int size,

            int usedBySnapshots)
        {
            Capacity = capacity;
            Description = description;
            Id = id;
            Name = name;
            Protocol = protocol;
            ServiceName = serviceName;
            Size = size;
            UsedBySnapshots = usedBySnapshots;
        }
    }
}
