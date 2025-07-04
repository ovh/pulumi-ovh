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
from .. import _utilities
from . import outputs

__all__ = [
    'VpsIam',
    'VpsModel',
    'VpsOrder',
    'VpsOrderDetail',
    'VpsPlan',
    'VpsPlanConfiguration',
    'VpsPlanOption',
    'VpsPlanOptionConfiguration',
]

@pulumi.output_type
class VpsIam(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "displayName":
            suggest = "display_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpsIam. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpsIam.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpsIam.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 display_name: Optional[builtins.str] = None,
                 id: Optional[builtins.str] = None,
                 tags: Optional[Mapping[str, builtins.str]] = None,
                 urn: Optional[builtins.str] = None):
        """
        :param builtins.str display_name: Custom display name
        :param builtins.str id: Unique identifier of the resource in the IAM
        :param Mapping[str, builtins.str] tags: Resource tags. Tags that were internally computed are prefixed with `ovh:`
        :param builtins.str urn: URN of the private database, used when writing IAM policies
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if urn is not None:
            pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[builtins.str]:
        """
        Custom display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> Optional[builtins.str]:
        """
        Unique identifier of the resource in the IAM
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, builtins.str]]:
        """
        Resource tags. Tags that were internally computed are prefixed with `ovh:`
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def urn(self) -> Optional[builtins.str]:
        """
        URN of the private database, used when writing IAM policies
        """
        return pulumi.get(self, "urn")


@pulumi.output_type
class VpsModel(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "availableOptions":
            suggest = "available_options"
        elif key == "maximumAdditionnalIp":
            suggest = "maximum_additionnal_ip"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpsModel. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpsModel.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpsModel.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 available_options: Optional[Sequence[builtins.str]] = None,
                 datacenters: Optional[Sequence[builtins.str]] = None,
                 disk: Optional[builtins.float] = None,
                 maximum_additionnal_ip: Optional[builtins.float] = None,
                 memory: Optional[builtins.float] = None,
                 name: Optional[builtins.str] = None,
                 offer: Optional[builtins.str] = None,
                 vcore: Optional[builtins.float] = None,
                 version: Optional[builtins.str] = None):
        """
        :param Sequence[builtins.str] available_options: All options the VPS can have (additionalDisk┃automatedBackup┃cpanel┃ftpbackup┃plesk┃snapshot┃veeam┃windows)
        :param Sequence[builtins.str] datacenters: Datacenters where this model is available
        :param builtins.float disk: Disk capacity of this VPS
        :param builtins.float maximum_additionnal_ip: Maximum number of additional IPs
        :param builtins.float memory: RAM of the VPS
        :param builtins.str name: Name of the VPS
        :param builtins.str offer: Description of this VPS offer
        :param builtins.float vcore: Number of vcores
        :param builtins.str version: All versions that VPS can have (2013v1┃2014v1┃2015v1┃2017v1┃2017v2┃2017v3┃2018v1┃2018v2┃2019v1)
        """
        if available_options is not None:
            pulumi.set(__self__, "available_options", available_options)
        if datacenters is not None:
            pulumi.set(__self__, "datacenters", datacenters)
        if disk is not None:
            pulumi.set(__self__, "disk", disk)
        if maximum_additionnal_ip is not None:
            pulumi.set(__self__, "maximum_additionnal_ip", maximum_additionnal_ip)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if offer is not None:
            pulumi.set(__self__, "offer", offer)
        if vcore is not None:
            pulumi.set(__self__, "vcore", vcore)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="availableOptions")
    def available_options(self) -> Optional[Sequence[builtins.str]]:
        """
        All options the VPS can have (additionalDisk┃automatedBackup┃cpanel┃ftpbackup┃plesk┃snapshot┃veeam┃windows)
        """
        return pulumi.get(self, "available_options")

    @property
    @pulumi.getter
    def datacenters(self) -> Optional[Sequence[builtins.str]]:
        """
        Datacenters where this model is available
        """
        return pulumi.get(self, "datacenters")

    @property
    @pulumi.getter
    def disk(self) -> Optional[builtins.float]:
        """
        Disk capacity of this VPS
        """
        return pulumi.get(self, "disk")

    @property
    @pulumi.getter(name="maximumAdditionnalIp")
    def maximum_additionnal_ip(self) -> Optional[builtins.float]:
        """
        Maximum number of additional IPs
        """
        return pulumi.get(self, "maximum_additionnal_ip")

    @property
    @pulumi.getter
    def memory(self) -> Optional[builtins.float]:
        """
        RAM of the VPS
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        Name of the VPS
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def offer(self) -> Optional[builtins.str]:
        """
        Description of this VPS offer
        """
        return pulumi.get(self, "offer")

    @property
    @pulumi.getter
    def vcore(self) -> Optional[builtins.float]:
        """
        Number of vcores
        """
        return pulumi.get(self, "vcore")

    @property
    @pulumi.getter
    def version(self) -> Optional[builtins.str]:
        """
        All versions that VPS can have (2013v1┃2014v1┃2015v1┃2017v1┃2017v2┃2017v3┃2018v1┃2018v2┃2019v1)
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class VpsOrder(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "expirationDate":
            suggest = "expiration_date"
        elif key == "orderId":
            suggest = "order_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpsOrder. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpsOrder.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpsOrder.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 date: Optional[builtins.str] = None,
                 details: Optional[Sequence['outputs.VpsOrderDetail']] = None,
                 expiration_date: Optional[builtins.str] = None,
                 order_id: Optional[builtins.float] = None):
        if date is not None:
            pulumi.set(__self__, "date", date)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if expiration_date is not None:
            pulumi.set(__self__, "expiration_date", expiration_date)
        if order_id is not None:
            pulumi.set(__self__, "order_id", order_id)

    @property
    @pulumi.getter
    def date(self) -> Optional[builtins.str]:
        return pulumi.get(self, "date")

    @property
    @pulumi.getter
    def details(self) -> Optional[Sequence['outputs.VpsOrderDetail']]:
        return pulumi.get(self, "details")

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[builtins.str]:
        return pulumi.get(self, "expiration_date")

    @property
    @pulumi.getter(name="orderId")
    def order_id(self) -> Optional[builtins.float]:
        return pulumi.get(self, "order_id")


