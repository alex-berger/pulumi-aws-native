// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Definition of the AWS::QuickSight::Analysis Resource Type.
 */
export class Analysis extends pulumi.CustomResource {
    /**
     * Get an existing Analysis resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Analysis {
        return new Analysis(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:quicksight:Analysis';

    /**
     * Returns true if the given object is an instance of Analysis.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Analysis {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Analysis.__pulumiType;
    }

    public readonly analysisId!: pulumi.Output<string>;
    /**
     * <p>The Amazon Resource Name (ARN) of the analysis.</p>
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly awsAccountId!: pulumi.Output<string>;
    /**
     * <p>The time that the analysis was created.</p>
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * <p>The ARNs of the datasets of the analysis.</p>
     */
    public /*out*/ readonly dataSetArns!: pulumi.Output<string[]>;
    /**
     * <p>Errors associated with the analysis.</p>
     */
    public readonly errors!: pulumi.Output<outputs.quicksight.AnalysisError[] | undefined>;
    /**
     * <p>The time that the analysis was last updated.</p>
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * <p>The descriptive name of the analysis.</p>
     */
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly parameters!: pulumi.Output<outputs.quicksight.AnalysisParameters | undefined>;
    /**
     * <p>A structure that describes the principals and the resource-level permissions on an
     *             analysis. You can use the <code>Permissions</code> structure to grant permissions by
     *             providing a list of AWS Identity and Access Management (IAM) action information for each
     *             principal listed by Amazon Resource Name (ARN). </p>
     *
     *         <p>To specify no permissions, omit <code>Permissions</code>.</p>
     */
    public readonly permissions!: pulumi.Output<outputs.quicksight.AnalysisResourcePermission[] | undefined>;
    /**
     * <p>A list of the associated sheets with the unique identifier and name of each sheet.</p>
     */
    public /*out*/ readonly sheets!: pulumi.Output<outputs.quicksight.AnalysisSheet[]>;
    public readonly sourceEntity!: pulumi.Output<outputs.quicksight.AnalysisSourceEntity>;
    public /*out*/ readonly status!: pulumi.Output<enums.quicksight.AnalysisResourceStatus>;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
     *             analysis.</p>
     */
    public readonly tags!: pulumi.Output<outputs.quicksight.AnalysisTag[] | undefined>;
    /**
     * <p>The ARN of the theme of the analysis.</p>
     */
    public readonly themeArn!: pulumi.Output<string | undefined>;

    /**
     * Create a Analysis resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalysisArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.analysisId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'analysisId'");
            }
            if ((!args || args.awsAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsAccountId'");
            }
            if ((!args || args.sourceEntity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEntity'");
            }
            resourceInputs["analysisId"] = args ? args.analysisId : undefined;
            resourceInputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            resourceInputs["errors"] = args ? args.errors : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["sourceEntity"] = args ? args.sourceEntity : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["themeArn"] = args ? args.themeArn : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["dataSetArns"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["sheets"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["analysisId"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["awsAccountId"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["dataSetArns"] = undefined /*out*/;
            resourceInputs["errors"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["permissions"] = undefined /*out*/;
            resourceInputs["sheets"] = undefined /*out*/;
            resourceInputs["sourceEntity"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["themeArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Analysis.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Analysis resource.
 */
export interface AnalysisArgs {
    analysisId: pulumi.Input<string>;
    awsAccountId: pulumi.Input<string>;
    /**
     * <p>Errors associated with the analysis.</p>
     */
    errors?: pulumi.Input<pulumi.Input<inputs.quicksight.AnalysisErrorArgs>[]>;
    /**
     * <p>The descriptive name of the analysis.</p>
     */
    name?: pulumi.Input<string>;
    parameters?: pulumi.Input<inputs.quicksight.AnalysisParametersArgs>;
    /**
     * <p>A structure that describes the principals and the resource-level permissions on an
     *             analysis. You can use the <code>Permissions</code> structure to grant permissions by
     *             providing a list of AWS Identity and Access Management (IAM) action information for each
     *             principal listed by Amazon Resource Name (ARN). </p>
     *
     *         <p>To specify no permissions, omit <code>Permissions</code>.</p>
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.AnalysisResourcePermissionArgs>[]>;
    sourceEntity: pulumi.Input<inputs.quicksight.AnalysisSourceEntityArgs>;
    /**
     * <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
     *             analysis.</p>
     */
    tags?: pulumi.Input<pulumi.Input<inputs.quicksight.AnalysisTagArgs>[]>;
    /**
     * <p>The ARN of the theme of the analysis.</p>
     */
    themeArn?: pulumi.Input<string>;
}
