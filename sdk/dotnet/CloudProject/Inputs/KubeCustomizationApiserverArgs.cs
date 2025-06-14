// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class KubeCustomizationApiserverArgs : global::Pulumi.ResourceArgs
    {
        [Input("admissionplugins")]
        private InputList<Inputs.KubeCustomizationApiserverAdmissionpluginArgs>? _admissionplugins;

        /// <summary>
        /// Kubernetes API server admission plugins customization
        /// </summary>
        public InputList<Inputs.KubeCustomizationApiserverAdmissionpluginArgs> Admissionplugins
        {
            get => _admissionplugins ?? (_admissionplugins = new InputList<Inputs.KubeCustomizationApiserverAdmissionpluginArgs>());
            set => _admissionplugins = value;
        }

        public KubeCustomizationApiserverArgs()
        {
        }
        public static new KubeCustomizationApiserverArgs Empty => new KubeCustomizationApiserverArgs();
    }
}
