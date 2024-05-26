// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// Using a custom template based on an OVHCloud template
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var rescue = Ovh.Dedicated.GetServerBoots.Invoke(new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///         BootType = "rescue",
    ///     });
    /// 
    ///     var debian = new Ovh.Me.InstallationTemplate("debian", new()
    ///     {
    ///         BaseTemplateName = "debian12_64",
    ///         TemplateName = "mydebian12",
    ///         Customization = new Ovh.Me.Inputs.InstallationTemplateCustomizationArgs
    ///         {
    ///             PostInstallationScriptLink = "http://test",
    ///             PostInstallationScriptReturn = "ok",
    ///         },
    ///     });
    /// 
    ///     var serverInstall = new Ovh.Dedicated.ServerInstallTask("serverInstall", new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///         TemplateName = debian.TemplateName,
    ///         BootidOnDestroy = rescue.Apply(getServerBootsResult =&gt; getServerBootsResult.Results[0]),
    ///         Details = new Ovh.Dedicated.Inputs.ServerInstallTaskDetailsArgs
    ///         {
    ///             CustomHostname = "mytest",
    ///         },
    ///         UserMetadatas = new[]
    ///         {
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "sshKey",
    ///                 Value = "ssh-ed25519 AAAAC3...",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Using a BringYourOwnLinux (BYOLinux) template (with userMetadata)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var server = Ovh.GetServer.Invoke(new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///     });
    /// 
    ///     var rescue = Ovh.Dedicated.GetServerBoots.Invoke(new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///         BootType = "rescue",
    ///     });
    /// 
    ///     var serverInstall = new Ovh.Dedicated.ServerInstallTask("serverInstall", new()
    ///     {
    ///         ServiceName = server.Apply(getServerResult =&gt; getServerResult.ServiceName),
    ///         TemplateName = "byolinux_64",
    ///         BootidOnDestroy = rescue.Apply(getServerBootsResult =&gt; getServerBootsResult.Results[0]),
    ///         Details = new Ovh.Dedicated.Inputs.ServerInstallTaskDetailsArgs
    ///         {
    ///             CustomHostname = "mytest",
    ///         },
    ///         UserMetadatas = new[]
    ///         {
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "imageURL",
    ///                 Value = "https://myimage.qcow2",
    ///             },
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "imageType",
    ///                 Value = "qcow2",
    ///             },
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "httpHeaders0Key",
    ///                 Value = "Authorization",
    ///             },
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "httpHeaders0Value",
    ///                 Value = "Basic bG9naW46xxxxxxx=",
    ///             },
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "imageChecksumType",
    ///                 Value = "sha512",
    ///             },
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "imageCheckSum",
    ///                 Value = "047122c9ff4d2a69512212104b06c678f5a9cdb22b75467353613ff87ccd03b57b38967e56d810e61366f9d22d6bd39ac0addf4e00a4c6445112a2416af8f225",
    ///             },
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "configDriveUserData",
    ///                 Value = @$"#cloud-config
    /// ssh_authorized_keys:
    ///   - {data.Ovh_me_ssh_key.Mykey.Key}
    /// 
    /// users:
    ///   - name: patient0
    ///     sudo: ALL=(ALL) NOPASSWD:ALL
    ///     groups: users, sudo
    ///     shell: /bin/bash
    ///     lock_passwd: false
    ///     ssh_authorized_keys:
    ///       - {data.Ovh_me_ssh_key.Mykey.Key}
    /// disable_root: false
    /// packages:
    ///   - vim
    ///   - tree
    /// final_message: The system is finally up, after $UPTIME seconds
    /// ",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Using a Microsoft Windows server OVHcloud template with a specific language
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var server = Ovh.GetServer.Invoke(new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///     });
    /// 
    ///     var rescue = Ovh.Dedicated.GetServerBoots.Invoke(new()
    ///     {
    ///         ServiceName = "nsxxxxxxx.ip-xx-xx-xx.eu",
    ///         BootType = "rescue",
    ///     });
    /// 
    ///     var serverInstall = new Ovh.Dedicated.ServerInstallTask("serverInstall", new()
    ///     {
    ///         ServiceName = server.Apply(getServerResult =&gt; getServerResult.ServiceName),
    ///         TemplateName = "win2019-std_64",
    ///         BootidOnDestroy = rescue.Apply(getServerBootsResult =&gt; getServerBootsResult.Results[0]),
    ///         Details = new Ovh.Dedicated.Inputs.ServerInstallTaskDetailsArgs
    ///         {
    ///             CustomHostname = "mytest",
    ///         },
    ///         UserMetadatas = new[]
    ///         {
    ///             new Ovh.Dedicated.Inputs.ServerInstallTaskUserMetadataArgs
    ///             {
    ///                 Key = "language",
    ///                 Value = "fr-fr",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Installation task can be imported using the `service_name` (`nsXXXX.ip...`) of the baremetal server, the `template_name` used  and ths `task_id`, separated by "/" E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:Dedicated/serverInstallTask:ServerInstallTask ovh_dedicated_server_install_task nsXXXX.ipXXXX/template_name/12345
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Dedicated/serverInstallTask:ServerInstallTask")]
    public partial class ServerInstallTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase.
        /// </summary>
        [Output("bootidOnDestroy")]
        public Output<int?> BootidOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Details of this task. (should be `Install asked`)
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// see `details` block below.
        /// </summary>
        [Output("details")]
        public Output<Outputs.ServerInstallTaskDetails?> Details { get; private set; } = null!;

        /// <summary>
        /// Completion date in RFC3339 format.
        /// </summary>
        [Output("doneDate")]
        public Output<string> DoneDate { get; private set; } = null!;

        /// <summary>
        /// Function name (should be `hardInstall`).
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// Last update in RFC3339 format.
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Output("partitionSchemeName")]
        public Output<string?> PartitionSchemeName { get; private set; } = null!;

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Task creation date in RFC3339 format.
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// Task status (should be `done`)
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Template name.
        /// </summary>
        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;

        /// <summary>
        /// see `user_metadata` block below.
        /// </summary>
        [Output("userMetadatas")]
        public Output<ImmutableArray<Outputs.ServerInstallTaskUserMetadata>> UserMetadatas { get; private set; } = null!;


        /// <summary>
        /// Create a ServerInstallTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerInstallTask(string name, ServerInstallTaskArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/serverInstallTask:ServerInstallTask", name, args ?? new ServerInstallTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerInstallTask(string name, Input<string> id, ServerInstallTaskState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/serverInstallTask:ServerInstallTask", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServerInstallTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerInstallTask Get(string name, Input<string> id, ServerInstallTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new ServerInstallTask(name, id, state, options);
        }
    }

    public sealed class ServerInstallTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase.
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        /// <summary>
        /// see `details` block below.
        /// </summary>
        [Input("details")]
        public Input<Inputs.ServerInstallTaskDetailsArgs>? Details { get; set; }

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Template name.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        [Input("userMetadatas")]
        private InputList<Inputs.ServerInstallTaskUserMetadataArgs>? _userMetadatas;

        /// <summary>
        /// see `user_metadata` block below.
        /// </summary>
        public InputList<Inputs.ServerInstallTaskUserMetadataArgs> UserMetadatas
        {
            get => _userMetadatas ?? (_userMetadatas = new InputList<Inputs.ServerInstallTaskUserMetadataArgs>());
            set => _userMetadatas = value;
        }

        public ServerInstallTaskArgs()
        {
        }
        public static new ServerInstallTaskArgs Empty => new ServerInstallTaskArgs();
    }

    public sealed class ServerInstallTaskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, reboot the server on the specified boot id during destroy phase.
        /// </summary>
        [Input("bootidOnDestroy")]
        public Input<int>? BootidOnDestroy { get; set; }

        /// <summary>
        /// Details of this task. (should be `Install asked`)
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// see `details` block below.
        /// </summary>
        [Input("details")]
        public Input<Inputs.ServerInstallTaskDetailsGetArgs>? Details { get; set; }

        /// <summary>
        /// Completion date in RFC3339 format.
        /// </summary>
        [Input("doneDate")]
        public Input<string>? DoneDate { get; set; }

        /// <summary>
        /// Function name (should be `hardInstall`).
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        /// <summary>
        /// Last update in RFC3339 format.
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// Partition scheme name.
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Task creation date in RFC3339 format.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// Task status (should be `done`)
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Template name.
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        [Input("userMetadatas")]
        private InputList<Inputs.ServerInstallTaskUserMetadataGetArgs>? _userMetadatas;

        /// <summary>
        /// see `user_metadata` block below.
        /// </summary>
        public InputList<Inputs.ServerInstallTaskUserMetadataGetArgs> UserMetadatas
        {
            get => _userMetadatas ?? (_userMetadatas = new InputList<Inputs.ServerInstallTaskUserMetadataGetArgs>());
            set => _userMetadatas = value;
        }

        public ServerInstallTaskState()
        {
        }
        public static new ServerInstallTaskState Empty => new ServerInstallTaskState();
    }
}
