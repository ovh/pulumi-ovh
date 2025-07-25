// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to retrieve data associated with a KMS credential, such as the PEM encoded certificate.
 */
export function getOkmsCredential(args: GetOkmsCredentialArgs, opts?: pulumi.InvokeOptions): Promise<GetOkmsCredentialResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Okms/getOkmsCredential:getOkmsCredential", {
        "id": args.id,
        "okmsId": args.okmsId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOkmsCredential.
 */
export interface GetOkmsCredentialArgs {
    /**
     * ID of the credential
     */
    id: string;
    /**
     * ID of the KMS
     */
    okmsId: string;
}

/**
 * A collection of values returned by getOkmsCredential.
 */
export interface GetOkmsCredentialResult {
    /**
     * (String) PEM encoded certificate of the credential
     */
    readonly certificatePem: string;
    /**
     * (String) Creation time of the credential
     */
    readonly createdAt: string;
    /**
     * (String) Description of the credential
     */
    readonly description: string;
    /**
     * (String) Expiration time of the credential
     */
    readonly expiredAt: string;
    /**
     * (Boolean) Is the credential generated from CSR
     */
    readonly fromCsr: boolean;
    readonly id: string;
    /**
     * (List of String) List of identity URNs associated with the credential
     */
    readonly identityUrns: string[];
    /**
     * (String) Name of the credential
     */
    readonly name: string;
    readonly okmsId: string;
    /**
     * (String) Status of the credential
     */
    readonly status: string;
}
/**
 * Use this data source to retrieve data associated with a KMS credential, such as the PEM encoded certificate.
 */
export function getOkmsCredentialOutput(args: GetOkmsCredentialOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOkmsCredentialResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ovh:Okms/getOkmsCredential:getOkmsCredential", {
        "id": args.id,
        "okmsId": args.okmsId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOkmsCredential.
 */
export interface GetOkmsCredentialOutputArgs {
    /**
     * ID of the credential
     */
    id: pulumi.Input<string>;
    /**
     * ID of the KMS
     */
    okmsId: pulumi.Input<string>;
}
