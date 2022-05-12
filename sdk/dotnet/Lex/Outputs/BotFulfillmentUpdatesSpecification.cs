// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Provides information for updating the user on the progress of fulfilling an intent.
    /// </summary>
    [OutputType]
    public sealed class BotFulfillmentUpdatesSpecification
    {
        /// <summary>
        /// Determines whether fulfillment updates are sent to the user. When this field is true, updates are sent.
        /// </summary>
        public readonly bool Active;
        public readonly Outputs.BotFulfillmentStartResponseSpecification? StartResponse;
        /// <summary>
        /// The length of time that the fulfillment Lambda function should run before it times out.
        /// </summary>
        public readonly int? TimeoutInSeconds;
        public readonly Outputs.BotFulfillmentUpdateResponseSpecification? UpdateResponse;

        [OutputConstructor]
        private BotFulfillmentUpdatesSpecification(
            bool active,

            Outputs.BotFulfillmentStartResponseSpecification? startResponse,

            int? timeoutInSeconds,

            Outputs.BotFulfillmentUpdateResponseSpecification? updateResponse)
        {
            Active = active;
            StartResponse = startResponse;
            TimeoutInSeconds = timeoutInSeconds;
            UpdateResponse = updateResponse;
        }
    }
}
