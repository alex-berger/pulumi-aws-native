// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// A structure that contains timestamp information. For more information, see [TimeInNanos](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_TimeInNanos.html) in the *AWS IoT SiteWise API Reference*.
    /// </summary>
    [OutputType]
    public sealed class DetectorModelAssetPropertyTimestamp
    {
        /// <summary>
        /// The timestamp, in seconds, in the Unix epoch format. The valid range is between `1-31556889864403199`. You can also specify an expression.
        /// </summary>
        public readonly string? OffsetInNanos;
        /// <summary>
        /// The nanosecond offset converted from `timeInSeconds`. The valid range is between `0-999999999`. You can also specify an expression.
        /// </summary>
        public readonly string TimeInSeconds;

        [OutputConstructor]
        private DetectorModelAssetPropertyTimestamp(
            string? offsetInNanos,

            string timeInSeconds)
        {
            OffsetInNanos = offsetInNanos;
            TimeInSeconds = timeInSeconds;
        }
    }
}
