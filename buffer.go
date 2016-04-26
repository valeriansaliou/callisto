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
)

type Buffers struct {
  Sphere                  Sphere
  Texture                 Texture

  AngleRotation           float32
  AngleRevolution         float32

  VBOSphereVertices       uint32
  VBOSphereVerticeNormals uint32
  VBOSphereTexture        uint32
  VBOSphereIndices        uint32
}

var __BUFFERS map[string]*Buffers = make(map[string]*Buffers)

func (buffers *Buffers) addToAngleRotation(angle float32) {
  buffers.AngleRotation += angle
}

func (buffers *Buffers) addToAngleRevolution(angle float32) {
  buffers.AngleRevolution += angle
}

func getBuffers(name string) (*Buffers) {
  return __BUFFERS[name]
}

func setBuffers(name string, buffers *Buffers) {
  __BUFFERS[name] = buffers
}

func createAllBuffers(objects *[]Object, program uint32, vao uint32) {
  // Object buffers
  for o := range *objects {
    // Create the object buffers
    createBuffers(&(*objects)[o], program, vao)
  }
}

func createBuffers(object *Object, program uint32, vao uint32) {
  var (
    err error
  )

  buffers := &Buffers{}

  // Zero angle
  buffers.AngleRotation = 0.0
  buffers.AngleRevolution = 0.0

  // Generate sphere
  buffers.Sphere = generateSphere(object)

  // Load texture
  buffers.Texture, err = loadTexture(object.Name)

  if err != nil {
    log.Fatalln(err)
  }

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

  gl.GenBuffers(1, &buffers.VBOSphereVerticeNormals)
  gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereVerticeNormals)
  gl.BufferData(gl.ARRAY_BUFFER, len(buffers.Sphere.VerticeNormals)*4, gl.Ptr(buffers.Sphere.VerticeNormals), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &buffers.VBOSphereTexture)
  gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereTexture)
  gl.BufferData(gl.ARRAY_BUFFER, len(buffers.Sphere.TextureCoords)*4, gl.Ptr(buffers.Sphere.TextureCoords), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &buffers.VBOSphereIndices)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers.VBOSphereIndices)
  gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(buffers.Sphere.Indices)*4, gl.Ptr(buffers.Sphere.Indices), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

  // Go deeper (if any child)
  for o := range (*object).Objects {
    createBuffers(&((*object).Objects[o]), program, vao)
  }

  // Store buffers
  // Notice: if a lower level buffer is set w/ the same name, the higher-level object will always override it
  setBuffers(object.Name, buffers)
}
