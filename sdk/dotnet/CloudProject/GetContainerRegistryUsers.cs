// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetContainerRegistryUsers
    {
        /// <summary>
        /// Use this data source to get the list of users of a container registry associated with a public cloud project.
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
        ///     var myRegistry = Ovh.CloudProject.GetContainerRegistry.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         RegistryId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     var users = Ovh.CloudProject.GetContainerRegistryUsers.Invoke(new()
        ///     {
        ///         ServiceName = ovh_cloud_project_containerregistry.My_registry.Service_name,
        ///         RegistryId = ovh_cloud_project_containerregistry.My_registry.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetContainerRegistryUsersResult> InvokeAsync(GetContainerRegistryUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContainerRegistryUsersResult>("ovh:CloudProject/getContainerRegistryUsers:getContainerRegistryUsers", args ?? new GetContainerRegistryUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of users of a container registry associated with a public cloud project.
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
        ///     var myRegistry = Ovh.CloudProject.GetContainerRegistry.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         RegistryId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     var users = Ovh.CloudProject.GetContainerRegistryUsers.Invoke(new()
        ///     {
        ///         ServiceName = ovh_cloud_project_containerregistry.My_registry.Service_name,
        ///         RegistryId = ovh_cloud_project_containerregistry.My_registry.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetContainerRegistryUsersResult> Invoke(GetContainerRegistryUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistryUsersResult>("ovh:CloudProject/getContainerRegistryUsers:getContainerRegistryUsers", args ?? new GetContainerRegistryUsersInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of users of a container registry associated with a public cloud project.
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
        ///     var myRegistry = Ovh.CloudProject.GetContainerRegistry.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         RegistryId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     var users = Ovh.CloudProject.GetContainerRegistryUsers.Invoke(new()
        ///     {
        ///         ServiceName = ovh_cloud_project_containerregistry.My_registry.Service_name,
        ///         RegistryId = ovh_cloud_project_containerregistry.My_registry.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetContainerRegistryUsersResult> Invoke(GetContainerRegistryUsersInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistryUsersResult>("ovh:CloudProject/getContainerRegistryUsers:getContainerRegistryUsers", args ?? new GetContainerRegistryUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerRegistryUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Registry ID
        /// </summary>
        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetContainerRegistryUsersArgs()
        {
        }
        public static new GetContainerRegistryUsersArgs Empty => new GetContainerRegistryUsersArgs();
    }

    public sealed class GetContainerRegistryUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Registry ID
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetContainerRegistryUsersInvokeArgs()
        {
        }
        public static new GetContainerRegistryUsersInvokeArgs Empty => new GetContainerRegistryUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetContainerRegistryUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string RegistryId;
        /// <summary>
        /// The list of users of the container registry associated with the project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContainerRegistryUsersResultResult> Results;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetContainerRegistryUsersResult(
            string id,

            string registryId,

            ImmutableArray<Outputs.GetContainerRegistryUsersResultResult> results,

            string serviceName)
        {
            Id = id;
            RegistryId = registryId;
            Results = results;
            ServiceName = serviceName;
        }
    }
}
