// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me.Inputs
{

    public sealed class InstallationTemplateCustomizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set up the server using the provided hostname instead of the default hostname.
        /// </summary>
        [Input("customHostname")]
        public Input<string>? CustomHostname { get; set; }

        /// <summary>
        /// Indicate the URL where your postinstall customisation script is located.
        /// </summary>
        [Input("postInstallationScriptLink")]
        public Input<string>? PostInstallationScriptLink { get; set; }

        /// <summary>
        /// indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is 'loh1Xee7eo OK OK OK UGh8Ang1Gu'.
        /// </summary>
        [Input("postInstallationScriptReturn")]
        public Input<string>? PostInstallationScriptReturn { get; set; }

        /// <summary>
        /// Rating.
        /// </summary>
        [Input("rating")]
        public Input<int>? Rating { get; set; }

        /// <summary>
        /// Name of the ssh key that should be installed. Password login will be disabled.
        /// </summary>
        [Input("sshKeyName")]
        public Input<string>? SshKeyName { get; set; }

        public InstallationTemplateCustomizationArgs()
        {
        }
        public static new InstallationTemplateCustomizationArgs Empty => new InstallationTemplateCustomizationArgs();
    }
}
