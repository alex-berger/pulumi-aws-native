// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticBeanstalk.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.html
    /// </summary>
    [OutputType]
    public sealed class ApplicationApplicationResourceLifecycleConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.html#cfn-elasticbeanstalk-application-applicationresourcelifecycleconfig-servicerole
        /// </summary>
        public readonly string? ServiceRole;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.html#cfn-elasticbeanstalk-application-applicationresourcelifecycleconfig-versionlifecycleconfig
        /// </summary>
        public readonly Outputs.ApplicationApplicationVersionLifecycleConfig? VersionLifecycleConfig;

        [OutputConstructor]
        private ApplicationApplicationResourceLifecycleConfig(
            string? serviceRole,

            Outputs.ApplicationApplicationVersionLifecycleConfig? versionLifecycleConfig)
        {
            ServiceRole = serviceRole;
            VersionLifecycleConfig = versionLifecycleConfig;
        }
    }
}
