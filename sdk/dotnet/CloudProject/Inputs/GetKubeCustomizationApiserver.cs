// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class GetKubeCustomizationApiserverArgs : global::Pulumi.InvokeArgs
    {
        [Input("admissionplugins", required: true)]
        private List<Inputs.GetKubeCustomizationApiserverAdmissionpluginArgs>? _admissionplugins;

        /// <summary>
        /// Kubernetes API server admission plugins customization
        /// </summary>
        public List<Inputs.GetKubeCustomizationApiserverAdmissionpluginArgs> Admissionplugins
        {
            get => _admissionplugins ?? (_admissionplugins = new List<Inputs.GetKubeCustomizationApiserverAdmissionpluginArgs>());
            set => _admissionplugins = value;
        }

        public GetKubeCustomizationApiserverArgs()
        {
        }
        public static new GetKubeCustomizationApiserverArgs Empty => new GetKubeCustomizationApiserverArgs();
    }
}
