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
    /// Creates a schema registry ACL for a Kafka cluster associated with a public cloud project.
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
    ///     var kafka = Ovh.CloudProjectDatabase.GetDatabase.Invoke(new()
    ///     {
    ///         ServiceName = "XXX",
    ///         Engine = "kafka",
    ///         Id = "ZZZ",
    ///     });
    /// 
    ///     var schemaRegistryAcl = new Ovh.CloudProjectDatabase.KafkaSchemaRegistryAcl("schemaRegistryAcl", new()
    ///     {
    ///         ServiceName = kafka.Apply(getDatabaseResult =&gt; getDatabaseResult.ServiceName),
    ///         ClusterId = kafka.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         Permission = "schema_registry_read",
    ///         Resource = "Subject:myResource",
    ///         Username = "johndoe",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud Managed Kafka clusters schema registry ACLs can be imported using the `service_name`, `cluster_id` and `id` of the schema registry ACL, separated by "/" E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl my_schemaRegistryAcl service_name/cluster_id/id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl")]
    public partial class KafkaSchemaRegistryAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Permission to give to this username on this resource. Available permissions:
        /// </summary>
        [Output("permission")]
        public Output<string> Permission { get; private set; } = null!;

        /// <summary>
        /// Resource affected by this schema registry ACL.
        /// </summary>
        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Username affected by this schema registry ACL.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a KafkaSchemaRegistryAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KafkaSchemaRegistryAcl(string name, KafkaSchemaRegistryAclArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl", name, args ?? new KafkaSchemaRegistryAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KafkaSchemaRegistryAcl(string name, Input<string> id, KafkaSchemaRegistryAclState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/kafkaSchemaRegistryAcl:KafkaSchemaRegistryAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KafkaSchemaRegistryAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KafkaSchemaRegistryAcl Get(string name, Input<string> id, KafkaSchemaRegistryAclState? state = null, CustomResourceOptions? options = null)
        {
            return new KafkaSchemaRegistryAcl(name, id, state, options);
        }
    }

    public sealed class KafkaSchemaRegistryAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Permission to give to this username on this resource. Available permissions:
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// Resource affected by this schema registry ACL.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Username affected by this schema registry ACL.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public KafkaSchemaRegistryAclArgs()
        {
        }
        public static new KafkaSchemaRegistryAclArgs Empty => new KafkaSchemaRegistryAclArgs();
    }

    public sealed class KafkaSchemaRegistryAclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Permission to give to this username on this resource. Available permissions:
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// Resource affected by this schema registry ACL.
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Username affected by this schema registry ACL.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public KafkaSchemaRegistryAclState()
        {
        }
        public static new KafkaSchemaRegistryAclState Empty => new KafkaSchemaRegistryAclState();
    }
}
