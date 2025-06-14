// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class LoadBalancerListenerPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pool algorithm to split traffic between members
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// Pool health monitor
        /// </summary>
        [Input("healthMonitor")]
        public Input<Inputs.LoadBalancerListenerPoolHealthMonitorArgs>? HealthMonitor { get; set; }

        [Input("members")]
        private InputList<Inputs.LoadBalancerListenerPoolMemberArgs>? _members;

        /// <summary>
        /// Pool members
        /// </summary>
        public InputList<Inputs.LoadBalancerListenerPoolMemberArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.LoadBalancerListenerPoolMemberArgs>());
            set => _members = value;
        }

        /// <summary>
        /// Name of the pool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol for the pool
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Pool session persistence
        /// </summary>
        [Input("sessionPersistence")]
        public Input<Inputs.LoadBalancerListenerPoolSessionPersistenceArgs>? SessionPersistence { get; set; }

        public LoadBalancerListenerPoolArgs()
        {
        }
        public static new LoadBalancerListenerPoolArgs Empty => new LoadBalancerListenerPoolArgs();
    }
}
