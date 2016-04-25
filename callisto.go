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
  "log"
  "os"
  "runtime"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/glfw/v3.1/glfw"
)

func init() {
  // GLFW event handling must run on the main OS thread
  runtime.LockOSThread()

  dir, err := importPathToDir("github.com/valeriansaliou/callisto")
  if err != nil {
    log.Fatalln("Unable to find Go package in your GOPATH; needed to load assets:", err)
  }

  err = os.Chdir(dir)
  if err != nil {
    log.Panicln("os.Chdir:", err)
  }
}

func main() {
  var (
    err     error
    window  *glfw.Window
    vao     uint32
  )

  // Create window
  if err = glfw.Init(); err != nil {
      log.Fatalln("Failed to initialize glfw:", err)
  }
  defer glfw.Terminate()

  glfw.WindowHint(glfw.Resizable, glfw.False)
  glfw.WindowHint(glfw.ContextVersionMajor, 4)
  glfw.WindowHint(glfw.ContextVersionMinor, 1)
  glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
  glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

  window, err = glfw.CreateWindow(WINDOW_WIDTH, WINDOW_HEIGHT, WINDOW_TITLE, nil, nil)
  if err != nil {
    panic(err)
  }

  window.MakeContextCurrent()
  window.SetKeyCallback(handleKey)
  window.SetCursorPosCallback(handleMouseCursor)

  // Initialize OpenGL
  gl.Init()

  // Configure the shaders program
  program, err := newProgram(SHADER_VERTEX, SHADER_FRAGMENT)
  if err != nil {
    panic(err)
  }

  gl.UseProgram(program)

  // Initialize stack matrix
  initializeMatrix()

  // Create environment
  createProjection(program)
  createCamera(program)

  // Create the VAO (Vertex Array Objects)
  // Notice: this stores links between attributes and active vertex data
  gl.GenVertexArrays(1, &vao)

  // Load the map of stellar objects
  objects := loadObjects("solar-system")

  // Create each object buffers
  createAllBuffers(objects, program, vao)

  // Initialize shaders
  initializeShaders(program)

  // Configure global settings
  gl.Enable(gl.DEPTH_TEST)
  gl.Enable(gl.TEXTURE_2D)
  gl.DepthFunc(gl.LESS)
  gl.ClearColor(0.025, 0.025, 0.025, 1.0)

  // Render loop
  for !window.ShouldClose() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    // Global routines
    updateElaspedTime(glfw.GetTime())
    gl.UseProgram(program)

    // Update context
    updateCamera()

    // Bind context
    bindProjection()
    bindCamera()

    // Render all objects in the map
    renderObjects(objects, program)

    glfw.PollEvents()
    window.SwapBuffers()
  }
}
