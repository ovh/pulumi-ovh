// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.IpLoadBalancing
{
    public static class GetVrackNetwork
    {
        /// <summary>
        /// Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVHcloud account.
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
        ///     var lbNetwork = Ovh.IpLoadBalancing.GetVrackNetwork.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         VrackNetworkId = "yyy",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetVrackNetworkResult> InvokeAsync(GetVrackNetworkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVrackNetworkResult>("ovh:IpLoadBalancing/getVrackNetwork:getVrackNetwork", args ?? new GetVrackNetworkArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVHcloud account.
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
        ///     var lbNetwork = Ovh.IpLoadBalancing.GetVrackNetwork.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         VrackNetworkId = "yyy",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVrackNetworkResult> Invoke(GetVrackNetworkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVrackNetworkResult>("ovh:IpLoadBalancing/getVrackNetwork:getVrackNetwork", args ?? new GetVrackNetworkInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVHcloud account.
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
        ///     var lbNetwork = Ovh.IpLoadBalancing.GetVrackNetwork.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         VrackNetworkId = "yyy",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetVrackNetworkResult> Invoke(GetVrackNetworkInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVrackNetworkResult>("ovh:IpLoadBalancing/getVrackNetwork:getVrackNetwork", args ?? new GetVrackNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVrackNetworkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network
        /// </summary>
        [Input("vrackNetworkId", required: true)]
        public int VrackNetworkId { get; set; }

        public GetVrackNetworkArgs()
        {
        }
        public static new GetVrackNetworkArgs Empty => new GetVrackNetworkArgs();
    }

    public sealed class GetVrackNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Internal Load Balancer identifier of the vRack private network
        /// </summary>
        [Input("vrackNetworkId", required: true)]
        public Input<int> VrackNetworkId { get; set; } = null!;

        public GetVrackNetworkInvokeArgs()
        {
        }
        public static new GetVrackNetworkInvokeArgs Empty => new GetVrackNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVrackNetworkResult
    {
        /// <summary>
        /// Human readable name for your vrack network
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
        /// </summary>
        public readonly string NatIp;
        public readonly string ServiceName;
        /// <summary>
        /// IP block of the private network in the vRack
        /// </summary>
        public readonly string Subnet;
        /// <summary>
        /// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
        /// </summary>
        public readonly int Vlan;
        public readonly int VrackNetworkId;

        [OutputConstructor]
        private GetVrackNetworkResult(
            string displayName,

            string id,

            string natIp,

            string serviceName,

            string subnet,

            int vlan,

            int vrackNetworkId)
        {
            DisplayName = displayName;
            Id = id;
            NatIp = natIp;
            ServiceName = serviceName;
            Subnet = subnet;
            Vlan = vlan;
            VrackNetworkId = vrackNetworkId;
        }
    }
}
