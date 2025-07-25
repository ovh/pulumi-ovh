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

__all__ = [
    'IpServiceOrderArgs',
    'IpServiceOrderArgsDict',
    'IpServiceOrderDetailArgs',
    'IpServiceOrderDetailArgsDict',
    'IpServicePlanArgs',
    'IpServicePlanArgsDict',
    'IpServicePlanConfigurationArgs',
    'IpServicePlanConfigurationArgsDict',
    'IpServicePlanOptionArgs',
    'IpServicePlanOptionArgsDict',
    'IpServicePlanOptionConfigurationArgs',
    'IpServicePlanOptionConfigurationArgsDict',
    'IpServiceRoutedToArgs',
    'IpServiceRoutedToArgsDict',
    'MoveRoutedToArgs',
    'MoveRoutedToArgsDict',
]

MYPY = False

if not MYPY:
    class IpServiceOrderArgsDict(TypedDict):
        date: NotRequired[pulumi.Input[builtins.str]]
        """
        date
        """
        details: NotRequired[pulumi.Input[Sequence[pulumi.Input['IpServiceOrderDetailArgsDict']]]]
        """
        Information about a Bill entry
        """
        expiration_date: NotRequired[pulumi.Input[builtins.str]]
        """
        expiration date
        """
        order_id: NotRequired[pulumi.Input[builtins.int]]
        """
        order id
        """
elif False:
    IpServiceOrderArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpServiceOrderArgs:
    def __init__(__self__, *,
                 date: Optional[pulumi.Input[builtins.str]] = None,
                 details: Optional[pulumi.Input[Sequence[pulumi.Input['IpServiceOrderDetailArgs']]]] = None,
                 expiration_date: Optional[pulumi.Input[builtins.str]] = None,
                 order_id: Optional[pulumi.Input[builtins.int]] = None):
        """
        :param pulumi.Input[builtins.str] date: date
        :param pulumi.Input[Sequence[pulumi.Input['IpServiceOrderDetailArgs']]] details: Information about a Bill entry
        :param pulumi.Input[builtins.str] expiration_date: expiration date
        :param pulumi.Input[builtins.int] order_id: order id
        """
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
    def date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        date
        """
        return pulumi.get(self, "date")

    @date.setter
    def date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "date", value)

    @property
    @pulumi.getter
    def details(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpServiceOrderDetailArgs']]]]:
        """
        Information about a Bill entry
        """
        return pulumi.get(self, "details")

    @details.setter
    def details(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpServiceOrderDetailArgs']]]]):
        pulumi.set(self, "details", value)

    @property
    @pulumi.getter(name="expirationDate")
    def expiration_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        expiration date
        """
        return pulumi.get(self, "expiration_date")

    @expiration_date.setter
    def expiration_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "expiration_date", value)

    @property
    @pulumi.getter(name="orderId")
    def order_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        order id
        """
        return pulumi.get(self, "order_id")

    @order_id.setter
    def order_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "order_id", value)


if not MYPY:
    class IpServiceOrderDetailArgsDict(TypedDict):
        description: NotRequired[pulumi.Input[builtins.str]]
        """
        Custom description on your ip.
        """
        domain: NotRequired[pulumi.Input[builtins.str]]
        """
        expiration date
        """
        order_detail_id: NotRequired[pulumi.Input[builtins.int]]
        """
        order detail id
        """
        quantity: NotRequired[pulumi.Input[builtins.str]]
        """
        quantity
        """
elif False:
    IpServiceOrderDetailArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpServiceOrderDetailArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain: Optional[pulumi.Input[builtins.str]] = None,
                 order_detail_id: Optional[pulumi.Input[builtins.int]] = None,
                 quantity: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] description: Custom description on your ip.
        :param pulumi.Input[builtins.str] domain: expiration date
        :param pulumi.Input[builtins.int] order_detail_id: order detail id
        :param pulumi.Input[builtins.str] quantity: quantity
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if order_detail_id is not None:
            pulumi.set(__self__, "order_detail_id", order_detail_id)
        if quantity is not None:
            pulumi.set(__self__, "quantity", quantity)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Custom description on your ip.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        expiration date
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="orderDetailId")
    def order_detail_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        order detail id
        """
        return pulumi.get(self, "order_detail_id")

    @order_detail_id.setter
    def order_detail_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "order_detail_id", value)

    @property
    @pulumi.getter
    def quantity(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        quantity
        """
        return pulumi.get(self, "quantity")

    @quantity.setter
    def quantity(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "quantity", value)


if not MYPY:
    class IpServicePlanArgsDict(TypedDict):
        duration: pulumi.Input[builtins.str]
        """
        duration
        """
        plan_code: pulumi.Input[builtins.str]
        """
        Plan code
        """
        pricing_mode: pulumi.Input[builtins.str]
        """
        Pricing model identifier
        """
        catalog_name: NotRequired[pulumi.Input[builtins.str]]
        """
        Catalog name
        """
        configurations: NotRequired[pulumi.Input[Sequence[pulumi.Input['IpServicePlanConfigurationArgsDict']]]]
        """
        Representation of a configuration item for personalizing product
        """
elif False:
    IpServicePlanArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpServicePlanArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[builtins.str],
                 plan_code: pulumi.Input[builtins.str],
                 pricing_mode: pulumi.Input[builtins.str],
                 catalog_name: Optional[pulumi.Input[builtins.str]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input['IpServicePlanConfigurationArgs']]]] = None):
        """
        :param pulumi.Input[builtins.str] duration: duration
        :param pulumi.Input[builtins.str] plan_code: Plan code
        :param pulumi.Input[builtins.str] pricing_mode: Pricing model identifier
        :param pulumi.Input[builtins.str] catalog_name: Catalog name
        :param pulumi.Input[Sequence[pulumi.Input['IpServicePlanConfigurationArgs']]] configurations: Representation of a configuration item for personalizing product
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "plan_code", plan_code)
        pulumi.set(__self__, "pricing_mode", pricing_mode)
        if catalog_name is not None:
            pulumi.set(__self__, "catalog_name", catalog_name)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[builtins.str]:
        """
        duration
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> pulumi.Input[builtins.str]:
        """
        Plan code
        """
        return pulumi.get(self, "plan_code")

    @plan_code.setter
    def plan_code(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "plan_code", value)

    @property
    @pulumi.getter(name="pricingMode")
    def pricing_mode(self) -> pulumi.Input[builtins.str]:
        """
        Pricing model identifier
        """
        return pulumi.get(self, "pricing_mode")

    @pricing_mode.setter
    def pricing_mode(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "pricing_mode", value)

    @property
    @pulumi.getter(name="catalogName")
    def catalog_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Catalog name
        """
        return pulumi.get(self, "catalog_name")

    @catalog_name.setter
    def catalog_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "catalog_name", value)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpServicePlanConfigurationArgs']]]]:
        """
        Representation of a configuration item for personalizing product
        """
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpServicePlanConfigurationArgs']]]]):
        pulumi.set(self, "configurations", value)


