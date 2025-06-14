// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase.Outputs
{

    [OutputType]
    public sealed class GetCapabilitiesEngineResult
    {
        /// <summary>
        /// Default version used for the engine.
        /// </summary>
        public readonly string DefaultVersion;
        /// <summary>
        /// Description of the plan.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Name of the plan.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// SSL modes for this engine.
        /// </summary>
        public readonly ImmutableArray<string> SslModes;
        /// <summary>
        /// Versions available for this engine.
        /// </summary>
        public readonly ImmutableArray<string> Versions;

        [OutputConstructor]
        private GetCapabilitiesEngineResult(
            string defaultVersion,

            string description,

            string name,

            ImmutableArray<string> sslModes,

            ImmutableArray<string> versions)
        {
            DefaultVersion = defaultVersion;
            Description = description;
            Name = name;
            SslModes = sslModes;
            Versions = versions;
        }
    }
}
