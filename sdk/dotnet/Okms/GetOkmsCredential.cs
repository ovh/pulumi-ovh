// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Okms
{
    public static class GetOkmsCredential
    {
        public static Task<GetOkmsCredentialResult> InvokeAsync(GetOkmsCredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOkmsCredentialResult>("ovh:Okms/getOkmsCredential:getOkmsCredential", args ?? new GetOkmsCredentialArgs(), options.WithDefaults());

        public static Output<GetOkmsCredentialResult> Invoke(GetOkmsCredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOkmsCredentialResult>("ovh:Okms/getOkmsCredential:getOkmsCredential", args ?? new GetOkmsCredentialInvokeArgs(), options.WithDefaults());

        public static Output<GetOkmsCredentialResult> Invoke(GetOkmsCredentialInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOkmsCredentialResult>("ovh:Okms/getOkmsCredential:getOkmsCredential", args ?? new GetOkmsCredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOkmsCredentialArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        [Input("okmsId", required: true)]
        public string OkmsId { get; set; } = null!;

        public GetOkmsCredentialArgs()
        {
        }
        public static new GetOkmsCredentialArgs Empty => new GetOkmsCredentialArgs();
    }

    public sealed class GetOkmsCredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("okmsId", required: true)]
        public Input<string> OkmsId { get; set; } = null!;

        public GetOkmsCredentialInvokeArgs()
        {
        }
        public static new GetOkmsCredentialInvokeArgs Empty => new GetOkmsCredentialInvokeArgs();
    }


    [OutputType]
    public sealed class GetOkmsCredentialResult
    {
        public readonly string CertificatePem;
        public readonly string CreatedAt;
        public readonly string Description;
        public readonly string ExpiredAt;
        public readonly bool FromCsr;
        public readonly string Id;
        public readonly ImmutableArray<string> IdentityUrns;
        public readonly string Name;
        public readonly string OkmsId;
        public readonly string Status;

        [OutputConstructor]
        private GetOkmsCredentialResult(
            string certificatePem,

            string createdAt,

            string description,

            string expiredAt,

            bool fromCsr,

            string id,

            ImmutableArray<string> identityUrns,

            string name,

            string okmsId,

            string status)
        {
            CertificatePem = certificatePem;
            CreatedAt = createdAt;
            Description = description;
            ExpiredAt = expiredAt;
            FromCsr = fromCsr;
            Id = id;
            IdentityUrns = identityUrns;
            Name = name;
            OkmsId = okmsId;
            Status = status;
        }
    }
}
