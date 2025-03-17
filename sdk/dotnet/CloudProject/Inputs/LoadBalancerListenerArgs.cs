// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class LoadBalancerListenerArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedCidrs")]
        private InputList<string>? _allowedCidrs;

        /// <summary>
        /// The allowed CIDRs
        /// </summary>
        public InputList<string> AllowedCidrs
        {
            get => _allowedCidrs ?? (_allowedCidrs = new InputList<string>());
            set => _allowedCidrs = value;
        }

        /// <summary>
        /// The description of the listener
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the listener
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Listener pool
        /// </summary>
        [Input("pool")]
        public Input<Inputs.LoadBalancerListenerPoolArgs>? Pool { get; set; }

        /// <summary>
        /// Listener port
        /// </summary>
        [Input("port", required: true)]
        public Input<double> Port { get; set; } = null!;

        /// <summary>
        /// Protocol for the listener
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Secret ID to get certificate for SSL listener creation
        /// </summary>
        [Input("secretId")]
        public Input<string>? SecretId { get; set; }

        /// <summary>
        /// Timeout client data of the listener
        /// </summary>
        [Input("timeoutClientData")]
        public Input<double>? TimeoutClientData { get; set; }

        /// <summary>
        /// Timeout member data of the listener
        /// </summary>
        [Input("timeoutMemberData")]
        public Input<double>? TimeoutMemberData { get; set; }

        [Input("tlsVersions")]
        private InputList<string>? _tlsVersions;

        /// <summary>
        /// TLS versions of the listener
        /// </summary>
        public InputList<string> TlsVersions
        {
            get => _tlsVersions ?? (_tlsVersions = new InputList<string>());
            set => _tlsVersions = value;
        }

        public LoadBalancerListenerArgs()
        {
        }
        public static new LoadBalancerListenerArgs Empty => new LoadBalancerListenerArgs();
    }
}
