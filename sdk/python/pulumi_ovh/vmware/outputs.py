# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
from .. import _utilities
from . import outputs

__all__ = [
    'GetCloudDirectorBackupCurrentStateResult',
    'GetCloudDirectorBackupCurrentStateOfferResult',
    'GetCloudDirectorBackupCurrentTaskResult',
    'GetCloudDirectorBackupIamResult',
    'GetCloudDirectorBackupTargetSpecResult',
    'GetCloudDirectorBackupTargetSpecOfferResult',
    'GetCloudDirectorOrganizationCurrentStateResult',
    'GetCloudDirectorOrganizationCurrentTaskResult',
    'GetCloudDirectorOrganizationIamResult',
    'GetCloudDirectorOrganizationTargetSpecResult',
]

@pulumi.output_type
class GetCloudDirectorBackupCurrentStateResult(dict):
    def __init__(__self__, *,
                 az_name: builtins.str,
                 offers: Sequence['outputs.GetCloudDirectorBackupCurrentStateOfferResult']):
        """
        :param builtins.str az_name: Availability zone of VMware Cloud Director organization backup
        :param Sequence['GetCloudDirectorBackupCurrentStateOfferArgs'] offers: List of your VMware Cloud Director organization backup offers
        """
        pulumi.set(__self__, "az_name", az_name)
        pulumi.set(__self__, "offers", offers)

    @property
    @pulumi.getter(name="azName")
    def az_name(self) -> builtins.str:
        """
        Availability zone of VMware Cloud Director organization backup
        """
        return pulumi.get(self, "az_name")

    @property
    @pulumi.getter
    def offers(self) -> Sequence['outputs.GetCloudDirectorBackupCurrentStateOfferResult']:
        """
        List of your VMware Cloud Director organization backup offers
        """
        return pulumi.get(self, "offers")


@pulumi.output_type
class GetCloudDirectorBackupCurrentStateOfferResult(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 protection_primary_region: builtins.str,
                 protection_replicated_region: builtins.str,
                 quota_in_tb: builtins.float,
                 status: builtins.str,
                 used_space_in_gb: builtins.float):
        """
        :param builtins.str name: Backup service offer type (BRONZE|SILVER|GOLD)
        :param builtins.str protection_primary_region: Backup repository primary region
        :param builtins.str protection_replicated_region: Backup repository replicated region
        :param builtins.float quota_in_tb: Backup repository quota in TB
        :param builtins.str status: Backup offer status
        :param builtins.float used_space_in_gb: Backup repository used space in GB
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "protection_primary_region", protection_primary_region)
        pulumi.set(__self__, "protection_replicated_region", protection_replicated_region)
        pulumi.set(__self__, "quota_in_tb", quota_in_tb)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "used_space_in_gb", used_space_in_gb)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Backup service offer type (BRONZE|SILVER|GOLD)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectionPrimaryRegion")
    def protection_primary_region(self) -> builtins.str:
        """
        Backup repository primary region
        """
        return pulumi.get(self, "protection_primary_region")

    @property
    @pulumi.getter(name="protectionReplicatedRegion")
    def protection_replicated_region(self) -> builtins.str:
        """
        Backup repository replicated region
        """
        return pulumi.get(self, "protection_replicated_region")

    @property
    @pulumi.getter(name="quotaInTb")
    def quota_in_tb(self) -> builtins.float:
        """
        Backup repository quota in TB
        """
        return pulumi.get(self, "quota_in_tb")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Backup offer status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="usedSpaceInGb")
    def used_space_in_gb(self) -> builtins.float:
        """
        Backup repository used space in GB
        """
        return pulumi.get(self, "used_space_in_gb")


@pulumi.output_type
class GetCloudDirectorBackupCurrentTaskResult(dict):
    def __init__(__self__, *,
                 id: builtins.str,
                 link: builtins.str,
                 status: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str id: Identifier of the current task
        :param builtins.str link: Link to the task details
        :param builtins.str status: Current global status of the current task
        :param builtins.str type: Type of the current task
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "link", link)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Identifier of the current task
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def link(self) -> builtins.str:
        """
        Link to the task details
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Current global status of the current task
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of the current task
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetCloudDirectorBackupIamResult(dict):
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
class GetCloudDirectorBackupTargetSpecResult(dict):
    def __init__(__self__, *,
                 offers: Sequence['outputs.GetCloudDirectorBackupTargetSpecOfferResult']):
        """
        :param Sequence['GetCloudDirectorBackupTargetSpecOfferArgs'] offers: List of your VMware Cloud Director backup offers
        """
        pulumi.set(__self__, "offers", offers)

    @property
    @pulumi.getter
    def offers(self) -> Sequence['outputs.GetCloudDirectorBackupTargetSpecOfferResult']:
        """
        List of your VMware Cloud Director backup offers
        """
        return pulumi.get(self, "offers")


