// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class KubeCustomizationKubeProxyIpvsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("minSyncPeriod")]
        public Input<string>? MinSyncPeriod { get; set; }

        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        [Input("syncPeriod")]
        public Input<string>? SyncPeriod { get; set; }

        [Input("tcpFinTimeout")]
        public Input<string>? TcpFinTimeout { get; set; }

        [Input("tcpTimeout")]
        public Input<string>? TcpTimeout { get; set; }

        [Input("udpTimeout")]
        public Input<string>? UdpTimeout { get; set; }

        public KubeCustomizationKubeProxyIpvsGetArgs()
        {
        }
        public static new KubeCustomizationKubeProxyIpvsGetArgs Empty => new KubeCustomizationKubeProxyIpvsGetArgs();
    }
}
