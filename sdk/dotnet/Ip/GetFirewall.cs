// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Ip
{
    public static class GetFirewall
    {
        /// <summary>
        /// Use this data source to retrieve information about an IP firewall.
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
        ///     var myFirewall = Ovh.Ip.GetFirewall.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         IpOnFirewall = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFirewallResult> InvokeAsync(GetFirewallArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFirewallResult>("ovh:Ip/getFirewall:getFirewall", args ?? new GetFirewallArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an IP firewall.
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
        ///     var myFirewall = Ovh.Ip.GetFirewall.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         IpOnFirewall = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFirewallResult> Invoke(GetFirewallInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallResult>("ovh:Ip/getFirewall:getFirewall", args ?? new GetFirewallInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an IP firewall.
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
        ///     var myFirewall = Ovh.Ip.GetFirewall.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         IpOnFirewall = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFirewallResult> Invoke(GetFirewallInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFirewallResult>("ovh:Ip/getFirewall:getFirewall", args ?? new GetFirewallInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFirewallArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP or the CIDR
        /// </summary>
        [Input("ip", required: true)]
        public string Ip { get; set; } = null!;

        /// <summary>
        /// IPv4 address
        /// </summary>
        [Input("ipOnFirewall", required: true)]
        public string IpOnFirewall { get; set; } = null!;

        public GetFirewallArgs()
        {
        }
        public static new GetFirewallArgs Empty => new GetFirewallArgs();
    }

    public sealed class GetFirewallInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP or the CIDR
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// IPv4 address
        /// </summary>
        [Input("ipOnFirewall", required: true)]
        public Input<string> IpOnFirewall { get; set; } = null!;

        public GetFirewallInvokeArgs()
        {
        }
        public static new GetFirewallInvokeArgs Empty => new GetFirewallInvokeArgs();
    }


    [OutputType]
    public sealed class GetFirewallResult
    {
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The IP or the CIDR
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// IPv4 address
        /// * `enabled ` - Whether firewall is enabled
        /// </summary>
        public readonly string IpOnFirewall;
        /// <summary>
        /// Current state of your ip on firewall
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetFirewallResult(
            bool enabled,

            string id,

            string ip,

            string ipOnFirewall,

            string state)
        {
            Enabled = enabled;
            Id = id;
            Ip = ip;
            IpOnFirewall = ipOnFirewall;
            State = state;
        }
    }
}
