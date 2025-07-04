// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    /// <summary>
    /// Creates a namespace for a M3DB cluster associated with a public cloud project.
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
    ///     var m3db = Ovh.CloudProjectDatabase.GetDatabase.Invoke(new()
    ///     {
    ///         ServiceName = "XXX",
    ///         Engine = "m3db",
    ///         Id = "ZZZ",
    ///     });
    /// 
    ///     var @namespace = new Ovh.CloudProjectDatabase.M3DbNamespace("namespace", new()
    ///     {
    ///         ServiceName = m3db.Apply(getDatabaseResult =&gt; getDatabaseResult.ServiceName),
    ///         ClusterId = m3db.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         Name = "mynamespace",
    ///         Resolution = "P2D",
    ///         RetentionPeriodDuration = "PT48H",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud Managed M3DB clusters namespaces can be imported using the `service_name`, `cluster_id` and `id` of the namespace, separated by "/" E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace my_namespace service_name/cluster_id/id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace")]
    public partial class M3DbNamespace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Name of the namespace. A namespace named "default" is mapped with already created default namespace instead of creating a new namespace.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Output("resolution")]
        public Output<string> Resolution { get; private set; } = null!;

        /// <summary>
        /// Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Output("retentionBlockDataExpirationDuration")]
        public Output<string?> RetentionBlockDataExpirationDuration { get; private set; } = null!;

        /// <summary>
        /// Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Output("retentionBlockSizeDuration")]
        public Output<string> RetentionBlockSizeDuration { get; private set; } = null!;

        /// <summary>
        /// Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Output("retentionBufferFutureDuration")]
        public Output<string?> RetentionBufferFutureDuration { get; private set; } = null!;

        /// <summary>
        /// Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Output("retentionBufferPastDuration")]
        public Output<string?> RetentionBufferPastDuration { get; private set; } = null!;

        /// <summary>
        /// Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Output("retentionPeriodDuration")]
        public Output<string> RetentionPeriodDuration { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Defines whether M3DB will create snapshot files for this namespace.
        /// </summary>
        [Output("snapshotEnabled")]
        public Output<bool> SnapshotEnabled { get; private set; } = null!;

        /// <summary>
        /// Type of namespace.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Defines whether M3DB will include writes to this namespace in the commit log.
        /// </summary>
        [Output("writesToCommitLogEnabled")]
        public Output<bool> WritesToCommitLogEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a M3DbNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public M3DbNamespace(string name, M3DbNamespaceArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace", name, args ?? new M3DbNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private M3DbNamespace(string name, Input<string> id, M3DbNamespaceState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing M3DbNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static M3DbNamespace Get(string name, Input<string> id, M3DbNamespaceState? state = null, CustomResourceOptions? options = null)
        {
            return new M3DbNamespace(name, id, state, options);
        }
    }

    public sealed class M3DbNamespaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the namespace. A namespace named "default" is mapped with already created default namespace instead of creating a new namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("resolution", required: true)]
        public Input<string> Resolution { get; set; } = null!;

        /// <summary>
        /// Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBlockDataExpirationDuration")]
        public Input<string>? RetentionBlockDataExpirationDuration { get; set; }

        /// <summary>
        /// Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBlockSizeDuration")]
        public Input<string>? RetentionBlockSizeDuration { get; set; }

        /// <summary>
        /// Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBufferFutureDuration")]
        public Input<string>? RetentionBufferFutureDuration { get; set; }

        /// <summary>
        /// Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBufferPastDuration")]
        public Input<string>? RetentionBufferPastDuration { get; set; }

        /// <summary>
        /// Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionPeriodDuration")]
        public Input<string>? RetentionPeriodDuration { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Defines whether M3DB will create snapshot files for this namespace.
        /// </summary>
        [Input("snapshotEnabled")]
        public Input<bool>? SnapshotEnabled { get; set; }

        /// <summary>
        /// Defines whether M3DB will include writes to this namespace in the commit log.
        /// </summary>
        [Input("writesToCommitLogEnabled")]
        public Input<bool>? WritesToCommitLogEnabled { get; set; }

        public M3DbNamespaceArgs()
        {
        }
        public static new M3DbNamespaceArgs Empty => new M3DbNamespaceArgs();
    }

    public sealed class M3DbNamespaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Name of the namespace. A namespace named "default" is mapped with already created default namespace instead of creating a new namespace.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("resolution")]
        public Input<string>? Resolution { get; set; }

        /// <summary>
        /// Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBlockDataExpirationDuration")]
        public Input<string>? RetentionBlockDataExpirationDuration { get; set; }

        /// <summary>
        /// Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBlockSizeDuration")]
        public Input<string>? RetentionBlockSizeDuration { get; set; }

        /// <summary>
        /// Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBufferFutureDuration")]
        public Input<string>? RetentionBufferFutureDuration { get; set; }

        /// <summary>
        /// Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionBufferPastDuration")]
        public Input<string>? RetentionBufferPastDuration { get; set; }

        /// <summary>
        /// Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
        /// </summary>
        [Input("retentionPeriodDuration")]
        public Input<string>? RetentionPeriodDuration { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Defines whether M3DB will create snapshot files for this namespace.
        /// </summary>
        [Input("snapshotEnabled")]
        public Input<bool>? SnapshotEnabled { get; set; }

        /// <summary>
        /// Type of namespace.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Defines whether M3DB will include writes to this namespace in the commit log.
        /// </summary>
        [Input("writesToCommitLogEnabled")]
        public Input<bool>? WritesToCommitLogEnabled { get; set; }

        public M3DbNamespaceState()
        {
        }
        public static new M3DbNamespaceState Empty => new M3DbNamespaceState();
    }
}
