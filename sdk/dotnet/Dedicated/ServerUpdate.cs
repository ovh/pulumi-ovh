// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    /// <summary>
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
    ///     var rescue = Ovh.Dedicated.GetServerBoots.Invoke(new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///         BootType = "rescue",
    ///         Kernel = "rescue64-pro",
    ///     });
    /// 
    ///     var server = new Ovh.Dedicated.ServerUpdate("server", new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///         BootId = rescue.Apply(getServerBootsResult =&gt; getServerBootsResult.Results[0]),
    ///         Monitoring = true,
    ///         State = "ok",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Dedicated/serverUpdate:ServerUpdate")]
    public partial class ServerUpdate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// boot id of the server
        /// </summary>
        [Output("bootId")]
        public Output<int> BootId { get; private set; } = null!;

        /// <summary>
        /// boot script of the server
        /// </summary>
        [Output("bootScript")]
        public Output<string?> BootScript { get; private set; } = null!;

        /// <summary>
        /// Icmp monitoring state
        /// </summary>
        [Output("monitoring")]
        public Output<bool> Monitoring { get; private set; } = null!;

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// error, hacked, hackedBlocked, ok
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a ServerUpdate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerUpdate(string name, ServerUpdateArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/serverUpdate:ServerUpdate", name, args ?? new ServerUpdateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerUpdate(string name, Input<string> id, ServerUpdateState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/serverUpdate:ServerUpdate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerUpdate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerUpdate Get(string name, Input<string> id, ServerUpdateState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerUpdate(name, id, state, options);
        }
    }

    public sealed class ServerUpdateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// boot id of the server
        /// </summary>
        [Input("bootId")]
        public Input<int>? BootId { get; set; }

        /// <summary>
        /// boot script of the server
        /// </summary>
        [Input("bootScript")]
        public Input<string>? BootScript { get; set; }

        /// <summary>
        /// Icmp monitoring state
        /// </summary>
        [Input("monitoring")]
        public Input<bool>? Monitoring { get; set; }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// error, hacked, hackedBlocked, ok
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ServerUpdateArgs()
        {
        }
        public static new ServerUpdateArgs Empty => new ServerUpdateArgs();
    }

    public sealed class ServerUpdateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// boot id of the server
        /// </summary>
        [Input("bootId")]
        public Input<int>? BootId { get; set; }

        /// <summary>
        /// boot script of the server
        /// </summary>
        [Input("bootScript")]
        public Input<string>? BootScript { get; set; }

        /// <summary>
        /// Icmp monitoring state
        /// </summary>
        [Input("monitoring")]
        public Input<bool>? Monitoring { get; set; }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// error, hacked, hackedBlocked, ok
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public ServerUpdateState()
        {
        }
        public static new ServerUpdateState Empty => new ServerUpdateState();
    }
}
