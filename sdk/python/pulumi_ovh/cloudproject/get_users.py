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
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, id=None, service_name=None, users=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        """
        The list of users of a public cloud project.
        """
        return pulumi.get(self, "users")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            id=self.id,
            service_name=self.service_name,
            users=self.users)


def get_users(service_name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    Get the list of all users of a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project_users = ovh.CloudProject.get_users(service_name="XXX")
    # Get the user ID of a previously created user with the description "S3-User"
    users = [user.user_id for user in project_users.users if user.description == "S3-User"]
    s3_user_id = users[0]
    pulumi.export("userId", s3_user_id)
    ```


    :param str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        users=pulumi.get(__ret__, 'users'))
def get_users_output(service_name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUsersResult]:
    """
    Get the list of all users of a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project_users = ovh.CloudProject.get_users(service_name="XXX")
    # Get the user ID of a previously created user with the description "S3-User"
    users = [user.user_id for user in project_users.users if user.description == "S3-User"]
    s3_user_id = users[0]
    pulumi.export("userId", s3_user_id)
    ```


    :param str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult)
    return __ret__.apply(lambda __response__: GetUsersResult(
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name'),
        users=pulumi.get(__response__, 'users')))
