// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Lambda::EventSourceMapping
 */
export class EventSourceMapping extends pulumi.CustomResource {
    /**
     * Get an existing EventSourceMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventSourceMapping {
        return new EventSourceMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lambda:EventSourceMapping';

    /**
     * Returns true if the given object is an instance of EventSourceMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSourceMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSourceMapping.__pulumiType;
    }

    /**
     * The maximum number of items to retrieve in a single batch.
     */
    public readonly batchSize!: pulumi.Output<number | undefined>;
    /**
     * (Streams) If the function returns an error, split the batch in two and retry.
     */
    public readonly bisectBatchOnFunctionError!: pulumi.Output<boolean | undefined>;
    /**
     * (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
     */
    public readonly destinationConfig!: pulumi.Output<outputs.lambda.EventSourceMappingDestinationConfig | undefined>;
    /**
     * Disables the event source mapping to pause polling and invocation.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the event source.
     */
    public readonly eventSourceArn!: pulumi.Output<string | undefined>;
    /**
     * The filter criteria to control event filtering.
     */
    public readonly filterCriteria!: pulumi.Output<outputs.lambda.EventSourceMappingFilterCriteria | undefined>;
    /**
     * The name of the Lambda function.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * (Streams) A list of response types supported by the function.
     */
    public readonly functionResponseTypes!: pulumi.Output<enums.lambda.EventSourceMappingFunctionResponseTypesItem[] | undefined>;
    /**
     * (Streams) The maximum amount of time to gather records before invoking the function, in seconds.
     */
    public readonly maximumBatchingWindowInSeconds!: pulumi.Output<number | undefined>;
    /**
     * (Streams) The maximum age of a record that Lambda sends to a function for processing.
     */
    public readonly maximumRecordAgeInSeconds!: pulumi.Output<number | undefined>;
    /**
     * (Streams) The maximum number of times to retry when the function returns an error.
     */
    public readonly maximumRetryAttempts!: pulumi.Output<number | undefined>;
    /**
     * (Streams) The number of batches to process from each shard concurrently.
     */
    public readonly parallelizationFactor!: pulumi.Output<number | undefined>;
    /**
     * (ActiveMQ) A list of ActiveMQ queues.
     */
    public readonly queues!: pulumi.Output<string[] | undefined>;
    /**
     * Self-managed event source endpoints.
     */
    public readonly selfManagedEventSource!: pulumi.Output<outputs.lambda.EventSourceMappingSelfManagedEventSource | undefined>;
    /**
     * A list of SourceAccessConfiguration.
     */
    public readonly sourceAccessConfigurations!: pulumi.Output<outputs.lambda.EventSourceMappingSourceAccessConfiguration[] | undefined>;
    /**
     * The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.
     */
    public readonly startingPosition!: pulumi.Output<string | undefined>;
    /**
     * With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
     */
    public readonly startingPositionTimestamp!: pulumi.Output<number | undefined>;
    /**
     * (Kafka) A list of Kafka topics.
     */
    public readonly topics!: pulumi.Output<string[] | undefined>;
    /**
     * (Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.
     */
    public readonly tumblingWindowInSeconds!: pulumi.Output<number | undefined>;

