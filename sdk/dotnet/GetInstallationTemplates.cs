// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh
{
    public static class GetInstallationTemplates
    {
        public static Task<GetInstallationTemplatesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstallationTemplatesResult>("ovh:index/getInstallationTemplates:getInstallationTemplates", InvokeArgs.Empty, options.WithDefaults());

        public static Output<GetInstallationTemplatesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstallationTemplatesResult>("ovh:index/getInstallationTemplates:getInstallationTemplates", InvokeArgs.Empty, options.WithDefaults());

        public static Output<GetInstallationTemplatesResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstallationTemplatesResult>("ovh:index/getInstallationTemplates:getInstallationTemplates", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetInstallationTemplatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Results;

        [OutputConstructor]
        private GetInstallationTemplatesResult(
            string id,

            ImmutableArray<string> results)
        {
            Id = id;
            Results = results;
        }
    }
}
