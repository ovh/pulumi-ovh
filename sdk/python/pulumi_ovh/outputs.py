# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetCloudProjectFlavorCapabilityResult',
    'GetCloudProjectFlavorPlanCodesResult',
    'GetCloudProjectRancherCapabilitiesPlanPlanResult',
    'GetCloudProjectRancherCapabilitiesVersionVersionResult',
    'GetCloudProjectSshKeysSshKeyResult',
    'GetInstallationTemplateInputResult',
    'GetInstallationTemplateLicenseResult',
    'GetInstallationTemplateLicenseOResult',
    'GetInstallationTemplateLicenseUsageResult',
    'GetInstallationTemplateProjectResult',
    'GetInstallationTemplateProjectOResult',
    'GetInstallationTemplateProjectUsageResult',
    'GetOvhcloudConnectConfigPopDatacenterExtrasExtraConfigResult',
    'GetOvhcloudConnectConfigPopDatacentersDatacenterConfigResult',
    'GetOvhcloudConnectConfigPopsPopConfigResult',
    'GetOvhcloudConnectDatacentersDatacenterResult',
    'GetServerVniResult',
    'GetStorageEfsIamResult',
    'GetStorageEfsShareAccessPathsAccessPathResult',
]

@pulumi.output_type
class GetCloudProjectFlavorCapabilityResult(dict):
    def __init__(__self__, *,
                 enabled: builtins.bool,
                 name: builtins.str):
        """
        :param builtins.bool enabled: Is the capability enabled
        :param builtins.str name: Name of the capability
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        """
        Is the capability enabled
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the capability
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetCloudProjectFlavorPlanCodesResult(dict):
    def __init__(__self__, *,
                 hourly: builtins.str,
                 monthly: builtins.str):
        """
        :param builtins.str hourly: Plan code to order hourly instance
        :param builtins.str monthly: Plan code to order monthly instance
        """
        pulumi.set(__self__, "hourly", hourly)
        pulumi.set(__self__, "monthly", monthly)

    @property
    @pulumi.getter
    def hourly(self) -> builtins.str:
        """
        Plan code to order hourly instance
        """
        return pulumi.get(self, "hourly")

    @property
    @pulumi.getter
    def monthly(self) -> builtins.str:
        """
        Plan code to order monthly instance
        """
        return pulumi.get(self, "monthly")


@pulumi.output_type
class GetCloudProjectRancherCapabilitiesPlanPlanResult(dict):
    def __init__(__self__, *,
                 cause: builtins.str,
                 message: builtins.str,
                 name: builtins.str,
                 status: builtins.str):
        """
        :param builtins.str cause: Cause for an unavailability
        :param builtins.str message: Human-readable description of the unavailability cause
        :param builtins.str name: Name of the plan
        :param builtins.str status: Status of the plan
        """
        pulumi.set(__self__, "cause", cause)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def cause(self) -> builtins.str:
        """
        Cause for an unavailability
        """
        return pulumi.get(self, "cause")

    @property
    @pulumi.getter
    def message(self) -> builtins.str:
        """
        Human-readable description of the unavailability cause
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the plan
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of the plan
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetCloudProjectRancherCapabilitiesVersionVersionResult(dict):
    def __init__(__self__, *,
                 cause: builtins.str,
                 changelog_url: builtins.str,
                 message: builtins.str,
                 name: builtins.str,
                 status: builtins.str):
        """
        :param builtins.str cause: Cause for an unavailability
        :param builtins.str changelog_url: Changelog URL of the version
        :param builtins.str message: Human-readable description of the unavailability cause
        :param builtins.str name: Name of the version
        :param builtins.str status: Status of the version
        """
        pulumi.set(__self__, "cause", cause)
        pulumi.set(__self__, "changelog_url", changelog_url)
        pulumi.set(__self__, "message", message)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def cause(self) -> builtins.str:
        """
        Cause for an unavailability
        """
        return pulumi.get(self, "cause")

    @property
    @pulumi.getter(name="changelogUrl")
    def changelog_url(self) -> builtins.str:
        """
        Changelog URL of the version
        """
        return pulumi.get(self, "changelog_url")

    @property
    @pulumi.getter
    def message(self) -> builtins.str:
        """
        Human-readable description of the unavailability cause
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the version
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of the version
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GetCloudProjectSshKeysSshKeyResult(dict):
    def __init__(__self__, *,
                 id: builtins.str,
                 name: builtins.str,
                 public_key: builtins.str,
                 regions: Sequence[builtins.str]):
        """
        :param builtins.str id: SSH key ID
        :param builtins.str name: SSH key name
        :param builtins.str public_key: SSH public key
        :param Sequence[builtins.str] regions: SSH key regions
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "public_key", public_key)
        pulumi.set(__self__, "regions", regions)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        SSH key ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        SSH key name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> builtins.str:
        """
        SSH public key
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter
    def regions(self) -> Sequence[builtins.str]:
        """
        SSH key regions
        """
        return pulumi.get(self, "regions")


