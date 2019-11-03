package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iam"
)

func (g *Getter) getIAMRoleAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.iam.GetRoleWithContext(ctx, &iam.GetRoleInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"Arn": nil,
		"RoleId": nil,
	}
	return attrs, nil
}
