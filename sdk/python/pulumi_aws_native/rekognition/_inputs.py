# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'CollectionTagArgs',
    'StreamProcessorBoundingBoxArgs',
    'StreamProcessorConnectedHomeSettingsArgs',
    'StreamProcessorDataSharingPreferenceArgs',
    'StreamProcessorFaceSearchSettingsArgs',
    'StreamProcessorKinesisDataStreamArgs',
    'StreamProcessorKinesisVideoStreamArgs',
    'StreamProcessorNotificationChannelArgs',
    'StreamProcessorPointArgs',
    'StreamProcessorS3DestinationArgs',
    'StreamProcessorTagArgs',
]

@pulumi.input_type
class CollectionTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class StreamProcessorBoundingBoxArgs:
    def __init__(__self__, *,
                 height: pulumi.Input[float],
                 left: pulumi.Input[float],
                 top: pulumi.Input[float],
                 width: pulumi.Input[float]):
        """
        A bounding box denoting a region of interest in the frame to be analyzed.
        """
        pulumi.set(__self__, "height", height)
        pulumi.set(__self__, "left", left)
        pulumi.set(__self__, "top", top)
        pulumi.set(__self__, "width", width)

    @property
    @pulumi.getter
    def height(self) -> pulumi.Input[float]:
        return pulumi.get(self, "height")

    @height.setter
    def height(self, value: pulumi.Input[float]):
        pulumi.set(self, "height", value)

    @property
    @pulumi.getter
    def left(self) -> pulumi.Input[float]:
        return pulumi.get(self, "left")

    @left.setter
    def left(self, value: pulumi.Input[float]):
        pulumi.set(self, "left", value)

    @property
    @pulumi.getter
    def top(self) -> pulumi.Input[float]:
        return pulumi.get(self, "top")

    @top.setter
    def top(self, value: pulumi.Input[float]):
        pulumi.set(self, "top", value)

    @property
    @pulumi.getter
    def width(self) -> pulumi.Input[float]:
        return pulumi.get(self, "width")

    @width.setter
    def width(self, value: pulumi.Input[float]):
        pulumi.set(self, "width", value)


@pulumi.input_type
class StreamProcessorConnectedHomeSettingsArgs:
    def __init__(__self__, *,
                 labels: pulumi.Input[Sequence[pulumi.Input[str]]],
                 min_confidence: Optional[pulumi.Input[float]] = None):
        """
        Connected home settings to use on a streaming video. Note that either ConnectedHomeSettings or FaceSearchSettings should be set. Not both
        :param pulumi.Input[float] min_confidence: Minimum object class match confidence score that must be met to return a result for a recognized object.
        """
        pulumi.set(__self__, "labels", labels)
        if min_confidence is not None:
            pulumi.set(__self__, "min_confidence", min_confidence)

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="minConfidence")
    def min_confidence(self) -> Optional[pulumi.Input[float]]:
        """
        Minimum object class match confidence score that must be met to return a result for a recognized object.
        """
        return pulumi.get(self, "min_confidence")

    @min_confidence.setter
    def min_confidence(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "min_confidence", value)


@pulumi.input_type
class StreamProcessorDataSharingPreferenceArgs:
    def __init__(__self__, *,
                 opt_in: pulumi.Input[bool]):
        """
        Indicates whether Rekognition is allowed to store the video stream data for model-training.
        :param pulumi.Input[bool] opt_in: Flag to enable data-sharing
        """
        pulumi.set(__self__, "opt_in", opt_in)

    @property
    @pulumi.getter(name="optIn")
    def opt_in(self) -> pulumi.Input[bool]:
        """
        Flag to enable data-sharing
        """
        return pulumi.get(self, "opt_in")

    @opt_in.setter
    def opt_in(self, value: pulumi.Input[bool]):
        pulumi.set(self, "opt_in", value)


