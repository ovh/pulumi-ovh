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
from . import outputs

__all__ = [
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    """
    A collection of values returned by getInstance.
    """
    def __init__(__self__, addresses=None, attached_volumes=None, availability_zone=None, flavor_id=None, flavor_name=None, id=None, image_id=None, instance_id=None, name=None, region=None, service_name=None, ssh_key=None, task_state=None):
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if attached_volumes and not isinstance(attached_volumes, list):
            raise TypeError("Expected argument 'attached_volumes' to be a list")
        pulumi.set(__self__, "attached_volumes", attached_volumes)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if flavor_id and not isinstance(flavor_id, str):
            raise TypeError("Expected argument 'flavor_id' to be a str")
        pulumi.set(__self__, "flavor_id", flavor_id)
        if flavor_name and not isinstance(flavor_name, str):
            raise TypeError("Expected argument 'flavor_name' to be a str")
        pulumi.set(__self__, "flavor_name", flavor_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if ssh_key and not isinstance(ssh_key, str):
            raise TypeError("Expected argument 'ssh_key' to be a str")
        pulumi.set(__self__, "ssh_key", ssh_key)
        if task_state and not isinstance(task_state, str):
            raise TypeError("Expected argument 'task_state' to be a str")
        pulumi.set(__self__, "task_state", task_state)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence['outputs.GetInstanceAddressResult']:
        """
        Instance IP addresses
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter(name="attachedVolumes")
    def attached_volumes(self) -> Sequence['outputs.GetInstanceAttachedVolumeResult']:
        """
        Volumes attached to the instance
        """
        return pulumi.get(self, "attached_volumes")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        Availability zone of the instance
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="flavorId")
    def flavor_id(self) -> str:
        """
        Flavor id
        """
        return pulumi.get(self, "flavor_id")

    @property
    @pulumi.getter(name="flavorName")
    def flavor_name(self) -> str:
        """
        Flavor name
        """
        return pulumi.get(self, "flavor_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        Image id
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Instance name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="sshKey")
    def ssh_key(self) -> str:
        """
        SSH Keypair
        """
        return pulumi.get(self, "ssh_key")

    @property
    @pulumi.getter(name="taskState")
    def task_state(self) -> str:
        """
        Instance task state
        """
        return pulumi.get(self, "task_state")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            addresses=self.addresses,
            attached_volumes=self.attached_volumes,
            availability_zone=self.availability_zone,
            flavor_id=self.flavor_id,
            flavor_name=self.flavor_name,
            id=self.id,
            image_id=self.image_id,
            instance_id=self.instance_id,
            name=self.name,
            region=self.region,
            service_name=self.service_name,
            ssh_key=self.ssh_key,
            task_state=self.task_state)


def get_instance(instance_id: Optional[str] = None,
                 region: Optional[str] = None,
                 service_name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    **This datasource uses a Beta API** Use this data source to get the instance of a public cloud project.

    ## Example Usage

    To get information of an instance:

    ```python
    import pulumi
    import pulumi_ovh as ovh

    instance = ovh.CloudProject.get_instance(instance_id="ZZZZZ",
        region="XXXX",
        service_name="YYYY")
    ```


    :param str instance_id: Instance id
    :param str region: Instance region
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getInstance:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        addresses=pulumi.get(__ret__, 'addresses'),
        attached_volumes=pulumi.get(__ret__, 'attached_volumes'),
        availability_zone=pulumi.get(__ret__, 'availability_zone'),
        flavor_id=pulumi.get(__ret__, 'flavor_id'),
        flavor_name=pulumi.get(__ret__, 'flavor_name'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name=pulumi.get(__ret__, 'name'),
        region=pulumi.get(__ret__, 'region'),
        service_name=pulumi.get(__ret__, 'service_name'),
        ssh_key=pulumi.get(__ret__, 'ssh_key'),
        task_state=pulumi.get(__ret__, 'task_state'))
def get_instance_output(instance_id: Optional[pulumi.Input[str]] = None,
                        region: Optional[pulumi.Input[str]] = None,
                        service_name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetInstanceResult]:
    """
    **This datasource uses a Beta API** Use this data source to get the instance of a public cloud project.

    ## Example Usage

    To get information of an instance:

    ```python
    import pulumi
    import pulumi_ovh as ovh

    instance = ovh.CloudProject.get_instance(instance_id="ZZZZZ",
        region="XXXX",
        service_name="YYYY")
    ```


    :param str instance_id: Instance id
    :param str region: Instance region
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['region'] = region
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getInstance:getInstance', __args__, opts=opts, typ=GetInstanceResult)
    return __ret__.apply(lambda __response__: GetInstanceResult(
        addresses=pulumi.get(__response__, 'addresses'),
        attached_volumes=pulumi.get(__response__, 'attached_volumes'),
        availability_zone=pulumi.get(__response__, 'availability_zone'),
        flavor_id=pulumi.get(__response__, 'flavor_id'),
        flavor_name=pulumi.get(__response__, 'flavor_name'),
        id=pulumi.get(__response__, 'id'),
        image_id=pulumi.get(__response__, 'image_id'),
        instance_id=pulumi.get(__response__, 'instance_id'),
        name=pulumi.get(__response__, 'name'),
        region=pulumi.get(__response__, 'region'),
        service_name=pulumi.get(__response__, 'service_name'),
        ssh_key=pulumi.get(__response__, 'ssh_key'),
        task_state=pulumi.get(__response__, 'task_state')))
