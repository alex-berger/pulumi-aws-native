// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Forecast.Outputs
{

    [OutputType]
    public sealed class AttributesItemProperties
    {
        /// <summary>
        /// Name of the dataset field
        /// </summary>
        public readonly string? AttributeName;
        /// <summary>
        /// Data type of the field
        /// </summary>
        public readonly Pulumi.AwsNative.Forecast.DatasetAttributesItemPropertiesAttributeType? AttributeType;

        [OutputConstructor]
        private AttributesItemProperties(
            string? attributeName,

            Pulumi.AwsNative.Forecast.DatasetAttributesItemPropertiesAttributeType? attributeType)
        {
            AttributeName = attributeName;
            AttributeType = attributeType;
        }
    }
}
