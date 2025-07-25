// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetOvhcloudConnectConfigPopDatacenterExtras
    {
        /// <summary>
        /// Use this data source to retrieve information about an Ovhcloud Connect Datacenter Extra configuration
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
        ///     var extraCfgs = Ovh.GetOvhcloudConnectConfigPopDatacenterExtras.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ConfigPopId = "YYY",
        ///         ConfigDatacenterId = "ZZZ",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOvhcloudConnectConfigPopDatacenterExtrasResult> InvokeAsync(GetOvhcloudConnectConfigPopDatacenterExtrasArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOvhcloudConnectConfigPopDatacenterExtrasResult>("ovh:index/getOvhcloudConnectConfigPopDatacenterExtras:getOvhcloudConnectConfigPopDatacenterExtras", args ?? new GetOvhcloudConnectConfigPopDatacenterExtrasArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an Ovhcloud Connect Datacenter Extra configuration
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
        ///     var extraCfgs = Ovh.GetOvhcloudConnectConfigPopDatacenterExtras.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ConfigPopId = "YYY",
        ///         ConfigDatacenterId = "ZZZ",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOvhcloudConnectConfigPopDatacenterExtrasResult> Invoke(GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOvhcloudConnectConfigPopDatacenterExtrasResult>("ovh:index/getOvhcloudConnectConfigPopDatacenterExtras:getOvhcloudConnectConfigPopDatacenterExtras", args ?? new GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an Ovhcloud Connect Datacenter Extra configuration
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
        ///     var extraCfgs = Ovh.GetOvhcloudConnectConfigPopDatacenterExtras.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ConfigPopId = "YYY",
        ///         ConfigDatacenterId = "ZZZ",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOvhcloudConnectConfigPopDatacenterExtrasResult> Invoke(GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOvhcloudConnectConfigPopDatacenterExtrasResult>("ovh:index/getOvhcloudConnectConfigPopDatacenterExtras:getOvhcloudConnectConfigPopDatacenterExtras", args ?? new GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOvhcloudConnectConfigPopDatacenterExtrasArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Datacenter Configuration
        /// </summary>
        [Input("configDatacenterId", required: true)]
        public double ConfigDatacenterId { get; set; }

        /// <summary>
        /// ID of the Pop Configuration
        /// </summary>
        [Input("configPopId", required: true)]
        public double ConfigPopId { get; set; }

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetOvhcloudConnectConfigPopDatacenterExtrasArgs()
        {
        }
        public static new GetOvhcloudConnectConfigPopDatacenterExtrasArgs Empty => new GetOvhcloudConnectConfigPopDatacenterExtrasArgs();
    }

    public sealed class GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the Datacenter Configuration
        /// </summary>
        [Input("configDatacenterId", required: true)]
        public Input<double> ConfigDatacenterId { get; set; } = null!;

        /// <summary>
        /// ID of the Pop Configuration
        /// </summary>
        [Input("configPopId", required: true)]
        public Input<double> ConfigPopId { get; set; } = null!;

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs()
        {
        }
        public static new GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs Empty => new GetOvhcloudConnectConfigPopDatacenterExtrasInvokeArgs();
    }


    [OutputType]
    public sealed class GetOvhcloudConnectConfigPopDatacenterExtrasResult
    {
        /// <summary>
        /// ID of the Datacenter Configuration
        /// </summary>
        public readonly double ConfigDatacenterId;
        /// <summary>
        /// ID of the Pop Configuration
        /// </summary>
        public readonly double ConfigPopId;
        public readonly ImmutableArray<Outputs.GetOvhcloudConnectConfigPopDatacenterExtrasExtraConfigResult> ExtraConfigs;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Service name
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetOvhcloudConnectConfigPopDatacenterExtrasResult(
            double configDatacenterId,

            double configPopId,

            ImmutableArray<Outputs.GetOvhcloudConnectConfigPopDatacenterExtrasExtraConfigResult> extraConfigs,

            string id,

            string serviceName)
        {
            ConfigDatacenterId = configDatacenterId;
            ConfigPopId = configPopId;
            ExtraConfigs = extraConfigs;
            Id = id;
            ServiceName = serviceName;
        }
    }
}
