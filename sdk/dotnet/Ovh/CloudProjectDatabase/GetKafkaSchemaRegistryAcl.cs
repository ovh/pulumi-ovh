// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Scraly.PulumiPackage.Ovh.CloudProjectDatabase
{
    public static class GetKafkaSchemaRegistryAcl
    {
        /// <summary>
        /// Use this data source to get information about a schema registry ACL of a kafka cluster associated with a public cloud project.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// data "ovh_cloud_project_database_kafka_schemaregistryacl" "schemaRegistryAcl" {
        ///   service_name  = "XXX"
        ///   cluster_id    = "YYY"
        ///   id            = "ZZZ"
        /// }
        /// 
        /// output "acl_permission" {
        ///   value = data.ovh_cloud_project_database_kafka_schemaregistryacl.schemaRegistryAcl.permission
        /// }
        /// ```
        /// </summary>
        public static Task<GetKafkaSchemaRegistryAclResult> InvokeAsync(GetKafkaSchemaRegistryAclArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKafkaSchemaRegistryAclResult>("ovh:CloudProjectDatabase/getKafkaSchemaRegistryAcl:getKafkaSchemaRegistryAcl", args ?? new GetKafkaSchemaRegistryAclArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a schema registry ACL of a kafka cluster associated with a public cloud project.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// data "ovh_cloud_project_database_kafka_schemaregistryacl" "schemaRegistryAcl" {
        ///   service_name  = "XXX"
        ///   cluster_id    = "YYY"
        ///   id            = "ZZZ"
        /// }
        /// 
        /// output "acl_permission" {
        ///   value = data.ovh_cloud_project_database_kafka_schemaregistryacl.schemaRegistryAcl.permission
        /// }
        /// ```
        /// </summary>
        public static Output<GetKafkaSchemaRegistryAclResult> Invoke(GetKafkaSchemaRegistryAclInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKafkaSchemaRegistryAclResult>("ovh:CloudProjectDatabase/getKafkaSchemaRegistryAcl:getKafkaSchemaRegistryAcl", args ?? new GetKafkaSchemaRegistryAclInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKafkaSchemaRegistryAclArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Schema registry ACL ID
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetKafkaSchemaRegistryAclArgs()
        {
        }
        public static new GetKafkaSchemaRegistryAclArgs Empty => new GetKafkaSchemaRegistryAclArgs();
    }

    public sealed class GetKafkaSchemaRegistryAclInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Schema registry ACL ID
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetKafkaSchemaRegistryAclInvokeArgs()
        {
        }
        public static new GetKafkaSchemaRegistryAclInvokeArgs Empty => new GetKafkaSchemaRegistryAclInvokeArgs();
    }


    [OutputType]
    public sealed class GetKafkaSchemaRegistryAclResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Permission to give to this username on this topic.
        /// </summary>
        public readonly string Permission;
        /// <summary>
        /// Resource affected by this ACL.
        /// </summary>
        public readonly string Resource;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Username affected by this ACL.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetKafkaSchemaRegistryAclResult(
            string clusterId,

            string id,

            string permission,

            string resource,

            string serviceName,

            string username)
        {
            ClusterId = clusterId;
            Id = id;
            Permission = permission;
            Resource = resource;
            ServiceName = serviceName;
            Username = username;
        }
    }
}
