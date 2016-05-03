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

// EventKeyState  Maps input key states
type EventKeyState struct {
  MoveTurbo bool

  MoveUp    bool
  MoveDown  bool
  MoveRight bool
  MoveLeft  bool

  WatchX    float32
  WatchY    float32
}

// InstanceEventKeyState  Stores input key states
var InstanceEventKeyState = EventKeyState{false, false, false, false, false, 0.0, 0.0}

func getEventKeyState() (*EventKeyState) {
  return &InstanceEventKeyState
}

func handleKey(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
  keyState := getEventKeyState()

  // Main controls
  if k == glfw.KeyEscape {
    if action == glfw.Press {
      window.SetShouldClose(true)
    }
  }

  // Camera control keys
  if k == glfw.KeyR {
    // Release?
    if action == glfw.Release {
      // Immediately reset camera
      resetCameraObject()
      resetCamera()
    }
  }

  if k == glfw.KeyTab {
    // Release?
    if action == glfw.Release {
      // Toggle to next camera object
      toggleNextCameraObject()
    }
  }

  // Camera controls
  if k == glfw.KeySpace {
    // Press or release?
    if action == glfw.Press {
      keyState.MoveTurbo = true
    } else if action == glfw.Release {
      keyState.MoveTurbo = false
    }
  }

  if k == glfw.KeyUp {
    // Press or release?
    if action == glfw.Press {
      keyState.MoveDown = false
      keyState.MoveUp = true
    } else if action == glfw.Release {
      keyState.MoveUp = false
    }
  }

  if k == glfw.KeyDown {
    // Press or release?
    if action == glfw.Press {
      keyState.MoveUp = false
      keyState.MoveDown = true
    } else if action == glfw.Release {
      keyState.MoveDown = false
    }
  }

  if k == glfw.KeyLeft {
    // Press or release?
    if action == glfw.Press {
      keyState.MoveRight = false
      keyState.MoveLeft = true
    } else if action == glfw.Release {
      keyState.MoveLeft = false
    }
  }

  if k == glfw.KeyRight {
    // Press or release?
    if action == glfw.Press {
      keyState.MoveLeft = false
      keyState.MoveRight = true
    } else if action == glfw.Release {
      keyState.MoveRight = false
    }
  }
}

func handleMouseCursor(window *glfw.Window, positionX float64, positionY float64) {
  keyState := getEventKeyState()
  windowData := getWindowData()

  // Bind new watch position
  if positionX >= 0 && positionX <= float64(windowData.Width) {
    keyState.WatchX = float32(positionX) * (1.0 / float32(windowData.Width)) - 0.5
  }

  if positionY >= 0 && positionY <= float64(windowData.Height) {
    keyState.WatchY = float32(positionY) * (1.0 / float32(windowData.Height)) - 0.5
  }
}

func handleMouseScroll(window *glfw.Window, offsetX float64, offsetY float64) {
  // Update scene simulation speed
  updateSpeedFactor(offsetY)
}

func resetMouseCursor() {
  keyState := getEventKeyState()

  keyState.WatchX = 0.0
  keyState.WatchY = 0.0
}
