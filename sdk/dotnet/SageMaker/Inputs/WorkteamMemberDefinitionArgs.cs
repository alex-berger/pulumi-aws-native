// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class WorkteamMemberDefinitionArgs : Pulumi.ResourceArgs
    {
        [Input("cognitoMemberDefinition", required: true)]
        public Input<Inputs.WorkteamCognitoMemberDefinitionArgs> CognitoMemberDefinition { get; set; } = null!;

        public WorkteamMemberDefinitionArgs()
        {
        }
    }
}
