// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class KubeCustomizationKubeProxyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kubernetes cluster kube-proxy customization of iptables specific config (durations format is RFC3339 duration, e.g. `PT60S`)
        /// </summary>
        [Input("iptables")]
        public Input<Inputs.KubeCustomizationKubeProxyIptablesArgs>? Iptables { get; set; }

        /// <summary>
        /// Kubernetes cluster kube-proxy customization of IPVS specific config (durations format is [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration, e.g. `PT60S`)
        /// </summary>
        [Input("ipvs")]
        public Input<Inputs.KubeCustomizationKubeProxyIpvsArgs>? Ipvs { get; set; }

        public KubeCustomizationKubeProxyArgs()
        {
        }
        public static new KubeCustomizationKubeProxyArgs Empty => new KubeCustomizationKubeProxyArgs();
    }
}
