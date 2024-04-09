// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogsOutputGraylogStream struct {
	pulumi.CustomResourceState

	// Indicates if the current user can create alert on the stream
	CanAlert pulumi.BoolOutput `pulumi:"canAlert"`
	// Cold storage compression method
	ColdStorageCompression pulumi.StringOutput `pulumi:"coldStorageCompression"`
	// ColdStorage content
	ColdStorageContent pulumi.StringOutput `pulumi:"coldStorageContent"`
	// Is Cold storage enabled?
	ColdStorageEnabled pulumi.BoolOutput `pulumi:"coldStorageEnabled"`
	// Notify on new Cold storage archive
	ColdStorageNotifyEnabled pulumi.BoolOutput `pulumi:"coldStorageNotifyEnabled"`
	// Cold storage retention in year
	ColdStorageRetention pulumi.IntOutput `pulumi:"coldStorageRetention"`
	// ColdStorage destination
	ColdStorageTarget pulumi.StringOutput `pulumi:"coldStorageTarget"`
	// Stream creation
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Stream description
	Description pulumi.StringOutput `pulumi:"description"`
	// Enable ES indexing
	IndexingEnabled pulumi.BoolOutput `pulumi:"indexingEnabled"`
	// Maximum indexing size (in GB)
	IndexingMaxSize pulumi.IntOutput `pulumi:"indexingMaxSize"`
	// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
	IndexingNotifyEnabled pulumi.BoolOutput `pulumi:"indexingNotifyEnabled"`
	// Indicates if you are allowed to edit entry
	IsEditable pulumi.BoolOutput `pulumi:"isEditable"`
	// Indicates if you are allowed to share entry
	IsShareable pulumi.BoolOutput `pulumi:"isShareable"`
	// Number of alert condition
	NbAlertCondition pulumi.IntOutput `pulumi:"nbAlertCondition"`
	// Number of coldstored archives
	NbArchive pulumi.IntOutput `pulumi:"nbArchive"`
	// Parent stream ID
	ParentStreamId pulumi.StringPtrOutput `pulumi:"parentStreamId"`
	// If set, pause indexing when maximum size is reach
	PauseIndexingOnMaxSize pulumi.BoolOutput `pulumi:"pauseIndexingOnMaxSize"`
	// Retention ID
	RetentionId pulumi.StringOutput `pulumi:"retentionId"`
	// The service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Stream ID
	StreamId pulumi.StringOutput `pulumi:"streamId"`
	// Stream description
	Title pulumi.StringOutput `pulumi:"title"`
	// Stream last update
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Enable Websocket
	WebSocketEnabled pulumi.BoolOutput `pulumi:"webSocketEnabled"`
}

