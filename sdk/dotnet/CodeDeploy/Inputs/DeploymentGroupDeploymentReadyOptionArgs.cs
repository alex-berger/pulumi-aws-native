// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Inputs
{

    public sealed class DeploymentGroupDeploymentReadyOptionArgs : Pulumi.ResourceArgs
    {
        [Input("actionOnTimeout")]
        public Input<string>? ActionOnTimeout { get; set; }

        [Input("waitTimeInMinutes")]
        public Input<int>? WaitTimeInMinutes { get; set; }

        public DeploymentGroupDeploymentReadyOptionArgs()
        {
        }
    }
}