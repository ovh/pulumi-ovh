// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// Use this data source to retrieve data associated with a KMS credential, such as the PEM encoded certificate.
        /// </summary>
        public static Task<GetOkmsCredentialResult> InvokeAsync(GetOkmsCredentialArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOkmsCredentialResult>("ovh:Okms/getOkmsCredential:getOkmsCredential", args ?? new GetOkmsCredentialArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve data associated with a KMS credential, such as the PEM encoded certificate.
        /// </summary>
        public static Output<GetOkmsCredentialResult> Invoke(GetOkmsCredentialInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOkmsCredentialResult>("ovh:Okms/getOkmsCredential:getOkmsCredential", args ?? new GetOkmsCredentialInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve data associated with a KMS credential, such as the PEM encoded certificate.
        /// </summary>
        public static Output<GetOkmsCredentialResult> Invoke(GetOkmsCredentialInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOkmsCredentialResult>("ovh:Okms/getOkmsCredential:getOkmsCredential", args ?? new GetOkmsCredentialInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOkmsCredentialArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the credential
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        /// <summary>
        /// ID of the KMS
        /// </summary>
        [Input("okmsId", required: true)]
        public string OkmsId { get; set; } = null!;

        public GetOkmsCredentialArgs()
        {
        }
        public static new GetOkmsCredentialArgs Empty => new GetOkmsCredentialArgs();
    }

    public sealed class GetOkmsCredentialInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the credential
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// ID of the KMS
        /// </summary>
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
        /// <summary>
        /// (String) PEM encoded certificate of the credential
        /// </summary>
        public readonly string CertificatePem;
        /// <summary>
        /// (String) Creation time of the credential
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// (String) Description of the credential
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (String) Expiration time of the credential
        /// </summary>
        public readonly string ExpiredAt;
        /// <summary>
        /// (Boolean) Is the credential generated from CSR
        /// </summary>
        public readonly bool FromCsr;
        public readonly string Id;
        /// <summary>
        /// (List of String) List of identity URNs associated with the credential
        /// </summary>
        public readonly ImmutableArray<string> IdentityUrns;
        /// <summary>
        /// (String) Name of the credential
        /// </summary>
        public readonly string Name;
        public readonly string OkmsId;
        /// <summary>
        /// (String) Status of the credential
        /// </summary>
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
