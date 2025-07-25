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
    public sealed class GetRancherPlanPlanResult
    {
        /// <summary>
        /// Cause for an unavailability
        /// </summary>
        public readonly string Cause;
        /// <summary>
        /// Human-readable description of the unavailability cause
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// Name of the plan
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Status of the plan
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetRancherPlanPlanResult(
            string cause,

            string message,

            string name,

            string status)
        {
            Cause = cause;
            Message = message;
            Name = name;
            Status = status;
        }
    }
}
