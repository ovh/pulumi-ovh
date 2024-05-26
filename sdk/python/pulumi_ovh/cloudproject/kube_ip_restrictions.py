# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KubeIpRestrictionsArgs', 'KubeIpRestrictions']

@pulumi.input_type
class KubeIpRestrictionsArgs:
    def __init__(__self__, *,
                 ips: pulumi.Input[Sequence[pulumi.Input[str]]],
                 kube_id: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a KubeIpRestrictions resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ips: List of CIDR authorized to interact with the managed Kubernetes cluster.
        :param pulumi.Input[str] kube_id: The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        pulumi.set(__self__, "ips", ips)
        pulumi.set(__self__, "kube_id", kube_id)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def ips(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of CIDR authorized to interact with the managed Kubernetes cluster.
        """
        return pulumi.get(self, "ips")

    @ips.setter
    def ips(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "ips", value)

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> pulumi.Input[str]:
        """
        The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "kube_id")

    @kube_id.setter
    def kube_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _KubeIpRestrictionsState:
    def __init__(__self__, *,
                 ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 kube_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KubeIpRestrictions resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ips: List of CIDR authorized to interact with the managed Kubernetes cluster.
        :param pulumi.Input[str] kube_id: The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        if ips is not None:
            pulumi.set(__self__, "ips", ips)
        if kube_id is not None:
            pulumi.set(__self__, "kube_id", kube_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of CIDR authorized to interact with the managed Kubernetes cluster.
        """
        return pulumi.get(self, "ips")

    @ips.setter
    def ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ips", value)

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "kube_id")

    @kube_id.setter
    def kube_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kube_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class KubeIpRestrictions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 kube_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Apply IP restrictions to an OVHcloud Managed Kubernetes cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_only = ovh.cloud_project.KubeIpRestrictions("vrackOnly",
            ips=["10.42.0.0/16"],
            kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
        ```

        ## Import

        OVHcloud Managed Kubernetes Service cluster IP restrictions can be imported using the `service_name` and the `id` of the cluster, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions iprestrictions service_name/kube_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ips: List of CIDR authorized to interact with the managed Kubernetes cluster.
        :param pulumi.Input[str] kube_id: The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KubeIpRestrictionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Apply IP restrictions to an OVHcloud Managed Kubernetes cluster.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ovh as ovh

        vrack_only = ovh.cloud_project.KubeIpRestrictions("vrackOnly",
            ips=["10.42.0.0/16"],
            kube_id="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
        ```

        ## Import

        OVHcloud Managed Kubernetes Service cluster IP restrictions can be imported using the `service_name` and the `id` of the cluster, separated by "/" E.g.,

        bash

        ```sh
        $ pulumi import ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions iprestrictions service_name/kube_id
        ```

        :param str resource_name: The name of the resource.
        :param KubeIpRestrictionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KubeIpRestrictionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 kube_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KubeIpRestrictionsArgs.__new__(KubeIpRestrictionsArgs)

            if ips is None and not opts.urn:
                raise TypeError("Missing required property 'ips'")
            __props__.__dict__["ips"] = ips
            if kube_id is None and not opts.urn:
                raise TypeError("Missing required property 'kube_id'")
            __props__.__dict__["kube_id"] = kube_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(KubeIpRestrictions, __self__).__init__(
            'ovh:CloudProject/kubeIpRestrictions:KubeIpRestrictions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            kube_id: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'KubeIpRestrictions':
        """
        Get an existing KubeIpRestrictions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ips: List of CIDR authorized to interact with the managed Kubernetes cluster.
        :param pulumi.Input[str] kube_id: The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KubeIpRestrictionsState.__new__(_KubeIpRestrictionsState)

        __props__.__dict__["ips"] = ips
        __props__.__dict__["kube_id"] = kube_id
        __props__.__dict__["service_name"] = service_name
        return KubeIpRestrictions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ips(self) -> pulumi.Output[Sequence[str]]:
        """
        List of CIDR authorized to interact with the managed Kubernetes cluster.
        """
        return pulumi.get(self, "ips")

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> pulumi.Output[str]:
        """
        The id of the managed Kubernetes cluster. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "kube_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
        """
        return pulumi.get(self, "service_name")

