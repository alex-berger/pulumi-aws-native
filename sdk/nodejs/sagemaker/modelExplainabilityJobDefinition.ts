// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelExplainabilityJobDefinition
 */
export class ModelExplainabilityJobDefinition extends pulumi.CustomResource {
    /**
     * Get an existing ModelExplainabilityJobDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ModelExplainabilityJobDefinition {
        return new ModelExplainabilityJobDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:ModelExplainabilityJobDefinition';

    /**
     * Returns true if the given object is an instance of ModelExplainabilityJobDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ModelExplainabilityJobDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ModelExplainabilityJobDefinition.__pulumiType;
    }

    /**
     * The time at which the job definition was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    public /*out*/ readonly jobDefinitionArn!: pulumi.Output<string>;
    public readonly jobDefinitionName!: pulumi.Output<string | undefined>;
    public readonly jobResources!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionMonitoringResources>;
    public readonly modelExplainabilityAppSpecification!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionModelExplainabilityAppSpecification>;
    public readonly modelExplainabilityBaselineConfig!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfig | undefined>;
    public readonly modelExplainabilityJobInput!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionModelExplainabilityJobInput>;
    public readonly modelExplainabilityJobOutputConfig!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionMonitoringOutputConfig>;
    public readonly networkConfig!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionNetworkConfig | undefined>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
     */
    public readonly roleArn!: pulumi.Output<string>;
    public readonly stoppingCondition!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionStoppingCondition | undefined>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.sagemaker.ModelExplainabilityJobDefinitionTag[] | undefined>;

    /**
     * Create a ModelExplainabilityJobDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ModelExplainabilityJobDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.jobResources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobResources'");
            }
            if ((!args || args.modelExplainabilityAppSpecification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelExplainabilityAppSpecification'");
            }
            if ((!args || args.modelExplainabilityJobInput === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelExplainabilityJobInput'");
            }
            if ((!args || args.modelExplainabilityJobOutputConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'modelExplainabilityJobOutputConfig'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["jobDefinitionName"] = args ? args.jobDefinitionName : undefined;
            resourceInputs["jobResources"] = args ? args.jobResources : undefined;
            resourceInputs["modelExplainabilityAppSpecification"] = args ? args.modelExplainabilityAppSpecification : undefined;
            resourceInputs["modelExplainabilityBaselineConfig"] = args ? args.modelExplainabilityBaselineConfig : undefined;
            resourceInputs["modelExplainabilityJobInput"] = args ? args.modelExplainabilityJobInput : undefined;
            resourceInputs["modelExplainabilityJobOutputConfig"] = args ? args.modelExplainabilityJobOutputConfig : undefined;
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["stoppingCondition"] = args ? args.stoppingCondition : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["jobDefinitionArn"] = undefined /*out*/;
        } else {
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["jobDefinitionArn"] = undefined /*out*/;
            resourceInputs["jobDefinitionName"] = undefined /*out*/;
            resourceInputs["jobResources"] = undefined /*out*/;
            resourceInputs["modelExplainabilityAppSpecification"] = undefined /*out*/;
            resourceInputs["modelExplainabilityBaselineConfig"] = undefined /*out*/;
            resourceInputs["modelExplainabilityJobInput"] = undefined /*out*/;
            resourceInputs["modelExplainabilityJobOutputConfig"] = undefined /*out*/;
            resourceInputs["networkConfig"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["stoppingCondition"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ModelExplainabilityJobDefinition.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ModelExplainabilityJobDefinition resource.
 */
export interface ModelExplainabilityJobDefinitionArgs {
    jobDefinitionName?: pulumi.Input<string>;
    jobResources: pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionMonitoringResourcesArgs>;
    modelExplainabilityAppSpecification: pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionModelExplainabilityAppSpecificationArgs>;
    modelExplainabilityBaselineConfig?: pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigArgs>;
    modelExplainabilityJobInput: pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionModelExplainabilityJobInputArgs>;
    modelExplainabilityJobOutputConfig: pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionMonitoringOutputConfigArgs>;
    networkConfig?: pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionNetworkConfigArgs>;
    /**
     * The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
     */
    roleArn: pulumi.Input<string>;
    stoppingCondition?: pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionStoppingConditionArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.ModelExplainabilityJobDefinitionTagArgs>[]>;
}
