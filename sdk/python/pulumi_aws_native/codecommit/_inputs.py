# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'RepositoryCodeArgs',
    'RepositoryRepositoryTriggerArgs',
    'RepositoryS3Args',
]

@pulumi.input_type
class RepositoryCodeArgs:
    def __init__(__self__, *,
                 s3: pulumi.Input['RepositoryS3Args'],
                 branch_name: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html
        :param pulumi.Input['RepositoryS3Args'] s3: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-s3
        :param pulumi.Input[str] branch_name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-branchname
        """
        pulumi.set(__self__, "s3", s3)
        if branch_name is not None:
            pulumi.set(__self__, "branch_name", branch_name)

    @property
    @pulumi.getter
    def s3(self) -> pulumi.Input['RepositoryS3Args']:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-s3
        """
        return pulumi.get(self, "s3")

    @s3.setter
    def s3(self, value: pulumi.Input['RepositoryS3Args']):
        pulumi.set(self, "s3", value)

    @property
    @pulumi.getter(name="branchName")
    def branch_name(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-branchname
        """
        return pulumi.get(self, "branch_name")

    @branch_name.setter
    def branch_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "branch_name", value)


@pulumi.input_type
class RepositoryRepositoryTriggerArgs:
    def __init__(__self__, *,
                 destination_arn: pulumi.Input[str],
                 events: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: pulumi.Input[str],
                 branches: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_data: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html
        :param pulumi.Input[str] destination_arn: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-destinationarn
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-events
        :param pulumi.Input[str] name: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-name
        :param pulumi.Input[Sequence[pulumi.Input[str]]] branches: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-branches
        :param pulumi.Input[str] custom_data: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-customdata
        """
        pulumi.set(__self__, "destination_arn", destination_arn)
        pulumi.set(__self__, "events", events)
        pulumi.set(__self__, "name", name)
        if branches is not None:
            pulumi.set(__self__, "branches", branches)
        if custom_data is not None:
            pulumi.set(__self__, "custom_data", custom_data)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-destinationarn
        """
        return pulumi.get(self, "destination_arn")

    @destination_arn.setter
    def destination_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_arn", value)

    @property
    @pulumi.getter
    def events(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-events
        """
        return pulumi.get(self, "events")

    @events.setter
    def events(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "events", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def branches(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-branches
        """
        return pulumi.get(self, "branches")

    @branches.setter
    def branches(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "branches", value)

    @property
    @pulumi.getter(name="customData")
    def custom_data(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-repositorytrigger.html#cfn-codecommit-repository-repositorytrigger-customdata
        """
        return pulumi.get(self, "custom_data")

    @custom_data.setter
    def custom_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_data", value)


@pulumi.input_type
class RepositoryS3Args:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 key: pulumi.Input[str],
                 object_version: Optional[pulumi.Input[str]] = None):
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-s3.html
        :param pulumi.Input[str] bucket: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-s3.html#cfn-codecommit-repository-s3-bucket
        :param pulumi.Input[str] key: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-s3.html#cfn-codecommit-repository-s3-key
        :param pulumi.Input[str] object_version: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-s3.html#cfn-codecommit-repository-s3-objectversion
        """
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        if object_version is not None:
            pulumi.set(__self__, "object_version", object_version)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-s3.html#cfn-codecommit-repository-s3-bucket
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-s3.html#cfn-codecommit-repository-s3-key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="objectVersion")
    def object_version(self) -> Optional[pulumi.Input[str]]:
        """
        http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-s3.html#cfn-codecommit-repository-s3-objectversion
        """
        return pulumi.get(self, "object_version")

    @object_version.setter
    def object_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_version", value)