if not MYPY:
    class IpServicePlanConfigurationArgsDict(TypedDict):
        label: pulumi.Input[builtins.str]
        """
        Identifier of the resource
        """
        value: pulumi.Input[builtins.str]
        """
        Path to the resource in API.OVH.COM
        """
elif False:
    IpServicePlanConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpServicePlanConfigurationArgs:
    def __init__(__self__, *,
                 label: pulumi.Input[builtins.str],
                 value: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] label: Identifier of the resource
        :param pulumi.Input[builtins.str] value: Path to the resource in API.OVH.COM
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[builtins.str]:
        """
        Identifier of the resource
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[builtins.str]:
        """
        Path to the resource in API.OVH.COM
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "value", value)


if not MYPY:
    class IpServicePlanOptionArgsDict(TypedDict):
        duration: pulumi.Input[builtins.str]
        """
        duration
        """
        plan_code: pulumi.Input[builtins.str]
        """
        Plan code
        """
        pricing_mode: pulumi.Input[builtins.str]
        """
        Pricing model identifier
        """
        catalog_name: NotRequired[pulumi.Input[builtins.str]]
        """
        Catalog name
        """
        configurations: NotRequired[pulumi.Input[Sequence[pulumi.Input['IpServicePlanOptionConfigurationArgsDict']]]]
        """
        Representation of a configuration item for personalizing product. The list of available configurations can be retrieved using call [GET /order/cart/{cartId}/item/{itemId}/requiredConfiguration](https://eu.api.ovh.com/console/?section=%2Forder&branch=v1#get-/order/cart/-cartId-/item/-itemId-/requiredConfiguration)
        """
