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

__all__ = ['HttpFarmServerArgs', 'HttpFarmServer']

@pulumi.input_type
class HttpFarmServerArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 farm_id: pulumi.Input[int],
                 service_name: pulumi.Input[str],
                 status: pulumi.Input[str],
                 backup: Optional[pulumi.Input[bool]] = None,
                 chain: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 on_marked_down: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 probe: Optional[pulumi.Input[bool]] = None,
                 proxy_protocol_version: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a HttpFarmServer resource.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[int] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        :param pulumi.Input[bool] backup: is it a backup server used in case of failure of all the non-backup backends
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[str] on_marked_down: enable action when backend marked down. (`shutdown-sessions`)
        :param pulumi.Input[int] port: Port that backend will respond on
        :param pulumi.Input[bool] probe: defines if backend will be probed to determine health and keep as active in farm if healthy
        :param pulumi.Input[str] proxy_protocol_version: version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        :param pulumi.Input[bool] ssl: is the connection ciphered with SSL (TLS)
        :param pulumi.Input[int] weight: used in loadbalancing algorithm
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "farm_id", farm_id)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "status", status)
        if backup is not None:
            pulumi.set(__self__, "backup", backup)
        if chain is not None:
            pulumi.set(__self__, "chain", chain)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if on_marked_down is not None:
            pulumi.set(__self__, "on_marked_down", on_marked_down)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if probe is not None:
            pulumi.set(__self__, "probe", probe)
        if proxy_protocol_version is not None:
            pulumi.set(__self__, "proxy_protocol_version", proxy_protocol_version)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

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
    def farm_id(self) -> pulumi.Input[int]:
        """
        ID of the farm this server is attached to
        """
        return pulumi.get(self, "farm_id")

    @farm_id.setter
    def farm_id(self, value: pulumi.Input[int]):
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
    @pulumi.getter
    def backup(self) -> Optional[pulumi.Input[bool]]:
        """
        is it a backup server used in case of failure of all the non-backup backends
        """
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter
    def chain(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "chain")

    @chain.setter
    def chain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "chain", value)

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
    @pulumi.getter(name="onMarkedDown")
    def on_marked_down(self) -> Optional[pulumi.Input[str]]:
        """
        enable action when backend marked down. (`shutdown-sessions`)
        """
        return pulumi.get(self, "on_marked_down")

    @on_marked_down.setter
    def on_marked_down(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "on_marked_down", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port that backend will respond on
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def probe(self) -> Optional[pulumi.Input[bool]]:
        """
        defines if backend will be probed to determine health and keep as active in farm if healthy
        """
        return pulumi.get(self, "probe")

    @probe.setter
    def probe(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "probe", value)

    @property
    @pulumi.getter(name="proxyProtocolVersion")
    def proxy_protocol_version(self) -> Optional[pulumi.Input[str]]:
        """
        version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        """
        return pulumi.get(self, "proxy_protocol_version")

    @proxy_protocol_version.setter
    def proxy_protocol_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_protocol_version", value)

    @property
    @pulumi.getter
    def ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        is the connection ciphered with SSL (TLS)
        """
        return pulumi.get(self, "ssl")

    @ssl.setter
    def ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssl", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        used in loadbalancing algorithm
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _HttpFarmServerState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 chain: Optional[pulumi.Input[str]] = None,
                 cookie: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_id: Optional[pulumi.Input[int]] = None,
                 on_marked_down: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 probe: Optional[pulumi.Input[bool]] = None,
                 proxy_protocol_version: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering HttpFarmServer resources.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[bool] backup: is it a backup server used in case of failure of all the non-backup backends
        :param pulumi.Input[str] cookie: Value of the stickiness cookie used for this backend.
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[int] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[str] on_marked_down: enable action when backend marked down. (`shutdown-sessions`)
        :param pulumi.Input[int] port: Port that backend will respond on
        :param pulumi.Input[bool] probe: defines if backend will be probed to determine health and keep as active in farm if healthy
        :param pulumi.Input[str] proxy_protocol_version: version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[bool] ssl: is the connection ciphered with SSL (TLS)
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        :param pulumi.Input[int] weight: used in loadbalancing algorithm
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if backup is not None:
            pulumi.set(__self__, "backup", backup)
        if chain is not None:
            pulumi.set(__self__, "chain", chain)
        if cookie is not None:
            pulumi.set(__self__, "cookie", cookie)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if farm_id is not None:
            pulumi.set(__self__, "farm_id", farm_id)
        if on_marked_down is not None:
            pulumi.set(__self__, "on_marked_down", on_marked_down)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if probe is not None:
            pulumi.set(__self__, "probe", probe)
        if proxy_protocol_version is not None:
            pulumi.set(__self__, "proxy_protocol_version", proxy_protocol_version)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

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
    @pulumi.getter
    def backup(self) -> Optional[pulumi.Input[bool]]:
        """
        is it a backup server used in case of failure of all the non-backup backends
        """
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter
    def chain(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "chain")

    @chain.setter
    def chain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "chain", value)

    @property
    @pulumi.getter
    def cookie(self) -> Optional[pulumi.Input[str]]:
        """
        Value of the stickiness cookie used for this backend.
        """
        return pulumi.get(self, "cookie")

    @cookie.setter
    def cookie(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cookie", value)

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
    def farm_id(self) -> Optional[pulumi.Input[int]]:
        """
        ID of the farm this server is attached to
        """
        return pulumi.get(self, "farm_id")

    @farm_id.setter
    def farm_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "farm_id", value)

    @property
    @pulumi.getter(name="onMarkedDown")
    def on_marked_down(self) -> Optional[pulumi.Input[str]]:
        """
        enable action when backend marked down. (`shutdown-sessions`)
        """
        return pulumi.get(self, "on_marked_down")

    @on_marked_down.setter
    def on_marked_down(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "on_marked_down", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        Port that backend will respond on
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def probe(self) -> Optional[pulumi.Input[bool]]:
        """
        defines if backend will be probed to determine health and keep as active in farm if healthy
        """
        return pulumi.get(self, "probe")

    @probe.setter
    def probe(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "probe", value)

    @property
    @pulumi.getter(name="proxyProtocolVersion")
    def proxy_protocol_version(self) -> Optional[pulumi.Input[str]]:
        """
        version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        """
        return pulumi.get(self, "proxy_protocol_version")

    @proxy_protocol_version.setter
    def proxy_protocol_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_protocol_version", value)

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
    def ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        is the connection ciphered with SSL (TLS)
        """
        return pulumi.get(self, "ssl")

    @ssl.setter
    def ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ssl", value)

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

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        used in loadbalancing algorithm
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class HttpFarmServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 chain: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_id: Optional[pulumi.Input[int]] = None,
                 on_marked_down: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 probe: Optional[pulumi.Input[bool]] = None,
                 proxy_protocol_version: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Creates a backend server entry linked to HTTP loadbalancing group (farm)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farmname = ovh.ip_load_balancing.HttpFarm("farmname",
            port=8080,
            service_name=lb.service_name,
            zone="all")
        backend = ovh.ip_load_balancing.HttpFarmServer("backend",
            address="4.5.6.7",
            backup=True,
            display_name="mybackend",
            farm_id=farmname.id,
            port=80,
            probe=True,
            proxy_protocol_version="v2",
            service_name=lb.service_name,
            ssl=False,
            status="active",
            weight=2)
        ```

        ## Import

        HTTP farm server can be imported using the following format `service_name`, the `id` of the farm and the `id` of the server separated by "/" e.g.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[bool] backup: is it a backup server used in case of failure of all the non-backup backends
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[int] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[str] on_marked_down: enable action when backend marked down. (`shutdown-sessions`)
        :param pulumi.Input[int] port: Port that backend will respond on
        :param pulumi.Input[bool] probe: defines if backend will be probed to determine health and keep as active in farm if healthy
        :param pulumi.Input[str] proxy_protocol_version: version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[bool] ssl: is the connection ciphered with SSL (TLS)
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        :param pulumi.Input[int] weight: used in loadbalancing algorithm
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HttpFarmServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a backend server entry linked to HTTP loadbalancing group (farm)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        lb = ovh.IpLoadBalancing.get_ip_load_balancing(service_name="ip-1.2.3.4",
            state="ok")
        farmname = ovh.ip_load_balancing.HttpFarm("farmname",
            port=8080,
            service_name=lb.service_name,
            zone="all")
        backend = ovh.ip_load_balancing.HttpFarmServer("backend",
            address="4.5.6.7",
            backup=True,
            display_name="mybackend",
            farm_id=farmname.id,
            port=80,
            probe=True,
            proxy_protocol_version="v2",
            service_name=lb.service_name,
            ssl=False,
            status="active",
            weight=2)
        ```

        ## Import

        HTTP farm server can be imported using the following format `service_name`, the `id` of the farm and the `id` of the server separated by "/" e.g.

        :param str resource_name: The name of the resource.
        :param HttpFarmServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HttpFarmServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 backup: Optional[pulumi.Input[bool]] = None,
                 chain: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 farm_id: Optional[pulumi.Input[int]] = None,
                 on_marked_down: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 probe: Optional[pulumi.Input[bool]] = None,
                 proxy_protocol_version: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[bool]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HttpFarmServerArgs.__new__(HttpFarmServerArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["backup"] = backup
            __props__.__dict__["chain"] = chain
            __props__.__dict__["display_name"] = display_name
            if farm_id is None and not opts.urn:
                raise TypeError("Missing required property 'farm_id'")
            __props__.__dict__["farm_id"] = farm_id
            __props__.__dict__["on_marked_down"] = on_marked_down
            __props__.__dict__["port"] = port
            __props__.__dict__["probe"] = probe
            __props__.__dict__["proxy_protocol_version"] = proxy_protocol_version
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["ssl"] = ssl
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["weight"] = weight
            __props__.__dict__["cookie"] = None
        super(HttpFarmServer, __self__).__init__(
            'ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            backup: Optional[pulumi.Input[bool]] = None,
            chain: Optional[pulumi.Input[str]] = None,
            cookie: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            farm_id: Optional[pulumi.Input[int]] = None,
            on_marked_down: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            probe: Optional[pulumi.Input[bool]] = None,
            proxy_protocol_version: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            ssl: Optional[pulumi.Input[bool]] = None,
            status: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[int]] = None) -> 'HttpFarmServer':
        """
        Get an existing HttpFarmServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Address of the backend server (IP from either internal or OVHcloud network)
        :param pulumi.Input[bool] backup: is it a backup server used in case of failure of all the non-backup backends
        :param pulumi.Input[str] cookie: Value of the stickiness cookie used for this backend.
        :param pulumi.Input[str] display_name: Label for the server
        :param pulumi.Input[int] farm_id: ID of the farm this server is attached to
        :param pulumi.Input[str] on_marked_down: enable action when backend marked down. (`shutdown-sessions`)
        :param pulumi.Input[int] port: Port that backend will respond on
        :param pulumi.Input[bool] probe: defines if backend will be probed to determine health and keep as active in farm if healthy
        :param pulumi.Input[str] proxy_protocol_version: version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        :param pulumi.Input[str] service_name: The internal name of your IP load balancing
        :param pulumi.Input[bool] ssl: is the connection ciphered with SSL (TLS)
        :param pulumi.Input[str] status: backend status - `active` or `inactive`
        :param pulumi.Input[int] weight: used in loadbalancing algorithm
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HttpFarmServerState.__new__(_HttpFarmServerState)

        __props__.__dict__["address"] = address
        __props__.__dict__["backup"] = backup
        __props__.__dict__["chain"] = chain
        __props__.__dict__["cookie"] = cookie
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["farm_id"] = farm_id
        __props__.__dict__["on_marked_down"] = on_marked_down
        __props__.__dict__["port"] = port
        __props__.__dict__["probe"] = probe
        __props__.__dict__["proxy_protocol_version"] = proxy_protocol_version
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["ssl"] = ssl
        __props__.__dict__["status"] = status
        __props__.__dict__["weight"] = weight
        return HttpFarmServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Address of the backend server (IP from either internal or OVHcloud network)
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def backup(self) -> pulumi.Output[Optional[bool]]:
        """
        is it a backup server used in case of failure of all the non-backup backends
        """
        return pulumi.get(self, "backup")

    @property
    @pulumi.getter
    def chain(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "chain")

    @property
    @pulumi.getter
    def cookie(self) -> pulumi.Output[str]:
        """
        Value of the stickiness cookie used for this backend.
        """
        return pulumi.get(self, "cookie")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Label for the server
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="farmId")
    def farm_id(self) -> pulumi.Output[int]:
        """
        ID of the farm this server is attached to
        """
        return pulumi.get(self, "farm_id")

    @property
    @pulumi.getter(name="onMarkedDown")
    def on_marked_down(self) -> pulumi.Output[Optional[str]]:
        """
        enable action when backend marked down. (`shutdown-sessions`)
        """
        return pulumi.get(self, "on_marked_down")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        Port that backend will respond on
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def probe(self) -> pulumi.Output[Optional[bool]]:
        """
        defines if backend will be probed to determine health and keep as active in farm if healthy
        """
        return pulumi.get(self, "probe")

    @property
    @pulumi.getter(name="proxyProtocolVersion")
    def proxy_protocol_version(self) -> pulumi.Output[Optional[str]]:
        """
        version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        """
        return pulumi.get(self, "proxy_protocol_version")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your IP load balancing
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def ssl(self) -> pulumi.Output[Optional[bool]]:
        """
        is the connection ciphered with SSL (TLS)
        """
        return pulumi.get(self, "ssl")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        backend status - `active` or `inactive`
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[Optional[int]]:
        """
        used in loadbalancing algorithm
        """
        return pulumi.get(self, "weight")