@pulumi.output_type
class GetCloudDirectorBackupTargetSpecOfferResult(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 quota_in_tb: builtins.float):
        """
        :param builtins.str name: Backup service offer type (BRONZE|SILVER|GOLD)
        :param builtins.float quota_in_tb: Backup repository quota in TB
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "quota_in_tb", quota_in_tb)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Backup service offer type (BRONZE|SILVER|GOLD)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="quotaInTb")
    def quota_in_tb(self) -> builtins.float:
        """
        Backup repository quota in TB
        """
        return pulumi.get(self, "quota_in_tb")


@pulumi.output_type
class GetCloudDirectorOrganizationCurrentStateResult(dict):
    def __init__(__self__, *,
                 api_url: builtins.str,
                 billing_type: builtins.str,
                 description: builtins.str,
                 full_name: builtins.str,
                 name: builtins.str,
                 region: builtins.str,
                 spla: builtins.bool,
                 web_interface_url: builtins.str):
        """
        :param builtins.str api_url: API URL to interact with your VMware Cloud Director organization at OVHcloud
        :param builtins.str billing_type: Billing type of your VMware Cloud Director project
        :param builtins.str description: Description of your VMware Cloud Director organization on OVHcloud
        :param builtins.str full_name: Human readable full name of your VMware Cloud Director organization
        :param builtins.str name: Name of your VMware Cloud Director organization
        :param builtins.str region: Datacenter where your VMware Cloud Director organization is physically located
        :param builtins.bool spla: SPLA licensing state
        :param builtins.str web_interface_url: URL to administrate your VMware Cloud Director organization at OVHcloud
        """
        pulumi.set(__self__, "api_url", api_url)
        pulumi.set(__self__, "billing_type", billing_type)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "full_name", full_name)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "spla", spla)
        pulumi.set(__self__, "web_interface_url", web_interface_url)

    @property
    @pulumi.getter(name="apiUrl")
    def api_url(self) -> builtins.str:
        """
        API URL to interact with your VMware Cloud Director organization at OVHcloud
        """
        return pulumi.get(self, "api_url")

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> builtins.str:
        """
        Billing type of your VMware Cloud Director project
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description of your VMware Cloud Director organization on OVHcloud
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> builtins.str:
        """
        Human readable full name of your VMware Cloud Director organization
        """
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of your VMware Cloud Director organization
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> builtins.str:
        """
        Datacenter where your VMware Cloud Director organization is physically located
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def spla(self) -> builtins.bool:
        """
        SPLA licensing state
        """
        return pulumi.get(self, "spla")

    @property
    @pulumi.getter(name="webInterfaceUrl")
    def web_interface_url(self) -> builtins.str:
        """
        URL to administrate your VMware Cloud Director organization at OVHcloud
        """
        return pulumi.get(self, "web_interface_url")


@pulumi.output_type
class GetCloudDirectorOrganizationCurrentTaskResult(dict):
    def __init__(__self__, *,
                 id: builtins.str,
                 link: builtins.str,
                 status: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str id: Identifier of the current task
        :param builtins.str link: Link to the task details
        :param builtins.str status: Current global status of the current task
        :param builtins.str type: Type of the current task
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "link", link)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Identifier of the current task
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def link(self) -> builtins.str:
        """
        Link to the task details
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Current global status of the current task
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Type of the current task
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetCloudDirectorOrganizationIamResult(dict):
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
class GetCloudDirectorOrganizationTargetSpecResult(dict):
    def __init__(__self__, *,
                 description: builtins.str,
                 full_name: builtins.str):
        """
        :param builtins.str description: Description of your VMware Cloud Director organization at OVHcloud
        :param builtins.str full_name: Human readable full name of your VMware Cloud Director organization
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "full_name", full_name)

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description of your VMware Cloud Director organization at OVHcloud
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> builtins.str:
        """
        Human readable full name of your VMware Cloud Director organization
        """
        return pulumi.get(self, "full_name")