elif False:
    IpServicePlanOptionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpServicePlanOptionArgs:
    def __init__(__self__, *,
                 duration: pulumi.Input[builtins.str],
                 plan_code: pulumi.Input[builtins.str],
                 pricing_mode: pulumi.Input[builtins.str],
                 catalog_name: Optional[pulumi.Input[builtins.str]] = None,
                 configurations: Optional[pulumi.Input[Sequence[pulumi.Input['IpServicePlanOptionConfigurationArgs']]]] = None):
        """
        :param pulumi.Input[builtins.str] duration: duration
        :param pulumi.Input[builtins.str] plan_code: Plan code
        :param pulumi.Input[builtins.str] pricing_mode: Pricing model identifier
        :param pulumi.Input[builtins.str] catalog_name: Catalog name
        :param pulumi.Input[Sequence[pulumi.Input['IpServicePlanOptionConfigurationArgs']]] configurations: Representation of a configuration item for personalizing product. The list of available configurations can be retrieved using call [GET /order/cart/{cartId}/item/{itemId}/requiredConfiguration](https://eu.api.ovh.com/console/?section=%2Forder&branch=v1#get-/order/cart/-cartId-/item/-itemId-/requiredConfiguration)
        """
        pulumi.set(__self__, "duration", duration)
        pulumi.set(__self__, "plan_code", plan_code)
        pulumi.set(__self__, "pricing_mode", pricing_mode)
        if catalog_name is not None:
            pulumi.set(__self__, "catalog_name", catalog_name)
        if configurations is not None:
            pulumi.set(__self__, "configurations", configurations)

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Input[builtins.str]:
        """
        duration
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> pulumi.Input[builtins.str]:
        """
        Plan code
        """
        return pulumi.get(self, "plan_code")

    @plan_code.setter
    def plan_code(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "plan_code", value)

    @property
    @pulumi.getter(name="pricingMode")
    def pricing_mode(self) -> pulumi.Input[builtins.str]:
        """
        Pricing model identifier
        """
        return pulumi.get(self, "pricing_mode")

    @pricing_mode.setter
    def pricing_mode(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "pricing_mode", value)

    @property
    @pulumi.getter(name="catalogName")
    def catalog_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Catalog name
        """
        return pulumi.get(self, "catalog_name")

    @catalog_name.setter
    def catalog_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "catalog_name", value)

    @property
    @pulumi.getter
    def configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IpServicePlanOptionConfigurationArgs']]]]:
        """
        Representation of a configuration item for personalizing product. The list of available configurations can be retrieved using call [GET /order/cart/{cartId}/item/{itemId}/requiredConfiguration](https://eu.api.ovh.com/console/?section=%2Forder&branch=v1#get-/order/cart/-cartId-/item/-itemId-/requiredConfiguration)
        """
        return pulumi.get(self, "configurations")

    @configurations.setter
    def configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IpServicePlanOptionConfigurationArgs']]]]):
        pulumi.set(self, "configurations", value)


if not MYPY:
    class IpServicePlanOptionConfigurationArgsDict(TypedDict):
        label: pulumi.Input[builtins.str]
        """
        Identifier of the resource
        """
        value: pulumi.Input[builtins.str]
        """
        Path to the resource in API.OVH.COM
        """
elif False:
    IpServicePlanOptionConfigurationArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpServicePlanOptionConfigurationArgs:
    def __init__(__self__, *,
                 label: pulumi.Input[builtins.str],
                 value: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] label: Identifier of the resource
        :param pulumi.Input[builtins.str] value: Path to the resource in API.OVH.COM
        """
        pulumi.set(__self__, "label", label)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def label(self) -> pulumi.Input[builtins.str]:
        """
        Identifier of the resource
        """
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[builtins.str]:
        """
        Path to the resource in API.OVH.COM
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "value", value)


if not MYPY:
    class IpServiceRoutedToArgsDict(TypedDict):
        service_name: NotRequired[pulumi.Input[builtins.str]]
        """
        service name
        """
elif False:
    IpServiceRoutedToArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class IpServiceRoutedToArgs:
    def __init__(__self__, *,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] service_name: service name
        """
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


if not MYPY:
    class MoveRoutedToArgsDict(TypedDict):
        service_name: pulumi.Input[builtins.str]
        """
        Name of the service to route the IP to. IP will be parked if this value is an empty string
        """
elif False:
    MoveRoutedToArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class MoveRoutedToArgs:
    def __init__(__self__, *,
                 service_name: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] service_name: Name of the service to route the IP to. IP will be parked if this value is an empty string
        """
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        Name of the service to route the IP to. IP will be parked if this value is an empty string
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)


