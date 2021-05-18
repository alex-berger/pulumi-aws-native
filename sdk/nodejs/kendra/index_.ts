// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html
 */
export class Index extends pulumi.CustomResource {
    /**
     * Get an existing Index resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Index {
        return new Index(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:Kendra:Index';

    /**
     * Returns true if the given object is an instance of Index.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Index {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Index.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-capacityunits
     */
    public readonly capacityUnits!: pulumi.Output<outputs.Kendra.IndexCapacityUnitsConfiguration | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-documentmetadataconfigurations
     */
    public readonly documentMetadataConfigurations!: pulumi.Output<outputs.Kendra.IndexDocumentMetadataConfigurationList | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-edition
     */
    public readonly edition!: pulumi.Output<string>;
    public /*out*/ readonly id!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-rolearn
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-serversideencryptionconfiguration
     */
    public readonly serverSideEncryptionConfiguration!: pulumi.Output<outputs.Kendra.IndexServerSideEncryptionConfiguration | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-tags
     */
    public readonly tags!: pulumi.Output<outputs.Kendra.IndexTagList | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-usercontextpolicy
     */
    public readonly userContextPolicy!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-usertokenconfigurations
     */
    public readonly userTokenConfigurations!: pulumi.Output<outputs.Kendra.IndexUserTokenConfigurationList | undefined>;

    /**
     * Create a Index resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IndexArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.edition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edition'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["capacityUnits"] = args ? args.capacityUnits : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["documentMetadataConfigurations"] = args ? args.documentMetadataConfigurations : undefined;
            inputs["edition"] = args ? args.edition : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["serverSideEncryptionConfiguration"] = args ? args.serverSideEncryptionConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userContextPolicy"] = args ? args.userContextPolicy : undefined;
            inputs["userTokenConfigurations"] = args ? args.userTokenConfigurations : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["capacityUnits"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["documentMetadataConfigurations"] = undefined /*out*/;
            inputs["edition"] = undefined /*out*/;
            inputs["id"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["roleArn"] = undefined /*out*/;
            inputs["serverSideEncryptionConfiguration"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["userContextPolicy"] = undefined /*out*/;
            inputs["userTokenConfigurations"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Index.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Index resource.
 */
export interface IndexArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-capacityunits
     */
    readonly capacityUnits?: pulumi.Input<inputs.Kendra.IndexCapacityUnitsConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-description
     */
    readonly description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-documentmetadataconfigurations
     */
    readonly documentMetadataConfigurations?: pulumi.Input<inputs.Kendra.IndexDocumentMetadataConfigurationListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-edition
     */
    readonly edition: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-name
     */
    readonly name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-rolearn
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-serversideencryptionconfiguration
     */
    readonly serverSideEncryptionConfiguration?: pulumi.Input<inputs.Kendra.IndexServerSideEncryptionConfigurationArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-tags
     */
    readonly tags?: pulumi.Input<inputs.Kendra.IndexTagListArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-usercontextpolicy
     */
    readonly userContextPolicy?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-index.html#cfn-kendra-index-usertokenconfigurations
     */
    readonly userTokenConfigurations?: pulumi.Input<inputs.Kendra.IndexUserTokenConfigurationListArgs>;
}
