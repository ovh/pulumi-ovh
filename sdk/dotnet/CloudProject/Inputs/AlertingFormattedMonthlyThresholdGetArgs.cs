// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class AlertingFormattedMonthlyThresholdGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("currencyCode")]
        public Input<string>? CurrencyCode { get; set; }

        [Input("text")]
        public Input<string>? Text { get; set; }

        [Input("value")]
        public Input<double>? Value { get; set; }

        public AlertingFormattedMonthlyThresholdGetArgs()
        {
        }
        public static new AlertingFormattedMonthlyThresholdGetArgs Empty => new AlertingFormattedMonthlyThresholdGetArgs();
    }
}
