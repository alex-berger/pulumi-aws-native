// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Transfer.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html
    /// </summary>
    public sealed class UserPosixProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html#cfn-transfer-user-posixprofile-gid
        /// </summary>
        [Input("gid", required: true)]
        public Input<double> Gid { get; set; } = null!;

        [Input("secondaryGids")]
        private InputList<double>? _secondaryGids;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html#cfn-transfer-user-posixprofile-secondarygids
        /// </summary>
        public InputList<double> SecondaryGids
        {
            get => _secondaryGids ?? (_secondaryGids = new InputList<double>());
            set => _secondaryGids = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-posixprofile.html#cfn-transfer-user-posixprofile-uid
        /// </summary>
        [Input("uid", required: true)]
        public Input<double> Uid { get; set; } = null!;

        public UserPosixProfileArgs()
        {
        }
    }
}
