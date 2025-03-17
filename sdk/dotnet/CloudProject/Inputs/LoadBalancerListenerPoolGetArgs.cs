// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class LoadBalancerListenerPoolGetArgs : global::Pulumi.ResourceArgs
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
        public Input<Inputs.LoadBalancerListenerPoolHealthMonitorGetArgs>? HealthMonitor { get; set; }

        [Input("members")]
        private InputList<Inputs.LoadBalancerListenerPoolMemberGetArgs>? _members;

        /// <summary>
        /// Pool members
        /// </summary>
        public InputList<Inputs.LoadBalancerListenerPoolMemberGetArgs> Members
        {
            get => _members ?? (_members = new InputList<Inputs.LoadBalancerListenerPoolMemberGetArgs>());
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
        public Input<Inputs.LoadBalancerListenerPoolSessionPersistenceGetArgs>? SessionPersistence { get; set; }

        public LoadBalancerListenerPoolGetArgs()
        {
        }
        public static new LoadBalancerListenerPoolGetArgs Empty => new LoadBalancerListenerPoolGetArgs();
    }
}
