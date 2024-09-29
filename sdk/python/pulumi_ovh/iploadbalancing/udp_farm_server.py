# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UdpFarmServerArgs', 'UdpFarmServer']

@pulumi.input_type
class UdpFarmServerArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 farm_id: pulumi.Input[float],
                 service_name: pulumi.Input[str],
                 status: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None):
        """
        The set of arguments for constructing a UdpFarmServer resource.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[float] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[float] port: Port that backend will respond on
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "farm_id", farm_id)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "status", status)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        Address of the backend server (IP from either internal or OVHcloud network)
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="farmId")
    def farm_id(self) -> pulumi.Input[float]:
        """
        ID of the farm this server is attached to
        """
        return pulumi.get(self, "farm_id")

    @farm_id.setter
    def farm_id(self, value: pulumi.Input[float]):
        pulumi.set(self, "farm_id", value)

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
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        backend status - `active` or `inactive`
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Label for the server
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[float]]:
        """
        Port that backend will respond on
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class _UdpFarmServerState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 backend_id: Optional[pulumi.Input[float]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_id: Optional[pulumi.Input[float]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 server_id: Optional[pulumi.Input[float]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UdpFarmServer resources.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[float] backend_id: Synonym for `farm_id`.
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[float] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[float] port: Port that backend will respond on
        :param pulumi.Input[float] server_id: Id of your server.
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if backend_id is not None:
            pulumi.set(__self__, "backend_id", backend_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if farm_id is not None:
            pulumi.set(__self__, "farm_id", farm_id)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the backend server (IP from either internal or OVHcloud network)
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> Optional[pulumi.Input[float]]:
        """
        Synonym for `farm_id`.
        """
        return pulumi.get(self, "backend_id")

    @backend_id.setter
    def backend_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "backend_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Label for the server
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="farmId")
    def farm_id(self) -> Optional[pulumi.Input[float]]:
        """
        ID of the farm this server is attached to
        """
        return pulumi.get(self, "farm_id")

    @farm_id.setter
    def farm_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "farm_id", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[float]]:
        """
        Port that backend will respond on
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[float]]:
        """
        Id of your server.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "server_id", value)

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
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        backend status - `active` or `inactive`
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class UdpFarmServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_id: Optional[pulumi.Input[float]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a backend server entry linked to loadbalancing group (farm)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farmname = ovh.ip_load_balancing.UdpFarm("farmname",
            display_name="ingress-8080-gra",
            port=80,
            service_name=lb.service_name,
            zone="gra")
        backend = ovh.ip_load_balancing.UdpFarmServer("backend",
            address="4.5.6.7",
            display_name="mybackend",
            farm_id=farmname.farm_id,
            port=80,
            service_name=lb.service_name,
            status="active")
        ```

        ## Import

        UDP farm server can be imported using the following format `service_name`, the `id` of the farm and the `id` of the server separated by "/" e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[float] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[float] port: Port that backend will respond on
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UdpFarmServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backend server entry linked to loadbalancing group (farm)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farmname = ovh.ip_load_balancing.UdpFarm("farmname",
            display_name="ingress-8080-gra",
            port=80,
            service_name=lb.service_name,
            zone="gra")
        backend = ovh.ip_load_balancing.UdpFarmServer("backend",
            address="4.5.6.7",
            display_name="mybackend",
            farm_id=farmname.farm_id,
            port=80,
            service_name=lb.service_name,
            status="active")
        ```

        ## Import

        UDP farm server can be imported using the following format `service_name`, the `id` of the farm and the `id` of the server separated by "/" e.g.

        :param str resource_name: The name of the resource.
        :param UdpFarmServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UdpFarmServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_id: Optional[pulumi.Input[float]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UdpFarmServerArgs.__new__(UdpFarmServerArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["display_name"] = display_name
            if farm_id is None and not opts.urn:
                raise TypeError("Missing required property 'farm_id'")
            __props__.__dict__["farm_id"] = farm_id
            __props__.__dict__["port"] = port
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["backend_id"] = None
            __props__.__dict__["server_id"] = None
        super(UdpFarmServer, __self__).__init__(
            'ovh:IpLoadBalancing/udpFarmServer:UdpFarmServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            backend_id: Optional[pulumi.Input[float]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            farm_id: Optional[pulumi.Input[float]] = None,
            port: Optional[pulumi.Input[float]] = None,
            server_id: Optional[pulumi.Input[float]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'UdpFarmServer':
        """
        Get an existing UdpFarmServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[float] backend_id: Synonym for `farm_id`.
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[float] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[float] port: Port that backend will respond on
        :param pulumi.Input[float] server_id: Id of your server.
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UdpFarmServerState.__new__(_UdpFarmServerState)

        __props__.__dict__["address"] = address
        __props__.__dict__["backend_id"] = backend_id
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["farm_id"] = farm_id
        __props__.__dict__["port"] = port
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        return UdpFarmServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Address of the backend server (IP from either internal or OVHcloud network)
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="backendId")
    def backend_id(self) -> pulumi.Output[float]:
        """
        Synonym for `farm_id`.
        """
        return pulumi.get(self, "backend_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Label for the server
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="farmId")
    def farm_id(self) -> pulumi.Output[float]:
        """
        ID of the farm this server is attached to
        """
        return pulumi.get(self, "farm_id")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[float]:
        """
        Port that backend will respond on
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[float]:
        """
        Id of your server.
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        backend status - `active` or `inactive`
        """
        return pulumi.get(self, "status")
