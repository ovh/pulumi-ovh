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

__all__ = ['IpV6Args', 'IpV6']

@pulumi.input_type
class IpV6Args:
    def __init__(__self__, *,
                 block: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a IpV6 resource.
        :param pulumi.Input[str] block: Your IPv6 block.
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        pulumi.set(__self__, "block", block)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def block(self) -> pulumi.Input[str]:
        """
        Your IPv6 block.
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: pulumi.Input[str]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _IpV6State:
    def __init__(__self__, *,
                 block: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpV6 resources.
        :param pulumi.Input[str] block: Your IPv6 block.
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        if block is not None:
            pulumi.set(__self__, "block", block)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def block(self) -> Optional[pulumi.Input[str]]:
        """
        Your IPv6 block.
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class IpV6(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attach an IPv6 block to a VRack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_block = ovh.vrack.IpV6("vrackBlock",
            block="<ipv6 block>",
            service_name="<vRack service name>")
        ```

        ## Import

        Attachment of an IPv6 block and a VRack can be imported using the `service_name` (vRack identifier) and the `block` (IPv6 block), separated by "," E.g.,

        bash

        ```sh
        $ pulumi import ovh:Vrack/ipV6:IpV6 myattach "<service_name>,<block>"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block: Your IPv6 block.
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpV6Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attach an IPv6 block to a VRack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_block = ovh.vrack.IpV6("vrackBlock",
            block="<ipv6 block>",
            service_name="<vRack service name>")
        ```

        ## Import

        Attachment of an IPv6 block and a VRack can be imported using the `service_name` (vRack identifier) and the `block` (IPv6 block), separated by "," E.g.,

        bash

        ```sh
        $ pulumi import ovh:Vrack/ipV6:IpV6 myattach "<service_name>,<block>"
        ```

        :param str resource_name: The name of the resource.
        :param IpV6Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpV6Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpV6Args.__new__(IpV6Args)

            if block is None and not opts.urn:
                raise TypeError("Missing required property 'block'")
            __props__.__dict__["block"] = block
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(IpV6, __self__).__init__(
            'ovh:Vrack/ipV6:IpV6',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'IpV6':
        """
        Get an existing IpV6 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block: Your IPv6 block.
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpV6State.__new__(_IpV6State)

        __props__.__dict__["block"] = block
        __props__.__dict__["service_name"] = service_name
        return IpV6(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def block(self) -> pulumi.Output[str]:
        """
        Your IPv6 block.
        """
        return pulumi.get(self, "block")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

