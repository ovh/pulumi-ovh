// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    ///     var farmName = new Ovh.IpLoadBalancing.TcpFarm("farm_name", new()
    ///     {
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.ServiceName),
    ///         DisplayName = "ingress-8080-gra",
    ///         Zone = "GRA",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// TCP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by "/" e.g.
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:IpLoadBalancing/tcpFarm:TcpFarm farmname service_name/farm_id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/tcpFarm:TcpFarm")]
    public partial class TcpFarm : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        /// </summary>
        [Output("balance")]
        public Output<string?> Balance { get; private set; } = null!;

        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Port attached to your farm ([1..49151]). Inherited from frontend if null
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// define a backend healthcheck probe
        /// </summary>
        [Output("probe")]
        public Output<Outputs.TcpFarmProbe?> Probe { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Stickiness type. No stickiness if null (`sourceIp`)
        /// </summary>
        [Output("stickiness")]
        public Output<string?> Stickiness { get; private set; } = null!;

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Output("vrackNetworkId")]
        public Output<int?> VrackNetworkId { get; private set; } = null!;

        /// <summary>
        /// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a TcpFarm resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TcpFarm(string name, TcpFarmArgs args, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/tcpFarm:TcpFarm", name, args ?? new TcpFarmArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TcpFarm(string name, Input<string> id, TcpFarmState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/tcpFarm:TcpFarm", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TcpFarm resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TcpFarm Get(string name, Input<string> id, TcpFarmState? state = null, CustomResourceOptions? options = null)
        {
            return new TcpFarm(name, id, state, options);
        }
    }

    public sealed class TcpFarmArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        /// </summary>
        [Input("balance")]
        public Input<string>? Balance { get; set; }

        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Port attached to your farm ([1..49151]). Inherited from frontend if null
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// define a backend healthcheck probe
        /// </summary>
        [Input("probe")]
        public Input<Inputs.TcpFarmProbeArgs>? Probe { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Stickiness type. No stickiness if null (`sourceIp`)
        /// </summary>
        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Input("vrackNetworkId")]
        public Input<int>? VrackNetworkId { get; set; }

        /// <summary>
        /// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public TcpFarmArgs()
        {
        }
        public static new TcpFarmArgs Empty => new TcpFarmArgs();
    }

    public sealed class TcpFarmState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
        /// </summary>
        [Input("balance")]
        public Input<string>? Balance { get; set; }

        /// <summary>
        /// Readable label for loadbalancer farm
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Port attached to your farm ([1..49151]). Inherited from frontend if null
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// define a backend healthcheck probe
        /// </summary>
        [Input("probe")]
        public Input<Inputs.TcpFarmProbeGetArgs>? Probe { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Stickiness type. No stickiness if null (`sourceIp`)
        /// </summary>
        [Input("stickiness")]
        public Input<string>? Stickiness { get; set; }

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
        /// </summary>
        [Input("vrackNetworkId")]
        public Input<int>? VrackNetworkId { get; set; }

        /// <summary>
        /// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public TcpFarmState()
        {
        }
        public static new TcpFarmState Empty => new TcpFarmState();
    }
}
