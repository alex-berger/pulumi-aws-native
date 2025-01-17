// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IVS.Outputs
{

    /// <summary>
    /// Recording Thumbnail Configuration.
    /// </summary>
    [OutputType]
    public sealed class RecordingConfigurationThumbnailConfiguration
    {
        /// <summary>
        /// Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.
        /// </summary>
        public readonly Pulumi.AwsNative.IVS.RecordingConfigurationThumbnailConfigurationRecordingMode RecordingMode;
        /// <summary>
        /// Thumbnail recording Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.
        /// </summary>
        public readonly int? TargetIntervalSeconds;

        [OutputConstructor]
        private RecordingConfigurationThumbnailConfiguration(
            Pulumi.AwsNative.IVS.RecordingConfigurationThumbnailConfigurationRecordingMode recordingMode,

            int? targetIntervalSeconds)
        {
            RecordingMode = recordingMode;
            TargetIntervalSeconds = targetIntervalSeconds;
        }
    }
}
