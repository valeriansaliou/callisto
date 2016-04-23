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
  "github.com/go-gl/mathgl/mgl32"
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
    vbo_sphere_vertices  uint32
    vbo_sphere_indices   uint32
    angle   float32
    time    float64
    elapsed float64
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

  // Initialize OpenGL
  gl.Init()

  // Configure the shaders program
  program, err := newProgram(vertexShader, fragmentShader)
  if err != nil {
    panic(err)
  }

  gl.UseProgram(program)

  // Create the view projection
  projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(WINDOW_WIDTH) / WINDOW_HEIGHT, 0.1, 10.0)
  projectionUniform := gl.GetUniformLocation(program, gl.Str("projectionUniform\x00"))

  // Create the normal matrix
  normal := mgl32.Mat4Normal(projection)
  normalUniform := gl.GetUniformLocation(program, gl.Str("normalUniform\x00"))

  // Create the camera
  camera := mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
  cameraUniform := gl.GetUniformLocation(program, gl.Str("cameraUniform\x00"))

  // Create the model (used for rotation)
  model := mgl32.Ident4()
  modelUniform := gl.GetUniformLocation(program, gl.Str("modelUniform\x00"))

  // Create lighting
  lighting := mgl32.Vec4{2, 2, -2, 1}
  lightingUniform := gl.GetUniformLocation(program, gl.Str("lightingUniform\x00"))

  // Color storage
  gl.BindFragDataLocation(program, 0, gl.Str("objectColor\x00"))

  // Generate the sphere
  var sphere_vertices, sphere_indices = generateSphere(128, 64)

  // Create the VAO (Vertex Array Objects)
  // Notice: this stores links between attributes and active vertex data
  gl.GenVertexArrays(1, &vao)
  gl.BindVertexArray(vao)

  // Create the VBO (Vertex Buffer Object)
  // Notice: this passes buffer data to the GPU (cache to GPU for I/O performance)
  gl.GenBuffers(1, &vbo_sphere_vertices)
  gl.BindBuffer(gl.ARRAY_BUFFER, vbo_sphere_vertices)
  gl.BufferData(gl.ARRAY_BUFFER, len(sphere_vertices)*4, gl.Ptr(sphere_vertices), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &vbo_sphere_indices)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo_sphere_indices)
  gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(sphere_indices)*4, gl.Ptr(sphere_indices), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

  // Bind buffer to shaders attributes
  vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertAttrib\x00")))

  // Configure global settings
  gl.Enable(gl.DEPTH_TEST)
  gl.DepthFunc(gl.LESS)
  gl.ClearColor(0.2, 0.2, 0.2, 1.0)

  // Model angle
  angle = 0.0
  previousTime := glfw.GetTime()

  // Render loop
  for !window.ShouldClose() {
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

    // Calculate angle (relative to elapsed time)
    time = glfw.GetTime()
    elapsed = time - previousTime
    previousTime = time
    angle += float32(elapsed / 4)

    // Process model
    model = mgl32.HomogRotate3D(angle, mgl32.Vec3{0, 1, 0})

    gl.UseProgram(program)

    // Process matrixes
    gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])
    gl.UniformMatrix4fv(normalUniform, 1, false, &normal[0])
    gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])
    gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])
    gl.UniformMatrix4fv(lightingUniform, 1, false, &lighting[0])
    gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

    // Render buffers
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo_sphere_vertices)
    gl.EnableVertexAttribArray(vertAttrib)
    gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, vbo_sphere_indices)

    // Draw elements
    gl.DrawElements(gl.TRIANGLES, int32(len(sphere_indices) * 2), gl.UNSIGNED_SHORT, gl.PtrOffset(0))

    // Reset buffers
    gl.DisableVertexAttribArray(vertAttrib)
    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

    glfw.PollEvents()
    window.SwapBuffers()
  }
}
