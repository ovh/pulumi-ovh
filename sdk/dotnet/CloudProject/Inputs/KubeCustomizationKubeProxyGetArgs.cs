// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class KubeCustomizationKubeProxyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("iptables")]
        public Input<Inputs.KubeCustomizationKubeProxyIptablesGetArgs>? Iptables { get; set; }

        [Input("ipvs")]
        public Input<Inputs.KubeCustomizationKubeProxyIpvsGetArgs>? Ipvs { get; set; }

        public KubeCustomizationKubeProxyGetArgs()
        {
        }
        public static new KubeCustomizationKubeProxyGetArgs Empty => new KubeCustomizationKubeProxyGetArgs();
    }
}
