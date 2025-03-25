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

__all__ = ['DedicatedCloudArgs', 'DedicatedCloud']

@pulumi.input_type
class DedicatedCloudArgs:
    def __init__(__self__, *,
                 dedicated_cloud: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a DedicatedCloud resource.
        :param pulumi.Input[str] dedicated_cloud: Your Dedicated Cloud service name
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        pulumi.set(__self__, "dedicated_cloud", dedicated_cloud)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="dedicatedCloud")
    def dedicated_cloud(self) -> pulumi.Input[str]:
        """
        Your Dedicated Cloud service name
        """
        return pulumi.get(self, "dedicated_cloud")

    @dedicated_cloud.setter
    def dedicated_cloud(self, value: pulumi.Input[str]):
        pulumi.set(self, "dedicated_cloud", value)

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
class _DedicatedCloudState:
    def __init__(__self__, *,
                 dedicated_cloud: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DedicatedCloud resources.
        :param pulumi.Input[str] dedicated_cloud: Your Dedicated Cloud service name
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        if dedicated_cloud is not None:
            pulumi.set(__self__, "dedicated_cloud", dedicated_cloud)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="dedicatedCloud")
    def dedicated_cloud(self) -> Optional[pulumi.Input[str]]:
        """
        Your Dedicated Cloud service name
        """
        return pulumi.get(self, "dedicated_cloud")

    @dedicated_cloud.setter
    def dedicated_cloud(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dedicated_cloud", value)

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


class DedicatedCloud(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_cloud: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attach a Dedicated Cloud to the vrack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_dedicated_cloud = ovh.vrack.DedicatedCloud("vrack-dedicatedCloud",
            dedicated_cloud="<Dedicated Cloud service name>",
            service_name="<vRack service name>")
        ```

        ## Import

        Attachment of a Dedicated Cloud and a vRack can be imported using the `service_name` (vRack identifier) and the `dedicated_cloud` (Dedicated Cloud service name), separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Vrack/dedicatedCloud:DedicatedCloud myattach "<vRack service name>/<Dedicated Cloud service name>"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dedicated_cloud: Your Dedicated Cloud service name
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedCloudArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attach a Dedicated Cloud to the vrack.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_dedicated_cloud = ovh.vrack.DedicatedCloud("vrack-dedicatedCloud",
            dedicated_cloud="<Dedicated Cloud service name>",
            service_name="<vRack service name>")
        ```

        ## Import

        Attachment of a Dedicated Cloud and a vRack can be imported using the `service_name` (vRack identifier) and the `dedicated_cloud` (Dedicated Cloud service name), separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Vrack/dedicatedCloud:DedicatedCloud myattach "<vRack service name>/<Dedicated Cloud service name>"
        ```

        :param str resource_name: The name of the resource.
        :param DedicatedCloudArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedCloudArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_cloud: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedCloudArgs.__new__(DedicatedCloudArgs)

            if dedicated_cloud is None and not opts.urn:
                raise TypeError("Missing required property 'dedicated_cloud'")
            __props__.__dict__["dedicated_cloud"] = dedicated_cloud
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(DedicatedCloud, __self__).__init__(
            'ovh:Vrack/dedicatedCloud:DedicatedCloud',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dedicated_cloud: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'DedicatedCloud':
        """
        Get an existing DedicatedCloud resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dedicated_cloud: Your Dedicated Cloud service name
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DedicatedCloudState.__new__(_DedicatedCloudState)

        __props__.__dict__["dedicated_cloud"] = dedicated_cloud
        __props__.__dict__["service_name"] = service_name
        return DedicatedCloud(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dedicatedCloud")
    def dedicated_cloud(self) -> pulumi.Output[str]:
        """
        Your Dedicated Cloud service name
        """
        return pulumi.get(self, "dedicated_cloud")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

