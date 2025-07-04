// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dbaas.Outputs
{

    [OutputType]
    public sealed class LogsInputConfigurationLogstash
    {
        /// <summary>
        /// The filter section of logstash.conf
        /// </summary>
        public readonly string? FilterSection;
        /// <summary>
        /// The filter section of logstash.conf
        /// </summary>
        public readonly string InputSection;
        /// <summary>
        /// The list of customs Grok patterns
        /// </summary>
        public readonly string? PatternSection;

        [OutputConstructor]
        private LogsInputConfigurationLogstash(
            string? filterSection,

            string inputSection,

            string? patternSection)
        {
            FilterSection = filterSection;
            InputSection = inputSection;
            PatternSection = patternSection;
        }
    }
}
