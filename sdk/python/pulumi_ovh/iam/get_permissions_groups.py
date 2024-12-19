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
    'GetPermissionsGroupsResult',
    'AwaitableGetPermissionsGroupsResult',
    'get_permissions_groups',
    'get_permissions_groups_output',
]

@pulumi.output_type
class GetPermissionsGroupsResult:
    """
    A collection of values returned by getPermissionsGroups.
    """
    def __init__(__self__, id=None, urns=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if urns and not isinstance(urns, list):
            raise TypeError("Expected argument 'urns' to be a list")
        pulumi.set(__self__, "urns", urns)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def urns(self) -> Sequence[str]:
        return pulumi.get(self, "urns")


class AwaitableGetPermissionsGroupsResult(GetPermissionsGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPermissionsGroupsResult(
            id=self.id,
            urns=self.urns)


def get_permissions_groups(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPermissionsGroupsResult:
    """
    Use this data source to retrieve an IAM permissions group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    website = ovh.Iam.get_permissions_group(urn="urn:v1:eu:permissionsGroup:ovh:controlPanelAccess")
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Iam/getPermissionsGroups:getPermissionsGroups', __args__, opts=opts, typ=GetPermissionsGroupsResult).value

    return AwaitableGetPermissionsGroupsResult(
        id=pulumi.get(__ret__, 'id'),
        urns=pulumi.get(__ret__, 'urns'))
def get_permissions_groups_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPermissionsGroupsResult]:
    """
    Use this data source to retrieve an IAM permissions group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    website = ovh.Iam.get_permissions_group(urn="urn:v1:eu:permissionsGroup:ovh:controlPanelAccess")
    ```
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Iam/getPermissionsGroups:getPermissionsGroups', __args__, opts=opts, typ=GetPermissionsGroupsResult)
    return __ret__.apply(lambda __response__: GetPermissionsGroupsResult(
        id=pulumi.get(__response__, 'id'),
        urns=pulumi.get(__response__, 'urns')))
