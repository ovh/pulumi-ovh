# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetCephResult',
    'AwaitableGetCephResult',
    'get_ceph',
    'get_ceph_output',
]

@pulumi.output_type
class GetCephResult:
    """
    A collection of values returned by getCeph.
    """
    def __init__(__self__, ceph_mons=None, ceph_version=None, crush_tunables=None, id=None, label=None, region=None, service_name=None, size=None, state=None, status=None, urn=None):
        if ceph_mons and not isinstance(ceph_mons, list):
            raise TypeError("Expected argument 'ceph_mons' to be a list")
        pulumi.set(__self__, "ceph_mons", ceph_mons)
        if ceph_version and not isinstance(ceph_version, str):
            raise TypeError("Expected argument 'ceph_version' to be a str")
        pulumi.set(__self__, "ceph_version", ceph_version)
        if crush_tunables and not isinstance(crush_tunables, str):
            raise TypeError("Expected argument 'crush_tunables' to be a str")
        pulumi.set(__self__, "crush_tunables", crush_tunables)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        pulumi.set(__self__, "label", label)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        pulumi.set(__self__, "size", size)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if urn and not isinstance(urn, str):
            raise TypeError("Expected argument 'urn' to be a str")
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="cephMons")
    def ceph_mons(self) -> Sequence[str]:
        """
        list of CEPH monitors IPs
        """
        return pulumi.get(self, "ceph_mons")

    @property
    @pulumi.getter(name="cephVersion")
    def ceph_version(self) -> str:
        """
        CEPH cluster version
        """
        return pulumi.get(self, "ceph_version")

    @property
    @pulumi.getter(name="crushTunables")
    def crush_tunables(self) -> str:
        """
        CRUSH algorithm settings. Possible values
        * OPTIMAL
        * DEFAULT
        * LEGACY
        * BOBTAIL
        * ARGONAUT
        * FIREFLY
        * HAMMER
        * JEWEL
        """
        return pulumi.get(self, "crush_tunables")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def label(self) -> str:
        """
        CEPH cluster label
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        cluster region
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def size(self) -> float:
        """
        Cluster size in TB
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        the state of the cluster
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        the status of the service
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        URN of the CEPH instance
        """
        return pulumi.get(self, "urn")


class AwaitableGetCephResult(GetCephResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCephResult(
            ceph_mons=self.ceph_mons,
            ceph_version=self.ceph_version,
            crush_tunables=self.crush_tunables,
            id=self.id,
            label=self.label,
            region=self.region,
            service_name=self.service_name,
            size=self.size,
            state=self.state,
            status=self.status,
            urn=self.urn)


def get_ceph(ceph_version: Optional[str] = None,
             service_name: Optional[str] = None,
             status: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCephResult:
    """
    Use this data source to retrieve information about a dedicated CEPH.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_ceph = ovh.Dedicated.get_ceph(service_name="XXXXXX")
    ```


    :param str ceph_version: CEPH cluster version
    :param str service_name: The service name of the dedicated CEPH cluster.
    :param str status: the status of the service
    """
    __args__ = dict()
    __args__['cephVersion'] = ceph_version
    __args__['serviceName'] = service_name
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dedicated/getCeph:getCeph', __args__, opts=opts, typ=GetCephResult).value

    return AwaitableGetCephResult(
        ceph_mons=pulumi.get(__ret__, 'ceph_mons'),
        ceph_version=pulumi.get(__ret__, 'ceph_version'),
        crush_tunables=pulumi.get(__ret__, 'crush_tunables'),
        id=pulumi.get(__ret__, 'id'),
        label=pulumi.get(__ret__, 'label'),
        region=pulumi.get(__ret__, 'region'),
        service_name=pulumi.get(__ret__, 'service_name'),
        size=pulumi.get(__ret__, 'size'),
        state=pulumi.get(__ret__, 'state'),
        status=pulumi.get(__ret__, 'status'),
        urn=pulumi.get(__ret__, 'urn'))


@_utilities.lift_output_func(get_ceph)
def get_ceph_output(ceph_version: Optional[pulumi.Input[Optional[str]]] = None,
                    service_name: Optional[pulumi.Input[str]] = None,
                    status: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCephResult]:
    """
    Use this data source to retrieve information about a dedicated CEPH.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    my_ceph = ovh.Dedicated.get_ceph(service_name="XXXXXX")
    ```


    :param str ceph_version: CEPH cluster version
    :param str service_name: The service name of the dedicated CEPH cluster.
    :param str status: the status of the service
    """
    ...
