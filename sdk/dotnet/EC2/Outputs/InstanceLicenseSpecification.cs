// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-licensespecification.html
    /// </summary>
    [OutputType]
    public sealed class InstanceLicenseSpecification
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-licensespecification.html#cfn-ec2-instance-licensespecification-licenseconfigurationarn
        /// </summary>
        public readonly string LicenseConfigurationArn;

        [OutputConstructor]
        private InstanceLicenseSpecification(string licenseConfigurationArn)
        {
            LicenseConfigurationArn = licenseConfigurationArn;
        }
    }
}