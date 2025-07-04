// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    /// <summary>
    /// Set the S3 Policy of a public cloud project user.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var user = new Ovh.CloudProject.User("user", new()
    ///     {
    ///         ServiceName = "XXX",
    ///         Description = "my user",
    ///         RoleNames = new[]
    ///         {
    ///             "objectstore_operator",
    ///         },
    ///     });
    /// 
    ///     var myS3Credentials = new Ovh.CloudProject.S3Credential("my_s3_credentials", new()
    ///     {
    ///         ServiceName = user.ServiceName,
    ///         UserId = user.Id,
    ///     });
    /// 
    ///     var policy = new Ovh.CloudProject.S3Policy("policy", new()
    ///     {
    ///         ServiceName = user.ServiceName,
    ///         UserId = user.Id,
    ///         Policy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["Statement"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["Sid"] = "RWContainer",
    ///                     ["Effect"] = "Allow",
    ///                     ["Action"] = new[]
    ///                     {
    ///                         "s3:GetObject",
    ///                         "s3:PutObject",
    ///                         "s3:DeleteObject",
    ///                         "s3:ListBucket",
    ///                         "s3:ListMultipartUploadParts",
    ///                         "s3:ListBucketMultipartUploads",
    ///                         "s3:AbortMultipartUpload",
    ///                         "s3:GetBucketLocation",
    ///                     },
    ///                     ["Resource"] = new[]
    ///                     {
    ///                         "arn:aws:s3:::hp-bucket",
    ///                         "arn:aws:s3:::hp-bucket/*",
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud User S3 Policy can be imported using the `service_name`, `user_id` of the policy, separated by "/" E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:CloudProject/s3Policy:S3Policy policy service_name/user_id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProject/s3Policy:S3Policy")]
    public partial class S3Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// The ID of a public cloud project's user.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a S3Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public S3Policy(string name, S3PolicyArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/s3Policy:S3Policy", name, args ?? new S3PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private S3Policy(string name, Input<string> id, S3PolicyState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/s3Policy:S3Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing S3Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static S3Policy Get(string name, Input<string> id, S3PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new S3Policy(name, id, state, options);
        }
    }

    public sealed class S3PolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The ID of a public cloud project's user.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public S3PolicyArgs()
        {
        }
        public static new S3PolicyArgs Empty => new S3PolicyArgs();
    }

    public sealed class S3PolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// The ID of a public cloud project's user.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public S3PolicyState()
        {
        }
        public static new S3PolicyState Empty => new S3PolicyState();
    }
}
