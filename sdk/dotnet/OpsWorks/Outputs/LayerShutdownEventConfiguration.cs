// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks.Outputs
{

    [OutputType]
    public sealed class LayerShutdownEventConfiguration
    {
        public readonly bool? DelayUntilElbConnectionsDrained;
        public readonly int? ExecutionTimeout;

        [OutputConstructor]
        private LayerShutdownEventConfiguration(
            bool? delayUntilElbConnectionsDrained,

            int? executionTimeout)
        {
            DelayUntilElbConnectionsDrained = delayUntilElbConnectionsDrained;
            ExecutionTimeout = executionTimeout;
        }
    }
}
