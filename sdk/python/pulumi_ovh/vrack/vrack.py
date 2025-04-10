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
from ._inputs import *

__all__ = ['VrackArgs', 'Vrack']

@pulumi.input_type
class VrackArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[builtins.str]] = None,
                 payment_mean: Optional[pulumi.Input[builtins.str]] = None,
                 plan: Optional[pulumi.Input['VrackPlanArgs']] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]] = None):
        """
        The set of arguments for constructing a Vrack resource.
        :param pulumi.Input[builtins.str] description: yourvrackdescription
        :param pulumi.Input[builtins.str] name: yourvrackname
        :param pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]] orders: Details about an Order
        :param pulumi.Input[builtins.str] ovh_subsidiary: OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        :param pulumi.Input[builtins.str] payment_mean: Ovh payment mode
        :param pulumi.Input['VrackPlanArgs'] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]] plan_options: Product Plan to order
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if orders is not None:
            pulumi.set(__self__, "orders", orders)
        if ovh_subsidiary is not None:
            pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        yourvrackdescription
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        yourvrackname
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def orders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @orders.setter
    def orders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]]):
        pulumi.set(self, "orders", value)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['VrackPlanArgs']]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['VrackPlanArgs']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)


@pulumi.input_type
class _VrackState:
    def __init__(__self__, *,
                 vrack_urn: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[builtins.str]] = None,
                 payment_mean: Optional[pulumi.Input[builtins.str]] = None,
                 plan: Optional[pulumi.Input['VrackPlanArgs']] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Vrack resources.
        :param pulumi.Input[builtins.str] vrack_urn: The URN of the vrack, used with IAM permissions
        :param pulumi.Input[builtins.str] description: yourvrackdescription
        :param pulumi.Input[builtins.str] name: yourvrackname
        :param pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]] orders: Details about an Order
        :param pulumi.Input[builtins.str] ovh_subsidiary: OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        :param pulumi.Input[builtins.str] payment_mean: Ovh payment mode
        :param pulumi.Input['VrackPlanArgs'] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]] plan_options: Product Plan to order
        :param pulumi.Input[builtins.str] service_name: The internal name of your vrack
        """
        if vrack_urn is not None:
            pulumi.set(__self__, "vrack_urn", vrack_urn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if orders is not None:
            pulumi.set(__self__, "orders", orders)
        if ovh_subsidiary is not None:
            pulumi.set(__self__, "ovh_subsidiary", ovh_subsidiary)
        if payment_mean is not None:
            warnings.warn("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""", DeprecationWarning)
            pulumi.log.warn("""payment_mean is deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
        if payment_mean is not None:
            pulumi.set(__self__, "payment_mean", payment_mean)
        if plan is not None:
            pulumi.set(__self__, "plan", plan)
        if plan_options is not None:
            pulumi.set(__self__, "plan_options", plan_options)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="VrackURN")
    def vrack_urn(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The URN of the vrack, used with IAM permissions
        """
        return pulumi.get(self, "vrack_urn")

    @vrack_urn.setter
    def vrack_urn(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vrack_urn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        yourvrackdescription
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        yourvrackname
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def orders(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @orders.setter
    def orders(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VrackOrderArgs']]]]):
        pulumi.set(self, "orders", value)

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        """
        return pulumi.get(self, "ovh_subsidiary")

    @ovh_subsidiary.setter
    def ovh_subsidiary(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ovh_subsidiary", value)

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @payment_mean.setter
    def payment_mean(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "payment_mean", value)

    @property
    @pulumi.getter
    def plan(self) -> Optional[pulumi.Input['VrackPlanArgs']]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @plan.setter
    def plan(self, value: Optional[pulumi.Input['VrackPlanArgs']]):
        pulumi.set(self, "plan", value)

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @plan_options.setter
    def plan_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VrackPlanOptionArgs']]]]):
        pulumi.set(self, "plan_options", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)


class Vrack(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VrackOrderArgs', 'VrackOrderArgsDict']]]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[builtins.str]] = None,
                 payment_mean: Optional[pulumi.Input[builtins.str]] = None,
                 plan: Optional[pulumi.Input[Union['VrackPlanArgs', 'VrackPlanArgsDict']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VrackPlanOptionArgs', 'VrackPlanOptionArgsDict']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_account = ovh.Me.get_me()
        my_cart = ovh.Order.get_cart(ovh_subsidiary=my_account.ovh_subsidiary)
        vrack_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=my_cart.id,
            price_capacity="renew",
            product="vrack",
            plan_code="vrack")
        vrack_vrack = ovh.vrack.Vrack("vrackVrack",
            ovh_subsidiary=my_cart.ovh_subsidiary,
            description="my vrack",
            plan={
                "duration": vrack_cart_product_plan.selected_prices[0].duration,
                "plan_code": vrack_cart_product_plan.plan_code,
                "pricing_mode": vrack_cart_product_plan.selected_prices[0].pricing_mode,
            })
        ```

        ## Import

        A vRack can be imported using the `service_name`. Using the following configuration:

        terraform

        import {

          to = ovh_vrack.vrack

          id = "<service name>"

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=vrack.tf

        $ pulumi up

        The file `vrack.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: yourvrackdescription
        :param pulumi.Input[builtins.str] name: yourvrackname
        :param pulumi.Input[Sequence[pulumi.Input[Union['VrackOrderArgs', 'VrackOrderArgsDict']]]] orders: Details about an Order
        :param pulumi.Input[builtins.str] ovh_subsidiary: OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        :param pulumi.Input[builtins.str] payment_mean: Ovh payment mode
        :param pulumi.Input[Union['VrackPlanArgs', 'VrackPlanArgsDict']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[Union['VrackPlanOptionArgs', 'VrackPlanOptionArgsDict']]]] plan_options: Product Plan to order
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[VrackArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        my_account = ovh.Me.get_me()
        my_cart = ovh.Order.get_cart(ovh_subsidiary=my_account.ovh_subsidiary)
        vrack_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=my_cart.id,
            price_capacity="renew",
            product="vrack",
            plan_code="vrack")
        vrack_vrack = ovh.vrack.Vrack("vrackVrack",
            ovh_subsidiary=my_cart.ovh_subsidiary,
            description="my vrack",
            plan={
                "duration": vrack_cart_product_plan.selected_prices[0].duration,
                "plan_code": vrack_cart_product_plan.plan_code,
                "pricing_mode": vrack_cart_product_plan.selected_prices[0].pricing_mode,
            })
        ```

        ## Import

        A vRack can be imported using the `service_name`. Using the following configuration:

        terraform

        import {

          to = ovh_vrack.vrack

          id = "<service name>"

        }

        You can then run:

        bash

        $ pulumi preview -generate-config-out=vrack.tf

        $ pulumi up

        The file `vrack.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.

        :param str resource_name: The name of the resource.
        :param VrackArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VrackArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VrackOrderArgs', 'VrackOrderArgsDict']]]]] = None,
                 ovh_subsidiary: Optional[pulumi.Input[builtins.str]] = None,
                 payment_mean: Optional[pulumi.Input[builtins.str]] = None,
                 plan: Optional[pulumi.Input[Union['VrackPlanArgs', 'VrackPlanArgsDict']]] = None,
                 plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VrackPlanOptionArgs', 'VrackPlanOptionArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VrackArgs.__new__(VrackArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["orders"] = orders
            __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
            __props__.__dict__["payment_mean"] = payment_mean
            __props__.__dict__["plan"] = plan
            __props__.__dict__["plan_options"] = plan_options
            __props__.__dict__["vrack_urn"] = None
            __props__.__dict__["service_name"] = None
        super(Vrack, __self__).__init__(
            'ovh:Vrack/vrack:Vrack',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vrack_urn: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            orders: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VrackOrderArgs', 'VrackOrderArgsDict']]]]] = None,
            ovh_subsidiary: Optional[pulumi.Input[builtins.str]] = None,
            payment_mean: Optional[pulumi.Input[builtins.str]] = None,
            plan: Optional[pulumi.Input[Union['VrackPlanArgs', 'VrackPlanArgsDict']]] = None,
            plan_options: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VrackPlanOptionArgs', 'VrackPlanOptionArgsDict']]]]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None) -> 'Vrack':
        """
        Get an existing Vrack resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] vrack_urn: The URN of the vrack, used with IAM permissions
        :param pulumi.Input[builtins.str] description: yourvrackdescription
        :param pulumi.Input[builtins.str] name: yourvrackname
        :param pulumi.Input[Sequence[pulumi.Input[Union['VrackOrderArgs', 'VrackOrderArgsDict']]]] orders: Details about an Order
        :param pulumi.Input[builtins.str] ovh_subsidiary: OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        :param pulumi.Input[builtins.str] payment_mean: Ovh payment mode
        :param pulumi.Input[Union['VrackPlanArgs', 'VrackPlanArgsDict']] plan: Product Plan to order
        :param pulumi.Input[Sequence[pulumi.Input[Union['VrackPlanOptionArgs', 'VrackPlanOptionArgsDict']]]] plan_options: Product Plan to order
        :param pulumi.Input[builtins.str] service_name: The internal name of your vrack
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VrackState.__new__(_VrackState)

        __props__.__dict__["vrack_urn"] = vrack_urn
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["orders"] = orders
        __props__.__dict__["ovh_subsidiary"] = ovh_subsidiary
        __props__.__dict__["payment_mean"] = payment_mean
        __props__.__dict__["plan"] = plan
        __props__.__dict__["plan_options"] = plan_options
        __props__.__dict__["service_name"] = service_name
        return Vrack(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="VrackURN")
    def vrack_urn(self) -> pulumi.Output[builtins.str]:
        """
        The URN of the vrack, used with IAM permissions
        """
        return pulumi.get(self, "vrack_urn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[builtins.str]:
        """
        yourvrackdescription
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        yourvrackname
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def orders(self) -> pulumi.Output[Sequence['outputs.VrackOrder']]:
        """
        Details about an Order
        """
        return pulumi.get(self, "orders")

    @property
    @pulumi.getter(name="ovhSubsidiary")
    def ovh_subsidiary(self) -> pulumi.Output[builtins.str]:
        """
        OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        """
        return pulumi.get(self, "ovh_subsidiary")

    @property
    @pulumi.getter(name="paymentMean")
    @_utilities.deprecated("""This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.""")
    def payment_mean(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Ovh payment mode
        """
        return pulumi.get(self, "payment_mean")

    @property
    @pulumi.getter
    def plan(self) -> pulumi.Output['outputs.VrackPlan']:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan")

    @property
    @pulumi.getter(name="planOptions")
    def plan_options(self) -> pulumi.Output[Optional[Sequence['outputs.VrackPlanOption']]]:
        """
        Product Plan to order
        """
        return pulumi.get(self, "plan_options")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

