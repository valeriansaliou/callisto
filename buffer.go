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

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type Buffers struct {
  Sphere              Sphere
  Texture             uint32

  Model               mgl32.Mat4
  ModelUniform        int32

  TextureUniform      int32

  VBOSphereVertices   uint32
  VBOSphereTexture    uint32
  VBOSphereIndices    uint32

  VertexAttributes    uint32
  VertexTextureCoords uint32
}

var BUFFERS map[string]Buffers = make(map[string]Buffers)

func getBuffers(name string) (Buffers) {
  return BUFFERS[name]
}

func createAllBuffers(objects []Object, program uint32, vao uint32) {
  for o := range objects {
    // Create the object buffers
    createBuffers(objects[o], program, vao)
  }
}

func createBuffers(object Object, program uint32, vao uint32) {
  var (
    buffers Buffers
    err     error
  )

  // Generate sphere
  buffers.Sphere = generateSphere(object.Name, object.Radius)

  // Load texture
  buffers.Texture, err = loadTexture(object.Name)

  if err != nil {
    log.Fatalln(err)
  }

  // Create the model (used for rotation)
  buffers.Model = mgl32.Ident4()
  buffers.ModelUniform = gl.GetUniformLocation(program, gl.Str("modelUniform\x00"))

  // Create the texture storage
  buffers.TextureUniform = gl.GetUniformLocation(program, gl.Str("textureUniform\x00"))
  gl.Uniform1i(buffers.TextureUniform, 0)

  // Color storage
  gl.BindFragDataLocation(program, 0, gl.Str("objectColor\x00"))

  // Bind object to VAO
  gl.BindVertexArray(vao)

  // Create the VBO (Vertex Buffer Object)
  // Notice: this passes buffer data to the GPU (cache to GPU for I/O performance)
  gl.GenBuffers(1, &buffers.VBOSphereVertices)
  gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereVertices)
  gl.BufferData(gl.ARRAY_BUFFER, len(buffers.Sphere.Vertices)*4, gl.Ptr(buffers.Sphere.Vertices), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &buffers.VBOSphereTexture)
  gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereTexture)
  gl.BufferData(gl.ARRAY_BUFFER, len(buffers.Sphere.TextureCoords)*4, gl.Ptr(buffers.Sphere.TextureCoords), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &buffers.VBOSphereIndices)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers.VBOSphereIndices)
  gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(buffers.Sphere.Indices)*4, gl.Ptr(buffers.Sphere.Indices), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

  // Bind buffer to shaders attributes
  buffers.VertexAttributes = uint32(gl.GetAttribLocation(program, gl.Str("vertexAttributes\x00")))
  buffers.VertexTextureCoords = uint32(gl.GetAttribLocation(program, gl.Str("vertexTextureCoords\x00")))

  // Store buffers
  BUFFERS[object.Name] = buffers
}
