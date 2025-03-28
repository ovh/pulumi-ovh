// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { GetInstallationTemplateArgs, GetInstallationTemplateResult, GetInstallationTemplateOutputArgs } from "./getInstallationTemplate";
export const getInstallationTemplate: typeof import("./getInstallationTemplate").getInstallationTemplate = null as any;
export const getInstallationTemplateOutput: typeof import("./getInstallationTemplate").getInstallationTemplateOutput = null as any;
utilities.lazyLoad(exports, ["getInstallationTemplate","getInstallationTemplateOutput"], () => require("./getInstallationTemplate"));

export { GetInstallationTemplatesResult } from "./getInstallationTemplates";
export const getInstallationTemplates: typeof import("./getInstallationTemplates").getInstallationTemplates = null as any;
export const getInstallationTemplatesOutput: typeof import("./getInstallationTemplates").getInstallationTemplatesOutput = null as any;
utilities.lazyLoad(exports, ["getInstallationTemplates","getInstallationTemplatesOutput"], () => require("./getInstallationTemplates"));

export { GetServerArgs, GetServerResult, GetServerOutputArgs } from "./getServer";
export const getServer: typeof import("./getServer").getServer = null as any;
export const getServerOutput: typeof import("./getServer").getServerOutput = null as any;
utilities.lazyLoad(exports, ["getServer","getServerOutput"], () => require("./getServer"));

export { GetServersResult } from "./getServers";
export const getServers: typeof import("./getServers").getServers = null as any;
export const getServersOutput: typeof import("./getServers").getServersOutput = null as any;
utilities.lazyLoad(exports, ["getServers","getServersOutput"], () => require("./getServers"));

export { GetVrackNetworksArgs, GetVrackNetworksResult, GetVrackNetworksOutputArgs } from "./getVrackNetworks";
export const getVrackNetworks: typeof import("./getVrackNetworks").getVrackNetworks = null as any;
export const getVrackNetworksOutput: typeof import("./getVrackNetworks").getVrackNetworksOutput = null as any;
utilities.lazyLoad(exports, ["getVrackNetworks","getVrackNetworksOutput"], () => require("./getVrackNetworks"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));


// Export sub-modules:
import * as cloud from "./cloud";
import * as cloudproject from "./cloudproject";
import * as cloudprojectdatabase from "./cloudprojectdatabase";
import * as config from "./config";
import * as dbaas from "./dbaas";
import * as dedicated from "./dedicated";
import * as domain from "./domain";
import * as hosting from "./hosting";
import * as iam from "./iam";
import * as ip from "./ip";
import * as iploadbalancing from "./iploadbalancing";
import * as me from "./me";
import * as okms from "./okms";
import * as order from "./order";
import * as ovhcloud from "./ovhcloud";
import * as savingsplan from "./savingsplan";
import * as types from "./types";
import * as vmware from "./vmware";
import * as vps from "./vps";
import * as vrack from "./vrack";

export {
    cloud,
    cloudproject,
    cloudprojectdatabase,
    config,
    dbaas,
    dedicated,
    domain,
    hosting,
    iam,
    ip,
    iploadbalancing,
    me,
    okms,
    order,
    ovhcloud,
    savingsplan,
    types,
    vmware,
    vps,
    vrack,
};
pulumi.runtime.registerResourcePackage("ovh", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:ovh") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
