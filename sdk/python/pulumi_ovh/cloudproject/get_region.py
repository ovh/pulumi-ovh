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

__all__ = [
    'GetRegionResult',
    'AwaitableGetRegionResult',
    'get_region',
    'get_region_output',
]

@pulumi.output_type
class GetRegionResult:
    """
    A collection of values returned by getRegion.
    """
    def __init__(__self__, continent_code=None, datacenter_location=None, id=None, name=None, service_name=None, services=None):
        if continent_code and not isinstance(continent_code, str):
            raise TypeError("Expected argument 'continent_code' to be a str")
        pulumi.set(__self__, "continent_code", continent_code)
        if datacenter_location and not isinstance(datacenter_location, str):
            raise TypeError("Expected argument 'datacenter_location' to be a str")
        pulumi.set(__self__, "datacenter_location", datacenter_location)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        pulumi.set(__self__, "services", services)

    @property
    @pulumi.getter(name="continentCode")
    def continent_code(self) -> builtins.str:
        """
        the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...
        """
        return pulumi.get(self, "continent_code")

    @property
    @pulumi.getter(name="datacenterLocation")
    def datacenter_location(self) -> builtins.str:
        """
        The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"
        """
        return pulumi.get(self, "datacenter_location")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        the name of the public cloud service
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def services(self) -> Sequence['outputs.GetRegionServiceResult']:
        """
        The list of public cloud services running within the region
        """
        return pulumi.get(self, "services")


class AwaitableGetRegionResult(GetRegionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionResult(
            continent_code=self.continent_code,
            datacenter_location=self.datacenter_location,
            id=self.id,
            name=self.name,
            service_name=self.service_name,
            services=self.services)


def get_region(name: Optional[builtins.str] = None,
               service_name: Optional[builtins.str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionResult:
    """
    Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    gra1 = ovh.CloudProject.get_region(service_name="XXXXXX",
        name="GRA1")
    ```


    :param builtins.str name: The name of the region associated with the public cloud project.
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getRegion:getRegion', __args__, opts=opts, typ=GetRegionResult).value

    return AwaitableGetRegionResult(
        continent_code=pulumi.get(__ret__, 'continent_code'),
        datacenter_location=pulumi.get(__ret__, 'datacenter_location'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        services=pulumi.get(__ret__, 'services'))
def get_region_output(name: Optional[pulumi.Input[builtins.str]] = None,
                      service_name: Optional[pulumi.Input[builtins.str]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRegionResult]:
    """
    Use this data source to retrieve information about a region associated with a public cloud project. The region must be associated with the project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    gra1 = ovh.CloudProject.get_region(service_name="XXXXXX",
        name="GRA1")
    ```


    :param builtins.str name: The name of the region associated with the public cloud project.
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getRegion:getRegion', __args__, opts=opts, typ=GetRegionResult)
    return __ret__.apply(lambda __response__: GetRegionResult(
        continent_code=pulumi.get(__response__, 'continent_code'),
        datacenter_location=pulumi.get(__response__, 'datacenter_location'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        service_name=pulumi.get(__response__, 'service_name'),
        services=pulumi.get(__response__, 'services')))