// NewLogsOutputGraylogStream registers a new resource with the given unique name, arguments, and options.
func NewLogsOutputGraylogStream(ctx *pulumi.Context,
	name string, args *LogsOutputGraylogStreamArgs, opts ...pulumi.ResourceOption) (*LogsOutputGraylogStream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogsOutputGraylogStream
	err := ctx.RegisterResource("ovh:Dbaas/logsOutputGraylogStream:LogsOutputGraylogStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogsOutputGraylogStream gets an existing LogsOutputGraylogStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogsOutputGraylogStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogsOutputGraylogStreamState, opts ...pulumi.ResourceOption) (*LogsOutputGraylogStream, error) {
	var resource LogsOutputGraylogStream
	err := ctx.ReadResource("ovh:Dbaas/logsOutputGraylogStream:LogsOutputGraylogStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogsOutputGraylogStream resources.
type logsOutputGraylogStreamState struct {
	// Indicates if the current user can create alert on the stream
	CanAlert *bool `pulumi:"canAlert"`
	// Cold storage compression method
	ColdStorageCompression *string `pulumi:"coldStorageCompression"`
	// ColdStorage content
	ColdStorageContent *string `pulumi:"coldStorageContent"`
	// Is Cold storage enabled?
	ColdStorageEnabled *bool `pulumi:"coldStorageEnabled"`
	// Notify on new Cold storage archive
	ColdStorageNotifyEnabled *bool `pulumi:"coldStorageNotifyEnabled"`
	// Cold storage retention in year
	ColdStorageRetention *int `pulumi:"coldStorageRetention"`
	// ColdStorage destination
	ColdStorageTarget *string `pulumi:"coldStorageTarget"`
	// Stream creation
	CreatedAt *string `pulumi:"createdAt"`
	// Stream description
	Description *string `pulumi:"description"`
	// Enable ES indexing
	IndexingEnabled *bool `pulumi:"indexingEnabled"`
	// Maximum indexing size (in GB)
	IndexingMaxSize *int `pulumi:"indexingMaxSize"`
	// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
	IndexingNotifyEnabled *bool `pulumi:"indexingNotifyEnabled"`
	// Indicates if you are allowed to edit entry
	IsEditable *bool `pulumi:"isEditable"`
	// Indicates if you are allowed to share entry
	IsShareable *bool `pulumi:"isShareable"`
	// Number of alert condition
	NbAlertCondition *int `pulumi:"nbAlertCondition"`
	// Number of coldstored archives
	NbArchive *int `pulumi:"nbArchive"`
	// Parent stream ID
	ParentStreamId *string `pulumi:"parentStreamId"`
	// If set, pause indexing when maximum size is reach
	PauseIndexingOnMaxSize *bool `pulumi:"pauseIndexingOnMaxSize"`
	// Retention ID
	RetentionId *string `pulumi:"retentionId"`
	// The service name
	ServiceName *string `pulumi:"serviceName"`
	// Stream ID
	StreamId *string `pulumi:"streamId"`
	// Stream description
	Title *string `pulumi:"title"`
	// Stream last update
	UpdatedAt *string `pulumi:"updatedAt"`
	// Enable Websocket
	WebSocketEnabled *bool `pulumi:"webSocketEnabled"`
}

type LogsOutputGraylogStreamState struct {
	// Indicates if the current user can create alert on the stream
	CanAlert pulumi.BoolPtrInput
	// Cold storage compression method
	ColdStorageCompression pulumi.StringPtrInput
	// ColdStorage content
	ColdStorageContent pulumi.StringPtrInput
	// Is Cold storage enabled?
	ColdStorageEnabled pulumi.BoolPtrInput
	// Notify on new Cold storage archive
	ColdStorageNotifyEnabled pulumi.BoolPtrInput
	// Cold storage retention in year
	ColdStorageRetention pulumi.IntPtrInput
	// ColdStorage destination
	ColdStorageTarget pulumi.StringPtrInput
	// Stream creation
	CreatedAt pulumi.StringPtrInput
	// Stream description
	Description pulumi.StringPtrInput
	// Enable ES indexing
	IndexingEnabled pulumi.BoolPtrInput
	// Maximum indexing size (in GB)
	IndexingMaxSize pulumi.IntPtrInput
	// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
	IndexingNotifyEnabled pulumi.BoolPtrInput
	// Indicates if you are allowed to edit entry
	IsEditable pulumi.BoolPtrInput
	// Indicates if you are allowed to share entry
	IsShareable pulumi.BoolPtrInput
	// Number of alert condition
	NbAlertCondition pulumi.IntPtrInput
	// Number of coldstored archives
	NbArchive pulumi.IntPtrInput
	// Parent stream ID
	ParentStreamId pulumi.StringPtrInput
	// If set, pause indexing when maximum size is reach
	PauseIndexingOnMaxSize pulumi.BoolPtrInput
	// Retention ID
	RetentionId pulumi.StringPtrInput
	// The service name
	ServiceName pulumi.StringPtrInput
	// Stream ID
	StreamId pulumi.StringPtrInput
	// Stream description
	Title pulumi.StringPtrInput
	// Stream last update
	UpdatedAt pulumi.StringPtrInput
	// Enable Websocket
	WebSocketEnabled pulumi.BoolPtrInput
}

func (LogsOutputGraylogStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*logsOutputGraylogStreamState)(nil)).Elem()
}

type logsOutputGraylogStreamArgs struct {
	// Cold storage compression method
	ColdStorageCompression *string `pulumi:"coldStorageCompression"`
	// ColdStorage content
	ColdStorageContent *string `pulumi:"coldStorageContent"`
	// Is Cold storage enabled?
	ColdStorageEnabled *bool `pulumi:"coldStorageEnabled"`
	// Notify on new Cold storage archive
	ColdStorageNotifyEnabled *bool `pulumi:"coldStorageNotifyEnabled"`
	// Cold storage retention in year
	ColdStorageRetention *int `pulumi:"coldStorageRetention"`
	// ColdStorage destination
	ColdStorageTarget *string `pulumi:"coldStorageTarget"`
	// Stream description
	Description string `pulumi:"description"`
	// Enable ES indexing
	IndexingEnabled *bool `pulumi:"indexingEnabled"`
	// Maximum indexing size (in GB)
	IndexingMaxSize *int `pulumi:"indexingMaxSize"`
	// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
	IndexingNotifyEnabled *bool `pulumi:"indexingNotifyEnabled"`
	// Parent stream ID
	ParentStreamId *string `pulumi:"parentStreamId"`
	// If set, pause indexing when maximum size is reach
	PauseIndexingOnMaxSize *bool `pulumi:"pauseIndexingOnMaxSize"`
	// Retention ID
	RetentionId *string `pulumi:"retentionId"`
	// The service name
	ServiceName string `pulumi:"serviceName"`
	// Stream description
	Title string `pulumi:"title"`
	// Enable Websocket
	WebSocketEnabled *bool `pulumi:"webSocketEnabled"`
}

// The set of arguments for constructing a LogsOutputGraylogStream resource.
type LogsOutputGraylogStreamArgs struct {
	// Cold storage compression method
	ColdStorageCompression pulumi.StringPtrInput
	// ColdStorage content
	ColdStorageContent pulumi.StringPtrInput
	// Is Cold storage enabled?
	ColdStorageEnabled pulumi.BoolPtrInput
	// Notify on new Cold storage archive
	ColdStorageNotifyEnabled pulumi.BoolPtrInput
	// Cold storage retention in year
	ColdStorageRetention pulumi.IntPtrInput
	// ColdStorage destination
	ColdStorageTarget pulumi.StringPtrInput
	// Stream description
	Description pulumi.StringInput
	// Enable ES indexing
	IndexingEnabled pulumi.BoolPtrInput
	// Maximum indexing size (in GB)
	IndexingMaxSize pulumi.IntPtrInput
	// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
	IndexingNotifyEnabled pulumi.BoolPtrInput
	// Parent stream ID
	ParentStreamId pulumi.StringPtrInput
	// If set, pause indexing when maximum size is reach
	PauseIndexingOnMaxSize pulumi.BoolPtrInput
	// Retention ID
	RetentionId pulumi.StringPtrInput
	// The service name
	ServiceName pulumi.StringInput
	// Stream description
	Title pulumi.StringInput
	// Enable Websocket
	WebSocketEnabled pulumi.BoolPtrInput
}

func (LogsOutputGraylogStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logsOutputGraylogStreamArgs)(nil)).Elem()
}

