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
  "math"
  "io/ioutil"
  "encoding/json"

  "github.com/go-gl/gl/v4.1-core/gl"
  "github.com/go-gl/mathgl/mgl32"
)

type Object struct {
  Name        string

  Radius      float32
  Inclination float32
  Revolution  float32
  Rotation    float32
  Distance    float32

  Objects     []Object
}

func loadObjects(map_name string) (*[]Object) {
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

  return &objects_map
}

func renderObjects(objects *[]Object, program uint32) {
  for o := range *objects {
    buffers := getBuffers((*objects)[o].Name)

    gl.BindTexture(gl.TEXTURE_2D, buffers.Texture.Ref)

    // Toggle to child context
    pushMatrix()

    // Update angles for object
    buffers.addToAngleRotation(rotationAngleSinceLast(&(*objects)[o]))
    buffers.addToAngleRevolution(revolutionAngleSinceLast(&(*objects)[o]))

    // Apply model transforms
    if (*objects)[o].Revolution != 0 {
      CURRENT_MATRIX = CURRENT_MATRIX.Mul4(mgl32.HomogRotate3D(buffers.AngleRevolution, mgl32.Vec3{0, 1, 0}))
    }

    if (*objects)[o].Distance > 0 {
      CURRENT_MATRIX = CURRENT_MATRIX.Mul4(mgl32.Translate3D(normalizeObjectDistance((*objects)[o].Distance), 0.0, 0.0))
    }

    // Toggle to unary context
    pushMatrix()

    // Apply object angles
    if (*objects)[o].Rotation != 0 {
      CURRENT_MATRIX = CURRENT_MATRIX.Mul4(mgl32.HomogRotate3D(buffers.AngleRotation, mgl32.Vec3{0, 1, 0}))
    }

    if (*objects)[o].Inclination > 0 {
      CURRENT_MATRIX = CURRENT_MATRIX.Mul4(mgl32.HomogRotate3D((*objects)[o].Inclination / 90.0, mgl32.Vec3{0, 0, 1}))
    }

    // Apply model
    gl.UniformMatrix4fv(MODEL_UNIFORM, 1, false, &CURRENT_MATRIX[0])

    // Render vertices
    gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereVertices)
    gl.VertexAttribPointer(SHADER_VERTEX_ATTRIBUTES, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

    // Render textures
    gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOSphereTexture)
    gl.VertexAttribPointer(SHADER_VERTEX_TEXTURE_COORDS, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

    // Render indices
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers.VBOSphereIndices)

    setMatrixUniforms(program)

    // Draw elements
    gl.DrawElements(gl.TRIANGLES, int32(len(buffers.Sphere.Indices) * 2), gl.UNSIGNED_INT, gl.PtrOffset(0))

    // Reset buffers
    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)

    // Toggle back from unary context
    popMatrix()

    // Render children (if any?)
    renderObjects(&((*objects)[o].Objects), program)

    // Toggle back to parent context
    popMatrix()
  }
}

func normalizeObjectRadius(radius float32) (float32) {
  return float32(math.Cbrt(float64(radius)) * OBJECT_FACTOR_RADIUS)
}

func normalizeObjectDistance(distance float32) (float32) {
  return float32(math.Sqrt(float64(distance)) * OBJECT_FACTOR_DISTANCE)
}
