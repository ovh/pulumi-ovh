# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['SavingsPlanArgs', 'SavingsPlan']

@pulumi.input_type
class SavingsPlanArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 flavor: pulumi.Input[str],
                 period: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 size: pulumi.Input[int],
                 auto_renewal: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SavingsPlan resource.
        :param pulumi.Input[str] display_name: Custom display name, used in invoices
        :param pulumi.Input[str] flavor: Savings Plan flavor. The list of available flavors can be retrieved in the next section.
        :param pulumi.Input[str] period: Periodicity of the Savings Plan
        :param pulumi.Input[str] service_name: ID of the public cloud project
        :param pulumi.Input[int] size: Size of the Savings Plan
        :param pulumi.Input[bool] auto_renewal: Whether Savings Plan should be renewed at the end of the period (defaults to false)
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "flavor", flavor)
        pulumi.set(__self__, "period", period)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "size", size)
        if auto_renewal is not None:
            pulumi.set(__self__, "auto_renewal", auto_renewal)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        Custom display name, used in invoices
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Input[str]:
        """
        Savings Plan flavor. The list of available flavors can be retrieved in the next section.
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: pulumi.Input[str]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter
    def period(self) -> pulumi.Input[str]:
        """
        Periodicity of the Savings Plan
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: pulumi.Input[str]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        ID of the public cloud project
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        """
        Size of the Savings Plan
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="autoRenewal")
    def auto_renewal(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Savings Plan should be renewed at the end of the period (defaults to false)
        """
        return pulumi.get(self, "auto_renewal")

    @auto_renewal.setter
    def auto_renewal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renewal", value)