type LogsOutputGraylogStreamInput interface {
	pulumi.Input

	ToLogsOutputGraylogStreamOutput() LogsOutputGraylogStreamOutput
	ToLogsOutputGraylogStreamOutputWithContext(ctx context.Context) LogsOutputGraylogStreamOutput
}

func (*LogsOutputGraylogStream) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsOutputGraylogStream)(nil)).Elem()
}

func (i *LogsOutputGraylogStream) ToLogsOutputGraylogStreamOutput() LogsOutputGraylogStreamOutput {
	return i.ToLogsOutputGraylogStreamOutputWithContext(context.Background())
}

func (i *LogsOutputGraylogStream) ToLogsOutputGraylogStreamOutputWithContext(ctx context.Context) LogsOutputGraylogStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputGraylogStreamOutput)
}

// LogsOutputGraylogStreamArrayInput is an input type that accepts LogsOutputGraylogStreamArray and LogsOutputGraylogStreamArrayOutput values.
// You can construct a concrete instance of `LogsOutputGraylogStreamArrayInput` via:
//
//	LogsOutputGraylogStreamArray{ LogsOutputGraylogStreamArgs{...} }
type LogsOutputGraylogStreamArrayInput interface {
	pulumi.Input

	ToLogsOutputGraylogStreamArrayOutput() LogsOutputGraylogStreamArrayOutput
	ToLogsOutputGraylogStreamArrayOutputWithContext(context.Context) LogsOutputGraylogStreamArrayOutput
}

