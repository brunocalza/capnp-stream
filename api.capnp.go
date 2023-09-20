// Code generated by capnpc-go. DO NOT EDIT.

package stream

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
)

type FileUploader capnp.Client

// FileUploader_TypeID is the unique identifier for the type FileUploader.
const FileUploader_TypeID = 0xb558a5aeed63f2a4

func (c FileUploader) Upload(ctx context.Context, params func(FileUploader_upload_Params) error) (FileUploader_upload_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xb558a5aeed63f2a4,
			MethodID:      0,
			InterfaceName: "api.capnp:FileUploader",
			MethodName:    "upload",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(FileUploader_upload_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return FileUploader_upload_Results_Future{Future: ans.Future()}, release

}

func (c FileUploader) WaitStreaming() error {
	return capnp.Client(c).WaitStreaming()
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c FileUploader) String() string {
	return "FileUploader(" + capnp.Client(c).String() + ")"
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c FileUploader) AddRef() FileUploader {
	return FileUploader(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c FileUploader) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c FileUploader) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c FileUploader) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (FileUploader) DecodeFromPtr(p capnp.Ptr) FileUploader {
	return FileUploader(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c FileUploader) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c FileUploader) IsSame(other FileUploader) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c FileUploader) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c FileUploader) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A FileUploader_Server is a FileUploader with a local implementation.
type FileUploader_Server interface {
	Upload(context.Context, FileUploader_upload) error
}

// FileUploader_NewServer creates a new Server from an implementation of FileUploader_Server.
func FileUploader_NewServer(s FileUploader_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(FileUploader_Methods(nil, s), s, c)
}

// FileUploader_ServerToClient creates a new Client from an implementation of FileUploader_Server.
// The caller is responsible for calling Release on the returned Client.
func FileUploader_ServerToClient(s FileUploader_Server) FileUploader {
	return FileUploader(capnp.NewClient(FileUploader_NewServer(s)))
}

// FileUploader_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func FileUploader_Methods(methods []server.Method, s FileUploader_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb558a5aeed63f2a4,
			MethodID:      0,
			InterfaceName: "api.capnp:FileUploader",
			MethodName:    "upload",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Upload(ctx, FileUploader_upload{call})
		},
	})

	return methods
}

// FileUploader_upload holds the state for a server call to FileUploader.upload.
// See server.Call for documentation.
type FileUploader_upload struct {
	*server.Call
}

// Args returns the call's arguments.
func (c FileUploader_upload) Args() FileUploader_upload_Params {
	return FileUploader_upload_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c FileUploader_upload) AllocResults() (FileUploader_upload_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FileUploader_upload_Results(r), err
}

// FileUploader_List is a list of FileUploader.
type FileUploader_List = capnp.CapList[FileUploader]

// NewFileUploader creates a new list of FileUploader.
func NewFileUploader_List(s *capnp.Segment, sz int32) (FileUploader_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[FileUploader](l), err
}

type FileUploader_Callback capnp.Client

// FileUploader_Callback_TypeID is the unique identifier for the type FileUploader_Callback.
const FileUploader_Callback_TypeID = 0xc0bb050d13ac1ce0

func (c FileUploader_Callback) Write(ctx context.Context, params func(FileUploader_Callback_write_Params) error) (FileUploader_Callback_write_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xc0bb050d13ac1ce0,
			MethodID:      0,
			InterfaceName: "api.capnp:FileUploader.Callback",
			MethodName:    "write",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(FileUploader_Callback_write_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return FileUploader_Callback_write_Results_Future{Future: ans.Future()}, release

}

func (c FileUploader_Callback) Done(ctx context.Context, params func(FileUploader_Callback_done_Params) error) (FileUploader_Callback_done_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xc0bb050d13ac1ce0,
			MethodID:      1,
			InterfaceName: "api.capnp:FileUploader.Callback",
			MethodName:    "done",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(FileUploader_Callback_done_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return FileUploader_Callback_done_Results_Future{Future: ans.Future()}, release

}

func (c FileUploader_Callback) WaitStreaming() error {
	return capnp.Client(c).WaitStreaming()
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c FileUploader_Callback) String() string {
	return "FileUploader_Callback(" + capnp.Client(c).String() + ")"
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c FileUploader_Callback) AddRef() FileUploader_Callback {
	return FileUploader_Callback(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c FileUploader_Callback) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c FileUploader_Callback) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c FileUploader_Callback) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (FileUploader_Callback) DecodeFromPtr(p capnp.Ptr) FileUploader_Callback {
	return FileUploader_Callback(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c FileUploader_Callback) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c FileUploader_Callback) IsSame(other FileUploader_Callback) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c FileUploader_Callback) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c FileUploader_Callback) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A FileUploader_Callback_Server is a FileUploader_Callback with a local implementation.
