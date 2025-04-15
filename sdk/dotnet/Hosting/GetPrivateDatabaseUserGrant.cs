// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Hosting
{
    public static class GetPrivateDatabaseUserGrant
    {
        /// <summary>
        /// Use this data source to retrieve information about an hosting privatedatabase user grant.
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
        ///     var userGrant = Ovh.Hosting.GetPrivateDatabaseUserGrant.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         DatabaseName = "XXXXXX",
        ///         UserName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPrivateDatabaseUserGrantResult> InvokeAsync(GetPrivateDatabaseUserGrantArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPrivateDatabaseUserGrantResult>("ovh:Hosting/getPrivateDatabaseUserGrant:getPrivateDatabaseUserGrant", args ?? new GetPrivateDatabaseUserGrantArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an hosting privatedatabase user grant.
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
        ///     var userGrant = Ovh.Hosting.GetPrivateDatabaseUserGrant.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         DatabaseName = "XXXXXX",
        ///         UserName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPrivateDatabaseUserGrantResult> Invoke(GetPrivateDatabaseUserGrantInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateDatabaseUserGrantResult>("ovh:Hosting/getPrivateDatabaseUserGrant:getPrivateDatabaseUserGrant", args ?? new GetPrivateDatabaseUserGrantInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an hosting privatedatabase user grant.
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
        ///     var userGrant = Ovh.Hosting.GetPrivateDatabaseUserGrant.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         DatabaseName = "XXXXXX",
        ///         UserName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPrivateDatabaseUserGrantResult> Invoke(GetPrivateDatabaseUserGrantInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPrivateDatabaseUserGrantResult>("ovh:Hosting/getPrivateDatabaseUserGrant:getPrivateDatabaseUserGrant", args ?? new GetPrivateDatabaseUserGrantInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPrivateDatabaseUserGrantArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database name on which grant the user
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// The user name
        /// </summary>
        [Input("userName", required: true)]
        public string UserName { get; set; } = null!;

        public GetPrivateDatabaseUserGrantArgs()
        {
        }
        public static new GetPrivateDatabaseUserGrantArgs Empty => new GetPrivateDatabaseUserGrantArgs();
    }

    public sealed class GetPrivateDatabaseUserGrantInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database name on which grant the user
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// The user name
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public GetPrivateDatabaseUserGrantInvokeArgs()
        {
        }
        public static new GetPrivateDatabaseUserGrantInvokeArgs Empty => new GetPrivateDatabaseUserGrantInvokeArgs();
    }


    [OutputType]
    public sealed class GetPrivateDatabaseUserGrantResult
    {
        /// <summary>
        /// Creation date of the database
        /// </summary>
        public readonly string CreationDate;
        public readonly string DatabaseName;
        /// <summary>
        /// Grant name
        /// </summary>
        public readonly string Grant;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ServiceName;
        public readonly string UserName;

        [OutputConstructor]
        private GetPrivateDatabaseUserGrantResult(
            string creationDate,

            string databaseName,

            string grant,

            string id,

            string serviceName,

            string userName)
        {
            CreationDate = creationDate;
            DatabaseName = databaseName;
            Grant = grant;
            Id = id;
            ServiceName = serviceName;
            UserName = userName;
        }
    }
}
