// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateMonitoring
    {
        public readonly bool? Enabled;

        [OutputConstructor]
        private LaunchTemplateMonitoring(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
