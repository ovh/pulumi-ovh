// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class GetKubeNodePoolTemplateMetadataResult
    {
        /// <summary>
        /// annotations
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// finalizers
        /// </summary>
        public readonly ImmutableArray<string> Finalizers;
        /// <summary>
        /// labels
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Labels;

        [OutputConstructor]
        private GetKubeNodePoolTemplateMetadataResult(
            ImmutableDictionary<string, string>? annotations,

            ImmutableArray<string> finalizers,

            ImmutableDictionary<string, string>? labels)
        {
            Annotations = annotations;
            Finalizers = finalizers;
            Labels = labels;
        }
    }
}