    /**
     * Create a EventSourceMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSourceMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            resourceInputs["batchSize"] = args ? args.batchSize : undefined;
            resourceInputs["bisectBatchOnFunctionError"] = args ? args.bisectBatchOnFunctionError : undefined;
            resourceInputs["destinationConfig"] = args ? args.destinationConfig : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventSourceArn"] = args ? args.eventSourceArn : undefined;
            resourceInputs["filterCriteria"] = args ? args.filterCriteria : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["functionResponseTypes"] = args ? args.functionResponseTypes : undefined;
            resourceInputs["maximumBatchingWindowInSeconds"] = args ? args.maximumBatchingWindowInSeconds : undefined;
            resourceInputs["maximumRecordAgeInSeconds"] = args ? args.maximumRecordAgeInSeconds : undefined;
            resourceInputs["maximumRetryAttempts"] = args ? args.maximumRetryAttempts : undefined;
            resourceInputs["parallelizationFactor"] = args ? args.parallelizationFactor : undefined;
            resourceInputs["queues"] = args ? args.queues : undefined;
            resourceInputs["selfManagedEventSource"] = args ? args.selfManagedEventSource : undefined;
            resourceInputs["sourceAccessConfigurations"] = args ? args.sourceAccessConfigurations : undefined;
            resourceInputs["startingPosition"] = args ? args.startingPosition : undefined;
            resourceInputs["startingPositionTimestamp"] = args ? args.startingPositionTimestamp : undefined;
            resourceInputs["topics"] = args ? args.topics : undefined;
            resourceInputs["tumblingWindowInSeconds"] = args ? args.tumblingWindowInSeconds : undefined;
        } else {
            resourceInputs["batchSize"] = undefined /*out*/;
            resourceInputs["bisectBatchOnFunctionError"] = undefined /*out*/;
            resourceInputs["destinationConfig"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["eventSourceArn"] = undefined /*out*/;
            resourceInputs["filterCriteria"] = undefined /*out*/;
            resourceInputs["functionName"] = undefined /*out*/;
            resourceInputs["functionResponseTypes"] = undefined /*out*/;
            resourceInputs["maximumBatchingWindowInSeconds"] = undefined /*out*/;
            resourceInputs["maximumRecordAgeInSeconds"] = undefined /*out*/;
            resourceInputs["maximumRetryAttempts"] = undefined /*out*/;
            resourceInputs["parallelizationFactor"] = undefined /*out*/;
            resourceInputs["queues"] = undefined /*out*/;
            resourceInputs["selfManagedEventSource"] = undefined /*out*/;
            resourceInputs["sourceAccessConfigurations"] = undefined /*out*/;
            resourceInputs["startingPosition"] = undefined /*out*/;
            resourceInputs["startingPositionTimestamp"] = undefined /*out*/;
            resourceInputs["topics"] = undefined /*out*/;
            resourceInputs["tumblingWindowInSeconds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventSourceMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventSourceMapping resource.
 */
export interface EventSourceMappingArgs {
    /**
     * The maximum number of items to retrieve in a single batch.
     */
    batchSize?: pulumi.Input<number>;
    /**
     * (Streams) If the function returns an error, split the batch in two and retry.
     */
    bisectBatchOnFunctionError?: pulumi.Input<boolean>;
    /**
     * (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded records.
     */
    destinationConfig?: pulumi.Input<inputs.lambda.EventSourceMappingDestinationConfigArgs>;
    /**
     * Disables the event source mapping to pause polling and invocation.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the event source.
     */
    eventSourceArn?: pulumi.Input<string>;
    /**
     * The filter criteria to control event filtering.
     */
    filterCriteria?: pulumi.Input<inputs.lambda.EventSourceMappingFilterCriteriaArgs>;
    /**
     * The name of the Lambda function.
     */
    functionName: pulumi.Input<string>;
    /**
     * (Streams) A list of response types supported by the function.
     */
    functionResponseTypes?: pulumi.Input<pulumi.Input<enums.lambda.EventSourceMappingFunctionResponseTypesItem>[]>;
    /**
     * (Streams) The maximum amount of time to gather records before invoking the function, in seconds.
     */
    maximumBatchingWindowInSeconds?: pulumi.Input<number>;
    /**
     * (Streams) The maximum age of a record that Lambda sends to a function for processing.
     */
    maximumRecordAgeInSeconds?: pulumi.Input<number>;
    /**
     * (Streams) The maximum number of times to retry when the function returns an error.
     */
    maximumRetryAttempts?: pulumi.Input<number>;
    /**
     * (Streams) The number of batches to process from each shard concurrently.
     */
    parallelizationFactor?: pulumi.Input<number>;
    /**
     * (ActiveMQ) A list of ActiveMQ queues.
     */
    queues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Self-managed event source endpoints.
     */
    selfManagedEventSource?: pulumi.Input<inputs.lambda.EventSourceMappingSelfManagedEventSourceArgs>;
    /**
     * A list of SourceAccessConfiguration.
     */
    sourceAccessConfigurations?: pulumi.Input<pulumi.Input<inputs.lambda.EventSourceMappingSourceAccessConfigurationArgs>[]>;
    /**
     * The position in a stream from which to start reading. Required for Amazon Kinesis and Amazon DynamoDB Streams sources.
     */
    startingPosition?: pulumi.Input<string>;
    /**
     * With StartingPosition set to AT_TIMESTAMP, the time from which to start reading, in Unix time seconds.
     */
    startingPositionTimestamp?: pulumi.Input<number>;
    /**
     * (Kafka) A list of Kafka topics.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (Streams) Tumbling window (non-overlapping time window) duration to perform aggregations.
     */
    tumblingWindowInSeconds?: pulumi.Input<number>;
}
