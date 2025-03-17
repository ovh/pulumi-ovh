// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class StorageReplicationArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.StorageReplicationRuleArgs>? _rules;

        /// <summary>
        /// Replication rules
        /// </summary>
        public InputList<Inputs.StorageReplicationRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.StorageReplicationRuleArgs>());
            set => _rules = value;
        }

        public StorageReplicationArgs()
        {
        }
        public static new StorageReplicationArgs Empty => new StorageReplicationArgs();
    }
}