type LogsOutputGraylogStreamArray []LogsOutputGraylogStreamInput

func (LogsOutputGraylogStreamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsOutputGraylogStream)(nil)).Elem()
}

func (i LogsOutputGraylogStreamArray) ToLogsOutputGraylogStreamArrayOutput() LogsOutputGraylogStreamArrayOutput {
	return i.ToLogsOutputGraylogStreamArrayOutputWithContext(context.Background())
}

func (i LogsOutputGraylogStreamArray) ToLogsOutputGraylogStreamArrayOutputWithContext(ctx context.Context) LogsOutputGraylogStreamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputGraylogStreamArrayOutput)
}

// LogsOutputGraylogStreamMapInput is an input type that accepts LogsOutputGraylogStreamMap and LogsOutputGraylogStreamMapOutput values.
// You can construct a concrete instance of `LogsOutputGraylogStreamMapInput` via:
//
//	LogsOutputGraylogStreamMap{ "key": LogsOutputGraylogStreamArgs{...} }
type LogsOutputGraylogStreamMapInput interface {
	pulumi.Input

	ToLogsOutputGraylogStreamMapOutput() LogsOutputGraylogStreamMapOutput
	ToLogsOutputGraylogStreamMapOutputWithContext(context.Context) LogsOutputGraylogStreamMapOutput
}

type LogsOutputGraylogStreamMap map[string]LogsOutputGraylogStreamInput

func (LogsOutputGraylogStreamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsOutputGraylogStream)(nil)).Elem()
}

func (i LogsOutputGraylogStreamMap) ToLogsOutputGraylogStreamMapOutput() LogsOutputGraylogStreamMapOutput {
	return i.ToLogsOutputGraylogStreamMapOutputWithContext(context.Background())
}

func (i LogsOutputGraylogStreamMap) ToLogsOutputGraylogStreamMapOutputWithContext(ctx context.Context) LogsOutputGraylogStreamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputGraylogStreamMapOutput)
}

type LogsOutputGraylogStreamOutput struct{ *pulumi.OutputState }

func (LogsOutputGraylogStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsOutputGraylogStream)(nil)).Elem()
}

func (o LogsOutputGraylogStreamOutput) ToLogsOutputGraylogStreamOutput() LogsOutputGraylogStreamOutput {
	return o
}

func (o LogsOutputGraylogStreamOutput) ToLogsOutputGraylogStreamOutputWithContext(ctx context.Context) LogsOutputGraylogStreamOutput {
	return o
}

// Indicates if the current user can create alert on the stream
func (o LogsOutputGraylogStreamOutput) CanAlert() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.CanAlert }).(pulumi.BoolOutput)
}

// Cold storage compression method
func (o LogsOutputGraylogStreamOutput) ColdStorageCompression() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.ColdStorageCompression }).(pulumi.StringOutput)
}

// ColdStorage content
func (o LogsOutputGraylogStreamOutput) ColdStorageContent() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.ColdStorageContent }).(pulumi.StringOutput)
}

// Is Cold storage enabled?
func (o LogsOutputGraylogStreamOutput) ColdStorageEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.ColdStorageEnabled }).(pulumi.BoolOutput)
}

// Notify on new Cold storage archive
func (o LogsOutputGraylogStreamOutput) ColdStorageNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.ColdStorageNotifyEnabled }).(pulumi.BoolOutput)
}

// Cold storage retention in year
func (o LogsOutputGraylogStreamOutput) ColdStorageRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.IntOutput { return v.ColdStorageRetention }).(pulumi.IntOutput)
}

// ColdStorage destination
func (o LogsOutputGraylogStreamOutput) ColdStorageTarget() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.ColdStorageTarget }).(pulumi.StringOutput)
}

// Stream creation
func (o LogsOutputGraylogStreamOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Stream description
func (o LogsOutputGraylogStreamOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Enable ES indexing
func (o LogsOutputGraylogStreamOutput) IndexingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.IndexingEnabled }).(pulumi.BoolOutput)
}

// Maximum indexing size (in GB)
func (o LogsOutputGraylogStreamOutput) IndexingMaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.IntOutput { return v.IndexingMaxSize }).(pulumi.IntOutput)
}

