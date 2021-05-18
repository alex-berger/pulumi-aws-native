// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html
 */
export class DataQualityJobDefinition extends pulumi.CustomResource {
    /**
     * Get an existing DataQualityJobDefinition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataQualityJobDefinition {
        return new DataQualityJobDefinition(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:SageMaker:DataQualityJobDefinition';

    /**
     * Returns true if the given object is an instance of DataQualityJobDefinition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataQualityJobDefinition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataQualityJobDefinition.__pulumiType;
    }

    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityappspecification
     */
    public readonly dataQualityAppSpecification!: pulumi.Output<outputs.SageMaker.DataQualityJobDefinitionDataQualityAppSpecification>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig
     */
    public readonly dataQualityBaselineConfig!: pulumi.Output<outputs.SageMaker.DataQualityJobDefinitionDataQualityBaselineConfig | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityjobinput
     */
    public readonly dataQualityJobInput!: pulumi.Output<outputs.SageMaker.DataQualityJobDefinitionDataQualityJobInput>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityjoboutputconfig
     */
    public readonly dataQualityJobOutputConfig!: pulumi.Output<outputs.SageMaker.DataQualityJobDefinitionMonitoringOutputConfig>;
    public /*out*/ readonly jobDefinitionArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-jobdefinitionname
     */
    public readonly jobDefinitionName!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-jobresources
     */
    public readonly jobResources!: pulumi.Output<outputs.SageMaker.DataQualityJobDefinitionMonitoringResources>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-networkconfig
     */
    public readonly networkConfig!: pulumi.Output<outputs.SageMaker.DataQualityJobDefinitionNetworkConfig | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-rolearn
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-stoppingcondition
     */
    public readonly stoppingCondition!: pulumi.Output<outputs.SageMaker.DataQualityJobDefinitionStoppingCondition | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a DataQualityJobDefinition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataQualityJobDefinitionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataQualityAppSpecification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataQualityAppSpecification'");
            }
            if ((!args || args.dataQualityJobInput === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataQualityJobInput'");
            }
            if ((!args || args.dataQualityJobOutputConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataQualityJobOutputConfig'");
            }
            if ((!args || args.jobResources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobResources'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["dataQualityAppSpecification"] = args ? args.dataQualityAppSpecification : undefined;
            inputs["dataQualityBaselineConfig"] = args ? args.dataQualityBaselineConfig : undefined;
            inputs["dataQualityJobInput"] = args ? args.dataQualityJobInput : undefined;
            inputs["dataQualityJobOutputConfig"] = args ? args.dataQualityJobOutputConfig : undefined;
            inputs["jobDefinitionName"] = args ? args.jobDefinitionName : undefined;
            inputs["jobResources"] = args ? args.jobResources : undefined;
            inputs["networkConfig"] = args ? args.networkConfig : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["stoppingCondition"] = args ? args.stoppingCondition : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["jobDefinitionArn"] = undefined /*out*/;
        } else {
            inputs["creationTime"] = undefined /*out*/;
            inputs["dataQualityAppSpecification"] = undefined /*out*/;
            inputs["dataQualityBaselineConfig"] = undefined /*out*/;
            inputs["dataQualityJobInput"] = undefined /*out*/;
            inputs["dataQualityJobOutputConfig"] = undefined /*out*/;
            inputs["jobDefinitionArn"] = undefined /*out*/;
            inputs["jobDefinitionName"] = undefined /*out*/;
            inputs["jobResources"] = undefined /*out*/;
            inputs["networkConfig"] = undefined /*out*/;
            inputs["roleArn"] = undefined /*out*/;
            inputs["stoppingCondition"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DataQualityJobDefinition.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataQualityJobDefinition resource.
 */
export interface DataQualityJobDefinitionArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityappspecification
     */
    readonly dataQualityAppSpecification: pulumi.Input<inputs.SageMaker.DataQualityJobDefinitionDataQualityAppSpecificationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualitybaselineconfig
     */
    readonly dataQualityBaselineConfig?: pulumi.Input<inputs.SageMaker.DataQualityJobDefinitionDataQualityBaselineConfigArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityjobinput
     */
    readonly dataQualityJobInput: pulumi.Input<inputs.SageMaker.DataQualityJobDefinitionDataQualityJobInputArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-dataqualityjoboutputconfig
     */
    readonly dataQualityJobOutputConfig: pulumi.Input<inputs.SageMaker.DataQualityJobDefinitionMonitoringOutputConfigArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-jobdefinitionname
     */
    readonly jobDefinitionName?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-jobresources
     */
    readonly jobResources: pulumi.Input<inputs.SageMaker.DataQualityJobDefinitionMonitoringResourcesArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-networkconfig
     */
    readonly networkConfig?: pulumi.Input<inputs.SageMaker.DataQualityJobDefinitionNetworkConfigArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-rolearn
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-stoppingcondition
     */
    readonly stoppingCondition?: pulumi.Input<inputs.SageMaker.DataQualityJobDefinitionStoppingConditionArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-dataqualityjobdefinition.html#cfn-sagemaker-dataqualityjobdefinition-tags
     */
    readonly tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}
