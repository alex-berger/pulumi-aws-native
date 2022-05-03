// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative
{
    public static class GetUrlSuffix
    {
        public static Task<GetUrlSuffixResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUrlSuffixResult>("aws-native:index:getUrlSuffix", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetUrlSuffixResult
    {
        public readonly string UrlSuffix;

        [OutputConstructor]
        private GetUrlSuffixResult(string urlSuffix)
        {
            UrlSuffix = urlSuffix;
        }
    }
}
