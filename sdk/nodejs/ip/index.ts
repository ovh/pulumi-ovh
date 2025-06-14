// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FirewallArgs, FirewallState } from "./firewall";
export type Firewall = import("./firewall").Firewall;
export const Firewall: typeof import("./firewall").Firewall = null as any;
utilities.lazyLoad(exports, ["Firewall"], () => require("./firewall"));

export { FirewallRuleArgs, FirewallRuleState } from "./firewallRule";
export type FirewallRule = import("./firewallRule").FirewallRule;
export const FirewallRule: typeof import("./firewallRule").FirewallRule = null as any;
utilities.lazyLoad(exports, ["FirewallRule"], () => require("./firewallRule"));

export { GetFirewallArgs, GetFirewallResult, GetFirewallOutputArgs } from "./getFirewall";
export const getFirewall: typeof import("./getFirewall").getFirewall = null as any;
export const getFirewallOutput: typeof import("./getFirewall").getFirewallOutput = null as any;
utilities.lazyLoad(exports, ["getFirewall","getFirewallOutput"], () => require("./getFirewall"));

export { GetFirewallRuleArgs, GetFirewallRuleResult, GetFirewallRuleOutputArgs } from "./getFirewallRule";
export const getFirewallRule: typeof import("./getFirewallRule").getFirewallRule = null as any;
export const getFirewallRuleOutput: typeof import("./getFirewallRule").getFirewallRuleOutput = null as any;
utilities.lazyLoad(exports, ["getFirewallRule","getFirewallRuleOutput"], () => require("./getFirewallRule"));

export { GetMitigationArgs, GetMitigationResult, GetMitigationOutputArgs } from "./getMitigation";
export const getMitigation: typeof import("./getMitigation").getMitigation = null as any;
export const getMitigationOutput: typeof import("./getMitigation").getMitigationOutput = null as any;
utilities.lazyLoad(exports, ["getMitigation","getMitigationOutput"], () => require("./getMitigation"));

export { GetServiceArgs, GetServiceResult, GetServiceOutputArgs } from "./getService";
export const getService: typeof import("./getService").getService = null as any;
export const getServiceOutput: typeof import("./getService").getServiceOutput = null as any;
utilities.lazyLoad(exports, ["getService","getServiceOutput"], () => require("./getService"));

export { IpServiceArgs, IpServiceState } from "./ipService";
export type IpService = import("./ipService").IpService;
export const IpService: typeof import("./ipService").IpService = null as any;
utilities.lazyLoad(exports, ["IpService"], () => require("./ipService"));

export { MitigationArgs, MitigationState } from "./mitigation";
export type Mitigation = import("./mitigation").Mitigation;
export const Mitigation: typeof import("./mitigation").Mitigation = null as any;
utilities.lazyLoad(exports, ["Mitigation"], () => require("./mitigation"));

export { MoveArgs, MoveState } from "./move";
export type Move = import("./move").Move;
export const Move: typeof import("./move").Move = null as any;
utilities.lazyLoad(exports, ["Move"], () => require("./move"));

export { ReverseArgs, ReverseState } from "./reverse";
export type Reverse = import("./reverse").Reverse;
export const Reverse: typeof import("./reverse").Reverse = null as any;
utilities.lazyLoad(exports, ["Reverse"], () => require("./reverse"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ovh:Ip/firewall:Firewall":
                return new Firewall(name, <any>undefined, { urn })
            case "ovh:Ip/firewallRule:FirewallRule":
                return new FirewallRule(name, <any>undefined, { urn })
            case "ovh:Ip/ipService:IpService":
                return new IpService(name, <any>undefined, { urn })
            case "ovh:Ip/mitigation:Mitigation":
                return new Mitigation(name, <any>undefined, { urn })
            case "ovh:Ip/move:Move":
                return new Move(name, <any>undefined, { urn })
            case "ovh:Ip/reverse:Reverse":
                return new Reverse(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ovh", "Ip/firewall", _module)
pulumi.runtime.registerResourceModule("ovh", "Ip/firewallRule", _module)
pulumi.runtime.registerResourceModule("ovh", "Ip/ipService", _module)
pulumi.runtime.registerResourceModule("ovh", "Ip/mitigation", _module)
pulumi.runtime.registerResourceModule("ovh", "Ip/move", _module)
pulumi.runtime.registerResourceModule("ovh", "Ip/reverse", _module)