type FileUploader_Callback_Server interface {
	Write(context.Context, FileUploader_Callback_write) error

	Done(context.Context, FileUploader_Callback_done) error
}

// FileUploader_Callback_NewServer creates a new Server from an implementation of FileUploader_Callback_Server.
func FileUploader_Callback_NewServer(s FileUploader_Callback_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(FileUploader_Callback_Methods(nil, s), s, c)
}

// FileUploader_Callback_ServerToClient creates a new Client from an implementation of FileUploader_Callback_Server.
// The caller is responsible for calling Release on the returned Client.
func FileUploader_Callback_ServerToClient(s FileUploader_Callback_Server) FileUploader_Callback {
	return FileUploader_Callback(capnp.NewClient(FileUploader_Callback_NewServer(s)))
}

// FileUploader_Callback_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func FileUploader_Callback_Methods(methods []server.Method, s FileUploader_Callback_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc0bb050d13ac1ce0,
			MethodID:      0,
			InterfaceName: "api.capnp:FileUploader.Callback",
			MethodName:    "write",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Write(ctx, FileUploader_Callback_write{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xc0bb050d13ac1ce0,
			MethodID:      1,
			InterfaceName: "api.capnp:FileUploader.Callback",
			MethodName:    "done",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Done(ctx, FileUploader_Callback_done{call})
		},
	})

	return methods
}

// FileUploader_Callback_write holds the state for a server call to FileUploader_Callback.write.
// See server.Call for documentation.
type FileUploader_Callback_write struct {
	*server.Call
}

// Args returns the call's arguments.
func (c FileUploader_Callback_write) Args() FileUploader_Callback_write_Params {
	return FileUploader_Callback_write_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c FileUploader_Callback_write) AllocResults() (FileUploader_Callback_write_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_write_Results(r), err
}

// FileUploader_Callback_done holds the state for a server call to FileUploader_Callback.done.
// See server.Call for documentation.
type FileUploader_Callback_done struct {
	*server.Call
}

// Args returns the call's arguments.
func (c FileUploader_Callback_done) Args() FileUploader_Callback_done_Params {
	return FileUploader_Callback_done_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c FileUploader_Callback_done) AllocResults() (FileUploader_Callback_done_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_done_Results(r), err
}

// FileUploader_Callback_List is a list of FileUploader_Callback.
type FileUploader_Callback_List = capnp.CapList[FileUploader_Callback]

// NewFileUploader_Callback creates a new list of FileUploader_Callback.
func NewFileUploader_Callback_List(s *capnp.Segment, sz int32) (FileUploader_Callback_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[FileUploader_Callback](l), err
}

type FileUploader_Callback_write_Params capnp.Struct

// FileUploader_Callback_write_Params_TypeID is the unique identifier for the type FileUploader_Callback_write_Params.
const FileUploader_Callback_write_Params_TypeID = 0xa68acdbb4df99282

func NewFileUploader_Callback_write_Params(s *capnp.Segment) (FileUploader_Callback_write_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FileUploader_Callback_write_Params(st), err
}

func NewRootFileUploader_Callback_write_Params(s *capnp.Segment) (FileUploader_Callback_write_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FileUploader_Callback_write_Params(st), err
}

func ReadRootFileUploader_Callback_write_Params(msg *capnp.Message) (FileUploader_Callback_write_Params, error) {
	root, err := msg.Root()
	return FileUploader_Callback_write_Params(root.Struct()), err
}

func (s FileUploader_Callback_write_Params) String() string {
	str, _ := text.Marshal(0xa68acdbb4df99282, capnp.Struct(s))
	return str
}

