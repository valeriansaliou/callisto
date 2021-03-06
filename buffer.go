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

// Buffers  Maps buffered object data (built once, accessed at every render)
type Buffers struct {
  Element                  ObjectElement
  Texture                  Texture

  AngleRotation            float32
  AngleRevolution          float32
  AngleTilt                float32

  VBOElementVertices       uint32
  VBOElementVerticeNormals uint32
  VBOElementTexture        uint32
  VBOElementIndices        uint32
}

// InstanceBuffers  Stores buffered object data (built once, accessed at every render)
var InstanceBuffers = make(map[string]*Buffers)

func (buffers *Buffers) addToAngleRotation(angle float32) {
  buffers.AngleRotation += angle
}

func (buffers *Buffers) addToAngleRevolution(angle float32) {
  buffers.AngleRevolution += angle
}

func getBuffers(name string) (*Buffers) {
  return InstanceBuffers[name]
}

func setBuffers(name string, buffers *Buffers) {
  InstanceBuffers[name] = buffers
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
  buffers.AngleRotation = rotationAngleSinceStart(object)
  buffers.AngleRevolution = revolutionAngleSinceStart(object)
  buffers.AngleTilt = object.Tilt * float32(ConfigMathDegreeToRadian)

  // Generate object
  switch object.Type {
    case "sphere":
      buffers.Element = ObjectElement(generateSphereFromObject(object))

    case "circle":
      buffers.Element = ObjectElement(generateCircleFromObject(object))

    case "circle-filled":
      buffers.Element = ObjectElement(generateCircleFilledFromObject(object))

    default:
      panic("Object type not supported")
  }

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
  gl.GenBuffers(1, &buffers.VBOElementVertices)
  gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOElementVertices)
  gl.BufferData(gl.ARRAY_BUFFER, len(buffers.Element.Vertices)*4, gl.Ptr(buffers.Element.Vertices), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &buffers.VBOElementVerticeNormals)
  gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOElementVerticeNormals)
  gl.BufferData(gl.ARRAY_BUFFER, len(buffers.Element.VerticeNormals)*4, gl.Ptr(buffers.Element.VerticeNormals), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &buffers.VBOElementTexture)
  gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOElementTexture)
  gl.BufferData(gl.ARRAY_BUFFER, len(buffers.Element.TextureCoords)*4, gl.Ptr(buffers.Element.TextureCoords), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ARRAY_BUFFER, 0)

  gl.GenBuffers(1, &buffers.VBOElementIndices)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers.VBOElementIndices)
  gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(buffers.Element.Indices)*4, gl.Ptr(buffers.Element.Indices), gl.STATIC_DRAW)
  gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

  // Go deeper (if any child)
  for o := range (*object).Objects {
    createBuffers(&((*object).Objects[o]), program, vao)
  }

  // Store buffers
  // Notice: if a lower level buffer is set w/ the same name, the higher-level object will always override it
  setBuffers(object.Name, buffers)
}
