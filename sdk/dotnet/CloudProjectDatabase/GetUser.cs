// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    public static class GetUser
    {
        /// <summary>
        /// Use this data source to get information about a user of a database cluster associated with a public cloud project.
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
        ///     var user = Ovh.CloudProjectDatabase.GetUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         Engine = "YYY",
        ///         ClusterId = "ZZZ",
        ///         Name = "UUU",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["userName"] = user.Apply(getUserResult =&gt; getUserResult.Name),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("ovh:CloudProjectDatabase/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a user of a database cluster associated with a public cloud project.
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
        ///     var user = Ovh.CloudProjectDatabase.GetUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         Engine = "YYY",
        ///         ClusterId = "ZZZ",
        ///         Name = "UUU",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["userName"] = user.Apply(getUserResult =&gt; getUserResult.Name),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("ovh:CloudProjectDatabase/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want user information. To get a full list of available engine visit :
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// Available engines:
        /// </summary>
        [Input("engine", required: true)]
        public string Engine { get; set; } = null!;

        /// <summary>
        /// Name of the user.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want user information. To get a full list of available engine visit :
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// Available engines:
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// Name of the user.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Date of the creation of the user.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Current status of the user.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetUserResult(
            string clusterId,

            string createdAt,

            string engine,

            string id,

            string name,

            string serviceName,

            string status)
        {
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Engine = engine;
            Id = id;
            Name = name;
            ServiceName = serviceName;
            Status = status;
        }
    }
}
