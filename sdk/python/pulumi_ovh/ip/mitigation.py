# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MitigationArgs', 'Mitigation']

@pulumi.input_type
class MitigationArgs:
    def __init__(__self__, *,
                 ip: pulumi.Input[str],
                 ip_on_mitigation: pulumi.Input[str],
                 permanent: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Mitigation resource.
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_mitigation: IPv4 address
               * `permanent ` - Set on true if the IP is on permanent mitigation
        :param pulumi.Input[bool] permanent: Set on true if your ip is on permanent mitigation
        """
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "ip_on_mitigation", ip_on_mitigation)
        if permanent is not None:
            pulumi.set(__self__, "permanent", permanent)

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Input[str]:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipOnMitigation")
    def ip_on_mitigation(self) -> pulumi.Input[str]:
        """
        IPv4 address
        * `permanent ` - Set on true if the IP is on permanent mitigation
        """
        return pulumi.get(self, "ip_on_mitigation")

    @ip_on_mitigation.setter
    def ip_on_mitigation(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_on_mitigation", value)

    @property
    @pulumi.getter
    def permanent(self) -> Optional[pulumi.Input[bool]]:
        """
        Set on true if your ip is on permanent mitigation
        """
        return pulumi.get(self, "permanent")

    @permanent.setter
    def permanent(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "permanent", value)


@pulumi.input_type
class _MitigationState:
    def __init__(__self__, *,
                 auto: Optional[pulumi.Input[bool]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_on_mitigation: Optional[pulumi.Input[str]] = None,
                 permanent: Optional[pulumi.Input[bool]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Mitigation resources.
        :param pulumi.Input[bool] auto: Set on true if the IP is on auto-mitigation
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_mitigation: IPv4 address
               * `permanent ` - Set on true if the IP is on permanent mitigation
        :param pulumi.Input[bool] permanent: Set on true if your ip is on permanent mitigation
        :param pulumi.Input[str] state: Current state of the IP on mitigation
        """
        if auto is not None:
            pulumi.set(__self__, "auto", auto)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip_on_mitigation is not None:
            pulumi.set(__self__, "ip_on_mitigation", ip_on_mitigation)
        if permanent is not None:
            pulumi.set(__self__, "permanent", permanent)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def auto(self) -> Optional[pulumi.Input[bool]]:
        """
        Set on true if the IP is on auto-mitigation
        """
        return pulumi.get(self, "auto")

    @auto.setter
    def auto(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipOnMitigation")
    def ip_on_mitigation(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 address
        * `permanent ` - Set on true if the IP is on permanent mitigation
        """
        return pulumi.get(self, "ip_on_mitigation")

    @ip_on_mitigation.setter
    def ip_on_mitigation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_on_mitigation", value)

    @property
    @pulumi.getter
    def permanent(self) -> Optional[pulumi.Input[bool]]:
        """
        Set on true if your ip is on permanent mitigation
        """
        return pulumi.get(self, "permanent")

    @permanent.setter
    def permanent(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "permanent", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Current state of the IP on mitigation
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class Mitigation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_on_mitigation: Optional[pulumi.Input[str]] = None,
                 permanent: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Use this resource to manage an IP permanent mitigation.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        mitigation = ovh.ip.Mitigation("mitigation",
            ip="XXXXXX",
            ip_on_mitigation="XXXXXX")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_mitigation: IPv4 address
               * `permanent ` - Set on true if the IP is on permanent mitigation
        :param pulumi.Input[bool] permanent: Set on true if your ip is on permanent mitigation
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MitigationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Use this resource to manage an IP permanent mitigation.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        mitigation = ovh.ip.Mitigation("mitigation",
            ip="XXXXXX",
            ip_on_mitigation="XXXXXX")
        ```

        :param str resource_name: The name of the resource.
        :param MitigationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MitigationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_on_mitigation: Optional[pulumi.Input[str]] = None,
                 permanent: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MitigationArgs.__new__(MitigationArgs)

            if ip is None and not opts.urn:
                raise TypeError("Missing required property 'ip'")
            __props__.__dict__["ip"] = ip
            if ip_on_mitigation is None and not opts.urn:
                raise TypeError("Missing required property 'ip_on_mitigation'")
            __props__.__dict__["ip_on_mitigation"] = ip_on_mitigation
            __props__.__dict__["permanent"] = permanent
            __props__.__dict__["auto"] = None
            __props__.__dict__["state"] = None
        super(Mitigation, __self__).__init__(
            'ovh:Ip/mitigation:Mitigation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto: Optional[pulumi.Input[bool]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ip_on_mitigation: Optional[pulumi.Input[str]] = None,
            permanent: Optional[pulumi.Input[bool]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'Mitigation':
        """
        Get an existing Mitigation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto: Set on true if the IP is on auto-mitigation
        :param pulumi.Input[str] ip: The IP or the CIDR
        :param pulumi.Input[str] ip_on_mitigation: IPv4 address
               * `permanent ` - Set on true if the IP is on permanent mitigation
        :param pulumi.Input[bool] permanent: Set on true if your ip is on permanent mitigation
        :param pulumi.Input[str] state: Current state of the IP on mitigation
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MitigationState.__new__(_MitigationState)

        __props__.__dict__["auto"] = auto
        __props__.__dict__["ip"] = ip
        __props__.__dict__["ip_on_mitigation"] = ip_on_mitigation
        __props__.__dict__["permanent"] = permanent
        __props__.__dict__["state"] = state
        return Mitigation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def auto(self) -> pulumi.Output[bool]:
        """
        Set on true if the IP is on auto-mitigation
        """
        return pulumi.get(self, "auto")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        The IP or the CIDR
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="ipOnMitigation")
    def ip_on_mitigation(self) -> pulumi.Output[str]:
        """
        IPv4 address
        * `permanent ` - Set on true if the IP is on permanent mitigation
        """
        return pulumi.get(self, "ip_on_mitigation")

    @property
    @pulumi.getter
    def permanent(self) -> pulumi.Output[bool]:
        """
        Set on true if your ip is on permanent mitigation
        """
        return pulumi.get(self, "permanent")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Current state of the IP on mitigation
        """
        return pulumi.get(self, "state")

