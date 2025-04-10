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
    'GetPrivateDatabaseUserResult',
    'AwaitableGetPrivateDatabaseUserResult',
    'get_private_database_user',
    'get_private_database_user_output',
]

@pulumi.output_type
class GetPrivateDatabaseUserResult:
    """
    A collection of values returned by getPrivateDatabaseUser.
    """
    def __init__(__self__, creation_date=None, databases=None, id=None, service_name=None, user_name=None):
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> builtins.str:
        """
        Creation date of the database
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def databases(self) -> Sequence['outputs.GetPrivateDatabaseUserDatabaseResult']:
        """
        Users granted to this database
        """
        return pulumi.get(self, "databases")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> builtins.str:
        return pulumi.get(self, "user_name")


class AwaitableGetPrivateDatabaseUserResult(GetPrivateDatabaseUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateDatabaseUserResult(
            creation_date=self.creation_date,
            databases=self.databases,
            id=self.id,
            service_name=self.service_name,
            user_name=self.user_name)


def get_private_database_user(service_name: Optional[builtins.str] = None,
                              user_name: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateDatabaseUserResult:
    """
    Use this data source to retrieve information about an hosting privatedatabase user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    user = ovh.Hosting.get_private_database_user(service_name="XXXXXX",
        user_name="XXXXXX")
    ```


    :param builtins.str service_name: The internal name of your private database
    :param builtins.str user_name: User name
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['userName'] = user_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Hosting/getPrivateDatabaseUser:getPrivateDatabaseUser', __args__, opts=opts, typ=GetPrivateDatabaseUserResult).value

    return AwaitableGetPrivateDatabaseUserResult(
        creation_date=pulumi.get(__ret__, 'creation_date'),
        databases=pulumi.get(__ret__, 'databases'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'),
        user_name=pulumi.get(__ret__, 'user_name'))
def get_private_database_user_output(service_name: Optional[pulumi.Input[builtins.str]] = None,
                                     user_name: Optional[pulumi.Input[builtins.str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPrivateDatabaseUserResult]:
    """
    Use this data source to retrieve information about an hosting privatedatabase user.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    user = ovh.Hosting.get_private_database_user(service_name="XXXXXX",
        user_name="XXXXXX")
    ```


    :param builtins.str service_name: The internal name of your private database
    :param builtins.str user_name: User name
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    __args__['userName'] = user_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Hosting/getPrivateDatabaseUser:getPrivateDatabaseUser', __args__, opts=opts, typ=GetPrivateDatabaseUserResult)
    return __ret__.apply(lambda __response__: GetPrivateDatabaseUserResult(
        creation_date=pulumi.get(__response__, 'creation_date'),
        databases=pulumi.get(__response__, 'databases'),
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name'),
        user_name=pulumi.get(__response__, 'user_name')))
