// Copyright 2022 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-defaults. DO NOT EDIT.

package test

import (
	"context"
	"errors"
	"io"

	"go.linka.cloud/protoc-gen-client/tests/pb/external"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ Test = (*clientTest)(nil)

type Test interface {
	UnaryEmpty(ctx context.Context, opts ...grpc.CallOption) (err error)
	UnaryReqParams(ctx context.Context, msg *Message, opts ...grpc.CallOption) (err error)
	UnaryResParams(ctx context.Context, opts ...grpc.CallOption) (Msg *Message, err error)
	UnaryOneOfParams(ctx context.Context, oneOf isUnaryOneOfParamsMsg_OneOf, opts ...grpc.CallOption) (OneOf isUnaryOneOfParamsMsg_OneOf, err error)
	UnaryParams(ctx context.Context, msg *Message, opts ...grpc.CallOption) (Msg *Message, err error)
	UnaryExternal(ctx context.Context, name ext.Name, value string, opts ...grpc.CallOption) (Name ext.Name, Value string, err error)
	UnaryParamsAny(ctx context.Context, any *anypb.Any, string *String, int64 *int64, opts ...grpc.CallOption) (Any *anypb.Any, String_ *String, Int64 *int64, err error)
	ServerStream(ctx context.Context, msg *Message, opts ...grpc.CallOption) (<-chan *ServerStreamMsg, error)
}

func NewTest(cc grpc.ClientConnInterface) Test {
	return &clientTest{c: NewTestClient(cc)}
}

type clientTest struct {
	c TestClient
}

// UnaryEmpty ...
func (x *clientTest) UnaryEmpty(ctx context.Context, opts ...grpc.CallOption) (err error) {
	_, err = x.c.UnaryEmpty(ctx, &UnaryEmptyRequest{}, opts...)
	err = x.unwrap(err)
	if err != nil {
		return
	}
	return nil
}

// UnaryReqParams ...
func (x *clientTest) UnaryReqParams(ctx context.Context, msg *Message, opts ...grpc.CallOption) (err error) {
	_, err = x.c.UnaryReqParams(ctx, &UnaryRequestParams{Msg: msg}, opts...)
	err = x.unwrap(err)
	if err != nil {
		return
	}
	return nil
}

// UnaryResParams ...
func (x *clientTest) UnaryResParams(ctx context.Context, opts ...grpc.CallOption) (Msg *Message, err error) {
	var res *UnaryResponseParams
	res, err = x.c.UnaryResParams(ctx, &UnaryEmptyRequest{}, opts...)
	err = x.unwrap(err)
	if err != nil {
		return
	}
	return res.Msg, nil
}

// UnaryOneOfParams ...
func (x *clientTest) UnaryOneOfParams(ctx context.Context, oneOf isUnaryOneOfParamsMsg_OneOf, opts ...grpc.CallOption) (OneOf isUnaryOneOfParamsMsg_OneOf, err error) {
	var res *UnaryOneOfParamsMsg
	res, err = x.c.UnaryOneOfParams(ctx, &UnaryOneOfParamsMsg{OneOf: oneOf}, opts...)
	err = x.unwrap(err)
	if err != nil {
		return
	}
	return res.OneOf, nil
}

// UnaryParams ...
func (x *clientTest) UnaryParams(ctx context.Context, msg *Message, opts ...grpc.CallOption) (Msg *Message, err error) {
	var res *UnaryResponseParams
	res, err = x.c.UnaryParams(ctx, &UnaryRequestParams{Msg: msg}, opts...)
	err = x.unwrap(err)
	if err != nil {
		return
	}
	return res.Msg, nil
}

// UnaryExternal ...
func (x *clientTest) UnaryExternal(ctx context.Context, name ext.Name, value string, opts ...grpc.CallOption) (Name ext.Name, Value string, err error) {
	var res *ext.External
	res, err = x.c.UnaryExternal(ctx, &ext.External{Key: name, Val: value}, opts...)
	err = x.unwrap(err)
	if err != nil {
		return
	}
	return res.Key, res.Val, nil
}

// UnaryParamsAny ...
func (x *clientTest) UnaryParamsAny(ctx context.Context, any *anypb.Any, string *String, int64 *int64, opts ...grpc.CallOption) (Any *anypb.Any, String_ *String, Int64 *int64, err error) {
	var res *UnaryResponseParamsAny
	res, err = x.c.UnaryParamsAny(ctx, &UnaryRequestParamsAny{Any: any, String_: (*String)(string), Int64: int64}, opts...)
	err = x.unwrap(err)
	if err != nil {
		return
	}
	return res.Any, res.String_, res.Int64, nil
}

type ServerStreamMsg struct {
	Msg *UnaryResponseParams
	Err error
}

func (x *clientTest) ServerStream(ctx context.Context, msg *Message, opts ...grpc.CallOption) (cn <-chan *ServerStreamMsg, err error) {
	ss, err := x.c.ServerStream(ctx, &UnaryRequestParams{Msg: msg}, opts...)
	if err != nil {
		return nil, err
	}
	ch := make(chan *ServerStreamMsg, 1)
	go func() {
		defer close(ch)
		for {
			res, err := ss.Recv()
			if errors.Is(io.EOF, err) {
				return
			}
			if err != nil {
				ch <- &ServerStreamMsg{Err: err}
				break
			}
			ch <- &ServerStreamMsg{Msg: res}
		}
	}()
	return ch, nil
}

// unwrap convert grpc status error to go error
func (x *clientTest) unwrap(err error) error {
	s, ok := status.FromError(err)
	if !ok {
		return err
	}
	if s != nil {
		return errors.New(s.Message())
	}
	return nil
}
