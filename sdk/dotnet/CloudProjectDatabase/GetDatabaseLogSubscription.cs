// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    public static class GetDatabaseLogSubscription
    {
        /// <summary>
        /// Use this data source to get information about a log subscription for a cluster associated with a public cloud project.
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
        ///     var subscription = Ovh.CloudProjectDatabase.GetDatabaseLogSubscription.Invoke(new()
        ///     {
        ///         ServiceName = "VVV",
        ///         Engine = "XXX",
        ///         ClusterId = "YYY",
        ///         Id = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["subscriptionLdpName"] = subscription.Apply(getDatabaseLogSubscriptionResult =&gt; getDatabaseLogSubscriptionResult.LdpServiceName),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabaseLogSubscriptionResult> InvokeAsync(GetDatabaseLogSubscriptionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseLogSubscriptionResult>("ovh:CloudProjectDatabase/getDatabaseLogSubscription:getDatabaseLogSubscription", args ?? new GetDatabaseLogSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a log subscription for a cluster associated with a public cloud project.
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
        ///     var subscription = Ovh.CloudProjectDatabase.GetDatabaseLogSubscription.Invoke(new()
        ///     {
        ///         ServiceName = "VVV",
        ///         Engine = "XXX",
        ///         ClusterId = "YYY",
        ///         Id = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["subscriptionLdpName"] = subscription.Apply(getDatabaseLogSubscriptionResult =&gt; getDatabaseLogSubscriptionResult.LdpServiceName),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseLogSubscriptionResult> Invoke(GetDatabaseLogSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseLogSubscriptionResult>("ovh:CloudProjectDatabase/getDatabaseLogSubscription:getDatabaseLogSubscription", args ?? new GetDatabaseLogSubscriptionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a log subscription for a cluster associated with a public cloud project.
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
        ///     var subscription = Ovh.CloudProjectDatabase.GetDatabaseLogSubscription.Invoke(new()
        ///     {
        ///         ServiceName = "VVV",
        ///         Engine = "XXX",
        ///         ClusterId = "YYY",
        ///         Id = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["subscriptionLdpName"] = subscription.Apply(getDatabaseLogSubscriptionResult =&gt; getDatabaseLogSubscriptionResult.LdpServiceName),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabaseLogSubscriptionResult> Invoke(GetDatabaseLogSubscriptionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseLogSubscriptionResult>("ovh:CloudProjectDatabase/getDatabaseLogSubscription:getDatabaseLogSubscription", args ?? new GetDatabaseLogSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseLogSubscriptionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The database engine for which you want to retrieve a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public string Engine { get; set; } = null!;

        /// <summary>
        /// Id of the log subscription.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetDatabaseLogSubscriptionArgs()
        {
        }
        public static new GetDatabaseLogSubscriptionArgs Empty => new GetDatabaseLogSubscriptionArgs();
    }

    public sealed class GetDatabaseLogSubscriptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The database engine for which you want to retrieve a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// Id of the log subscription.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetDatabaseLogSubscriptionInvokeArgs()
        {
        }
        public static new GetDatabaseLogSubscriptionInvokeArgs Empty => new GetDatabaseLogSubscriptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseLogSubscriptionResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Creation date of the subscription.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// ID of the log subscription.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Log kind name of this subscription.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the destination log service.
        /// </summary>
        public readonly string LdpServiceName;
        /// <summary>
        /// Name of subscribed resource, where the logs come from.
        /// </summary>
        public readonly string ResourceName;
        /// <summary>
        /// Type of subscribed resource, where the logs come from.
        /// </summary>
        public readonly string ResourceType;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string StreamId;
        /// <summary>
        /// Last update date of the subscription.
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetDatabaseLogSubscriptionResult(
            string clusterId,

            string createdAt,

            string engine,

            string id,

            string kind,

            string ldpServiceName,

            string resourceName,

            string resourceType,

            string serviceName,

            string streamId,

            string updatedAt)
        {
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Engine = engine;
            Id = id;
            Kind = kind;
            LdpServiceName = ldpServiceName;
            ResourceName = resourceName;
            ResourceType = resourceType;
            ServiceName = serviceName;
            StreamId = streamId;
            UpdatedAt = updatedAt;
        }
    }
}
