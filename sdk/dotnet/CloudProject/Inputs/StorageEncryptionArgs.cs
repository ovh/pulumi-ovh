// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class StorageEncryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Encryption algorithm
        /// </summary>
        [Input("sseAlgorithm")]
        public Input<string>? SseAlgorithm { get; set; }

        public StorageEncryptionArgs()
        {
        }
        public static new StorageEncryptionArgs Empty => new StorageEncryptionArgs();
    }
}
