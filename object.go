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
  Type        string

  Radius      float32
  Inclination float32
  Revolution  float32
  Rotation    float32
  Distance    float32
  Center      bool
  Radiate     bool
  Cosmic      bool

  Objects     []Object
}

type ObjectElement struct {
  Vertices       []float32
  VerticeNormals []float32
  Indices        []int32
  TextureCoords  []float32
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
  // Acquire shader
  shader := getShader()
  light := getLight()
  matrix_uniforms := getMatrixUniforms()

  // Iterate on current-level objects
  for o := range *objects {
    buffers := getBuffers((*objects)[o].Name)

    gl.BindTexture(gl.TEXTURE_2D, buffers.Texture.Ref)

    // Toggle to child context
    pushMatrix()

    // Update angles for object
    buffers.addToAngleRotation(rotationAngleSinceLast(&(*objects)[o]))
    buffers.addToAngleRevolution(revolutionAngleSinceLast(&(*objects)[o]))

    // Apply model transforms
    current_matrix_shared := getMatrix()

    if (*objects)[o].Revolution != 0 {
      *current_matrix_shared = (*current_matrix_shared).Mul4(mgl32.HomogRotate3D(buffers.AngleRevolution, mgl32.Vec3{0, 1, 0}))
    }

    if (*objects)[o].Distance > 0 && (*objects)[o].Center != true {
      *current_matrix_shared = (*current_matrix_shared).Mul4(mgl32.Translate3D(normalizeObjectSize((*objects)[o].Distance), 0.0, 0.0))
    }

    setMatrix(current_matrix_shared)

    // Toggle to unary context
    pushMatrix()

    // Apply object angles
    current_matrix_self := getMatrix()

    if (*objects)[o].Rotation != 0 {
      *current_matrix_self = (*current_matrix_self).Mul4(mgl32.HomogRotate3D(buffers.AngleRotation, mgl32.Vec3{0, 1, 0}))
    }

    if (*objects)[o].Inclination > 0 {
      *current_matrix_self = (*current_matrix_self).Mul4(mgl32.HomogRotate3D((*objects)[o].Inclination / 90.0, mgl32.Vec3{0, 0, 1}))
    }

    setMatrix(current_matrix_shared)

    // Process normal to model matrix
    normal_matrix := mgl32.Mat4Normal(*current_matrix_self)

    // Apply model + normal
    gl.UniformMatrix4fv(matrix_uniforms.Model, 1, false, &((*current_matrix_self)[0]))
    gl.UniformMatrix3fv(matrix_uniforms.Normal, 1, false, &normal_matrix[0])

    // Render vertices
    gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOElementVertices)
    gl.VertexAttribPointer(shader.VertexAttributes, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

    // Render textures
    gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOElementTexture)
    gl.VertexAttribPointer(shader.VertexTextureCoords, 2, gl.FLOAT, false, 0, gl.PtrOffset(0))

    // Render vertice lightings
    gl.BindBuffer(gl.ARRAY_BUFFER, buffers.VBOElementVerticeNormals)
    gl.VertexAttribPointer(shader.NormalAttributes, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

    // Light emitter? (eg: Sun)
    if (*objects)[o].Radiate == true {
      gl.Uniform1i(light.IsLightEmitterUniform, 1)

      gl.Uniform3f(light.PointLightingLocationUniform, 0, 0, 0);
      gl.Uniform3f(light.PointLightingColorUniform, 1, 1, 1);
    }

    // Light receiver? (eg: planet, moon)
    if (*objects)[o].Cosmic == true {
      // It is a far-away cosmic object, dont light it from emitter
      gl.Uniform1i(light.IsLightReceiverUniform, 0)
    } else {
      gl.Uniform1i(light.IsLightReceiverUniform, 1)
    }

    // Render indices
    gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers.VBOElementIndices)

    // Draw elements
    gl.DrawElements(getObjectDrawMode(&(*objects)[o]), int32(len(buffers.Element.Indices) * 2), gl.UNSIGNED_INT, gl.PtrOffset(0))

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

func normalizeObjectSize(size float32) (float32) {
  return float32(math.Sqrt(float64(size)) * OBJECT_FACTOR_SIZE)
}

func getObjectDrawMode(object *Object) (uint32) {
  if object.Type == "circle" {
    return gl.LINES
  }

  return gl.TRIANGLES
}
