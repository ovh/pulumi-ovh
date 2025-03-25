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

__all__ = ['DedicatedServerArgs', 'DedicatedServer']

@pulumi.input_type
class DedicatedServerArgs:
    def __init__(__self__, *,
                 server_id: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a DedicatedServer resource.
        :param pulumi.Input[str] server_id: The id of the dedicated server.
        :param pulumi.Input[str] service_name: The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        """
        pulumi.set(__self__, "server_id", server_id)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Input[str]:
        """
        The id of the dedicated server.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _DedicatedServerState:
    def __init__(__self__, *,
                 server_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DedicatedServer resources.
        :param pulumi.Input[str] server_id: The id of the dedicated server.
        :param pulumi.Input[str] service_name: The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        """
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the dedicated server.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class DedicatedServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attach a legacy dedicated server to a vRack.

        > **NOTE:** The resource `Vrack.DedicatedServer` is intended to be used for legacy dedicated servers.<br /> Dedicated servers that have configurable network interfaces MUST use the resource `Vrack.DedicatedServerInterface` instead.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vds = ovh.vrack.DedicatedServer("vds",
            server_id="67890",
            service_name="XXXX")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] server_id: The id of the dedicated server.
        :param pulumi.Input[str] service_name: The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DedicatedServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attach a legacy dedicated server to a vRack.

        > **NOTE:** The resource `Vrack.DedicatedServer` is intended to be used for legacy dedicated servers.<br /> Dedicated servers that have configurable network interfaces MUST use the resource `Vrack.DedicatedServerInterface` instead.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vds = ovh.vrack.DedicatedServer("vds",
            server_id="67890",
            service_name="XXXX")
        ```

        :param str resource_name: The name of the resource.
        :param DedicatedServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DedicatedServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DedicatedServerArgs.__new__(DedicatedServerArgs)

            if server_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_id'")
            __props__.__dict__["server_id"] = server_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(DedicatedServer, __self__).__init__(
            'ovh:Vrack/dedicatedServer:DedicatedServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'DedicatedServer':
        """
        Get an existing DedicatedServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] server_id: The id of the dedicated server.
        :param pulumi.Input[str] service_name: The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DedicatedServerState.__new__(_DedicatedServerState)

        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["service_name"] = service_name
        return DedicatedServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The id of the dedicated server.
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

