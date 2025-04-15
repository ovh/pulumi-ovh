// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    /// <summary>
    /// Deprecated: Use ip_restriction field in cloud_project_database resource instead. Continuing to use the ovh.CloudProjectDatabase.IpRestriction resource to add an IP restriction to a cloud_project_database resource will cause the cloud_project_database resource to be updated on every apply
    /// 
    /// Apply IP restrictions to an OVHcloud Managed Database cluster.
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
    ///     var db = Ovh.CloudProjectDatabase.GetDatabase.Invoke(new()
    ///     {
    ///         ServiceName = "XXXX",
    ///         Engine = "YYYY",
    ///         Id = "ZZZZ",
    ///     });
    /// 
    ///     var ipRestriction = new Ovh.CloudProjectDatabase.IpRestriction("ip_restriction", new()
    ///     {
    ///         ServiceName = db.Apply(getDatabaseResult =&gt; getDatabaseResult.ServiceName),
    ///         Engine = db.Apply(getDatabaseResult =&gt; getDatabaseResult.Engine),
    ///         ClusterId = db.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         Ip = "178.97.6.0/24",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud Managed database cluster IP restrictions can be imported using the `service_name`, `engine`, `cluster_id` and the `ip`, separated by "/" E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:CloudProjectDatabase/ipRestriction:IpRestriction my_ip_restriction service_name/engine/cluster_id/178.97.6.0/24
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProjectDatabase/ipRestriction:IpRestriction")]
    public partial class IpRestriction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Description of the IP restriction.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Authorized IP.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Current status of the IP restriction.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a IpRestriction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpRestriction(string name, IpRestrictionArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/ipRestriction:IpRestriction", name, args ?? new IpRestrictionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpRestriction(string name, Input<string> id, IpRestrictionState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/ipRestriction:IpRestriction", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpRestriction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpRestriction Get(string name, Input<string> id, IpRestrictionState? state = null, CustomResourceOptions? options = null)
        {
            return new IpRestriction(name, id, state, options);
        }
    }

    public sealed class IpRestrictionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Description of the IP restriction.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// Authorized IP.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public IpRestrictionArgs()
        {
        }
        public static new IpRestrictionArgs Empty => new IpRestrictionArgs();
    }

    public sealed class IpRestrictionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Description of the IP restriction.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Authorized IP.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Current status of the IP restriction.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IpRestrictionState()
        {
        }
        public static new IpRestrictionState Empty => new IpRestrictionState();
    }
}
