// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Hosting
{
    [OvhResourceType("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb")]
    public partial class PrivateDatabaseDb : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of your new database
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a PrivateDatabaseDb resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PrivateDatabaseDb(string name, PrivateDatabaseDbArgs args, CustomResourceOptions? options = null)
            : base("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, args ?? new PrivateDatabaseDbArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PrivateDatabaseDb(string name, Input<string> id, PrivateDatabaseDbState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PrivateDatabaseDb resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PrivateDatabaseDb Get(string name, Input<string> id, PrivateDatabaseDbState? state = null, CustomResourceOptions? options = null)
        {
            return new PrivateDatabaseDb(name, id, state, options);
        }
    }

    public sealed class PrivateDatabaseDbArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of your new database
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public PrivateDatabaseDbArgs()
        {
        }
        public static new PrivateDatabaseDbArgs Empty => new PrivateDatabaseDbArgs();
    }

    public sealed class PrivateDatabaseDbState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of your new database
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The internal name of your private database
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public PrivateDatabaseDbState()
        {
        }
        public static new PrivateDatabaseDbState Empty => new PrivateDatabaseDbState();
    }
}
