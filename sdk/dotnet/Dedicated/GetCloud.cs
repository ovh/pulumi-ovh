// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    public static class GetCloud
    {
        /// <summary>
        /// Get information about a Managed VMware service. 
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
        ///     var pcc = Ovh.Dedicated.GetCloud.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;Dedicated Cloud service name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCloudResult> InvokeAsync(GetCloudArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudResult>("ovh:Dedicated/getCloud:getCloud", args ?? new GetCloudArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Managed VMware service. 
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
        ///     var pcc = Ovh.Dedicated.GetCloud.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;Dedicated Cloud service name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCloudResult> Invoke(GetCloudInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudResult>("ovh:Dedicated/getCloud:getCloud", args ?? new GetCloudInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Managed VMware service. 
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
        ///     var pcc = Ovh.Dedicated.GetCloud.Invoke(new()
        ///     {
        ///         ServiceName = "&lt;Dedicated Cloud service name&gt;",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCloudResult> Invoke(GetCloudInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudResult>("ovh:Dedicated/getCloud:getCloud", args ?? new GetCloudInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Domain of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudArgs()
        {
        }
        public static new GetCloudArgs Empty => new GetCloudArgs();
    }

    public sealed class GetCloudInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Domain of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudInvokeArgs()
        {
        }
        public static new GetCloudInvokeArgs Empty => new GetCloudInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudResult
    {
        /// <summary>
        /// Advanced security state
        /// </summary>
        public readonly bool AdvancedSecurity;
        /// <summary>
        /// The current bandwidth of your VMware on OVHcloud
        /// </summary>
        public readonly string Bandwidth;
        /// <summary>
        /// Billing type of your VMware on OVHcloud
        /// </summary>
        public readonly string BillingType;
        /// <summary>
        /// Can the PCC be migrated to VCD
        /// </summary>
        public readonly bool CanMigrateToVcd;
        /// <summary>
        /// Url to the VMware on OVHcloud certified interface
        /// </summary>
        public readonly string CertifiedInterfaceUrl;
        /// <summary>
        /// The current version of your VMware on OVHcloud
        /// </summary>
        public readonly string CommercialRange;
        /// <summary>
        /// Description of your VMware on OVHcloud
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Generation of your VMware on OVHcloud
        /// </summary>
        public readonly string Generation;
        /// <summary>
        /// IAM resource metadata
        /// </summary>
        public readonly Outputs.GetCloudIamResult Iam;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Datacenter where your VMware on OVHcloud is physically located
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The management interface name
        /// </summary>
        public readonly string ManagementInterface;
        /// <summary>
        /// The reference universe information for your VMware on OVHcloud
        /// </summary>
        public readonly string ProductReference;
        /// <summary>
        /// Domain of the service
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Name of the current service pack
        /// </summary>
        public readonly string ServicePackName;
        /// <summary>
        /// SPLA licensing state
        /// </summary>
        public readonly bool Spla;
        /// <summary>
        /// Enable SSL v3 support. Warning : this option is not recommended as it was recognized as a security breach. If this is enabled, we advise you to enable the filtered User access policy
        /// </summary>
        public readonly bool SslV3;
        /// <summary>
        /// Current state of your VMware on OVHcloud
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Access policy of your VMware on OVHcloud : opened to every IPs or filtered
        /// </summary>
        public readonly string UserAccessPolicy;
        /// <summary>
        /// The maximum amount of connected users allowed on the VMware on OVHcloud management interface
        /// </summary>
        public readonly double UserLimitConcurrentSession;
        /// <summary>
        /// Which user will be disconnected first in case of quota of maximum connection is reached
        /// </summary>
        public readonly string UserLogoutPolicy;
        /// <summary>
        /// The timeout (in seconds) for the user sessions on the VMware on OVHcloud management interface. 0 value disable the timeout
        /// </summary>
        public readonly double UserSessionTimeout;
        /// <summary>
        /// Url to the VMware on OVHcloud vScope interface
        /// </summary>
        public readonly string VScopeUrl;
        /// <summary>
        /// Version of the management interface
        /// </summary>
        public readonly Outputs.GetCloudVersionResult Version;
        /// <summary>
        /// Url to the VMware on OVHcloud web interface
        /// </summary>
        public readonly string WebInterfaceUrl;

        [OutputConstructor]
        private GetCloudResult(
            bool advancedSecurity,

            string bandwidth,

            string billingType,

            bool canMigrateToVcd,

            string certifiedInterfaceUrl,

            string commercialRange,

            string description,

            string generation,

            Outputs.GetCloudIamResult iam,

            string id,

            string location,

            string managementInterface,

            string productReference,

            string serviceName,

            string servicePackName,

            bool spla,

            bool sslV3,

            string state,

            string userAccessPolicy,

            double userLimitConcurrentSession,

            string userLogoutPolicy,

            double userSessionTimeout,

            string vScopeUrl,

            Outputs.GetCloudVersionResult version,

            string webInterfaceUrl)
        {
            AdvancedSecurity = advancedSecurity;
            Bandwidth = bandwidth;
            BillingType = billingType;
            CanMigrateToVcd = canMigrateToVcd;
            CertifiedInterfaceUrl = certifiedInterfaceUrl;
            CommercialRange = commercialRange;
            Description = description;
            Generation = generation;
            Iam = iam;
            Id = id;
            Location = location;
            ManagementInterface = managementInterface;
            ProductReference = productReference;
            ServiceName = serviceName;
            ServicePackName = servicePackName;
            Spla = spla;
            SslV3 = sslV3;
            State = state;
            UserAccessPolicy = userAccessPolicy;
            UserLimitConcurrentSession = userLimitConcurrentSession;
            UserLogoutPolicy = userLogoutPolicy;
            UserSessionTimeout = userSessionTimeout;
            VScopeUrl = vScopeUrl;
            Version = version;
            WebInterfaceUrl = webInterfaceUrl;
        }
    }
}
