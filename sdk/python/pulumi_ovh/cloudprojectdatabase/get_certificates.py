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
    'GetCertificatesResult',
    'AwaitableGetCertificatesResult',
    'get_certificates',
    'get_certificates_output',
]

@pulumi.output_type
class GetCertificatesResult:
    """
    A collection of values returned by getCertificates.
    """
    def __init__(__self__, ca=None, cluster_id=None, engine=None, id=None, service_name=None):
        if ca and not isinstance(ca, str):
            raise TypeError("Expected argument 'ca' to be a str")
        pulumi.set(__self__, "ca", ca)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        pulumi.set(__self__, "engine", engine)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def ca(self) -> str:
        """
        CA certificate used for the service.
        """
        return pulumi.get(self, "ca")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def engine(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "engine")

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
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")


class AwaitableGetCertificatesResult(GetCertificatesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCertificatesResult(
            ca=self.ca,
            cluster_id=self.cluster_id,
            engine=self.engine,
            id=self.id,
            service_name=self.service_name)


def get_certificates(cluster_id: Optional[str] = None,
                     engine: Optional[str] = None,
                     service_name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCertificatesResult:
    """
    Use this data source to get information about certificates of a cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    certificates = ovh.CloudProjectDatabase.get_certificates(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ")
    pulumi.export("certificatesCa", certificates.ca)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want database information. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProjectDatabase/getCertificates:getCertificates', __args__, opts=opts, typ=GetCertificatesResult).value

    return AwaitableGetCertificatesResult(
        ca=pulumi.get(__ret__, 'ca'),
        cluster_id=pulumi.get(__ret__, 'cluster_id'),
        engine=pulumi.get(__ret__, 'engine'),
        id=pulumi.get(__ret__, 'id'),
        service_name=pulumi.get(__ret__, 'service_name'))
def get_certificates_output(cluster_id: Optional[pulumi.Input[str]] = None,
                            engine: Optional[pulumi.Input[str]] = None,
                            service_name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCertificatesResult]:
    """
    Use this data source to get information about certificates of a cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    certificates = ovh.CloudProjectDatabase.get_certificates(service_name="XXX",
        engine="YYY",
        cluster_id="ZZZ")
    pulumi.export("certificatesCa", certificates.ca)
    ```


    :param str cluster_id: Cluster ID
    :param str engine: The engine of the database cluster you want database information. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
    :param str service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['engine'] = engine
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:CloudProjectDatabase/getCertificates:getCertificates', __args__, opts=opts, typ=GetCertificatesResult)
    return __ret__.apply(lambda __response__: GetCertificatesResult(
        ca=pulumi.get(__response__, 'ca'),
        cluster_id=pulumi.get(__response__, 'cluster_id'),
        engine=pulumi.get(__response__, 'engine'),
        id=pulumi.get(__response__, 'id'),
        service_name=pulumi.get(__response__, 'service_name')))
