// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IAM.Outputs
{

    [OutputType]
    public sealed class UserPolicy
    {
        public readonly object PolicyDocument;
        public readonly string PolicyName;

        [OutputConstructor]
        private UserPolicy(
            object policyDocument,

            string policyName)
        {
            PolicyDocument = policyDocument;
            PolicyName = policyName;
        }
    }
}
