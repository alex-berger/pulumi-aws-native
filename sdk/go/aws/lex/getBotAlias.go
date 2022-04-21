// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lex

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Bot Alias enables you to change the version of a bot without updating applications that use the bot
func LookupBotAlias(ctx *pulumi.Context, args *LookupBotAliasArgs, opts ...pulumi.InvokeOption) (*LookupBotAliasResult, error) {
	var rv LookupBotAliasResult
	err := ctx.Invoke("aws-native:lex:getBotAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBotAliasArgs struct {
	BotAliasId string `pulumi:"botAliasId"`
	BotId      string `pulumi:"botId"`
}

type LookupBotAliasResult struct {
	Arn                     *string                          `pulumi:"arn"`
	BotAliasId              *string                          `pulumi:"botAliasId"`
	BotAliasLocaleSettings  []BotAliasLocaleSettingsItem     `pulumi:"botAliasLocaleSettings"`
	BotAliasName            *string                          `pulumi:"botAliasName"`
	BotAliasStatus          *BotAliasStatus                  `pulumi:"botAliasStatus"`
	BotVersion              *string                          `pulumi:"botVersion"`
	ConversationLogSettings *BotAliasConversationLogSettings `pulumi:"conversationLogSettings"`
	Description             *string                          `pulumi:"description"`
	// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
	SentimentAnalysisSettings *SentimentAnalysisSettingsProperties `pulumi:"sentimentAnalysisSettings"`
}

func LookupBotAliasOutput(ctx *pulumi.Context, args LookupBotAliasOutputArgs, opts ...pulumi.InvokeOption) LookupBotAliasResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBotAliasResult, error) {
			args := v.(LookupBotAliasArgs)
			r, err := LookupBotAlias(ctx, &args, opts...)
			var s LookupBotAliasResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBotAliasResultOutput)
}

type LookupBotAliasOutputArgs struct {
	BotAliasId pulumi.StringInput `pulumi:"botAliasId"`
	BotId      pulumi.StringInput `pulumi:"botId"`
}

func (LookupBotAliasOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotAliasArgs)(nil)).Elem()
}

type LookupBotAliasResultOutput struct{ *pulumi.OutputState }

func (LookupBotAliasResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBotAliasResult)(nil)).Elem()
}

func (o LookupBotAliasResultOutput) ToLookupBotAliasResultOutput() LookupBotAliasResultOutput {
	return o
}

func (o LookupBotAliasResultOutput) ToLookupBotAliasResultOutputWithContext(ctx context.Context) LookupBotAliasResultOutput {
	return o
}

func (o LookupBotAliasResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupBotAliasResultOutput) BotAliasId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *string { return v.BotAliasId }).(pulumi.StringPtrOutput)
}

func (o LookupBotAliasResultOutput) BotAliasLocaleSettings() BotAliasLocaleSettingsItemArrayOutput {
	return o.ApplyT(func(v LookupBotAliasResult) []BotAliasLocaleSettingsItem { return v.BotAliasLocaleSettings }).(BotAliasLocaleSettingsItemArrayOutput)
}

func (o LookupBotAliasResultOutput) BotAliasName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *string { return v.BotAliasName }).(pulumi.StringPtrOutput)
}

func (o LookupBotAliasResultOutput) BotAliasStatus() BotAliasStatusPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *BotAliasStatus { return v.BotAliasStatus }).(BotAliasStatusPtrOutput)
}

func (o LookupBotAliasResultOutput) BotVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *string { return v.BotVersion }).(pulumi.StringPtrOutput)
}

func (o LookupBotAliasResultOutput) ConversationLogSettings() BotAliasConversationLogSettingsPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *BotAliasConversationLogSettings { return v.ConversationLogSettings }).(BotAliasConversationLogSettingsPtrOutput)
}

func (o LookupBotAliasResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Determines whether Amazon Lex will use Amazon Comprehend to detect the sentiment of user utterances.
func (o LookupBotAliasResultOutput) SentimentAnalysisSettings() SentimentAnalysisSettingsPropertiesPtrOutput {
	return o.ApplyT(func(v LookupBotAliasResult) *SentimentAnalysisSettingsProperties { return v.SentimentAnalysisSettings }).(SentimentAnalysisSettingsPropertiesPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBotAliasResultOutput{})
}
