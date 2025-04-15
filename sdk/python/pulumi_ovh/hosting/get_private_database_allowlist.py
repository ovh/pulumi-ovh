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

__all__ = [
    'GetPrivateDatabaseAllowlistResult',
    'AwaitableGetPrivateDatabaseAllowlistResult',
    'get_private_database_allowlist',
    'get_private_database_allowlist_output',
]

@pulumi.output_type
class GetPrivateDatabaseAllowlistResult:
    """
    A collection of values returned by getPrivateDatabaseAllowlist.
    """
    def __init__(__self__, creation_date=None, id=None, ip=None, last_update=None, name=None, service=None, service_name=None, sftp=None, status=None):
        if creation_date and not isinstance(creation_date, str):
            raise TypeError("Expected argument 'creation_date' to be a str")
        pulumi.set(__self__, "creation_date", creation_date)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if last_update and not isinstance(last_update, str):
            raise TypeError("Expected argument 'last_update' to be a str")
        pulumi.set(__self__, "last_update", last_update)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service and not isinstance(service, bool):
            raise TypeError("Expected argument 'service' to be a bool")
        pulumi.set(__self__, "service", service)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if sftp and not isinstance(sftp, bool):
            raise TypeError("Expected argument 'sftp' to be a bool")
        pulumi.set(__self__, "sftp", sftp)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> builtins.str:
        """
        Creation date of the database
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> Optional[builtins.str]:
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="lastUpdate")
    def last_update(self) -> builtins.str:
        """
        The last update date of this whitelist
        """
        return pulumi.get(self, "last_update")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Custom name for your Whitelisted IP
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def service(self) -> builtins.bool:
        """
        Authorize this IP to access service port
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def sftp(self) -> builtins.bool:
        """
        Authorize this IP to access SFTP port
        """
        return pulumi.get(self, "sftp")

    @property
    @pulumi.getter
    def status(self) -> builtins.str:
        """
        Whitelist status
        """
        return pulumi.get(self, "status")


class AwaitableGetPrivateDatabaseAllowlistResult(GetPrivateDatabaseAllowlistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateDatabaseAllowlistResult(
            creation_date=self.creation_date,
            id=self.id,
            ip=self.ip,
            last_update=self.last_update,
            name=self.name,
            service=self.service,
            service_name=self.service_name,
            sftp=self.sftp,
            status=self.status)


def get_private_database_allowlist(ip: Optional[builtins.str] = None,
                                   service_name: Optional[builtins.str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateDatabaseAllowlistResult:
    """
    Use this data source to retrieve information about an hosting privatedatabase whitelist.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    whitelist = ovh.Hosting.get_private_database_allowlist(service_name="XXXXXX",
        ip="XXXXXX")
    ```


    :param builtins.str ip: The whitelisted IP in your instance
    :param builtins.str service_name: The internal name of your private database
    """
    __args__ = dict()
    __args__['ip'] = ip
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist', __args__, opts=opts, typ=GetPrivateDatabaseAllowlistResult).value

    return AwaitableGetPrivateDatabaseAllowlistResult(
        creation_date=pulumi.get(__ret__, 'creation_date'),
        id=pulumi.get(__ret__, 'id'),
        ip=pulumi.get(__ret__, 'ip'),
        last_update=pulumi.get(__ret__, 'last_update'),
        name=pulumi.get(__ret__, 'name'),
        service=pulumi.get(__ret__, 'service'),
        service_name=pulumi.get(__ret__, 'service_name'),
        sftp=pulumi.get(__ret__, 'sftp'),
        status=pulumi.get(__ret__, 'status'))
def get_private_database_allowlist_output(ip: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                          service_name: Optional[pulumi.Input[builtins.str]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPrivateDatabaseAllowlistResult]:
    """
    Use this data source to retrieve information about an hosting privatedatabase whitelist.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    whitelist = ovh.Hosting.get_private_database_allowlist(service_name="XXXXXX",
        ip="XXXXXX")
    ```


    :param builtins.str ip: The whitelisted IP in your instance
    :param builtins.str service_name: The internal name of your private database
    """
    __args__ = dict()
    __args__['ip'] = ip
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist', __args__, opts=opts, typ=GetPrivateDatabaseAllowlistResult)
    return __ret__.apply(lambda __response__: GetPrivateDatabaseAllowlistResult(
        creation_date=pulumi.get(__response__, 'creation_date'),
        id=pulumi.get(__response__, 'id'),
        ip=pulumi.get(__response__, 'ip'),
        last_update=pulumi.get(__response__, 'last_update'),
        name=pulumi.get(__response__, 'name'),
        service=pulumi.get(__response__, 'service'),
        service_name=pulumi.get(__response__, 'service_name'),
        sftp=pulumi.get(__response__, 'sftp'),
        status=pulumi.get(__response__, 'status')))
