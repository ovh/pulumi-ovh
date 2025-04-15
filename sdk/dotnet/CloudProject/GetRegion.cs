// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetRegion
    {
        /// <summary>
        /// Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.
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
        ///     var GRA1 = Ovh.CloudProject.GetRegion.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Name = "GRA1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRegionResult> InvokeAsync(GetRegionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegionResult>("ovh:CloudProject/getRegion:getRegion", args ?? new GetRegionArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.
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
        ///     var GRA1 = Ovh.CloudProject.GetRegion.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Name = "GRA1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRegionResult> Invoke(GetRegionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionResult>("ovh:CloudProject/getRegion:getRegion", args ?? new GetRegionInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.
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
        ///     var GRA1 = Ovh.CloudProject.GetRegion.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Name = "GRA1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRegionResult> Invoke(GetRegionInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionResult>("ovh:CloudProject/getRegion:getRegion", args ?? new GetRegionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the region associated with the public cloud project.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetRegionArgs()
        {
        }
        public static new GetRegionArgs Empty => new GetRegionArgs();
    }

    public sealed class GetRegionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the region associated with the public cloud project.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetRegionInvokeArgs()
        {
        }
        public static new GetRegionInvokeArgs Empty => new GetRegionInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionResult
    {
        /// <summary>
        /// the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...
        /// </summary>
        public readonly string ContinentCode;
        /// <summary>
        /// The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"
        /// </summary>
        public readonly string DatacenterLocation;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// the name of the public cloud service
        /// </summary>
        public readonly string Name;
        public readonly string ServiceName;
        /// <summary>
        /// The list of public cloud services running within the region
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRegionServiceResult> Services;

        [OutputConstructor]
        private GetRegionResult(
            string continentCode,

            string datacenterLocation,

            string id,

            string name,

            string serviceName,

            ImmutableArray<Outputs.GetRegionServiceResult> services)
        {
            ContinentCode = continentCode;
            DatacenterLocation = datacenterLocation;
            Id = id;
            Name = name;
            ServiceName = serviceName;
            Services = services;
        }
    }
}
