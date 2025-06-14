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
from ._inputs import *

__all__ = ['TcpRouteArgs', 'TcpRoute']

@pulumi.input_type
class TcpRouteArgs:
    def __init__(__self__, *,
                 action: pulumi.Input['TcpRouteActionArgs'],
                 service_name: pulumi.Input[builtins.str],
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 frontend_id: Optional[pulumi.Input[builtins.int]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None):
        """
        The set of arguments for constructing a TcpRoute resource.
        :param pulumi.Input['TcpRouteActionArgs'] action: Action triggered when all rules match
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.str] display_name: Human readable name for your route, this field is for you
        :param pulumi.Input[builtins.int] frontend_id: Route traffic for this frontend
        :param pulumi.Input[builtins.int] weight: Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "service_name", service_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if frontend_id is not None:
            pulumi.set(__self__, "frontend_id", frontend_id)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input['TcpRouteActionArgs']:
        """
        Action triggered when all rules match
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input['TcpRouteActionArgs']):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Human readable name for your route, this field is for you
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Route traffic for this frontend
        """
        return pulumi.get(self, "frontend_id")

    @frontend_id.setter
    def frontend_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "frontend_id", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _TcpRouteState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input['TcpRouteActionArgs']] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 frontend_id: Optional[pulumi.Input[builtins.int]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['TcpRouteRuleArgs']]]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 status: Optional[pulumi.Input[builtins.str]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering TcpRoute resources.
        :param pulumi.Input['TcpRouteActionArgs'] action: Action triggered when all rules match
        :param pulumi.Input[builtins.str] display_name: Human readable name for your route, this field is for you
        :param pulumi.Input[builtins.int] frontend_id: Route traffic for this frontend
        :param pulumi.Input[Sequence[pulumi.Input['TcpRouteRuleArgs']]] rules: List of rules to match to trigger action
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.str] status: Route status. Routes in "ok" state are ready to operate
        :param pulumi.Input[builtins.int] weight: Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if frontend_id is not None:
            pulumi.set(__self__, "frontend_id", frontend_id)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['TcpRouteActionArgs']]:
        """
        Action triggered when all rules match
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['TcpRouteActionArgs']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Human readable name for your route, this field is for you
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Route traffic for this frontend
        """
        return pulumi.get(self, "frontend_id")

    @frontend_id.setter
    def frontend_id(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "frontend_id", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TcpRouteRuleArgs']]]]:
        """
        List of rules to match to trigger action
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TcpRouteRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Route status. Routes in "ok" state are ready to operate
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "weight", value)


@pulumi.type_token("ovh:IpLoadBalancing/tcpRoute:TcpRoute")
class TcpRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[Union['TcpRouteActionArgs', 'TcpRouteActionArgsDict']]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 frontend_id: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        Manage TCP route for a loadbalancer service

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        tcp_reject = ovh.ip_load_balancing.TcpRoute("tcp_reject",
            service_name="loadbalancer-xxxxxxxxxxxxxxxxxx",
            weight=1,
            action={
                "type": "reject",
            })
        ```

        ## Import

        TCP route can be imported using the following format `service_name` and the `id` of the route separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:IpLoadBalancing/tcpRoute:TcpRoute tcpreject service_name/route_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['TcpRouteActionArgs', 'TcpRouteActionArgsDict']] action: Action triggered when all rules match
        :param pulumi.Input[builtins.str] display_name: Human readable name for your route, this field is for you
        :param pulumi.Input[builtins.int] frontend_id: Route traffic for this frontend
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.int] weight: Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TcpRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage TCP route for a loadbalancer service

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        tcp_reject = ovh.ip_load_balancing.TcpRoute("tcp_reject",
            service_name="loadbalancer-xxxxxxxxxxxxxxxxxx",
            weight=1,
            action={
                "type": "reject",
            })
        ```

        ## Import

        TCP route can be imported using the following format `service_name` and the `id` of the route separated by "/" e.g.

        bash

        ```sh
        $ pulumi import ovh:IpLoadBalancing/tcpRoute:TcpRoute tcpreject service_name/route_id
        ```

        :param str resource_name: The name of the resource.
        :param TcpRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TcpRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[Union['TcpRouteActionArgs', 'TcpRouteActionArgsDict']]] = None,
                 display_name: Optional[pulumi.Input[builtins.str]] = None,
                 frontend_id: Optional[pulumi.Input[builtins.int]] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 weight: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TcpRouteArgs.__new__(TcpRouteArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["frontend_id"] = frontend_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["weight"] = weight
            __props__.__dict__["rules"] = None
            __props__.__dict__["status"] = None
        super(TcpRoute, __self__).__init__(
            'ovh:IpLoadBalancing/tcpRoute:TcpRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[Union['TcpRouteActionArgs', 'TcpRouteActionArgsDict']]] = None,
            display_name: Optional[pulumi.Input[builtins.str]] = None,
            frontend_id: Optional[pulumi.Input[builtins.int]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['TcpRouteRuleArgs', 'TcpRouteRuleArgsDict']]]]] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            status: Optional[pulumi.Input[builtins.str]] = None,
            weight: Optional[pulumi.Input[builtins.int]] = None) -> 'TcpRoute':
        """
        Get an existing TcpRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['TcpRouteActionArgs', 'TcpRouteActionArgsDict']] action: Action triggered when all rules match
        :param pulumi.Input[builtins.str] display_name: Human readable name for your route, this field is for you
        :param pulumi.Input[builtins.int] frontend_id: Route traffic for this frontend
        :param pulumi.Input[Sequence[pulumi.Input[Union['TcpRouteRuleArgs', 'TcpRouteRuleArgsDict']]]] rules: List of rules to match to trigger action
        :param pulumi.Input[builtins.str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[builtins.str] status: Route status. Routes in "ok" state are ready to operate
        :param pulumi.Input[builtins.int] weight: Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TcpRouteState.__new__(_TcpRouteState)

        __props__.__dict__["action"] = action
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["frontend_id"] = frontend_id
        __props__.__dict__["rules"] = rules
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        __props__.__dict__["weight"] = weight
        return TcpRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output['outputs.TcpRouteAction']:
        """
        Action triggered when all rules match
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Human readable name for your route, this field is for you
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="frontendId")
    def frontend_id(self) -> pulumi.Output[builtins.int]:
        """
        Route traffic for this frontend
        """
        return pulumi.get(self, "frontend_id")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.TcpRouteRule']]:
        """
        List of rules to match to trigger action
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[builtins.str]:
        """
        Route status. Routes in "ok" state are ready to operate
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[builtins.int]:
        """
        Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        """
        return pulumi.get(self, "weight")