@pulumi.output_type
class GetInstallationTemplateInputResult(dict):
    def __init__(__self__, *,
                 default: builtins.str,
                 description: builtins.str,
                 enums: Sequence[builtins.str],
                 mandatory: builtins.bool,
                 name: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str description: Information about this template.
        """
        pulumi.set(__self__, "default", default)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "enums", enums)
        pulumi.set(__self__, "mandatory", mandatory)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def default(self) -> builtins.str:
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Information about this template.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enums(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "enums")

    @property
    @pulumi.getter
    def mandatory(self) -> builtins.bool:
        return pulumi.get(self, "mandatory")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        return pulumi.get(self, "type")


@pulumi.output_type
class GetInstallationTemplateLicenseResult(dict):
    def __init__(__self__, *,
                 os: Sequence['outputs.GetInstallationTemplateLicenseOResult'],
                 usages: Sequence['outputs.GetInstallationTemplateLicenseUsageResult']):
        pulumi.set(__self__, "os", os)
        pulumi.set(__self__, "usages", usages)

    @property
    @pulumi.getter
    def os(self) -> Sequence['outputs.GetInstallationTemplateLicenseOResult']:
        return pulumi.get(self, "os")

    @property
    @pulumi.getter
    def usages(self) -> Sequence['outputs.GetInstallationTemplateLicenseUsageResult']:
        return pulumi.get(self, "usages")


@pulumi.output_type
class GetInstallationTemplateLicenseOResult(dict):
    def __init__(__self__, *,
                 names: Sequence[builtins.str],
                 url: builtins.str):
        pulumi.set(__self__, "names", names)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def names(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        return pulumi.get(self, "url")


@pulumi.output_type
class GetInstallationTemplateLicenseUsageResult(dict):
    def __init__(__self__, *,
                 names: Sequence[builtins.str],
                 url: builtins.str):
        pulumi.set(__self__, "names", names)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def names(self) -> Sequence[builtins.str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        return pulumi.get(self, "url")


@pulumi.output_type
class GetInstallationTemplateProjectResult(dict):
    def __init__(__self__, *,
                 os: Sequence['outputs.GetInstallationTemplateProjectOResult'],
                 usages: Sequence['outputs.GetInstallationTemplateProjectUsageResult']):
        """
        :param Sequence['GetInstallationTemplateProjectOArgs'] os: OS template project OS details
        """
        pulumi.set(__self__, "os", os)
        pulumi.set(__self__, "usages", usages)

    @property
    @pulumi.getter
    def os(self) -> Sequence['outputs.GetInstallationTemplateProjectOResult']:
        """
        OS template project OS details
        """
        return pulumi.get(self, "os")

    @property
    @pulumi.getter
    def usages(self) -> Sequence['outputs.GetInstallationTemplateProjectUsageResult']:
        return pulumi.get(self, "usages")


@pulumi.output_type
class GetInstallationTemplateProjectOResult(dict):
    def __init__(__self__, *,
                 governances: Sequence[builtins.str],
                 name: builtins.str,
                 release_notes: builtins.str,
                 url: builtins.str,
                 version: builtins.str):
        """
        :param Sequence[builtins.str] governances: OS template project item governance
        :param builtins.str name: OS template project item name
        :param builtins.str release_notes: OS template project item release notes
        :param builtins.str url: OS template project item url
        :param builtins.str version: OS template project item version
        """
        pulumi.set(__self__, "governances", governances)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "release_notes", release_notes)
        pulumi.set(__self__, "url", url)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def governances(self) -> Sequence[builtins.str]:
        """
        OS template project item governance
        """
        return pulumi.get(self, "governances")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        OS template project item name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="releaseNotes")
    def release_notes(self) -> builtins.str:
        """
        OS template project item release notes
        """
        return pulumi.get(self, "release_notes")

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        """
        OS template project item url
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> builtins.str:
        """
        OS template project item version
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class GetInstallationTemplateProjectUsageResult(dict):
    def __init__(__self__, *,
                 governances: Sequence[builtins.str],
                 name: builtins.str,
                 release_notes: builtins.str,
                 url: builtins.str,
                 version: builtins.str):
        """
        :param Sequence[builtins.str] governances: OS template project item governance
        :param builtins.str name: OS template project item name
        :param builtins.str release_notes: OS template project item release notes
        :param builtins.str url: OS template project item url
        :param builtins.str version: OS template project item version
        """
        pulumi.set(__self__, "governances", governances)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "release_notes", release_notes)
        pulumi.set(__self__, "url", url)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def governances(self) -> Sequence[builtins.str]:
        """
        OS template project item governance
        """
        return pulumi.get(self, "governances")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        OS template project item name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="releaseNotes")
    def release_notes(self) -> builtins.str:
        """
        OS template project item release notes
        """
        return pulumi.get(self, "release_notes")

    @property
    @pulumi.getter
    def url(self) -> builtins.str:
        """
        OS template project item url
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> builtins.str:
        """
        OS template project item version
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class GetOvhcloudConnectConfigPopDatacenterExtrasExtraConfigResult(dict):
    def __init__(__self__, *,
                 bgp_neighbor_area: builtins.float,
                 bgp_neighbor_ip: builtins.str,
                 id: builtins.float,
                 next_hop: builtins.str,
                 status: builtins.str,
                 subnet: builtins.str,
                 type: builtins.str):
        """
        :param builtins.float bgp_neighbor_area: BGP AS number
        :param builtins.str bgp_neighbor_ip: Router IP for BGP
        :param builtins.float id: ID of the extra configuration
        :param builtins.str next_hop: Static route next hop
        :param builtins.str status: Status of the pop configuration
        :param builtins.str subnet: Static route ip
        :param builtins.str type: Type of the configuration
        """
        pulumi.set(__self__, "bgp_neighbor_area", bgp_neighbor_area)
        pulumi.set(__self__, "bgp_neighbor_ip", bgp_neighbor_ip)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "next_hop", next_hop)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet", subnet)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="bgpNeighborArea")
    def bgp_neighbor_area(self) -> builtins.float:
        """
        BGP AS number
        """
        return pulumi.get(self, "bgp_neighbor_area")

    @property
    @pulumi.getter(name="bgpNeighborIp")
    def bgp_neighbor_ip(self) -> builtins.str:
        """
        Router IP for BGP
        """
        return pulumi.get(self, "bgp_neighbor_ip")

    @property
    @pulumi.getter
    def id(self) -> builtins.float:
        """
        ID of the extra configuration
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nextHop")
    def next_hop(self) -> builtins.str:
        """
        Static route next hop
        """
        return pulumi.get(self, "next_hop")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of the pop configuration
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def subnet(self) -> builtins.str:
        """
        Static route ip
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of the configuration
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetOvhcloudConnectConfigPopDatacentersDatacenterConfigResult(dict):
    def __init__(__self__, *,
                 datacenter_id: builtins.float,
                 id: builtins.float,
                 ovh_bgp_area: builtins.float,
                 status: builtins.str,
                 subnet: builtins.str):
        """
        :param builtins.float datacenter_id: Datacenter ID
        :param builtins.float id: ID of the Datacenter configuration
        :param builtins.float ovh_bgp_area: OVH Private AS
        :param builtins.str status: Status of the pop configuration
        :param builtins.str subnet: Subnet should be a /28 min
        """
        pulumi.set(__self__, "datacenter_id", datacenter_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "ovh_bgp_area", ovh_bgp_area)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet", subnet)

    @property
    @pulumi.getter(name="datacenterId")
    def datacenter_id(self) -> builtins.float:
        """
        Datacenter ID
        """
        return pulumi.get(self, "datacenter_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.float:
        """
        ID of the Datacenter configuration
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ovhBgpArea")
    def ovh_bgp_area(self) -> builtins.float:
        """
        OVH Private AS
        """
        return pulumi.get(self, "ovh_bgp_area")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of the pop configuration
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def subnet(self) -> builtins.str:
        """
        Subnet should be a /28 min
        """
        return pulumi.get(self, "subnet")


@pulumi.output_type
class GetOvhcloudConnectConfigPopsPopConfigResult(dict):
    def __init__(__self__, *,
                 customer_bgp_area: builtins.float,
                 id: builtins.float,
                 interface_id: builtins.float,
                 ovh_bgp_area: builtins.float,
                 status: builtins.str,
                 subnet: builtins.str,
                 type: builtins.str):
        """
        :param builtins.float customer_bgp_area: Customer Private AS
        :param builtins.float id: ID of the Pop Configuration
        :param builtins.float interface_id: ID of the interface
        :param builtins.float ovh_bgp_area: OVH Private AS
        :param builtins.str status: Status of the pop configuration
        :param builtins.str subnet: Subnet should be a /30, first IP for OVH, second IP for customer
        :param builtins.str type: Type of the pop configuration
        """
        pulumi.set(__self__, "customer_bgp_area", customer_bgp_area)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "interface_id", interface_id)
        pulumi.set(__self__, "ovh_bgp_area", ovh_bgp_area)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet", subnet)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="customerBgpArea")
    def customer_bgp_area(self) -> builtins.float:
        """
        Customer Private AS
        """
        return pulumi.get(self, "customer_bgp_area")

    @property
    @pulumi.getter
    def id(self) -> builtins.float:
        """
        ID of the Pop Configuration
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> builtins.float:
        """
        ID of the interface
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter(name="ovhBgpArea")
    def ovh_bgp_area(self) -> builtins.float:
        """
        OVH Private AS
        """
        return pulumi.get(self, "ovh_bgp_area")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Status of the pop configuration
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def subnet(self) -> builtins.str:
        """
        Subnet should be a /30, first IP for OVH, second IP for customer
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of the pop configuration
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetOvhcloudConnectDatacentersDatacenterResult(dict):
    def __init__(__self__, *,
                 available: builtins.bool,
                 id: builtins.float,
                 name: builtins.str,
                 region: builtins.str,
                 region_type: builtins.str):
        """
        :param builtins.bool available: Get availability to add new configuration on it
        :param builtins.float id: Id
        :param builtins.str name: name of the datacenter
        :param builtins.str region: region of the datacenter
        :param builtins.str region_type: region type of the datacenter
        """
        pulumi.set(__self__, "available", available)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "region_type", region_type)

    @property
    @pulumi.getter
    def available(self) -> builtins.bool:
        """
        Get availability to add new configuration on it
        """
        return pulumi.get(self, "available")

    @property
    @pulumi.getter
    def id(self) -> builtins.float:
        """
        Id
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        name of the datacenter
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        """
        region of the datacenter
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionType")
    def region_type(self) -> builtins.str:
        """
        region type of the datacenter
        """
        return pulumi.get(self, "region_type")


@pulumi.output_type
class GetServerVniResult(dict):
    def __init__(__self__, *,
                 enabled: builtins.bool,
                 mode: builtins.str,
                 name: builtins.str,
                 nics: Sequence[builtins.str],
                 server_name: builtins.str,
                 uuid: builtins.str,
                 vrack: builtins.str):
        """
        :param builtins.bool enabled: VirtualNetworkInterface activation state
        :param builtins.str mode: VirtualNetworkInterface mode (public,vrack,vrack_aggregation)
        :param builtins.str name: User defined VirtualNetworkInterface name
        :param Sequence[builtins.str] nics: NetworkInterfaceControllers bound to this VirtualNetworkInterface
        :param builtins.str server_name: Server bound to this VirtualNetworkInterface
        :param builtins.str uuid: VirtualNetworkInterface unique id
        :param builtins.str vrack: vRack name
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "nics", nics)
        pulumi.set(__self__, "server_name", server_name)
        pulumi.set(__self__, "uuid", uuid)
        pulumi.set(__self__, "vrack", vrack)

    @property
    @pulumi.getter
    def enabled(self) -> builtins.bool:
        """
        VirtualNetworkInterface activation state
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def mode(self) -> builtins.str:
        """
        VirtualNetworkInterface mode (public,vrack,vrack_aggregation)
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        User defined VirtualNetworkInterface name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def nics(self) -> Sequence[builtins.str]:
        """
        NetworkInterfaceControllers bound to this VirtualNetworkInterface
        """
        return pulumi.get(self, "nics")

    @property
    @pulumi.getter(name="serverName")
    def server_name(self) -> builtins.str:
        """
        Server bound to this VirtualNetworkInterface
        """
        return pulumi.get(self, "server_name")

    @property
    @pulumi.getter
    def uuid(self) -> builtins.str:
        """
        VirtualNetworkInterface unique id
        """
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vrack(self) -> builtins.str:
        """
        vRack name
        """
        return pulumi.get(self, "vrack")


