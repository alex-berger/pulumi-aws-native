// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Inputs
{

    public sealed class DeploymentConfigTimeBasedLinearArgs : Pulumi.ResourceArgs
    {
        [Input("linearInterval", required: true)]
        public Input<int> LinearInterval { get; set; } = null!;

        [Input("linearPercentage", required: true)]
        public Input<int> LinearPercentage { get; set; } = null!;

        public DeploymentConfigTimeBasedLinearArgs()
        {
        }
    }
}
