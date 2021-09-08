// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-autorollbackconfig.html
    /// </summary>
    public sealed class EndpointAutoRollbackConfigArgs : Pulumi.ResourceArgs
    {
        [Input("alarms", required: true)]
        private InputList<Inputs.EndpointAlarmArgs>? _alarms;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-autorollbackconfig.html#cfn-sagemaker-endpoint-autorollbackconfig-alarms
        /// </summary>
        public InputList<Inputs.EndpointAlarmArgs> Alarms
        {
            get => _alarms ?? (_alarms = new InputList<Inputs.EndpointAlarmArgs>());
            set => _alarms = value;
        }

        public EndpointAutoRollbackConfigArgs()
        {
        }
    }
}
