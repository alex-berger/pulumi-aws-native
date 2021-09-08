// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html
 */
export class CloudFormationProduct extends pulumi.CustomResource {
    /**
     * Get an existing CloudFormationProduct resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CloudFormationProduct {
        return new CloudFormationProduct(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicecatalog:CloudFormationProduct';

    /**
     * Returns true if the given object is an instance of CloudFormationProduct.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudFormationProduct {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudFormationProduct.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-acceptlanguage
     */
    public readonly acceptLanguage!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-distributor
     */
    public readonly distributor!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-owner
     */
    public readonly owner!: pulumi.Output<string>;
    public /*out*/ readonly productName!: pulumi.Output<string>;
    public /*out*/ readonly provisioningArtifactIds!: pulumi.Output<string>;
    public /*out*/ readonly provisioningArtifactNames!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-provisioningartifactparameters
     */
    public readonly provisioningArtifactParameters!: pulumi.Output<outputs.servicecatalog.CloudFormationProductProvisioningArtifactProperties[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-replaceprovisioningartifacts
     */
    public readonly replaceProvisioningArtifacts!: pulumi.Output<boolean | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportdescription
     */
    public readonly supportDescription!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportemail
     */
    public readonly supportEmail!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supporturl
     */
    public readonly supportUrl!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;

    /**
     * Create a CloudFormationProduct resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudFormationProductArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.owner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'owner'");
            }
            if ((!args || args.provisioningArtifactParameters === undefined) && !opts.urn) {
                throw new Error("Missing required property 'provisioningArtifactParameters'");
            }
            inputs["acceptLanguage"] = args ? args.acceptLanguage : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["distributor"] = args ? args.distributor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["provisioningArtifactParameters"] = args ? args.provisioningArtifactParameters : undefined;
            inputs["replaceProvisioningArtifacts"] = args ? args.replaceProvisioningArtifacts : undefined;
            inputs["supportDescription"] = args ? args.supportDescription : undefined;
            inputs["supportEmail"] = args ? args.supportEmail : undefined;
            inputs["supportUrl"] = args ? args.supportUrl : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["productName"] = undefined /*out*/;
            inputs["provisioningArtifactIds"] = undefined /*out*/;
            inputs["provisioningArtifactNames"] = undefined /*out*/;
        } else {
            inputs["acceptLanguage"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["distributor"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["owner"] = undefined /*out*/;
            inputs["productName"] = undefined /*out*/;
            inputs["provisioningArtifactIds"] = undefined /*out*/;
            inputs["provisioningArtifactNames"] = undefined /*out*/;
            inputs["provisioningArtifactParameters"] = undefined /*out*/;
            inputs["replaceProvisioningArtifacts"] = undefined /*out*/;
            inputs["supportDescription"] = undefined /*out*/;
            inputs["supportEmail"] = undefined /*out*/;
            inputs["supportUrl"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CloudFormationProduct.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CloudFormationProduct resource.
 */
export interface CloudFormationProductArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-acceptlanguage
     */
    acceptLanguage?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-description
     */
    description?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-distributor
     */
    distributor?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-name
     */
    name: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-owner
     */
    owner: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-provisioningartifactparameters
     */
    provisioningArtifactParameters: pulumi.Input<pulumi.Input<inputs.servicecatalog.CloudFormationProductProvisioningArtifactPropertiesArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-replaceprovisioningartifacts
     */
    replaceProvisioningArtifacts?: pulumi.Input<boolean>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportdescription
     */
    supportDescription?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supportemail
     */
    supportEmail?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-supporturl
     */
    supportUrl?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationproduct.html#cfn-servicecatalog-cloudformationproduct-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}