// If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
func (o LogsOutputGraylogStreamOutput) IndexingNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.IndexingNotifyEnabled }).(pulumi.BoolOutput)
}

// Indicates if you are allowed to edit entry
func (o LogsOutputGraylogStreamOutput) IsEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.IsEditable }).(pulumi.BoolOutput)
}

// Indicates if you are allowed to share entry
func (o LogsOutputGraylogStreamOutput) IsShareable() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.IsShareable }).(pulumi.BoolOutput)
}

// Number of alert condition
func (o LogsOutputGraylogStreamOutput) NbAlertCondition() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.IntOutput { return v.NbAlertCondition }).(pulumi.IntOutput)
}

// Number of coldstored archives
func (o LogsOutputGraylogStreamOutput) NbArchive() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.IntOutput { return v.NbArchive }).(pulumi.IntOutput)
}

// Parent stream ID
func (o LogsOutputGraylogStreamOutput) ParentStreamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringPtrOutput { return v.ParentStreamId }).(pulumi.StringPtrOutput)
}

// If set, pause indexing when maximum size is reach
func (o LogsOutputGraylogStreamOutput) PauseIndexingOnMaxSize() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.PauseIndexingOnMaxSize }).(pulumi.BoolOutput)
}

// Retention ID
func (o LogsOutputGraylogStreamOutput) RetentionId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.RetentionId }).(pulumi.StringOutput)
}

// The service name
func (o LogsOutputGraylogStreamOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Stream ID
func (o LogsOutputGraylogStreamOutput) StreamId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.StreamId }).(pulumi.StringOutput)
}

// Stream description
func (o LogsOutputGraylogStreamOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// Stream last update
func (o LogsOutputGraylogStreamOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Enable Websocket
func (o LogsOutputGraylogStreamOutput) WebSocketEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputGraylogStream) pulumi.BoolOutput { return v.WebSocketEnabled }).(pulumi.BoolOutput)
}

type LogsOutputGraylogStreamArrayOutput struct{ *pulumi.OutputState }

func (LogsOutputGraylogStreamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsOutputGraylogStream)(nil)).Elem()
}

func (o LogsOutputGraylogStreamArrayOutput) ToLogsOutputGraylogStreamArrayOutput() LogsOutputGraylogStreamArrayOutput {
	return o
}

func (o LogsOutputGraylogStreamArrayOutput) ToLogsOutputGraylogStreamArrayOutputWithContext(ctx context.Context) LogsOutputGraylogStreamArrayOutput {
	return o
}

func (o LogsOutputGraylogStreamArrayOutput) Index(i pulumi.IntInput) LogsOutputGraylogStreamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogsOutputGraylogStream {
		return vs[0].([]*LogsOutputGraylogStream)[vs[1].(int)]
	}).(LogsOutputGraylogStreamOutput)
}

type LogsOutputGraylogStreamMapOutput struct{ *pulumi.OutputState }

func (LogsOutputGraylogStreamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsOutputGraylogStream)(nil)).Elem()
}

func (o LogsOutputGraylogStreamMapOutput) ToLogsOutputGraylogStreamMapOutput() LogsOutputGraylogStreamMapOutput {
	return o
}

func (o LogsOutputGraylogStreamMapOutput) ToLogsOutputGraylogStreamMapOutputWithContext(ctx context.Context) LogsOutputGraylogStreamMapOutput {
	return o
}

func (o LogsOutputGraylogStreamMapOutput) MapIndex(k pulumi.StringInput) LogsOutputGraylogStreamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogsOutputGraylogStream {
		return vs[0].(map[string]*LogsOutputGraylogStream)[vs[1].(string)]
	}).(LogsOutputGraylogStreamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputGraylogStreamInput)(nil)).Elem(), &LogsOutputGraylogStream{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputGraylogStreamArrayInput)(nil)).Elem(), LogsOutputGraylogStreamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputGraylogStreamMapInput)(nil)).Elem(), LogsOutputGraylogStreamMap{})
	pulumi.RegisterOutputType(LogsOutputGraylogStreamOutput{})
	pulumi.RegisterOutputType(LogsOutputGraylogStreamArrayOutput{})
	pulumi.RegisterOutputType(LogsOutputGraylogStreamMapOutput{})
}
