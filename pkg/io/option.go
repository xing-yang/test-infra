/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package io

import (
	"cloud.google.com/go/storage"
	"gocloud.dev/blob"
	"google.golang.org/api/googleapi"
)

// WriterOptions are options for the opener Writer method
type WriterOptions struct {
	BufferSize      *int64
	ContentEncoding *string
	ContentType     *string
	Metadata        map[string]string
}

// Apply applies the WriterOptions to storage.Writer and blob.WriterOptions
// Both arguments are allowed to be nil.
func (wo WriterOptions) Apply(writer *storage.Writer, o *blob.WriterOptions) {
	if writer != nil {
		if wo.BufferSize != nil {
			if *wo.BufferSize < googleapi.DefaultUploadChunkSize {
				writer.ChunkSize = int(*wo.BufferSize)
			}
		}
		if wo.ContentEncoding != nil {
			writer.ObjectAttrs.ContentEncoding = *wo.ContentEncoding
		}
		if wo.ContentType != nil {
			writer.ObjectAttrs.ContentType = *wo.ContentType
		}
		if wo.Metadata != nil {
			writer.ObjectAttrs.Metadata = wo.Metadata
		}
	}

	if o == nil {
		return
	}

	if wo.BufferSize != nil {
		o.BufferSize = int(*wo.BufferSize)
	}
	if wo.ContentEncoding != nil {
		o.ContentEncoding = *wo.ContentEncoding
	}
	if wo.ContentType != nil {
		o.ContentType = *wo.ContentType
	}
	if wo.Metadata != nil {
		o.Metadata = wo.Metadata
	}
}