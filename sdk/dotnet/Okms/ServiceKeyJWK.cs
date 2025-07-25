// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Okms
{
    /// <summary>
    /// Import an existing key in the JWK format in an OVHcloud KMS.
    /// </summary>
    [OvhResourceType("ovh:Okms/serviceKeyJWK:ServiceKeyJWK")]
    public partial class ServiceKeyJWK : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Context of the key
        /// </summary>
        [Output("context")]
        public Output<string?> Context { get; private set; } = null!;

        /// <summary>
        /// Creation time of the key
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Key deactivation reason
        /// </summary>
        [Output("deactivationReason")]
        public Output<string> DeactivationReason { get; private set; } = null!;

        /// <summary>
        /// IAM resource metadata
        /// </summary>
        [Output("iam")]
        public Output<Outputs.ServiceKeyJWKIam> Iam { get; private set; } = null!;

        /// <summary>
        /// Set of JSON Web Keys to import
        /// </summary>
        [Output("keys")]
        public Output<ImmutableArray<Outputs.ServiceKeyJWKKey>> Keys { get; private set; } = null!;

        /// <summary>
        /// Key name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Okms ID
        /// </summary>
        [Output("okmsId")]
        public Output<string> OkmsId { get; private set; } = null!;

        /// <summary>
        /// Size of the key to be created
        /// </summary>
        [Output("size")]
        public Output<double> Size { get; private set; } = null!;

        /// <summary>
        /// State of the key
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Type of the key to be created
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceKeyJWK resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceKeyJWK(string name, ServiceKeyJWKArgs args, CustomResourceOptions? options = null)
            : base("ovh:Okms/serviceKeyJWK:ServiceKeyJWK", name, args ?? new ServiceKeyJWKArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceKeyJWK(string name, Input<string> id, ServiceKeyJWKState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Okms/serviceKeyJWK:ServiceKeyJWK", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceKeyJWK resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceKeyJWK Get(string name, Input<string> id, ServiceKeyJWKState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceKeyJWK(name, id, state, options);
        }
    }

    public sealed class ServiceKeyJWKArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Context of the key
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        [Input("keys", required: true)]
        private InputList<Inputs.ServiceKeyJWKKeyArgs>? _keys;

        /// <summary>
        /// Set of JSON Web Keys to import
        /// </summary>
        public InputList<Inputs.ServiceKeyJWKKeyArgs> Keys
        {
            get => _keys ?? (_keys = new InputList<Inputs.ServiceKeyJWKKeyArgs>());
            set => _keys = value;
        }

        /// <summary>
        /// Key name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Okms ID
        /// </summary>
        [Input("okmsId", required: true)]
        public Input<string> OkmsId { get; set; } = null!;

        public ServiceKeyJWKArgs()
        {
        }
        public static new ServiceKeyJWKArgs Empty => new ServiceKeyJWKArgs();
    }

    public sealed class ServiceKeyJWKState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Context of the key
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// Creation time of the key
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Key deactivation reason
        /// </summary>
        [Input("deactivationReason")]
        public Input<string>? DeactivationReason { get; set; }

        /// <summary>
        /// IAM resource metadata
        /// </summary>
        [Input("iam")]
        public Input<Inputs.ServiceKeyJWKIamGetArgs>? Iam { get; set; }

        [Input("keys")]
        private InputList<Inputs.ServiceKeyJWKKeyGetArgs>? _keys;

        /// <summary>
        /// Set of JSON Web Keys to import
        /// </summary>
        public InputList<Inputs.ServiceKeyJWKKeyGetArgs> Keys
        {
            get => _keys ?? (_keys = new InputList<Inputs.ServiceKeyJWKKeyGetArgs>());
            set => _keys = value;
        }

        /// <summary>
        /// Key name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Okms ID
        /// </summary>
        [Input("okmsId")]
        public Input<string>? OkmsId { get; set; }

        /// <summary>
        /// Size of the key to be created
        /// </summary>
        [Input("size")]
        public Input<double>? Size { get; set; }

        /// <summary>
        /// State of the key
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Type of the key to be created
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ServiceKeyJWKState()
        {
        }
        public static new ServiceKeyJWKState Empty => new ServiceKeyJWKState();
    }
}
