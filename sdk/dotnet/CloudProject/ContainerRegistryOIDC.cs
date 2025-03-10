// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    [OvhResourceType("ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC")]
    public partial class ContainerRegistryOIDC : global::Pulumi.CustomResource
    {
        [Output("deleteUsers")]
        public Output<bool?> DeleteUsers { get; private set; } = null!;

        [Output("oidcAdminGroup")]
        public Output<string?> OidcAdminGroup { get; private set; } = null!;

        [Output("oidcAutoOnboard")]
        public Output<bool?> OidcAutoOnboard { get; private set; } = null!;

        [Output("oidcClientId")]
        public Output<string> OidcClientId { get; private set; } = null!;

        [Output("oidcClientSecret")]
        public Output<string> OidcClientSecret { get; private set; } = null!;

        [Output("oidcEndpoint")]
        public Output<string> OidcEndpoint { get; private set; } = null!;

        [Output("oidcGroupsClaim")]
        public Output<string?> OidcGroupsClaim { get; private set; } = null!;

        [Output("oidcName")]
        public Output<string> OidcName { get; private set; } = null!;

        [Output("oidcScope")]
        public Output<string> OidcScope { get; private set; } = null!;

        [Output("oidcUserClaim")]
        public Output<string?> OidcUserClaim { get; private set; } = null!;

        [Output("oidcVerifyCert")]
        public Output<bool?> OidcVerifyCert { get; private set; } = null!;

        [Output("registryId")]
        public Output<string> RegistryId { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerRegistryOIDC resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerRegistryOIDC(string name, ContainerRegistryOIDCArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC", name, args ?? new ContainerRegistryOIDCArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerRegistryOIDC(string name, Input<string> id, ContainerRegistryOIDCState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC", name, state, MakeResourceOptions(options, id))
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
                    "oidcClientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ContainerRegistryOIDC resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerRegistryOIDC Get(string name, Input<string> id, ContainerRegistryOIDCState? state = null, CustomResourceOptions? options = null)
        {
            return new ContainerRegistryOIDC(name, id, state, options);
        }
    }

    public sealed class ContainerRegistryOIDCArgs : global::Pulumi.ResourceArgs
    {
        [Input("deleteUsers")]
        public Input<bool>? DeleteUsers { get; set; }

        [Input("oidcAdminGroup")]
        public Input<string>? OidcAdminGroup { get; set; }

        [Input("oidcAutoOnboard")]
        public Input<bool>? OidcAutoOnboard { get; set; }

        [Input("oidcClientId", required: true)]
        public Input<string> OidcClientId { get; set; } = null!;

        [Input("oidcClientSecret", required: true)]
        private Input<string>? _oidcClientSecret;
        public Input<string>? OidcClientSecret
        {
            get => _oidcClientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _oidcClientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("oidcEndpoint", required: true)]
        public Input<string> OidcEndpoint { get; set; } = null!;

        [Input("oidcGroupsClaim")]
        public Input<string>? OidcGroupsClaim { get; set; }

        [Input("oidcName", required: true)]
        public Input<string> OidcName { get; set; } = null!;

        [Input("oidcScope", required: true)]
        public Input<string> OidcScope { get; set; } = null!;

        [Input("oidcUserClaim")]
        public Input<string>? OidcUserClaim { get; set; }

        [Input("oidcVerifyCert")]
        public Input<bool>? OidcVerifyCert { get; set; }

        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ContainerRegistryOIDCArgs()
        {
        }
        public static new ContainerRegistryOIDCArgs Empty => new ContainerRegistryOIDCArgs();
    }

    public sealed class ContainerRegistryOIDCState : global::Pulumi.ResourceArgs
    {
        [Input("deleteUsers")]
        public Input<bool>? DeleteUsers { get; set; }

        [Input("oidcAdminGroup")]
        public Input<string>? OidcAdminGroup { get; set; }

        [Input("oidcAutoOnboard")]
        public Input<bool>? OidcAutoOnboard { get; set; }

        [Input("oidcClientId")]
        public Input<string>? OidcClientId { get; set; }

        [Input("oidcClientSecret")]
        private Input<string>? _oidcClientSecret;
        public Input<string>? OidcClientSecret
        {
            get => _oidcClientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _oidcClientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("oidcEndpoint")]
        public Input<string>? OidcEndpoint { get; set; }

        [Input("oidcGroupsClaim")]
        public Input<string>? OidcGroupsClaim { get; set; }

        [Input("oidcName")]
        public Input<string>? OidcName { get; set; }

        [Input("oidcScope")]
        public Input<string>? OidcScope { get; set; }

        [Input("oidcUserClaim")]
        public Input<string>? OidcUserClaim { get; set; }

        [Input("oidcVerifyCert")]
        public Input<bool>? OidcVerifyCert { get; set; }

        [Input("registryId")]
        public Input<string>? RegistryId { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public ContainerRegistryOIDCState()
        {
        }
        public static new ContainerRegistryOIDCState Empty => new ContainerRegistryOIDCState();
    }
}
