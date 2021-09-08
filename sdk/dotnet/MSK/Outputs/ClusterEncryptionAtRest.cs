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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionatrest.html
    /// </summary>
    [OutputType]
    public sealed class ClusterEncryptionAtRest
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptionatrest.html#cfn-msk-cluster-encryptionatrest-datavolumekmskeyid
        /// </summary>
        public readonly string DataVolumeKMSKeyId;

        [OutputConstructor]
        private ClusterEncryptionAtRest(string dataVolumeKMSKeyId)
        {
            DataVolumeKMSKeyId = dataVolumeKMSKeyId;
        }
    }
}