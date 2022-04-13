// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticBeanstalk.Inputs
{

    public sealed class ApplicationResourceLifecycleConfigArgs : Pulumi.ResourceArgs
    {
        [Input("serviceRole")]
        public Input<string>? ServiceRole { get; set; }

        [Input("versionLifecycleConfig")]
        public Input<Inputs.ApplicationVersionLifecycleConfigArgs>? VersionLifecycleConfig { get; set; }

        public ApplicationResourceLifecycleConfigArgs()
        {
        }
    }
}
