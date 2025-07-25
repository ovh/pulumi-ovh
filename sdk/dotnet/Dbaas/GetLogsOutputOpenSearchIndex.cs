// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dbaas
{
    public static class GetLogsOutputOpenSearchIndex
    {
        /// <summary>
        /// Use this data source to retrieve information about a DBaas logs output opensearch index.
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
        ///     var index = Ovh.Dbaas.GetLogsOutputOpenSearchIndex.Invoke(new()
        ///     {
        ///         ServiceName = "ldp-xx-xxxxx",
        ///         Name = "index-name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLogsOutputOpenSearchIndexResult> InvokeAsync(GetLogsOutputOpenSearchIndexArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogsOutputOpenSearchIndexResult>("ovh:Dbaas/getLogsOutputOpenSearchIndex:getLogsOutputOpenSearchIndex", args ?? new GetLogsOutputOpenSearchIndexArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a DBaas logs output opensearch index.
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
        ///     var index = Ovh.Dbaas.GetLogsOutputOpenSearchIndex.Invoke(new()
        ///     {
        ///         ServiceName = "ldp-xx-xxxxx",
        ///         Name = "index-name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLogsOutputOpenSearchIndexResult> Invoke(GetLogsOutputOpenSearchIndexInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogsOutputOpenSearchIndexResult>("ovh:Dbaas/getLogsOutputOpenSearchIndex:getLogsOutputOpenSearchIndex", args ?? new GetLogsOutputOpenSearchIndexInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a DBaas logs output opensearch index.
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
        ///     var index = Ovh.Dbaas.GetLogsOutputOpenSearchIndex.Invoke(new()
        ///     {
        ///         ServiceName = "ldp-xx-xxxxx",
        ///         Name = "index-name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLogsOutputOpenSearchIndexResult> Invoke(GetLogsOutputOpenSearchIndexInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogsOutputOpenSearchIndexResult>("ovh:Dbaas/getLogsOutputOpenSearchIndex:getLogsOutputOpenSearchIndex", args ?? new GetLogsOutputOpenSearchIndexInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogsOutputOpenSearchIndexArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Index name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Number of shard
        /// </summary>
        [Input("nbShard")]
        public int? NbShard { get; set; }

        /// <summary>
        /// The service name. It's the ID of your Logs Data Platform instance.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetLogsOutputOpenSearchIndexArgs()
        {
        }
        public static new GetLogsOutputOpenSearchIndexArgs Empty => new GetLogsOutputOpenSearchIndexArgs();
    }

    public sealed class GetLogsOutputOpenSearchIndexInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Index name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Number of shard
        /// </summary>
        [Input("nbShard")]
        public Input<int>? NbShard { get; set; }

        /// <summary>
        /// The service name. It's the ID of your Logs Data Platform instance.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetLogsOutputOpenSearchIndexInvokeArgs()
        {
        }
        public static new GetLogsOutputOpenSearchIndexInvokeArgs Empty => new GetLogsOutputOpenSearchIndexInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogsOutputOpenSearchIndexResult
    {
        /// <summary>
        /// If set, notify when size is near 80, 90 or 100 % of its maximum capacity
        /// </summary>
        public readonly bool AlertNotifyEnabled;
        /// <summary>
        /// Index creation
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Current index size (in bytes)
        /// </summary>
        public readonly int CurrentSize;
        /// <summary>
        /// Index description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Index ID
        /// </summary>
        public readonly string IndexId;
        /// <summary>
        /// Indicates if you are allowed to edit entry
        /// </summary>
        public readonly bool IsEditable;
        /// <summary>
        /// Maximum index size (in bytes)
        /// </summary>
        public readonly int MaxSize;
        /// <summary>
        /// Index name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of shard
        /// </summary>
        public readonly int NbShard;
        public readonly string ServiceName;
        /// <summary>
        /// Index last update
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetLogsOutputOpenSearchIndexResult(
            bool alertNotifyEnabled,

            string createdAt,

            int currentSize,

            string description,

            string id,

            string indexId,

            bool isEditable,

            int maxSize,

            string name,

            int nbShard,

            string serviceName,

            string updatedAt)
        {
            AlertNotifyEnabled = alertNotifyEnabled;
            CreatedAt = createdAt;
            CurrentSize = currentSize;
            Description = description;
            Id = id;
            IndexId = indexId;
            IsEditable = isEditable;
            MaxSize = maxSize;
            Name = name;
            NbShard = nbShard;
            ServiceName = serviceName;
            UpdatedAt = updatedAt;
        }
    }
}
