// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class MLTransformTransformEncryptionArgs : Pulumi.ResourceArgs
    {
        [Input("mLUserDataEncryption")]
        public Input<Inputs.MLTransformMLUserDataEncryptionArgs>? MLUserDataEncryption { get; set; }

        [Input("taskRunSecurityConfigurationName")]
        public Input<string>? TaskRunSecurityConfigurationName { get; set; }

        public MLTransformTransformEncryptionArgs()
        {
        }
    }
}
