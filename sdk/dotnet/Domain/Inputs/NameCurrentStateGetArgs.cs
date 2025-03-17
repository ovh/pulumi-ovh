// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain.Inputs
{

    public sealed class NameCurrentStateGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalStates")]
        private InputList<string>? _additionalStates;

        /// <summary>
        /// Domain additional states
        /// </summary>
        public InputList<string> AdditionalStates
        {
            get => _additionalStates ?? (_additionalStates = new InputList<string>());
            set => _additionalStates = value;
        }

        /// <summary>
        /// The domain DNS configuration
        /// </summary>
        [Input("dnsConfiguration")]
        public Input<Inputs.NameCurrentStateDnsConfigurationGetArgs>? DnsConfiguration { get; set; }

        /// <summary>
        /// Extension of the domain name
        /// </summary>
        [Input("extension")]
        public Input<string>? Extension { get; set; }

        /// <summary>
        /// Domain main state
        /// </summary>
        [Input("mainState")]
        public Input<string>? MainState { get; set; }

        /// <summary>
        /// Domain name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Domain protection state
        /// </summary>
        [Input("protectionState")]
        public Input<string>? ProtectionState { get; set; }

        /// <summary>
        /// Domain suspension state
        /// </summary>
        [Input("suspensionState")]
        public Input<string>? SuspensionState { get; set; }

        public NameCurrentStateGetArgs()
        {
        }
        public static new NameCurrentStateGetArgs Empty => new NameCurrentStateGetArgs();
    }
}
