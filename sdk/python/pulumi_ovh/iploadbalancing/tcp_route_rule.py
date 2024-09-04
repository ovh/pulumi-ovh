# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TcpRouteRuleInitArgs', 'TcpRouteRule']

@pulumi.input_type
class TcpRouteRuleInitArgs:
    def __init__(__self__, *,
                 field: pulumi.Input[str],
                 match: pulumi.Input[str],
                 route_id: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 negate: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 sub_field: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TcpRouteRule resource.
        :param pulumi.Input[str] field: Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        :param pulumi.Input[str] match: Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        :param pulumi.Input[str] route_id: The route to apply this rule
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] display_name: Human readable name for your rule, this field is for you
        :param pulumi.Input[bool] negate: Invert the matching operator effect
        :param pulumi.Input[str] pattern: Value to match against this match. Interpretation if this field depends on the match and field
        :param pulumi.Input[str] sub_field: Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        """
        pulumi.set(__self__, "field", field)
        pulumi.set(__self__, "match", match)
        pulumi.set(__self__, "route_id", route_id)
        pulumi.set(__self__, "service_name", service_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if negate is not None:
            pulumi.set(__self__, "negate", negate)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if sub_field is not None:
            pulumi.set(__self__, "sub_field", sub_field)

    @property
    @pulumi.getter
    def field(self) -> pulumi.Input[str]:
        """
        Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: pulumi.Input[str]):
        pulumi.set(self, "field", value)

    @property
    @pulumi.getter
    def match(self) -> pulumi.Input[str]:
        """
        Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: pulumi.Input[str]):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Input[str]:
        """
        The route to apply this rule
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for your rule, this field is for you
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def negate(self) -> Optional[pulumi.Input[bool]]:
        """
        Invert the matching operator effect
        """
        return pulumi.get(self, "negate")

    @negate.setter
    def negate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "negate", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Value to match against this match. Interpretation if this field depends on the match and field
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter(name="subField")
    def sub_field(self) -> Optional[pulumi.Input[str]]:
        """
        Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        """
        return pulumi.get(self, "sub_field")

    @sub_field.setter
    def sub_field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_field", value)


