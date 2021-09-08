// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterserverappsettings.html
    /// </summary>
    [OutputType]
    public sealed class UserProfileJupyterServerAppSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-jupyterserverappsettings.html#cfn-sagemaker-userprofile-jupyterserverappsettings-defaultresourcespec
        /// </summary>
        public readonly Outputs.UserProfileResourceSpec? DefaultResourceSpec;

        [OutputConstructor]
        private UserProfileJupyterServerAppSettings(Outputs.UserProfileResourceSpec? defaultResourceSpec)
        {
            DefaultResourceSpec = defaultResourceSpec;
        }
    }
}