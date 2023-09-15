// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    /// <summary>
    /// ## Import
    /// 
    /// OVHcloud Managed OpenSearch clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g., bash
    /// 
    /// ```sh
    ///  $ pulumi import ovh:CloudProjectDatabase/opensearchUser:OpensearchUser my_user service_name/cluster_id/id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProjectDatabase/opensearchUser:OpensearchUser")]
    public partial class OpensearchUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Acls of the user.
        /// </summary>
        [Output("acls")]
        public Output<ImmutableArray<Outputs.OpensearchUserAcl>> Acls { get; private set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Date of the creation of the user.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Username affected by this acl. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Sensitive) Password of the user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        [Output("passwordReset")]
        public Output<string?> PasswordReset { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Current status of the user.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a OpensearchUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OpensearchUser(string name, OpensearchUserArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/opensearchUser:OpensearchUser", name, args ?? new OpensearchUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OpensearchUser(string name, Input<string> id, OpensearchUserState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/opensearchUser:OpensearchUser", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OpensearchUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OpensearchUser Get(string name, Input<string> id, OpensearchUserState? state = null, CustomResourceOptions? options = null)
        {
            return new OpensearchUser(name, id, state, options);
        }
    }

    public sealed class OpensearchUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.OpensearchUserAclArgs>? _acls;

        /// <summary>
        /// Acls of the user.
        /// </summary>
        public InputList<Inputs.OpensearchUserAclArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.OpensearchUserAclArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Username affected by this acl. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        [Input("passwordReset")]
        public Input<string>? PasswordReset { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public OpensearchUserArgs()
        {
        }
        public static new OpensearchUserArgs Empty => new OpensearchUserArgs();
    }

    public sealed class OpensearchUserState : global::Pulumi.ResourceArgs
    {
        [Input("acls")]
        private InputList<Inputs.OpensearchUserAclGetArgs>? _acls;

        /// <summary>
        /// Acls of the user.
        /// </summary>
        public InputList<Inputs.OpensearchUserAclGetArgs> Acls
        {
            get => _acls ?? (_acls = new InputList<Inputs.OpensearchUserAclGetArgs>());
            set => _acls = value;
        }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Date of the creation of the user.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Username affected by this acl. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// (Sensitive) Password of the user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        [Input("passwordReset")]
        public Input<string>? PasswordReset { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Current status of the user.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public OpensearchUserState()
        {
        }
        public static new OpensearchUserState Empty => new OpensearchUserState();
    }
}
