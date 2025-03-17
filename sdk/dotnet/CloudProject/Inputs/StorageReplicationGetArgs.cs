// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class StorageReplicationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.StorageReplicationRuleGetArgs>? _rules;

        /// <summary>
        /// Replication rules
        /// </summary>
        public InputList<Inputs.StorageReplicationRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.StorageReplicationRuleGetArgs>());
            set => _rules = value;
        }

        public StorageReplicationGetArgs()
        {
        }
        public static new StorageReplicationGetArgs Empty => new StorageReplicationGetArgs();
    }
}
