# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetServerSpecificationsHardwareResult',
    'AwaitableGetServerSpecificationsHardwareResult',
    'get_server_specifications_hardware',
    'get_server_specifications_hardware_output',
]

@pulumi.output_type
class GetServerSpecificationsHardwareResult:
    """
    A collection of values returned by getServerSpecificationsHardware.
    """
    def __init__(__self__, boot_mode=None, cores_per_processor=None, default_hardware_raid_size=None, default_hardware_raid_type=None, description=None, disk_groups=None, expansion_cards=None, form_factor=None, id=None, memory_size=None, motherboard=None, number_of_processors=None, processor_architecture=None, processor_name=None, service_name=None, threads_per_processor=None, usb_keys=None):
        if boot_mode and not isinstance(boot_mode, str):
            raise TypeError("Expected argument 'boot_mode' to be a str")
        pulumi.set(__self__, "boot_mode", boot_mode)
        if cores_per_processor and not isinstance(cores_per_processor, float):
            raise TypeError("Expected argument 'cores_per_processor' to be a float")
        pulumi.set(__self__, "cores_per_processor", cores_per_processor)
        if default_hardware_raid_size and not isinstance(default_hardware_raid_size, dict):
            raise TypeError("Expected argument 'default_hardware_raid_size' to be a dict")
        pulumi.set(__self__, "default_hardware_raid_size", default_hardware_raid_size)
        if default_hardware_raid_type and not isinstance(default_hardware_raid_type, str):
            raise TypeError("Expected argument 'default_hardware_raid_type' to be a str")
        pulumi.set(__self__, "default_hardware_raid_type", default_hardware_raid_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_groups and not isinstance(disk_groups, list):
            raise TypeError("Expected argument 'disk_groups' to be a list")
        pulumi.set(__self__, "disk_groups", disk_groups)
        if expansion_cards and not isinstance(expansion_cards, list):
            raise TypeError("Expected argument 'expansion_cards' to be a list")
        pulumi.set(__self__, "expansion_cards", expansion_cards)
        if form_factor and not isinstance(form_factor, str):
            raise TypeError("Expected argument 'form_factor' to be a str")
        pulumi.set(__self__, "form_factor", form_factor)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if memory_size and not isinstance(memory_size, dict):
            raise TypeError("Expected argument 'memory_size' to be a dict")
        pulumi.set(__self__, "memory_size", memory_size)
        if motherboard and not isinstance(motherboard, str):
            raise TypeError("Expected argument 'motherboard' to be a str")
        pulumi.set(__self__, "motherboard", motherboard)
        if number_of_processors and not isinstance(number_of_processors, float):
            raise TypeError("Expected argument 'number_of_processors' to be a float")
        pulumi.set(__self__, "number_of_processors", number_of_processors)
        if processor_architecture and not isinstance(processor_architecture, str):
            raise TypeError("Expected argument 'processor_architecture' to be a str")
        pulumi.set(__self__, "processor_architecture", processor_architecture)
        if processor_name and not isinstance(processor_name, str):
            raise TypeError("Expected argument 'processor_name' to be a str")
        pulumi.set(__self__, "processor_name", processor_name)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if threads_per_processor and not isinstance(threads_per_processor, float):
            raise TypeError("Expected argument 'threads_per_processor' to be a float")
        pulumi.set(__self__, "threads_per_processor", threads_per_processor)
        if usb_keys and not isinstance(usb_keys, list):
            raise TypeError("Expected argument 'usb_keys' to be a list")
        pulumi.set(__self__, "usb_keys", usb_keys)

    @property
    @pulumi.getter(name="bootMode")
    def boot_mode(self) -> str:
        """
        Server boot mode
        """
        return pulumi.get(self, "boot_mode")

    @property
    @pulumi.getter(name="coresPerProcessor")
    def cores_per_processor(self) -> float:
        """
        Number of cores per processor
        """
        return pulumi.get(self, "cores_per_processor")

    @property
    @pulumi.getter(name="defaultHardwareRaidSize")
    def default_hardware_raid_size(self) -> 'outputs.GetServerSpecificationsHardwareDefaultHardwareRaidSizeResult':
        """
        Default hardware raid size for this disk group
        """
        return pulumi.get(self, "default_hardware_raid_size")

    @property
    @pulumi.getter(name="defaultHardwareRaidType")
    def default_hardware_raid_type(self) -> str:
        """
        Default hardware raid type for this disk group
        """
        return pulumi.get(self, "default_hardware_raid_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Expansion card description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskGroups")
    def disk_groups(self) -> Sequence['outputs.GetServerSpecificationsHardwareDiskGroupResult']:
        """
        Details about the groups of disks in the server
        """
        return pulumi.get(self, "disk_groups")

    @property
    @pulumi.getter(name="expansionCards")
    def expansion_cards(self) -> Sequence['outputs.GetServerSpecificationsHardwareExpansionCardResult']:
        """
        Details about the server's expansion cards
        """
        return pulumi.get(self, "expansion_cards")

    @property
    @pulumi.getter(name="formFactor")
    def form_factor(self) -> str:
        """
        Server form factor
        """
        return pulumi.get(self, "form_factor")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="memorySize")
    def memory_size(self) -> 'outputs.GetServerSpecificationsHardwareMemorySizeResult':
        """
        RAM capacity
        """
        return pulumi.get(self, "memory_size")

    @property
    @pulumi.getter
    def motherboard(self) -> str:
        """
        Server motherboard
        """
        return pulumi.get(self, "motherboard")

    @property
    @pulumi.getter(name="numberOfProcessors")
    def number_of_processors(self) -> float:
        """
        Number of processors in this dedicated server
        """
        return pulumi.get(self, "number_of_processors")

    @property
    @pulumi.getter(name="processorArchitecture")
    def processor_architecture(self) -> str:
        """
        Processor architecture bit
        """
        return pulumi.get(self, "processor_architecture")

    @property
    @pulumi.getter(name="processorName")
    def processor_name(self) -> str:
        """
        Processor name
        """
        return pulumi.get(self, "processor_name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="threadsPerProcessor")
    def threads_per_processor(self) -> float:
        """
        Number of threads per processor
        """
        return pulumi.get(self, "threads_per_processor")

    @property
    @pulumi.getter(name="usbKeys")
    def usb_keys(self) -> Sequence['outputs.GetServerSpecificationsHardwareUsbKeyResult']:
        """
        Capacity of the USB keys installed on your server, if any
        """
        return pulumi.get(self, "usb_keys")


class AwaitableGetServerSpecificationsHardwareResult(GetServerSpecificationsHardwareResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerSpecificationsHardwareResult(
            boot_mode=self.boot_mode,
            cores_per_processor=self.cores_per_processor,
            default_hardware_raid_size=self.default_hardware_raid_size,
            default_hardware_raid_type=self.default_hardware_raid_type,
            description=self.description,
            disk_groups=self.disk_groups,
            expansion_cards=self.expansion_cards,
            form_factor=self.form_factor,
            id=self.id,
            memory_size=self.memory_size,
            motherboard=self.motherboard,
            number_of_processors=self.number_of_processors,
            processor_architecture=self.processor_architecture,
            processor_name=self.processor_name,
            service_name=self.service_name,
            threads_per_processor=self.threads_per_processor,
            usb_keys=self.usb_keys)


def get_server_specifications_hardware(service_name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerSpecificationsHardwareResult:
    """
    Use this data source to get the hardward information about a dedicated server associated with your OVHcloud Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    spec = ovh.Dedicated.get_server_specifications_hardware(service_name="myserver")
    ```


    :param str service_name: The internal name of your dedicated server.
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Dedicated/getServerSpecificationsHardware:getServerSpecificationsHardware', __args__, opts=opts, typ=GetServerSpecificationsHardwareResult).value

    return AwaitableGetServerSpecificationsHardwareResult(
        boot_mode=pulumi.get(__ret__, 'boot_mode'),
        cores_per_processor=pulumi.get(__ret__, 'cores_per_processor'),
        default_hardware_raid_size=pulumi.get(__ret__, 'default_hardware_raid_size'),
        default_hardware_raid_type=pulumi.get(__ret__, 'default_hardware_raid_type'),
        description=pulumi.get(__ret__, 'description'),
        disk_groups=pulumi.get(__ret__, 'disk_groups'),
        expansion_cards=pulumi.get(__ret__, 'expansion_cards'),
        form_factor=pulumi.get(__ret__, 'form_factor'),
        id=pulumi.get(__ret__, 'id'),
        memory_size=pulumi.get(__ret__, 'memory_size'),
        motherboard=pulumi.get(__ret__, 'motherboard'),
        number_of_processors=pulumi.get(__ret__, 'number_of_processors'),
        processor_architecture=pulumi.get(__ret__, 'processor_architecture'),
        processor_name=pulumi.get(__ret__, 'processor_name'),
        service_name=pulumi.get(__ret__, 'service_name'),
        threads_per_processor=pulumi.get(__ret__, 'threads_per_processor'),
        usb_keys=pulumi.get(__ret__, 'usb_keys'))


@_utilities.lift_output_func(get_server_specifications_hardware)
def get_server_specifications_hardware_output(service_name: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServerSpecificationsHardwareResult]:
    """
    Use this data source to get the hardward information about a dedicated server associated with your OVHcloud Account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    spec = ovh.Dedicated.get_server_specifications_hardware(service_name="myserver")
    ```


    :param str service_name: The internal name of your dedicated server.
    """
    ...
