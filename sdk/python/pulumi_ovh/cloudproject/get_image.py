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
    'GetImageResult',
    'AwaitableGetImageResult',
    'get_image',
    'get_image_output',
]

@pulumi.output_type
class GetImageResult:
    """
    A collection of values returned by getImage.
    """
    def __init__(__self__, creation_date=None, flavor_type=None, id=None, image_id=None, min_disk=None, min_ram=None, name=None, plan_code=None, region=None, service_name=None, size=None, status=None, tags=None, type=None, user=None, visibility=None):
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if flavor_type and not isinstance(flavor_type, str):
            raise TypeError("Expected argument 'flavor_type' to be a str")
        pulumi.set(__self__, "flavor_type", flavor_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image_id and not isinstance(image_id, str):
            raise TypeError("Expected argument 'image_id' to be a str")
        pulumi.set(__self__, "image_id", image_id)
        if min_disk and not isinstance(min_disk, float):
            raise TypeError("Expected argument 'min_disk' to be a float")
        pulumi.set(__self__, "min_disk", min_disk)
        if min_ram and not isinstance(min_ram, float):
            raise TypeError("Expected argument 'min_ram' to be a float")
        pulumi.set(__self__, "min_ram", min_ram)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if plan_code and not isinstance(plan_code, str):
            raise TypeError("Expected argument 'plan_code' to be a str")
        pulumi.set(__self__, "plan_code", plan_code)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        pulumi.set(__self__, "size", size)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user and not isinstance(user, str):
            raise TypeError("Expected argument 'user' to be a str")
        pulumi.set(__self__, "user", user)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        Image creation date
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="flavorType")
    def flavor_type(self) -> str:
        """
        Image usable only for this type of flavor if not null
        """
        return pulumi.get(self, "flavor_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Image ID
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        """
        Image ID
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="minDisk")
    def min_disk(self) -> float:
        """
        Minimum disks required to use image
        """
        return pulumi.get(self, "min_disk")

    @property
    @pulumi.getter(name="minRam")
    def min_ram(self) -> float:
        """
        Minimum RAM required to use image
        """
        return pulumi.get(self, "min_ram")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Image name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="planCode")
    def plan_code(self) -> str:
        """
        Order plan code
        """
        return pulumi.get(self, "plan_code")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        Image region
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        Public cloud project ID
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def size(self) -> float:
        """
        Image size (in GiB)
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Image status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        Tags about the image
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Image type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def user(self) -> str:
        """
        User to connect with
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        Image visibility
        """
        return pulumi.get(self, "visibility")


class AwaitableGetImageResult(GetImageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageResult(
            creation_date=self.creation_date,
            flavor_type=self.flavor_type,
            id=self.id,
            image_id=self.image_id,
            min_disk=self.min_disk,
            min_ram=self.min_ram,
            name=self.name,
            plan_code=self.plan_code,
            region=self.region,
            service_name=self.service_name,
            size=self.size,
            status=self.status,
            tags=self.tags,
            type=self.type,
            user=self.user,
            visibility=self.visibility)


def get_image(image_id: Optional[str] = None,
              service_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageResult:
    """
    Get information about an image in the given public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    image = ovh.CloudProject.get_image(image_id="<image ID>",
        service_name="<public cloud project ID>")
    ```


    :param str image_id: Image ID
    :param str service_name: Public cloud project ID
    """
    __args__ = dict()
    __args__['imageId'] = image_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getImage:getImage', __args__, opts=opts, typ=GetImageResult).value

    return AwaitableGetImageResult(
        creation_date=pulumi.get(__ret__, 'creation_date'),
        flavor_type=pulumi.get(__ret__, 'flavor_type'),
        id=pulumi.get(__ret__, 'id'),
        image_id=pulumi.get(__ret__, 'image_id'),
        min_disk=pulumi.get(__ret__, 'min_disk'),
        min_ram=pulumi.get(__ret__, 'min_ram'),
        name=pulumi.get(__ret__, 'name'),
        plan_code=pulumi.get(__ret__, 'plan_code'),
        region=pulumi.get(__ret__, 'region'),
        service_name=pulumi.get(__ret__, 'service_name'),
        size=pulumi.get(__ret__, 'size'),
        status=pulumi.get(__ret__, 'status'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'),
        user=pulumi.get(__ret__, 'user'),
        visibility=pulumi.get(__ret__, 'visibility'))
def get_image_output(image_id: Optional[pulumi.Input[str]] = None,
                     service_name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetImageResult]:
    """
    Get information about an image in the given public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    image = ovh.CloudProject.get_image(image_id="<image ID>",
        service_name="<public cloud project ID>")
    ```


    :param str image_id: Image ID
    :param str service_name: Public cloud project ID
    """
    __args__ = dict()
    __args__['imageId'] = image_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getImage:getImage', __args__, opts=opts, typ=GetImageResult)
    return __ret__.apply(lambda __response__: GetImageResult(
        creation_date=pulumi.get(__response__, 'creation_date'),
        flavor_type=pulumi.get(__response__, 'flavor_type'),
        id=pulumi.get(__response__, 'id'),
        image_id=pulumi.get(__response__, 'image_id'),
        min_disk=pulumi.get(__response__, 'min_disk'),
        min_ram=pulumi.get(__response__, 'min_ram'),
        name=pulumi.get(__response__, 'name'),
        plan_code=pulumi.get(__response__, 'plan_code'),
        region=pulumi.get(__response__, 'region'),
        service_name=pulumi.get(__response__, 'service_name'),
        size=pulumi.get(__response__, 'size'),
        status=pulumi.get(__response__, 'status'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type'),
        user=pulumi.get(__response__, 'user'),
        visibility=pulumi.get(__response__, 'visibility')))
