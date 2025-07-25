// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    public static class GetAPIOAuth2Client
    {
        /// <summary>
        /// Use this data source to retrieve information about an existing OAuth2 service account.
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
        ///     var myOauth2Client = Ovh.Me.GetAPIOAuth2Client.Invoke(new()
        ///     {
        ///         ClientId = "5f8969a993ec8b4b",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAPIOAuth2ClientResult> InvokeAsync(GetAPIOAuth2ClientArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAPIOAuth2ClientResult>("ovh:Me/getAPIOAuth2Client:getAPIOAuth2Client", args ?? new GetAPIOAuth2ClientArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an existing OAuth2 service account.
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
        ///     var myOauth2Client = Ovh.Me.GetAPIOAuth2Client.Invoke(new()
        ///     {
        ///         ClientId = "5f8969a993ec8b4b",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAPIOAuth2ClientResult> Invoke(GetAPIOAuth2ClientInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAPIOAuth2ClientResult>("ovh:Me/getAPIOAuth2Client:getAPIOAuth2Client", args ?? new GetAPIOAuth2ClientInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an existing OAuth2 service account.
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
        ///     var myOauth2Client = Ovh.Me.GetAPIOAuth2Client.Invoke(new()
        ///     {
        ///         ClientId = "5f8969a993ec8b4b",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAPIOAuth2ClientResult> Invoke(GetAPIOAuth2ClientInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAPIOAuth2ClientResult>("ovh:Me/getAPIOAuth2Client:getAPIOAuth2Client", args ?? new GetAPIOAuth2ClientInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAPIOAuth2ClientArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Client ID of an existing OAuth2 service account.
        /// </summary>
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        public GetAPIOAuth2ClientArgs()
        {
        }
        public static new GetAPIOAuth2ClientArgs Empty => new GetAPIOAuth2ClientArgs();
    }

    public sealed class GetAPIOAuth2ClientInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Client ID of an existing OAuth2 service account.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        public GetAPIOAuth2ClientInvokeArgs()
        {
        }
        public static new GetAPIOAuth2ClientInvokeArgs Empty => new GetAPIOAuth2ClientInvokeArgs();
    }


    [OutputType]
    public sealed class GetAPIOAuth2ClientResult
    {
        /// <summary>
        /// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
        /// </summary>
        public readonly ImmutableArray<string> CallbackUrls;
        /// <summary>
        /// Client ID of the created service account.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// OAuth2 client description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
        /// </summary>
        public readonly string Flow;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Identity;
        /// <summary>
        /// OAuth2 client name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetAPIOAuth2ClientResult(
            ImmutableArray<string> callbackUrls,

            string clientId,

            string description,

            string flow,

            string id,

            string identity,

            string name)
        {
            CallbackUrls = callbackUrls;
            ClientId = clientId;
            Description = description;
            Flow = flow;
            Id = id;
            Identity = identity;
            Name = name;
        }
    }
}
