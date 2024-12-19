// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetFailoverIpAttach
    {
        /// <summary>
        /// Use this data source to get the details of a failover IP address of a service in a public cloud project.
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
        ///     var myFailoverIp = Ovh.CloudProject.GetFailoverIpAttach.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetFailoverIpAttachResult> InvokeAsync(GetFailoverIpAttachArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFailoverIpAttachResult>("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", args ?? new GetFailoverIpAttachArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the details of a failover IP address of a service in a public cloud project.
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
        ///     var myFailoverIp = Ovh.CloudProject.GetFailoverIpAttach.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFailoverIpAttachResult> Invoke(GetFailoverIpAttachInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFailoverIpAttachResult>("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", args ?? new GetFailoverIpAttachInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the details of a failover IP address of a service in a public cloud project.
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
        ///     var myFailoverIp = Ovh.CloudProject.GetFailoverIpAttach.Invoke(new()
        ///     {
        ///         Ip = "XXXXXX",
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetFailoverIpAttachResult> Invoke(GetFailoverIpAttachInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFailoverIpAttachResult>("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", args ?? new GetFailoverIpAttachInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFailoverIpAttachArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP block
        /// * `continentCode` - The Ip continent
        /// </summary>
        [Input("block")]
        public string? Block { get; set; }

        [Input("continentCode")]
        public string? ContinentCode { get; set; }

        [Input("geoLoc")]
        public string? GeoLoc { get; set; }

        /// <summary>
        /// The failover ip address to query
        /// </summary>
        [Input("ip")]
        public string? Ip { get; set; }

        [Input("routedTo")]
        public string? RoutedTo { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetFailoverIpAttachArgs()
        {
        }
        public static new GetFailoverIpAttachArgs Empty => new GetFailoverIpAttachArgs();
    }

    public sealed class GetFailoverIpAttachInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The IP block
        /// * `continentCode` - The Ip continent
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        [Input("continentCode")]
        public Input<string>? ContinentCode { get; set; }

        [Input("geoLoc")]
        public Input<string>? GeoLoc { get; set; }

        /// <summary>
        /// The failover ip address to query
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("routedTo")]
        public Input<string>? RoutedTo { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetFailoverIpAttachInvokeArgs()
        {
        }
        public static new GetFailoverIpAttachInvokeArgs Empty => new GetFailoverIpAttachInvokeArgs();
    }


    [OutputType]
    public sealed class GetFailoverIpAttachResult
    {
        /// <summary>
        /// The IP block
        /// * `continentCode` - The Ip continent
        /// </summary>
        public readonly string Block;
        public readonly string ContinentCode;
        public readonly string GeoLoc;
        /// <summary>
        /// The Ip id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Ip Address
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Current operation progress in percent
        /// * `routedTo` - Instance where ip is routed to
        /// </summary>
        public readonly int Progress;
        public readonly string RoutedTo;
        public readonly string ServiceName;
        /// <summary>
        /// Ip status, can be `ok` or `operationPending`
        /// * `subType` - IP sub type, can be `cloud` or `ovh`
        /// </summary>
        public readonly string Status;
        public readonly string SubType;

        [OutputConstructor]
        private GetFailoverIpAttachResult(
            string block,

            string continentCode,

            string geoLoc,

            string id,

            string ip,

            int progress,

            string routedTo,

            string serviceName,

            string status,

            string subType)
        {
            Block = block;
            ContinentCode = continentCode;
            GeoLoc = geoLoc;
            Id = id;
            Ip = ip;
            Progress = progress;
            RoutedTo = routedTo;
            ServiceName = serviceName;
            Status = status;
            SubType = subType;
        }
    }
}
