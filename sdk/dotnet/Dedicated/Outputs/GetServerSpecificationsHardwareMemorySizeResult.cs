// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Outputs
{

    [OutputType]
    public sealed class GetServerSpecificationsHardwareMemorySizeResult
    {
        public readonly string Unit;
        public readonly double Value;

        [OutputConstructor]
        private GetServerSpecificationsHardwareMemorySizeResult(
            string unit,

            double value)
        {
            Unit = unit;
            Value = value;
        }
    }
}
