// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDB
{
    /// <summary>
    /// Resource Type definition for AWS::DynamoDB::Table
    /// </summary>
    [Obsolete(@"Table is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:dynamodb:Table")]
    public partial class Table : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("attributeDefinitions")]
        public Output<ImmutableArray<Outputs.TableAttributeDefinition>> AttributeDefinitions { get; private set; } = null!;

        [Output("billingMode")]
        public Output<string?> BillingMode { get; private set; } = null!;

        [Output("contributorInsightsSpecification")]
        public Output<Outputs.TableContributorInsightsSpecification?> ContributorInsightsSpecification { get; private set; } = null!;

        [Output("globalSecondaryIndexes")]
        public Output<ImmutableArray<Outputs.TableGlobalSecondaryIndex>> GlobalSecondaryIndexes { get; private set; } = null!;

        [Output("keySchema")]
        public Output<ImmutableArray<Outputs.TableKeySchema>> KeySchema { get; private set; } = null!;

        [Output("kinesisStreamSpecification")]
        public Output<Outputs.TableKinesisStreamSpecification?> KinesisStreamSpecification { get; private set; } = null!;

        [Output("localSecondaryIndexes")]
        public Output<ImmutableArray<Outputs.TableLocalSecondaryIndex>> LocalSecondaryIndexes { get; private set; } = null!;

        [Output("pointInTimeRecoverySpecification")]
        public Output<Outputs.TablePointInTimeRecoverySpecification?> PointInTimeRecoverySpecification { get; private set; } = null!;

        [Output("provisionedThroughput")]
        public Output<Outputs.TableProvisionedThroughput?> ProvisionedThroughput { get; private set; } = null!;

        [Output("sSESpecification")]
        public Output<Outputs.TableSSESpecification?> SSESpecification { get; private set; } = null!;

        [Output("streamArn")]
        public Output<string> StreamArn { get; private set; } = null!;

        [Output("streamSpecification")]
        public Output<Outputs.TableStreamSpecification?> StreamSpecification { get; private set; } = null!;

        [Output("tableClass")]
        public Output<string?> TableClass { get; private set; } = null!;

        [Output("tableName")]
        public Output<string?> TableName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.TableTag>> Tags { get; private set; } = null!;

        [Output("timeToLiveSpecification")]
        public Output<Outputs.TableTimeToLiveSpecification?> TimeToLiveSpecification { get; private set; } = null!;


        /// <summary>
        /// Create a Table resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Table(string name, TableArgs args, CustomResourceOptions? options = null)
            : base("aws-native:dynamodb:Table", name, args ?? new TableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Table(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:dynamodb:Table", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Table resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Table Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Table(name, id, options);
        }
    }

    public sealed class TableArgs : Pulumi.ResourceArgs
    {
        [Input("attributeDefinitions")]
        private InputList<Inputs.TableAttributeDefinitionArgs>? _attributeDefinitions;
        public InputList<Inputs.TableAttributeDefinitionArgs> AttributeDefinitions
        {
            get => _attributeDefinitions ?? (_attributeDefinitions = new InputList<Inputs.TableAttributeDefinitionArgs>());
            set => _attributeDefinitions = value;
        }

        [Input("billingMode")]
        public Input<string>? BillingMode { get; set; }

        [Input("contributorInsightsSpecification")]
        public Input<Inputs.TableContributorInsightsSpecificationArgs>? ContributorInsightsSpecification { get; set; }

        [Input("globalSecondaryIndexes")]
        private InputList<Inputs.TableGlobalSecondaryIndexArgs>? _globalSecondaryIndexes;
        public InputList<Inputs.TableGlobalSecondaryIndexArgs> GlobalSecondaryIndexes
        {
            get => _globalSecondaryIndexes ?? (_globalSecondaryIndexes = new InputList<Inputs.TableGlobalSecondaryIndexArgs>());
            set => _globalSecondaryIndexes = value;
        }

        [Input("keySchema", required: true)]
        private InputList<Inputs.TableKeySchemaArgs>? _keySchema;
        public InputList<Inputs.TableKeySchemaArgs> KeySchema
        {
            get => _keySchema ?? (_keySchema = new InputList<Inputs.TableKeySchemaArgs>());
            set => _keySchema = value;
        }

        [Input("kinesisStreamSpecification")]
        public Input<Inputs.TableKinesisStreamSpecificationArgs>? KinesisStreamSpecification { get; set; }

        [Input("localSecondaryIndexes")]
        private InputList<Inputs.TableLocalSecondaryIndexArgs>? _localSecondaryIndexes;
        public InputList<Inputs.TableLocalSecondaryIndexArgs> LocalSecondaryIndexes
        {
            get => _localSecondaryIndexes ?? (_localSecondaryIndexes = new InputList<Inputs.TableLocalSecondaryIndexArgs>());
            set => _localSecondaryIndexes = value;
        }

        [Input("pointInTimeRecoverySpecification")]
        public Input<Inputs.TablePointInTimeRecoverySpecificationArgs>? PointInTimeRecoverySpecification { get; set; }

        [Input("provisionedThroughput")]
        public Input<Inputs.TableProvisionedThroughputArgs>? ProvisionedThroughput { get; set; }

        [Input("sSESpecification")]
        public Input<Inputs.TableSSESpecificationArgs>? SSESpecification { get; set; }

        [Input("streamSpecification")]
        public Input<Inputs.TableStreamSpecificationArgs>? StreamSpecification { get; set; }

        [Input("tableClass")]
        public Input<string>? TableClass { get; set; }

        [Input("tableName")]
        public Input<string>? TableName { get; set; }

        [Input("tags")]
        private InputList<Inputs.TableTagArgs>? _tags;
        public InputList<Inputs.TableTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TableTagArgs>());
            set => _tags = value;
        }

        [Input("timeToLiveSpecification")]
        public Input<Inputs.TableTimeToLiveSpecificationArgs>? TimeToLiveSpecification { get; set; }

        public TableArgs()
        {
        }
    }
}