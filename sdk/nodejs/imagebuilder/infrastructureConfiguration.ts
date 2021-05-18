// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html
 */
export class InfrastructureConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing InfrastructureConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InfrastructureConfiguration {
        return new InfrastructureConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ImageBuilder:InfrastructureConfiguration';

    /**
     * Returns true if the given object is an instance of InfrastructureConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InfrastructureConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InfrastructureConfiguration.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instanceprofilename
     */
    public readonly instanceProfileName!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instancetypes
     */
    public readonly instanceTypes!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-keypair
     */
    public readonly keyPair!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-logging
     */
    public readonly logging!: pulumi.Output<any | string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-resourcetags
     */
    public readonly resourceTags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-securitygroupids
     */
    public readonly securityGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-snstopicarn
     */
    public readonly snsTopicArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-subnetid
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-terminateinstanceonfailure
     */
    public readonly terminateInstanceOnFailure!: pulumi.Output<boolean | undefined>;

    /**
     * Create a InfrastructureConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InfrastructureConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.instanceProfileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceProfileName'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["instanceProfileName"] = args ? args.instanceProfileName : undefined;
            inputs["instanceTypes"] = args ? args.instanceTypes : undefined;
            inputs["keyPair"] = args ? args.keyPair : undefined;
            inputs["logging"] = args ? args.logging : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceTags"] = args ? args.resourceTags : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["terminateInstanceOnFailure"] = args ? args.terminateInstanceOnFailure : undefined;
            inputs["arn"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["instanceProfileName"] = undefined /*out*/;
            inputs["instanceTypes"] = undefined /*out*/;
            inputs["keyPair"] = undefined /*out*/;
            inputs["logging"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["resourceTags"] = undefined /*out*/;
            inputs["securityGroupIds"] = undefined /*out*/;
            inputs["snsTopicArn"] = undefined /*out*/;
            inputs["subnetId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["terminateInstanceOnFailure"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InfrastructureConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a InfrastructureConfiguration resource.
 */
export interface InfrastructureConfigurationArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instanceprofilename
     */
    readonly instanceProfileName: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-instancetypes
     */
    readonly instanceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-keypair
     */
    readonly keyPair?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-logging
     */
    readonly logging?: pulumi.Input<any | string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-name
     */
    readonly name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-resourcetags
     */
    readonly resourceTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-securitygroupids
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-snstopicarn
     */
    readonly snsTopicArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-subnetid
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-infrastructureconfiguration.html#cfn-imagebuilder-infrastructureconfiguration-terminateinstanceonfailure
     */
    readonly terminateInstanceOnFailure?: pulumi.Input<boolean>;
}
