// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class GetKubeCustomizationKubeProxyIpvsInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        [Input("minSyncPeriod")]
        public Input<string>? MinSyncPeriod { get; set; }

        /// <summary>
        /// IPVS scheduler.
        /// </summary>
        [Input("scheduler")]
        public Input<string>? Scheduler { get; set; }

        /// <summary>
        /// Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.
        /// </summary>
        [Input("syncPeriod")]
        public Input<string>? SyncPeriod { get; set; }

        /// <summary>
        /// Timeout value used for IPVS TCP sessions after receiving a FIN in RFC3339 duration.
        /// </summary>
        [Input("tcpFinTimeout")]
        public Input<string>? TcpFinTimeout { get; set; }

        /// <summary>
        /// Timeout value used for idle IPVS TCP sessions in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        [Input("tcpTimeout")]
        public Input<string>? TcpTimeout { get; set; }

        /// <summary>
        /// timeout value used for IPVS UDP packets in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        [Input("udpTimeout")]
        public Input<string>? UdpTimeout { get; set; }

        public GetKubeCustomizationKubeProxyIpvsInputArgs()
        {
        }
        public static new GetKubeCustomizationKubeProxyIpvsInputArgs Empty => new GetKubeCustomizationKubeProxyIpvsInputArgs();
    }
}
