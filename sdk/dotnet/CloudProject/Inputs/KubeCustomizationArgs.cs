// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class KubeCustomizationArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiservers")]
        private InputList<Inputs.KubeCustomizationApiserverArgs>? _apiservers;
        [Obsolete(@"Use customization_apiserver instead")]
        public InputList<Inputs.KubeCustomizationApiserverArgs> Apiservers
        {
            get => _apiservers ?? (_apiservers = new InputList<Inputs.KubeCustomizationApiserverArgs>());
            set => _apiservers = value;
        }

        public KubeCustomizationArgs()
        {
        }
        public static new KubeCustomizationArgs Empty => new KubeCustomizationArgs();
    }
}
