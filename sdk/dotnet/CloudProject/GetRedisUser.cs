// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetRedisUser
    {
        /// <summary>
        /// Use this data source to get information about a user of a redis cluster associated with a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var redisuser = Ovh.CloudProject.GetRedisUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         Name = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["redisuserCommands"] = redisuser.Apply(getRedisUserResult =&gt; getRedisUserResult.Commands),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRedisUserResult> InvokeAsync(GetRedisUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRedisUserResult>("ovh:CloudProject/getRedisUser:getRedisUser", args ?? new GetRedisUserArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a user of a redis cluster associated with a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var redisuser = Ovh.CloudProject.GetRedisUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         Name = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["redisuserCommands"] = redisuser.Apply(getRedisUserResult =&gt; getRedisUserResult.Commands),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRedisUserResult> Invoke(GetRedisUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRedisUserResult>("ovh:CloudProject/getRedisUser:getRedisUser", args ?? new GetRedisUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRedisUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the user
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetRedisUserArgs()
        {
        }
        public static new GetRedisUserArgs Empty => new GetRedisUserArgs();
    }

    public sealed class GetRedisUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the user
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetRedisUserInvokeArgs()
        {
        }
        public static new GetRedisUserInvokeArgs Empty => new GetRedisUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetRedisUserResult
    {
        /// <summary>
        /// Categories of the user.
        /// </summary>
        public readonly ImmutableArray<string> Categories;
        /// <summary>
        /// Channels of the user.
        /// </summary>
        public readonly ImmutableArray<string> Channels;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Commands of the user.
        /// </summary>
        public readonly ImmutableArray<string> Commands;
        /// <summary>
        /// Date of the creation of the user.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Keys of the user.
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Current status of the user.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Current status of the user.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetRedisUserResult(
            ImmutableArray<string> categories,

            ImmutableArray<string> channels,

            string clusterId,

            ImmutableArray<string> commands,

            string createdAt,

            string id,

            ImmutableArray<string> keys,

            string name,

            string serviceName,

            string status)
        {
            Categories = categories;
            Channels = channels;
            ClusterId = clusterId;
            Commands = commands;
            CreatedAt = createdAt;
            Id = id;
            Keys = keys;
            Name = name;
            ServiceName = serviceName;
            Status = status;
        }
    }
}
