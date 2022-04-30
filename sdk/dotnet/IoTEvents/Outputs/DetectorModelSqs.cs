// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    [OutputType]
    public sealed class DetectorModelSqs
    {
        public readonly Outputs.DetectorModelPayload? Payload;
        /// <summary>
        /// The URL of the SQS queue where the data is written.
        /// </summary>
        public readonly string QueueUrl;
        /// <summary>
        /// Set this to `TRUE` if you want the data to be base-64 encoded before it is written to the queue. Otherwise, set this to `FALSE`.
        /// </summary>
        public readonly bool? UseBase64;

        [OutputConstructor]
        private DetectorModelSqs(
            Outputs.DetectorModelPayload? payload,

            string queueUrl,

            bool? useBase64)
        {
            Payload = payload;
            QueueUrl = queueUrl;
            UseBase64 = useBase64;
        }
    }
}
