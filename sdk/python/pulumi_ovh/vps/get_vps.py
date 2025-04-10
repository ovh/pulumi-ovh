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
    'GetVpsResult',
    'AwaitableGetVpsResult',
    'get_vps',
    'get_vps_output',
]

@pulumi.output_type
class GetVpsResult:
    """
    A collection of values returned by getVps.
    """
    def __init__(__self__, vps_urn=None, cluster=None, datacenter=None, displayname=None, id=None, ips=None, keymap=None, memory=None, model=None, name=None, netbootmode=None, offertype=None, service_name=None, slamonitoring=None, state=None, type=None, vcore=None, zone=None):
        if vps_urn and not isinstance(vps_urn, str):
            raise TypeError("Expected argument 'vps_urn' to be a str")
        pulumi.set(__self__, "vps_urn", vps_urn)
        if cluster and not isinstance(cluster, str):
            raise TypeError("Expected argument 'cluster' to be a str")
        pulumi.set(__self__, "cluster", cluster)
        if datacenter and not isinstance(datacenter, dict):
            raise TypeError("Expected argument 'datacenter' to be a dict")
        pulumi.set(__self__, "datacenter", datacenter)
        if displayname and not isinstance(displayname, str):
            raise TypeError("Expected argument 'displayname' to be a str")
        pulumi.set(__self__, "displayname", displayname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ips and not isinstance(ips, list):
            raise TypeError("Expected argument 'ips' to be a list")
        pulumi.set(__self__, "ips", ips)
        if keymap and not isinstance(keymap, str):
            raise TypeError("Expected argument 'keymap' to be a str")
        pulumi.set(__self__, "keymap", keymap)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if model and not isinstance(model, dict):
            raise TypeError("Expected argument 'model' to be a dict")
        pulumi.set(__self__, "model", model)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if netbootmode and not isinstance(netbootmode, str):
            raise TypeError("Expected argument 'netbootmode' to be a str")
        pulumi.set(__self__, "netbootmode", netbootmode)
        if offertype and not isinstance(offertype, str):
            raise TypeError("Expected argument 'offertype' to be a str")
        pulumi.set(__self__, "offertype", offertype)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if slamonitoring and not isinstance(slamonitoring, bool):
            raise TypeError("Expected argument 'slamonitoring' to be a bool")
        pulumi.set(__self__, "slamonitoring", slamonitoring)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vcore and not isinstance(vcore, int):
            raise TypeError("Expected argument 'vcore' to be a int")
        pulumi.set(__self__, "vcore", vcore)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="VpsURN")
    def vps_urn(self) -> builtins.str:
        """
        The URN of the vps
        """
        return pulumi.get(self, "vps_urn")

    @property
    @pulumi.getter
    def cluster(self) -> builtins.str:
        """
        The OVHcloud cluster the vps is in
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter
    def datacenter(self) -> Mapping[str, builtins.str]:
        """
        The datacenter in which the vps is located
        """
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter
    def displayname(self) -> builtins.str:
        """
        The displayed name in the OVHcloud web admin
        """
        return pulumi.get(self, "displayname")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ips(self) -> Sequence[builtins.str]:
        """
        The list of IPs addresses attached to the vps
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter
    def keymap(self) -> builtins.str:
        """
        The keymap for the ip kvm, valid values "", "fr", "us"
        """
        return pulumi.get(self, "keymap")

    @property
    @pulumi.getter
    def memory(self) -> builtins.int:
        """
        The amount of memory in MB of the vps.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def model(self) -> Mapping[str, builtins.str]:
        """
        A dict describing the type of vps.
        """
        return pulumi.get(self, "model")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netbootmode(self) -> builtins.str:
        """
        The source of the boot kernel
        """
        return pulumi.get(self, "netbootmode")

    @property
    @pulumi.getter
    def offertype(self) -> builtins.str:
        """
        The type of offer (ssd, cloud, classic)
        """
        return pulumi.get(self, "offertype")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> builtins.str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def slamonitoring(self) -> builtins.bool:
        """
        A boolean to indicate if OVHcloud SLA monitoring is active.
        """
        return pulumi.get(self, "slamonitoring")

    @property
    @pulumi.getter
    def state(self) -> builtins.str:
        """
        The state of the vps
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of server
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vcore(self) -> builtins.int:
        """
        The number of vcore of the vps
        """
        return pulumi.get(self, "vcore")

    @property
    @pulumi.getter
    def zone(self) -> builtins.str:
        """
        The OVHcloud zone where the vps is
        """
        return pulumi.get(self, "zone")