@pulumi.output_type
class GetStorageEfsIamResult(dict):
    def __init__(__self__, *,
                 display_name: builtins.str,
                 id: builtins.str,
                 tags: Mapping[str, builtins.str],
                 urn: builtins.str):
        """
        :param builtins.str display_name: Resource display name
        :param builtins.str id: Unique identifier of the resource
        :param Mapping[str, builtins.str] tags: Resource tags. Tags that were internally computed are prefixed with ovh:
        :param builtins.str urn: Unique resource name used in policies
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> builtins.str:
        """
        Resource display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Unique identifier of the resource
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, builtins.str]:
        """
        Resource tags. Tags that were internally computed are prefixed with ovh:
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> builtins.str:
        """
        Unique resource name used in policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class GetStorageEfsShareAccessPathsAccessPathResult(dict):
    def __init__(__self__, *,
                 id: builtins.str,
                 path: builtins.str,
                 preferred: builtins.bool):
        """
        :param builtins.str id: Access path ID
        :param builtins.str path: Access path
        :param builtins.bool preferred: Is this the preferred access path?
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "preferred", preferred)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Access path ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def path(self) -> builtins.str:
        """
        Access path
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def preferred(self) -> builtins.bool:
        """
        Is this the preferred access path?
        """
        return pulumi.get(self, "preferred")


