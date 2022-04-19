// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Evidently::Feature.
 */
export class Feature extends pulumi.CustomResource {
    /**
     * Get an existing Feature resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Feature {
        return new Feature(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:evidently:Feature';

    /**
     * Returns true if the given object is an instance of Feature.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Feature {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Feature.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly defaultVariation!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly entityOverrides!: pulumi.Output<outputs.evidently.FeatureEntityOverride[] | undefined>;
    public readonly evaluationStrategy!: pulumi.Output<enums.evidently.FeatureEvaluationStrategy | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.evidently.FeatureTag[] | undefined>;
    public readonly variations!: pulumi.Output<outputs.evidently.FeatureVariationObject[]>;

    /**
     * Create a Feature resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FeatureArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.variations === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variations'");
            }
            resourceInputs["defaultVariation"] = args ? args.defaultVariation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["entityOverrides"] = args ? args.entityOverrides : undefined;
            resourceInputs["evaluationStrategy"] = args ? args.evaluationStrategy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["variations"] = args ? args.variations : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["defaultVariation"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["entityOverrides"] = undefined /*out*/;
            resourceInputs["evaluationStrategy"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["variations"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Feature.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Feature resource.
 */
export interface FeatureArgs {
    defaultVariation?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    entityOverrides?: pulumi.Input<pulumi.Input<inputs.evidently.FeatureEntityOverrideArgs>[]>;
    evaluationStrategy?: pulumi.Input<enums.evidently.FeatureEvaluationStrategy>;
    name?: pulumi.Input<string>;
    project: pulumi.Input<string>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.evidently.FeatureTagArgs>[]>;
    variations: pulumi.Input<pulumi.Input<inputs.evidently.FeatureVariationObjectArgs>[]>;
}
