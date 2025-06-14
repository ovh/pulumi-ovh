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
    ///     var newKms = new Ovh.Okms.Okms("new_kms", new()
    ///     {
    ///         OvhSubsidiary = "FR",
    ///         Region = "eu-west-rbx",
    ///         DisplayName = "terraformed KMS",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Okms/okms:Okms")]
    public partial class Okms : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (String) Resource display name
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// (Attributes) IAM resource metadata (see below for nested schema)
        /// </summary>
        [Output("iam")]
        public Output<Outputs.OkmsIam> Iam { get; private set; } = null!;

        /// <summary>
        /// (String) KMS kmip API endpoint
        /// </summary>
        [Output("kmipEndpoint")]
        public Output<string> KmipEndpoint { get; private set; } = null!;

        /// <summary>
        /// OVH subsidiaries
        /// </summary>
        [Output("ovhSubsidiary")]
        public Output<string> OvhSubsidiary { get; private set; } = null!;

        /// <summary>
        /// (String) KMS public CA (Certificate Authority)
        /// </summary>
        [Output("publicCa")]
        public Output<string> PublicCa { get; private set; } = null!;

        /// <summary>
        /// KMS region
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// (String) KMS rest API endpoint
        /// </summary>
        [Output("restEndpoint")]
        public Output<string> RestEndpoint { get; private set; } = null!;

        /// <summary>
        /// (String) KMS rest API swagger UI
        /// </summary>
        [Output("swaggerEndpoint")]
        public Output<string> SwaggerEndpoint { get; private set; } = null!;


        /// <summary>
        /// Create a Okms resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Okms(string name, OkmsArgs args, CustomResourceOptions? options = null)
            : base("ovh:Okms/okms:Okms", name, args ?? new OkmsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Okms(string name, Input<string> id, OkmsState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Okms/okms:Okms", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Okms resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Okms Get(string name, Input<string> id, OkmsState? state = null, CustomResourceOptions? options = null)
        {
            return new Okms(name, id, state, options);
        }
    }

    public sealed class OkmsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (String) Resource display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// OVH subsidiaries
        /// </summary>
        [Input("ovhSubsidiary", required: true)]
        public Input<string> OvhSubsidiary { get; set; } = null!;

        /// <summary>
        /// KMS region
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public OkmsArgs()
        {
        }
        public static new OkmsArgs Empty => new OkmsArgs();
    }

    public sealed class OkmsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (String) Resource display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// (Attributes) IAM resource metadata (see below for nested schema)
        /// </summary>
        [Input("iam")]
        public Input<Inputs.OkmsIamGetArgs>? Iam { get; set; }

        /// <summary>
        /// (String) KMS kmip API endpoint
        /// </summary>
        [Input("kmipEndpoint")]
        public Input<string>? KmipEndpoint { get; set; }

        /// <summary>
        /// OVH subsidiaries
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// (String) KMS public CA (Certificate Authority)
        /// </summary>
        [Input("publicCa")]
        public Input<string>? PublicCa { get; set; }

        /// <summary>
        /// KMS region
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// (String) KMS rest API endpoint
        /// </summary>
        [Input("restEndpoint")]
        public Input<string>? RestEndpoint { get; set; }

        /// <summary>
        /// (String) KMS rest API swagger UI
        /// </summary>
        [Input("swaggerEndpoint")]
        public Input<string>? SwaggerEndpoint { get; set; }

        public OkmsState()
        {
        }
        public static new OkmsState Empty => new OkmsState();
    }
}
