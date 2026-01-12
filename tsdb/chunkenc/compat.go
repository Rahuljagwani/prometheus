// Copyright The Prometheus Authors
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

package chunkenc

import (
	"github.com/prometheus/prometheus/model/histogram"
)

// CompatAppenderV2 provides AppenderV2 interface for chunks that do not
// natively support it.
type CompatAppenderV2 struct {
	Appender Appender
}

func (a *CompatAppenderV2) Append(_, t int64, v float64) {
	a.Appender.Append(t, v)
}

func (a *CompatAppenderV2) AppendHistogram(prev *HistogramAppender, _, t int64, h *histogram.Histogram, appendOnly bool) (c Chunk, isRecoded bool, app Appender, err error) {
	return a.Appender.AppendHistogram(prev, t, h, appendOnly)
}

func (a *CompatAppenderV2) AppendFloatHistogram(prev *FloatHistogramAppender, _, t int64, h *histogram.FloatHistogram, appendOnly bool) (c Chunk, isRecoded bool, app Appender, err error) {
	return a.Appender.AppendFloatHistogram(prev, t, h, appendOnly)
}
