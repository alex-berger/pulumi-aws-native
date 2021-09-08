// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html
    /// </summary>
    [OutputType]
    public sealed class ClusterClientAuthentication
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-sasl
        /// </summary>
        public readonly Outputs.ClusterSasl? Sasl;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-clientauthentication.html#cfn-msk-cluster-clientauthentication-tls
        /// </summary>
        public readonly Outputs.ClusterTls? Tls;

        [OutputConstructor]
        private ClusterClientAuthentication(
            Outputs.ClusterSasl? sasl,

            Outputs.ClusterTls? tls)
        {
            Sasl = sasl;
            Tls = tls;
        }
    }
}