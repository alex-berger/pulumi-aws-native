// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaStore.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html
    /// </summary>
    public sealed class ContainerCorsRuleArgs : Pulumi.ResourceArgs
    {
        [Input("allowedHeaders")]
        private InputList<string>? _allowedHeaders;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-allowedheaders
        /// </summary>
        public InputList<string> AllowedHeaders
        {
            get => _allowedHeaders ?? (_allowedHeaders = new InputList<string>());
            set => _allowedHeaders = value;
        }

        [Input("allowedMethods")]
        private InputList<string>? _allowedMethods;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-allowedmethods
        /// </summary>
        public InputList<string> AllowedMethods
        {
            get => _allowedMethods ?? (_allowedMethods = new InputList<string>());
            set => _allowedMethods = value;
        }

        [Input("allowedOrigins")]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-allowedorigins
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        [Input("exposeHeaders")]
        private InputList<string>? _exposeHeaders;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-exposeheaders
        /// </summary>
        public InputList<string> ExposeHeaders
        {
            get => _exposeHeaders ?? (_exposeHeaders = new InputList<string>());
            set => _exposeHeaders = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediastore-container-corsrule.html#cfn-mediastore-container-corsrule-maxageseconds
        /// </summary>
        [Input("maxAgeSeconds")]
        public Input<int>? MaxAgeSeconds { get; set; }

        public ContainerCorsRuleArgs()
        {
        }
    }
}
