// Copyright (c) 2008-2018, Hazelcast, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protocol

import (
	. "github.com/hazelcast/hazelcast-go-client/internal/common"
	. "github.com/hazelcast/hazelcast-go-client/internal/serialization"
)

func QueueDrainToMaxSizeCalculateSize(name *string, maxSize int32) int {
	// Calculates the request payload size
	dataSize := 0
	dataSize += StringCalculateSize(name)
	dataSize += INT32_SIZE_IN_BYTES
	return dataSize
}

func QueueDrainToMaxSizeEncodeRequest(name *string, maxSize int32) *ClientMessage {
	// Encode request into clientMessage
	clientMessage := NewClientMessage(nil, QueueDrainToMaxSizeCalculateSize(name, maxSize))
	clientMessage.SetMessageType(QUEUE_DRAINTOMAXSIZE)
	clientMessage.IsRetryable = false
	clientMessage.AppendString(name)
	clientMessage.AppendInt32(maxSize)
	clientMessage.UpdateFrameLength()
	return clientMessage
}

func QueueDrainToMaxSizeDecodeResponse(clientMessage *ClientMessage) func() (response []*Data) {
	// Decode response from client message
	return func() (response []*Data) {
		responseSize := clientMessage.ReadInt32()
		response = make([]*Data, responseSize)
		for responseIndex := 0; responseIndex < int(responseSize); responseIndex++ {
			responseItem := clientMessage.ReadData()
			response[responseIndex] = responseItem
		}
		return
	}
}
