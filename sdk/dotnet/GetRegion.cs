// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    public static class GetRegion
    {
        public static Task<GetRegionResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegionResult>("aws-native:index:getRegion", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetRegionResult
    {
        public readonly string Region;

        [OutputConstructor]
        private GetRegionResult(string region)
        {
            Region = region;
        }
    }
}