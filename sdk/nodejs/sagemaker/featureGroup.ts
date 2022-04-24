// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::FeatureGroup
 */
export class FeatureGroup extends pulumi.CustomResource {
    /**
     * Get an existing FeatureGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FeatureGroup {
        return new FeatureGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:FeatureGroup';

    /**
     * Returns true if the given object is an instance of FeatureGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FeatureGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FeatureGroup.__pulumiType;
    }

    /**
     * Description about the FeatureGroup.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Event Time Feature Name.
     */
    public readonly eventTimeFeatureName!: pulumi.Output<string>;
    /**
     * An Array of Feature Definition
     */
    public readonly featureDefinitions!: pulumi.Output<outputs.sagemaker.FeatureGroupFeatureDefinition[]>;
    /**
     * The Name of the FeatureGroup.
     */
    public readonly featureGroupName!: pulumi.Output<string>;
    public readonly offlineStoreConfig!: pulumi.Output<outputs.sagemaker.OfflineStoreConfigProperties | undefined>;
    public readonly onlineStoreConfig!: pulumi.Output<outputs.sagemaker.OnlineStoreConfigProperties | undefined>;
    /**
     * The Record Identifier Feature Name.
     */
    public readonly recordIdentifierFeatureName!: pulumi.Output<string>;
    /**
     * Role Arn
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * An array of key-value pair to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.sagemaker.FeatureGroupTag[] | undefined>;

    /**
     * Create a FeatureGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FeatureGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.eventTimeFeatureName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventTimeFeatureName'");
            }
            if ((!args || args.featureDefinitions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'featureDefinitions'");
            }
            if ((!args || args.recordIdentifierFeatureName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'recordIdentifierFeatureName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eventTimeFeatureName"] = args ? args.eventTimeFeatureName : undefined;
            resourceInputs["featureDefinitions"] = args ? args.featureDefinitions : undefined;
            resourceInputs["featureGroupName"] = args ? args.featureGroupName : undefined;
            resourceInputs["offlineStoreConfig"] = args ? args.offlineStoreConfig : undefined;
            resourceInputs["onlineStoreConfig"] = args ? args.onlineStoreConfig : undefined;
            resourceInputs["recordIdentifierFeatureName"] = args ? args.recordIdentifierFeatureName : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["eventTimeFeatureName"] = undefined /*out*/;
            resourceInputs["featureDefinitions"] = undefined /*out*/;
            resourceInputs["featureGroupName"] = undefined /*out*/;
            resourceInputs["offlineStoreConfig"] = undefined /*out*/;
            resourceInputs["onlineStoreConfig"] = undefined /*out*/;
            resourceInputs["recordIdentifierFeatureName"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FeatureGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FeatureGroup resource.
 */
export interface FeatureGroupArgs {
    /**
     * Description about the FeatureGroup.
     */
    description?: pulumi.Input<string>;
    /**
     * The Event Time Feature Name.
     */
    eventTimeFeatureName: pulumi.Input<string>;
    /**
     * An Array of Feature Definition
     */
    featureDefinitions: pulumi.Input<pulumi.Input<inputs.sagemaker.FeatureGroupFeatureDefinitionArgs>[]>;
    /**
     * The Name of the FeatureGroup.
     */
    featureGroupName?: pulumi.Input<string>;
    offlineStoreConfig?: pulumi.Input<inputs.sagemaker.OfflineStoreConfigPropertiesArgs>;
    onlineStoreConfig?: pulumi.Input<inputs.sagemaker.OnlineStoreConfigPropertiesArgs>;
    /**
     * The Record Identifier Feature Name.
     */
    recordIdentifierFeatureName: pulumi.Input<string>;
    /**
     * Role Arn
     */
    roleArn?: pulumi.Input<string>;
    /**
     * An array of key-value pair to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.FeatureGroupTagArgs>[]>;
}
