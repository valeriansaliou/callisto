/* Callisto - Yet another Solar System simulator
 *
 * Copyright (c) 2016, Valerian Saliou <valerian@valeriansaliou.name>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *   * Redistributions of source code must retain the above copyright notice,
 *     this list of conditions and the following disclaimer.
 *   * Redistributions in binary form must reproduce the above copyright
 *     notice, this list of conditions and the following disclaimer in the
 *     documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
  "github.com/go-gl/glfw/v3.1/glfw"
)

type EventKeyState struct {
  MoveUp    bool
  MoveDown  bool
  MoveRight bool
  MoveLeft  bool

  WatchX    float32
  WatchY    float32
}

var EVENT_KEY_STATE = EventKeyState{false, false, false, false, float32(WINDOW_WIDTH) / 2.0, float32(WINDOW_HEIGHT) / 2.0}

func getEventKeyState() (*EventKeyState) {
  return &EVENT_KEY_STATE
}

func handleKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
  // Main controls
  if k == glfw.KeyEscape {
    if action == glfw.Press {
      window.SetShouldClose(true)
    }
  }

  // Camera controls
  if k == glfw.KeyUp {
    // Press or release?
    if action == glfw.Press {
      EVENT_KEY_STATE.MoveDown = false
      EVENT_KEY_STATE.MoveUp = true
    } else if action == glfw.Release {
      EVENT_KEY_STATE.MoveUp = false
    }
  }

  if k == glfw.KeyDown {
    // Press or release?
    if action == glfw.Press {
      EVENT_KEY_STATE.MoveUp = false
      EVENT_KEY_STATE.MoveDown = true
    } else if action == glfw.Release {
      EVENT_KEY_STATE.MoveDown = false
    }
  }

  if k == glfw.KeyLeft {
    // Press or release?
    if action == glfw.Press {
      EVENT_KEY_STATE.MoveRight = false
      EVENT_KEY_STATE.MoveLeft = true
    } else if action == glfw.Release {
      EVENT_KEY_STATE.MoveLeft = false
    }
  }

  if k == glfw.KeyRight {
    // Press or release?
    if action == glfw.Press {
      EVENT_KEY_STATE.MoveLeft = false
      EVENT_KEY_STATE.MoveRight = true
    } else if action == glfw.Release {
      EVENT_KEY_STATE.MoveRight = false
    }
  }
}

func handleMouseCursor(window *glfw.Window, position_x float64, position_y float64) {
  EVENT_KEY_STATE.WatchX = float32(position_x) * CAMERA_WATCH_REDUCER
  EVENT_KEY_STATE.WatchY = float32(position_y) * CAMERA_WATCH_REDUCER
}