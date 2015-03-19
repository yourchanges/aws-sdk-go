package sqs

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// AddPermissionRequest generates a request for the AddPermission operation.
func (c *SQS) AddPermissionRequest(input *AddPermissionInput) (req *aws.Request, output *AddPermissionOutput) {
	if opAddPermission == nil {
		opAddPermission = &aws.Operation{
			Name:       "AddPermission",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAddPermission, input, output)
	output = &AddPermissionOutput{}
	req.Data = output
	return
}

func (c *SQS) AddPermission(input *AddPermissionInput) (output *AddPermissionOutput, err error) {
	req, out := c.AddPermissionRequest(input)
	output = out
	err = req.Send()
	return
}

var opAddPermission *aws.Operation

// ChangeMessageVisibilityRequest generates a request for the ChangeMessageVisibility operation.
func (c *SQS) ChangeMessageVisibilityRequest(input *ChangeMessageVisibilityInput) (req *aws.Request, output *ChangeMessageVisibilityOutput) {
	if opChangeMessageVisibility == nil {
		opChangeMessageVisibility = &aws.Operation{
			Name:       "ChangeMessageVisibility",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opChangeMessageVisibility, input, output)
	output = &ChangeMessageVisibilityOutput{}
	req.Data = output
	return
}

func (c *SQS) ChangeMessageVisibility(input *ChangeMessageVisibilityInput) (output *ChangeMessageVisibilityOutput, err error) {
	req, out := c.ChangeMessageVisibilityRequest(input)
	output = out
	err = req.Send()
	return
}

var opChangeMessageVisibility *aws.Operation

// ChangeMessageVisibilityBatchRequest generates a request for the ChangeMessageVisibilityBatch operation.
func (c *SQS) ChangeMessageVisibilityBatchRequest(input *ChangeMessageVisibilityBatchInput) (req *aws.Request, output *ChangeMessageVisibilityBatchOutput) {
	if opChangeMessageVisibilityBatch == nil {
		opChangeMessageVisibilityBatch = &aws.Operation{
			Name:       "ChangeMessageVisibilityBatch",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opChangeMessageVisibilityBatch, input, output)
	output = &ChangeMessageVisibilityBatchOutput{}
	req.Data = output
	return
}

func (c *SQS) ChangeMessageVisibilityBatch(input *ChangeMessageVisibilityBatchInput) (output *ChangeMessageVisibilityBatchOutput, err error) {
	req, out := c.ChangeMessageVisibilityBatchRequest(input)
	output = out
	err = req.Send()
	return
}

var opChangeMessageVisibilityBatch *aws.Operation

// CreateQueueRequest generates a request for the CreateQueue operation.
func (c *SQS) CreateQueueRequest(input *CreateQueueInput) (req *aws.Request, output *CreateQueueOutput) {
	if opCreateQueue == nil {
		opCreateQueue = &aws.Operation{
			Name:       "CreateQueue",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateQueue, input, output)
	output = &CreateQueueOutput{}
	req.Data = output
	return
}

func (c *SQS) CreateQueue(input *CreateQueueInput) (output *CreateQueueOutput, err error) {
	req, out := c.CreateQueueRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateQueue *aws.Operation

// DeleteMessageRequest generates a request for the DeleteMessage operation.
func (c *SQS) DeleteMessageRequest(input *DeleteMessageInput) (req *aws.Request, output *DeleteMessageOutput) {
	if opDeleteMessage == nil {
		opDeleteMessage = &aws.Operation{
			Name:       "DeleteMessage",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteMessage, input, output)
	output = &DeleteMessageOutput{}
	req.Data = output
	return
}

func (c *SQS) DeleteMessage(input *DeleteMessageInput) (output *DeleteMessageOutput, err error) {
	req, out := c.DeleteMessageRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteMessage *aws.Operation

// DeleteMessageBatchRequest generates a request for the DeleteMessageBatch operation.
func (c *SQS) DeleteMessageBatchRequest(input *DeleteMessageBatchInput) (req *aws.Request, output *DeleteMessageBatchOutput) {
	if opDeleteMessageBatch == nil {
		opDeleteMessageBatch = &aws.Operation{
			Name:       "DeleteMessageBatch",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteMessageBatch, input, output)
	output = &DeleteMessageBatchOutput{}
	req.Data = output
	return
}

func (c *SQS) DeleteMessageBatch(input *DeleteMessageBatchInput) (output *DeleteMessageBatchOutput, err error) {
	req, out := c.DeleteMessageBatchRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteMessageBatch *aws.Operation

// DeleteQueueRequest generates a request for the DeleteQueue operation.
func (c *SQS) DeleteQueueRequest(input *DeleteQueueInput) (req *aws.Request, output *DeleteQueueOutput) {
	if opDeleteQueue == nil {
		opDeleteQueue = &aws.Operation{
			Name:       "DeleteQueue",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteQueue, input, output)
	output = &DeleteQueueOutput{}
	req.Data = output
	return
}

func (c *SQS) DeleteQueue(input *DeleteQueueInput) (output *DeleteQueueOutput, err error) {
	req, out := c.DeleteQueueRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteQueue *aws.Operation

// GetQueueAttributesRequest generates a request for the GetQueueAttributes operation.
func (c *SQS) GetQueueAttributesRequest(input *GetQueueAttributesInput) (req *aws.Request, output *GetQueueAttributesOutput) {
	if opGetQueueAttributes == nil {
		opGetQueueAttributes = &aws.Operation{
			Name:       "GetQueueAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetQueueAttributes, input, output)
	output = &GetQueueAttributesOutput{}
	req.Data = output
	return
}

func (c *SQS) GetQueueAttributes(input *GetQueueAttributesInput) (output *GetQueueAttributesOutput, err error) {
	req, out := c.GetQueueAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetQueueAttributes *aws.Operation

// GetQueueURLRequest generates a request for the GetQueueURL operation.
func (c *SQS) GetQueueURLRequest(input *GetQueueURLInput) (req *aws.Request, output *GetQueueURLOutput) {
	if opGetQueueURL == nil {
		opGetQueueURL = &aws.Operation{
			Name:       "GetQueueUrl",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetQueueURL, input, output)
	output = &GetQueueURLOutput{}
	req.Data = output
	return
}

func (c *SQS) GetQueueURL(input *GetQueueURLInput) (output *GetQueueURLOutput, err error) {
	req, out := c.GetQueueURLRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetQueueURL *aws.Operation

// ListDeadLetterSourceQueuesRequest generates a request for the ListDeadLetterSourceQueues operation.
func (c *SQS) ListDeadLetterSourceQueuesRequest(input *ListDeadLetterSourceQueuesInput) (req *aws.Request, output *ListDeadLetterSourceQueuesOutput) {
	if opListDeadLetterSourceQueues == nil {
		opListDeadLetterSourceQueues = &aws.Operation{
			Name:       "ListDeadLetterSourceQueues",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListDeadLetterSourceQueues, input, output)
	output = &ListDeadLetterSourceQueuesOutput{}
	req.Data = output
	return
}

func (c *SQS) ListDeadLetterSourceQueues(input *ListDeadLetterSourceQueuesInput) (output *ListDeadLetterSourceQueuesOutput, err error) {
	req, out := c.ListDeadLetterSourceQueuesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListDeadLetterSourceQueues *aws.Operation

// ListQueuesRequest generates a request for the ListQueues operation.
func (c *SQS) ListQueuesRequest(input *ListQueuesInput) (req *aws.Request, output *ListQueuesOutput) {
	if opListQueues == nil {
		opListQueues = &aws.Operation{
			Name:       "ListQueues",
			HTTPMethod: "POST",
			HTTPPath:   "/",
			Paginator: &aws.Paginator{
				InputToken:      "",
				OutputToken:     "",
				LimitToken:      "",
				TruncationToken: "",
			},
		}
	}

	req = aws.NewRequest(c.Service, opListQueues, input, output)
	output = &ListQueuesOutput{}
	req.Data = output
	return
}

func (c *SQS) ListQueues(input *ListQueuesInput) (output *ListQueuesOutput, err error) {
	req, out := c.ListQueuesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListQueues *aws.Operation

// PurgeQueueRequest generates a request for the PurgeQueue operation.
func (c *SQS) PurgeQueueRequest(input *PurgeQueueInput) (req *aws.Request, output *PurgeQueueOutput) {
	if opPurgeQueue == nil {
		opPurgeQueue = &aws.Operation{
			Name:       "PurgeQueue",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPurgeQueue, input, output)
	output = &PurgeQueueOutput{}
	req.Data = output
	return
}

func (c *SQS) PurgeQueue(input *PurgeQueueInput) (output *PurgeQueueOutput, err error) {
	req, out := c.PurgeQueueRequest(input)
	output = out
	err = req.Send()
	return
}

var opPurgeQueue *aws.Operation

// ReceiveMessageRequest generates a request for the ReceiveMessage operation.
func (c *SQS) ReceiveMessageRequest(input *ReceiveMessageInput) (req *aws.Request, output *ReceiveMessageOutput) {
	if opReceiveMessage == nil {
		opReceiveMessage = &aws.Operation{
			Name:       "ReceiveMessage",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opReceiveMessage, input, output)
	output = &ReceiveMessageOutput{}
	req.Data = output
	return
}

func (c *SQS) ReceiveMessage(input *ReceiveMessageInput) (output *ReceiveMessageOutput, err error) {
	req, out := c.ReceiveMessageRequest(input)
	output = out
	err = req.Send()
	return
}

var opReceiveMessage *aws.Operation

// RemovePermissionRequest generates a request for the RemovePermission operation.
func (c *SQS) RemovePermissionRequest(input *RemovePermissionInput) (req *aws.Request, output *RemovePermissionOutput) {
	if opRemovePermission == nil {
		opRemovePermission = &aws.Operation{
			Name:       "RemovePermission",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRemovePermission, input, output)
	output = &RemovePermissionOutput{}
	req.Data = output
	return
}

func (c *SQS) RemovePermission(input *RemovePermissionInput) (output *RemovePermissionOutput, err error) {
	req, out := c.RemovePermissionRequest(input)
	output = out
	err = req.Send()
	return
}

var opRemovePermission *aws.Operation

// SendMessageRequest generates a request for the SendMessage operation.
func (c *SQS) SendMessageRequest(input *SendMessageInput) (req *aws.Request, output *SendMessageOutput) {
	if opSendMessage == nil {
		opSendMessage = &aws.Operation{
			Name:       "SendMessage",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSendMessage, input, output)
	output = &SendMessageOutput{}
	req.Data = output
	return
}

func (c *SQS) SendMessage(input *SendMessageInput) (output *SendMessageOutput, err error) {
	req, out := c.SendMessageRequest(input)
	output = out
	err = req.Send()
	return
}

var opSendMessage *aws.Operation

// SendMessageBatchRequest generates a request for the SendMessageBatch operation.
func (c *SQS) SendMessageBatchRequest(input *SendMessageBatchInput) (req *aws.Request, output *SendMessageBatchOutput) {
	if opSendMessageBatch == nil {
		opSendMessageBatch = &aws.Operation{
			Name:       "SendMessageBatch",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSendMessageBatch, input, output)
	output = &SendMessageBatchOutput{}
	req.Data = output
	return
}

func (c *SQS) SendMessageBatch(input *SendMessageBatchInput) (output *SendMessageBatchOutput, err error) {
	req, out := c.SendMessageBatchRequest(input)
	output = out
	err = req.Send()
	return
}

var opSendMessageBatch *aws.Operation

// SetQueueAttributesRequest generates a request for the SetQueueAttributes operation.
func (c *SQS) SetQueueAttributesRequest(input *SetQueueAttributesInput) (req *aws.Request, output *SetQueueAttributesOutput) {
	if opSetQueueAttributes == nil {
		opSetQueueAttributes = &aws.Operation{
			Name:       "SetQueueAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetQueueAttributes, input, output)
	output = &SetQueueAttributesOutput{}
	req.Data = output
	return
}

func (c *SQS) SetQueueAttributes(input *SetQueueAttributesInput) (output *SetQueueAttributesOutput, err error) {
	req, out := c.SetQueueAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetQueueAttributes *aws.Operation

type AddPermissionInput struct {
	AWSAccountIDs []*string `locationName:"AWSAccountIds" locationNameList:"AWSAccountId" type:"list" flattened:"true" required:"true"`
	Actions       []*string `locationNameList:"ActionName" type:"list" flattened:"true" required:"true"`
	Label         *string   `type:"string" required:"true"`
	QueueURL      *string   `locationName:"QueueUrl" type:"string" required:"true"`

	metadataAddPermissionInput `json:"-", xml:"-"`
}

type metadataAddPermissionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type AddPermissionOutput struct {
	metadataAddPermissionOutput `json:"-", xml:"-"`
}

type metadataAddPermissionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type BatchResultErrorEntry struct {
	Code        *string `type:"string" required:"true"`
	ID          *string `locationName:"Id" type:"string" required:"true"`
	Message     *string `type:"string"`
	SenderFault *bool   `type:"boolean" required:"true"`

	metadataBatchResultErrorEntry `json:"-", xml:"-"`
}

type metadataBatchResultErrorEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type ChangeMessageVisibilityBatchInput struct {
	Entries  []*ChangeMessageVisibilityBatchRequestEntry `locationNameList:"ChangeMessageVisibilityBatchRequestEntry" type:"list" flattened:"true" required:"true"`
	QueueURL *string                                     `locationName:"QueueUrl" type:"string" required:"true"`

	metadataChangeMessageVisibilityBatchInput `json:"-", xml:"-"`
}

type metadataChangeMessageVisibilityBatchInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ChangeMessageVisibilityBatchOutput struct {
	Failed     []*BatchResultErrorEntry                   `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`
	Successful []*ChangeMessageVisibilityBatchResultEntry `locationNameList:"ChangeMessageVisibilityBatchResultEntry" type:"list" flattened:"true" required:"true"`

	metadataChangeMessageVisibilityBatchOutput `json:"-", xml:"-"`
}

type metadataChangeMessageVisibilityBatchOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ChangeMessageVisibilityBatchRequestEntry struct {
	ID                *string `locationName:"Id" type:"string" required:"true"`
	ReceiptHandle     *string `type:"string" required:"true"`
	VisibilityTimeout *int64  `type:"integer"`

	metadataChangeMessageVisibilityBatchRequestEntry `json:"-", xml:"-"`
}

type metadataChangeMessageVisibilityBatchRequestEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type ChangeMessageVisibilityBatchResultEntry struct {
	ID *string `locationName:"Id" type:"string" required:"true"`

	metadataChangeMessageVisibilityBatchResultEntry `json:"-", xml:"-"`
}

type metadataChangeMessageVisibilityBatchResultEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type ChangeMessageVisibilityInput struct {
	QueueURL          *string `locationName:"QueueUrl" type:"string" required:"true"`
	ReceiptHandle     *string `type:"string" required:"true"`
	VisibilityTimeout *int64  `type:"integer" required:"true"`

	metadataChangeMessageVisibilityInput `json:"-", xml:"-"`
}

type metadataChangeMessageVisibilityInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ChangeMessageVisibilityOutput struct {
	metadataChangeMessageVisibilityOutput `json:"-", xml:"-"`
}

type metadataChangeMessageVisibilityOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateQueueInput struct {
	Attributes *map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
	QueueName  *string             `type:"string" required:"true"`

	metadataCreateQueueInput `json:"-", xml:"-"`
}

type metadataCreateQueueInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateQueueOutput struct {
	QueueURL *string `locationName:"QueueUrl" type:"string"`

	metadataCreateQueueOutput `json:"-", xml:"-"`
}

type metadataCreateQueueOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMessageBatchInput struct {
	Entries  []*DeleteMessageBatchRequestEntry `locationNameList:"DeleteMessageBatchRequestEntry" type:"list" flattened:"true" required:"true"`
	QueueURL *string                           `locationName:"QueueUrl" type:"string" required:"true"`

	metadataDeleteMessageBatchInput `json:"-", xml:"-"`
}

type metadataDeleteMessageBatchInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMessageBatchOutput struct {
	Failed     []*BatchResultErrorEntry         `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`
	Successful []*DeleteMessageBatchResultEntry `locationNameList:"DeleteMessageBatchResultEntry" type:"list" flattened:"true" required:"true"`

	metadataDeleteMessageBatchOutput `json:"-", xml:"-"`
}

type metadataDeleteMessageBatchOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMessageBatchRequestEntry struct {
	ID            *string `locationName:"Id" type:"string" required:"true"`
	ReceiptHandle *string `type:"string" required:"true"`

	metadataDeleteMessageBatchRequestEntry `json:"-", xml:"-"`
}

type metadataDeleteMessageBatchRequestEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMessageBatchResultEntry struct {
	ID *string `locationName:"Id" type:"string" required:"true"`

	metadataDeleteMessageBatchResultEntry `json:"-", xml:"-"`
}

type metadataDeleteMessageBatchResultEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMessageInput struct {
	QueueURL      *string `locationName:"QueueUrl" type:"string" required:"true"`
	ReceiptHandle *string `type:"string" required:"true"`

	metadataDeleteMessageInput `json:"-", xml:"-"`
}

type metadataDeleteMessageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteMessageOutput struct {
	metadataDeleteMessageOutput `json:"-", xml:"-"`
}

type metadataDeleteMessageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteQueueInput struct {
	QueueURL *string `locationName:"QueueUrl" type:"string" required:"true"`

	metadataDeleteQueueInput `json:"-", xml:"-"`
}

type metadataDeleteQueueInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteQueueOutput struct {
	metadataDeleteQueueOutput `json:"-", xml:"-"`
}

type metadataDeleteQueueOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetQueueAttributesInput struct {
	AttributeNames []*string `locationNameList:"AttributeName" type:"list" flattened:"true"`
	QueueURL       *string   `locationName:"QueueUrl" type:"string" required:"true"`

	metadataGetQueueAttributesInput `json:"-", xml:"-"`
}

type metadataGetQueueAttributesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetQueueAttributesOutput struct {
	Attributes *map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	metadataGetQueueAttributesOutput `json:"-", xml:"-"`
}

type metadataGetQueueAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetQueueURLInput struct {
	QueueName              *string `type:"string" required:"true"`
	QueueOwnerAWSAccountID *string `locationName:"QueueOwnerAWSAccountId" type:"string"`

	metadataGetQueueURLInput `json:"-", xml:"-"`
}

type metadataGetQueueURLInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetQueueURLOutput struct {
	QueueURL *string `locationName:"QueueUrl" type:"string"`

	metadataGetQueueURLOutput `json:"-", xml:"-"`
}

type metadataGetQueueURLOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDeadLetterSourceQueuesInput struct {
	QueueURL *string `locationName:"QueueUrl" type:"string" required:"true"`

	metadataListDeadLetterSourceQueuesInput `json:"-", xml:"-"`
}

type metadataListDeadLetterSourceQueuesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListDeadLetterSourceQueuesOutput struct {
	QueueURLs []*string `locationName:"queueUrls" locationNameList:"QueueUrl" type:"list" flattened:"true" required:"true"`

	metadataListDeadLetterSourceQueuesOutput `json:"-", xml:"-"`
}

type metadataListDeadLetterSourceQueuesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListQueuesInput struct {
	QueueNamePrefix *string `type:"string"`

	metadataListQueuesInput `json:"-", xml:"-"`
}

type metadataListQueuesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListQueuesOutput struct {
	QueueURLs []*string `locationName:"QueueUrls" locationNameList:"QueueUrl" type:"list" flattened:"true"`

	metadataListQueuesOutput `json:"-", xml:"-"`
}

type metadataListQueuesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Message struct {
	Attributes             *map[string]*string                `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
	Body                   *string                            `type:"string"`
	MD5OfBody              *string                            `type:"string"`
	MD5OfMessageAttributes *string                            `type:"string"`
	MessageAttributes      *map[string]*MessageAttributeValue `locationName:"MessageAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
	MessageID              *string                            `locationName:"MessageId" type:"string"`
	ReceiptHandle          *string                            `type:"string"`

	metadataMessage `json:"-", xml:"-"`
}

type metadataMessage struct {
	SDKShapeTraits bool `type:"structure"`
}

type MessageAttributeValue struct {
	BinaryListValues [][]byte  `locationName:"BinaryListValue" locationNameList:"BinaryListValue" type:"list" flattened:"true"`
	BinaryValue      []byte    `type:"blob"`
	DataType         *string   `type:"string" required:"true"`
	StringListValues []*string `locationName:"StringListValue" locationNameList:"StringListValue" type:"list" flattened:"true"`
	StringValue      *string   `type:"string"`

	metadataMessageAttributeValue `json:"-", xml:"-"`
}

type metadataMessageAttributeValue struct {
	SDKShapeTraits bool `type:"structure"`
}

type PurgeQueueInput struct {
	QueueURL *string `locationName:"QueueUrl" type:"string" required:"true"`

	metadataPurgeQueueInput `json:"-", xml:"-"`
}

type metadataPurgeQueueInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type PurgeQueueOutput struct {
	metadataPurgeQueueOutput `json:"-", xml:"-"`
}

type metadataPurgeQueueOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ReceiveMessageInput struct {
	AttributeNames        []*string `locationNameList:"AttributeName" type:"list" flattened:"true"`
	MaxNumberOfMessages   *int64    `type:"integer"`
	MessageAttributeNames []*string `locationNameList:"MessageAttributeName" type:"list" flattened:"true"`
	QueueURL              *string   `locationName:"QueueUrl" type:"string" required:"true"`
	VisibilityTimeout     *int64    `type:"integer"`
	WaitTimeSeconds       *int64    `type:"integer"`

	metadataReceiveMessageInput `json:"-", xml:"-"`
}

type metadataReceiveMessageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ReceiveMessageOutput struct {
	Messages []*Message `locationNameList:"Message" type:"list" flattened:"true"`

	metadataReceiveMessageOutput `json:"-", xml:"-"`
}

type metadataReceiveMessageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RemovePermissionInput struct {
	Label    *string `type:"string" required:"true"`
	QueueURL *string `locationName:"QueueUrl" type:"string" required:"true"`

	metadataRemovePermissionInput `json:"-", xml:"-"`
}

type metadataRemovePermissionInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RemovePermissionOutput struct {
	metadataRemovePermissionOutput `json:"-", xml:"-"`
}

type metadataRemovePermissionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendMessageBatchInput struct {
	Entries  []*SendMessageBatchRequestEntry `locationNameList:"SendMessageBatchRequestEntry" type:"list" flattened:"true" required:"true"`
	QueueURL *string                         `locationName:"QueueUrl" type:"string" required:"true"`

	metadataSendMessageBatchInput `json:"-", xml:"-"`
}

type metadataSendMessageBatchInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendMessageBatchOutput struct {
	Failed     []*BatchResultErrorEntry       `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`
	Successful []*SendMessageBatchResultEntry `locationNameList:"SendMessageBatchResultEntry" type:"list" flattened:"true" required:"true"`

	metadataSendMessageBatchOutput `json:"-", xml:"-"`
}

type metadataSendMessageBatchOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendMessageBatchRequestEntry struct {
	DelaySeconds      *int64                             `type:"integer"`
	ID                *string                            `locationName:"Id" type:"string" required:"true"`
	MessageAttributes *map[string]*MessageAttributeValue `locationName:"MessageAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
	MessageBody       *string                            `type:"string" required:"true"`

	metadataSendMessageBatchRequestEntry `json:"-", xml:"-"`
}

type metadataSendMessageBatchRequestEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendMessageBatchResultEntry struct {
	ID                     *string `locationName:"Id" type:"string" required:"true"`
	MD5OfMessageAttributes *string `type:"string"`
	MD5OfMessageBody       *string `type:"string" required:"true"`
	MessageID              *string `locationName:"MessageId" type:"string" required:"true"`

	metadataSendMessageBatchResultEntry `json:"-", xml:"-"`
}

type metadataSendMessageBatchResultEntry struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendMessageInput struct {
	DelaySeconds      *int64                             `type:"integer"`
	MessageAttributes *map[string]*MessageAttributeValue `locationName:"MessageAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
	MessageBody       *string                            `type:"string" required:"true"`
	QueueURL          *string                            `locationName:"QueueUrl" type:"string" required:"true"`

	metadataSendMessageInput `json:"-", xml:"-"`
}

type metadataSendMessageInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SendMessageOutput struct {
	MD5OfMessageAttributes *string `type:"string"`
	MD5OfMessageBody       *string `type:"string"`
	MessageID              *string `locationName:"MessageId" type:"string"`

	metadataSendMessageOutput `json:"-", xml:"-"`
}

type metadataSendMessageOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetQueueAttributesInput struct {
	Attributes *map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true" required:"true"`
	QueueURL   *string             `locationName:"QueueUrl" type:"string" required:"true"`

	metadataSetQueueAttributesInput `json:"-", xml:"-"`
}

type metadataSetQueueAttributesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetQueueAttributesOutput struct {
	metadataSetQueueAttributesOutput `json:"-", xml:"-"`
}

type metadataSetQueueAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}