class AwaitableGetVpsResult(GetVpsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpsResult(
            vps_urn=self.vps_urn,
            cluster=self.cluster,
            datacenter=self.datacenter,
            displayname=self.displayname,
            id=self.id,
            ips=self.ips,
            keymap=self.keymap,
            memory=self.memory,
            model=self.model,
            name=self.name,
            netbootmode=self.netbootmode,
            offertype=self.offertype,
            service_name=self.service_name,
            slamonitoring=self.slamonitoring,
            state=self.state,
            type=self.type,
            vcore=self.vcore,
            zone=self.zone)


def get_vps(service_name: Optional[builtins.str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpsResult:
    """
    Use this data source to retrieve information about a vps associated with your OVHcloud Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    server = ovh.Vps.get_vps(service_name="XXXXXX")
    ```


    :param builtins.str service_name: The service_name of your dedicated server.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Vps/getVps:getVps', __args__, opts=opts, typ=GetVpsResult).value

    return AwaitableGetVpsResult(
        vps_urn=pulumi.get(__ret__, 'vps_urn'),
        cluster=pulumi.get(__ret__, 'cluster'),
        datacenter=pulumi.get(__ret__, 'datacenter'),
        displayname=pulumi.get(__ret__, 'displayname'),
        id=pulumi.get(__ret__, 'id'),
        ips=pulumi.get(__ret__, 'ips'),
        keymap=pulumi.get(__ret__, 'keymap'),
        memory=pulumi.get(__ret__, 'memory'),
        model=pulumi.get(__ret__, 'model'),
        name=pulumi.get(__ret__, 'name'),
        netbootmode=pulumi.get(__ret__, 'netbootmode'),
        offertype=pulumi.get(__ret__, 'offertype'),
        service_name=pulumi.get(__ret__, 'service_name'),
        slamonitoring=pulumi.get(__ret__, 'slamonitoring'),
        state=pulumi.get(__ret__, 'state'),
        type=pulumi.get(__ret__, 'type'),
        vcore=pulumi.get(__ret__, 'vcore'),
        zone=pulumi.get(__ret__, 'zone'))
def get_vps_output(service_name: Optional[pulumi.Input[builtins.str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVpsResult]:
    """
    Use this data source to retrieve information about a vps associated with your OVHcloud Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    server = ovh.Vps.get_vps(service_name="XXXXXX")
    ```


    :param builtins.str service_name: The service_name of your dedicated server.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ovh:Vps/getVps:getVps', __args__, opts=opts, typ=GetVpsResult)
    return __ret__.apply(lambda __response__: GetVpsResult(
        vps_urn=pulumi.get(__response__, 'vps_urn'),
        cluster=pulumi.get(__response__, 'cluster'),
        datacenter=pulumi.get(__response__, 'datacenter'),
        displayname=pulumi.get(__response__, 'displayname'),
        id=pulumi.get(__response__, 'id'),
        ips=pulumi.get(__response__, 'ips'),
        keymap=pulumi.get(__response__, 'keymap'),
        memory=pulumi.get(__response__, 'memory'),
        model=pulumi.get(__response__, 'model'),
        name=pulumi.get(__response__, 'name'),
        netbootmode=pulumi.get(__response__, 'netbootmode'),
        offertype=pulumi.get(__response__, 'offertype'),
        service_name=pulumi.get(__response__, 'service_name'),
        slamonitoring=pulumi.get(__response__, 'slamonitoring'),
        state=pulumi.get(__response__, 'state'),
        type=pulumi.get(__response__, 'type'),
        vcore=pulumi.get(__response__, 'vcore'),
        zone=pulumi.get(__response__, 'zone')))
