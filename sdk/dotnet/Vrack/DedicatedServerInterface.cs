// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vrack
{
    /// <summary>
    /// Attach a Dedicated Server Network Interface to a vRack.
    /// 
    /// &gt; **NOTE:** The resource `ovh.Vrack.DedicatedServerInterface` is intended to be used for dedicated servers that have configurable network interfaces.&lt;br /&gt; Legacy Dedicated servers that do not have configurable network interfaces MUST use the resource `ovh.Vrack.DedicatedServer` instead.
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
    ///     var server = Ovh.GetServer.Invoke(new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///     });
    /// 
    ///     var vdsi = new Ovh.Vrack.DedicatedServerInterface("vdsi", new()
    ///     {
    ///         ServiceName = "pn-xxxxxxx",
    ///         InterfaceId = server.Apply(getServerResult =&gt; getServerResult.EnabledVrackVnis[0]),
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface")]
    public partial class DedicatedServerInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of dedicated server network interface.
        /// </summary>
        [Output("interfaceId")]
        public Output<string> InterfaceId { get; private set; } = null!;

        /// <summary>
        /// The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedServerInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedServerInterface(string name, DedicatedServerInterfaceArgs args, CustomResourceOptions? options = null)
            : base("ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface", name, args ?? new DedicatedServerInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedServerInterface(string name, Input<string> id, DedicatedServerInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DedicatedServerInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedServerInterface Get(string name, Input<string> id, DedicatedServerInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedServerInterface(name, id, state, options);
        }
    }

    public sealed class DedicatedServerInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of dedicated server network interface.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        /// <summary>
        /// The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public DedicatedServerInterfaceArgs()
        {
        }
        public static new DedicatedServerInterfaceArgs Empty => new DedicatedServerInterfaceArgs();
    }

    public sealed class DedicatedServerInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of dedicated server network interface.
        /// </summary>
        [Input("interfaceId")]
        public Input<string>? InterfaceId { get; set; }

        /// <summary>
        /// The id of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public DedicatedServerInterfaceState()
        {
        }
        public static new DedicatedServerInterfaceState Empty => new DedicatedServerInterfaceState();
    }
}
