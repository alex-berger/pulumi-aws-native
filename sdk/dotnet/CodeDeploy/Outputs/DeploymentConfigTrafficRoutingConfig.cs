// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Outputs
{

    [OutputType]
    public sealed class DeploymentConfigTrafficRoutingConfig
    {
        public readonly Outputs.DeploymentConfigTimeBasedCanary? TimeBasedCanary;
        public readonly Outputs.DeploymentConfigTimeBasedLinear? TimeBasedLinear;
        public readonly string Type;

        [OutputConstructor]
        private DeploymentConfigTrafficRoutingConfig(
            Outputs.DeploymentConfigTimeBasedCanary? timeBasedCanary,

            Outputs.DeploymentConfigTimeBasedLinear? timeBasedLinear,

            string type)
        {
            TimeBasedCanary = timeBasedCanary;
            TimeBasedLinear = timeBasedLinear;
            Type = type;
        }
    }
}
