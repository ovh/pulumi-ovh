// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    [OvhResourceType("ovh:Dedicated/serverReinstallTask:ServerReinstallTask")]
    public partial class ServerReinstallTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase
        /// </summary>
        [Output("bootidOnDestroy")]
        public Output<int?> BootidOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Details of this task
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// OS reinstallation customizations
        /// </summary>
        [Output("customizations")]
        public Output<Outputs.ServerReinstallTaskCustomizations?> Customizations { get; private set; } = null!;

        /// <summary>
        /// Completion date
        /// </summary>
        [Output("doneDate")]
        public Output<string> DoneDate { get; private set; } = null!;

        /// <summary>
        /// Function name
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// Last update
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// Operating System name
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// Arbitrary properties to pass to cloud-init's config drive datasource
        /// </summary>
        [Output("properties")]
        public Output<ImmutableDictionary<string, string>?> Properties { get; private set; } = null!;

        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Task Creation date
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// Task status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Storage configuration
        /// </summary>
        [Output("storages")]
        public Output<ImmutableArray<Outputs.ServerReinstallTaskStorage>> Storages { get; private set; } = null!;


        /// <summary>
        /// Create a ServerReinstallTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerReinstallTask(string name, ServerReinstallTaskArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/serverReinstallTask:ServerReinstallTask", name, args ?? new ServerReinstallTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerReinstallTask(string name, Input<string> id, ServerReinstallTaskState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/serverReinstallTask:ServerReinstallTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerReinstallTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerReinstallTask Get(string name, Input<string> id, ServerReinstallTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerReinstallTask(name, id, state, options);
        }
    }

    public sealed class ServerReinstallTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        /// <summary>
        /// OS reinstallation customizations
        /// </summary>
        [Input("customizations")]
        public Input<Inputs.ServerReinstallTaskCustomizationsArgs>? Customizations { get; set; }

        /// <summary>
        /// Operating System name
        /// </summary>
        [Input("os", required: true)]
        public Input<string> Os { get; set; } = null!;

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Arbitrary properties to pass to cloud-init's config drive datasource
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("storages")]
        private InputList<Inputs.ServerReinstallTaskStorageArgs>? _storages;

        /// <summary>
        /// Storage configuration
        /// </summary>
        public InputList<Inputs.ServerReinstallTaskStorageArgs> Storages
        {
            get => _storages ?? (_storages = new InputList<Inputs.ServerReinstallTaskStorageArgs>());
            set => _storages = value;
        }

        public ServerReinstallTaskArgs()
        {
        }
        public static new ServerReinstallTaskArgs Empty => new ServerReinstallTaskArgs();
    }

    public sealed class ServerReinstallTaskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        /// <summary>
        /// Details of this task
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// OS reinstallation customizations
        /// </summary>
        [Input("customizations")]
        public Input<Inputs.ServerReinstallTaskCustomizationsGetArgs>? Customizations { get; set; }

        /// <summary>
        /// Completion date
        /// </summary>
        [Input("doneDate")]
        public Input<string>? DoneDate { get; set; }

        /// <summary>
        /// Function name
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        /// <summary>
        /// Last update
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// Operating System name
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// Arbitrary properties to pass to cloud-init's config drive datasource
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        /// <summary>
        /// The internal name of your dedicated server.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Task Creation date
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// Task status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("storages")]
        private InputList<Inputs.ServerReinstallTaskStorageGetArgs>? _storages;

        /// <summary>
        /// Storage configuration
        /// </summary>
        public InputList<Inputs.ServerReinstallTaskStorageGetArgs> Storages
        {
            get => _storages ?? (_storages = new InputList<Inputs.ServerReinstallTaskStorageGetArgs>());
            set => _storages = value;
        }

        public ServerReinstallTaskState()
        {
        }
        public static new ServerReinstallTaskState Empty => new ServerReinstallTaskState();
    }
}
