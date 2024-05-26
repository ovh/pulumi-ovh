// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me.Outputs
{

    [OutputType]
    public sealed class GetInstallationTemplateCustomizationResult
    {
        /// <summary>
        /// Set up the server using the provided hostname instead of the default hostname.
        /// </summary>
        public readonly string CustomHostname;
        /// <summary>
        /// Indicate the URL where your postinstall customisation script is located.
        /// </summary>
        public readonly string PostInstallationScriptLink;
        /// <summary>
        /// Indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is 'loh1Xee7eo OK OK OK UGh8Ang1Gu'.
        /// </summary>
        public readonly string PostInstallationScriptReturn;

        [OutputConstructor]
        private GetInstallationTemplateCustomizationResult(
            string customHostname,

            string postInstallationScriptLink,

            string postInstallationScriptReturn)
        {
            CustomHostname = customHostname;
            PostInstallationScriptLink = postInstallationScriptLink;
            PostInstallationScriptReturn = postInstallationScriptReturn;
        }
    }
}
