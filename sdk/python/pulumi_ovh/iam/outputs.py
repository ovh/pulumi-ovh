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
    'GetReferenceActionsActionResult',
]

@pulumi.output_type
class GetReferenceActionsActionResult(dict):
    def __init__(__self__, *,
                 action: str,
                 categories: Sequence[str],
                 description: str,
                 resource_type: str):
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "categories", categories)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def categories(self) -> Sequence[str]:
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")


