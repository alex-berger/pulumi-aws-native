// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// Information needed to configure the payload.
    /// 
    /// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
    /// </summary>
    [OutputType]
    public sealed class AlarmModelPayload
    {
        /// <summary>
        /// The content of the payload. You can use a string expression that includes quoted strings (`'&lt;string&gt;'`), variables (`$variable.&lt;variable-name&gt;`), input values (`$input.&lt;input-name&gt;.&lt;path-to-datum&gt;`), string concatenations, and quoted strings that contain `${}` as the content. The recommended maximum size of a content expression is 1 KB.
        /// </summary>
        public readonly string ContentExpression;
        /// <summary>
        /// The value of the payload type can be either `STRING` or `JSON`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AlarmModelPayload(
            string contentExpression,

            string type)
        {
            ContentExpression = contentExpression;
            Type = type;
        }
    }
}
