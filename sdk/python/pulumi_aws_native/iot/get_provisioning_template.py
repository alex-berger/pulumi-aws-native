# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetProvisioningTemplateResult',
    'AwaitableGetProvisioningTemplateResult',
    'get_provisioning_template',
    'get_provisioning_template_output',
]

@pulumi.output_type
class GetProvisioningTemplateResult:
    def __init__(__self__, description=None, enabled=None, pre_provisioning_hook=None, provisioning_role_arn=None, tags=None, template_arn=None, template_body=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if pre_provisioning_hook and not isinstance(pre_provisioning_hook, dict):
            raise TypeError("Expected argument 'pre_provisioning_hook' to be a dict")
        pulumi.set(__self__, "pre_provisioning_hook", pre_provisioning_hook)
        if provisioning_role_arn and not isinstance(provisioning_role_arn, str):
            raise TypeError("Expected argument 'provisioning_role_arn' to be a str")
        pulumi.set(__self__, "provisioning_role_arn", provisioning_role_arn)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if template_arn and not isinstance(template_arn, str):
            raise TypeError("Expected argument 'template_arn' to be a str")
        pulumi.set(__self__, "template_arn", template_arn)
        if template_body and not isinstance(template_body, str):
            raise TypeError("Expected argument 'template_body' to be a str")
        pulumi.set(__self__, "template_body", template_body)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="preProvisioningHook")
    def pre_provisioning_hook(self) -> Optional['outputs.ProvisioningTemplateProvisioningHook']:
        return pulumi.get(self, "pre_provisioning_hook")

    @property
    @pulumi.getter(name="provisioningRoleArn")
    def provisioning_role_arn(self) -> Optional[str]:
        return pulumi.get(self, "provisioning_role_arn")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ProvisioningTemplateTag']]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateArn")
    def template_arn(self) -> Optional[str]:
        return pulumi.get(self, "template_arn")

    @property
    @pulumi.getter(name="templateBody")
    def template_body(self) -> Optional[str]:
        return pulumi.get(self, "template_body")


class AwaitableGetProvisioningTemplateResult(GetProvisioningTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProvisioningTemplateResult(
            description=self.description,
            enabled=self.enabled,
            pre_provisioning_hook=self.pre_provisioning_hook,
            provisioning_role_arn=self.provisioning_role_arn,
            tags=self.tags,
            template_arn=self.template_arn,
            template_body=self.template_body)


def get_provisioning_template(template_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProvisioningTemplateResult:
    """
    Creates a fleet provisioning template.
    """
    __args__ = dict()
    __args__['templateName'] = template_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:iot:getProvisioningTemplate', __args__, opts=opts, typ=GetProvisioningTemplateResult).value

    return AwaitableGetProvisioningTemplateResult(
        description=__ret__.description,
        enabled=__ret__.enabled,
        pre_provisioning_hook=__ret__.pre_provisioning_hook,
        provisioning_role_arn=__ret__.provisioning_role_arn,
        tags=__ret__.tags,
        template_arn=__ret__.template_arn,
        template_body=__ret__.template_body)


@_utilities.lift_output_func(get_provisioning_template)
def get_provisioning_template_output(template_name: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProvisioningTemplateResult]:
    """
    Creates a fleet provisioning template.
    """
    ...