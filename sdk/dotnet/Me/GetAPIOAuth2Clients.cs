// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    public static class GetAPIOAuth2Clients
    {
        /// <summary>
        /// Use this data source to retrieve information the list of existing OAuth2 service account IDs.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myOauth2Clients = Ovh.Me.GetAPIOAuth2Client.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAPIOAuth2ClientsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAPIOAuth2ClientsResult>("ovh:Me/getAPIOAuth2Clients:getAPIOAuth2Clients", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information the list of existing OAuth2 service account IDs.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myOauth2Clients = Ovh.Me.GetAPIOAuth2Client.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAPIOAuth2ClientsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAPIOAuth2ClientsResult>("ovh:Me/getAPIOAuth2Clients:getAPIOAuth2Clients", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAPIOAuth2ClientsResult
    {
        /// <summary>
        /// The list of all the existing client IDs.
        /// </summary>
        public readonly ImmutableArray<string> ClientIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAPIOAuth2ClientsResult(
            ImmutableArray<string> clientIds,

            string id)
        {
            ClientIds = clientIds;
            Id = id;
        }
    }
}