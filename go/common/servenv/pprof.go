// Copyright 2019 The Vitess Authors.
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
//
// Modifications Copyright 2025 Supabase, Inc.

package servenv

import (
	"fmt"
	"io"
	"log/slog"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
)

type profmode string

const (
	profileCPU       profmode = "cpu"
	profileMemHeap   profmode = "mem_heap"
	profileMemAllocs profmode = "mem_allocs"
	profileMutex     profmode = "mutex"
	profileBlock     profmode = "block"
	profileTrace     profmode = "trace"
	profileThreads   profmode = "threads"
	profileGoroutine profmode = "goroutine"
)

func (p profmode) filename() string { _ = "STUB: not implemented"; return "" }

type profile struct {
	mode    profmode
	rate    int
	path    string
	quiet   bool
	waitSig bool
}

func (sv *ServEnv) parseProfileFlag(pf []string) (*profile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var profileStarted uint32

// isProfileStarted returns true if profiling is currently active.
// This function uses atomic.LoadUint32 to safely read the profile state.
func isProfileStarted() bool { _ = "STUB: not implemented"; return false }

func startCallback(start func() error) func() error { _ = "STUB: not implemented"; return nil }

func stopCallback(stop func()) func() { _ = "STUB: not implemented"; return nil }

func (prof *profile) mkprofile() (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// init returns a start function that begins the configured profiling process and
// returns a cleanup function that must be executed before process termination to
// flush the profile to disk.
// Based on the profiling code in github.com/pkg/profile
func (prof *profile) init() (start func() error, stop func()) {
	var pf io.WriteCloser

	switch prof.mode {
	case profileCPU:
		start = startCallback(func() error {
			var err error
			pf, err = prof.mkprofile()
			if err != nil {
				return err
			}
			if err := pprof.StartCPUProfile(pf); err != nil {
				return fmt.Errorf("pprof: could not start CPU profile: %w", err)
			}
			return nil
		})
		stop = stopCallback(func() {
			pprof.StopCPUProfile()
			pf.Close()
		})
		return start, stop

	case profileMemHeap, profileMemAllocs:
		old := runtime.MemProfileRate
		start = startCallback(func() error {
			var err error
			pf, err = prof.mkprofile()
			if err != nil {
				return err
			}
			runtime.MemProfileRate = prof.rate
			return nil
		})
		stop = stopCallback(func() {
			tt := "heap"
			if prof.mode == profileMemAllocs {
				tt = "allocs"
			}
			if err := pprof.Lookup(tt).WriteTo(pf, 0); err != nil {
				slog.Error("pprof: could not write memory profile", "err", err)
			}
			pf.Close()
			runtime.MemProfileRate = old
		})
		return start, stop

	case profileMutex:
		start = startCallback(func() error {
			var err error
			pf, err = prof.mkprofile()
			if err != nil {
				return err
			}
			runtime.SetMutexProfileFraction(prof.rate)
			return nil
		})
		stop = stopCallback(func() {
			if mp := pprof.Lookup("mutex"); mp != nil {
				if err := mp.WriteTo(pf, 0); err != nil {
					slog.Error("pprof: could not write mutex profile", "err", err)
				}
			}
			pf.Close()
			runtime.SetMutexProfileFraction(0)
		})
		return start, stop

	case profileBlock:
		start = startCallback(func() error {
			var err error
			pf, err = prof.mkprofile()
			if err != nil {
				return err
			}
			runtime.SetBlockProfileRate(prof.rate)
			return nil
		})
		stop = stopCallback(func() {
			if err := pprof.Lookup("block").WriteTo(pf, 0); err != nil {
				slog.Error("pprof: could not write block profile", "err", err)
			}
			pf.Close()
			runtime.SetBlockProfileRate(0)
		})
		return start, stop

	case profileThreads:
		start = startCallback(func() error {
			var err error
			pf, err = prof.mkprofile()
			return err
		})
		stop = stopCallback(func() {
			if mp := pprof.Lookup("threadcreate"); mp != nil {
				if err := mp.WriteTo(pf, 0); err != nil {
					slog.Error("pprof: could not write thread profile", "err", err)
				}
			}
			pf.Close()
		})
		return start, stop

	case profileTrace:
		start = startCallback(func() error {
			var err error
			pf, err = prof.mkprofile()
			if err != nil {
				return err
			}
			if err := trace.Start(pf); err != nil {
				return fmt.Errorf("pprof: could not start trace: %w", err)
			}
			return nil
		})
		stop = stopCallback(func() {
			trace.Stop()
			pf.Close()
		})
		return start, stop

	case profileGoroutine:
		start = startCallback(func() error {
			var err error
			pf, err = prof.mkprofile()
			return err
		})
		stop = stopCallback(func() {
			if mp := pprof.Lookup("goroutine"); mp != nil {
				if err := mp.WriteTo(pf, 0); err != nil {
					slog.Error("pprof: could not write goroutine profile", "err", err)
				}
			}
			pf.Close()
		})
		return start, stop

	default:
		panic("unsupported profile mode")
	}
}
