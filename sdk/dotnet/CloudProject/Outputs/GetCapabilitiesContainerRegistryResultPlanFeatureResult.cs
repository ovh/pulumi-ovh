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
    public sealed class GetCapabilitiesContainerRegistryResultPlanFeatureResult
    {
        /// <summary>
        /// Vulnerability scanning
        /// </summary>
        public readonly bool Vulnerability;

        [OutputConstructor]
        private GetCapabilitiesContainerRegistryResultPlanFeatureResult(bool vulnerability)
        {
            Vulnerability = vulnerability;
        }
    }
}
