// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    [OutputType]
    public sealed class DetectorModelAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-cleartimer
        /// </summary>
        public readonly Outputs.DetectorModelClearTimer? ClearTimer;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-dynamodb
        /// </summary>
        public readonly Outputs.DetectorModelDynamoDB? DynamoDB;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-dynamodbv2
        /// </summary>
        public readonly Outputs.DetectorModelDynamoDBv2? DynamoDBv2;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-firehose
        /// </summary>
        public readonly Outputs.DetectorModelFirehose? Firehose;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-iotevents
        /// </summary>
        public readonly Outputs.DetectorModelIotEvents? IotEvents;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-iotsitewise
        /// </summary>
        public readonly Outputs.DetectorModelIotSiteWise? IotSiteWise;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-iottopicpublish
        /// </summary>
        public readonly Outputs.DetectorModelIotTopicPublish? IotTopicPublish;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-lambda
        /// </summary>
        public readonly Outputs.DetectorModelLambda? Lambda;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-resettimer
        /// </summary>
        public readonly Outputs.DetectorModelResetTimer? ResetTimer;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-settimer
        /// </summary>
        public readonly Outputs.DetectorModelSetTimer? SetTimer;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-setvariable
        /// </summary>
        public readonly Outputs.DetectorModelSetVariable? SetVariable;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-sns
        /// </summary>
        public readonly Outputs.DetectorModelSns? Sns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-action.html#cfn-iotevents-detectormodel-action-sqs
        /// </summary>
        public readonly Outputs.DetectorModelSqs? Sqs;

        [OutputConstructor]
        private DetectorModelAction(
            Outputs.DetectorModelClearTimer? clearTimer,

            Outputs.DetectorModelDynamoDB? dynamoDB,

            Outputs.DetectorModelDynamoDBv2? dynamoDBv2,

            Outputs.DetectorModelFirehose? firehose,

            Outputs.DetectorModelIotEvents? iotEvents,

            Outputs.DetectorModelIotSiteWise? iotSiteWise,

            Outputs.DetectorModelIotTopicPublish? iotTopicPublish,

            Outputs.DetectorModelLambda? lambda,

            Outputs.DetectorModelResetTimer? resetTimer,

            Outputs.DetectorModelSetTimer? setTimer,

            Outputs.DetectorModelSetVariable? setVariable,

            Outputs.DetectorModelSns? sns,

            Outputs.DetectorModelSqs? sqs)
        {
            ClearTimer = clearTimer;
            DynamoDB = dynamoDB;
            DynamoDBv2 = dynamoDBv2;
            Firehose = firehose;
            IotEvents = iotEvents;
            IotSiteWise = iotSiteWise;
            IotTopicPublish = iotTopicPublish;
            Lambda = lambda;
            ResetTimer = resetTimer;
            SetTimer = setTimer;
            SetVariable = setVariable;
            Sns = sns;
            Sqs = sqs;
        }
    }
}