@pulumi.output_type
class VpsOrderDetail(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "detailType":
            suggest = "detail_type"
        elif key == "orderDetailId":
            suggest = "order_detail_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpsOrderDetail. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpsOrderDetail.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpsOrderDetail.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[builtins.str] = None,
                 detail_type: Optional[builtins.str] = None,
                 domain: Optional[builtins.str] = None,
                 order_detail_id: Optional[builtins.float] = None,
                 quantity: Optional[builtins.str] = None):
        """
        :param builtins.str detail_type: Product type of item in order
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if detail_type is not None:
            pulumi.set(__self__, "detail_type", detail_type)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if order_detail_id is not None:
            pulumi.set(__self__, "order_detail_id", order_detail_id)
        if quantity is not None:
            pulumi.set(__self__, "quantity", quantity)

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="detailType")
    def detail_type(self) -> Optional[builtins.str]:
        """
        Product type of item in order
        """
        return pulumi.get(self, "detail_type")

    @property
    @pulumi.getter
    def domain(self) -> Optional[builtins.str]:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="orderDetailId")
    def order_detail_id(self) -> Optional[builtins.float]:
        return pulumi.get(self, "order_detail_id")

    @property
    @pulumi.getter
    def quantity(self) -> Optional[builtins.str]:
        return pulumi.get(self, "quantity")


@pulumi.output_type
class VpsPlan(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "planCode":
            suggest = "plan_code"
        elif key == "pricingMode":
            suggest = "pricing_mode"
        elif key == "itemId":
            suggest = "item_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpsPlan. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpsPlan.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpsPlan.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 duration: builtins.str,
                 plan_code: builtins.str,
                 pricing_mode: builtins.str,
                 configurations: Optional[Sequence['outputs.VpsPlanConfiguration']] = None,
                 item_id: Optional[builtins.float] = None,
                 quantity: Optional[builtins.float] = None):
        """
        :param builtins.str duration: duration
        :param builtins.str plan_code: Plan code
        :param builtins.str pricing_mode: Pricing model identifier
        :param Sequence['VpsPlanConfigurationArgs'] configurations: Representation of a configuration item for personalizing product
        :param builtins.float item_id: Cart item to be linked
        :param builtins.float quantity: Quantity of product desired
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "plan_code", plan_code)
        pulumi.set(__self__, "pricing_mode", pricing_mode)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)
        if item_id is not None:
            pulumi.set(__self__, "item_id", item_id)
        if quantity is not None:
            pulumi.set(__self__, "quantity", quantity)

    @property
    @pulumi.getter
    def duration(self) -> builtins.str:
        """
        duration
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> builtins.str:
        """
        Plan code
        """
        return pulumi.get(self, "plan_code")

    @property
    @pulumi.getter(name="pricingMode")
    def pricing_mode(self) -> builtins.str:
        """
        Pricing model identifier
        """
        return pulumi.get(self, "pricing_mode")

    @property
    @pulumi.getter
    def configurations(self) -> Optional[Sequence['outputs.VpsPlanConfiguration']]:
        """
        Representation of a configuration item for personalizing product
        """
        return pulumi.get(self, "configurations")

    @property
    @pulumi.getter(name="itemId")
    def item_id(self) -> Optional[builtins.float]:
        """
        Cart item to be linked
        """
        return pulumi.get(self, "item_id")

    @property
    @pulumi.getter
    def quantity(self) -> Optional[builtins.float]:
        """
        Quantity of product desired
        """
        return pulumi.get(self, "quantity")


@pulumi.output_type
class VpsPlanConfiguration(dict):
    def __init__(__self__, *,
                 label: builtins.str,
                 value: builtins.str):
        """
        :param builtins.str label: Identifier of the resource
        :param builtins.str value: Path to the resource in api.ovh.com
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def label(self) -> builtins.str:
        """
        Identifier of the resource
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def value(self) -> builtins.str:
        """
        Path to the resource in api.ovh.com
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class VpsPlanOption(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "planCode":
            suggest = "plan_code"
        elif key == "pricingMode":
            suggest = "pricing_mode"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpsPlanOption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpsPlanOption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpsPlanOption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 duration: builtins.str,
                 plan_code: builtins.str,
                 pricing_mode: builtins.str,
                 quantity: builtins.float,
                 configurations: Optional[Sequence['outputs.VpsPlanOptionConfiguration']] = None):
        """
        :param builtins.str duration: duration
        :param builtins.str plan_code: Plan code
        :param builtins.str pricing_mode: Pricing model identifier
        :param builtins.float quantity: Quantity of product desired
        :param Sequence['VpsPlanOptionConfigurationArgs'] configurations: Representation of a configuration item for personalizing product
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "plan_code", plan_code)
        pulumi.set(__self__, "pricing_mode", pricing_mode)
        pulumi.set(__self__, "quantity", quantity)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)

    @property
    @pulumi.getter
    def duration(self) -> builtins.str:
        """
        duration
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> builtins.str:
        """
        Plan code
        """
        return pulumi.get(self, "plan_code")

    @property
    @pulumi.getter(name="pricingMode")
    def pricing_mode(self) -> builtins.str:
        """
        Pricing model identifier
        """
        return pulumi.get(self, "pricing_mode")

    @property
    @pulumi.getter
    def quantity(self) -> builtins.float:
        """
        Quantity of product desired
        """
        return pulumi.get(self, "quantity")

    @property
    @pulumi.getter
    def configurations(self) -> Optional[Sequence['outputs.VpsPlanOptionConfiguration']]:
        """
        Representation of a configuration item for personalizing product
        """
        return pulumi.get(self, "configurations")


@pulumi.output_type
class VpsPlanOptionConfiguration(dict):
    def __init__(__self__, *,
                 label: builtins.str,
                 value: builtins.str):
        """
        :param builtins.str label: Identifier of the resource
        :param builtins.str value: Path to the resource in api.ovh.com
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def label(self) -> builtins.str:
        """
        Identifier of the resource
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def value(self) -> builtins.str:
        """
        Path to the resource in api.ovh.com
        """
        return pulumi.get(self, "value")


