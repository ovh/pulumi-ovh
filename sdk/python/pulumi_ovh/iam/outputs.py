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
    'GetReferenceActionsActionResult',
]

@pulumi.output_type
class GetReferenceActionsActionResult(dict):
    def __init__(__self__, *,
                 action: builtins.str,
                 categories: Sequence[builtins.str],
                 description: builtins.str,
                 resource_type: builtins.str):
        """
        :param builtins.str action: Name of the action
        :param Sequence[builtins.str] categories: List of the categories of the action
        :param builtins.str description: Description of the action
        :param builtins.str resource_type: Resource type the action is related to
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "categories", categories)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def action(self) -> builtins.str:
        """
        Name of the action
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def categories(self) -> Sequence[builtins.str]:
        """
        List of the categories of the action
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description of the action
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> builtins.str:
        """
        Resource type the action is related to
        """
        return pulumi.get(self, "resource_type")


