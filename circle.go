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
  "math"
)

// Circle  Maps a circle object
type Circle ObjectElement

func generateCircleFilledFromObject(object *Object) (Circle) {
  return generateCircle(object.Distance, object.Radius, object.Radiate)
}

func generateCircleFromObject(object *Object) (Circle) {
  return generateCircle(object.Radius, 0.0, object.Radiate)
}

func generateCircle(radius float32, thickness float32, radiate bool) (Circle) {
  var (
    circle                 Circle

    i                      int
    j                      int
    k                      int
    l                      int

    accumulatorMainSize    int32
    accumulatorIndicesSize int32
    normalDirection        float32

    radiusInsideN          float32
    radiusOutsideN         float32
    nbVertices             int32

    angle                  float64
    angleMax               int32
  )

  angleMax = 360

  if thickness > 0 {
    accumulatorMainSize = 2
    accumulatorIndicesSize = 6
  } else {
    accumulatorMainSize = 1
    accumulatorIndicesSize = 2
  }

  circle.Vertices = make([]float32, 3 * (angleMax + 1) * accumulatorMainSize)
  circle.VerticeNormals = make([]float32, 3 * (angleMax + 1) * accumulatorMainSize)
  circle.Indices = make([]int32, (angleMax + 1) * accumulatorIndicesSize)
  circle.TextureCoords = make([]float32, 2 * (angleMax + 1) * accumulatorMainSize)

  i = 0
  j = 0
  k = 0
  l = 0

  nbVertices = 1

  radiusInsideN = normalizeObjectSize(radius)
  radiusOutsideN = normalizeObjectSize(radius + thickness)

  // Normal is -1 if sun, which is the light source, to avoid any self-shadow effect
  if radiate == true {
    normalDirection = -1.0
  } else {
    normalDirection = 1.0
  }

  for angle = 0.0; angle <= float64(angleMax); angle++ {
    // Generate inner circle object
    generateCircleObject(&circle, radiusInsideN, thickness, angle, angleMax, normalDirection, nbVertices, 0, &i, &j, &k, &l)

    if thickness > 0.0 {
      // Generate outer circle object? (if not last)
      generateCircleObject(&circle, radiusOutsideN, thickness, angle, angleMax, normalDirection, nbVertices, 1, &i, &j, &k, &l)

      nbVertices++
    }

    nbVertices++
  }

  return circle
}

func generateCircleObject(circle *Circle, radiusN float32, thickness float32, angle float64, angleMax int32, normalDirection float32, nbVertices int32, passIndex int32, i *int, j *int, k *int, l *int) {
  var (
    vertexPositionX float32
    vertexPositionY float32
    vertexPositionZ float32
  )

  // Generate inside circle vertices
  vertexPositionX = float32(math.Cos(ConfigMathDegreeToRadian * angle))
  vertexPositionY = 0.0
  vertexPositionZ = float32(math.Sin(ConfigMathDegreeToRadian * angle))

  // Bind inside circle vertices
  circle.Vertices[*i] = radiusN * vertexPositionX
  circle.Vertices[*i + 1] = radiusN * vertexPositionY
  circle.Vertices[*i + 2] = radiusN * vertexPositionZ

  *i += 3

  // Bind circle vertice normals
  circle.VerticeNormals[*j] = normalDirection * vertexPositionX
  circle.VerticeNormals[*j + 1] = normalDirection * vertexPositionY
  circle.VerticeNormals[*j + 2] = normalDirection * vertexPositionZ

  *j += 3

  // Bind circle indices
  if thickness > 0.0 {
    circle.Indices[*k] = (nbVertices % (angleMax * 2)) + 1
    circle.Indices[*k + 1] = ((nbVertices + 1 + passIndex) % (angleMax * 2)) + 1
    circle.Indices[*k + 2] = ((nbVertices + 3) % (angleMax * 2)) + 1

    *k += 3
  } else {
    circle.Indices[*k] = ((nbVertices) % angleMax) + 1
    circle.Indices[*k + 1] = ((nbVertices + 1) % angleMax) + 1

    *k += 2
  }

  // Bind circle texture coordinates
  circle.TextureCoords[*l] = float32(passIndex)
  circle.TextureCoords[*l + 1] = 0.0

  *l += 2
}
