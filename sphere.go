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

// Sphere  Maps a sphere object
type Sphere ObjectElement

func generateSphereFromObject(object *Object) (Sphere) {
  return generateSphere(object.Radius, object.Radiate)
}

func generateSphere(radius float32, radiate bool) (Sphere) {
  var (
    sphere          Sphere

    i               int
    j               int
    k               int
    l               int

    normalDirection float32

    radiusN         float32
    nbVertices      int32

    unarySizeFull   int
    unarySizeShort  int

    resLongitude    int32

    longitude       int
    latitude        int

    longitudeRF     float64
    latitudeRF      float64

    vertexPositionX float32
    vertexPositionY float32
    vertexPositionZ float32
  )

  unarySizeFull = (2 * ConfigObjectTexturePhiMax / ConfigObjectTextureStepLatitude + 1) * (ConfigObjectTextureThetaMax / ConfigObjectTextureStepLongitude + 1)
  unarySizeShort = (2 * ConfigObjectTexturePhiMax / ConfigObjectTextureStepLatitude) * (ConfigObjectTextureThetaMax / ConfigObjectTextureStepLongitude)

  sphere.Vertices = make([]float32, 3 * unarySizeFull)
  sphere.VerticeNormals = make([]float32, 3 * unarySizeFull)
  sphere.Indices = make([]int32, 6 * unarySizeShort)
  sphere.TextureCoords = make([]float32, 2 * unarySizeFull)

  i = 0
  j = 0
  k = 0
  l = 0

  radiusN = normalizeObjectSize(radius)

  nbVertices = 0
  resLongitude = int32(float32(ConfigObjectTextureThetaMax) / float32(ConfigObjectTextureStepLongitude) + 1.0);

  // Normal is -1 if sun, which is the light source, to avoid any self-shadow effect
  if radiate == true {
    normalDirection = -1.0
  } else {
    normalDirection = 1.0
  }

  // Map sphere data
  for latitude = -90; latitude <= ConfigObjectTexturePhiMax; latitude += ConfigObjectTextureStepLatitude {
    for longitude = 0; longitude <= ConfigObjectTextureThetaMax; longitude += ConfigObjectTextureStepLongitude {
      // Convert latitude & longitude to radians
      longitudeRF = float64(ConfigMathDegreeToRadian) * float64(longitude)
      latitudeRF = float64(ConfigMathDegreeToRadian) * float64(latitude)

      // Process vertex positions
      vertexPositionX = float32(math.Sin(longitudeRF) * math.Cos(latitudeRF))
      vertexPositionY = float32(math.Sin(latitudeRF))
      vertexPositionZ = float32(math.Cos(latitudeRF) * math.Cos(longitudeRF))

      // Bind sphere vertices
      sphere.Vertices[i] = radiusN * vertexPositionX
      sphere.Vertices[i + 1] = radiusN * vertexPositionY
      sphere.Vertices[i + 2] = radiusN * vertexPositionZ

      i += 3

      // Bind sphere vertice normals
      sphere.VerticeNormals[j] = normalDirection * vertexPositionX
      sphere.VerticeNormals[j + 1] = normalDirection * vertexPositionY
      sphere.VerticeNormals[j + 2] = normalDirection * vertexPositionZ

      j += 3

      // Bind sphere indices
      if longitude != ConfigObjectTextureThetaMax && latitude < ConfigObjectTexturePhiMax {
        sphere.Indices[k] = nbVertices
        sphere.Indices[k + 1] = nbVertices + 1
        sphere.Indices[k + 2] = nbVertices + 1 + resLongitude

        sphere.Indices[k + 3] = nbVertices
        sphere.Indices[k + 4] = nbVertices + 1 + resLongitude
        sphere.Indices[k + 5] = nbVertices + resLongitude

        k += 6
      }

      nbVertices += 1

      // Bind sphere texture coordinates
      sphere.TextureCoords[l] = float32(longitude) / float32(ConfigObjectTextureThetaMax)
      sphere.TextureCoords[l + 1] = -1.0 * float32(90.0 + latitude) / float32(90.0 + ConfigObjectTexturePhiMax)

      l += 2
    }
  }

  return sphere
}
