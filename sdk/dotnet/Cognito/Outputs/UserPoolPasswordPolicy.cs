// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Outputs
{

    [OutputType]
    public sealed class UserPoolPasswordPolicy
    {
        public readonly int? MinimumLength;
        public readonly bool? RequireLowercase;
        public readonly bool? RequireNumbers;
        public readonly bool? RequireSymbols;
        public readonly bool? RequireUppercase;
        public readonly int? TemporaryPasswordValidityDays;

        [OutputConstructor]
        private UserPoolPasswordPolicy(
            int? minimumLength,

            bool? requireLowercase,

            bool? requireNumbers,

            bool? requireSymbols,

            bool? requireUppercase,

            int? temporaryPasswordValidityDays)
        {
            MinimumLength = minimumLength;
            RequireLowercase = requireLowercase;
            RequireNumbers = requireNumbers;
            RequireSymbols = requireSymbols;
            RequireUppercase = requireUppercase;
            TemporaryPasswordValidityDays = temporaryPasswordValidityDays;
        }
    }
}
