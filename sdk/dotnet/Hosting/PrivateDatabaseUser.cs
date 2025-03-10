// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Hosting
{
    [OvhResourceType("ovh:Hosting/privateDatabaseUser:PrivateDatabaseUser")]
    public partial class PrivateDatabaseUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Password for the new user ( alphanumeric and 8 characters minimum )
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// User name used to connect on your databases
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateDatabaseUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateDatabaseUser(string name, PrivateDatabaseUserArgs args, CustomResourceOptions? options = null)
            : base("ovh:Hosting/privateDatabaseUser:PrivateDatabaseUser", name, args ?? new PrivateDatabaseUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateDatabaseUser(string name, Input<string> id, PrivateDatabaseUserState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Hosting/privateDatabaseUser:PrivateDatabaseUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateDatabaseUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateDatabaseUser Get(string name, Input<string> id, PrivateDatabaseUserState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateDatabaseUser(name, id, state, options);
        }
    }

    public sealed class PrivateDatabaseUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Password for the new user ( alphanumeric and 8 characters minimum )
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
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// User name used to connect on your databases
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public PrivateDatabaseUserArgs()
        {
        }
        public static new PrivateDatabaseUserArgs Empty => new PrivateDatabaseUserArgs();
    }

    public sealed class PrivateDatabaseUserState : global::Pulumi.ResourceArgs
    {
        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the new user ( alphanumeric and 8 characters minimum )
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
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// User name used to connect on your databases
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public PrivateDatabaseUserState()
        {
        }
        public static new PrivateDatabaseUserState Empty => new PrivateDatabaseUserState();
    }
}
