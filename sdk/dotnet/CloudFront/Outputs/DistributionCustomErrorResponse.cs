// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    [OutputType]
    public sealed class DistributionCustomErrorResponse
    {
        public readonly double? ErrorCachingMinTTL;
        public readonly int ErrorCode;
        public readonly int? ResponseCode;
        public readonly string? ResponsePagePath;

        [OutputConstructor]
        private DistributionCustomErrorResponse(
            double? errorCachingMinTTL,

            int errorCode,

            int? responseCode,

            string? responsePagePath)
        {
            ErrorCachingMinTTL = errorCachingMinTTL;
            ErrorCode = errorCode;
            ResponseCode = responseCode;
            ResponsePagePath = responsePagePath;
        }
    }
}
