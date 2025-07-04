// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetCloudProjectSshKeys
    {
        /// <summary>
        /// Get SSH keys in a Public Cloud project.
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
        ///     var keys = Ovh.GetCloudProjectSshKeys.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCloudProjectSshKeysResult> InvokeAsync(GetCloudProjectSshKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectSshKeysResult>("ovh:index/getCloudProjectSshKeys:getCloudProjectSshKeys", args ?? new GetCloudProjectSshKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Get SSH keys in a Public Cloud project.
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
        ///     var keys = Ovh.GetCloudProjectSshKeys.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCloudProjectSshKeysResult> Invoke(GetCloudProjectSshKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudProjectSshKeysResult>("ovh:index/getCloudProjectSshKeys:getCloudProjectSshKeys", args ?? new GetCloudProjectSshKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get SSH keys in a Public Cloud project.
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
        ///     var keys = Ovh.GetCloudProjectSshKeys.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCloudProjectSshKeysResult> Invoke(GetCloudProjectSshKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudProjectSshKeysResult>("ovh:index/getCloudProjectSshKeys:getCloudProjectSshKeys", args ?? new GetCloudProjectSshKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectSshKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Region
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectSshKeysArgs()
        {
        }
        public static new GetCloudProjectSshKeysArgs Empty => new GetCloudProjectSshKeysArgs();
    }

    public sealed class GetCloudProjectSshKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Region
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectSshKeysInvokeArgs()
        {
        }
        public static new GetCloudProjectSshKeysInvokeArgs Empty => new GetCloudProjectSshKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectSshKeysResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Region
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Service name
        /// </summary>
        public readonly string ServiceName;
        public readonly ImmutableArray<Outputs.GetCloudProjectSshKeysSshKeyResult> SshKeys;

        [OutputConstructor]
        private GetCloudProjectSshKeysResult(
            string id,

            string region,

            string serviceName,

            ImmutableArray<Outputs.GetCloudProjectSshKeysSshKeyResult> sshKeys)
        {
            Id = id;
            Region = region;
            ServiceName = serviceName;
            SshKeys = sshKeys;
        }
    }
}
