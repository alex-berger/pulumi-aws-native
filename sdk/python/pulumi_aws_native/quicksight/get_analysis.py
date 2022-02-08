# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'GetAnalysisResult',
    'AwaitableGetAnalysisResult',
    'get_analysis',
    'get_analysis_output',
]

@pulumi.output_type
class GetAnalysisResult:
    def __init__(__self__, arn=None, created_time=None, data_set_arns=None, errors=None, name=None, permissions=None, tags=None, theme_arn=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if created_time and not isinstance(created_time, str):
            raise TypeError("Expected argument 'created_time' to be a str")
        pulumi.set(__self__, "created_time", created_time)
        if data_set_arns and not isinstance(data_set_arns, list):
            raise TypeError("Expected argument 'data_set_arns' to be a list")
        pulumi.set(__self__, "data_set_arns", data_set_arns)
        if errors and not isinstance(errors, list):
            raise TypeError("Expected argument 'errors' to be a list")
        pulumi.set(__self__, "errors", errors)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if permissions and not isinstance(permissions, list):
            raise TypeError("Expected argument 'permissions' to be a list")
        pulumi.set(__self__, "permissions", permissions)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if theme_arn and not isinstance(theme_arn, str):
            raise TypeError("Expected argument 'theme_arn' to be a str")
        pulumi.set(__self__, "theme_arn", theme_arn)

    @property
    @pulumi.getter
    def arn(self) -> Optional[str]:
        """
        <p>The Amazon Resource Name (ARN) of the analysis.</p>
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> Optional[str]:
        """
        <p>The time that the analysis was created.</p>
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="dataSetArns")
    def data_set_arns(self) -> Optional[Sequence[str]]:
        """
        <p>The ARNs of the datasets of the analysis.</p>
        """
        return pulumi.get(self, "data_set_arns")

    @property
    @pulumi.getter
    def errors(self) -> Optional[Sequence['outputs.AnalysisError']]:
        """
        <p>Errors associated with the analysis.</p>
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        <p>The descriptive name of the analysis.</p>
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> Optional[Sequence['outputs.AnalysisResourcePermission']]:
        """
        <p>A structure that describes the principals and the resource-level permissions on an
                    analysis. You can use the <code>Permissions</code> structure to grant permissions by
                    providing a list of AWS Identity and Access Management (IAM) action information for each
                    principal listed by Amazon Resource Name (ARN). </p>

                <p>To specify no permissions, omit <code>Permissions</code>.</p>
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.AnalysisTag']]:
        """
        <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
                    analysis.</p>
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="themeArn")
    def theme_arn(self) -> Optional[str]:
        """
        <p>The ARN of the theme of the analysis.</p>
        """
        return pulumi.get(self, "theme_arn")


class AwaitableGetAnalysisResult(GetAnalysisResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAnalysisResult(
            arn=self.arn,
            created_time=self.created_time,
            data_set_arns=self.data_set_arns,
            errors=self.errors,
            name=self.name,
            permissions=self.permissions,
            tags=self.tags,
            theme_arn=self.theme_arn)


def get_analysis(analysis_id: Optional[str] = None,
                 aws_account_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAnalysisResult:
    """
    Definition of the AWS::QuickSight::Analysis Resource Type.
    """
    __args__ = dict()
    __args__['analysisId'] = analysis_id
    __args__['awsAccountId'] = aws_account_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws-native:quicksight:getAnalysis', __args__, opts=opts, typ=GetAnalysisResult).value

    return AwaitableGetAnalysisResult(
        arn=__ret__.arn,
        created_time=__ret__.created_time,
        data_set_arns=__ret__.data_set_arns,
        errors=__ret__.errors,
        name=__ret__.name,
        permissions=__ret__.permissions,
        tags=__ret__.tags,
        theme_arn=__ret__.theme_arn)


@_utilities.lift_output_func(get_analysis)
def get_analysis_output(analysis_id: Optional[pulumi.Input[str]] = None,
                        aws_account_id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAnalysisResult]:
    """
    Definition of the AWS::QuickSight::Analysis Resource Type.
    """
    ...