// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
    /// </summary>
    [OutputType]
    public sealed class RolePolicy
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
        /// </summary>
        public readonly Union<System.Text.Json.JsonElement, string> PolicyDocument;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
        /// </summary>
        public readonly string PolicyName;

        [OutputConstructor]
        private RolePolicy(
            Union<System.Text.Json.JsonElement, string> policyDocument,

            string policyName)
        {
            PolicyDocument = policyDocument;
            PolicyName = policyName;
        }
    }
}