@pulumi.input_type
class StreamProcessorFaceSearchSettingsArgs:
    def __init__(__self__, *,
                 collection_id: pulumi.Input[str],
                 face_match_threshold: Optional[pulumi.Input[float]] = None):
        """
        Face search settings to use on a streaming video. Note that either FaceSearchSettings or ConnectedHomeSettings should be set. Not both
        :param pulumi.Input[str] collection_id: The ID of a collection that contains faces that you want to search for.
        :param pulumi.Input[float] face_match_threshold: Minimum face match confidence score percentage that must be met to return a result for a recognized face. The default is 80. 0 is the lowest confidence. 100 is the highest confidence. Values between 0 and 100 are accepted.
        """
        pulumi.set(__self__, "collection_id", collection_id)
        if face_match_threshold is not None:
            pulumi.set(__self__, "face_match_threshold", face_match_threshold)

    @property
    @pulumi.getter(name="collectionId")
    def collection_id(self) -> pulumi.Input[str]:
        """
        The ID of a collection that contains faces that you want to search for.
        """
        return pulumi.get(self, "collection_id")

    @collection_id.setter
    def collection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "collection_id", value)

    @property
    @pulumi.getter(name="faceMatchThreshold")
    def face_match_threshold(self) -> Optional[pulumi.Input[float]]:
        """
        Minimum face match confidence score percentage that must be met to return a result for a recognized face. The default is 80. 0 is the lowest confidence. 100 is the highest confidence. Values between 0 and 100 are accepted.
        """
        return pulumi.get(self, "face_match_threshold")

    @face_match_threshold.setter
    def face_match_threshold(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "face_match_threshold", value)


@pulumi.input_type
class StreamProcessorKinesisDataStreamArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str]):
        """
        The Amazon Kinesis Data Stream stream to which the Amazon Rekognition stream processor streams the analysis results, as part of face search feature.
        :param pulumi.Input[str] arn: ARN of the Kinesis Data Stream stream.
        """
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        """
        ARN of the Kinesis Data Stream stream.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "arn", value)


@pulumi.input_type
class StreamProcessorKinesisVideoStreamArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str]):
        """
        The Kinesis Video Stream that streams the source video.
        :param pulumi.Input[str] arn: ARN of the Kinesis Video Stream that streams the source video.
        """
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        """
        ARN of the Kinesis Video Stream that streams the source video.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "arn", value)


@pulumi.input_type
class StreamProcessorNotificationChannelArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str]):
        """
        The ARN of the SNS notification channel where events of interests are published, as part of connected home feature.
        :param pulumi.Input[str] arn: ARN of the SNS topic.
        """
        pulumi.set(__self__, "arn", arn)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        """
        ARN of the SNS topic.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "arn", value)


@pulumi.input_type
class StreamProcessorPointArgs:
    def __init__(__self__, *,
                 x: pulumi.Input[float],
                 y: pulumi.Input[float]):
        """
        An (X, Y) cartesian coordinate denoting a point on the frame
        :param pulumi.Input[float] x: The X coordinate of the point.
        :param pulumi.Input[float] y: The Y coordinate of the point.
        """
        pulumi.set(__self__, "x", x)
        pulumi.set(__self__, "y", y)

    @property
    @pulumi.getter
    def x(self) -> pulumi.Input[float]:
        """
        The X coordinate of the point.
        """
        return pulumi.get(self, "x")

    @x.setter
    def x(self, value: pulumi.Input[float]):
        pulumi.set(self, "x", value)

    @property
    @pulumi.getter
    def y(self) -> pulumi.Input[float]:
        """
        The Y coordinate of the point.
        """
        return pulumi.get(self, "y")

    @y.setter
    def y(self, value: pulumi.Input[float]):
        pulumi.set(self, "y", value)


@pulumi.input_type
class StreamProcessorS3DestinationArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 object_key_prefix: Optional[pulumi.Input[str]] = None):
        """
        The S3 location in customer's account where inference output & artifacts are stored, as part of connected home feature.
        :param pulumi.Input[str] bucket_name: Name of the S3 bucket.
        :param pulumi.Input[str] object_key_prefix: The object key prefix path where the results will be stored. Default is no prefix path
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        if object_key_prefix is not None:
            pulumi.set(__self__, "object_key_prefix", object_key_prefix)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        """
        Name of the S3 bucket.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="objectKeyPrefix")
    def object_key_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        The object key prefix path where the results will be stored. Default is no prefix path
        """
        return pulumi.get(self, "object_key_prefix")

    @object_key_prefix.setter
    def object_key_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_key_prefix", value)


@pulumi.input_type
class StreamProcessorTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        :param pulumi.Input[str] value: The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

