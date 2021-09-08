// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html
    /// </summary>
    [OutputType]
    public sealed class GraphQLApiOpenIDConnectConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-authttl
        /// </summary>
        public readonly double? AuthTTL;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-clientid
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-iatttl
        /// </summary>
        public readonly double? IatTTL;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-openidconnectconfig.html#cfn-appsync-graphqlapi-openidconnectconfig-issuer
        /// </summary>
        public readonly string? Issuer;

        [OutputConstructor]
        private GraphQLApiOpenIDConnectConfig(
            double? authTTL,

            string? clientId,

            double? iatTTL,

            string? issuer)
        {
            AuthTTL = authTTL;
            ClientId = clientId;
            IatTTL = iatTTL;
            Issuer = issuer;
        }
    }
}
