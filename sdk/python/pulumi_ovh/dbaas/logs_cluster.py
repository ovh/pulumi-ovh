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

__all__ = ['LogsClusterArgs', 'LogsCluster']

@pulumi.input_type
class LogsClusterArgs:
    def __init__(__self__, *,
                 service_name: pulumi.Input[str],
                 archive_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 direct_input_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 query_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a LogsCluster resource.
        :param pulumi.Input[str] service_name: The service name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] archive_allowed_networks: List of IP blocks
        :param pulumi.Input[str] cluster_id: Cluster ID. If not provided, the default cluster_id is used
        :param pulumi.Input[Sequence[pulumi.Input[str]]] direct_input_allowed_networks: List of IP blocks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_allowed_networks: List of IP blocks
        """
        pulumi.set(__self__, "service_name", service_name)
        if archive_allowed_networks is not None:
            pulumi.set(__self__, "archive_allowed_networks", archive_allowed_networks)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if direct_input_allowed_networks is not None:
            pulumi.set(__self__, "direct_input_allowed_networks", direct_input_allowed_networks)
        if query_allowed_networks is not None:
            pulumi.set(__self__, "query_allowed_networks", query_allowed_networks)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="archiveAllowedNetworks")
    def archive_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "archive_allowed_networks")

    @archive_allowed_networks.setter
    def archive_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "archive_allowed_networks", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID. If not provided, the default cluster_id is used
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="directInputAllowedNetworks")
    def direct_input_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "direct_input_allowed_networks")

    @direct_input_allowed_networks.setter
    def direct_input_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "direct_input_allowed_networks", value)

    @property
    @pulumi.getter(name="queryAllowedNetworks")
    def query_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "query_allowed_networks")

    @query_allowed_networks.setter
    def query_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "query_allowed_networks", value)


@pulumi.input_type
class _LogsClusterState:
    def __init__(__self__, *,
                 archive_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 cluster_type: Optional[pulumi.Input[str]] = None,
                 dedicated_input_pem: Optional[pulumi.Input[str]] = None,
                 direct_input_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 direct_input_pem: Optional[pulumi.Input[str]] = None,
                 hostname: Optional[pulumi.Input[str]] = None,
                 initial_archive_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 initial_direct_input_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 initial_query_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_default: Optional[pulumi.Input[bool]] = None,
                 is_unlocked: Optional[pulumi.Input[bool]] = None,
                 query_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogsCluster resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] archive_allowed_networks: List of IP blocks
        :param pulumi.Input[str] cluster_id: Cluster ID. If not provided, the default cluster_id is used
        :param pulumi.Input[str] cluster_type: type of cluster (DEDICATED, PRO or TRIAL)
        :param pulumi.Input[str] dedicated_input_pem: PEM for dedicated inputs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] direct_input_allowed_networks: List of IP blocks
        :param pulumi.Input[str] direct_input_pem: PEM for direct inputs
        :param pulumi.Input[str] hostname: cluster hostname hosting tenant
        :param pulumi.Input[Sequence[pulumi.Input[str]]] initial_archive_allowed_networks: Initial allowed networks for ARCHIVE flow type
        :param pulumi.Input[Sequence[pulumi.Input[str]]] initial_direct_input_allowed_networks: Initial allowed networks for DIRECT_INPUT flow type
        :param pulumi.Input[Sequence[pulumi.Input[str]]] initial_query_allowed_networks: Initial allowed networks for QUERY flow type
        :param pulumi.Input[bool] is_default: true if all content generated by given service will be placed on this cluster
        :param pulumi.Input[bool] is_unlocked: true if given service can perform advanced operations on cluster
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_allowed_networks: List of IP blocks
        :param pulumi.Input[str] region: datacenter localization
        :param pulumi.Input[str] service_name: The service name
        """
        if archive_allowed_networks is not None:
            pulumi.set(__self__, "archive_allowed_networks", archive_allowed_networks)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if cluster_type is not None:
            pulumi.set(__self__, "cluster_type", cluster_type)
        if dedicated_input_pem is not None:
            pulumi.set(__self__, "dedicated_input_pem", dedicated_input_pem)
        if direct_input_allowed_networks is not None:
            pulumi.set(__self__, "direct_input_allowed_networks", direct_input_allowed_networks)
        if direct_input_pem is not None:
            pulumi.set(__self__, "direct_input_pem", direct_input_pem)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)
        if initial_archive_allowed_networks is not None:
            pulumi.set(__self__, "initial_archive_allowed_networks", initial_archive_allowed_networks)
        if initial_direct_input_allowed_networks is not None:
            pulumi.set(__self__, "initial_direct_input_allowed_networks", initial_direct_input_allowed_networks)
        if initial_query_allowed_networks is not None:
            pulumi.set(__self__, "initial_query_allowed_networks", initial_query_allowed_networks)
        if is_default is not None:
            pulumi.set(__self__, "is_default", is_default)
        if is_unlocked is not None:
            pulumi.set(__self__, "is_unlocked", is_unlocked)
        if query_allowed_networks is not None:
            pulumi.set(__self__, "query_allowed_networks", query_allowed_networks)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="archiveAllowedNetworks")
    def archive_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "archive_allowed_networks")

    @archive_allowed_networks.setter
    def archive_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "archive_allowed_networks", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID. If not provided, the default cluster_id is used
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> Optional[pulumi.Input[str]]:
        """
        type of cluster (DEDICATED, PRO or TRIAL)
        """
        return pulumi.get(self, "cluster_type")

    @cluster_type.setter
    def cluster_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_type", value)

    @property
    @pulumi.getter(name="dedicatedInputPem")
    def dedicated_input_pem(self) -> Optional[pulumi.Input[str]]:
        """
        PEM for dedicated inputs
        """
        return pulumi.get(self, "dedicated_input_pem")

    @dedicated_input_pem.setter
    def dedicated_input_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dedicated_input_pem", value)

    @property
    @pulumi.getter(name="directInputAllowedNetworks")
    def direct_input_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "direct_input_allowed_networks")

    @direct_input_allowed_networks.setter
    def direct_input_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "direct_input_allowed_networks", value)

    @property
    @pulumi.getter(name="directInputPem")
    def direct_input_pem(self) -> Optional[pulumi.Input[str]]:
        """
        PEM for direct inputs
        """
        return pulumi.get(self, "direct_input_pem")

    @direct_input_pem.setter
    def direct_input_pem(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direct_input_pem", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        """
        cluster hostname hosting tenant
        """
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)

    @property
    @pulumi.getter(name="initialArchiveAllowedNetworks")
    def initial_archive_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Initial allowed networks for ARCHIVE flow type
        """
        return pulumi.get(self, "initial_archive_allowed_networks")

    @initial_archive_allowed_networks.setter
    def initial_archive_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "initial_archive_allowed_networks", value)

    @property
    @pulumi.getter(name="initialDirectInputAllowedNetworks")
    def initial_direct_input_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Initial allowed networks for DIRECT_INPUT flow type
        """
        return pulumi.get(self, "initial_direct_input_allowed_networks")

    @initial_direct_input_allowed_networks.setter
    def initial_direct_input_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "initial_direct_input_allowed_networks", value)

    @property
    @pulumi.getter(name="initialQueryAllowedNetworks")
    def initial_query_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Initial allowed networks for QUERY flow type
        """
        return pulumi.get(self, "initial_query_allowed_networks")

    @initial_query_allowed_networks.setter
    def initial_query_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "initial_query_allowed_networks", value)

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> Optional[pulumi.Input[bool]]:
        """
        true if all content generated by given service will be placed on this cluster
        """
        return pulumi.get(self, "is_default")

    @is_default.setter
    def is_default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_default", value)

    @property
    @pulumi.getter(name="isUnlocked")
    def is_unlocked(self) -> Optional[pulumi.Input[bool]]:
        """
        true if given service can perform advanced operations on cluster
        """
        return pulumi.get(self, "is_unlocked")

    @is_unlocked.setter
    def is_unlocked(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_unlocked", value)

    @property
    @pulumi.getter(name="queryAllowedNetworks")
    def query_allowed_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "query_allowed_networks")

    @query_allowed_networks.setter
    def query_allowed_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "query_allowed_networks", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        datacenter localization
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class LogsCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 direct_input_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 query_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        ldp = ovh.dbaas.LogsCluster("ldp",
            archive_allowed_networks=["10.0.0.0/16"],
            cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
            direct_input_allowed_networks=["10.0.0.0/16"],
            query_allowed_networks=["10.0.0.0/16"],
            service_name="ldp-xx-xxxxx")
        ```

        ## Import

        OVHcloud DBaaS Log Data Platform clusters can be imported using the `service_name` and `cluster_id` of the cluster, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Dbaas/logsCluster:LogsCluster ldp service_name/cluster_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] archive_allowed_networks: List of IP blocks
        :param pulumi.Input[str] cluster_id: Cluster ID. If not provided, the default cluster_id is used
        :param pulumi.Input[Sequence[pulumi.Input[str]]] direct_input_allowed_networks: List of IP blocks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_allowed_networks: List of IP blocks
        :param pulumi.Input[str] service_name: The service name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogsClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        ldp = ovh.dbaas.LogsCluster("ldp",
            archive_allowed_networks=["10.0.0.0/16"],
            cluster_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
            direct_input_allowed_networks=["10.0.0.0/16"],
            query_allowed_networks=["10.0.0.0/16"],
            service_name="ldp-xx-xxxxx")
        ```

        ## Import

        OVHcloud DBaaS Log Data Platform clusters can be imported using the `service_name` and `cluster_id` of the cluster, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:Dbaas/logsCluster:LogsCluster ldp service_name/cluster_id
        ```

        :param str resource_name: The name of the resource.
        :param LogsClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogsClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 direct_input_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 query_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogsClusterArgs.__new__(LogsClusterArgs)

            __props__.__dict__["archive_allowed_networks"] = archive_allowed_networks
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["direct_input_allowed_networks"] = direct_input_allowed_networks
            __props__.__dict__["query_allowed_networks"] = query_allowed_networks
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["cluster_type"] = None
            __props__.__dict__["dedicated_input_pem"] = None
            __props__.__dict__["direct_input_pem"] = None
            __props__.__dict__["hostname"] = None
            __props__.__dict__["initial_archive_allowed_networks"] = None
            __props__.__dict__["initial_direct_input_allowed_networks"] = None
            __props__.__dict__["initial_query_allowed_networks"] = None
            __props__.__dict__["is_default"] = None
            __props__.__dict__["is_unlocked"] = None
            __props__.__dict__["region"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["dedicatedInputPem", "directInputPem", "initialArchiveAllowedNetworks", "initialDirectInputAllowedNetworks", "initialQueryAllowedNetworks"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LogsCluster, __self__).__init__(
            'ovh:Dbaas/logsCluster:LogsCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            archive_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            cluster_type: Optional[pulumi.Input[str]] = None,
            dedicated_input_pem: Optional[pulumi.Input[str]] = None,
            direct_input_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            direct_input_pem: Optional[pulumi.Input[str]] = None,
            hostname: Optional[pulumi.Input[str]] = None,
            initial_archive_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            initial_direct_input_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            initial_query_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            is_default: Optional[pulumi.Input[bool]] = None,
            is_unlocked: Optional[pulumi.Input[bool]] = None,
            query_allowed_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'LogsCluster':
        """
        Get an existing LogsCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] archive_allowed_networks: List of IP blocks
        :param pulumi.Input[str] cluster_id: Cluster ID. If not provided, the default cluster_id is used
        :param pulumi.Input[str] cluster_type: type of cluster (DEDICATED, PRO or TRIAL)
        :param pulumi.Input[str] dedicated_input_pem: PEM for dedicated inputs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] direct_input_allowed_networks: List of IP blocks
        :param pulumi.Input[str] direct_input_pem: PEM for direct inputs
        :param pulumi.Input[str] hostname: cluster hostname hosting tenant
        :param pulumi.Input[Sequence[pulumi.Input[str]]] initial_archive_allowed_networks: Initial allowed networks for ARCHIVE flow type
        :param pulumi.Input[Sequence[pulumi.Input[str]]] initial_direct_input_allowed_networks: Initial allowed networks for DIRECT_INPUT flow type
        :param pulumi.Input[Sequence[pulumi.Input[str]]] initial_query_allowed_networks: Initial allowed networks for QUERY flow type
        :param pulumi.Input[bool] is_default: true if all content generated by given service will be placed on this cluster
        :param pulumi.Input[bool] is_unlocked: true if given service can perform advanced operations on cluster
        :param pulumi.Input[Sequence[pulumi.Input[str]]] query_allowed_networks: List of IP blocks
        :param pulumi.Input[str] region: datacenter localization
        :param pulumi.Input[str] service_name: The service name
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogsClusterState.__new__(_LogsClusterState)

        __props__.__dict__["archive_allowed_networks"] = archive_allowed_networks
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["cluster_type"] = cluster_type
        __props__.__dict__["dedicated_input_pem"] = dedicated_input_pem
        __props__.__dict__["direct_input_allowed_networks"] = direct_input_allowed_networks
        __props__.__dict__["direct_input_pem"] = direct_input_pem
        __props__.__dict__["hostname"] = hostname
        __props__.__dict__["initial_archive_allowed_networks"] = initial_archive_allowed_networks
        __props__.__dict__["initial_direct_input_allowed_networks"] = initial_direct_input_allowed_networks
        __props__.__dict__["initial_query_allowed_networks"] = initial_query_allowed_networks
        __props__.__dict__["is_default"] = is_default
        __props__.__dict__["is_unlocked"] = is_unlocked
        __props__.__dict__["query_allowed_networks"] = query_allowed_networks
        __props__.__dict__["region"] = region
        __props__.__dict__["service_name"] = service_name
        return LogsCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="archiveAllowedNetworks")
    def archive_allowed_networks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "archive_allowed_networks")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[Optional[str]]:
        """
        Cluster ID. If not provided, the default cluster_id is used
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="clusterType")
    def cluster_type(self) -> pulumi.Output[str]:
        """
        type of cluster (DEDICATED, PRO or TRIAL)
        """
        return pulumi.get(self, "cluster_type")

    @property
    @pulumi.getter(name="dedicatedInputPem")
    def dedicated_input_pem(self) -> pulumi.Output[str]:
        """
        PEM for dedicated inputs
        """
        return pulumi.get(self, "dedicated_input_pem")

    @property
    @pulumi.getter(name="directInputAllowedNetworks")
    def direct_input_allowed_networks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "direct_input_allowed_networks")

    @property
    @pulumi.getter(name="directInputPem")
    def direct_input_pem(self) -> pulumi.Output[str]:
        """
        PEM for direct inputs
        """
        return pulumi.get(self, "direct_input_pem")

    @property
    @pulumi.getter
    def hostname(self) -> pulumi.Output[str]:
        """
        cluster hostname hosting tenant
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="initialArchiveAllowedNetworks")
    def initial_archive_allowed_networks(self) -> pulumi.Output[Sequence[str]]:
        """
        Initial allowed networks for ARCHIVE flow type
        """
        return pulumi.get(self, "initial_archive_allowed_networks")

    @property
    @pulumi.getter(name="initialDirectInputAllowedNetworks")
    def initial_direct_input_allowed_networks(self) -> pulumi.Output[Sequence[str]]:
        """
        Initial allowed networks for DIRECT_INPUT flow type
        """
        return pulumi.get(self, "initial_direct_input_allowed_networks")

    @property
    @pulumi.getter(name="initialQueryAllowedNetworks")
    def initial_query_allowed_networks(self) -> pulumi.Output[Sequence[str]]:
        """
        Initial allowed networks for QUERY flow type
        """
        return pulumi.get(self, "initial_query_allowed_networks")

    @property
    @pulumi.getter(name="isDefault")
    def is_default(self) -> pulumi.Output[bool]:
        """
        true if all content generated by given service will be placed on this cluster
        """
        return pulumi.get(self, "is_default")

    @property
    @pulumi.getter(name="isUnlocked")
    def is_unlocked(self) -> pulumi.Output[bool]:
        """
        true if given service can perform advanced operations on cluster
        """
        return pulumi.get(self, "is_unlocked")

    @property
    @pulumi.getter(name="queryAllowedNetworks")
    def query_allowed_networks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of IP blocks
        """
        return pulumi.get(self, "query_allowed_networks")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        datacenter localization
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The service name
        """
        return pulumi.get(self, "service_name")

