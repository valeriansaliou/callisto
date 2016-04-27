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

// WindowData  Maps data on current window
type WindowData struct {
  Width  int
  Height int
}

// InstanceWindowData  Stores data on current window
var InstanceWindowData WindowData

func getWindowData() (*WindowData) {
  return &InstanceWindowData
}

func initializeWindow(monitor *glfw.Monitor) {
  windowData := getWindowData()
  videoMode := monitor.GetVideoMode()

  // Initialize window size
  windowData.Width = videoMode.Width
  windowData.Height = videoMode.Height

  // Lock window framerate to monitor framerate
  getSpeed().setFramerate(videoMode.RefreshRate)
}

func adjustWindow(window *glfw.Window) {
  framebufferWidth, framebufferHeight := window.GetFramebufferSize()

  handleAdjustWindow(window, framebufferWidth, framebufferHeight)
}

func handleAdjustWindow(window *glfw.Window, width int, height int) {
  windowData := getWindowData()

  // Adjust window size
  windowData.Width = width
  windowData.Height = height
}
