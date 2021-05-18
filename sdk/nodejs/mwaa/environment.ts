// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:MWAA:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-airflowconfigurationoptions
     */
    public readonly airflowConfigurationOptions!: pulumi.Output<outputs.MWAA.EnvironmentAirflowConfigurationOptions | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-airflowversion
     */
    public readonly airflowVersion!: pulumi.Output<string | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-dags3path
     */
    public readonly dagS3Path!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-environmentclass
     */
    public readonly environmentClass!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-executionrolearn
     */
    public readonly executionRoleArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-kmskey
     */
    public readonly kmsKey!: pulumi.Output<string | undefined>;
    public /*out*/ readonly lastUpdate!: pulumi.Output<outputs.MWAA.EnvironmentLastUpdate>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-loggingconfiguration
     */
    public readonly loggingConfiguration!: pulumi.Output<outputs.MWAA.EnvironmentLoggingConfiguration | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-maxworkers
     */
    public readonly maxWorkers!: pulumi.Output<number | undefined>;
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-networkconfiguration
     */
    public readonly networkConfiguration!: pulumi.Output<outputs.MWAA.EnvironmentNetworkConfiguration | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-pluginss3objectversion
     */
    public readonly pluginsS3ObjectVersion!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-pluginss3path
     */
    public readonly pluginsS3Path!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-requirementss3objectversion
     */
    public readonly requirementsS3ObjectVersion!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-requirementss3path
     */
    public readonly requirementsS3Path!: pulumi.Output<string | undefined>;
    public /*out*/ readonly serviceRoleArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-sourcebucketarn
     */
    public readonly sourceBucketArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-tags
     */
    public readonly tags!: pulumi.Output<outputs.MWAA.EnvironmentTagMap | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-webserveraccessmode
     */
    public readonly webserverAccessMode!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-webserverurl
     */
    public readonly webserverUrl!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-weeklymaintenancewindowstart
     */
    public readonly weeklyMaintenanceWindowStart!: pulumi.Output<string | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["airflowConfigurationOptions"] = args ? args.airflowConfigurationOptions : undefined;
            inputs["airflowVersion"] = args ? args.airflowVersion : undefined;
            inputs["dagS3Path"] = args ? args.dagS3Path : undefined;
            inputs["environmentClass"] = args ? args.environmentClass : undefined;
            inputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            inputs["kmsKey"] = args ? args.kmsKey : undefined;
            inputs["loggingConfiguration"] = args ? args.loggingConfiguration : undefined;
            inputs["maxWorkers"] = args ? args.maxWorkers : undefined;
            inputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            inputs["pluginsS3ObjectVersion"] = args ? args.pluginsS3ObjectVersion : undefined;
            inputs["pluginsS3Path"] = args ? args.pluginsS3Path : undefined;
            inputs["requirementsS3ObjectVersion"] = args ? args.requirementsS3ObjectVersion : undefined;
            inputs["requirementsS3Path"] = args ? args.requirementsS3Path : undefined;
            inputs["sourceBucketArn"] = args ? args.sourceBucketArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["webserverAccessMode"] = args ? args.webserverAccessMode : undefined;
            inputs["webserverUrl"] = args ? args.webserverUrl : undefined;
            inputs["weeklyMaintenanceWindowStart"] = args ? args.weeklyMaintenanceWindowStart : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["lastUpdate"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["serviceRoleArn"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        } else {
            inputs["airflowConfigurationOptions"] = undefined /*out*/;
            inputs["airflowVersion"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["dagS3Path"] = undefined /*out*/;
            inputs["environmentClass"] = undefined /*out*/;
            inputs["executionRoleArn"] = undefined /*out*/;
            inputs["kmsKey"] = undefined /*out*/;
            inputs["lastUpdate"] = undefined /*out*/;
            inputs["loggingConfiguration"] = undefined /*out*/;
            inputs["maxWorkers"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkConfiguration"] = undefined /*out*/;
            inputs["pluginsS3ObjectVersion"] = undefined /*out*/;
            inputs["pluginsS3Path"] = undefined /*out*/;
            inputs["requirementsS3ObjectVersion"] = undefined /*out*/;
            inputs["requirementsS3Path"] = undefined /*out*/;
            inputs["serviceRoleArn"] = undefined /*out*/;
            inputs["sourceBucketArn"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["webserverAccessMode"] = undefined /*out*/;
            inputs["webserverUrl"] = undefined /*out*/;
            inputs["weeklyMaintenanceWindowStart"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Environment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-airflowconfigurationoptions
     */
    readonly airflowConfigurationOptions?: pulumi.Input<inputs.MWAA.EnvironmentAirflowConfigurationOptionsArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-airflowversion
     */
    readonly airflowVersion?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-dags3path
     */
    readonly dagS3Path?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-environmentclass
     */
    readonly environmentClass?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-executionrolearn
     */
    readonly executionRoleArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-kmskey
     */
    readonly kmsKey?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-loggingconfiguration
     */
    readonly loggingConfiguration?: pulumi.Input<inputs.MWAA.EnvironmentLoggingConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-maxworkers
     */
    readonly maxWorkers?: pulumi.Input<number>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-networkconfiguration
     */
    readonly networkConfiguration?: pulumi.Input<inputs.MWAA.EnvironmentNetworkConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-pluginss3objectversion
     */
    readonly pluginsS3ObjectVersion?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-pluginss3path
     */
    readonly pluginsS3Path?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-requirementss3objectversion
     */
    readonly requirementsS3ObjectVersion?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-requirementss3path
     */
    readonly requirementsS3Path?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-sourcebucketarn
     */
    readonly sourceBucketArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-tags
     */
    readonly tags?: pulumi.Input<inputs.MWAA.EnvironmentTagMapArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-webserveraccessmode
     */
    readonly webserverAccessMode?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-webserverurl
     */
    readonly webserverUrl?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#cfn-mwaa-environment-weeklymaintenancewindowstart
     */
    readonly weeklyMaintenanceWindowStart?: pulumi.Input<string>;
}
