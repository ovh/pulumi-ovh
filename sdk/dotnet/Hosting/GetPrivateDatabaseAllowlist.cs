// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Hosting
{
    public static class GetPrivateDatabaseAllowlist
    {
        /// <summary>
        /// Use this data source to retrieve information about an hosting privatedatabase whitelist.
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
        ///     var whitelist = Ovh.Hosting.GetPrivateDatabaseAllowlist.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Ip = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPrivateDatabaseAllowlistResult> InvokeAsync(GetPrivateDatabaseAllowlistArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateDatabaseAllowlistResult>("ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist", args ?? new GetPrivateDatabaseAllowlistArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an hosting privatedatabase whitelist.
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
        ///     var whitelist = Ovh.Hosting.GetPrivateDatabaseAllowlist.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Ip = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPrivateDatabaseAllowlistResult> Invoke(GetPrivateDatabaseAllowlistInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateDatabaseAllowlistResult>("ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist", args ?? new GetPrivateDatabaseAllowlistInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an hosting privatedatabase whitelist.
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
        ///     var whitelist = Ovh.Hosting.GetPrivateDatabaseAllowlist.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Ip = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPrivateDatabaseAllowlistResult> Invoke(GetPrivateDatabaseAllowlistInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateDatabaseAllowlistResult>("ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist", args ?? new GetPrivateDatabaseAllowlistInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateDatabaseAllowlistArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The whitelisted IP in your instance
        /// </summary>
        [Input("ip")]
        public string? Ip { get; set; }

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetPrivateDatabaseAllowlistArgs()
        {
        }
        public static new GetPrivateDatabaseAllowlistArgs Empty => new GetPrivateDatabaseAllowlistArgs();
    }

    public sealed class GetPrivateDatabaseAllowlistInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The whitelisted IP in your instance
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetPrivateDatabaseAllowlistInvokeArgs()
        {
        }
        public static new GetPrivateDatabaseAllowlistInvokeArgs Empty => new GetPrivateDatabaseAllowlistInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateDatabaseAllowlistResult
    {
        /// <summary>
        /// Creation date of the database
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Ip;
        /// <summary>
        /// The last update date of this whitelist
        /// </summary>
        public readonly string LastUpdate;
        /// <summary>
        /// Custom name for your Whitelisted IP
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Authorize this IP to access service port
        /// </summary>
        public readonly bool Service;
        public readonly string ServiceName;
        /// <summary>
        /// Authorize this IP to access SFTP port
        /// </summary>
        public readonly bool Sftp;
        /// <summary>
        /// Whitelist status
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetPrivateDatabaseAllowlistResult(
            string creationDate,

            string id,

            string? ip,

            string lastUpdate,

            string name,

            bool service,

            string serviceName,

            bool sftp,

            string status)
        {
            CreationDate = creationDate;
            Id = id;
            Ip = ip;
            LastUpdate = lastUpdate;
            Name = name;
            Service = service;
            ServiceName = serviceName;
            Sftp = sftp;
            Status = status;
        }
    }
}
