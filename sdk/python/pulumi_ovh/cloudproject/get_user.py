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
from . import outputs

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, creation_date=None, description=None, id=None, roles=None, service_name=None, status=None, user_id=None, username=None):
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> builtins.str:
        """
        the date the user was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        description of the role
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def roles(self) -> Sequence['outputs.GetUserRoleResult']:
        """
        A list of roles associated with the user.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        the status of the user. should be normally set to 'ok'.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> builtins.str:
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter
    def username(self) -> builtins.str:
        """
        the username generated for the user. This username can be used with the Openstack API.
        """
        return pulumi.get(self, "username")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            creation_date=self.creation_date,
            description=self.description,
            id=self.id,
            roles=self.roles,
            service_name=self.service_name,
            status=self.status,
            user_id=self.user_id,
            username=self.username)


def get_user(service_name: Optional[builtins.str] = None,
             user_id: Optional[builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Get the user details of a previously created public cloud project user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project_users = ovh.CloudProject.get_users(service_name="XXX")
    # Get the user ID of a previously created user with the description "S3-User"
    users = [user.user_id for user in project_users.users if user.description == "S3-User"]
    s3_user_id = users[0]
    my_user = ovh.CloudProject.get_user(service_name=project_users.service_name,
        user_id=s3_user_id)
    ```


    :param builtins.str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param builtins.str user_id: The ID of a public cloud project's user.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        creation_date=pulumi.get(__ret__, 'creation_date'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        roles=pulumi.get(__ret__, 'roles'),
        service_name=pulumi.get(__ret__, 'service_name'),
        status=pulumi.get(__ret__, 'status'),
        user_id=pulumi.get(__ret__, 'user_id'),
        username=pulumi.get(__ret__, 'username'))
def get_user_output(service_name: Optional[pulumi.Input[builtins.str]] = None,
                    user_id: Optional[pulumi.Input[builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserResult]:
    """
    Get the user details of a previously created public cloud project user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    project_users = ovh.CloudProject.get_users(service_name="XXX")
    # Get the user ID of a previously created user with the description "S3-User"
    users = [user.user_id for user in project_users.users if user.description == "S3-User"]
    s3_user_id = users[0]
    my_user = ovh.CloudProject.get_user(service_name=project_users.service_name,
        user_id=s3_user_id)
    ```


    :param builtins.str service_name: The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    :param builtins.str user_id: The ID of a public cloud project's user.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['userId'] = user_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProject/getUser:getUser', __args__, opts=opts, typ=GetUserResult)
    return __ret__.apply(lambda __response__: GetUserResult(
        creation_date=pulumi.get(__response__, 'creation_date'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        roles=pulumi.get(__response__, 'roles'),
        service_name=pulumi.get(__response__, 'service_name'),
        status=pulumi.get(__response__, 'status'),
        user_id=pulumi.get(__response__, 'user_id'),
        username=pulumi.get(__response__, 'username')))
