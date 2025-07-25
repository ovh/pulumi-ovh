// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.IpLoadBalancing
{
    /// <summary>
    /// Creates a backend server group (farm) to be used by loadbalancing frontend(s)
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
    ///     var lb = Ovh.IpLoadBalancing.GetIpLoadBalancing.Invoke(new()
    ///     {
    ///         ServiceName = "ip-1.2.3.4",
    ///         State = "ok",
    ///     });
    /// 
    ///     var farmName = new Ovh.IpLoadBalancing.UdpFarm("farm_name", new()
    ///     {
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.ServiceName),
    ///         DisplayName = "ingress-8080-gra",
    ///         Zone = "gra",
    ///         Port = 80,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// UDP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by "/" e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:IpLoadBalancing/udpFarm:UdpFarm farmname service_name/farm_id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/udpFarm:UdpFarm")]
    public partial class UdpFarm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Id of your farm.
        /// </summary>
        [Output("farmId")]
        public Output<double> FarmId { get; private set; } = null!;

        /// <summary>
        /// Port attached to your farm ([1..49151]). Inherited from frontend if null
        /// </summary>
        [Output("port")]
        public Output<double> Port { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Output("vrackNetworkId")]
        public Output<double?> VrackNetworkId { get; private set; } = null!;

        /// <summary>
        /// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a UdpFarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UdpFarm(string name, UdpFarmArgs args, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/udpFarm:UdpFarm", name, args ?? new UdpFarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UdpFarm(string name, Input<string> id, UdpFarmState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/udpFarm:UdpFarm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UdpFarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UdpFarm Get(string name, Input<string> id, UdpFarmState? state = null, CustomResourceOptions? options = null)
        {
            return new UdpFarm(name, id, state, options);
        }
    }

    public sealed class UdpFarmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Port attached to your farm ([1..49151]). Inherited from frontend if null
        /// </summary>
        [Input("port", required: true)]
        public Input<double> Port { get; set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Input("vrackNetworkId")]
        public Input<double>? VrackNetworkId { get; set; }

        /// <summary>
        /// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public UdpFarmArgs()
        {
        }
        public static new UdpFarmArgs Empty => new UdpFarmArgs();
    }

    public sealed class UdpFarmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Id of your farm.
        /// </summary>
        [Input("farmId")]
        public Input<double>? FarmId { get; set; }

        /// <summary>
        /// Port attached to your farm ([1..49151]). Inherited from frontend if null
        /// </summary>
        [Input("port")]
        public Input<double>? Port { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Input("vrackNetworkId")]
        public Input<double>? VrackNetworkId { get; set; }

        /// <summary>
        /// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public UdpFarmState()
        {
        }
        public static new UdpFarmState Empty => new UdpFarmState();
    }
}
