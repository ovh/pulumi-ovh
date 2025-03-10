// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dbaas
{
    [OvhResourceType("ovh:Dbaas/logsInput:LogsInput")]
    public partial class LogsInput : global::Pulumi.CustomResource
    {
        /// <summary>
        /// IP blocks
        /// </summary>
        [Output("allowedNetworks")]
        public Output<ImmutableArray<string>> AllowedNetworks { get; private set; } = null!;

        /// <summary>
        /// Whether the workload is auto-scaled
        /// </summary>
        [Output("autoscale")]
        public Output<bool?> Autoscale { get; private set; } = null!;

        /// <summary>
        /// Input configuration
        /// </summary>
        [Output("configuration")]
        public Output<Outputs.LogsInputConfiguration> Configuration { get; private set; } = null!;

        /// <summary>
        /// Input creation
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Number of instance running (returned by the API)
        /// </summary>
        [Output("currentNbInstance")]
        public Output<int> CurrentNbInstance { get; private set; } = null!;

        /// <summary>
        /// Input description
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Input engine ID
        /// </summary>
        [Output("engineId")]
        public Output<string> EngineId { get; private set; } = null!;

        /// <summary>
        /// Port
        /// </summary>
        [Output("exposedPort")]
        public Output<string> ExposedPort { get; private set; } = null!;

        /// <summary>
        /// Hostname
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// Input ID
        /// </summary>
        [Output("inputId")]
        public Output<string> InputId { get; private set; } = null!;

        /// <summary>
        /// Indicate if input need to be restarted
        /// </summary>
        [Output("isRestartRequired")]
        public Output<bool> IsRestartRequired { get; private set; } = null!;

        /// <summary>
        /// Maximum number of instances in auto-scaled mode
        /// </summary>
        [Output("maxScaleInstance")]
        public Output<int?> MaxScaleInstance { get; private set; } = null!;

        /// <summary>
        /// Minimum number of instances in auto-scaled mode
        /// </summary>
        [Output("minScaleInstance")]
        public Output<int?> MinScaleInstance { get; private set; } = null!;

        /// <summary>
        /// Number of instance running
        /// </summary>
        [Output("nbInstance")]
        public Output<int?> NbInstance { get; private set; } = null!;

        /// <summary>
        /// Input IP address
        /// </summary>
        [Output("publicAddress")]
        public Output<string> PublicAddress { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Input SSL certificate
        /// </summary>
        [Output("sslCertificate")]
        public Output<string> SslCertificate { get; private set; } = null!;

        /// <summary>
        /// init: configuration required, pending: ready to start, running: available
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Associated Graylog stream
        /// </summary>
        [Output("streamId")]
        public Output<string> StreamId { get; private set; } = null!;

        /// <summary>
        /// Input title
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;

        /// <summary>
        /// Input last update
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a LogsInput resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogsInput(string name, LogsInputArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dbaas/logsInput:LogsInput", name, args ?? new LogsInputArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogsInput(string name, Input<string> id, LogsInputState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dbaas/logsInput:LogsInput", name, state, MakeResourceOptions(options, id))
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
                    "sslCertificate",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LogsInput resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogsInput Get(string name, Input<string> id, LogsInputState? state = null, CustomResourceOptions? options = null)
        {
            return new LogsInput(name, id, state, options);
        }
    }

    public sealed class LogsInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedNetworks")]
        private InputList<string>? _allowedNetworks;

        /// <summary>
        /// IP blocks
        /// </summary>
        public InputList<string> AllowedNetworks
        {
            get => _allowedNetworks ?? (_allowedNetworks = new InputList<string>());
            set => _allowedNetworks = value;
        }

        /// <summary>
        /// Whether the workload is auto-scaled
        /// </summary>
        [Input("autoscale")]
        public Input<bool>? Autoscale { get; set; }

        /// <summary>
        /// Input configuration
        /// </summary>
        [Input("configuration", required: true)]
        public Input<Inputs.LogsInputConfigurationArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// Input description
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Input engine ID
        /// </summary>
        [Input("engineId", required: true)]
        public Input<string> EngineId { get; set; } = null!;

        /// <summary>
        /// Port
        /// </summary>
        [Input("exposedPort")]
        public Input<string>? ExposedPort { get; set; }

        /// <summary>
        /// Maximum number of instances in auto-scaled mode
        /// </summary>
        [Input("maxScaleInstance")]
        public Input<int>? MaxScaleInstance { get; set; }

        /// <summary>
        /// Minimum number of instances in auto-scaled mode
        /// </summary>
        [Input("minScaleInstance")]
        public Input<int>? MinScaleInstance { get; set; }

        /// <summary>
        /// Number of instance running
        /// </summary>
        [Input("nbInstance")]
        public Input<int>? NbInstance { get; set; }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Associated Graylog stream
        /// </summary>
        [Input("streamId", required: true)]
        public Input<string> StreamId { get; set; } = null!;

        /// <summary>
        /// Input title
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public LogsInputArgs()
        {
        }
        public static new LogsInputArgs Empty => new LogsInputArgs();
    }

    public sealed class LogsInputState : global::Pulumi.ResourceArgs
    {
        [Input("allowedNetworks")]
        private InputList<string>? _allowedNetworks;

        /// <summary>
        /// IP blocks
        /// </summary>
        public InputList<string> AllowedNetworks
        {
            get => _allowedNetworks ?? (_allowedNetworks = new InputList<string>());
            set => _allowedNetworks = value;
        }

        /// <summary>
        /// Whether the workload is auto-scaled
        /// </summary>
        [Input("autoscale")]
        public Input<bool>? Autoscale { get; set; }

        /// <summary>
        /// Input configuration
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.LogsInputConfigurationGetArgs>? Configuration { get; set; }

        /// <summary>
        /// Input creation
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Number of instance running (returned by the API)
        /// </summary>
        [Input("currentNbInstance")]
        public Input<int>? CurrentNbInstance { get; set; }

        /// <summary>
        /// Input description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Input engine ID
        /// </summary>
        [Input("engineId")]
        public Input<string>? EngineId { get; set; }

        /// <summary>
        /// Port
        /// </summary>
        [Input("exposedPort")]
        public Input<string>? ExposedPort { get; set; }

        /// <summary>
        /// Hostname
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// Input ID
        /// </summary>
        [Input("inputId")]
        public Input<string>? InputId { get; set; }

        /// <summary>
        /// Indicate if input need to be restarted
        /// </summary>
        [Input("isRestartRequired")]
        public Input<bool>? IsRestartRequired { get; set; }

        /// <summary>
        /// Maximum number of instances in auto-scaled mode
        /// </summary>
        [Input("maxScaleInstance")]
        public Input<int>? MaxScaleInstance { get; set; }

        /// <summary>
        /// Minimum number of instances in auto-scaled mode
        /// </summary>
        [Input("minScaleInstance")]
        public Input<int>? MinScaleInstance { get; set; }

        /// <summary>
        /// Number of instance running
        /// </summary>
        [Input("nbInstance")]
        public Input<int>? NbInstance { get; set; }

        /// <summary>
        /// Input IP address
        /// </summary>
        [Input("publicAddress")]
        public Input<string>? PublicAddress { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("sslCertificate")]
        private Input<string>? _sslCertificate;

        /// <summary>
        /// Input SSL certificate
        /// </summary>
        public Input<string>? SslCertificate
        {
            get => _sslCertificate;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sslCertificate = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// init: configuration required, pending: ready to start, running: available
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Associated Graylog stream
        /// </summary>
        [Input("streamId")]
        public Input<string>? StreamId { get; set; }

        /// <summary>
        /// Input title
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        /// <summary>
        /// Input last update
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public LogsInputState()
        {
        }
        public static new LogsInputState Empty => new LogsInputState();
    }
}
