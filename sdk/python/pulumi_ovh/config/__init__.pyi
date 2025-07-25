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

accessToken: Optional[str]
"""
The OVH API Access Token
"""

apiRateLimit: Optional[int]
"""
Specify the API request rate limit, X operations by seconds (default: unlimited)
"""

applicationKey: Optional[str]
"""
The OVH API Application Key
"""

applicationSecret: Optional[str]
"""
The OVH API Application Secret
"""

clientId: Optional[str]
"""
OAuth 2.0 application's ID
"""

clientSecret: Optional[str]
"""
OAuth 2.0 application's secret
"""

consumerKey: Optional[str]
"""
The OVH API Consumer Key
"""

endpoint: Optional[str]
"""
The OVH API endpoint to target (ex: "ovh-eu")
"""

userAgentExtra: Optional[str]
"""
Extra information to append to the user-agent
"""

