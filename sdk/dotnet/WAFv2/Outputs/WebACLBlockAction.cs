// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Block traffic towards application.
    /// </summary>
    [OutputType]
    public sealed class WebACLBlockAction
    {
        public readonly Outputs.WebACLCustomResponse? CustomResponse;

        [OutputConstructor]
        private WebACLBlockAction(Outputs.WebACLCustomResponse? customResponse)
        {
            CustomResponse = customResponse;
        }
    }
}