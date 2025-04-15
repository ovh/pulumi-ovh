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

__all__ = ['VrackservicesArgs', 'Vrackservices']

@pulumi.input_type
class VrackservicesArgs:
    def __init__(__self__, *,
                 service_name: pulumi.Input[builtins.str],
                 vrack_services: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a Vrackservices resource.
        :param pulumi.Input[builtins.str] service_name: The internal name of your vrack
        :param pulumi.Input[builtins.str] vrack_services: Your vrackServices service name.
        """
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "vrack_services", vrack_services)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[builtins.str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="vrackServices")
    def vrack_services(self) -> pulumi.Input[builtins.str]:
        """
        Your vrackServices service name.
        """
        return pulumi.get(self, "vrack_services")

    @vrack_services.setter
    def vrack_services(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "vrack_services", value)


@pulumi.input_type
class _VrackservicesState:
    def __init__(__self__, *,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 vrack_services: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Vrackservices resources.
        :param pulumi.Input[builtins.str] service_name: The internal name of your vrack
        :param pulumi.Input[builtins.str] vrack_services: Your vrackServices service name.
        """
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if vrack_services is not None:
            pulumi.set(__self__, "vrack_services", vrack_services)

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

    @property
    @pulumi.getter(name="vrackServices")
    def vrack_services(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Your vrackServices service name.
        """
        return pulumi.get(self, "vrack_services")

    @vrack_services.setter
    def vrack_services(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "vrack_services", value)


class Vrackservices(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 vrack_services: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Attach a vrackServices to the vrack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_vrackservices = ovh.vrack.Vrackservices("vrack_vrackservices",
            service_name="<vRack service name>",
            vrack_services="<vrackServices service name>")
        ```

        ## Import

        Attachment of a vrackServices and a vRack can be imported using the `service_name` (vRack identifier) and the `vrack_services` (vrackServices service name), separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Vrack/vrackservices:Vrackservices myattach "<service_name>/<vrackServices service name>"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] service_name: The internal name of your vrack
        :param pulumi.Input[builtins.str] vrack_services: Your vrackServices service name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VrackservicesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attach a vrackServices to the vrack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_vrackservices = ovh.vrack.Vrackservices("vrack_vrackservices",
            service_name="<vRack service name>",
            vrack_services="<vrackServices service name>")
        ```

        ## Import

        Attachment of a vrackServices and a vRack can be imported using the `service_name` (vRack identifier) and the `vrack_services` (vrackServices service name), separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Vrack/vrackservices:Vrackservices myattach "<service_name>/<vrackServices service name>"
        ```

        :param str resource_name: The name of the resource.
        :param VrackservicesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VrackservicesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_name: Optional[pulumi.Input[builtins.str]] = None,
                 vrack_services: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VrackservicesArgs.__new__(VrackservicesArgs)

            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if vrack_services is None and not opts.urn:
                raise TypeError("Missing required property 'vrack_services'")
            __props__.__dict__["vrack_services"] = vrack_services
        super(Vrackservices, __self__).__init__(
            'ovh:Vrack/vrackservices:Vrackservices',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            service_name: Optional[pulumi.Input[builtins.str]] = None,
            vrack_services: Optional[pulumi.Input[builtins.str]] = None) -> 'Vrackservices':
        """
        Get an existing Vrackservices resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] service_name: The internal name of your vrack
        :param pulumi.Input[builtins.str] vrack_services: Your vrackServices service name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VrackservicesState.__new__(_VrackservicesState)

        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["vrack_services"] = vrack_services
        return Vrackservices(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[builtins.str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="vrackServices")
    def vrack_services(self) -> pulumi.Output[builtins.str]:
        """
        Your vrackServices service name.
        """
        return pulumi.get(self, "vrack_services")