@pulumi.input_type
class _SavingsPlanState:
    def __init__(__self__, *,
                 auto_renewal: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 end_date: Optional[pulumi.Input[str]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 period_end_action: Optional[pulumi.Input[str]] = None,
                 period_end_date: Optional[pulumi.Input[str]] = None,
                 period_start_date: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[int]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 start_date: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SavingsPlan resources.
        :param pulumi.Input[bool] auto_renewal: Whether Savings Plan should be renewed at the end of the period (defaults to false)
        :param pulumi.Input[str] display_name: Custom display name, used in invoices
        :param pulumi.Input[str] end_date: End date of the Savings Plan
        :param pulumi.Input[str] flavor: Savings Plan flavor. The list of available flavors can be retrieved in the next section.
        :param pulumi.Input[str] period: Periodicity of the Savings Plan
        :param pulumi.Input[str] period_end_action: Action performed when reaching the end of the period (controles by the `auto_renewal` parameter)
        :param pulumi.Input[str] period_end_date: End date of the current period
        :param pulumi.Input[str] period_start_date: Start date of the current period
        :param pulumi.Input[int] service_id: Billing ID of the service
        :param pulumi.Input[str] service_name: ID of the public cloud project
        :param pulumi.Input[int] size: Size of the Savings Plan
        :param pulumi.Input[str] start_date: Start date of the Savings Plan
        :param pulumi.Input[str] status: Status of the Savings Plan
        """
        if auto_renewal is not None:
            pulumi.set(__self__, "auto_renewal", auto_renewal)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if end_date is not None:
            pulumi.set(__self__, "end_date", end_date)
        if flavor is not None:
            pulumi.set(__self__, "flavor", flavor)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_end_action is not None:
            pulumi.set(__self__, "period_end_action", period_end_action)
        if period_end_date is not None:
            pulumi.set(__self__, "period_end_date", period_end_date)
        if period_start_date is not None:
            pulumi.set(__self__, "period_start_date", period_start_date)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if start_date is not None:
            pulumi.set(__self__, "start_date", start_date)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="autoRenewal")
    def auto_renewal(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Savings Plan should be renewed at the end of the period (defaults to false)
        """
        return pulumi.get(self, "auto_renewal")

    @auto_renewal.setter
    def auto_renewal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renewal", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Custom display name, used in invoices
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> Optional[pulumi.Input[str]]:
        """
        End date of the Savings Plan
        """
        return pulumi.get(self, "end_date")

    @end_date.setter
    def end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_date", value)

    @property
    @pulumi.getter
    def flavor(self) -> Optional[pulumi.Input[str]]:
        """
        Savings Plan flavor. The list of available flavors can be retrieved in the next section.
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[str]]:
        """
        Periodicity of the Savings Plan
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodEndAction")
    def period_end_action(self) -> Optional[pulumi.Input[str]]:
        """
        Action performed when reaching the end of the period (controles by the `auto_renewal` parameter)
        """
        return pulumi.get(self, "period_end_action")

    @period_end_action.setter
    def period_end_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_end_action", value)

    @property
    @pulumi.getter(name="periodEndDate")
    def period_end_date(self) -> Optional[pulumi.Input[str]]:
        """
        End date of the current period
        """
        return pulumi.get(self, "period_end_date")

    @period_end_date.setter
    def period_end_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_end_date", value)

    @property
    @pulumi.getter(name="periodStartDate")
    def period_start_date(self) -> Optional[pulumi.Input[str]]:
        """
        Start date of the current period
        """
        return pulumi.get(self, "period_start_date")

    @period_start_date.setter
    def period_start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_start_date", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[int]]:
        """
        Billing ID of the service
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the public cloud project
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        Size of the Savings Plan
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> Optional[pulumi.Input[str]]:
        """
        Start date of the Savings Plan
        """
        return pulumi.get(self, "start_date")

    @start_date.setter
    def start_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_date", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of the Savings Plan
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class SavingsPlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renewal: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create and manage an OVHcloud Savings Plan

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        plan = ovh.savings_plan.SavingsPlan("plan",
            auto_renewal=True,
            display_name="one_month_rancher_savings_plan",
            flavor="Rancher",
            period="P1M",
            service_name="<public cloud project ID>",
            size=2)
        ```

        ## Import

        A savings plan can be imported using the following format: `service_name` and `id` of the savings plan, separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:SavingsPlan/savingsPlan:SavingsPlan plan service_name/savings_plan_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renewal: Whether Savings Plan should be renewed at the end of the period (defaults to false)
        :param pulumi.Input[str] display_name: Custom display name, used in invoices
        :param pulumi.Input[str] flavor: Savings Plan flavor. The list of available flavors can be retrieved in the next section.
        :param pulumi.Input[str] period: Periodicity of the Savings Plan
        :param pulumi.Input[str] service_name: ID of the public cloud project
        :param pulumi.Input[int] size: Size of the Savings Plan
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SavingsPlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create and manage an OVHcloud Savings Plan

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        plan = ovh.savings_plan.SavingsPlan("plan",
            auto_renewal=True,
            display_name="one_month_rancher_savings_plan",
            flavor="Rancher",
            period="P1M",
            service_name="<public cloud project ID>",
            size=2)
        ```

        ## Import

        A savings plan can be imported using the following format: `service_name` and `id` of the savings plan, separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:SavingsPlan/savingsPlan:SavingsPlan plan service_name/savings_plan_id
        ```

        :param str resource_name: The name of the resource.
        :param SavingsPlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SavingsPlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renewal: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SavingsPlanArgs.__new__(SavingsPlanArgs)

            __props__.__dict__["auto_renewal"] = auto_renewal
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if flavor is None and not opts.urn:
                raise TypeError("Missing required property 'flavor'")
            __props__.__dict__["flavor"] = flavor
            if period is None and not opts.urn:
                raise TypeError("Missing required property 'period'")
            __props__.__dict__["period"] = period
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__.__dict__["size"] = size
            __props__.__dict__["end_date"] = None
            __props__.__dict__["period_end_action"] = None
            __props__.__dict__["period_end_date"] = None
            __props__.__dict__["period_start_date"] = None
            __props__.__dict__["service_id"] = None
            __props__.__dict__["start_date"] = None
            __props__.__dict__["status"] = None
        super(SavingsPlan, __self__).__init__(
            'ovh:SavingsPlan/savingsPlan:SavingsPlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_renewal: Optional[pulumi.Input[bool]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            end_date: Optional[pulumi.Input[str]] = None,
            flavor: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[str]] = None,
            period_end_action: Optional[pulumi.Input[str]] = None,
            period_end_date: Optional[pulumi.Input[str]] = None,
            period_start_date: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[int]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            start_date: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'SavingsPlan':
        """
        Get an existing SavingsPlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renewal: Whether Savings Plan should be renewed at the end of the period (defaults to false)
        :param pulumi.Input[str] display_name: Custom display name, used in invoices
        :param pulumi.Input[str] end_date: End date of the Savings Plan
        :param pulumi.Input[str] flavor: Savings Plan flavor. The list of available flavors can be retrieved in the next section.
        :param pulumi.Input[str] period: Periodicity of the Savings Plan
        :param pulumi.Input[str] period_end_action: Action performed when reaching the end of the period (controles by the `auto_renewal` parameter)
        :param pulumi.Input[str] period_end_date: End date of the current period
        :param pulumi.Input[str] period_start_date: Start date of the current period
        :param pulumi.Input[int] service_id: Billing ID of the service
        :param pulumi.Input[str] service_name: ID of the public cloud project
        :param pulumi.Input[int] size: Size of the Savings Plan
        :param pulumi.Input[str] start_date: Start date of the Savings Plan
        :param pulumi.Input[str] status: Status of the Savings Plan
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SavingsPlanState.__new__(_SavingsPlanState)

        __props__.__dict__["auto_renewal"] = auto_renewal
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["end_date"] = end_date
        __props__.__dict__["flavor"] = flavor
        __props__.__dict__["period"] = period
        __props__.__dict__["period_end_action"] = period_end_action
        __props__.__dict__["period_end_date"] = period_end_date
        __props__.__dict__["period_start_date"] = period_start_date
        __props__.__dict__["service_id"] = service_id
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["size"] = size
        __props__.__dict__["start_date"] = start_date
        __props__.__dict__["status"] = status
        return SavingsPlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRenewal")
    def auto_renewal(self) -> pulumi.Output[bool]:
        """
        Whether Savings Plan should be renewed at the end of the period (defaults to false)
        """
        return pulumi.get(self, "auto_renewal")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Custom display name, used in invoices
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> pulumi.Output[str]:
        """
        End date of the Savings Plan
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Output[str]:
        """
        Savings Plan flavor. The list of available flavors can be retrieved in the next section.
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[str]:
        """
        Periodicity of the Savings Plan
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="periodEndAction")
    def period_end_action(self) -> pulumi.Output[str]:
        """
        Action performed when reaching the end of the period (controles by the `auto_renewal` parameter)
        """
        return pulumi.get(self, "period_end_action")

    @property
    @pulumi.getter(name="periodEndDate")
    def period_end_date(self) -> pulumi.Output[str]:
        """
        End date of the current period
        """
        return pulumi.get(self, "period_end_date")

    @property
    @pulumi.getter(name="periodStartDate")
    def period_start_date(self) -> pulumi.Output[str]:
        """
        Start date of the current period
        """
        return pulumi.get(self, "period_start_date")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[int]:
        """
        Billing ID of the service
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        ID of the public cloud project
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        Size of the Savings Plan
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> pulumi.Output[str]:
        """
        Start date of the Savings Plan
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of the Savings Plan
        """
        return pulumi.get(self, "status")

