// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Outputs
{

    [OutputType]
    public sealed class ServerIdentityProviderDetails
    {
        public readonly string? DirectoryId;
        public readonly string? Function;
        public readonly string? InvocationRole;
        public readonly string? Url;

        [OutputConstructor]
        private ServerIdentityProviderDetails(
            string? directoryId,

            string? function,

            string? invocationRole,

            string? url)
        {
            DirectoryId = directoryId;
            Function = function;
            InvocationRole = invocationRole;
            Url = url;
        }
    }
}