@pulumi.input_type
class _TcpRouteRuleState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 match: Optional[pulumi.Input[str]] = None,
                 negate: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sub_field: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TcpRouteRule resources.
        :param pulumi.Input[str] display_name: Human readable name for your rule, this field is for you
        :param pulumi.Input[str] field: Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        :param pulumi.Input[str] match: Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        :param pulumi.Input[bool] negate: Invert the matching operator effect
        :param pulumi.Input[str] pattern: Value to match against this match. Interpretation if this field depends on the match and field
        :param pulumi.Input[str] route_id: The route to apply this rule
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] sub_field: Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if field is not None:
            pulumi.set(__self__, "field", field)
        if match is not None:
            pulumi.set(__self__, "match", match)
        if negate is not None:
            pulumi.set(__self__, "negate", negate)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if route_id is not None:
            pulumi.set(__self__, "route_id", route_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if sub_field is not None:
            pulumi.set(__self__, "sub_field", sub_field)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable name for your rule, this field is for you
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def field(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        """
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field", value)

    @property
    @pulumi.getter
    def match(self) -> Optional[pulumi.Input[str]]:
        """
        Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter
    def negate(self) -> Optional[pulumi.Input[bool]]:
        """
        Invert the matching operator effect
        """
        return pulumi.get(self, "negate")

    @negate.setter
    def negate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "negate", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Value to match against this match. Interpretation if this field depends on the match and field
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> Optional[pulumi.Input[str]]:
        """
        The route to apply this rule
        """
        return pulumi.get(self, "route_id")

    @route_id.setter
    def route_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="subField")
    def sub_field(self) -> Optional[pulumi.Input[str]]:
        """
        Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        """
        return pulumi.get(self, "sub_field")

    @sub_field.setter
    def sub_field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_field", value)


class TcpRouteRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 match: Optional[pulumi.Input[str]] = None,
                 negate: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sub_field: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage rules for TCP route.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        reject = ovh.ip_load_balancing.TcpRoute("reject",
            service_name="loadbalancer-xxxxxxxxxxxxxxxxxx",
            weight=1,
            frontend_id=11111,
            action={
                "type": "reject",
            })
        examplerule = ovh.ip_load_balancing.TcpRouteRule("examplerule",
            service_name="loadbalancer-xxxxxxxxxxxxxxxxxx",
            route_id=reject.id,
            display_name="Match example.com host",
            field="sni",
            match="is",
            negate=False,
            pattern="example.com")
        ```

        ## Import

        TCP route rule can be imported using the following format `service_name`, the `id` of the route and the `id` of the rule separated by "/" e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human readable name for your rule, this field is for you
        :param pulumi.Input[str] field: Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        :param pulumi.Input[str] match: Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        :param pulumi.Input[bool] negate: Invert the matching operator effect
        :param pulumi.Input[str] pattern: Value to match against this match. Interpretation if this field depends on the match and field
        :param pulumi.Input[str] route_id: The route to apply this rule
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] sub_field: Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TcpRouteRuleInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage rules for TCP route.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        reject = ovh.ip_load_balancing.TcpRoute("reject",
            service_name="loadbalancer-xxxxxxxxxxxxxxxxxx",
            weight=1,
            frontend_id=11111,
            action={
                "type": "reject",
            })
        examplerule = ovh.ip_load_balancing.TcpRouteRule("examplerule",
            service_name="loadbalancer-xxxxxxxxxxxxxxxxxx",
            route_id=reject.id,
            display_name="Match example.com host",
            field="sni",
            match="is",
            negate=False,
            pattern="example.com")
        ```

        ## Import

        TCP route rule can be imported using the following format `service_name`, the `id` of the route and the `id` of the rule separated by "/" e.g.

        :param str resource_name: The name of the resource.
        :param TcpRouteRuleInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TcpRouteRuleInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 field: Optional[pulumi.Input[str]] = None,
                 match: Optional[pulumi.Input[str]] = None,
                 negate: Optional[pulumi.Input[bool]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 route_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sub_field: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TcpRouteRuleInitArgs.__new__(TcpRouteRuleInitArgs)

            __props__.__dict__["display_name"] = display_name
            if field is None and not opts.urn:
                raise TypeError("Missing required property 'field'")
            __props__.__dict__["field"] = field
            if match is None and not opts.urn:
                raise TypeError("Missing required property 'match'")
            __props__.__dict__["match"] = match
            __props__.__dict__["negate"] = negate
            __props__.__dict__["pattern"] = pattern
            if route_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_id'")
            __props__.__dict__["route_id"] = route_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["sub_field"] = sub_field
        super(TcpRouteRule, __self__).__init__(
            'ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            field: Optional[pulumi.Input[str]] = None,
            match: Optional[pulumi.Input[str]] = None,
            negate: Optional[pulumi.Input[bool]] = None,
            pattern: Optional[pulumi.Input[str]] = None,
            route_id: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            sub_field: Optional[pulumi.Input[str]] = None) -> 'TcpRouteRule':
        """
        Get an existing TcpRouteRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: Human readable name for your rule, this field is for you
        :param pulumi.Input[str] field: Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        :param pulumi.Input[str] match: Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        :param pulumi.Input[bool] negate: Invert the matching operator effect
        :param pulumi.Input[str] pattern: Value to match against this match. Interpretation if this field depends on the match and field
        :param pulumi.Input[str] route_id: The route to apply this rule
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] sub_field: Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TcpRouteRuleState.__new__(_TcpRouteRuleState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["field"] = field
        __props__.__dict__["match"] = match
        __props__.__dict__["negate"] = negate
        __props__.__dict__["pattern"] = pattern
        __props__.__dict__["route_id"] = route_id
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["sub_field"] = sub_field
        return TcpRouteRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Human readable name for your rule, this field is for you
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def field(self) -> pulumi.Output[str]:
        """
        Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        """
        return pulumi.get(self, "field")

    @property
    @pulumi.getter
    def match(self) -> pulumi.Output[str]:
        """
        Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        """
        return pulumi.get(self, "match")

    @property
    @pulumi.getter
    def negate(self) -> pulumi.Output[bool]:
        """
        Invert the matching operator effect
        """
        return pulumi.get(self, "negate")

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[Optional[str]]:
        """
        Value to match against this match. Interpretation if this field depends on the match and field
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter(name="routeId")
    def route_id(self) -> pulumi.Output[str]:
        """
        The route to apply this rule
        """
        return pulumi.get(self, "route_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="subField")
    def sub_field(self) -> pulumi.Output[Optional[str]]:
        """
        Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        """
        return pulumi.get(self, "sub_field")

