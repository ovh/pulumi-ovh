// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    /// <summary>
    /// Provides a resource for managing **snapshot** to partitions on HA-NAS services
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
    ///     var myPartition = new Ovh.Dedicated.NasHAPartitionSnapshot("my_partition", new()
    ///     {
    ///         ServiceName = "zpool-12345",
    ///         PartitionName = "my-partition",
    ///         Type = "day-3",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// HA-NAS partition snapshot can be imported using the `{service_name}/{partition_name}/{type}`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import ovh:Dedicated/nasHAPartitionSnapshot:NasHAPartitionSnapshot my-partition zpool-12345/my-partition/day-3`
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Dedicated/nasHAPartitionSnapshot:NasHAPartitionSnapshot")]
    public partial class NasHAPartitionSnapshot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// name of the partition
        /// </summary>
        [Output("partitionName")]
        public Output<string> PartitionName { get; private set; } = null!;

        /// <summary>
        /// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Snapshot interval, allowed : day-1, day-2, day-3, day-7, hour-1, hour-6
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a NasHAPartitionSnapshot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NasHAPartitionSnapshot(string name, NasHAPartitionSnapshotArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/nasHAPartitionSnapshot:NasHAPartitionSnapshot", name, args ?? new NasHAPartitionSnapshotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NasHAPartitionSnapshot(string name, Input<string> id, NasHAPartitionSnapshotState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/nasHAPartitionSnapshot:NasHAPartitionSnapshot", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NasHAPartitionSnapshot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NasHAPartitionSnapshot Get(string name, Input<string> id, NasHAPartitionSnapshotState? state = null, CustomResourceOptions? options = null)
        {
            return new NasHAPartitionSnapshot(name, id, state, options);
        }
    }

    public sealed class NasHAPartitionSnapshotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of the partition
        /// </summary>
        [Input("partitionName", required: true)]
        public Input<string> PartitionName { get; set; } = null!;

        /// <summary>
        /// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Snapshot interval, allowed : day-1, day-2, day-3, day-7, hour-1, hour-6
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public NasHAPartitionSnapshotArgs()
        {
        }
        public static new NasHAPartitionSnapshotArgs Empty => new NasHAPartitionSnapshotArgs();
    }

    public sealed class NasHAPartitionSnapshotState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// name of the partition
        /// </summary>
        [Input("partitionName")]
        public Input<string>? PartitionName { get; set; }

        /// <summary>
        /// The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Snapshot interval, allowed : day-1, day-2, day-3, day-7, hour-1, hour-6
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NasHAPartitionSnapshotState()
        {
        }
        public static new NasHAPartitionSnapshotState Empty => new NasHAPartitionSnapshotState();
    }
}
