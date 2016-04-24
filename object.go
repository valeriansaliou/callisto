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
  "fmt"
  "io/ioutil"
  "encoding/json"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type Object struct {
  Name       string
  Radius     float32
  Angle      float32
  Velocity   float32
  Revolution float32

  Objects    []Object
}

func loadObjects(map_name string) ([]Object) {
  var objects_map []Object

  // Load objects map
  filePath := fmt.Sprintf("maps/%s.json", map_name)

  file, err := ioutil.ReadFile(filePath)
  if err != nil {
    panic(err)
  }

  // Transform JSON map into Go map
  err = json.Unmarshal(file, &objects_map)

  if err != nil {
    panic(err)
  }

  return objects_map
}

func renderObjects(objects []Object, angle float32) {
  for o := range objects {
    buffers := getBuffers(objects[o].Name)

    // Process model
    buffers.Model = mgl32.HomogRotate3D(angle, mgl32.Vec3{0, 1, 0})
    gl.UniformMatrix4fv(buffers.ModelUniform, 1, false, &buffers.Model[0])

    // Render vertices
    gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereVertices)
    gl.EnableVertexAttribArray(buffers.VertexAttributes)
    gl.VertexAttribPointer(buffers.VertexAttributes, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

    // Render textures
    gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereTexture)
    gl.EnableVertexAttribArray(buffers.VertexTextureCoords)
    gl.VertexAttribPointer(buffers.VertexTextureCoords, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

    // Render indices
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers.VBOSphereIndices)

    gl.BindTexture(gl.TEXTURE_2D, buffers.Texture)

    // Draw elements
    gl.DrawElements(gl.TRIANGLES, int32(len(buffers.Sphere.Indices) * 2), gl.UNSIGNED_INT, gl.PtrOffset(0))

    // Reset buffers
    gl.DisableVertexAttribArray(buffers.VertexAttributes)
    gl.DisableVertexAttribArray(buffers.VertexTextureCoords)
    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
  }
}
