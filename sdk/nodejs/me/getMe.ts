// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about the current OVHcloud account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myaccount = ovh.Me.getMe({});
 * ```
 */
export function getMe(opts?: pulumi.InvokeOptions): Promise<GetMeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Me/getMe:getMe", {
    }, opts);
}

/**
 * A collection of values returned by getMe.
 */
export interface GetMeResult {
    /**
     * Postal address of the account
     */
    readonly address: string;
    /**
     * Area of the account
     */
    readonly area: string;
    /**
     * City of birth
     */
    readonly birthCity: string;
    /**
     * Birth date
     */
    readonly birthDay: string;
    /**
     * City of the account
     */
    readonly city: string;
    /**
     * This is the national identification number of the company that possess this account
     */
    readonly companyNationalIdentificationNumber: string;
    /**
     * Type of corporation
     */
    readonly corporationType: string;
    /**
     * Country of the account
     */
    readonly country: string;
    readonly currencies: outputs.Me.GetMeCurrency[];
    /**
     * The customer code of this account (a numerical value used for identification when contacting support via phone call)
     */
    readonly customerCode: string;
    /**
     * Email address
     */
    readonly email: string;
    /**
     * Fax number
     */
    readonly fax: string;
    /**
     * First name
     */
    readonly firstname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Italian SDI
     */
    readonly italianSdi: string;
    /**
     * Preferred language for this account
     */
    readonly language: string;
    /**
     * Legal form of the account
     */
    readonly legalform: string;
    /**
     * Name of the account holder
     */
    readonly name: string;
    /**
     * National Identification Number of this account
     */
    readonly nationalIdentificationNumber: string;
    /**
     * Nic handle / customer identifier
     */
    readonly nichandle: string;
    /**
     * Name of the organisation for this account
     */
    readonly organisation: string;
    /**
     * OVHcloud subsidiary
     */
    readonly ovhCompany: string;
    /**
     * OVHcloud subsidiary
     */
    readonly ovhSubsidiary: string;
    /**
     * Phone number
     */
    readonly phone: string;
    /**
     * Country code of the phone number
     */
    readonly phoneCountry: string;
    /**
     * Gender of the account holder
     */
    readonly sex: string;
    /**
     * Backup email address
     */
    readonly spareEmail: string;
    /**
     * State of the postal address
     */
    readonly state: string;
    /**
     * The resource URN of the account, to be used when writing IAM policies
     */
    readonly urn: string;
    /**
     * VAT number
     */
    readonly vat: string;
    /**
     * Zipcode of the address
     */
    readonly zip: string;
}
/**
 * Use this data source to get information about the current OVHcloud account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const myaccount = ovh.Me.getMe({});
 * ```
 */
export function getMeOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetMeResult> {
    return pulumi.output(getMe(opts))
}
