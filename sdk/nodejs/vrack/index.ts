// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CloudProjectArgs, CloudProjectState } from "./cloudProject";
export type CloudProject = import("./cloudProject").CloudProject;
export const CloudProject: typeof import("./cloudProject").CloudProject = null as any;
utilities.lazyLoad(exports, ["CloudProject"], () => require("./cloudProject"));

export { DedicatedServerArgs, DedicatedServerState } from "./dedicatedServer";
export type DedicatedServer = import("./dedicatedServer").DedicatedServer;
export const DedicatedServer: typeof import("./dedicatedServer").DedicatedServer = null as any;
utilities.lazyLoad(exports, ["DedicatedServer"], () => require("./dedicatedServer"));

export { DedicatedServerInterfaceArgs, DedicatedServerInterfaceState } from "./dedicatedServerInterface";
export type DedicatedServerInterface = import("./dedicatedServerInterface").DedicatedServerInterface;
export const DedicatedServerInterface: typeof import("./dedicatedServerInterface").DedicatedServerInterface = null as any;
utilities.lazyLoad(exports, ["DedicatedServerInterface"], () => require("./dedicatedServerInterface"));

export { GetVracksResult } from "./getVracks";
export const getVracks: typeof import("./getVracks").getVracks = null as any;
export const getVracksOutput: typeof import("./getVracks").getVracksOutput = null as any;
utilities.lazyLoad(exports, ["getVracks","getVracksOutput"], () => require("./getVracks"));

export { IpAddressArgs, IpAddressState } from "./ipAddress";
export type IpAddress = import("./ipAddress").IpAddress;
export const IpAddress: typeof import("./ipAddress").IpAddress = null as any;
utilities.lazyLoad(exports, ["IpAddress"], () => require("./ipAddress"));

export { IpLoadbalancingArgs, IpLoadbalancingState } from "./ipLoadbalancing";
export type IpLoadbalancing = import("./ipLoadbalancing").IpLoadbalancing;
export const IpLoadbalancing: typeof import("./ipLoadbalancing").IpLoadbalancing = null as any;
utilities.lazyLoad(exports, ["IpLoadbalancing"], () => require("./ipLoadbalancing"));

export { IpV6Args, IpV6State } from "./ipV6";
export type IpV6 = import("./ipV6").IpV6;
export const IpV6: typeof import("./ipV6").IpV6 = null as any;
utilities.lazyLoad(exports, ["IpV6"], () => require("./ipV6"));

export { OVHcloudConnectArgs, OVHcloudConnectState } from "./ovhcloudConnect";
export type OVHcloudConnect = import("./ovhcloudConnect").OVHcloudConnect;
export const OVHcloudConnect: typeof import("./ovhcloudConnect").OVHcloudConnect = null as any;
utilities.lazyLoad(exports, ["OVHcloudConnect"], () => require("./ovhcloudConnect"));

export { VrackArgs, VrackState } from "./vrack";
export type Vrack = import("./vrack").Vrack;
export const Vrack: typeof import("./vrack").Vrack = null as any;
utilities.lazyLoad(exports, ["Vrack"], () => require("./vrack"));

export { VrackservicesArgs, VrackservicesState } from "./vrackservices";
export type Vrackservices = import("./vrackservices").Vrackservices;
export const Vrackservices: typeof import("./vrackservices").Vrackservices = null as any;
utilities.lazyLoad(exports, ["Vrackservices"], () => require("./vrackservices"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ovh:Vrack/cloudProject:CloudProject":
                return new CloudProject(name, <any>undefined, { urn })
            case "ovh:Vrack/dedicatedServer:DedicatedServer":
                return new DedicatedServer(name, <any>undefined, { urn })
            case "ovh:Vrack/dedicatedServerInterface:DedicatedServerInterface":
                return new DedicatedServerInterface(name, <any>undefined, { urn })
            case "ovh:Vrack/ipAddress:IpAddress":
                return new IpAddress(name, <any>undefined, { urn })
            case "ovh:Vrack/ipLoadbalancing:IpLoadbalancing":
                return new IpLoadbalancing(name, <any>undefined, { urn })
            case "ovh:Vrack/ipV6:IpV6":
                return new IpV6(name, <any>undefined, { urn })
            case "ovh:Vrack/oVHcloudConnect:OVHcloudConnect":
                return new OVHcloudConnect(name, <any>undefined, { urn })
            case "ovh:Vrack/vrack:Vrack":
                return new Vrack(name, <any>undefined, { urn })
            case "ovh:Vrack/vrackservices:Vrackservices":
                return new Vrackservices(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ovh", "Vrack/cloudProject", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/dedicatedServer", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/dedicatedServerInterface", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/ipAddress", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/ipLoadbalancing", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/ipV6", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/oVHcloudConnect", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/vrack", _module)
pulumi.runtime.registerResourceModule("ovh", "Vrack/vrackservices", _module)
