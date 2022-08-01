package appmesh

import (
	"github.com/weaveworks/goformation/v4/cloudformation/types"

	"github.com/weaveworks/goformation/v4/cloudformation/policies"
)

// VirtualGateway_VirtualGatewayHttp2ConnectionPool AWS CloudFormation Resource (AWS::AppMesh::VirtualGateway.VirtualGatewayHttp2ConnectionPool)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhttp2connectionpool.html
type VirtualGateway_VirtualGatewayHttp2ConnectionPool struct {

	// MaxRequests AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewayhttp2connectionpool.html#cfn-appmesh-virtualgateway-virtualgatewayhttp2connectionpool-maxrequests
	MaxRequests *types.Value `json:"MaxRequests"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *VirtualGateway_VirtualGatewayHttp2ConnectionPool) AWSCloudFormationType() string {
	return "AWS::AppMesh::VirtualGateway.VirtualGatewayHttp2ConnectionPool"
}