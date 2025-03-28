// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetM3dbUser
    {
        /// <summary>
        /// Use this data source to get information about a user of a M3DB cluster associated with a public cloud project.
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
        ///     var m3dbUser = Ovh.CloudProject.GetM3dbUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         Name = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["m3dbUserGroup"] = m3dbUser.Apply(getM3dbUserResult =&gt; getM3dbUserResult.Group),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetM3dbUserResult> InvokeAsync(GetM3dbUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetM3dbUserResult>("ovh:CloudProject/getM3dbUser:getM3dbUser", args ?? new GetM3dbUserArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a user of a M3DB cluster associated with a public cloud project.
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
        ///     var m3dbUser = Ovh.CloudProject.GetM3dbUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         Name = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["m3dbUserGroup"] = m3dbUser.Apply(getM3dbUserResult =&gt; getM3dbUserResult.Group),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetM3dbUserResult> Invoke(GetM3dbUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetM3dbUserResult>("ovh:CloudProject/getM3dbUser:getM3dbUser", args ?? new GetM3dbUserInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a user of a M3DB cluster associated with a public cloud project.
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
        ///     var m3dbUser = Ovh.CloudProject.GetM3dbUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         Name = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["m3dbUserGroup"] = m3dbUser.Apply(getM3dbUserResult =&gt; getM3dbUserResult.Group),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetM3dbUserResult> Invoke(GetM3dbUserInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetM3dbUserResult>("ovh:CloudProject/getM3dbUser:getM3dbUser", args ?? new GetM3dbUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetM3dbUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the user.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetM3dbUserArgs()
        {
        }
        public static new GetM3dbUserArgs Empty => new GetM3dbUserArgs();
    }

    public sealed class GetM3dbUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the user.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetM3dbUserInvokeArgs()
        {
        }
        public static new GetM3dbUserInvokeArgs Empty => new GetM3dbUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetM3dbUserResult
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
        public readonly string Group;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
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
        private GetM3dbUserResult(
            string clusterId,

            string createdAt,

            string group,

            string id,

            string name,

            string serviceName,

            string status)
        {
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Group = group;
            Id = id;
            Name = name;
            ServiceName = serviceName;
            Status = status;
        }
    }
}
