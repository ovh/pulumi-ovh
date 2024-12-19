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

__all__ = [
    'GetGatewayInterfaceResult',
    'AwaitableGetGatewayInterfaceResult',
    'get_gateway_interface',
    'get_gateway_interface_output',
]

@pulumi.output_type
class GetGatewayInterfaceResult:
    """
    A collection of values returned by getGatewayInterface.
    """
    def __init__(__self__, id=None, interface_id=None, ip=None, network_id=None, region=None, service_name=None, subnet_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if network_id and not isinstance(network_id, str):
            raise TypeError("Expected argument 'network_id' to be a str")
        pulumi.set(__self__, "network_id", network_id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        ID of the gateway
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> str:
        """
        ID of the interface
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        IP of the interface
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> str:
        """
        Network ID of the interface
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Region of the gateway
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        ID of the cloud project
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        ID of the subnet to add
        """
        return pulumi.get(self, "subnet_id")


class AwaitableGetGatewayInterfaceResult(GetGatewayInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayInterfaceResult(
            id=self.id,
            interface_id=self.interface_id,
            ip=self.ip,
            network_id=self.network_id,
            region=self.region,
            service_name=self.service_name,
            subnet_id=self.subnet_id)


def get_gateway_interface(id: Optional[str] = None,
                          interface_id: Optional[str] = None,
                          region: Optional[str] = None,
                          service_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayInterfaceResult:
    """
    Use this datasource to get a public cloud project Gateway Interface.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    interface = ovh.CloudProject.get_gateway_interface(id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        interface_id="yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyy",
        region="GRA11",
        service_name="XXXXXX")
    ```


    :param str id: ID of the gateway
    :param str interface_id: ID of the interface
    :param str region: Region of the gateway
    :param str service_name: ID of the cloud project
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['interfaceId'] = interface_id
    __args__['region'] = region
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getGatewayInterface:getGatewayInterface', __args__, opts=opts, typ=GetGatewayInterfaceResult).value

    return AwaitableGetGatewayInterfaceResult(
        id=pulumi.get(__ret__, 'id'),
        interface_id=pulumi.get(__ret__, 'interface_id'),
        ip=pulumi.get(__ret__, 'ip'),
        network_id=pulumi.get(__ret__, 'network_id'),
        region=pulumi.get(__ret__, 'region'),
        service_name=pulumi.get(__ret__, 'service_name'),
        subnet_id=pulumi.get(__ret__, 'subnet_id'))
def get_gateway_interface_output(id: Optional[pulumi.Input[str]] = None,
                                 interface_id: Optional[pulumi.Input[str]] = None,
                                 region: Optional[pulumi.Input[str]] = None,
                                 service_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetGatewayInterfaceResult]:
    """
    Use this datasource to get a public cloud project Gateway Interface.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    interface = ovh.CloudProject.get_gateway_interface(id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        interface_id="yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyy",
        region="GRA11",
        service_name="XXXXXX")
    ```


    :param str id: ID of the gateway
    :param str interface_id: ID of the interface
    :param str region: Region of the gateway
    :param str service_name: ID of the cloud project
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['interfaceId'] = interface_id
    __args__['region'] = region
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getGatewayInterface:getGatewayInterface', __args__, opts=opts, typ=GetGatewayInterfaceResult)
    return __ret__.apply(lambda __response__: GetGatewayInterfaceResult(
        id=pulumi.get(__response__, 'id'),
        interface_id=pulumi.get(__response__, 'interface_id'),
        ip=pulumi.get(__response__, 'ip'),
        network_id=pulumi.get(__response__, 'network_id'),
        region=pulumi.get(__response__, 'region'),
        service_name=pulumi.get(__response__, 'service_name'),
        subnet_id=pulumi.get(__response__, 'subnet_id')))
