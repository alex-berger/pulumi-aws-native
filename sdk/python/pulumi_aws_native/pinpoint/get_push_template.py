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
    'GetPushTemplateResult',
    'AwaitableGetPushTemplateResult',
    'get_push_template',
    'get_push_template_output',
]

@pulumi.output_type
class GetPushTemplateResult:
    def __init__(__self__, a_dm=None, a_pns=None, arn=None, baidu=None, default=None, default_substitutions=None, g_cm=None, id=None, tags=None, template_description=None):
        if a_dm and not isinstance(a_dm, dict):
            raise TypeError("Expected argument 'a_dm' to be a dict")
        pulumi.set(__self__, "a_dm", a_dm)
        if a_pns and not isinstance(a_pns, dict):
            raise TypeError("Expected argument 'a_pns' to be a dict")
        pulumi.set(__self__, "a_pns", a_pns)
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if baidu and not isinstance(baidu, dict):
            raise TypeError("Expected argument 'baidu' to be a dict")
        pulumi.set(__self__, "baidu", baidu)
        if default and not isinstance(default, dict):
            raise TypeError("Expected argument 'default' to be a dict")
        pulumi.set(__self__, "default", default)
        if default_substitutions and not isinstance(default_substitutions, str):
            raise TypeError("Expected argument 'default_substitutions' to be a str")
        pulumi.set(__self__, "default_substitutions", default_substitutions)
        if g_cm and not isinstance(g_cm, dict):
            raise TypeError("Expected argument 'g_cm' to be a dict")
        pulumi.set(__self__, "g_cm", g_cm)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if template_description and not isinstance(template_description, str):
            raise TypeError("Expected argument 'template_description' to be a str")
        pulumi.set(__self__, "template_description", template_description)

    @property
    @pulumi.getter(name="aDM")
    def a_dm(self) -> Optional['outputs.PushTemplateAndroidPushNotificationTemplate']:
        return pulumi.get(self, "a_dm")

    @property
    @pulumi.getter(name="aPNS")
    def a_pns(self) -> Optional['outputs.PushTemplateAPNSPushNotificationTemplate']:
        return pulumi.get(self, "a_pns")

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def baidu(self) -> Optional['outputs.PushTemplateAndroidPushNotificationTemplate']:
        return pulumi.get(self, "baidu")

    @property
    @pulumi.getter
    def default(self) -> Optional['outputs.PushTemplateDefaultPushNotificationTemplate']:
        return pulumi.get(self, "default")

    @property
    @pulumi.getter(name="defaultSubstitutions")
    def default_substitutions(self) -> Optional[str]:
        return pulumi.get(self, "default_substitutions")

    @property
    @pulumi.getter(name="gCM")
    def g_cm(self) -> Optional['outputs.PushTemplateAndroidPushNotificationTemplate']:
        return pulumi.get(self, "g_cm")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateDescription")
    def template_description(self) -> Optional[str]:
        return pulumi.get(self, "template_description")


class AwaitableGetPushTemplateResult(GetPushTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPushTemplateResult(
            a_dm=self.a_dm,
            a_pns=self.a_pns,
            arn=self.arn,
            baidu=self.baidu,
            default=self.default,
            default_substitutions=self.default_substitutions,
            g_cm=self.g_cm,
            id=self.id,
            tags=self.tags,
            template_description=self.template_description)


def get_push_template(id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPushTemplateResult:
    """
    Resource Type definition for AWS::Pinpoint::PushTemplate
    """
    __args__ = dict()
    __args__['id'] = id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:pinpoint:getPushTemplate', __args__, opts=opts, typ=GetPushTemplateResult).value

    return AwaitableGetPushTemplateResult(
        a_dm=__ret__.a_dm,
        a_pns=__ret__.a_pns,
        arn=__ret__.arn,
        baidu=__ret__.baidu,
        default=__ret__.default,
        default_substitutions=__ret__.default_substitutions,
        g_cm=__ret__.g_cm,
        id=__ret__.id,
        tags=__ret__.tags,
        template_description=__ret__.template_description)


@_utilities.lift_output_func(get_push_template)
def get_push_template_output(id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPushTemplateResult]:
    """
    Resource Type definition for AWS::Pinpoint::PushTemplate
    """
    ...
