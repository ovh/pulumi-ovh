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
    'GetIdentityUsersResult',
    'AwaitableGetIdentityUsersResult',
    'get_identity_users',
    'get_identity_users_output',
]

@pulumi.output_type
class GetIdentityUsersResult:
    """
    A collection of values returned by getIdentityUsers.
    """
    def __init__(__self__, id=None, users=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
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
    @pulumi.getter
    def users(self) -> Sequence[str]:
        return pulumi.get(self, "users")


class AwaitableGetIdentityUsersResult(GetIdentityUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdentityUsersResult(
            id=self.id,
            users=self.users)


def get_identity_users(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdentityUsersResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Me/getIdentityUsers:getIdentityUsers', __args__, opts=opts, typ=GetIdentityUsersResult).value

    return AwaitableGetIdentityUsersResult(
        id=pulumi.get(__ret__, 'id'),
        users=pulumi.get(__ret__, 'users'))
def get_identity_users_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIdentityUsersResult]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Me/getIdentityUsers:getIdentityUsers', __args__, opts=opts, typ=GetIdentityUsersResult)
    return __ret__.apply(lambda __response__: GetIdentityUsersResult(
        id=pulumi.get(__response__, 'id'),
        users=pulumi.get(__response__, 'users')))