func (s FileUploader_Callback_write_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (FileUploader_Callback_write_Params) DecodeFromPtr(p capnp.Ptr) FileUploader_Callback_write_Params {
	return FileUploader_Callback_write_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s FileUploader_Callback_write_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s FileUploader_Callback_write_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s FileUploader_Callback_write_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s FileUploader_Callback_write_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s FileUploader_Callback_write_Params) Chunk() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s FileUploader_Callback_write_Params) HasChunk() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s FileUploader_Callback_write_Params) SetChunk(v []byte) error {
	return capnp.Struct(s).SetData(0, v)
}

// FileUploader_Callback_write_Params_List is a list of FileUploader_Callback_write_Params.
type FileUploader_Callback_write_Params_List = capnp.StructList[FileUploader_Callback_write_Params]

// NewFileUploader_Callback_write_Params creates a new list of FileUploader_Callback_write_Params.
func NewFileUploader_Callback_write_Params_List(s *capnp.Segment, sz int32) (FileUploader_Callback_write_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[FileUploader_Callback_write_Params](l), err
}

// FileUploader_Callback_write_Params_Future is a wrapper for a FileUploader_Callback_write_Params promised by a client call.
type FileUploader_Callback_write_Params_Future struct{ *capnp.Future }

func (f FileUploader_Callback_write_Params_Future) Struct() (FileUploader_Callback_write_Params, error) {
	p, err := f.Future.Ptr()
	return FileUploader_Callback_write_Params(p.Struct()), err
}

type FileUploader_Callback_write_Results capnp.Struct

// FileUploader_Callback_write_Results_TypeID is the unique identifier for the type FileUploader_Callback_write_Results.
const FileUploader_Callback_write_Results_TypeID = 0x922c3107c5f7024b

func NewFileUploader_Callback_write_Results(s *capnp.Segment) (FileUploader_Callback_write_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_write_Results(st), err
}

func NewRootFileUploader_Callback_write_Results(s *capnp.Segment) (FileUploader_Callback_write_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_write_Results(st), err
}

func ReadRootFileUploader_Callback_write_Results(msg *capnp.Message) (FileUploader_Callback_write_Results, error) {
	root, err := msg.Root()
	return FileUploader_Callback_write_Results(root.Struct()), err
}

func (s FileUploader_Callback_write_Results) String() string {
	str, _ := text.Marshal(0x922c3107c5f7024b, capnp.Struct(s))
	return str
}

