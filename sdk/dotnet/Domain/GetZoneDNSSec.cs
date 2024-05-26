// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain
{
    public static class GetZoneDNSSec
    {
        /// <summary>
        /// Use this data source to retrieve information about a domain zone DNSSEC status.
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
        ///     var dnssec = Ovh.Domain.GetZoneDNSSec.Invoke(new()
        ///     {
        ///         ZoneName = "mysite.ovh",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetZoneDNSSecResult> InvokeAsync(GetZoneDNSSecArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZoneDNSSecResult>("ovh:Domain/getZoneDNSSec:getZoneDNSSec", args ?? new GetZoneDNSSecArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a domain zone DNSSEC status.
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
        ///     var dnssec = Ovh.Domain.GetZoneDNSSec.Invoke(new()
        ///     {
        ///         ZoneName = "mysite.ovh",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetZoneDNSSecResult> Invoke(GetZoneDNSSecInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZoneDNSSecResult>("ovh:Domain/getZoneDNSSec:getZoneDNSSec", args ?? new GetZoneDNSSecInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZoneDNSSecArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain zone
        /// </summary>
        [Input("zoneName", required: true)]
        public string ZoneName { get; set; } = null!;

        public GetZoneDNSSecArgs()
        {
        }
        public static new GetZoneDNSSecArgs Empty => new GetZoneDNSSecArgs();
    }

    public sealed class GetZoneDNSSecInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the domain zone
        /// </summary>
        [Input("zoneName", required: true)]
        public Input<string> ZoneName { get; set; } = null!;

        public GetZoneDNSSecInvokeArgs()
        {
        }
        public static new GetZoneDNSSecInvokeArgs Empty => new GetZoneDNSSecInvokeArgs();
    }


    [OutputType]
    public sealed class GetZoneDNSSecResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// DNSSEC status (`disableInProgress`, `disabled`, `enableInProgress` or `enabled`)
        /// </summary>
        public readonly string Status;
        public readonly string ZoneName;

        [OutputConstructor]
        private GetZoneDNSSecResult(
            string id,

            string status,

            string zoneName)
        {
            Id = id;
            Status = status;
            ZoneName = zoneName;
        }
    }
}
