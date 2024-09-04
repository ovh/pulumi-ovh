// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    /// <summary>
    /// Apply IP restrictions container registry associated with a public cloud project on Harbor UI and API.
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
    ///     var registry = Ovh.CloudProject.GetContainerRegistry.Invoke(new()
    ///     {
    ///         ServiceName = "XXXXXX",
    ///         RegistryId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
    ///     });
    /// 
    ///     var my_mgt_iprestrictions = new Ovh.CloudProject.ContainerRegistryIPRestrictionsManagement("my-mgt-iprestrictions", new()
    ///     {
    ///         ServiceName = ovh_cloud_project_containerregistry.Registry.Service_name,
    ///         RegistryId = ovh_cloud_project_containerregistry.Registry.Id,
    ///         IpRestrictions = new[]
    ///         {
    ///             
    ///             {
    ///                 { "ip_block", "xxx.xxx.xxx.xxx/xx" },
    ///                 { "description", "xxxxxxx" },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProject/containerRegistryIPRestrictionsManagement:ContainerRegistryIPRestrictionsManagement")]
    public partial class ContainerRegistryIPRestrictionsManagement : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP restrictions applied on Harbor UI and API.
        /// </summary>
        [Output("ipRestrictions")]
        public Output<ImmutableArray<ImmutableDictionary<string, string>>> IpRestrictions { get; private set; } = null!;

        /// <summary>
        /// The id of the Managed Private Registry.
        /// </summary>
        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerRegistryIPRestrictionsManagement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerRegistryIPRestrictionsManagement(string name, ContainerRegistryIPRestrictionsManagementArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/containerRegistryIPRestrictionsManagement:ContainerRegistryIPRestrictionsManagement", name, args ?? new ContainerRegistryIPRestrictionsManagementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerRegistryIPRestrictionsManagement(string name, Input<string> id, ContainerRegistryIPRestrictionsManagementState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/containerRegistryIPRestrictionsManagement:ContainerRegistryIPRestrictionsManagement", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ContainerRegistryIPRestrictionsManagement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerRegistryIPRestrictionsManagement Get(string name, Input<string> id, ContainerRegistryIPRestrictionsManagementState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerRegistryIPRestrictionsManagement(name, id, state, options);
        }
    }

    public sealed class ContainerRegistryIPRestrictionsManagementArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipRestrictions", required: true)]
        private InputList<ImmutableDictionary<string, string>>? _ipRestrictions;

        /// <summary>
        /// IP restrictions applied on Harbor UI and API.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> IpRestrictions
        {
            get => _ipRestrictions ?? (_ipRestrictions = new InputList<ImmutableDictionary<string, string>>());
            set => _ipRestrictions = value;
        }

        /// <summary>
        /// The id of the Managed Private Registry.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ContainerRegistryIPRestrictionsManagementArgs()
        {
        }
        public static new ContainerRegistryIPRestrictionsManagementArgs Empty => new ContainerRegistryIPRestrictionsManagementArgs();
    }

    public sealed class ContainerRegistryIPRestrictionsManagementState : global::Pulumi.ResourceArgs
    {
        [Input("ipRestrictions")]
        private InputList<ImmutableDictionary<string, string>>? _ipRestrictions;

        /// <summary>
        /// IP restrictions applied on Harbor UI and API.
        /// </summary>
        public InputList<ImmutableDictionary<string, string>> IpRestrictions
        {
            get => _ipRestrictions ?? (_ipRestrictions = new InputList<ImmutableDictionary<string, string>>());
            set => _ipRestrictions = value;
        }

        /// <summary>
        /// The id of the Managed Private Registry.
        /// </summary>
        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public ContainerRegistryIPRestrictionsManagementState()
        {
        }
        public static new ContainerRegistryIPRestrictionsManagementState Empty => new ContainerRegistryIPRestrictionsManagementState();
    }
}
