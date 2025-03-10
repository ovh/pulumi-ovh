// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    [OvhResourceType("ovh:CloudProjectDatabase/postgresSqlUser:PostgresSqlUser")]
    public partial class PostgresSqlUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Id of the database cluster
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Date of the creation of the user
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Name of the user
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password of the user
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Arbitrary string to change to trigger a password update
        /// </summary>
        [Output("passwordReset")]
        public Output<string?> PasswordReset { get; private set; } = null!;

        /// <summary>
        /// Roles the user belongs to
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Current status of the user
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a PostgresSqlUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PostgresSqlUser(string name, PostgresSqlUserArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/postgresSqlUser:PostgresSqlUser", name, args ?? new PostgresSqlUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PostgresSqlUser(string name, Input<string> id, PostgresSqlUserState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/postgresSqlUser:PostgresSqlUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PostgresSqlUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PostgresSqlUser Get(string name, Input<string> id, PostgresSqlUserState? state = null, CustomResourceOptions? options = null)
        {
            return new PostgresSqlUser(name, id, state, options);
        }
    }

    public sealed class PostgresSqlUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the database cluster
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the user
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Arbitrary string to change to trigger a password update
        /// </summary>
        [Input("passwordReset")]
        public Input<string>? PasswordReset { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Roles the user belongs to
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public PostgresSqlUserArgs()
        {
        }
        public static new PostgresSqlUserArgs Empty => new PostgresSqlUserArgs();
    }

    public sealed class PostgresSqlUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the database cluster
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Date of the creation of the user
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Name of the user
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password of the user
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
        /// Arbitrary string to change to trigger a password update
        /// </summary>
        [Input("passwordReset")]
        public Input<string>? PasswordReset { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// Roles the user belongs to
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Current status of the user
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PostgresSqlUserState()
        {
        }
        public static new PostgresSqlUserState Empty => new PostgresSqlUserState();
    }
}
