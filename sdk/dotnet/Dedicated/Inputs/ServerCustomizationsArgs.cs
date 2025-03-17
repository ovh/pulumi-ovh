// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Inputs
{

    public sealed class ServerCustomizationsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Config Drive UserData
        /// </summary>
        [Input("configDriveUserData")]
        public Input<string>? ConfigDriveUserData { get; set; }

        /// <summary>
        /// Path of the EFI bootloader
        /// </summary>
        [Input("efiBootloaderPath")]
        public Input<string>? EfiBootloaderPath { get; set; }

        /// <summary>
        /// Custom hostname
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        [Input("httpHeaders")]
        private InputMap<string>? _httpHeaders;

        /// <summary>
        /// Image HTTP Headers
        /// </summary>
        public InputMap<string> HttpHeaders
        {
            get => _httpHeaders ?? (_httpHeaders = new InputMap<string>());
            set => _httpHeaders = value;
        }

        /// <summary>
        /// Image checksum
        /// </summary>
        [Input("imageCheckSum")]
        public Input<string>? ImageCheckSum { get; set; }

        /// <summary>
        /// Checksum type
        /// </summary>
        [Input("imageCheckSumType")]
        public Input<string>? ImageCheckSumType { get; set; }

        /// <summary>
        /// Image Type
        /// </summary>
        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        /// <summary>
        /// Image URL
        /// </summary>
        [Input("imageUrl")]
        public Input<string>? ImageUrl { get; set; }

        /// <summary>
        /// Display Language
        /// </summary>
        [Input("language")]
        public Input<string>? Language { get; set; }

        /// <summary>
        /// Post-Installation Script
        /// </summary>
        [Input("postInstallationScript")]
        public Input<string>? PostInstallationScript { get; set; }

        /// <summary>
        /// Post-Installation Script File Extension
        /// </summary>
        [Input("postInstallationScriptExtension")]
        public Input<string>? PostInstallationScriptExtension { get; set; }

        /// <summary>
        /// SSH Public Key
        /// </summary>
        [Input("sshKey")]
        public Input<string>? SshKey { get; set; }

        public ServerCustomizationsArgs()
        {
        }
        public static new ServerCustomizationsArgs Empty => new ServerCustomizationsArgs();
    }
}
