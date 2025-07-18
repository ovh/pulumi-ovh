// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dbaas
{
    /// <summary>
    /// Allows to manipulate LDP tokens.
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
    ///     var token = new Ovh.Dbaas.LogsToken("token", new()
    ///     {
    ///         ServiceName = "ldp-xx-xxxxx",
    ///         Name = "ExampleToken",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Dbaas/logsToken:LogsToken")]
    public partial class LogsToken : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID. If not provided, the default cluster_id is used
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Token creation date
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the token
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The LDP service name
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// ID of the token
        /// </summary>
        [Output("tokenId")]
        public Output<string> TokenId { get; private set; } = null!;

        /// <summary>
        /// Token last update date
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// Token value
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a LogsToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogsToken(string name, LogsTokenArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dbaas/logsToken:LogsToken", name, args ?? new LogsTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogsToken(string name, Input<string> id, LogsTokenState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dbaas/logsToken:LogsToken", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
                AdditionalSecretOutputs =
                {
                    "value",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LogsToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogsToken Get(string name, Input<string> id, LogsTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new LogsToken(name, id, state, options);
        }
    }

    public sealed class LogsTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID. If not provided, the default cluster_id is used
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Name of the token
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The LDP service name
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public LogsTokenArgs()
        {
        }
        public static new LogsTokenArgs Empty => new LogsTokenArgs();
    }

    public sealed class LogsTokenState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID. If not provided, the default cluster_id is used
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Token creation date
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Name of the token
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The LDP service name
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// ID of the token
        /// </summary>
        [Input("tokenId")]
        public Input<string>? TokenId { get; set; }

        /// <summary>
        /// Token last update date
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        [Input("value")]
        private Input<string>? _value;

        /// <summary>
        /// Token value
        /// </summary>
        public Input<string>? Value
        {
            get => _value;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _value = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public LogsTokenState()
        {
        }
        public static new LogsTokenState Empty => new LogsTokenState();
    }
}
