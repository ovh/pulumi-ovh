// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    [OvhResourceType("ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic")]
    public partial class KafkaTopic : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Id of the database cluster
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Minimum insync replica accepted for this topic
        /// </summary>
        [Output("minInsyncReplicas")]
        public Output<int> MinInsyncReplicas { get; private set; } = null!;

        /// <summary>
        /// Name of the topic
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of partitions for this topic
        /// </summary>
        [Output("partitions")]
        public Output<int> Partitions { get; private set; } = null!;

        /// <summary>
        /// Number of replication for this topic
        /// </summary>
        [Output("replication")]
        public Output<int> Replication { get; private set; } = null!;

        /// <summary>
        /// Number of bytes for the retention of the data for this topic
        /// </summary>
        [Output("retentionBytes")]
        public Output<int> RetentionBytes { get; private set; } = null!;

        /// <summary>
        /// Number of hours for the retention of the data for this topic
        /// </summary>
        [Output("retentionHours")]
        public Output<int> RetentionHours { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a KafkaTopic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KafkaTopic(string name, KafkaTopicArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic", name, args ?? new KafkaTopicArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KafkaTopic(string name, Input<string> id, KafkaTopicState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KafkaTopic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KafkaTopic Get(string name, Input<string> id, KafkaTopicState? state = null, CustomResourceOptions? options = null)
        {
            return new KafkaTopic(name, id, state, options);
        }
    }

    public sealed class KafkaTopicArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the database cluster
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Minimum insync replica accepted for this topic
        /// </summary>
        [Input("minInsyncReplicas")]
        public Input<int>? MinInsyncReplicas { get; set; }

        /// <summary>
        /// Name of the topic
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of partitions for this topic
        /// </summary>
        [Input("partitions")]
        public Input<int>? Partitions { get; set; }

        /// <summary>
        /// Number of replication for this topic
        /// </summary>
        [Input("replication")]
        public Input<int>? Replication { get; set; }

        /// <summary>
        /// Number of bytes for the retention of the data for this topic
        /// </summary>
        [Input("retentionBytes")]
        public Input<int>? RetentionBytes { get; set; }

        /// <summary>
        /// Number of hours for the retention of the data for this topic
        /// </summary>
        [Input("retentionHours")]
        public Input<int>? RetentionHours { get; set; }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public KafkaTopicArgs()
        {
        }
        public static new KafkaTopicArgs Empty => new KafkaTopicArgs();
    }

    public sealed class KafkaTopicState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the database cluster
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Minimum insync replica accepted for this topic
        /// </summary>
        [Input("minInsyncReplicas")]
        public Input<int>? MinInsyncReplicas { get; set; }

        /// <summary>
        /// Name of the topic
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of partitions for this topic
        /// </summary>
        [Input("partitions")]
        public Input<int>? Partitions { get; set; }

        /// <summary>
        /// Number of replication for this topic
        /// </summary>
        [Input("replication")]
        public Input<int>? Replication { get; set; }

        /// <summary>
        /// Number of bytes for the retention of the data for this topic
        /// </summary>
        [Input("retentionBytes")]
        public Input<int>? RetentionBytes { get; set; }

        /// <summary>
        /// Number of hours for the retention of the data for this topic
        /// </summary>
        [Input("retentionHours")]
        public Input<int>? RetentionHours { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public KafkaTopicState()
        {
        }
        public static new KafkaTopicState Empty => new KafkaTopicState();
    }
}
