// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteVirtualMFADeviceInput struct {
	_ struct{} `type:"structure"`

	// The serial number that uniquely identifies the MFA device. For virtual MFA
	// devices, the serial number is the same as the ARN.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: =,.@:/-
	//
	// SerialNumber is a required field
	SerialNumber *string `min:"9" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVirtualMFADeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVirtualMFADeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVirtualMFADeviceInput"}

	if s.SerialNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("SerialNumber"))
	}
	if s.SerialNumber != nil && len(*s.SerialNumber) < 9 {
		invalidParams.Add(aws.NewErrParamMinLen("SerialNumber", 9))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteVirtualMFADeviceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVirtualMFADeviceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteVirtualMFADevice = "DeleteVirtualMFADevice"

// DeleteVirtualMFADeviceRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes a virtual MFA device.
//
// You must deactivate a user's virtual MFA device before you can delete it.
// For information about deactivating MFA devices, see DeactivateMFADevice.
//
//    // Example sending a request using DeleteVirtualMFADeviceRequest.
//    req := client.DeleteVirtualMFADeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteVirtualMFADevice
func (c *Client) DeleteVirtualMFADeviceRequest(input *DeleteVirtualMFADeviceInput) DeleteVirtualMFADeviceRequest {
	op := &aws.Operation{
		Name:       opDeleteVirtualMFADevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVirtualMFADeviceInput{}
	}

	req := c.newRequest(op, input, &DeleteVirtualMFADeviceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVirtualMFADeviceRequest{Request: req, Input: input, Copy: c.DeleteVirtualMFADeviceRequest}
}

// DeleteVirtualMFADeviceRequest is the request type for the
// DeleteVirtualMFADevice API operation.
type DeleteVirtualMFADeviceRequest struct {
	*aws.Request
	Input *DeleteVirtualMFADeviceInput
	Copy  func(*DeleteVirtualMFADeviceInput) DeleteVirtualMFADeviceRequest
}

// Send marshals and sends the DeleteVirtualMFADevice API request.
func (r DeleteVirtualMFADeviceRequest) Send(ctx context.Context) (*DeleteVirtualMFADeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVirtualMFADeviceResponse{
		DeleteVirtualMFADeviceOutput: r.Request.Data.(*DeleteVirtualMFADeviceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVirtualMFADeviceResponse is the response type for the
// DeleteVirtualMFADevice API operation.
type DeleteVirtualMFADeviceResponse struct {
	*DeleteVirtualMFADeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVirtualMFADevice request.
func (r *DeleteVirtualMFADeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}