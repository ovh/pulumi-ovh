// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Hosting.Outputs
{

    [OutputType]
    public sealed class GetPrivateDatabaseDbUserResult
    {
        /// <summary>
        /// User's rights on this database
        /// </summary>
        public readonly string GrantType;
        /// <summary>
        /// User's name granted on this database
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetPrivateDatabaseDbUserResult(
            string grantType,

            string userName)
        {
            GrantType = grantType;
            UserName = userName;
        }
    }
}
