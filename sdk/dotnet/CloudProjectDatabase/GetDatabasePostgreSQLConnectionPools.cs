// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    public static class GetDatabasePostgreSQLConnectionPools
    {
        /// <summary>
        /// Use this data source to get the list of connection pools of a postgresql cluster associated with a public cloud project.
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
        ///     var testPools = Ovh.CloudProjectDatabase.GetDatabasePostgreSQLConnectionPools.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["connectionPoolIds"] = testPools.Apply(getDatabasePostgreSQLConnectionPoolsResult =&gt; getDatabasePostgreSQLConnectionPoolsResult.ConnectionPoolIds),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatabasePostgreSQLConnectionPoolsResult> InvokeAsync(GetDatabasePostgreSQLConnectionPoolsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasePostgreSQLConnectionPoolsResult>("ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools", args ?? new GetDatabasePostgreSQLConnectionPoolsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of connection pools of a postgresql cluster associated with a public cloud project.
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
        ///     var testPools = Ovh.CloudProjectDatabase.GetDatabasePostgreSQLConnectionPools.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["connectionPoolIds"] = testPools.Apply(getDatabasePostgreSQLConnectionPoolsResult =&gt; getDatabasePostgreSQLConnectionPoolsResult.ConnectionPoolIds),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabasePostgreSQLConnectionPoolsResult> Invoke(GetDatabasePostgreSQLConnectionPoolsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasePostgreSQLConnectionPoolsResult>("ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools", args ?? new GetDatabasePostgreSQLConnectionPoolsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of connection pools of a postgresql cluster associated with a public cloud project.
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
        ///     var testPools = Ovh.CloudProjectDatabase.GetDatabasePostgreSQLConnectionPools.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["connectionPoolIds"] = testPools.Apply(getDatabasePostgreSQLConnectionPoolsResult =&gt; getDatabasePostgreSQLConnectionPoolsResult.ConnectionPoolIds),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatabasePostgreSQLConnectionPoolsResult> Invoke(GetDatabasePostgreSQLConnectionPoolsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasePostgreSQLConnectionPoolsResult>("ovh:CloudProjectDatabase/getDatabasePostgreSQLConnectionPools:getDatabasePostgreSQLConnectionPools", args ?? new GetDatabasePostgreSQLConnectionPoolsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasePostgreSQLConnectionPoolsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetDatabasePostgreSQLConnectionPoolsArgs()
        {
        }
        public static new GetDatabasePostgreSQLConnectionPoolsArgs Empty => new GetDatabasePostgreSQLConnectionPoolsArgs();
    }

    public sealed class GetDatabasePostgreSQLConnectionPoolsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetDatabasePostgreSQLConnectionPoolsInvokeArgs()
        {
        }
        public static new GetDatabasePostgreSQLConnectionPoolsInvokeArgs Empty => new GetDatabasePostgreSQLConnectionPoolsInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabasePostgreSQLConnectionPoolsResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The list of patterns ids of the opensearch cluster associated with the project.
        /// </summary>
        public readonly ImmutableArray<string> ConnectionPoolIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetDatabasePostgreSQLConnectionPoolsResult(
            string clusterId,

            ImmutableArray<string> connectionPoolIds,

            string id,

            string serviceName)
        {
            ClusterId = clusterId;
            ConnectionPoolIds = connectionPoolIds;
            Id = id;
            ServiceName = serviceName;
        }
    }
}
