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
    'InstallationTemplateCustomizationArgs',
]

@pulumi.input_type
class InstallationTemplateCustomizationArgs:
    def __init__(__self__, *,
                 change_log: Optional[pulumi.Input[str]] = None,
                 custom_hostname: Optional[pulumi.Input[str]] = None,
                 post_installation_script_link: Optional[pulumi.Input[str]] = None,
                 post_installation_script_return: Optional[pulumi.Input[str]] = None,
                 rating: Optional[pulumi.Input[int]] = None,
                 ssh_key_name: Optional[pulumi.Input[str]] = None,
                 use_distribution_kernel: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] change_log: Template change log details.
        :param pulumi.Input[str] custom_hostname: Set up the server using the provided hostname instead of the default hostname.
        :param pulumi.Input[str] post_installation_script_link: Indicate the URL where your postinstall customisation script is located.
        :param pulumi.Input[str] post_installation_script_return: indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is 'loh1Xee7eo OK OK OK UGh8Ang1Gu'.
        :param pulumi.Input[int] rating: Rating.
        :param pulumi.Input[str] ssh_key_name: Name of the ssh key that should be installed. Password login will be disabled.
        :param pulumi.Input[bool] use_distribution_kernel: Use the distribution's native kernel instead of the recommended OV
        """
        if change_log is not None:
            warnings.warn("""field is not used anymore""", DeprecationWarning)
            pulumi.log.warn("""change_log is deprecated: field is not used anymore""")
        if change_log is not None:
            pulumi.set(__self__, "change_log", change_log)
        if custom_hostname is not None:
            pulumi.set(__self__, "custom_hostname", custom_hostname)
        if post_installation_script_link is not None:
            pulumi.set(__self__, "post_installation_script_link", post_installation_script_link)
        if post_installation_script_return is not None:
            pulumi.set(__self__, "post_installation_script_return", post_installation_script_return)
        if rating is not None:
            warnings.warn("""field is not used anymore""", DeprecationWarning)
            pulumi.log.warn("""rating is deprecated: field is not used anymore""")
        if rating is not None:
            pulumi.set(__self__, "rating", rating)
        if ssh_key_name is not None:
            pulumi.set(__self__, "ssh_key_name", ssh_key_name)
        if use_distribution_kernel is not None:
            pulumi.set(__self__, "use_distribution_kernel", use_distribution_kernel)

    @property
    @pulumi.getter(name="changeLog")
    def change_log(self) -> Optional[pulumi.Input[str]]:
        """
        Template change log details.
        """
        warnings.warn("""field is not used anymore""", DeprecationWarning)
        pulumi.log.warn("""change_log is deprecated: field is not used anymore""")

        return pulumi.get(self, "change_log")

    @change_log.setter
    def change_log(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "change_log", value)

    @property
    @pulumi.getter(name="customHostname")
    def custom_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        Set up the server using the provided hostname instead of the default hostname.
        """
        return pulumi.get(self, "custom_hostname")

    @custom_hostname.setter
    def custom_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_hostname", value)

    @property
    @pulumi.getter(name="postInstallationScriptLink")
    def post_installation_script_link(self) -> Optional[pulumi.Input[str]]:
        """
        Indicate the URL where your postinstall customisation script is located.
        """
        return pulumi.get(self, "post_installation_script_link")

    @post_installation_script_link.setter
    def post_installation_script_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "post_installation_script_link", value)

    @property
    @pulumi.getter(name="postInstallationScriptReturn")
    def post_installation_script_return(self) -> Optional[pulumi.Input[str]]:
        """
        indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is 'loh1Xee7eo OK OK OK UGh8Ang1Gu'.
        """
        return pulumi.get(self, "post_installation_script_return")

    @post_installation_script_return.setter
    def post_installation_script_return(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "post_installation_script_return", value)

    @property
    @pulumi.getter
    def rating(self) -> Optional[pulumi.Input[int]]:
        """
        Rating.
        """
        warnings.warn("""field is not used anymore""", DeprecationWarning)
        pulumi.log.warn("""rating is deprecated: field is not used anymore""")

        return pulumi.get(self, "rating")

    @rating.setter
    def rating(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rating", value)

    @property
    @pulumi.getter(name="sshKeyName")
    def ssh_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the ssh key that should be installed. Password login will be disabled.
        """
        return pulumi.get(self, "ssh_key_name")

    @ssh_key_name.setter
    def ssh_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssh_key_name", value)

    @property
    @pulumi.getter(name="useDistributionKernel")
    def use_distribution_kernel(self) -> Optional[pulumi.Input[bool]]:
        """
        Use the distribution's native kernel instead of the recommended OV
        """
        return pulumi.get(self, "use_distribution_kernel")

    @use_distribution_kernel.setter
    def use_distribution_kernel(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_distribution_kernel", value)

