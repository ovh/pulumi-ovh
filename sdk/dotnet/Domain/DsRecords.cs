// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain
{
    [OvhResourceType("ovh:Domain/dsRecords:DsRecords")]
    public partial class DsRecords : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Domain name
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// DS Records for the domain
        /// </summary>
        [Output("dsRecords")]
        public Output<ImmutableArray<Outputs.DsRecordsDsRecord>> DsRecords { get; private set; } = null!;


        /// <summary>
        /// Create a DsRecords resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DsRecords(string name, DsRecordsArgs args, CustomResourceOptions? options = null)
            : base("ovh:Domain/dsRecords:DsRecords", name, args ?? new DsRecordsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DsRecords(string name, Input<string> id, DsRecordsState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Domain/dsRecords:DsRecords", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DsRecords resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DsRecords Get(string name, Input<string> id, DsRecordsState? state = null, CustomResourceOptions? options = null)
        {
            return new DsRecords(name, id, state, options);
        }
    }

    public sealed class DsRecordsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain name
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("dsRecords", required: true)]
        private InputList<Inputs.DsRecordsDsRecordArgs>? _dsRecords;

        /// <summary>
        /// DS Records for the domain
        /// </summary>
        public InputList<Inputs.DsRecordsDsRecordArgs> DsRecords
        {
            get => _dsRecords ?? (_dsRecords = new InputList<Inputs.DsRecordsDsRecordArgs>());
            set => _dsRecords = value;
        }

        public DsRecordsArgs()
        {
        }
        public static new DsRecordsArgs Empty => new DsRecordsArgs();
    }

    public sealed class DsRecordsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain name
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("dsRecords")]
        private InputList<Inputs.DsRecordsDsRecordGetArgs>? _dsRecords;

        /// <summary>
        /// DS Records for the domain
        /// </summary>
        public InputList<Inputs.DsRecordsDsRecordGetArgs> DsRecords
        {
            get => _dsRecords ?? (_dsRecords = new InputList<Inputs.DsRecordsDsRecordGetArgs>());
            set => _dsRecords = value;
        }

        public DsRecordsState()
        {
        }
        public static new DsRecordsState Empty => new DsRecordsState();
    }
}
