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
    public sealed class GetKubeCustomizationResult
    {
        /// <summary>
        /// Kubernetes API server customization
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubeCustomizationApiserverResult> Apiservers;

        [OutputConstructor]
        private GetKubeCustomizationResult(ImmutableArray<Outputs.GetKubeCustomizationApiserverResult> apiservers)
        {
            Apiservers = apiservers;
        }
    }
}
