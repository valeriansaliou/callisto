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
  Up    bool
  Down  bool
  Right bool
  Left  bool
}

var EVENT_KEY_STATE = EventKeyState{false, false, false, false}

func getEventKeyState() (*EventKeyState) {
  return &EVENT_KEY_STATE
}

func handleKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
  if action != glfw.Press {
    return
  }

  // Main controls
  if k == glfw.KeyEscape {
    window.SetShouldClose(true)
  }

  // Camera controls
  if k == glfw.KeyUp {
    EVENT_KEY_STATE.Down = false
    EVENT_KEY_STATE.Up = flipBool(EVENT_KEY_STATE.Up)
  }
  if k == glfw.KeyDown {
    EVENT_KEY_STATE.Up = false
    EVENT_KEY_STATE.Down = flipBool(EVENT_KEY_STATE.Down)
  }
  if k == glfw.KeyLeft {
    EVENT_KEY_STATE.Right = false
    EVENT_KEY_STATE.Left = flipBool(EVENT_KEY_STATE.Left)
  }
  if k == glfw.KeyRight {
    EVENT_KEY_STATE.Left = false
    EVENT_KEY_STATE.Right = flipBool(EVENT_KEY_STATE.Right)
  }
}
