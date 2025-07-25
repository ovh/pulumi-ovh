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

__all__ = [
    'GetKafkaAclsResult',
    'AwaitableGetKafkaAclsResult',
    'get_kafka_acls',
    'get_kafka_acls_output',
]

@pulumi.output_type
class GetKafkaAclsResult:
    """
    A collection of values returned by getKafkaAcls.
    """
    def __init__(__self__, acl_ids=None, cluster_id=None, id=None, service_name=None):
        if acl_ids and not isinstance(acl_ids, list):
            raise TypeError("Expected argument 'acl_ids' to be a list")
        pulumi.set(__self__, "acl_ids", acl_ids)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="aclIds")
    def acl_ids(self) -> Sequence[builtins.str]:
        """
        The list of ACLs ids of the kafka cluster associated with the project.
        """
        return pulumi.get(self, "acl_ids")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> builtins.str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

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
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetKafkaAclsResult(GetKafkaAclsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKafkaAclsResult(
            acl_ids=self.acl_ids,
            cluster_id=self.cluster_id,
            id=self.id,
            service_name=self.service_name)


def get_kafka_acls(cluster_id: Optional[builtins.str] = None,
                   service_name: Optional[builtins.str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKafkaAclsResult:
    """
    Use this data source to get the list of ACLs of a kafka cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    acls = ovh.CloudProjectDatabase.get_kafka_acls(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("aclIds", acls.acl_ids)
    ```


    :param builtins.str cluster_id: Cluster ID
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getKafkaAcls:getKafkaAcls', __args__, opts=opts, typ=GetKafkaAclsResult).value

    return AwaitableGetKafkaAclsResult(
        acl_ids=pulumi.get(__ret__, 'acl_ids'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_kafka_acls_output(cluster_id: Optional[pulumi.Input[builtins.str]] = None,
                          service_name: Optional[pulumi.Input[builtins.str]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetKafkaAclsResult]:
    """
    Use this data source to get the list of ACLs of a kafka cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    acls = ovh.CloudProjectDatabase.get_kafka_acls(service_name="XXX",
        cluster_id="YYY")
    pulumi.export("aclIds", acls.acl_ids)
    ```


    :param builtins.str cluster_id: Cluster ID
    :param builtins.str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProjectDatabase/getKafkaAcls:getKafkaAcls', __args__, opts=opts, typ=GetKafkaAclsResult)
    return __ret__.apply(lambda __response__: GetKafkaAclsResult(
        acl_ids=pulumi.get(__response__, 'acl_ids'),
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name')))
