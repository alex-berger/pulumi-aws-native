// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html
 */
export class FlowLog extends pulumi.CustomResource {
    /**
     * Get an existing FlowLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FlowLog {
        return new FlowLog(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:EC2:FlowLog';

    /**
     * Returns true if the given object is an instance of FlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowLog.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-deliverlogspermissionarn
     */
    public readonly deliverLogsPermissionArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly id!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestination
     */
    public readonly logDestination!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestinationtype
     */
    public readonly logDestinationType!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logformat
     */
    public readonly logFormat!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-loggroupname
     */
    public readonly logGroupName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-maxaggregationinterval
     */
    public readonly maxAggregationInterval!: pulumi.Output<number | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourceid
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype
     */
    public readonly trafficType!: pulumi.Output<string>;

    /**
     * Create a FlowLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowLogArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            if ((!args || args.trafficType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trafficType'");
            }
            inputs["deliverLogsPermissionArn"] = args ? args.deliverLogsPermissionArn : undefined;
            inputs["logDestination"] = args ? args.logDestination : undefined;
            inputs["logDestinationType"] = args ? args.logDestinationType : undefined;
            inputs["logFormat"] = args ? args.logFormat : undefined;
            inputs["logGroupName"] = args ? args.logGroupName : undefined;
            inputs["maxAggregationInterval"] = args ? args.maxAggregationInterval : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["resourceType"] = args ? args.resourceType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["trafficType"] = args ? args.trafficType : undefined;
            inputs["id"] = undefined /*out*/;
        } else {
            inputs["deliverLogsPermissionArn"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
            inputs["logDestination"] = undefined /*out*/;
            inputs["logDestinationType"] = undefined /*out*/;
            inputs["logFormat"] = undefined /*out*/;
            inputs["logGroupName"] = undefined /*out*/;
            inputs["maxAggregationInterval"] = undefined /*out*/;
            inputs["resourceId"] = undefined /*out*/;
            inputs["resourceType"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["trafficType"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FlowLog.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FlowLog resource.
 */
export interface FlowLogArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-deliverlogspermissionarn
     */
    readonly deliverLogsPermissionArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestination
     */
    readonly logDestination?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logdestinationtype
     */
    readonly logDestinationType?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-logformat
     */
    readonly logFormat?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-loggroupname
     */
    readonly logGroupName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-maxaggregationinterval
     */
    readonly maxAggregationInterval?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourceid
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype
     */
    readonly resourceType: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-tags
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype
     */
    readonly trafficType: pulumi.Input<string>;
}
