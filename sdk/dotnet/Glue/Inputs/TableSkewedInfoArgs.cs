// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html
    /// </summary>
    public sealed class TableSkewedInfoArgs : Pulumi.ResourceArgs
    {
        [Input("skewedColumnNames")]
        private InputList<string>? _skewedColumnNames;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html#cfn-glue-table-skewedinfo-skewedcolumnnames
        /// </summary>
        public InputList<string> SkewedColumnNames
        {
            get => _skewedColumnNames ?? (_skewedColumnNames = new InputList<string>());
            set => _skewedColumnNames = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html#cfn-glue-table-skewedinfo-skewedcolumnvaluelocationmaps
        /// </summary>
        [Input("skewedColumnValueLocationMaps")]
        public InputUnion<System.Text.Json.JsonElement, string>? SkewedColumnValueLocationMaps { get; set; }

        [Input("skewedColumnValues")]
        private InputList<string>? _skewedColumnValues;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-skewedinfo.html#cfn-glue-table-skewedinfo-skewedcolumnvalues
        /// </summary>
        public InputList<string> SkewedColumnValues
        {
            get => _skewedColumnValues ?? (_skewedColumnValues = new InputList<string>());
            set => _skewedColumnValues = value;
        }

        public TableSkewedInfoArgs()
        {
        }
    }
}
