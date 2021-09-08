// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RoboMaker.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html
    /// </summary>
    public sealed class SimulationApplicationSourceConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html#cfn-robomaker-simulationapplication-sourceconfig-architecture
        /// </summary>
        [Input("architecture", required: true)]
        public Input<string> Architecture { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html#cfn-robomaker-simulationapplication-sourceconfig-s3bucket
        /// </summary>
        [Input("s3Bucket", required: true)]
        public Input<string> S3Bucket { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-robomaker-simulationapplication-sourceconfig.html#cfn-robomaker-simulationapplication-sourceconfig-s3key
        /// </summary>
        [Input("s3Key", required: true)]
        public Input<string> S3Key { get; set; } = null!;

        public SimulationApplicationSourceConfigArgs()
        {
        }
    }
}