func (s FileUploader_Callback_write_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (FileUploader_Callback_write_Results) DecodeFromPtr(p capnp.Ptr) FileUploader_Callback_write_Results {
	return FileUploader_Callback_write_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s FileUploader_Callback_write_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s FileUploader_Callback_write_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s FileUploader_Callback_write_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s FileUploader_Callback_write_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// FileUploader_Callback_write_Results_List is a list of FileUploader_Callback_write_Results.
type FileUploader_Callback_write_Results_List = capnp.StructList[FileUploader_Callback_write_Results]

// NewFileUploader_Callback_write_Results creates a new list of FileUploader_Callback_write_Results.
func NewFileUploader_Callback_write_Results_List(s *capnp.Segment, sz int32) (FileUploader_Callback_write_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[FileUploader_Callback_write_Results](l), err
}

// FileUploader_Callback_write_Results_Future is a wrapper for a FileUploader_Callback_write_Results promised by a client call.
type FileUploader_Callback_write_Results_Future struct{ *capnp.Future }

func (f FileUploader_Callback_write_Results_Future) Struct() (FileUploader_Callback_write_Results, error) {
	p, err := f.Future.Ptr()
	return FileUploader_Callback_write_Results(p.Struct()), err
}

type FileUploader_Callback_done_Params capnp.Struct

// FileUploader_Callback_done_Params_TypeID is the unique identifier for the type FileUploader_Callback_done_Params.
const FileUploader_Callback_done_Params_TypeID = 0x93363d0da3c75d18

func NewFileUploader_Callback_done_Params(s *capnp.Segment) (FileUploader_Callback_done_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_done_Params(st), err
}

func NewRootFileUploader_Callback_done_Params(s *capnp.Segment) (FileUploader_Callback_done_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_done_Params(st), err
}

func ReadRootFileUploader_Callback_done_Params(msg *capnp.Message) (FileUploader_Callback_done_Params, error) {
	root, err := msg.Root()
	return FileUploader_Callback_done_Params(root.Struct()), err
}

func (s FileUploader_Callback_done_Params) String() string {
	str, _ := text.Marshal(0x93363d0da3c75d18, capnp.Struct(s))
	return str
}

func (s FileUploader_Callback_done_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (FileUploader_Callback_done_Params) DecodeFromPtr(p capnp.Ptr) FileUploader_Callback_done_Params {
	return FileUploader_Callback_done_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s FileUploader_Callback_done_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s FileUploader_Callback_done_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s FileUploader_Callback_done_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s FileUploader_Callback_done_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// FileUploader_Callback_done_Params_List is a list of FileUploader_Callback_done_Params.
type FileUploader_Callback_done_Params_List = capnp.StructList[FileUploader_Callback_done_Params]

// NewFileUploader_Callback_done_Params creates a new list of FileUploader_Callback_done_Params.
func NewFileUploader_Callback_done_Params_List(s *capnp.Segment, sz int32) (FileUploader_Callback_done_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[FileUploader_Callback_done_Params](l), err
}

// FileUploader_Callback_done_Params_Future is a wrapper for a FileUploader_Callback_done_Params promised by a client call.
type FileUploader_Callback_done_Params_Future struct{ *capnp.Future }

func (f FileUploader_Callback_done_Params_Future) Struct() (FileUploader_Callback_done_Params, error) {
	p, err := f.Future.Ptr()
	return FileUploader_Callback_done_Params(p.Struct()), err
}

type FileUploader_Callback_done_Results capnp.Struct

// FileUploader_Callback_done_Results_TypeID is the unique identifier for the type FileUploader_Callback_done_Results.
const FileUploader_Callback_done_Results_TypeID = 0xe41ec3aacd3dcd3b

func NewFileUploader_Callback_done_Results(s *capnp.Segment) (FileUploader_Callback_done_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_done_Results(st), err
}

func NewRootFileUploader_Callback_done_Results(s *capnp.Segment) (FileUploader_Callback_done_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return FileUploader_Callback_done_Results(st), err
}

func ReadRootFileUploader_Callback_done_Results(msg *capnp.Message) (FileUploader_Callback_done_Results, error) {
	root, err := msg.Root()
	return FileUploader_Callback_done_Results(root.Struct()), err
}

func (s FileUploader_Callback_done_Results) String() string {
	str, _ := text.Marshal(0xe41ec3aacd3dcd3b, capnp.Struct(s))
	return str
}

func (s FileUploader_Callback_done_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (FileUploader_Callback_done_Results) DecodeFromPtr(p capnp.Ptr) FileUploader_Callback_done_Results {
	return FileUploader_Callback_done_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s FileUploader_Callback_done_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s FileUploader_Callback_done_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s FileUploader_Callback_done_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s FileUploader_Callback_done_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// FileUploader_Callback_done_Results_List is a list of FileUploader_Callback_done_Results.
type FileUploader_Callback_done_Results_List = capnp.StructList[FileUploader_Callback_done_Results]

// NewFileUploader_Callback_done_Results creates a new list of FileUploader_Callback_done_Results.
func NewFileUploader_Callback_done_Results_List(s *capnp.Segment, sz int32) (FileUploader_Callback_done_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[FileUploader_Callback_done_Results](l), err
}

// FileUploader_Callback_done_Results_Future is a wrapper for a FileUploader_Callback_done_Results promised by a client call.
type FileUploader_Callback_done_Results_Future struct{ *capnp.Future }

func (f FileUploader_Callback_done_Results_Future) Struct() (FileUploader_Callback_done_Results, error) {
	p, err := f.Future.Ptr()
	return FileUploader_Callback_done_Results(p.Struct()), err
}

type FileUploader_upload_Params capnp.Struct

// FileUploader_upload_Params_TypeID is the unique identifier for the type FileUploader_upload_Params.
const FileUploader_upload_Params_TypeID = 0x837f1f0ba3d44020

func NewFileUploader_upload_Params(s *capnp.Segment) (FileUploader_upload_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FileUploader_upload_Params(st), err
}

func NewRootFileUploader_upload_Params(s *capnp.Segment) (FileUploader_upload_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FileUploader_upload_Params(st), err
}

func ReadRootFileUploader_upload_Params(msg *capnp.Message) (FileUploader_upload_Params, error) {
	root, err := msg.Root()
	return FileUploader_upload_Params(root.Struct()), err
}

func (s FileUploader_upload_Params) String() string {
	str, _ := text.Marshal(0x837f1f0ba3d44020, capnp.Struct(s))
	return str
}

func (s FileUploader_upload_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (FileUploader_upload_Params) DecodeFromPtr(p capnp.Ptr) FileUploader_upload_Params {
	return FileUploader_upload_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s FileUploader_upload_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s FileUploader_upload_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s FileUploader_upload_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s FileUploader_upload_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s FileUploader_upload_Params) Filename() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s FileUploader_upload_Params) HasFilename() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s FileUploader_upload_Params) FilenameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s FileUploader_upload_Params) SetFilename(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// FileUploader_upload_Params_List is a list of FileUploader_upload_Params.
type FileUploader_upload_Params_List = capnp.StructList[FileUploader_upload_Params]

// NewFileUploader_upload_Params creates a new list of FileUploader_upload_Params.
func NewFileUploader_upload_Params_List(s *capnp.Segment, sz int32) (FileUploader_upload_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[FileUploader_upload_Params](l), err
}

// FileUploader_upload_Params_Future is a wrapper for a FileUploader_upload_Params promised by a client call.
type FileUploader_upload_Params_Future struct{ *capnp.Future }

func (f FileUploader_upload_Params_Future) Struct() (FileUploader_upload_Params, error) {
	p, err := f.Future.Ptr()
	return FileUploader_upload_Params(p.Struct()), err
}

type FileUploader_upload_Results capnp.Struct

// FileUploader_upload_Results_TypeID is the unique identifier for the type FileUploader_upload_Results.
const FileUploader_upload_Results_TypeID = 0xf6aa737f6605f8e2

func NewFileUploader_upload_Results(s *capnp.Segment) (FileUploader_upload_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FileUploader_upload_Results(st), err
}

func NewRootFileUploader_upload_Results(s *capnp.Segment) (FileUploader_upload_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return FileUploader_upload_Results(st), err
}

func ReadRootFileUploader_upload_Results(msg *capnp.Message) (FileUploader_upload_Results, error) {
	root, err := msg.Root()
	return FileUploader_upload_Results(root.Struct()), err
}

func (s FileUploader_upload_Results) String() string {
	str, _ := text.Marshal(0xf6aa737f6605f8e2, capnp.Struct(s))
	return str
}

func (s FileUploader_upload_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (FileUploader_upload_Results) DecodeFromPtr(p capnp.Ptr) FileUploader_upload_Results {
	return FileUploader_upload_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s FileUploader_upload_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s FileUploader_upload_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s FileUploader_upload_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s FileUploader_upload_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s FileUploader_upload_Results) Callback() FileUploader_Callback {
	p, _ := capnp.Struct(s).Ptr(0)
	return FileUploader_Callback(p.Interface().Client())
}

func (s FileUploader_upload_Results) HasCallback() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s FileUploader_upload_Results) SetCallback(v FileUploader_Callback) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// FileUploader_upload_Results_List is a list of FileUploader_upload_Results.
type FileUploader_upload_Results_List = capnp.StructList[FileUploader_upload_Results]

// NewFileUploader_upload_Results creates a new list of FileUploader_upload_Results.
func NewFileUploader_upload_Results_List(s *capnp.Segment, sz int32) (FileUploader_upload_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[FileUploader_upload_Results](l), err
}

// FileUploader_upload_Results_Future is a wrapper for a FileUploader_upload_Results promised by a client call.
type FileUploader_upload_Results_Future struct{ *capnp.Future }

func (f FileUploader_upload_Results_Future) Struct() (FileUploader_upload_Results, error) {
	p, err := f.Future.Ptr()
	return FileUploader_upload_Results(p.Struct()), err
}
func (p FileUploader_upload_Results_Future) Callback() FileUploader_Callback {
	return FileUploader_Callback(p.Future.Field(0, nil).Client())
}

const schema_9cf9878fd3dd8473 = "x\xda\x8c\x92?h\x13q\x1c\xc5\xdf\xfb\xde\x9dW5" +
	"G\xf8qB*\x88\xe9\x90\"\x16\x0dM\x0a\x0e\x95\x90" +
	"\x80\xa0P)\xe4\x0a\x05\x1d\x1c\xae\xe9\x15\xd3^\xd2\x90" +
	"k\x10\xa7\xe2\x1f(\x08\"\xd4\xd5\xad\xd6\xa1\x88\x9b\x8b" +
	"]\x04\x11\xa7L\xba\x0a\"\xae.\x0ej\x069\xb9\x84" +
	"KR1\xda\xed\xe0\xfb\xbd\xcf\xef}\xdf{\xd3\x16K" +
	"z\xce:\xa3C\x9ci\xe3H8Q\xfa\xb0s<\xbd" +
	"y\x0f*E\xc0\xa0\x09\xcct8E\xd0\xa6\x14\xc1\xf0" +
	"\xaa\xfcxk\xe6\xcemCM\x10\xd0\xa3\xf9\xa44\x09" +
	"=\x1c\xbf\xf1n\xc7*\\x<4Q\xb2\x14M\xee" +
	"nw\xe6\xf7\xdb\x0f\x9e\xf5&=\xe8/\xaeF\xd0\xa3" +
	"]\xe8\xd3o\x95\xaf/v\xaf\xbd\x84:\xa6\x85\xc1\xfd" +
	"\x8f\xef\x1fmu\x9e\x00\xb4\xcf\xca+;')\xc0." +
	"\xc8\x15\xdb\x13\x13\x08?\x9dzn[\xc6\xfek\xa8\x94" +
	"6\xf8\x13\x9c\x99\x17\xa1}=\xda\xb1\x17e\xcb\xde\xed" +
	"n_l\x17\xda{oN\x7f\x19R\xf5PV#U" +
	"\x9f\x7f\x1a+\x9b\xc1\xde\xf7\xe1SoK>RuG" +
	"\x8a8\x1f\xba\x8dj\xb6\xe26\xeaZc\xf6r\xd5\xf7" +
	"\x16\x1b\xfe\xba\xbb\xec5\xb3\xad\xeeG\xa6\xec&\x9bn" +
	"-ptM\x07t\x02\xca\x9a\x03\x9c\x84Fg\\\x18" +
	"\xaeT}\xaf\xee\xd6<\x00L@\x98\x00\xfbD\xfd\x0f" +
	"\xe2%\xd7\xf7\x97\xdc\xcaZ\xf6V\xb3\xba\xe1e\x16\xbc" +
	"t\xd0\xf27\x82\xff\xef/\xaf\xd7\xbdL\xd9m\x9an" +
	"-8,\xfd/\xba\xf3\x803\xa6\xd19!LWn" +
	"\xb6\xeak\xb4 \xb4\x86\x143f\x16{PG'\x87" +
	"\x92\xe0\\\x18\xbf\x02\xc0\xd15\x03\xe8w\x89\xb1\xd3J" +
	"\xcdB\x94a\x16{\x0e\x96X&G\xba\xdc\xe7\x95I" +
	"g\xac\x0b\x8c{\xc4\xb8\x84*\x97\x87\xa8I\x93\xec\xb7" +
	"\x8fq\xe0\xea\xe4\x14DYf\xba{u\x89\xc9\xc8\xab" +
	"\x83O\xfe\xdb\xd6\x05/H\x1eHaD\x0f\xe2\xb4F" +
	"\x15\xa120\x86j\xe0\x19H\x05\xfe\x0e\x00\x00\xff\xff" +
	"|~\x02D"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_9cf9878fd3dd8473,
		Nodes: []uint64{
			0x837f1f0ba3d44020,
			0x922c3107c5f7024b,
			0x93363d0da3c75d18,
			0xa68acdbb4df99282,
			0xb558a5aeed63f2a4,
			0xc0bb050d13ac1ce0,
			0xe41ec3aacd3dcd3b,
			0xf6aa737f6605f8e2,
		},
		Compressed: true,
	})
}
